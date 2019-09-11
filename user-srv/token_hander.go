package main

type Authable interface {
	Decode(token string) (interface{}, error)
	Encode(data interface{}) (string, error)
}

type TokenHandler struct {
	repo Repository
}

func (th *TokenHandler) Decode(token string) (interface{}, error) {
	return "", nil
}

func (th *TokenHandler) Encode(data interface{}) (string, error) {
	return "", nil
}
