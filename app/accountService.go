package app

import (
	"errors"
	"net/http"

	"github.com/fajaralmu/go_part4_web/entities"
	"github.com/fajaralmu/go_part4_web/repository"
)

func getUserByUsernameAndPassword(user *entities.User) *entities.User {
	list, count := repository.Filter(&[]entities.User{}, entities.Filter{
		FieldsFilter: map[string]interface{}{
			"Username": user.Username,
			"Password": user.Password,
		},
		Exacts: true,
	})
	if count != 1 {
		return nil
	}
	dbUser, ok := list[0].(entities.User)
	if ok {
		return &dbUser
	}
	return nil
}

func authenticateUser(request entities.WebRequest, w http.ResponseWriter, r *http.Request) (response entities.WebResponse, err error) {

	user := request.User
	dbUser := getUserByUsernameAndPassword(user)
	if dbUser == nil {
		return response, errors.New("User Not Found")
	}
	setUserToSession(w, r, dbUser)

	return webResponse("00", "success"), nil
}

func isUsernameAvailable(username string) bool {

	resultList := repository.FilterByKey(&[]entities.User{}, "Username", username)

	return len(resultList) == 0
}
