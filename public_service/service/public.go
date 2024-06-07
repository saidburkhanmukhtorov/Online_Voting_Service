// public_service.go
package service

import (
	"context"
	"fmt"
	"log/slog"

	public "github.com/myfirstgo/online_voting_service/public_service/genproto"
	"github.com/myfirstgo/online_voting_service/public_service/storage/postgres"
)

// PublicService implements the PublicServiceServer interface
type PublicService struct {
	stg *postgres.Storage
	public.UnimplementedPublicServiceServer
}

// NewPublicService creates a new PublicService instance
func NewPublicService(stg *postgres.Storage) *PublicService {
	return &PublicService{stg: stg}
}

// Create creates a new public
func (p *PublicService) Create(ctx context.Context, publicReq *public.PublicCreate) (*public.Public, error) {

	isValid, err := p.stg.PartyS.IsDeleted(ctx, &publicReq.PartyId)
	if err != nil || isValid {
		return nil, err
	}
	slog.Info("CreatePublic Service", "public", publicReq)
	publicRes, err := p.stg.PublicS.Create(ctx, publicReq)
	if err != nil {
		slog.Error("Error creating public", "err", err)
		return nil, err
	}
	return publicRes, nil
}

// Update updates an existing public
func (p *PublicService) Update(ctx context.Context, publicReq *public.PublicUpdate) (*public.Void, error) {

	isValid, err := p.stg.PartyS.IsDeleted(ctx, &publicReq.PartyId)
	if err != nil || isValid {
		return nil, err
	}
	slog.Info("UpdatePublic Service", "public", publicReq)
	err = p.stg.PublicS.Update(ctx, publicReq)
	if err != nil {
		slog.Error("Error updating public", "err", err)
		return &public.Void{}, err
	}

	return &public.Void{}, nil
}

// Delete deletes a public
func (p *PublicService) Delete(ctx context.Context, publicReq *public.PublicDelete) (*public.Void, error) {
	slog.Info("DeletePublic Service", "public id", publicReq.Id)
	err := p.stg.PublicS.Delete(ctx, publicReq)
	if err != nil {
		slog.Error("Error deleting public", "err", err)
		return &public.Void{}, err
	}

	return &public.Void{}, nil
}

// GetById retrieves a public by its ID
func (p *PublicService) GetById(ctx context.Context, publicReq *public.PublicById) (*public.Public, error) {
	slog.Info("GetPublicById Service", "public id", publicReq.Id)
	publicRes, err := p.stg.PublicS.GetById(ctx, publicReq)
	if err != nil {
		slog.Error("Error getting public by ID", "err", err)
		return nil, err
	}

	return publicRes, nil
}

// GetAll retrieves all publics
func (p *PublicService) GetAll(ctx context.Context, publicReq *public.GetAllPublicReq) (*public.GetAllPublicRes, error) {
	slog.Info("GetAllPublic Service", "public req", publicReq)
	publicRes, err := p.stg.PublicS.GetAll(ctx, publicReq)
	if err != nil {
		slog.Error("Error getting all publics", "err", err)
		return nil, err
	}

	return publicRes, nil
}

func (s *PublicService) IsValidPublic(ctx context.Context, req *public.ValidPublicReq) (*public.ValidPublicRes, error) {
	isValid, err := s.stg.PublicS.IsValidPublic(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to check public record validity: %w", err)
	}
	return isValid, nil
}
