package service

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"

	// "github.com/gogo/protobuf/types"
	"database/sql"
	// biz "pes/common"
	pb "github.com/1token/email-services/email-apis/generated/go"
)

const (
	draftTable = "draft"
)

type DraftServerImpl struct {
	DB *sql.DB
}

func (s *DraftServerImpl) GetDraft(ctx context.Context, in *pb.GetDraftRequest) (*pb.Draft, error) {
	draft := &pb.Draft{}
	return draft, nil
}

func (s *DraftServerImpl) ListDrafts(ctx context.Context, in *pb.ListDraftsRequest) (*pb.ListDraftsResponse, error) {
	drafts := &pb.ListDraftsResponse{}
	return drafts, nil
}

func (s *DraftServerImpl) CreateDraft(ctx context.Context, in *pb.CreateDraftRequest) (*pb.Draft, error) {
	draft := &pb.Draft{}
	return draft, nil
}

func (s *DraftServerImpl) UpdateDraft(ctx context.Context, in *pb.UpdateDraftRequest) (*pb.Draft, error) {
	draft := &pb.Draft{}
	return draft, nil
}

func (s *DraftServerImpl) DeleteDraft(ctx context.Context, in *pb.DeleteDraftRequest) (*empty.Empty, error) {
	return nil, nil
}

func (s *DraftServerImpl) SendDraft(ctx context.Context, in *pb.SendDraftRequest) (*pb.Email, error) {
	email := &pb.Email{}
	return email, nil
}
