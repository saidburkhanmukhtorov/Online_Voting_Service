// party_service.go
package service

import (
	"context"
	"log/slog"

	public "github.com/myfirstgo/online_voting_service/public_service/genproto"
	"github.com/myfirstgo/online_voting_service/public_service/storage/postgres"
)

// PartyService implements the PartyServiceServer interface
type PartyService struct {
	stg *postgres.Storage
	public.UnimplementedPartyServiceServer
}

// NewPartyService creates a new PartyService instance
func NewPartyService(stg *postgres.Storage) *PartyService {
	return &PartyService{stg: stg}
}

// Create creates a new party
func (p *PartyService) Create(ctx context.Context, partyReq *public.PartyCreate) (*public.Party, error) {
	slog.Info("CreateParty Service", "party", partyReq)
	partyRes, err := p.stg.PartyS.Create(ctx, partyReq)
	if err != nil {
		slog.Error("Error creating party", "err", err)
		return nil, err
	}

	return partyRes, nil
}

// Update updates an existing party
func (p *PartyService) Update(ctx context.Context, partyReq *public.PartyUpdate) (*public.Void, error) {
	slog.Info("UpdateParty Service", "party", partyReq)
	err := p.stg.PartyS.Update(ctx, partyReq)
	if err != nil {
		slog.Error("Error updating party", "err", err)
		return &public.Void{}, err
	}

	return &public.Void{}, nil
}

// Delete deletes a party
func (p *PartyService) Delete(ctx context.Context, partyReq *public.PartyDelete) (*public.Void, error) {
	slog.Info("DeleteParty Service", "party id", partyReq.Id)
	err := p.stg.PartyS.Delete(ctx, partyReq)
	if err != nil {
		slog.Error("Error deleting party", "err", err)
		return &public.Void{}, err
	}

	return &public.Void{}, nil
}

// GetById retrieves a party by its ID
func (p *PartyService) GetById(ctx context.Context, partyReq *public.PartyById) (*public.Party, error) {
	slog.Info("GetPartyById Service", "party id", partyReq.Id)
	partyRes, err := p.stg.PartyS.GetById(ctx, partyReq)
	if err != nil {
		slog.Error("Error getting party by ID", "err", err)
		return nil, err
	}

	return partyRes, nil
}

// GetAll retrieves all parties
func (p *PartyService) GetAll(ctx context.Context, partyReq *public.GetAllPartyRequest) (*public.GetAllPartyResponse, error) {
	slog.Info("GetAllParty Service", "party req", partyReq)

	partyRes, err := p.stg.PartyS.GetAll(ctx, partyReq)
	if err != nil {
		slog.Error("Error getting all parties", "err", err)
		return nil, err
	}

	return partyRes, nil
}
