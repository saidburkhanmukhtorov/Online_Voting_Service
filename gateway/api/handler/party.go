package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/online_voting_service/gateway/genproto/public"
)

// @Summary Create a new party
// @Description Creates a new party.
// @Tags Parties
// @Accept json
// @Produce json
// @Param party body public.PartyCreate true "Party data"
// @Success 200 {object} public.Party
// @Failure 400 {object} string "Invalid request body"
// @Failure 500 {object} string "Internal server error"
// @Router /party [post]
func (h *HandlerStruct) CreatePartyHandler(c *gin.Context) {
	var (
		partyReq public.PartyCreate
		err      error
	)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	if err = c.BindJSON(&partyReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error Binding data: " + err.Error()})
		return
	}

	partyRes, err := h.Party.Create(ctx, &partyReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create party: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"party": partyRes})
}

// @Summary Update a party
// @Description Updates an existing party.
// @Tags Parties
// @Accept json
// @Produce json
// @Param party body public.PartyUpdate true "Party data"
// @Success 200 {object} vote.Void
// @Failure 400 {object} string "Invalid request body"
// @Failure 404 {object} string "Resource not found"
// @Failure 500 {object} string "Internal server error"
// @Router /party [put]
func (h *HandlerStruct) UpdatePartyHandler(c *gin.Context) {
	var (
		partyReq public.PartyUpdate
		err      error
	)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	if err = c.BindJSON(&partyReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error Binding data: " + err.Error()})
		return
	}

	_, err = h.Party.Update(ctx, &partyReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update party: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Party updated successfully"})
}

// @Summary Delete a party
// @Description Deletes a party by its ID.
// @Tags Parties
// @Accept json
// @Produce json
// @Param id query string true "Party ID"
// @Success 200 {object} vote.Void
// @Failure 400 {object} string "Invalid request parameters"
// @Failure 404 {object} string "Resource not found"
// @Failure 500 {object} string "Internal server error"
// @Router /party [delete]
func (h *HandlerStruct) DeletePartyHandler(c *gin.Context) {
	var (
		partyReq public.PartyDelete
		err      error
	)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	partyReq.Id = c.Query("id")

	_, err = h.Party.Delete(ctx, &partyReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete party: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Party deleted successfully"})
}

// @Summary Get a party by its ID
// @Description Retrieves a party by its ID.
// @Tags Parties
// @Accept json
// @Produce json
// @Param id query string true "Party ID"
// @Success 200 {object} public.Party
// @Failure 400 {object} string "Invalid request parameters"
// @Failure 404 {object} string "Resource not found"
// @Failure 500 {object} string "Internal server error"
// @Router /party/id [get]
func (h *HandlerStruct) GetPartyByIdHandler(c *gin.Context) {
	var (
		partyReq public.PartyById
		err      error
	)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	partyReq.Id = c.Query("id")

	partyRes, err := h.Party.GetById(ctx, &partyReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get party: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"party": partyRes})
}

// @Summary Get all parties
// @Description Retrieves all parties.
// @Tags Parties
// @Accept json
// @Produce json
// @Param opened_date query string false "Opened date (optional)"
// @Param name query string false "Name (optional)"
// @Param slogan query string false "Slogan (optional)"
// @Param description query string false "Description (optional)"
// @Success 200 {object} public.GetAllPartyResponse
// @Failure 400 {object} string "Invalid request parameters"
// @Failure 500 {object} string "Internal server error"
// @Router /party/all [get]
func (h *HandlerStruct) GetAllPartiesHandler(c *gin.Context) {
	var (
		partyReq public.GetAllPartyRequest
		err      error
	)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 50*time.Second)
	defer cancel()

	partyReq.OpenedDate = c.Query("opened_date")
	partyReq.Name = c.Query("name")
	partyReq.Slogan = c.Query("slogan")
	partyReq.Description = c.Query("description")

	partyRes, err := h.Party.GetAll(ctx, &partyReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get parties: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"parties": partyRes})
}
