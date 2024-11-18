package handler

import (
	"github.com/Olmosbek510/todo-app"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Create Item
// @Security ApiKeyAuth
// @Tags items
// @Description Create a todo item in a specific list
// @ID create-item
// @Accept json
// @Produce json
// @Param id path int true "List ID"
// @Param input body todo.TodoItem true "Item Input"
// @Success 200 {object} map[string]interface{} "ID of the created item"
// @Failure 400 {object} errorResponse "Invalid request"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /api/lists/{id}/items [post]
func (h *Handler) createItem(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
	}
	var input todo.TodoItem
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	id, err := h.services.TodoItem.Create(userId, listId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Get All Items
// @Security ApiKeyAuth
// @Tags items
// @Description Get all todo items from a specific list
// @ID get-all-items
// @Accept json
// @Produce json
// @Param id path int true "List ID"
// @Success 200 {array} todo.TodoItem
// @Failure 400 {object} errorResponse "Invalid list ID parameter"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /api/lists/{id}/items [get]
func (h *Handler) getAllItems(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
	}

	items, err := h.services.TodoItem.GetAll(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, items)
}

// @Summary Get Item By ID
// @Security ApiKeyAuth
// @Tags items
// @Description Get a specific todo item by its ID
// @ID get-item-by-id
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} todo.TodoItem
// @Failure 400 {object} errorResponse "Invalid item ID parameter"
// @Failure 404 {object} errorResponse "Item not found"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /api/items/{id} [get]
func (h *Handler) getItemById(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
	}

	item, err := h.services.TodoItem.GetById(userId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, item)
}

// @Summary Update Item
// @Security ApiKeyAuth
// @Tags items
// @Description Update a todo item by its ID
// @ID update-item
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Param input body todo.UpdateItemInput true "Update Item Input"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse "Invalid request"
// @Failure 404 {object} errorResponse "Item not found"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /api/items/{id} [put]
func (h *Handler) updateItem(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input todo.UpdateItemInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.TodoItem.Update(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{Status: "ok"})
}

// @Summary Delete Item
// @Security ApiKeyAuth
// @Tags items
// @Description Delete a todo item by its ID
// @ID delete-item
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse "Invalid item ID parameter"
// @Failure 404 {object} errorResponse "Item not found"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /api/items/{id} [delete]
func (h *Handler) deleteItem(c *gin.Context) {
	userId, err := h.getUserId(c)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
	}

	err = h.services.TodoItem.Delete(userId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{Status: "ok"})
}
