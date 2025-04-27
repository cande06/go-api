package sale

import (
	"time"

	"math/rand"

	"fmt"
	"net/http"

	"github.com/google/uuid"
)

// Funciones necesarias
// Genera un status random
func randomStatus() string {
	statuses := []string{"pending", "approved", "rejected"}
	return statuses[rand.Intn(len(statuses))]
}

// Verifica que el ID del usuario exista
func validateUser(userID string) error {
	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/users/%s", userID))
	//resp almacena la respuesta del servidor
	//err almacenara un error en caso de que el servidor no pueda responder
	if err != nil {
		return fmt.Errorf("failed to check user: %w", err)
	}
	defer resp.Body.Close() //cierra resp.body al final de la funcion

	if resp.StatusCode == http.StatusNotFound {
		return fmt.Errorf("user not found")
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected error checking user: %s", resp.Status)
	}

	return nil
}

type Service struct {
	storage *LocalStorage
}

func NewService(storage *LocalStorage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) Create(sale *Sale) error {
	now := time.Now()

	//Controles
	err := validateUser(sale.User_id)
	if err != nil {
		return err
	}
	if sale.Amount <= 0 {
		return fmt.Errorf("amount must be greater than 0")
	}

	sale.ID = uuid.NewString()
	sale.Status = randomStatus()
	sale.CreatedAt = now
	sale.UpdatedAt = now
	sale.Version = 1

	return s.storage.Set(sale)
}

func (s *Service) Get(id string) (*Sale, error) {
	return s.storage.Read(id)
}

func (s *Service) Update(id string, sale *UpdateFields) (*Sale, error) {
	existing, err := s.storage.Read(id)
	if err != nil {
		return nil, err
	}

	if existing.Status != "pending" {
		return nil, fmt.Errorf("only pending sales can be updated")
	}

	if sale.Status == nil {
		return nil, fmt.Errorf("status is required")
	}

	if *sale.Status != "approved" && *sale.Status != "rejected" {
		return nil, fmt.Errorf("invalid status")
	}

	existing.Status = *sale.Status
	existing.UpdatedAt = time.Now()
	existing.Version++

	if err := s.storage.Set(existing); err != nil {
		return nil, err
	}

	return existing, nil
}
