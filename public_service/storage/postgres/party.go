package postgres

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	public "github.com/myfirstgo/online_voting_service/public_service/genproto"
	"github.com/myfirstgo/online_voting_service/public_service/helper"
)

// PartyDb provides database operation for parties
type PartyDb struct {
	Db *pgx.Conn
}

// NewParty creates a new instance of PartyDb
func NewParty(db *pgx.Conn) *PartyDb {
	return &PartyDb{Db: db}
}

// Create creates a new party in the database
func (part *PartyDb) Create(ctx context.Context, partyReq *public.PartyCreate) (*public.Party, error) {
	// Generate a new UUID for the party if id is not exists
	if len(partyReq.Id) == 0 {
		partyReq.Id = uuid.New().String()
	}
	// Parse the opened date string to time.Time
	openedDate, err := parseDate(partyReq.OpenedDate)
	if err != nil {
		return nil, err
	}
	// Insert into party table
	partyQuery := `
        INSERT INTO party (
			id, 
			name, 
			slogan, 
			opened_date, 
			description)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING 
			id, 
			name, 
			slogan, 
			opened_date, 
			description, 
			created_at, 
			updated_at, 
			deleted_at
    `
	var partyRes public.Party
	var createdAt, updatedAt time.Time
	err = part.Db.QueryRow(ctx, partyQuery,
		partyReq.Id,
		partyReq.Name,
		partyReq.Slogan,
		openedDate,
		partyReq.Description,
	).Scan(
		&partyRes.Id,
		&partyRes.Name,
		&partyRes.Slogan,
		&openedDate,
		&partyRes.Description,
		&createdAt,
		&updatedAt,
		&partyRes.DeletedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to insert into party: %w", err)
	}
	partyRes.OpenedDate = openedDate.String()
	partyRes.CreatedAt = createdAt.Format(time.RFC3339)
	partyRes.UpdatedAt = updatedAt.Format(time.RFC3339)

	return &partyRes, nil
}

// Update updates an existing party in the database
func (part *PartyDb) Update(ctx context.Context, partyReq *public.PartyUpdate) error {
	// Parse the opened date string to time.Time
	openedDate, err := parseDate(partyReq.OpenedDate)
	if err != nil {
		return err
	}
	updateQuery := `
		UPDATE party
		SET name = $2, slogan = $3, opened_date = $4, description = $5, updated_at = NOW()
		WHERE id = $1
	`

	_, err = part.Db.Exec(ctx, updateQuery,
		partyReq.Id,
		partyReq.Name,
		partyReq.Slogan,
		openedDate,
		partyReq.Description,
	)
	if err != nil {
		return fmt.Errorf("failed to update party: %w", err)
	}

	return nil
}

// Delete deletes a party from the database
func (part *PartyDb) Delete(ctx context.Context, partyReq *public.PartyDelete) error {
	deleteQuery := `
		UPDATE party
		SET deleted_at = $1, updated_at = NOW()
		WHERE id = $2
	`
	_, err := part.Db.Exec(ctx, deleteQuery, time.Now().Unix(), partyReq.Id)
	if err != nil {
		return fmt.Errorf("failed to delete party: %w", err)
	}

	return nil
}

// GetById retrieves a party by its ID from the database
func (part *PartyDb) GetById(ctx context.Context, partyReq *public.PartyById) (*public.Party, error) {
	log.Println(partyReq.Id)
	var partyRes public.Party
	var createdAt, updatedAt, openedDate time.Time

	query := `
		SELECT 
			id, 
			name, 
			slogan, 
			opened_date, 
			description, 
			created_at, 
			updated_at, 
			deleted_at
		FROM 
			party
		WHERE 
			id = $1
		AND 
			deleted_at = 0
	`

	err := part.Db.QueryRow(ctx, query, partyReq.Id).Scan(
		&partyRes.Id,
		&partyRes.Name,
		&partyRes.Slogan,
		&openedDate,
		&partyRes.Description,
		&createdAt,
		&updatedAt,
		&partyRes.DeletedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("party not found for ID: %s", partyReq.Id)
		}
		return nil, fmt.Errorf("failed to fetch party by ID: %w", err)
	}
	partyRes.OpenedDate = openedDate.String()
	partyRes.CreatedAt = createdAt.Format(time.RFC3339)
	partyRes.UpdatedAt = updatedAt.Format(time.RFC3339)

	return &partyRes, nil
}

// GetAll retrieves all parties from the database based on optional filters
func (part *PartyDb) GetAll(ctx context.Context, partyReq *public.GetAllPartyRequest) (*public.GetAllPartyResponse, error) {
	var partyRes public.GetAllPartyResponse
	query := `
		SELECT id, name, slogan, opened_date, description, created_at, updated_at, deleted_at
		FROM party
	`

	// filter adds filter to the query
	// In here I use key and params to avoid sql injection
	params := make(map[string]interface{}, 0)
	var arr []interface{}
	filter := `WHERE deleted_at = 0`

	if partyReq.OpenedDate != "" {
		openedDate, err := parseDate(partyReq.OpenedDate)
		if err != nil {
			return nil, err
		}
		params["opened_date"] = openedDate
		filter += " AND opened_date = :opened_date"
	}

	if partyReq.Name != "" {
		params["name"] = partyReq.Name
		filter += " AND name = :name"
	}

	if partyReq.Slogan != "" {
		params["slogan"] = partyReq.Slogan
		filter += " AND slogan = :slogan"
	}

	if partyReq.Description != "" {
		params["description"] = partyReq.Description
		filter += " AND description = :description"
	}
	query += filter
	query, arr = helper.ReplaceQueryParams(query, params) // Assuming you have a helper.ReplaceQueryParams function
	rows, err := part.Db.Query(ctx, query, arr...)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch parties: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		log.Println(1)
		var party public.Party
		var createdAt, updatedAt, openedDate time.Time

		err := rows.Scan(
			&party.Id,
			&party.Name,
			&party.Slogan,
			&openedDate,
			&party.Description,
			&createdAt,
			&updatedAt,
			&party.DeletedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan party row: %w", err)
		}

		party.OpenedDate = openedDate.Format(time.RFC3339)
		party.CreatedAt = createdAt.Format(time.RFC3339)
		party.UpdatedAt = updatedAt.Format(time.RFC3339)

		partyRes.Parties = append(partyRes.Parties, &party)
	}
	return &partyRes, nil
}

func parseDate(DateString string) (time.Time, error) {
	// Define the layout for the date and time string
	layout := "2006-01-02"

	// Parse the string into a time.Time object
	date, err := time.Parse(layout, DateString)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse opened date: %w", err)
	}

	return date, nil
}

func (part *PartyDb) IsDeleted(ctx context.Context, partyId *string) (bool, error) {
	query := `
		SELECT 
			deleted_at 
		FROM 
			party 
		WHERE 
			id = $1
	`

	var deletedAt int64
	err := part.Db.QueryRow(ctx, query, partyId).Scan(&deletedAt)
	if err != nil {
		return false, fmt.Errorf("failed to check party deletion status: %w", err)
	}

	return deletedAt != 0, nil // If deleted_at is not 0, then it's deleted
}
