package app

import (
	"log"
	"net/http"
)

func mvcPreHandle(w http.ResponseWriter, r *http.Request) bool {
	log.Println("mvcPreHandle: ", r.RequestURI)
	return true
}
func apiPreHandle(w http.ResponseWriter, r *http.Request) bool {
	log.Println("apiPreHandle: ", r.RequestURI)
	// var request entities.WebRequest
	// err := nil
	// log.Printf("apiPreHandle result: %v", err == nil)
	// return err == nil
	return true
}
