package usecase

import (

	"time"
	"github.com/golang-jwt/jwt/v4"
	"github.com/shiori-42/textbook_change_app/model"
	"github.com/shiori-42/textbook_change_app/repository"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(user *model.User) (model.UserResponse, error)
	Login(user model.User) (string, error)
}

type userUsecase struct {
	ur repository.IUserRepository
}

func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &userUsecase{ur}
}

func (uu *userUsecase) SignUp(user *model.User) (model.UserResponse, error) {
	hash,err := bcrypt.GenerateFromPassword([]byte(user.Password),10)
	if err != nil {
		return model.UserResponse{}, err
	}
	newUserUser:=model.User{Email: user.Email,Password: string(hash)}
	if err:=uu.ur.CreateUser(&newUserUser); err != nil {
		return model.UserResponse{}, err
	}
	resUser:=model.UserResponse{
		ID:	   newUserUser.ID,
		Email: newUserUser.Email,
	}
	return resUser,nil

}

func (uu *userUsecase) Login(user model.User) (string, error) {
	storeUser:=model.User{}
	if err:=uu.ur.GetUserByEmail(&storeUser,user.Email);err!=nil{
		return "",err
	}
	err:=bcrypt.CompareHashAndPassword([]byte(storeUser.Password),[]byte(user.Password))
	if err!=nil{
	return "", nil
	}
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"user_id":storeUser.ID,
		"exp":time.Now().Add(time.Hour*12).Unix(),
	})
	tokenString,err:=token.SignedString([]byte("secret"))
	if err!=nil{
		return "",err
	}
	return tokenString,nil
}