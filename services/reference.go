package services

import (
	"encoding/gob"
	"fmt"
	"os"
)

var ReferenceMap map[string]string

func saveReferenceFile() {
	ReferenceMap = make(map[string]string)
	ReferenceMap["BFR"] = "Body Fat Rate"
	ReferenceMap["DoE"] = "Date of Exercise"
	ReferenceMap["DP"] = "Diastolic Pressure"
	ReferenceMap["E3"] = "Exercise 3"
	ReferenceMap["EHR"] = "Exercise Heart Rate"
	ReferenceMap["FBG"] = "Fasting Blood Glucose"
	ReferenceMap["FR"] = "Fat Ratio"
	ReferenceMap["PBG"] = "Postprandial Blood Glucose"
	ReferenceMap["RHR"] = "Resting Heart Rate"
	ReferenceMap["SEX"] = "Sex"
	ReferenceMap["SP"] = "Systolic Pressure"
	ReferenceMap["ST"] = "Step Count"
	ReferenceMap["W"] = "Weight"
	file, err := os.Create("reference.bin")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	if err := encoder.Encode(ReferenceMap); err != nil {
		fmt.Println("Error encoding data:", err)
		return
	}
	fmt.Println("Data has been saved to data.bin")
}

func LoadReferenceFile() {
	file, err := os.Open("reference.bin")
	if err != nil {
		fmt.Println("Error opening file:", err)
		saveReferenceFile()
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	if err := decoder.Decode(&ReferenceMap); err != nil {
		fmt.Println("Error decoding data:", err)
		return
	}

}
