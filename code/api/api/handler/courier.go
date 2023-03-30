package handler

import (
	"app/api/models"
	"app/pkg/helper"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create Courier godoc
// @ID create_courier
// @Router /courier [POST]
// @Summary Create Courier
// @Description Create Courier
// @Tags Courier
// @Accept json
// @Produce json
// @Param courier body models.CreateCourier true "CreateCourier"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateCourier(c *gin.Context) {

	var createCourier models.CreateCourier

	err := c.ShouldBindJSON(&createCourier)
	if err != nil {
		h.handlerResponse(c, "create Courier", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.Courier().CreateCourier(context.Background(), &createCourier)
	if err != nil {
		h.handlerResponse(c, "storage.Courier.create", http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println("id", id)

	resp, err := h.storages.Courier().GetByIdCourier(context.Background(), &models.CourierPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Courier.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Courier", http.StatusCreated, resp)
}

// Get By ID Courier godoc
// @ID get_by_id_courier
// @Router /courier/{id} [GET]
// @Summary Get By ID Courier
// @Description Get By ID Courier
// @Tags Courier
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIdCourier(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id Courier", http.StatusBadRequest, "invalid Courier id")
		return
	}

	resp, err := h.storages.Courier().GetByIdCourier(context.Background(), &models.CourierPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Courier.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Courier", http.StatusCreated, resp)
}

// Get List Courier godoc
// @ID get_list_courier
// @Router /courier [GET]
// @Summary Get List Courier
// @Description Get List Courier
// @Tags Courier
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListCourier(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list Courier", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list Courier", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.Courier().GetListCourier(context.Background(), &models.GetListCourierRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.Courier.getlist", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list Courier response", http.StatusOK, resp)
}

// update Courier godoc
// @ID update_courier
// @Router /courier/{id} [PUT]
// @Summary update Courier
// @Description update Courier
// @Tags Courier
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param courier body models.UpdateCourier true "UpdateCourier"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateCourier(c *gin.Context) {

	var updateCourier models.UpdateCourier

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id Courier", http.StatusBadRequest, "invalid Courier id")
		return
	}

	err := c.ShouldBindJSON(&updateCourier)
	if err != nil {
		h.handlerResponse(c, "update Courier", http.StatusBadRequest, err.Error())
		return
	}

	updateCourier.Id = id

	rowsAffected, err := h.storages.Courier().UpdateCourier(context.Background(), &updateCourier)
	if err != nil {
		h.handlerResponse(c, "storage.Courier.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.Courier.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.Courier().GetByIdCourier(context.Background(), &models.CourierPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Courier.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update Courier", http.StatusAccepted, resp)
}

// Update Patch Courier godoc
// @ID updat_patch_courier
// @Router /courier/{id} [PATCH]
// @Summary update Patch Courier
// @Description Update Patch Courier
// @Tags Courier
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param courier body models.PatchRequest true "UpdatPatchCourier"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdatePatchCourier(c *gin.Context) {

	var object models.PatchRequest

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id Courier", http.StatusBadRequest, "invalid Courier id")
		return
	}

	err := c.ShouldBindJSON(&object)
	if err != nil {
		h.handlerResponse(c, "update patch Courier", http.StatusBadRequest, err.Error())
		return
	}

	object.ID = id

	rowsAffected, err := h.storages.Courier().PatchCourier(context.Background(), &object)
	if err != nil {
		h.handlerResponse(c, "storage.Courier.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.Courier.patch", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.storages.Courier().GetByIdCourier(context.Background(), &models.CourierPrimaryKey{Id: object.ID})
	if err != nil {
		h.handlerResponse(c, "storage.Courier.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update patch Courier", http.StatusAccepted, resp)
}

// Delete Courier godoc
// @ID delete_courier
// @Router /courier/{id} [DELETE]
// @Summary Delete Courier
// @Description Delete Courier
// @Tags Courier
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteCourier(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "get by id Courier", http.StatusBadRequest, "invalid Courier id")
		return
	}

	err := h.storages.Courier().DeleteCourier(context.Background(), &models.CourierPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Courier.update", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "update Courier", http.StatusAccepted, nil)
}
