package services

import (
	"encoding/gob"
	"fmt"
	"oldsouth/models"
	"os"
	"time"
)

var MetaDatas []models.MetaData

func SaveData(code string, float float32) {
	newData := models.MetaData{
		Id:        len(MetaDatas) + 1,
		MetaCode:  code,
		MetaFloat: float,
		MetaTime:  time.Now(),
	}
	MetaDatas = append(MetaDatas, newData)
	saveData()
	fmt.Println(MetaDatas)
}

func SaveDataByDate(code string, float float32, time time.Time) {
	newData := models.MetaData{
		Id:        len(MetaDatas) + 1,
		MetaCode:  code,
		MetaFloat: float,
		MetaTime:  time,
	}
	MetaDatas = append(MetaDatas, newData)
	saveData()
	fmt.Println(MetaDatas)
}

func saveData() {
	file, err := os.Create("data.bin")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	if err := encoder.Encode(MetaDatas); err != nil {
		fmt.Println("Error encoding data:", err)
		return
	}
	fmt.Println("Data has been saved to data.bin")
}

func LoadData() {
	file, err := os.Open("data.bin")
	if err != nil {
		fmt.Println("Error opening file:", err)
		saveData()
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	if err := decoder.Decode(&MetaDatas); err != nil {
		fmt.Println("Error decoding data:", err)
		return
	}
}

func GetData(metaCode string) models.MetaDatas {
	now := time.Now()
	var tmpList []models.MetaData
	var timeList []string
	var dataList []float32
	var referenceList []string
	for _, metaData := range MetaDatas {
		if metaCode == metaData.MetaCode && metaData.MetaTime.After(now.AddDate(0, 0, -7)) && metaData.MetaTime.Before(now) {
			fmt.Println(metaData)
			tmpList = append(tmpList, metaData)
		}
	}
	for _, dataInfo := range tmpList {
		timeList = append(timeList, dataInfo.MetaTime.Format("20060102-15"))
		dataList = append(dataList, dataInfo.MetaFloat)
		referenceList = append(referenceList, dataInfo.MetaCode)
	}
	result := models.MetaDatas{
		TimeList:      timeList,
		DataList:      dataList,
		ReferenceList: referenceList,
		Title:         metaCode,
	}
	return result
}
