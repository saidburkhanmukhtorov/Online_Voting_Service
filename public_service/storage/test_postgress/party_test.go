package postgres

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	public "github.com/myfirstgo/online_voting_service/public_service/genproto"
	"github.com/myfirstgo/online_voting_service/public_service/storage/postgres"
	"github.com/stretchr/testify/assert"
)

// Create a test database connection pool for PartyDb
func newTestParty(t *testing.T) *postgres.PartyDb {
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		"sayyidmuhammad",
		"root",
		"localhost",
		5432,
		"public",
	)

	db, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	return &postgres.PartyDb{Db: db}
}

// Create a test party object
func createTestParty() *public.PartyCreate {
	return &public.PartyCreate{
		Id:          uuid.New().String(),
		Name:        "Test Party",
		Slogan:      "Test Party Slogan",
		OpenedDate:  "2024-12-31",
		Description: "Test Party Description",
	}
}

// TestCreateParty tests the Create method
func TestCreateParty(t *testing.T) {
	stgParty := newTestParty(t)
	testParty := createTestParty()

	partyRes, err := stgParty.Create(context.Background(), testParty)
	if err != nil {
		t.Fatalf("Error creating party: %v", err)
	}

	assert.NotEmpty(t, partyRes.Id)
	assert.Equal(t, testParty.Name, partyRes.Name)
	assert.Equal(t, testParty.Slogan, partyRes.Slogan)
	assert.Equal(t, testParty.OpenedDate, partyRes.OpenedDate[:10])
	assert.Equal(t, testParty.Description, partyRes.Description)
}

// TestUpdateParty tests the Update method
func TestUpdateParty(t *testing.T) {
	stgParty := newTestParty(t)
	testParty := createTestParty()

	// Create a party first
	createdParty, err := stgParty.Create(context.Background(), testParty)
	if err != nil {
		t.Fatalf("Error creating party: %v", err)
	}

	// Update the party
	updateParty := &public.PartyUpdate{
		Id:          createdParty.Id,
		Name:        "Updated Test Party",
		Slogan:      "Updated Test Party Slogan",
		OpenedDate:  "2025-01-01",
		Description: "Updated Test Party Description",
	}
	err = stgParty.Update(context.Background(), updateParty)
	if err != nil {
		t.Fatalf("Error updating party: %v", err)
	}

	// Get the updated party
	updatedParty, err := stgParty.GetById(context.Background(), &public.PartyById{Id: createdParty.Id})
	if err != nil {
		t.Fatalf("Error getting party by ID: %v", err)
	}

	assert.Equal(t, updateParty.Name, updatedParty.Name)
	assert.Equal(t, updateParty.Slogan, updatedParty.Slogan)
	assert.Equal(t, updateParty.OpenedDate, updatedParty.OpenedDate[:10])
	assert.Equal(t, updateParty.Description, updatedParty.Description)
}

// TestDeleteParty tests the Delete method
func TestDeleteParty(t *testing.T) {
	stgParty := newTestParty(t)
	testParty := createTestParty()

	// Create a party first
	createdParty, err := stgParty.Create(context.Background(), testParty)
	if err != nil {
		t.Fatalf("Error creating party: %v", err)
	}

	// Delete the party
	err = stgParty.Delete(context.Background(), &public.PartyDelete{Id: createdParty.Id})
	if err != nil {
		t.Fatalf("Error deleting party: %v", err)
	}

	// Try to get the deleted party
	_, err = stgParty.GetById(context.Background(), &public.PartyById{Id: createdParty.Id})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "party not found")
}

// TestGetByIdParty tests the GetById method
func TestGetByIdParty(t *testing.T) {
	stgParty := newTestParty(t)
	testParty := createTestParty()

	// Create a party first
	createdParty, err := stgParty.Create(context.Background(), testParty)
	if err != nil {
		t.Fatalf("Error creating party: %v", err)
	}

	// Get the party by ID
	retrievedParty, err := stgParty.GetById(context.Background(), &public.PartyById{Id: createdParty.Id})
	if err != nil {
		t.Fatalf("Error getting party by ID: %v", err)
	}

	assert.Equal(t, createdParty.Id, retrievedParty.Id)
	assert.Equal(t, createdParty.Name, retrievedParty.Name)
	assert.Equal(t, createdParty.Slogan, retrievedParty.Slogan)
	assert.Equal(t, createdParty.OpenedDate, retrievedParty.OpenedDate)
	assert.Equal(t, createdParty.Description, retrievedParty.Description)
}

// TestGetAllParties tests the GetAll method
func TestGetAllParties(t *testing.T) {
	stgParty := newTestParty(t)

	// Create a few parties
	for i := 0; i < 3; i++ {

		_, err := stgParty.Create(context.Background(), createTestParty())
		if err != nil {
			t.Fatalf("Error creating party: %v", err)
		}
	}

	// Get all parties
	allParties, err := stgParty.GetAll(context.Background(), &public.GetAllPartyRequest{})
	if err != nil {
		t.Fatalf("Error getting all parties: %v", err)
	}
	assert.GreaterOrEqual(t, len(allParties.Parties), 3)

	// Get parties with specific filters
	filteredParties, err := stgParty.GetAll(context.Background(), &public.GetAllPartyRequest{
		Name: "Test Party",
	})

	if err != nil {
		t.Fatalf("Error getting filtered parties: %v", err)
	}
	assert.GreaterOrEqual(t, len(filteredParties.Parties), 1)
}
