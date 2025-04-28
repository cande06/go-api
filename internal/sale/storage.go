package sale

import "errors"

// ErrNotFound retorna cuando la venta proporcionada no existe (sale not found)
var ErrNotFound = errors.New("not found")

// ErrEmptyID retorna cuando se intenta almacenar una venta con ID vacio (id empty)
var ErrEmptyID = errors.New("bad request")

type LocalStorage struct {
	m map[string]*Sale
}

func NewLocalStorage() *LocalStorage {
	return &LocalStorage{
		m: map[string]*Sale{},
	}
}

func (l *LocalStorage) Set(sale *Sale) error {
	if sale.ID == "" {
		return ErrEmptyID
	}

	l.m[sale.ID] = sale
	return nil
}

func (l *LocalStorage) Read(id string) (*Sale, error) {
	u, ok := l.m[id]
	if !ok {
		return nil, ErrNotFound
	}

	return u, nil
}
