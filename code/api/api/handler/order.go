package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create Order godoc
// @ID create_order
// @Router /order [POST]
// @Summary Create Order
// @Description Create Order
// @Tags Order
// @Accept json
// @Produce json
// @Param order body models.CreateOrder true "CreateOrder"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateOrder(c *gin.Context) {

	var createOrder models.CreateOrder

	err := c.ShouldBindJSON(&createOrder)
	if err != nil {
		h.handlerResponse(c, "create Order", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.Order().CreateOrder(context.Background(), &createOrder)
	if err != nil {
		h.handlerResponse(c, "storage.Order.create", http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println("id", id)

	resp, err := h.storages.Order().GetByIdOrder(context.Background(), &models.OrderPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Order.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Order", http.StatusCreated, resp)
}

// Get By ID Order godoc
// @ID get_by_id_order
// @Router /order/{id} [GET]
// @Summary Get By ID Order
// @Description Get By ID Order
// @Tags Order
// @Accept json
// @Produce json
// @Param cd path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIdOrder(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id Order", http.StatusBadRequest, "invalid Order id")
		return
	}

	resp, err := h.storages.Order().GetByIdOrder(context.Background(), &models.OrderPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Order.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Order", http.StatusCreated, resp)
}

// Get List Order godoc
// @ID get_list_order
// @Router /order [GET]
// @Summary Get List Order
// @Description Get List Order
// @Tags Order
// @Accept json
// @Produce json
// @Param cffset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListOrder(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list Order", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list Order", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.Order().GetListOrder(context.Background(), &models.GetListOrderRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.Order.getlist", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list Order response", http.StatusOK, resp)
}

// update Order godoc
// @ID update_order
// @Router /order/{id} [PUT]
// @Summary update Order
// @Description update Order
// @Tags Order
// @Accept json
// @Produce json
// @Param cd path string true "id"
// @Param order body models.UpdateOrder true "UpdateOrder"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateOrder(c *gin.Context) {

	var updateOrder models.UpdateOrder

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id Order", http.StatusBadRequest, "invalid Order id")
		return
	}

	err := c.ShouldBindJSON(&updateOrder)
	if err != nil {
		h.handlerResponse(c, "update Order", http.StatusBadRequest, err.Error())
		return
	}

	updateOrder.Id = id

	rowsAffected, err := h.storages.Order().UpdateOrder(context.Background(), &updateOrder)
	if err != nil {
		h.handlerResponse(c, "storage.Order.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.Order.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.Order().GetByIdOrder(context.Background(), &models.OrderPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Order.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update Order", http.StatusAccepted, resp)
}

// Update Patch Order godoc
// @ID updat_patch_order
// @Router /order/{id} [PATCH]
// @Summary update Patch Order
// @Description Update Patch Order
// @Tags Order
// @Accept json
// @Produce json
// @Param cd path string true "id"
// @Param Order body models.PatchRequest true "UpdatPatchOrder"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdatePatchOrder(c *gin.Context) {

	var object models.PatchRequest

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id Order", http.StatusBadRequest, "invalid Order id")
		return
	}

	err := c.ShouldBindJSON(&object)
	if err != nil {
		h.handlerResponse(c, "update patch Order", http.StatusBadRequest, err.Error())
		return
	}

	object.ID = id

	rowsAffected, err := h.storages.Order().PatchOrder(context.Background(), &object)
	if err != nil {
		h.handlerResponse(c, "storage.Order.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.Order.patch", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.Order().GetByIdOrder(context.Background(), &models.OrderPrimaryKey{Id: object.ID})
	if err != nil {
		h.handlerResponse(c, "storage.Order.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update patch Order", http.StatusAccepted, resp)
}

// Delete Order godoc
// @ID delete_order
// @Router /order/{id} [DELETE]
// @Summary Delete Order
// @Description Delete Order
// @Tags Order
// @Accept json
// @Produce json
// @Param cd path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteOrder(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id Order", http.StatusBadRequest, "invalid Order id")
		return
	}

	err := h.storages.Order().DeleteOrder(context.Background(), &models.OrderPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Order.update", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update Order", http.StatusAccepted, nil)
}
