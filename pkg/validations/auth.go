package validations

import (
	"encoding/json"
	"github.com/bednyak/go-react-jwt-auth/pkg/errors"
	"github.com/bednyak/go-react-jwt-auth/pkg/structs"
	"net/http"
)

func IsSignUpRequestValid(w http.ResponseWriter, r *http.Request) (structs.User, error) {
	var user structs.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		var err errors.Error
		err = errors.SetError(err, "Error in reading payload.")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
	}

	if user.Email == "" {
		var err errors.Error
		err = errors.SetError(err, "Field 'email' can not be empty.")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
	}

	if user.Name == "" {
		var err errors.Error
		err = errors.SetError(err, "Field 'name' can not be empty.")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
	}

	if user.Role == "" {
		var err errors.Error
		err = errors.SetError(err, "Field 'role' can not be empty.")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
	}

	if user.Password == "" {
		var err errors.Error
		err = errors.SetError(err, "Field 'password' can not be empty.")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
	}

	return user, err
}

func IsSignInRequestValid(w http.ResponseWriter, r *http.Request) (structs.User, error) {
	var user structs.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		var err errors.Error
		err = errors.SetError(err, "Error in reading payload.")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
	}

	if user.Email == "" {
		var err errors.Error
		err = errors.SetError(err, "Field 'email' can not be empty.")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
	}

	if user.Password == "" {
		var err errors.Error
		err = errors.SetError(err, "Field 'password' can not be empty.")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
	}

	return user, err
}
