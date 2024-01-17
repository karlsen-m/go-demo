package main

import (
	"errors"
	"fmt"
	"github.com/tealeg/xlsx"
	"regexp"
	"strings"
)

var successData [][]string
var failData [][]string

func main() {
	fileName := "demo.xlsx"
	startRowNum := 1
	endRowNum := 115216
	GetFileData(fileName, startRowNum, endRowNum)
}

func GetFileData(fileName string, startRowNum int, endRowNum int) (data [][]string) {
	var err error
	jobData := &WhitelistImportData{}
	xlFile, err := xlsx.OpenFile(fileName)
	if err != nil {
		fmt.Println("文件名包含非法字符，请修改后重新上传")
	} else {
		for _, sheet := range xlFile.Sheets {
			rowNum := 0
			for _, row := range sheet.Rows {
				rowNum++
				fmt.Println("runNum:", rowNum)
				if rowNum >= startRowNum {
					strs := make([]string, 5)
					for k, cell := range row.Cells {
						text := cell.String()
						text = strings.TrimSpace(text)
						strs[k] = text
					}
					if rowNum == 1 {
						failData = append(failData, strs)
						successData = append(successData, strs)
						continue
					}
					jobData = &WhitelistImportData{
						Name:          strs[0],
						Mobile:        strs[1],
						IdCard:        strs[2],
						Remarks:       strs[3],
						FounderMobile: strs[4],
					}
					err = selectData(jobData)
					if err != nil {
						failData = append(failData, strs)
					} else {
						successData = append(successData, strs)
					}
				}
				if rowNum >= endRowNum {
					break
				}
			}
		}
	}
	AddWriteToXlsx(successData, "success1.xlsx")
	AddWriteToXlsx(failData, "fail1.xlsx")
	return
}

func AddWriteToXlsx(data [][]string, fileName string) {
	var setFile *xlsx.File
	var setSheet *xlsx.Sheet
	var setRow *xlsx.Row
	var setCell *xlsx.Cell
	var err error
	setFile = xlsx.NewFile()
	setSheet, err = setFile.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}
	for _, v := range data {
		setRow = setSheet.AddRow()
		for _, strV := range v {
			setCell = setRow.AddCell()
			setCell.Value = strV
		}
	}
	err = setFile.Save(fileName)
	if err != nil {
		fmt.Printf(err.Error())
	}
}

type WhitelistImportData struct {
	Name          string
	Mobile        string
	FounderMobile string
	IdCard        string
	Remarks       string
}

func selectData(hitelistImportData *WhitelistImportData) (err error) {
	if hitelistImportData.Name == "" && hitelistImportData.Mobile == "" && hitelistImportData.FounderMobile == "" && hitelistImportData.IdCard == "" && hitelistImportData.Remarks == "" {
		return nil
	}
	if hitelistImportData.Name == "" || hitelistImportData.Mobile == "" || hitelistImportData.FounderMobile == "" || hitelistImportData.IdCard == "" || hitelistImportData.Remarks == "" {
		err = errors.New("姓名手机号创建人手机号数据都不能为空")
		return err
	}
	//姓名
	r, _ := regexp.Compile(`^[^\da-zA-Z]+$`)
	if !r.MatchString(hitelistImportData.Name) {
		err = errors.New("姓名不合法")
		return err
	}
	//手机号
	r, _ = regexp.Compile(`^[1][3-9][0-9]{9}$`)
	if !r.MatchString(hitelistImportData.Mobile) {
		err = errors.New("手机号不合法")
		return err
	}
	//身份证的
	r, _ = regexp.Compile("^[1-9][0-9]{5}[1-2][0-9]{3}[0-1][0-9][0-3][0-9][0-9]{3}[0-9X]$")
	if !r.MatchString(hitelistImportData.IdCard) {
		//粤港澳通行证
		r, _ = regexp.Compile("^[HMhm][0-9]{10}$")
		r2, _ := regexp.Compile("^[HMhm][0-9]{8}$")
		if !r.MatchString(hitelistImportData.IdCard) && !r2.MatchString(hitelistImportData.IdCard) {
			//台湾通行证
			r, _ = regexp.Compile("^[A-Z][0-9A-Z][0-9A-Z]{6}$")
			if !r.MatchString(hitelistImportData.IdCard) {
				err = errors.New("证件号格式不正确")
				return err
			}
		}
	}
	return nil
}
