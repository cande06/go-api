package sale

import "errors"

// ErrNotFound retorna cuando la venta proporcionada no existe
var ErrNotFound = errors.New("sale not found")

// ErrEmptyID retorna cuando se intenta almacenar una venta con ID vacio
var ErrEmptyID = errors.New("empty sale ID")

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
