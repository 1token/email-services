package database

import (
	"context"
	"database/sql"
	"github.com/golang/protobuf/ptypes/empty"

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
	/*if err := database.List(s.DB, "draft", &drafts, "order by data->'$.created' desc"); err != nil {
		return nil, err
	}*/
	draft := &pb.Draft{
		Id:       "abcd",
		Snipped:  "Hello",
		Envelope: nil,
	}
	drafts.Draft = append(drafts.Draft, draft)
	draft = &pb.Draft{
		Id:       "efgh",
		Snipped:  "World",
		Envelope: nil,
	}
	drafts.Draft = append(drafts.Draft, draft)
	return drafts, nil
}

func (s *DraftServerImpl) CreateDraft(ctx context.Context, in *pb.CreateDraftRequest) (*pb.Draft, error) {
	draft := &pb.Draft{}

	/*insertCommentQuery := `INSERT INTO comments(ticket_id, owner, content, metadata, created_at, updated_at) VALUES ($1, $2, $3, $4, now(), now())`
	updateTicketQuery := `UPDATE tickets SET updated_at=now() WHERE id=$1`

	batch := &pgx.Batch{}
	batch.Queue("BEGIN")
	batch.Queue(insertCommentQuery, comment.TicketId, comment.Owner, comment.Content, comment.Metadata)
	batch.Queue(updateTicketQuery, comment.TicketId)
	batch.Queue("COMMIT")

	results := service.db.SendBatch(context, batch)
	if err := results.Close(); err != nil {
		if strings.Contains(err.Error(), "comments_ticket_id_fkey") {
			return status.Error(codes.InvalidArgument, "create_comment.ticket_not_exists")
		}

		service.logger.Error("error on inserting new comment: %v", err)
		return status.Error(codes.Internal, "create_comment.failed")
	}*/

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
