package app

import (
	"encoding/gob"
	"log"
	"net/http"

	"github.com/fajaralmu/go_part4_web/entities"

	"github.com/gorilla/sessions"
)

// Note: Don't store your key in your source code. Pass it via an
// environmental variable, or flag (or both), and don't accidentally commit it
// alongside your code. Ensure your key is sufficiently random - i.e. use Go's
// crypto/rand or securecookie.GenerateRandomKey(32) and persist the result.
// Ensure SESSION_KEY exists in the environment, or sessions will fail.
var store *sessions.CookieStore = sessions.NewCookieStore([]byte("DONT_PUT_THE_KEY_HERE_THIS_IS_TESTING_PURPOSE"))

func registerSessions() {
	gob.Register(&entities.User{})
}

func getSessionValue(r *http.Request, sessionName string) (*sessions.Session, error) {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func validateSessionn(w http.ResponseWriter, r *http.Request) bool {
	userSession := getUserFromSession(w, r)
	if nil == userSession {
		return false
	}
	return getUserByUsernameAndPassword(userSession) != nil
}

func getUserFromSession(w http.ResponseWriter, r *http.Request) *entities.User {
	session, err := getSessionValue(r, "APP_SESSION")
	if err != nil {
		return nil
	}

	sessVal := session.Values["logged_user"]

	if _, ok := sessVal.(*entities.User); !ok {
		// Handle the case that it's not an expected type

		log.Printf("Cannot get user from session")
		return nil
	}

	if session.Values["valid"] == true {
		return sessVal.(*entities.User)
	}
	return nil
}

func setUserToSession(w http.ResponseWriter, r *http.Request, user *entities.User) (sessionUpdated bool) {
	if nil != user {
		sessionUpdated = setSessionValue(w, r, "APP_SESSION", "logged_user", user)
		sessionUpdated = setSessionValue(w, r, "APP_SESSION", "valid", true)
	} else {
		sessionUpdated = setSessionValue(w, r, "APP_SESSION", "valid", false)
	}

	log.Printf("sessionUpdated: %v", sessionUpdated)
	return sessionUpdated
}

func setSessionValue(w http.ResponseWriter, r *http.Request, sessionName string, sessionKey string, sessionValue interface{}) bool {
	session, err := getSessionValue(r, sessionName)
	if err != nil {
		log.Printf("getSessionValue err: %v", err.Error())
		return false
	}

	session.Values[sessionKey] = sessionValue
	err2 := session.Save(r, w)
	if err2 != nil {
		log.Printf("Save err: %v", err2.Error())
	}
	log.Printf("Saving session : %v", err2 == nil)
	return err2 == nil
}