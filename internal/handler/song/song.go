package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/galileoarius8/final-project-edspert/internal/entity"
	"github.com/galileoarius8/final-project-edspert/internal/helper"
	"github.com/gin-gonic/gin"
)

// It will call the function Get in song usecase
func (handler *songHandler) Get(context *gin.Context) {
	// Get id from request param
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// Call the usecase
	song, err := handler.songUsecase.Get(context, id)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", song)
	context.JSON(http.StatusOK, res)
}

// It will call the function Create in song usecase
func (handler *songHandler) Create(context *gin.Context) {
	var requestBody entity.Song

	// Get request body from user
	err := context.BindJSON(&requestBody)
	if err != nil {
		res := helper.BuildErrorResponse("Bad request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	// Call the usecase
	song, err := handler.songUsecase.Create(context, &requestBody)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", song)
	context.JSON(http.StatusOK, res)
}

// It will call the function GetAllSong in song usecase
func (handler *songHandler) GetAllSong(context *gin.Context) {
	// Get all songs from usecase
	limit := context.Query("limit")
	if limit == "" {
		limit = "1000"
	}
	offset := context.Query("offset")
	if offset == "" {
		offset = "0"
	}
	filter_album_id := context.Query("album_id")
	if filter_album_id == "" {
		filter_album_id = "nothing"
	}

	fmt.Println(offset, limit, filter_album_id)
	songs, err := handler.songUsecase.GetAllSong(context, offset, limit, filter_album_id)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", songs)
	context.JSON(http.StatusOK, res)
}

// It will call the function BatchCreate in song usecase
func (handler *songHandler) BatchCreate(context *gin.Context) {
	var requestBody []entity.Song

	// Get request body from user
	err := context.BindJSON(&requestBody)
	if err != nil {
		res := helper.BuildErrorResponse("Bad request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	songs, err := handler.songUsecase.BatchCreate(context, requestBody)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", songs)
	context.JSON(http.StatusOK, res)
}

// It will call the function Update in song usecase
func (handler *songHandler) Update(context *gin.Context) {
	var requestBody entity.Song

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
	song, err := handler.songUsecase.Update(context, requestBody)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", song)
	context.JSON(http.StatusOK, res)
}

// It will call the function Delete in song usecase
func (handler *songHandler) Delete(context *gin.Context) {
	// Get id from request param
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	err = handler.songUsecase.Delete(context, id)
	if err != nil {
		res := helper.BuildErrorResponse("Internal Server Error", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := helper.BuildResponse(true, "OK!", helper.EmptyObj{})
	context.JSON(http.StatusOK, res)
}
