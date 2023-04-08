package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Create User godoc
// @ID create_user
// @Router /user [POST]
// @Summary Create User
// @Description Create User
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.CreateUser true "CreateUserRequest"
// @Success 201 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateUser(c *gin.Context) {

	var createUser models.CreateUser
	id := c.Param("id")

	err := c.ShouldBindJSON(&createUser) // parse req body to given type struct
	if err != nil {
		h.handlerResponse(c, "create User", http.StatusBadRequest, err.Error())
		return
	}

	id, err = h.storages.User().CreateUser(context.Background(), &createUser)
	if err != nil {
		h.handlerResponse(c, "storage.User.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.storages.User().GetByIDUser(context.Background(), &models.UserPrimaryKey{UserId: id})
	if err != nil {
		h.handlerResponse(c, "storage.User.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create User", http.StatusCreated, resp)
}

// Get By ID User godoc
// @ID get_by_id_user
// @Router /user/{id} [GET]
// @Summary Get By ID User
// @Description Get By ID User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIdUser(c *gin.Context) {

	id := c.Param("id")

	resp, err := h.storages.User().GetByIDUser(context.Background(), &models.UserPrimaryKey{UserId: id})
	if err != nil {
		h.handlerResponse(c, "storage.User.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get User by id", http.StatusCreated, resp)
}

// Get List User godoc
// @ID get_list_user
// @Router /user [GET]
// @Summary Get List User
// @Description Get List User
// @Tags User
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListUser(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list User", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list User", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.User().GetListUser(context.Background(), &models.GetListUserRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.User.getlist", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list User response", http.StatusOK, resp)
}

// Login godoc
// @ID login
// @Router /login [POST]
// @Summary Create Login
// @Description Create Login
// @Tags Login
// @Accept json
// @Produce json
// @Param Login body models.Login true "LoginRequestBody"
// @Success 201 {object} models.LoginResponse "GetLoginBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *Handler) Login(c *gin.Context) {
	var login models.Login

	err := c.ShouldBindJSON(&login)
	if err != nil {
		log.Printf("error whiling create: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.storages.User().GetByIDUser(
		context.Background(),
		&models.UserPrimaryKey{Login: login.Login},
	)

	if err != nil {
		log.Printf("error whiling GetByPKey: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling GetByPKey").Error())
		return
	}

	if login.Password != resp.Password {
		c.JSON(http.StatusInternalServerError, errors.New("error password is not correct").Error())
		return
	}

	data := map[string]interface{}{
		"id": resp.UserId,
	}

	token, err := helper.GenerateJWT(data, time.Hour*24,"secret")
	if err != nil {
		log.Printf("error whiling GenerateJWT: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling GenerateJWT").Error())
		return
	}

	c.JSON(http.StatusCreated, models.LoginResponse{Token: token})
}

// Register godoc
// @ID register
// @Router /register [POST]
// @Summary Create Register
// @Description Create Register
// @Tags Register
// @Accept json
// @Produce json
// @Param Regester body models.User true "RegisterRequestBody"
// @Success 201 {object} models.RegisterResponse "GetRegisterBody"
// @Response 400 {object} string "Invalid Argument"
// @Failure 500 {object} string "Server Error"
func (h *Handler) Register(c *gin.Context) {
	var register models.User

	err := c.ShouldBindJSON(&register)
	if err != nil {
		log.Printf("error whiling create: %v\n", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.User().CreateUser(
		context.Background(),
		&models.CreateUser{
			FirstName:   register.FirstName,
			LastName:    register.LastName,
			Login:       register.Login,
			Password:    register.Password,
			PhoneNumber: register.PhoneNumber,
		},
	)

	if err != nil {
		log.Printf("error whiling GetByPKey: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling GetByPKey").Error())
		return
	}

	data := map[string]interface{}{
		"id": id,
	}

	token, err := helper.GenerateJWT(data, time.Hour*24, h.cfg.AuthSecretKey)
	if err != nil {
		log.Printf("error whiling GenerateJWT: %v\n", err)
		c.JSON(http.StatusInternalServerError, errors.New("error whiling GenerateJWT").Error())
		return
	}

	c.JSON(http.StatusCreated, models.RegisterResponse{Register: token})
}
