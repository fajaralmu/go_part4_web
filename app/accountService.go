package app

import (
	"errors"
	"net/http"

	"github.com/fajaralmu/go_part4_web/repository"

	"github.com/fajaralmu/go_part4_web/entities"
)

func Login(request entities.WebRequest, w http.ResponseWriter, r *http.Request) (response entities.WebResponse, err error) {

	user := request.User
	list, count := repository.Filter(&[]entities.User{}, entities.Filter{
		FieldsFilter: map[string]interface{}{
			"Username": user.Username,
			"Password": user.Password,
		},
		Exacts: true,
	})
	if count != 1 {
		return response, errors.New("User Not Found")
	}
	dbUser := list[0].(entities.User)
	setUserToSession(w, r, (&dbUser))
	return webResponse("00", "success"), nil
}
