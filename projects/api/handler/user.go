package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateUser(c *gin.Context) {

	var CreateUser models.CreateUser

	err := c.ShouldBindJSON(&CreateUser)
	if err != nil {
		h.handlerResponse(c, "create user", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.User().CreateUser(&CreateUser)
	if err != nil {
		h.handlerResponse(c, "storage.user.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.storages.User().GetByIDUser(&models.UserPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.user.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create user", http.StatusCreated, resp)
}

func (h *Handler) UpdateUser(c *gin.Context) {
	var updateUser models.UpdateUser

	err := c.ShouldBindJSON(&updateUser)

	if err != nil {
		h.handlerResponse(c, "create User", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.User().UpdateUser(&updateUser)
	if err != nil {
		h.handlerResponse(c, "storage book update", http.StatusInternalServerError, err.Error())
		return
	}
	resp, err := h.storages.User().GetByIDUser(&models.UserPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.book.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "Update book", http.StatusAccepted, resp)
}
func (h *Handler) DeleteUser(c *gin.Context) {
	var deleteUser models.DeleteUser

	err := c.ShouldBindJSON(&deleteUser)

	if err != nil {
		h.handlerResponse(c, "create User", http.StatusBadRequest, err.Error())
		return
	}

	err = h.storages.User().DeleteUser(&deleteUser)
	if err != nil {
		h.handlerResponse(c, "storage User delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "delete user", http.StatusOK, nil)
}

func (h *Handler) GetByIdUser(c *gin.Context) {


	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id user", http.StatusBadRequest, "invalid user id")
		return
	}

	resp, err := h.storages.User().GetByIDUser(&models.UserPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.user.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create user", http.StatusCreated, resp)
}

func (h *Handler) GetListUser(c *gin.Context) {

	offset, err := h.getOffSetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list User", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list User", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.User().GetListUser(&models.GetListUserRequest{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		h.handlerResponse(c, "storage.User.getlist", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list User response", http.StatusOK, resp)
}