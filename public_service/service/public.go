package service

import (
	"context"
	"log/slog"

	pub "github.com/myfirstgo/online_voting_service/public_service/genproto"
	"github.com/myfirstgo/online_voting_service/public_service/storage/postgres"
)

type publicService struct {
	stg *postgres.Storage
	pub.UnimplementedPublicServiceServer
}

func NewPublicService(stg *postgres.Storage) *publicService {
	return &publicService{stg: stg}
}

func (p *publicService) Create(ctx context.Context, partyReq *pub.PublicCreate) (*pub.Void, error) {
	slog.Info("CreateParty publicService", "Public", partyReq)
	_, err := p.stg.PublicS.Create(ctx, partyReq)
	return nil, err
}

func (p *publicService) Update(ctx context.Context, partyReq *pub.PublicUpdate) (*pub.Void, error) {
	slog.Info("UpdateParty publicService", "Public", partyReq)
    _, err := p.stg.PublicS.Update(ctx, partyReq)
    return nil, err
}

func (p *publicService) Delete(ctx context.Context, partyReq *pub.PublicDelete) (*pub.Void, error) {
	slog.Info("DeleteParty publicService", "Public", partyReq)
    _, err := p.stg.PublicS.Delete(ctx, partyReq)
    return nil, err
}

func (p *publicService) GetById(ctx context.Context, partyReq *pub.PublicById) (*pub.Public, error) {
    slog.Info("GetByIdParty publicService", "Public", partyReq)
    return p.stg.PublicS.GetById(ctx, partyReq)
}


func (p *publicService) GetAll(ctx context.Context, partyReq *pub.GetAllPublicReq) (*pub.GetAllPublicRes, error) {
    slog.Info("GetAllParty publicService", "Public", partyReq)
    return p.stg.PublicS.GetAll(ctx, partyReq)
}