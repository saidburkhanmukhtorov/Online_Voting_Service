package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/online_voting_service/gateway/genproto/public"
	// Assuming Void is defined here
)

// @Summary Create a new public
// @Description Creates a new public.
// @Tags Publics
// @Accept json
// @Produce json
// @Param public body public.PublicCreate true "Public data"
// @Success 200 {object} public.Public
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Router /public [post]
func (h *HandlerStruct) CreatePublicHandler(c *gin.Context) {
	var (
		publicReq public.PublicCreate
		err       error
	)
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	if err = c.BindJSON(&publicReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error Binding data: " + err.Error()})
		return
	}

	publicRes, err := h.Public.Create(ctx, &publicReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create public: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"public": publicRes})
}

// @Summary Update a public
// @Description Updates an existing public.
// @Tags Publics
// @Accept json
// @Produce json
// @Param public body public.PublicUpdate true "Public data"
// @Success 200 {object} vote.Void
// @Failure 400 {object} string "Invalid request body"
// @Failure 404 {object} string "Resource not found"
// @Failure 500 {object} string "Internal server error"
// @Router /public [put]
func (h *HandlerStruct) UpdatePublicHandler(c *gin.Context) {
	var (
		publicReq public.PublicUpdate
		err       error
	)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	if err = c.BindJSON(&publicReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error Binding data: " + err.Error()})
		return
	}

	_, err = h.Public.Update(ctx, &publicReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update public: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Public updated successfully"})
}

// @Summary Delete a public
// @Description Deletes a public by its ID.
// @Tags Publics
// @Accept json
// @Produce json
// @Param id query string true "Public ID"
// @Success 200 {object} vote.Void
// @Failure 400 {object} string "Invalid request parameters"
// @Failure 404 {object} string "Resource not found"
// @Failure 500 {object} string "Internal server error"
// @Router /public [delete]
func (h *HandlerStruct) DeletePublicHandler(c *gin.Context) {
	var (
		publicReq public.PublicDelete
		err       error
	)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	publicReq.Id = c.Query("id")

	_, err = h.Public.Delete(ctx, &publicReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete public: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Public deleted successfully"})
}

// @Summary Get a public by its ID
// @Description Retrieves a public by its ID.
// @Tags Publics
// @Accept json
// @Produce json
// @Param id query string true "Public ID"
// @Success 200 {object} public.Public
// @Failure 400 {object} string "Invalid request parameters"
// @Failure 404 {object} string "Resource not found"
// @Failure 500 {object} string "Internal server error"
// @Router /public/id [get]
func (h *HandlerStruct) GetPublicByIdHandler(c *gin.Context) {
	var (
		publicReq public.PublicById
		err       error
	)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	publicReq.Id = c.Query("id")

	publicRes, err := h.Public.GetById(ctx, &publicReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get public: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"public": publicRes})
}

// @Summary Get all publics
// @Description Retrieves all publics.
// @Tags Publics
// @Accept json
// @Produce json
// @Param party_id query string false "Party ID (optional)"
// @Param name query string false "Name (optional)"
// @Param last_name query string false "Last name (optional)"
// @Param phone query string false "Phone (optional)"
// @Param email query string false "Email (optional)"
// @Param birthday query string false "Birthday (optional)"
// @Param gender query string false "Gender (optional)"
// @Success 200 {object} public.GetAllPublicRes
// @Failure 400 {object} string "Invalid request parameters"
// @Failure 500 {object} string "Internal server error"
// @Router /public/all [get]
func (h *HandlerStruct) GetAllPublicsHandler(c *gin.Context) {
	var (
		publicReq public.GetAllPublicReq
		err       error
	)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 50*time.Second)
	defer cancel()

	publicReq.PartyId = c.Query("party_id")
	publicReq.Name = c.Query("name")
	publicReq.LastName = c.Query("last_name")
	publicReq.Phone = c.Query("phone")
	publicReq.Email = c.Query("email")
	publicReq.Birthday = c.Query("birthday")
	publicReq.Gender = c.Query("gender")

	publicRes, err := h.Public.GetAll(ctx, &publicReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get publics: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"publics": publicRes})
}
