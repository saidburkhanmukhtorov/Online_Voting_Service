package service

import (
	"context"
	"log/slog"

	pub "github.com/myfirstgo/online_voting_service/public_service/genproto"
	"github.com/myfirstgo/online_voting_service/public_service/storage/postgres"
)

type Service struct {
	stg *postgres.Storage
	pub.UnimplementedPartyServiceServer
}

func NewPartyService(stg *postgres.Storage) *Service {
	return &Service{stg: stg}
}

func (s *Service) Create(ctx context.Context, partyReq *pub.PartyCreate) (*pub.Void, error) {
	slog.Info("CreateParty Service", "Party", partyReq)
	_, err := s.stg.PartyS.Create(ctx, partyReq)
	return nil, err
}

func (s *Service) Update(ctx context.Context, partyReq *pub.PartyUpdate) (*pub.Void, error) {
	slog.Info("UpdateParty Service", "Party", partyReq)
    _, err := s.stg.PartyS.Update(ctx, partyReq)
    return nil, err
}

func (s *Service) Delete(ctx context.Context, partyReq *pub.PartyDelete) (*pub.Void, error) {
	slog.Info("DeleteParty Service", "Party", partyReq)
    _, err := s.stg.PartyS.Delete(ctx, partyReq)
    return nil, err
}


func (s *Service) GetById(ctx context.Context, partyReq *pub.PartyById) (*pub.Party, error) {
	slog.Info("GetByIdParty Service", "Party", partyReq)
    party, err := s.stg.PartyS.GetById(ctx, partyReq)
    return party, err
}


func (s *Service) GetAll(ctx context.Context, partyReq *pub.GetAllPartyRequest) (*pub.GetAllPartyResponse, error) {
	slog.Info("GetAllParty Service", "Party", partyReq)
    party, err := s.stg.PartyS.GetAll(ctx, partyReq)
    return party, err
}
