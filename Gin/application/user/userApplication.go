package user

import (
	"errors"
	"esm-backend/models/domain"
	"esm-backend/models/in"
	"esm-backend/models/out"
	"esm-backend/utilities"

	"gorm.io/gorm"
)

func Login(credentials in.CredentialsInput, db *gorm.DB) (out.ResponseOutput, error) {
	User := domain.User{}
	var error error
	msg := "Invalid username/password combination"

	if results := db.Where("username = ? OR email = ?", credentials.UserName, credentials.UserName).
		First(&User); results.Error != nil || results.RowsAffected < 1 {
		error = errors.New(msg)
		return out.ResponseOutput{}, error
	}

	check := utilities.CheckPasswordHash(credentials.Password, User.Password)

	if !check {
		error = errors.New(msg)
		return out.ResponseOutput{}, error
	}

	payload := out.PayloadOut{
		Username: User.Username,
		Email:    User.Email,
		Id:       User.ID,
	}

	token, error := utilities.GenerateJwtToken(payload)

	if error != nil {
		return out.ResponseOutput{}, error
	}

	userOut := out.UserOut{
		Name:     User.Username,
		LastName: User.LastName,
		Email:    User.Email,
		Username: User.Username,
	}

	response := out.ResponseOutput{
		User:  userOut,
		Token: token,
	}

	return response, nil
}

func Signup(userInput in.UserInput, db *gorm.DB) error {
	if len(userInput.Name) < 3 {
		return errors.New("Name should be at least 3 characters long!")
	}

	if len(userInput.LastName) < 3 {
		return errors.New("Lastname should be at least 3 characters long!")
	}

	if len(userInput.Username) < 3 {
		return errors.New("Username should be at least 3 characters long!")
	}

	if len(userInput.Email) < 3 {
		return errors.New("Email should be at least 3 characters long!")
	}

	if len(userInput.Password) < 3 {
		return errors.New("Password should be at least 3 characters long!")
	}

	hashedPassword, err := utilities.GeneratehashPassword(userInput.Password)

	if err != nil {
		return err
	}

	user := domain.User{}
	// mapping to entity
	user.Name = userInput.Name
	user.LastName = userInput.LastName
	user.Username = userInput.Username
	user.Email = userInput.Email
	user.Password = hashedPassword

	if result := db.Create(&user); result.Error != nil {
		return errors.New("Failed To Add new User in database! \n" + result.Error.Error())
	}

	return nil
}

func Verify(token string, db *gorm.DB) (out.ClaimsOutput, error) {
	//last10 := token[len(token)-10:]
	//fmt.Println("verifiying token ended with..." + last10)

	claims, err := utilities.VerifyJwtToken(token)

	if err != nil {
		return out.ClaimsOutput{}, err
	}

	return *claims, nil
}
