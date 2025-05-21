package sale

import "errors"

// ErrNotFound retorna cuando la venta proporcionada no existe (sale not found)
var ErrNotFound = errors.New("sale not found")

// ErrEmptyID retorna cuando se intenta almacenar una venta con ID vacio (id empty)
var ErrEmptyID = errors.New("empty ID")

// ErrInvalidStatus retorna cuando se intenta buscar una venta con un estado invalido
var ErrInvalidStatus = errors.New("invalid sale status")

var ErrInvalidAmount = errors.New("amount must be greater than 0")

var ErrInvalidRequest = errors.New("sale to update must have pending status")

var ErrInvalidUpdate = errors.New("invalid update status")

type Storage interface {
	Set(sale *Sale) error
	Read(id string) (*Sale, error)
	// Delete(id string) error
	FindSale(id string, st string) ([]*Sale, error)
}

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
	if sale.Amount <= 0 {
		return ErrInvalidAmount
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

func (l *LocalStorage) FindSale(id string, st string) ([]*Sale, error) {
	// busco errores
	if id == "" {
		return nil, ErrEmptyID
	}

	if st != "" && st != "approved" && st != "rejected" && st != "pending" {
		return nil, ErrInvalidStatus
	}

	//creo un arreglo para guardar las coincidencias
	var results []*Sale

	// recorro y busco
	for _, sale := range l.m {
		if st != "" {
			if sale.UserID == id && sale.Status == st {
				results = append(results, sale)
			}
		} else {
			if sale.UserID == id {
				results = append(results, sale)
			}
		}
	}

	if results == nil {
		return []*Sale{}, nil
	}

	return results, nil
}
