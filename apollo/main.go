package main

import (
	"bytes"
	"fmt"
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"time"
)

func main() {
	c := &config.AppConfig{
		AppID:          "SampleApp",
		Cluster:        "dev",
		IP:             "http://localhost:8080",
		NamespaceName:  "appjson.json",
		IsBackupConfig: true,
		BackupConfigPath: "./app.json",
		Secret:         "39428d39f5d942888856bb80e59ebea6",
		MustStart:      true,
	}
	notifies := make([]*config.Notification, 0)
	notifies = append(notifies, &config.Notification{
		NamespaceName: "appjson.json",
		NotificationID: -1,
	})
	c.Init()
	c.GetNotificationsMap().UpdateAllNotifications(notifies)
	client, _ := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})
	fmt.Println("初始化Apollo配置成功")
	//Use your apollo key to test
	cache := client.GetConfigCache(c.NamespaceName)
	value, _ := cache.Get("content")
	go func() {
		var oldValue string
		oldValue = value.(string)
		for  {
			time.Sleep(time.Second * 1)
			value, _ = cache.Get("content")
			if oldValue != value.(string) {
				Init([]byte(value.(string)))
			}
		}
	}()
	configData := []byte(value.(string))
	viper.SetConfigType("json")
	if err := viper.ReadConfig(bytes.NewReader(configData)); err != nil {
		log.Fatal("加载配置文件失败;信息:" + err.Error())
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件变更", e.Name)
	})
	num := 0
	for  {
		num++
		fmt.Println(fmt.Sprintf("第%d次,获取到的配置为:%s", num, viper.GetString("Name")))
		time.Sleep(5 * time.Second)
	}
}

func Init(configData []byte) {
	viper.SetConfigType("json")
	if err := viper.ReadConfig(bytes.NewReader(configData)); err != nil {
		log.Fatal("加载配置文件失败;信息:" + err.Error())
	}
}