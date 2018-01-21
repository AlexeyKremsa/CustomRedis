package storage

import "errors"

const (
	errWrongType = "Operation against a key holding the wrong kind of value"
)

type Str struct {
	Value string
}

type Storage struct {
	keyValues map[string]interface{}
}

func Init() *Storage {
	strg := &Storage{}
	strg.keyValues = make(map[string]interface{})
	return strg
}

func (s *Storage) SetStr(key, value string) {
	item := Str{Value: value}
	s.keyValues[key] = item
}

func (s *Storage) GetStr(key string) (string, error) {
	if v, ok := s.keyValues[key]; ok {
		if s, ok := v.(Str); ok {
			return s.Value, nil
		}
		return "", errors.New(errWrongType)
	}

	return "", nil
}
