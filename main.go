package main

import (
	"fmt"
	"net/http"
	"oldsouth/services"
)

func init() {

}
func main() {
	fmt.Println("starting")
	services.LoadData()
	services.LoadReferenceFile()
	fmt.Println(services.GetData("W"))

	http.HandleFunc("/meta", services.MetaHandler)
	http.HandleFunc("/metas", services.MetasHandler)
	http.HandleFunc("/health", services.HealthHandler)
	fmt.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Server error: %v", err)
	}
}
