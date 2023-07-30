package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/bednyak/go-react-jwt-auth/pkg/errors"
	"github.com/bednyak/go-react-jwt-auth/pkg/models"
	"github.com/bednyak/go-react-jwt-auth/pkg/structs"
	"github.com/bednyak/go-react-jwt-auth/pkg/validations"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"time"
)

func SignUp(w http.ResponseWriter, r *http.Request) {

	var user, err = validations.IsSignUpRequestValid(w, r)

	var dbuser, _ = models.GetUserByEmail(user.Email)

	if dbuser.Email != "" {
		var err errors.Error
		err = errors.SetError(err, "Email already in use")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	user.Password, err = generateHashPassword(user.Password)
	if err != nil {
		log.Fatalln("Error in password hashing.")
	}

	models.CreateUser(&models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func SignIn(w http.ResponseWriter, r *http.Request) {

	var authDetails, err = validations.IsSignInRequestValid(w, r)
	var authUser, _ = models.GetUserByEmail(authDetails.Email)

	if authUser.Email == "" {
		var err errors.Error
		err = errors.SetError(err, "Username or Password is incorrect")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	check := checkPasswordHash(authDetails.Password, authUser.Password)

	if !check {
		var err errors.Error
		err = errors.SetError(err, "Username or Password is incorrect")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	validToken, err := generateJWT(authUser.Email, authUser.Role)
	if err != nil {
		var err errors.Error
		err = errors.SetError(err, "Failed to generate token")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	var token structs.Token
	token.Email = authUser.Email
	token.Role = authUser.Role
	token.TokenString = validToken
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}

func generateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateJWT(email, role string) (string, error) {
	var mySigningKey = []byte(os.Getenv("JWT_SECRET_KEY"))
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Errorf("Something went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}
