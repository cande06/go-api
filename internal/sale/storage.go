package sale

import "errors"

// ErrNotFound retorna cuando la venta proporcionada no existe (sale not found)
var ErrNotFound = errors.New("sale not found")

// ErrEmptyID retorna cuando se ingresa un sale ID vacio (id empty)
var ErrEmptyID = errors.New("empty ID")

// ErrInvalidAmount retorna cuando se intenta ingresar un monto menor o igual a 0
var ErrInvalidAmount = errors.New("amount must be greater than 0")

// ErrInvalidStatus retorna cuando se ingresa un estado mal escrito o vacío
var ErrInvalidStatus = errors.New("invalid sale status")

// ErrSameStatus retorna cuando el estado para actualizar es pending
var ErrSameStatus = errors.New("status is already 'pending'")

// ErrInvalidUpdate retorna cuando se intenta realizar un cambio de estado no válido
var ErrInvalidUpdate = errors.New("only sales with a 'pending' status can be updated")

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
	s, ok := l.m[id]
	if !ok {
		return nil, ErrNotFound
	}

	return s, nil
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
