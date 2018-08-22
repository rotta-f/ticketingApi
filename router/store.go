package router

type Store struct {
	s map[string]interface{}
}

var data map[string]*Store

func init() {
	data = map[string]*Store{}
}

func InitStore(id string) {
	data[id] = &Store{}
	data[id].s = map[string]interface{}{}
}

func GetStore(id string) *Store {
	if data[id] != nil {
		return data[id]
	}
	return nil
}

func (s *Store) Set(key string, value interface{}) {
	s.s[key] = value
}

func (s *Store) Get(key string) interface{} {
	return s.s[key]
}
