package handlers

import (
	"ESM-backend-app/pkg/helpers"
	"ESM-backend-app/pkg/models"
	"ESM-backend-app/pkg/models/out"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func (h Handler) SignupUser(w http.ResponseWriter, r *http.Request) {

	User := models.User{}
	json.NewDecoder(r.Body).Decode(&User)
	log.Println("Signin up user " + User.Username)
	if len(User.Name) < 3 {
		helpers.ApiError(w, http.StatusBadRequest, "Name should be at least 3 characters long!")
		return
	}

	if len(User.Username) < 3 {
		helpers.ApiError(w, http.StatusBadRequest, "Username should be at least 3 characters long!")
		return
	}

	if len(User.Email) < 3 {
		helpers.ApiError(w, http.StatusBadRequest, "Email should be at least 3 characters long!")
		return
	}

	if len(User.Password) < 3 {
		helpers.ApiError(w, http.StatusBadRequest, "Password should be at least 3 characters long!")
		return
	}

	hashedPassword, err := helpers.GeneratehashPassword(User.Password)
	if err != nil {
		log.Fatalln(err)
		return
	}

	User.Password = hashedPassword

	if result := h.DB.Create(&User); result.Error != nil {
		helpers.ApiError(w, http.StatusInternalServerError, "Failed To Add new User in database! \n"+result.Error.Error())
		return
	}

	payload := out.Payload{
		Username: User.Username,
		Email:    User.Email,
		Id:       User.ID,
	}

	token, err := helpers.GenerateJwtToken(payload)
	if err != nil {
		helpers.ApiError(w, http.StatusInternalServerError, "Failed To Generate New JWT Token!")
		return
	}

	helpers.RespondWithJSON(w, out.ResponseOutput{
		Token: token,
		User:  User,
	})
}

func (h Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	User := models.User{}

	type Credentials struct {
		UserName string `json:"userName"`
		Password string `json:"password"`
	}
	credentials := Credentials{}
	json.NewDecoder(r.Body).Decode(&credentials)

	if len(credentials.UserName) < 3 {
		helpers.ApiError(w, http.StatusBadRequest, "Invalid Username/Email!")
		return
	}

	if len(credentials.Password) < 3 {
		helpers.ApiError(w, http.StatusBadRequest, "Invalid Password!")
		return
	}

	if results := h.DB.Where("username = ? OR email = ?", credentials.UserName, credentials.UserName).First(&User); results.Error != nil || results.RowsAffected < 1 {
		helpers.ApiError(w, http.StatusNotFound, "Invalid Username/Email, Please Signup!")
		return
	}

	//TODO : changing to  hashing
	check := helpers.CheckPasswordHash(credentials.Password, User.Password)
	if !check {
		helpers.ApiError(w, http.StatusNotFound, "Invalid Credentials!")
		return
	}

	payload := out.Payload{
		Username: User.Username,
		Email:    User.Email,
		Id:       User.ID,
	}

	token, err := helpers.GenerateJwtToken(payload)
	if err != nil {
		helpers.ApiError(w, http.StatusInternalServerError, "Failed To Generate New JWT Token!")
		return
	}

	helpers.RespondWithJSON(w, out.ResponseOutput{
		Token: token,
		User:  User,
	})
}

func (h Handler) Verify(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) < 2 {
		helpers.ApiError(w, http.StatusForbidden, "Token not provided!")
		return
	}

	token := bearerToken[1]
	last10 := token[len(token)-10:]
	log.Println("verifiying token ended with..." + last10)

	claims, err := helpers.VerifyJwtToken(token)

	if err != nil {
		helpers.ApiError(w, http.StatusForbidden, err.Error())
		return
	}
	helpers.RespondWithJSON(w, claims)
}
