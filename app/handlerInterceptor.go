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
	return true
}
