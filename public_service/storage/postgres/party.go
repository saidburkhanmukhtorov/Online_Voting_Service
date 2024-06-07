package postgres

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	public "github.com/myfirstgo/online_voting_service/public_service/genproto"
)

// PartyDb provides database operation for candidates

var ErrPartyNotFound = errors.New("party not found")

type PartyDb struct {
	Db *pgx.Conn
}

// NewCandidate creates a new instance of PartyDb
func NewParty(db *pgx.Conn) *PartyDb {
	return &PartyDb{Db: db}
}

func (part *PartyDb) Create(ctx context.Context, partyReq *public.PartyCreate) (*public.Party, error) {
	partyID := uuid.New().String()

	openDate, err := time.Parse("2006-01-02", partyReq.OpenedDate)
	if err != nil {
		slog.Error("Unable to parse date", err)
		return nil, err
	}

	query := `
		INSERT INTO
			parties (
				id,
				name,
				slogan,
				opened_date,
				description
			)
			VALUES (
				$1,
				$2,
				$3,
				$4,
				$5
			)
			RETURNING
				id,
				name,
				slogan,
				opened_date,
				description
	`

	var openedDate time.Time
	res := public.Party{}
	if err := part.Db.QueryRow(
		ctx,
		query,
		partyID,
		partyReq.Name,
		partyReq.Slogan,
		openDate,
		partyReq.Description,
	).Scan(
		&res.Id,
		&res.Name,
		&res.Slogan,
		&openedDate,
		&res.Description,
	); err != nil {
		return nil, err
	}
	res.OpenedDate = openedDate.Format("2006-01-02")

	return &res, nil
}

func (part *PartyDb) Update(ctx context.Context, partyReq *public.PartyUpdate) error {
	var args []interface{}
	count := 1
	query := `
		UPDATE	
			parties
		SET`
	filter := ``

	if len(partyReq.Name) > 0 {
		filter += fmt.Sprintf(" name = $%d,", count)
		args = append(args, partyReq.Name)
		count++

	}

	if len(partyReq.Slogan) > 0 {
		filter += fmt.Sprintf(" slogan = $%d,", count)
		args = append(args, partyReq.Slogan)
		count++
	}
	if len(partyReq.OpenedDate) > 0 {
		openeDate, err := time.Parse("2006-01-02", partyReq.OpenedDate)
		if err != nil {
			slog.Error("Unable to parse date", err)
			return err
		}
		filter += fmt.Sprintf(" opened_date = $%d,", count)
		args = append(args, openeDate.String()[:10])
		count++
	}

	if len(partyReq.Description) > 0 {
		filter += fmt.Sprintf(" description = $%d,", count)
		args = append(args, partyReq.Description)
		count++
	}
	query += filter
	query += " updated_at = now()"
	query += fmt.Sprintf(" WHERE id = $%d", count)
	args = append(args, partyReq.Id)
	fmt.Println(query)
	_, err := part.Db.Query(ctx, query, args...)
	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Error("Party not found", err)
			return ErrPartyNotFound
		}
		slog.Error("Unable to update party", err)
		return err
	}

	return nil
}

func (part *PartyDb) Delete(ctx context.Context, partyReq *public.PartyDelete) error {
	query := `
		UPDATE
			parties
		SET
			deleted_at = $1
		WHERE	
			id = $2
		`
	_, err := part.Db.Exec(ctx, query, time.Now().Unix(), partyReq.Id)
	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Error("Party not found", err)
			return ErrPartyNotFound
		}
		slog.Error("Unable to delete party", err)
		return err
	}

	return nil
}
func (PartyDb *PartyDb) GetById(ctx context.Context, id *public.PartyById) (*public.Party, error) {
	var (
		party       public.Party
		openDate    time.Time
		slogan      string
		description string
		creates     time.Time
		updated     time.Time
	)

	query := `
		SELECT
		id,
		name,
		slogan,
		opened_date,
		description,
		created_at,
		updated_at
		FROM
		parties
		WHERE
		deleted_at = 0
		AND
		id = $1
		`
	err := PartyDb.Db.QueryRow(ctx, query, id.Id).Scan(
		&party.Id,
		&party.Name,
		&slogan,
		&openDate,
		&description,
		&creates,
		&updated,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Error("Party not found", err)
			return nil, ErrPartyNotFound
		}
		log.Println(1)
		slog.Error("Unable to get party", err)
		return nil, err
	}

	party.CreatedAt = creates.String()
	party.UpdatedAt = updated.String()
	party.OpenedDate = openDate.String()[:10]
	party.Slogan = slogan
	party.Description = description
	return &party, nil
}

func (part *PartyDb) GetAll(ctx context.Context, partyReq *public.GetAllPartyRequest) (*public.GetAllPartyResponse, error) {
	var partyRes public.GetAllPartyResponse

	query := `
		SELECT
			id,
			name,
			slogan,
			opened_date,
			description
		FROM
			parties
		WHERE
			deleted_at = 0	
		`

	params := make(map[string]interface{}, 0)

	var arr []interface{}

	if len(partyReq.Name) > 0 {
		query += " AND name = :name"
		params["name"] = partyReq.Name
	}

	if len(partyReq.Slogan) > 0 {
		query += " AND slogan = :slogan"
		params["slogan"] = partyReq.Slogan
	}
	if len(partyReq.OpenedDate) > 0 {
		openDate, err := time.Parse("2006-01-02", partyReq.OpenedDate)
		if err != nil {
			slog.Error("Unable to parse date", err)
			return nil, err
		}
		query += " AND opened_date = :opened_date"
		params["opened_date"] = openDate
	}
	if len(partyReq.Description) > 0 {
		query += " AND description = :description"
		params["description"] = partyReq.Description
	}
	query += "ORDER BY name"
	rows, err := part.Db.Query(ctx, query, arr...)
	if err != nil {
		if err == pgx.ErrNoRows {
			slog.Error("Party not found", err)
			return nil, ErrPartyNotFound
		}
		slog.Error("Unable to get party", err)
		return nil, err
	}
	for rows.Next() {
		var (
			id          string
			name        string
			slogan      string
			opened_date time.Time
			description string
			party       public.Party
		)
		err := rows.Scan(&id, &name, &slogan, &opened_date, &description)
		if err != nil {
			slog.Error("Unable to get party", err)
			return nil, err
		}
		party.Id = id
		party.Name = name
		party.Slogan = slogan
		party.OpenedDate = opened_date.String()
		party.Description = description
		partyRes.Parties = append(partyRes.Parties, &party)
	}

	return &partyRes, nil
}
