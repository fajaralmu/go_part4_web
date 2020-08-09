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

const appSession string = "APP_SESSION"
const requestURI string = "request-uri"
const keyLoggedUser string = "logged_user"
const keyAuthValid string = "valid"

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

func getLatestURI(w http.ResponseWriter, r *http.Request) string {
	savedURI := getSessionVal(w, r, appSession, requestURI)
	if savedURI == nil {
		return ""
	}
	if _, ok := (savedURI).(string); ok {
		return (savedURI).(string)
	}
	return ""
}

func setLatestURI(w http.ResponseWriter, r *http.Request, path string) bool {
	sessionUpdated := setSessionValue(w, r, appSession, requestURI, path)
	log.Println("setLatestURI:", path, " sessionUpdated: ", sessionUpdated)
	return sessionUpdated
}

func validateSessionn(w http.ResponseWriter, r *http.Request) bool {
	userSession := getUserFromSession(w, r)
	if nil == userSession {
		return false
	}
	return getUserByUsernameAndPassword(userSession) != nil
}

func getSessionVal(w http.ResponseWriter, r *http.Request, sessionName string, sessionKey string) interface{} {
	session, err := getSessionValue(r, sessionName)
	if err != nil {
		return nil
	}
	return session.Values[sessionKey]
}

func getUserFromSession(w http.ResponseWriter, r *http.Request) *entities.User {
	session, err := getSessionValue(r, appSession)
	if err != nil {
		return nil
	}

	sessVal := session.Values[keyLoggedUser]

	if _, ok := sessVal.(*entities.User); !ok {
		// Handle the case that it's not an expected type

		log.Printf("Cannot get user from session")
		return nil
	}

	if session.Values[keyAuthValid] == true {
		return sessVal.(*entities.User)
	}
	return nil
}

func setUserToSession(w http.ResponseWriter, r *http.Request, user *entities.User) (updated bool) {
	if nil != user {
		updated = setSessionValue(w, r, appSession, keyLoggedUser, user)
		updated = setSessionValue(w, r, appSession, keyAuthValid, true)
	} else {
		updated = setSessionValue(w, r, appSession, keyAuthValid, false)
	}

	log.Printf("sessionUpdated: %v", updated)
	return updated
}

func setSessionValue(w http.ResponseWriter, r *http.Request, sessionName string, sessionKey string, sessionValue interface{}) bool {
	session, err := getSessionValue(r, sessionName)
	if err != nil {
		log.Printf("get Session Value err: %v", err.Error())
		return false
	}

	session.Values[sessionKey] = sessionValue
	err = session.Save(r, w)
	if err != nil {
		log.Printf("Save session err: %v", err.Error())
	}
	log.Printf("Session updated : %v", err == nil)
	return err == nil
}
