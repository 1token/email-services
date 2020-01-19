package service

import (
	// "context"

	// "github.com/gogo/protobuf/types"
	"database/sql"
	// biz "pes/common"
	pb "github.com/1token/email-services/email-apis/generated/go"
)

const (
	userTable = "users"
)

type DraftServerImpl struct {
	DB *sql.DB
}

func (s *DraftServerImpl) List(in *pb.Draft, stream pb.ListDraftsRequest) error {
	/*users := []*pb.User{}
	if err := biz.List(s.DB, userTable, &users, "order by data->'$.created' desc"); err != nil {
		return err
	}

	for _, v := range users {
		if err := stream.Send(v); err != nil {
			return err
		}
	}*/

	return nil
}
