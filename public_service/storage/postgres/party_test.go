package postgres

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	public "github.com/myfirstgo/online_voting_service/public_service/genproto"
	"github.com/stretchr/testify/assert"
)

// Create a test database connection poolCREATE TYPE gender AS ENUM ('m', 'f');

func NewTestParty(t *testing.T) *PartyDb {

	connString := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		"postgres",  // Replace with your database username
		"0101",      // Replace with your database password
		"localhost", // Replace with your database host
		5432,        // Replace with your database port
		"public",    // Replace with your database name
	)

	db, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	return &PartyDb{Db: db}
}

func createTestParty() *public.PartyCreate {

	return &public.PartyCreate{
		Id:          uuid.NewString(),
		Name:        "test",
		Slogan:      "test",
		OpenedDate:  "2022-01-01",
		Description: "test",
	}
}

func TestCreateParty(t *testing.T) {
	stgParty := NewTestParty(t)
	defer stgParty.Db.Close(context.Background())
	testParty := createTestParty()

	partyRes, err := stgParty.Create(context.TODO(), testParty)

	if err != nil {
		t.Fatal("Failed to create party: ", err)
	}
	assert.NotEqual(t, partyRes, nil)
	assert.Equal(t, partyRes.Name, testParty.Name)
	assert.Equal(t, partyRes.Slogan, testParty.Slogan)
	assert.Equal(t, partyRes.OpenedDate, testParty.OpenedDate)
	assert.Equal(t, partyRes.Description, testParty.Description)

}
func TestGetPartyById(t *testing.T) {
	stgParty := NewTestParty(t)

	testParty := createTestParty()
	partyRes, err := stgParty.Create(context.TODO(), testParty)
	if err != nil {
		t.Fatal("Failed to create party: ", err)
	}
	partyRes, err = stgParty.GetById(context.TODO(), &public.PartyById{Id: partyRes.Id})
	if err != nil {
		t.Fatal("Failed to get party: ", err)
	}
	assert.NotEqual(t, partyRes, nil)
	assert.Equal(t, partyRes.Name, testParty.Name)
	assert.Equal(t, partyRes.Slogan, testParty.Slogan)
	assert.Equal(t, partyRes.OpenedDate, testParty.OpenedDate)
	assert.Equal(t, partyRes.Description, testParty.Description)
}

func TestDeleteParty(t *testing.T) {
	stgParty := NewTestParty(t)

	testParty := createTestParty()
	partyRes, err := stgParty.Create(context.TODO(), testParty)
	if err != nil {
		t.Fatal("Failed to create party: ", err)
	}
	err = stgParty.Delete(context.TODO(), &public.PartyDelete{Id: partyRes.Id})
	if err != nil {
		t.Fatal("Failed to delete party: ", err)
	}

	partyRes, err = stgParty.GetById(context.TODO(), &public.PartyById{Id: partyRes.Id})

	if err == nil {
		if err == pgx.ErrNoRows {
			return
		}
		t.Fatal("Unexpected error when getting party: ", err)
	}
	assert.Contains(t, err.Error(), "party not found")
}

func TestUpdateParty(t *testing.T) {
	stgParty := NewTestParty(t)

	testParty := createTestParty()
	partyRes, err := stgParty.Create(context.TODO(), testParty)
	if err != nil {
		t.Fatal("Failed to create party: ", err)
	}
	updateParty := &public.PartyUpdate{
		Id:          partyRes.Id,
		Name:        "test2",
		Slogan:      "test2",
		OpenedDate:  "2022-01-01",
		Description: "qqqqqqqqqqqqqqqt2",
	}
	err = stgParty.Update(context.TODO(), updateParty)

	if err != nil {
		t.Fatal("Failed to update party: ", err)
	}
	stgParty.Db.Close(context.Background())

	stgParty = NewTestParty(t)
	updateRes, err := stgParty.GetById(context.Background(), &public.PartyById{Id: partyRes.Id})

	if err != nil {
		if err == pgx.ErrNoRows {
			return
		}
		t.Fatal("Unexpected error when getting party: ", err)
	}
	stgParty.Db.Close(context.Background())
	assert.NotEqual(t, testParty, updateRes)
	assert.Equal(t, updateRes.Name, updateParty.Name)
	assert.Equal(t, updateRes.Slogan, updateParty.Slogan)
	assert.Equal(t, updateRes.OpenedDate, updateParty.OpenedDate)
	assert.Equal(t, updateRes.Description, updateParty.Description)
}

func TestGetAllParties(t *testing.T) {
	stgParty := NewTestParty(t)

	for i := 0; i < 10; i++ {
		testParty := &public.PartyCreate{
			Id:          uuid.NewString(),
			Name:        "test",
			Slogan:      "test",
			OpenedDate:  "2022-01-01",
			Description: "test",
		}
		_, err := stgParty.Create(context.TODO(), testParty)
		if err != nil {
			t.Fatal("Failed to create party: ", err)
		}
	}
	partyRes, err := stgParty.GetAll(context.TODO(), &public.GetAllPartyRequest{})
	if err != nil {
		t.Fatal("Failed to get all parties: ", err)
	}
	
	assert.GreaterOrEqual(t, len(partyRes.Parties), 9)
}
