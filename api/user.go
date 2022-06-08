package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "multitenant.com/app/db/sqlc"
	pubsub "multitenant.com/app/pubsub"
)

type createUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	UserName string `json:"userName" binding:"required,alphanum"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.CreateUserParams{
		Email:    req.Email,
		UserName: req.UserName,
	}

	user, err := server.store.CreateUser(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	//TODO: publish User Read Event to Topic
	pubsub.PublishToTopic()

	ctx.JSON(http.StatusCreated, gin.H{"id": user.ID})
}

type getUserRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type userResponseDto struct {
	ID               int64     `json:"id"`
	FirstName        string    `json:"first_name"`
	LastName         string    `json:"last_name"`
	FullName         string    `json:"full_name"`
	IsEmailConfirmed bool      `json:"is_email_confirmed"`
	Email            string    `json:"email"`
	UserName         string    `json:"user_name"`
	RoleID           int32     `json:"role_id"`
	CreatedBy        int64     `json:"created_by"`
	UpdatedBy        int64     `json:"updated_by"`
	CreatedTimestamp time.Time `json:"created_timestamp"`
	UpdatedTimestamp time.Time `json:"updated_timestamp"`
	IsArchived       bool      `json:"is_archived"`
}

func (server *Server) getUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	//TODO: check if there is automapper
	var response userResponseDto
	response.ID = user.ID
	response.FirstName = user.FirstName.String
	response.LastName = user.LastName.String
	response.FullName = user.FullName.String
	response.IsEmailConfirmed = user.IsEmailConfirmed.Bool
	response.Email = user.Email
	response.UserName = user.UserName
	response.RoleID = user.RoleID.Int32
	response.CreatedBy = user.CreatedBy.Int64
	response.UpdatedBy = user.UpdatedBy.Int64
	response.CreatedTimestamp = user.CreatedTimestamp.Time
	response.UpdatedTimestamp = user.UpdatedTimestamp.Time
	response.IsArchived = user.IsArchived.Bool

	ctx.JSON(http.StatusOK, response)
}

type getUsersRequest struct {
	PageID   int32 `form:"pageId" binding:"required,min=1"`
	PageSize int32 `form:"pageSize" binding:"required,min=1,max=100"`
}

func (server *Server) getUsers(ctx *gin.Context) {
	var req getUsersRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListUsersParams{Limit: req.PageSize, Offset: (req.PageID - 1) * req.PageSize}
	users, err := server.store.ListUsers(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	usersResponse := []userResponseDto{}

	for _, user := range users {
		response := userResponseDto{
			ID:               user.ID,
			FirstName:        user.FirstName.String,
			LastName:         user.LastName.String,
			FullName:         user.FullName.String,
			IsEmailConfirmed: user.IsEmailConfirmed.Bool,
			Email:            user.Email,
			UserName:         user.UserName,
			RoleID:           user.RoleID.Int32,
			CreatedBy:        user.CreatedBy.Int64,
			UpdatedBy:        user.UpdatedBy.Int64,
			CreatedTimestamp: user.CreatedTimestamp.Time,
			UpdatedTimestamp: user.UpdatedTimestamp.Time,
			IsArchived:       user.IsArchived.Bool,
		}

		usersResponse = append(usersResponse, response)
	}

	ctx.JSON(http.StatusOK, usersResponse)
}
