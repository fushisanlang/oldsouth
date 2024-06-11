package models

import "time"

type MetaData struct {
	Id        int       `json:"Id"`
	MetaCode  string    `json:"MetaCode"`
	MetaFloat float32   `json:"MetaFloat"`
	MetaTime  time.Time `json:"MetaTime"`
}
type MetaDatas struct {
	DataList      []float32
	ReferenceList []string
	TimeList      []string

	Title string
}
