package services

import (
	"github.com/1token/email-services/database"
	"net/http"
)

type DraftServerImpl struct {
	DB *database.DatabaseX
}

func (s *DraftServerImpl) ListDrafts(w http.ResponseWriter, r *http.Request) {

}
