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

// PublicDb provides database operation for publics
type PublicDb struct {
	Db *pgx.Conn
}

// NewPublic creates a new instance of PublicDb
func NewPublic(db *pgx.Conn) *PublicDb {
	return &PublicDb{Db: db}
}

// Create creates a new public in the database
func (pub *PublicDb) Create(ctx context.Context, publicReq *public.PublicCreate) (*public.Public, error) {
	// Generate a new UUID for the public if id is not exists
	if len(publicReq.Id) == 0 {
		publicReq.Id = uuid.New().String()
	}
	log.Println(publicReq.Id)
	// Insert into public table
	publicQuery := `
        INSERT INTO public (
			id, 
			name, 
			last_name, 
			phone, 
			email, 
			birthday, 
			gender, 
			party_id)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
        RETURNING 
			id, 
			name, 
			last_name, 
			phone, 
			email, 
			birthday, 
			gender, 
			party_id, 
			created_at, 
			updated_at, 
			deleted_at
    `
	var publicRes public.Public
	var createdAt, updatedAt, birthday time.Time
	birthday, err := parseDate(publicReq.Birthday)
	if err != nil {
		return nil, err
	}
	err = pub.Db.QueryRow(ctx, publicQuery,
		publicReq.Id,
		publicReq.Name,
		publicReq.LastName,
		publicReq.Phone,
		publicReq.Email,
		birthday,
		publicReq.Gender,
		publicReq.PartyId,
	).Scan(
		&publicRes.Id,
		&publicRes.Name,
		&publicRes.LastName,
		&publicRes.Phone,
		&publicRes.Email,
		&birthday,
		&publicRes.Gender,
		&publicRes.PartyId,
		&createdAt,
		&updatedAt,
		&publicRes.DeletedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to insert into public: %w", err)
	}
	publicRes.Birthday = birthday.Format(time.RFC3339)
	publicRes.CreatedAt = createdAt.Format(time.RFC3339)
	publicRes.UpdatedAt = updatedAt.Format(time.RFC3339)

	return &publicRes, nil
}

// Update updates an existing public in the database
func (pub *PublicDb) Update(ctx context.Context, publicReq *public.PublicUpdate) error {
	// Parse the birthday string to time.Time
	birthday, err := parseDate(publicReq.Birthday)
	if err != nil {
		return err
	}
	updateQuery := `
		UPDATE public
		SET name = $2, last_name = $3, phone = $4, email = $5, birthday = $6, gender = $7, party_id = $8, updated_at = NOW()
		WHERE id = $1
	`

	_, err = pub.Db.Exec(ctx, updateQuery,
		publicReq.Id,
		publicReq.Name,
		publicReq.LastName,
		publicReq.Phone,
		publicReq.Email,
		birthday,
		publicReq.Gender,
		publicReq.PartyId,
	)
	if err != nil {
		return fmt.Errorf("failed to update public: %w", err)
	}

	return nil
}

// Delete deletes a public from the database
func (pub *PublicDb) Delete(ctx context.Context, publicReq *public.PublicDelete) error {
	deleteQuery := `
		UPDATE public
		SET deleted_at = $1, updated_at = NOW()
		WHERE id = $2
	`
	_, err := pub.Db.Exec(ctx, deleteQuery, time.Now().Unix(), publicReq.Id)
	if err != nil {
		return fmt.Errorf("failed to delete public: %w", err)
	}

	return nil
}

// GetById retrieves a public by its ID from the database
func (pub *PublicDb) GetById(ctx context.Context, publicReq *public.PublicById) (*public.Public, error) {
	var publicRes public.Public
	var createdAt, updatedAt, birthday time.Time

	query := `
		SELECT 
			id, 
			name, 
			last_name, 
			phone, 
			email, 
			birthday, 
			gender, 
			party_id, 
			created_at, 
			updated_at, 
			deleted_at
		FROM 
			public
		WHERE 
			id = $1
		AND 
			deleted_at = 0
	`

	err := pub.Db.QueryRow(ctx, query, publicReq.Id).Scan(
		&publicRes.Id,
		&publicRes.Name,
		&publicRes.LastName,
		&publicRes.Phone,
		&publicRes.Email,
		&birthday,
		&publicRes.Gender,
		&publicRes.PartyId,
		&createdAt,
		&updatedAt,
		&publicRes.DeletedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("public not found for ID: %s", publicReq.Id)
		}
		return nil, fmt.Errorf("failed to fetch public by ID: %w", err)
	}
	publicRes.Birthday = birthday.Format(time.RFC3339)
	publicRes.CreatedAt = createdAt.Format(time.RFC3339)
	publicRes.UpdatedAt = updatedAt.Format(time.RFC3339)

	return &publicRes, nil
}

// GetAll retrieves all publics from the database based on optional filters
func (pub *PublicDb) GetAll(ctx context.Context, publicReq *public.GetAllPublicReq) (*public.GetAllPublicRes, error) {
	var publicResAll public.GetAllPublicRes
	// publicRes.Publics = []*public.Public{}
	query := `
		SELECT id, name, last_name, phone, email, birthday, gender, party_id, created_at, updated_at, deleted_at
		FROM public
	`

	// filter adds filter to the query
	// In here I use key and params to avoid sql injection
	params := make(map[string]interface{}, 0)
	var arr []interface{}
	filter := `WHERE deleted_at = 0`

	if publicReq.PartyId != "" {
		params["party_id"] = publicReq.PartyId
		filter += " AND party_id = :party_id"
	}

	if publicReq.Name != "" {
		params["name"] = publicReq.Name
		filter += " AND name = :name"
	}

	if publicReq.LastName != "" {
		params["last_name"] = publicReq.LastName
		filter += " AND last_name = :last_name"
	}

	if publicReq.Phone != "" {
		params["phone"] = publicReq.Phone
		filter += " AND phone = :phone"
	}

	if publicReq.Email != "" {
		params["email"] = publicReq.Email
		filter += " AND email = :email"
	}

	if publicReq.Birthday != "" {
		birthday, err := parseDate(publicReq.Birthday)
		if err != nil {
			return nil, err
		}
		params["birthday"] = birthday
		filter += " AND birthday = :birthday"
	}

	if publicReq.Gender != "" {
		params["gender"] = publicReq.Gender
		filter += " AND gender = :gender"
	}

	query += filter
	query, arr = helper.ReplaceQueryParams(query, params) // Assuming you have a helper.ReplaceQueryParams function
	rows, err := pub.Db.Query(ctx, query, arr...)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch publics: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var publicRes public.Public
		var createdAt, updatedAt, birthday time.Time

		err := rows.Scan(
			&publicRes.Id,
			&publicRes.Name,
			&publicRes.LastName,
			&publicRes.Phone,
			&publicRes.Email,
			&birthday,
			&publicRes.Gender,
			&publicRes.PartyId,
			&createdAt,
			&updatedAt,
			&publicRes.DeletedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan public row: %w", err)
		}
		publicRes.Birthday = birthday.Format(time.RFC3339)
		publicRes.CreatedAt = createdAt.Format(time.RFC3339)
		publicRes.UpdatedAt = updatedAt.Format(time.RFC3339)

		publicResAll.Publics = append(publicResAll.Publics, &publicRes)
	}

	return &publicResAll, nil
}

func (pub *PublicDb) IsValidPublic(ctx context.Context, in *public.ValidPublicReq) (*public.ValidPublicRes, error) {
	query := `
		SELECT 
			deleted_at 
		FROM 
			public 
		WHERE 
			id = $1
	`

	var deletedAt int64
	err := pub.Db.QueryRow(ctx, query, in.Id).Scan(&deletedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to check party deletion status: %w", err)
	}
	if deletedAt != 0 {
		return &public.ValidPublicRes{Valid: false}, nil
	}
	return &public.ValidPublicRes{Valid: true}, nil
}
