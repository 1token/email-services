package services

import (
	"database/sql"
	"encoding/json"
	pb "github.com/1token/email-services/email-apis/generated/go"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type DraftServerImpl struct {
	DB *sql.DB
}

func (s *DraftServerImpl) ListDrafts(w http.ResponseWriter, r *http.Request) {
	drafts := &pb.ListDraftsResponse{}
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

	/*response, err := proto.Marshal(drafts)
	if err != nil {
		log.Fatalf("Unable to marshal response : %v", err)
	}
	w.Write(response)*/

	response, err := json.Marshal(drafts)
	if err != nil {
		log.Fatalf("Unable to marshal response : %v", err)
	}
	w.Write(response)
}
