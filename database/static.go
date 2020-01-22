package database

type staticDraftsStorage struct {
	DatabaseX

	// A read-only set of drafts.
	drafts     []Draft
	draftsByID map[string]Draft
}

func (s staticDraftsStorage) isStatic(id string) bool {
	_, ok := s.draftsByID[id]
	return ok
}

func (s staticDraftsStorage) ListDrafts() ([]Draft, error) {
	drafts, err := s.DatabaseX.ListDrafts()
	if err != nil {
		return nil, err
	}
	n := 0
	for _, draft := range drafts {
		// If a draft in the backing storage has the same ID as a static draft
		// prefer the static draft.
		if !s.isStatic(draft.ID) {
			drafts[n] = draft
			n++
		}
	}
	return append(drafts[:n], s.drafts...), nil
}
