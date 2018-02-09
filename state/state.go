package state

type State interface {
	Version() uint64
	Hash() []byte
	Commit(version uint64) ([]byte, error)

	Account() AccountAdapter
}

func NewState(db DB) State {
	return &state{db}
}

type state struct {
	db DB
}

func (s *state) Version() uint64 {
	return s.db.Version()
}

func (s *state) Hash() []byte {
	return s.db.Hash()
}

func (s *state) Commit(version uint64) ([]byte, error) {
	return s.db.Commit(version)
}

func (s *state) Account() AccountAdapter {
	return NewAccountAdapter(s.db)
}