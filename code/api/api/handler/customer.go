package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


// Create Customer godoc
// @ID create_customer
// @Router /customer [POST]
// @Summary Create Customer
// @Description Create Customer
// @Tags Customer
// @Accept json
// @Produce json
// @Param customer body models.CreateCustomer true "CreateCustomer"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateCustomer(c *gin.Context) {

	var createCustomer models.CreateCustomer

	err := c.ShouldBindJSON(&createCustomer)
	if err != nil {
		h.handlerResponse(c, "create Customer", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.Customer().CreateCustomer(context.Background(), &createCustomer)
	if err != nil {
		h.handlerResponse(c, "storage.Customer.create", http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println("id", id)

	resp, err := h.storages.Customer().GetByIdCustomer(context.Background(), &models.CustomerPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Customer.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Customer", http.StatusCreated, resp)
}

// Get By ID Customer godoc
// @ID get_by_id_customer
// @Router /customer/{id} [GET]
// @Summary Get By ID Customer
// @Description Get By ID Customer
// @Tags Customer
// @Accept json
// @Produce json
// @Param cd path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIdCustomer(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id Customer", http.StatusBadRequest, "invalid Customer id")
		return
	}

	resp, err := h.storages.Customer().GetByIdCustomer(context.Background(), &models.CustomerPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Customer.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Customer", http.StatusCreated, resp)
}

// Get List Customer godoc
// @ID get_list_customer
// @Router /customer [GET]
// @Summary Get List Customer
// @Description Get List Customer
// @Tags Customer
// @Accept json
// @Produce json
// @Param cffset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListCustomer(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list Customer", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list Customer", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.Customer().GetListCustomer(context.Background(), &models.GetListCustomerRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.Customer.getlist", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list Customer response", http.StatusOK, resp)
}

// update Customer godoc
// @ID update_customer
// @Router /customer/{id} [PUT]
// @Summary update Customer
// @Description update Customer
// @Tags Customer
// @Accept json
// @Produce json
// @Param cd path string true "id"
// @Param Customer body models.UpdateCustomer true "UpdateCustomer"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateCustomer(c *gin.Context) {

	var updateCustomer models.UpdateCustomer

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id Customer", http.StatusBadRequest, "invalid Customer id")
		return
	}

	err := c.ShouldBindJSON(&updateCustomer)
	if err != nil {
		h.handlerResponse(c, "update Customer", http.StatusBadRequest, err.Error())
		return
	}

	updateCustomer.Id = id

	rowsAffected, err := h.storages.Customer().UpdateCustomer(context.Background(), &updateCustomer)
	if err != nil {
		h.handlerResponse(c, "storage.Customer.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.Customer.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.Customer().GetByIdCustomer(context.Background(), &models.CustomerPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Customer.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update Customer", http.StatusAccepted, resp)
}

// Update Patch Customer godoc
// @ID updat_patch_Ccstomer
// @Router /customer/{id} [PATCH]
// @Summary update Patch Customer
// @Description Update Patch Customer
// @Tags Customer
// @Accept json
// @Produce json
// @Param cd path string true "id"
// @Param customer body models.PatchRequest true "UpdatPatchCustomer"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdatePatchCustomer(c *gin.Context) {

	var object models.PatchRequest

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id Customer", http.StatusBadRequest, "invalid Customer id")
		return
	}

	err := c.ShouldBindJSON(&object)
	if err != nil {
		h.handlerResponse(c, "update patch Customer", http.StatusBadRequest, err.Error())
		return
	}

	object.ID = id

	rowsAffected, err := h.storages.Customer().PatchCustomer(context.Background(), &object)
	if err != nil {
		h.handlerResponse(c, "storage.Customer.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.Customer.patch", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.Customer().GetByIdCustomer(context.Background(), &models.CustomerPrimaryKey{Id: object.ID})
	if err != nil {
		h.handlerResponse(c, "storage.Customer.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update patch Customer", http.StatusAccepted, resp)
}

// Delete Customer godoc
// @ID delete_customer
// @Router /customer/{id} [DELETE]
// @Summary Delete Customer
// @Description Delete Customer
// @Tags Customer
// @Accept json
// @Produce json
// @Param cd path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteCustomer(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id Customer", http.StatusBadRequest, "invalid Customer id")
		return
	}

	err := h.storages.Customer().DeleteCustomer(context.Background(), &models.CustomerPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Customer.update", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update Customer", http.StatusAccepted, nil)
}
