package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/spf13/cast"
)

var (
	totalNum      = 496216
	runTotalNum   = 0
	yfMerchantNos = []string{}
	pageTotalSize = 200
	pageSize      = 0
)

func main() {
	excelFileName := "./shujuliu.xlsx"

	// 打开Excel文件
	f, _err := excelize.OpenFile(excelFileName)
	if _err != nil {
		fmt.Printf("无法打开Excel文件: %s\n", _err)
		return
	}

	// 遍历每个Sheet
	for _, sheet := range f.GetSheetMap() {
		fmt.Printf("Sheet名称: %s\n", sheet)

		// 获取Sheet的数据流
		r, err := f.Rows(sheet)
		if err != nil {
			fmt.Printf("获取Sheet数据流错误: %s\n", err)
			return
		}

		// 处理数据流中的行数据
		err = processRows(r)
		if err != nil {
			fmt.Printf("处理行数据错误: %s\n", err)
			return
		}
	}
}

func processRows(rows *excelize.Rows) error {
	for rows.Next() {
		runTotalNum++
		if runTotalNum == 1 {
			continue
		}
		if runTotalNum > totalNum {
			break
		}
		// 读取当前行的单元格数据
		cells := rows.Columns()
		pageSize++
		yfMerchantNos = append(yfMerchantNos, cells[0])
		fmt.Println("runNum：" + cast.ToString(runTotalNum) + "，yfMerchantNos：|" + cells[0] + "|")
		if pageSize < pageTotalSize && runTotalNum < totalNum {
			continue
		}
		//查询数据库

	}

	if err := rows.Error(); err != nil {
		return err
	}

	return nil
}
