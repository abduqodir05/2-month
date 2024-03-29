package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"fmt"
	"net/http"

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
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateUser(c *gin.Context) {

	var createUser models.CreateUser

	err := c.ShouldBindJSON(&createUser)
	if err != nil {
		h.handlerResponse(c, "create User", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.User().CreateUser(context.Background(), &createUser)
	if err != nil {
		h.handlerResponse(c, "storage.User.create", http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println("id", id)

	resp, err := h.storages.User().GetByIdUser(context.Background(), &models.UserPrimaryKey{Id: id})
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

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id User", http.StatusBadRequest, "invalid User id")
		return
	}

	resp, err := h.storages.User().GetByIdUser(context.Background(), &models.UserPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.User.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create User", http.StatusCreated, resp)
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

// update User godoc
// @ID update_user
// @Router /user/{id} [PUT]
// @Summary update User
// @Description update User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param user body models.UpdateUser true "UpdateUserRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateUser(c *gin.Context) {

	var updateUser models.UpdateUser

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id User", http.StatusBadRequest, "invalid User id")
		return
	}

	err := c.ShouldBindJSON(&updateUser)
	if err != nil {
		h.handlerResponse(c, "update User", http.StatusBadRequest, err.Error())
		return
	}

	updateUser.Id = id

	rowsAffected, err := h.storages.User().UpdateUser(context.Background(), &updateUser)
	if err != nil {
		h.handlerResponse(c, "storage.User.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.User.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.User().GetByIdUser(context.Background(), &models.UserPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.User.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update User", http.StatusAccepted, resp)
}

// Update Patch User godoc
// @ID updat_patch_User
// @Router /user/{id} [PATCH]
// @Summary update Patch User
// @Description Update Patch User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param user body models.PatchRequest true "UpdatPatchUserRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdatePatchUser(c *gin.Context) {

	var object models.PatchRequest

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id User", http.StatusBadRequest, "invalid User id")
		return
	}

	err := c.ShouldBindJSON(&object)
	if err != nil {
		h.handlerResponse(c, "update patch User", http.StatusBadRequest, err.Error())
		return
	}

	object.ID = id

	rowsAffected, err := h.storages.User().PatchUser(context.Background(), &object)
	if err != nil {
		h.handlerResponse(c, "storage.User.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.User.patch", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.User().GetByIdUser(context.Background(), &models.UserPrimaryKey{Id: object.ID})
	if err != nil {
		h.handlerResponse(c, "storage.User.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update patch User", http.StatusAccepted, resp)
}

// Delete User godoc
// @ID delete_user
// @Router /user/{id} [DELETE]
// @Summary Delete User
// @Description Delete User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteUser(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id User", http.StatusBadRequest, "invalid User id")
		return
	}

	err := h.storages.User().DeleteUser(context.Background(), &models.UserPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.User.update", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update User", http.StatusAccepted, nil)
}
