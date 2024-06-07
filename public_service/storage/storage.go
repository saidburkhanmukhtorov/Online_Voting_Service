package storage

import (
	"context"

	public "github.com/myfirstgo/online_voting_service/public_service/genproto"
)

type StorageI interface {
	Party() PartyI
	Public() PublicI
}
type PartyI interface {
	Create(ctx context.Context, partyReq *public.PartyCreate) (*public.Void, error)
	Update(ctx context.Context, partyReq *public.PartyUpdate) (*public.Void, error)
	Delete(ctx context.Context, partyReq *public.PartyDelete) (*public.Void, error)
	GetById(ctx context.Context, partyReq *public.PartyById) (*public.Party, error)
	GetAll(ctx context.Context, partyReq *public.GetAllPartyRequest) (*public.GetAllPartyResponse, error)
}

type PublicI interface {
	Create(ctx context.Context, publicReq *public.PublicCreate) (*public.Void, error)
	Update(ctx context.Context, publicReq *public.PublicUpdate) (*public.Void, error)
	Delete(ctx context.Context, publicReq *public.PublicDelete) (*public.Void, error)
	GetById(ctx context.Context, publicReq *public.PublicById) (*public.Public, error)
	GetAll(ctx context.Context, publicReq *public.GetAllPublicReq) (*public.GetAllPublicRes, error)
}
