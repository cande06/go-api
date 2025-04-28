package sale

import "errors"

// ErrNotFound retorna cuando la venta proporcionada no existe (sale not found)
var ErrNotFound = errors.New("not found")

// ErrEmptyID retorna cuando se intenta almacenar una venta con ID vacio (id empty)
var ErrEmptyID = errors.New("bad request")

// ErrInvalidStatus retorna cuando se intenta buscar una venta con un estado invalido
var ErrInvalidStatus = errors.New("invalid sale status")

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
			if sale.User_id == id && sale.Status == st {
				results = append(results, sale)
			}
		} else {
			if sale.User_id == id {
				results = append(results, sale)
			}
		}
	}

	if results == nil {
		return []*Sale{}, nil
	}

	return results, nil
}
