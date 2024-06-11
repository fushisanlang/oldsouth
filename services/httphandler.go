package services

import (
	"encoding/json"
	"net/http"
	"oldsouth/models"
	"time"
)

func checkToken(token string) bool {
	if token == "2Ux~BOIBCDU+6!hgDCp8" {
		return true
	} else {
		return false
	}
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {

	response := models.HelloResponse{
		Message: "Healthy me and you!",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func MetasHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		token := r.Header.Get("token")
		metaCode := r.Header.Get("metacode")
		if checkToken(token) {
			_, ok := ReferenceMap[metaCode]
			if ok {
				response := GetData(metaCode)

				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(response)
			} else {
				response := models.HelloResponse{
					Message: "err code",
				}

				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(response)
			}
		} else {
			response := models.HelloResponse{
				Message: "err token",
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		}

	} else {
		response := models.HelloResponse{
			Message: "err method",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func MetaHandler(w http.ResponseWriter, r *http.Request) {
	var metaData models.MetaData
	err := json.NewDecoder(r.Body).Decode(&metaData)
	if err != nil {
		response := models.HelloResponse{
			Message: "err json",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
	if r.Method == http.MethodPost {

		token := r.Header.Get("token")
		if checkToken(token) {
			_, ok := ReferenceMap[metaData.MetaCode]
			if ok {
				response := models.HelloResponse{
					Message: "done",
				}
				if metaData.MetaTime == (time.Time{}) {
					SaveData(metaData.MetaCode, metaData.MetaFloat)
				} else {
					SaveDataByDate(metaData.MetaCode, metaData.MetaFloat, metaData.MetaTime)
				}
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(response)
			} else {
				response := models.HelloResponse{
					Message: "err code",
				}

				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(response)
			}
		} else {
			response := models.HelloResponse{
				Message: "err token",
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		}

	} else {
		response := models.HelloResponse{
			Message: "err method",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
