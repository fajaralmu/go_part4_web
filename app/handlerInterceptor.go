package app

import (
	"log"
	"net/http"
)

func mvcPreHandle(w http.ResponseWriter, r *http.Request, authenticated bool) bool {
	log.Println("mvcPreHandle: ", r.RequestURI)
	if authenticated {
		sessionValid := validateSessionn(w, r)
		log.Println("Session IS VALID: ", sessionValid)
		return sessionValid
	}
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

func sendRedirect(w http.ResponseWriter, r *http.Request, path string) {
	log.Println("sendRedirect: ", path)
	w.Header().Add("location", path)
	w.WriteHeader(302)
}
