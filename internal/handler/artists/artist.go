package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/galileoarius8/final-project-edspert/internal/entity"
	"github.com/galileoarius8/final-project-edspert/internal/helper"
	"github.com/gin-gonic/gin"
)

// It will call the function Get in artist usecase
func (handler *artistHandler) Get(context *gin.Context) {
	// Get id from request param
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// Call the usecase
	artist, err := handler.artistUsecase.Get(context, id)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", artist)
	context.JSON(http.StatusOK, res)
}

// It will call the function Create in artist usecase
func (handler *artistHandler) Create(context *gin.Context) {
	var requestBody entity.Artist

	// Get request body from user
	err := context.BindJSON(&requestBody)
	if err != nil {
		res := helper.BuildErrorResponse("Bad request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// Call the usecase
	artist, err := handler.artistUsecase.Create(context, &requestBody)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	if artist.Name == "" {
		res := helper.BuildErrorResponse("Bad request", "name property can not be null", helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", artist)
	context.JSON(http.StatusOK, res)
}

// It will call the function GetAllArtist in artist usecase
func (handler *artistHandler) GetAllArtist(context *gin.Context) {
	// Get all artists from usecase
	limit := context.Query("limit")
	if limit == "" {
		limit = "1000"
	}
	offset := context.Query("offset")
	if offset == "" {
		offset = "0"
	}
	fmt.Println(offset, limit)
	artists, err := handler.artistUsecase.GetAllArtist(context, offset, limit)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", artists)
	context.JSON(http.StatusOK, res)
}

// It will call the function BatchCreate in artist usecase
func (handler *artistHandler) BatchCreate(context *gin.Context) {
	var requestBody []entity.Artist

	// Get request body from user
	err := context.BindJSON(&requestBody)
	if err != nil {
		res := helper.BuildErrorResponse("Bad request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	artists, err := handler.artistUsecase.BatchCreate(context, requestBody)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", artists)
	context.JSON(http.StatusOK, res)
}

// It will call the function Update in artist usecase
func (handler *artistHandler) Update(context *gin.Context) {
	var requestBody entity.Artist

	// Get id from request param
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// Get request body from user
	err = context.BindJSON(&requestBody)
	if err != nil {
		res := helper.BuildErrorResponse("Bad request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// Set id from params
	requestBody.ID = id

	// Call the usecase
	artist, err := handler.artistUsecase.Update(context, requestBody)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", artist)
	context.JSON(http.StatusOK, res)
}

// It will call the function Delete in artist usecase
func (handler *artistHandler) Delete(context *gin.Context) {
	// Get id from request param
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	err = handler.artistUsecase.Delete(context, id)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", helper.EmptyObj{})
	context.JSON(http.StatusOK, res)
}
