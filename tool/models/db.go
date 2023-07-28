package db

import (
	"context"
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"math/rand"
	"time"
)

const (
	DSN                = `%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Asia%%2FShanghai`
	ConfigTypeCluster  = `cluster`
	ClusterTypeSource  = `source`  // 集群的类型主库
	ClusterTypeReplica = `replica` // 集群的类型从库
)

// Config 数据库配置结构体
type Config struct {
	Type            string    `json:"type"`
	Username        string    `json:"username"`
	Host            string    `json:"host"`
	Port            int       `json:"port"`
	Password        string    `json:"password"`
	Database        string    `json:"database"`
	Prefix          string    `json:"prefix"` // 表名前缀
	ConnMaxLifetime int       `json:"connMaxLifetime"`
	ClusterType     string    `json:"clusterType"`
	ClusterConfigs  []*Config `json:"clusterConfigs"`
}

type Client struct {
	db     *gorm.DB
	config *Config
	driver string
}

var clients map[string]*Client

var defaultDriver string

func InitWithEnv() (err error) {
	clients = make(map[string]*Client)
	// 默认驱动
	defaultDriver = "default"

	// 生成所有连接
	dc := &Config{
		Type:            "mysql",
		Username:        "root",
		Host:            "127.0.0.1",
		Port:            3306,
		Password:        "123456",
		Database:        "test",
		Prefix:          "",
		ConnMaxLifetime: 3600,
	}
	db, _err := ConnectMysqlWithGorm(dc)
	if _err != nil {
		return _err
	}
	clients[defaultDriver] = &Client{
		db:     db,
		config: dc,
		driver: defaultDriver,
	}

	return nil
}

func GetDefaultClient() *Client {
	return GetClient(defaultDriver)
}
func GetClient(driver string) *Client {
	if client, ok := clients[driver]; ok {
		return client
	} else {
		panic("client not found")
	}
}
func GetClients() map[string]*Client {
	return clients
}

func (c *Client) GetDB() *gorm.DB {
	return c.db
}
func (c *Client) GetDBSession(ctx context.Context) *gorm.DB {
	return c.GetDB().WithContext(ctx)
}
func (c *Client) GetConfig() *Config {
	return c.config
}
func (c *Client) GetDriver() string {
	return c.driver
}

func NewGormDialector(c *Config) gorm.Dialector {
	return mysql.New(mysql.Config{
		DSN:                      fmt.Sprintf(DSN, c.Username, c.Password, c.Host, c.Port, c.Database), // DSN data source name
		DisableDatetimePrecision: true,                                                                 // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:   true,                                                                 // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		// DontSupportRenameColumn: true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	})
}

// ConnectMysqlWithGorm 获取数据库连接
func ConnectMysqlWithGorm(tmpConf *Config) (db *gorm.DB, err error) {
	db, err = gorm.Open(NewGormDialector(tmpConf), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: tmpConf.Prefix,
		},
	})
	if err != nil {
		return
	}
	if tmpConf.Type == ConfigTypeCluster {
		// 集群添加配置
		clusterConfigs := tmpConf.ClusterConfigs
		sds := make([]gorm.Dialector, 0)
		rds := make([]gorm.Dialector, 0)
		sds = append(sds, NewGormDialector(tmpConf))
		for _, clusterConfig := range clusterConfigs {
			if clusterConfig.ClusterType == ClusterTypeSource {
				// 主库
				sds = append(sds, NewGormDialector(clusterConfig))
			} else {
				// 从库
				rds = append(rds, NewGormDialector(clusterConfig))
			}
		}
		_ = db.Use(dbresolver.Register(dbresolver.Config{
			Sources:  sds,
			Replicas: rds,
			Policy:   ClusterPolicy{},
		}).
			SetMaxIdleConns(10).
			SetMaxOpenConns(1000).
			SetConnMaxLifetime(time.Duration(tmpConf.ConnMaxLifetime) * time.Second))
	} else {
		// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
		sqlDB, _err := db.DB()
		if _err != nil {
			err = _err
			return
		}
		// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
		sqlDB.SetMaxIdleConns(10)
		// SetMaxOpenConns 设置打开数据库连接的最大数量。
		sqlDB.SetMaxOpenConns(1000)
		// SetConnMaxLifetime 设置了连接可复用的最大时间。
		sqlDB.SetConnMaxLifetime(time.Duration(tmpConf.ConnMaxLifetime) * time.Second)
	}
	// 添加链路跟踪插件
	err = db.Use(&OpentracingPlugin{})
	if err != nil {
		return
	}
	return
}

type ClusterPolicy struct {
}

func (p ClusterPolicy) Resolve(connPools []gorm.ConnPool) gorm.ConnPool {
	index := rand.Intn(len(connPools))
	tmp := connPools[index].(*sql.DB)
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer func() {
		cancel()
	}()
	err := tmp.PingContext(ctx)
	if err != nil {
		return p.Resolve(connPools)
	}
	return tmp
}
