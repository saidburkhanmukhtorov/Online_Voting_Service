package postgres

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	public "github.com/myfirstgo/online_voting_service/public_service/genproto"
	"github.com/myfirstgo/online_voting_service/public_service/storage/postgres"
	"github.com/stretchr/testify/assert"
)

// Create a test database connection pool for PublicDb
func newTestPublic(t *testing.T) *postgres.PublicDb {
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
	return &postgres.PublicDb{Db: db}
}

// Create a test public object
func createTestPublic() *public.PublicCreate {
	return &public.PublicCreate{
		Id:       uuid.New().String(),
		Name:     "Test Public",
		LastName: "Public",
		Phone:    "1" + strconv.Itoa(time.Now().Nanosecond()),
		Email:    "test.public@example.com" + strconv.Itoa(time.Now().Nanosecond()),
		Birthday: "2000-01-01",
		Gender:   "m",
		PartyId:  "0a8aeaef-37f5-42ff-9002-e331fa686188",
	}
}

// TestCreatePublic tests the Create method
func TestCreatePublic(t *testing.T) {
	stgPublic := newTestPublic(t)
	testPublic := createTestPublic()

	publicRes, err := stgPublic.Create(context.Background(), testPublic)
	if err != nil {
		t.Fatalf("Error creating public: %v", err)
	}

	assert.NotEmpty(t, publicRes.Id)
	assert.Equal(t, testPublic.Name, publicRes.Name)
	assert.Equal(t, testPublic.LastName, publicRes.LastName)
	assert.Equal(t, testPublic.Phone, publicRes.Phone)
	assert.Equal(t, testPublic.Email, publicRes.Email)
	assert.Equal(t, testPublic.Birthday[:10], publicRes.Birthday[:10]) // Compare only date part
	assert.Equal(t, testPublic.Gender, publicRes.Gender)
	assert.Equal(t, testPublic.PartyId, publicRes.PartyId)
}

// TestUpdatePublic tests the Update method
func TestUpdatePublic(t *testing.T) {
	stgPublic := newTestPublic(t)
	testPublic := createTestPublic()

	// Create a public first
	createdPublic, err := stgPublic.Create(context.Background(), testPublic)
	if err != nil {
		t.Fatalf("Error creating public: %v", err)
	}

	// Update the public
	updatePublic := &public.PublicUpdate{
		Id:       createdPublic.Id,
		Name:     "Updated Test Public",
		LastName: "Public",
		Phone:    "1" + strconv.Itoa(time.Now().Nanosecond()),
		Email:    "test.public@example.com" + strconv.Itoa(time.Now().Nanosecond()),
		Birthday: "2001-02-02",
		Gender:   "f",
		PartyId:  "9dbc90ca-893e-4233-86cf-bee0cdc18353", // Replace with a valid party ID from your database
	}
	err = stgPublic.Update(context.Background(), updatePublic)
	if err != nil {
		t.Fatalf("Error updating public: %v", err)
	}

	// Get the updated public
	updatedPublic, err := stgPublic.GetById(context.Background(), &public.PublicById{Id: createdPublic.Id})
	if err != nil {
		t.Fatalf("Error getting public by ID: %v", err)
	}

	assert.Equal(t, updatePublic.Name, updatedPublic.Name)
	assert.Equal(t, updatePublic.LastName, updatedPublic.LastName)
	assert.Equal(t, updatePublic.Phone, updatedPublic.Phone)
	assert.Equal(t, updatePublic.Email, updatedPublic.Email)
	assert.Equal(t, updatePublic.Birthday[:10], updatedPublic.Birthday[:10]) // Compare only date part
	assert.Equal(t, updatePublic.Gender, updatedPublic.Gender)
	assert.Equal(t, updatePublic.PartyId, updatedPublic.PartyId)
}

// TestDeletePublic tests the Delete method
func TestDeletePublic(t *testing.T) {
	stgPublic := newTestPublic(t)
	testPublic := createTestPublic()

	// Create a public first
	createdPublic, err := stgPublic.Create(context.Background(), testPublic)
	if err != nil {
		t.Fatalf("Error creating public: %v", err)
	}

	// Delete the public
	err = stgPublic.Delete(context.Background(), &public.PublicDelete{Id: createdPublic.Id})
	if err != nil {
		t.Fatalf("Error deleting public: %v", err)
	}

	// Try to get the deleted public
	_, err = stgPublic.GetById(context.Background(), &public.PublicById{Id: createdPublic.Id})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "public not found")
}

// TestGetByIdPublic tests the GetById method
func TestGetByIdPublic(t *testing.T) {
	stgPublic := newTestPublic(t)
	testPublic := createTestPublic()

	// Create a public first
	createdPublic, err := stgPublic.Create(context.Background(), testPublic)
	if err != nil {
		t.Fatalf("Error creating public: %v", err)
	}

	// Get the public by ID
	retrievedPublic, err := stgPublic.GetById(context.Background(), &public.PublicById{Id: createdPublic.Id})
	if err != nil {
		t.Fatalf("Error getting public by ID: %v", err)
	}

	assert.Equal(t, createdPublic.Id, retrievedPublic.Id)
	assert.Equal(t, createdPublic.Name, retrievedPublic.Name)
	assert.Equal(t, createdPublic.LastName, retrievedPublic.LastName)
	assert.Equal(t, createdPublic.Phone, retrievedPublic.Phone)
	assert.Equal(t, createdPublic.Email, retrievedPublic.Email)
	assert.Equal(t, createdPublic.Birthday, retrievedPublic.Birthday)
	assert.Equal(t, createdPublic.Gender, retrievedPublic.Gender)
	assert.Equal(t, createdPublic.PartyId, retrievedPublic.PartyId)
}

// TestGetAllPublics tests the GetAll method
func TestGetAllPublics(t *testing.T) {
	stgPublic := newTestPublic(t)

	// Create a few publics
	for i := 0; i < 3; i++ {
		testPublic := createTestPublic()
		_, err := stgPublic.Create(context.Background(), testPublic)
		if err != nil {
			t.Fatalf("Error creating public: %v", err)
		}
	}

	// Get all publics
	allPublics, err := stgPublic.GetAll(context.Background(), &public.GetAllPublicReq{})
	if err != nil {
		t.Fatalf("Error getting all publics: %v", err)
	}
	assert.GreaterOrEqual(t, len(allPublics.Publics), 3)

	// Get publics with specific filters
	filteredPublics, err := stgPublic.GetAll(context.Background(), &public.GetAllPublicReq{
		Name: "Test Public",
	})
	if err != nil {
		t.Fatalf("Error getting filtered publics: %v", err)
	}
	assert.GreaterOrEqual(t, len(filteredPublics.Publics), 1)
}
