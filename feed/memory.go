package feed

type MemoryStore struct {
	db        []interface{}
	positions map[string]uint64
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{positions: make(map[string]uint64)}
}

func (s *MemoryStore) GetPosition(t string) uint64 {
	if v, ok := s.positions[t]; ok {
		return v
	}

	return 0
}

func (s *MemoryStore) SetPosition(t string, after uint64) error {
	s.positions[t] = after
	return nil
}

func (s *MemoryStore) Save(i interface{}) error {
	s.db = append(s.db, i)
	return nil
}

func (s *MemoryStore) Get() interface{} {
	return s.db
}
