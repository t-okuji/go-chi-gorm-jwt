package usecase

import (
	"errors"
	"os"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/go-chi/jwtauth/v5"
	"github.com/t-okuji/go-chi-gorm-jwt/model"
	"github.com/t-okuji/go-chi-gorm-jwt/repository"
)

type IUserUsecase interface {
	SignUp(user model.User) (model.UserResponse, error)
	Login(user model.User) (string, error)
}

type userUsecase struct {
	ur repository.IUserRepository
}

func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &userUsecase{ur}
}

func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	hash, err := argon2id.CreateHash(user.Password, argon2id.DefaultParams)
	if err != nil {
		return model.UserResponse{}, err
	}
	newUser := model.User{Email: user.Email, Password: hash}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}
	resUser := model.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	}
	return resUser, nil
}

func (uu *userUsecase) Login(user model.User) (string, error) {
	storedUser := model.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}
	match, err := argon2id.ComparePasswordAndHash(user.Password, storedUser.Password)
	if err != nil {
		return "", err
	}
	if !match {
		return "", errors.New("invalid password")
	}

	tokenAuth := jwtauth.New("HS256", []byte(os.Getenv("TOKEN_SECRET")), nil)
	claims := map[string]interface{}{"user_id": storedUser.ID}
	// add {exp: time.Time}
	jwtauth.SetExpiry(claims, time.Now().Add(time.Hour*12))
	_, tokenString, err := tokenAuth.Encode(claims)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}
