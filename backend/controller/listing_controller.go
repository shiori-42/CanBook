package controller

import (
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/shiori-42/textbook_change_app/model"
	"github.com/shiori-42/textbook_change_app/usecase"
)

type IListingController interface {
	GetAllMyListings(c echo.Context) error
	GetMyListingById(c echo.Context) error
	CreateListing(c echo.Context) error
	UpdateListing(c echo.Context) error
	DeleteListing(c echo.Context) error
}

type listingController struct {
	lu usecase.IListingUsecase
}

func NewListingController(lu usecase.IListingUsecase) IListingController {
	return &listingController{lu}
}

func (lc *listingController) GetAllMyListings(c echo.Context) error {
	user:=c.Get("user").(*jwt.Token)
	claims:=user.Claims.(jwt.MapClaims)
	userId:=claims["user_id"]
	listingRes,err:=lc.lu.GetAllMyListings(uint(userId.(float64)))
	if err!=nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK,listingRes)
}

func (lc *listingController) GetMyListingById(c echo.Context) error {
	user:=c.Get("user").(*jwt.Token)
	claims:=user.Claims.(jwt.MapClaims)
	userId:=claims["user_id"]
	id:=c.Param("id")
	listingId,_:=strconv.Atoi(id)
	listingRes,err:=lc.lu.GetMyListingById(uint(userId.(float64)),uint(listingId))
	if err!=nil{
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK,listingRes)
}

func (lc *listingController) CreateListing(c echo.Context) error {
	user:=c.Get("user").(*jwt.Token)
	claims:=user.Claims.(jwt.MapClaims)
	userId:=claims["user_id"]
	listing:=model.Listing{}
	if err:=c.Bind(&listing);err!=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	listing.UserID=uint(userId.(float64))
	listingRes,err:=lc.lu.CreateListing(listing)
	if err!=nil{
		return c.JSON(http.StatusInternalServerError,err.Error())
	}
	return c.JSON(http.StatusCreated,listingRes)
}

func (lc *listingController) UpdateListing(c echo.Context) error {
	user:=c.Get("user").(*jwt.Token)
	claims:=user.Claims.(jwt.MapClaims)
	userId:=claims["user_id"]
	id:=c.Param("id")
	listingId,_:=strconv.Atoi(id)

	listing:=model.Listing{}
	if err:=c.Bind(&listing);err!=nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	listingRes,err:=lc.lu.UpdateListing(listing,uint(userId.(float64)),uint(listingId))
	if err!=nil{
		return c.JSON(http.StatusInternalServerError,err.Error())
	}
	return c.JSON(http.StatusOK,listingRes)
}

func (lc *listingController) DeleteListing(c echo.Context) error {
	user:=c.Get("user").(*jwt.Token)
	claims:=user.Claims.(jwt.MapClaims)
	userId:=claims["user_id"]
	id:=c.Param("id")
	listingId,_:=strconv.Atoi(id)
	err:=lc.lu.DeleteListing(uint(userId.(float64)),uint(listingId))
	if err!=nil{
		return c.JSON(http.StatusInternalServerError,err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}