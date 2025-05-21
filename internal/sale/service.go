package sale

import (
	"time"

	"math/rand"

	"errors"
	// "fmt"

	"go-api/internal/user"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

var ErrInexistentUser = errors.New("user does not exist")

type Service struct {
	storage     Storage
	userStorage user.Storage
	logger      *zap.Logger
}

func NewService(saleStorage Storage, userStorage user.Storage, logger *zap.Logger) *Service {
	if logger == nil {
		logger, _ = zap.NewProduction()
		defer logger.Sync()
	}

	return &Service{
		storage:     saleStorage,
		userStorage: userStorage,
		logger:      logger,
	}
}

// ## Funciones necesarias
// Genera un status random
func randomStatus() string {
	statuses := []string{"pending", "approved", "rejected"}
	return statuses[rand.Intn(len(statuses))]
}

// Verifica que el usuario exista sin hacer una consulta http al localhost
func (s *Service) validateUser(userID string) error {
	_, err := s.userStorage.Read(userID)
	if err != nil {
		return ErrInexistentUser
	}
	return nil
}

func (s *Service) Create(sale *Sale) error {
	now := time.Now()

	//Controles
	err := s.validateUser(sale.UserID)
	if err != nil {
		return err
	}

	sale.ID = uuid.NewString()
	sale.Status = randomStatus()
	sale.CreatedAt = now
	sale.UpdatedAt = now
	sale.Version = 1

	if err := s.storage.Set(sale); err != nil {
		return err
	}
	return s.storage.Set(sale)
}

func (s *Service) Get(id string, st string) ([]*Sale, error) {
	return s.storage.FindSale(id, st)
}

func (s *Service) Update(id string, sale *UpdateFields) (*Sale, error) {
	//sale exists
	existing, err := s.storage.Read(id)
	if err != nil {
		// if errors.Is(err, ErrNotFound) {
		// return nil, ErrNotFound
		// }

		return nil, err
	}
	//sale status must be pending
	if existing.Status != "pending" {
		return nil, ErrInvalidRequest
	}
	// validate body
	if (*sale.Status != "approved") && (*sale.Status != "rejected") || (sale.Status == nil) {
		return nil, ErrInvalidUpdate
	}

	existing.Status = *sale.Status
	existing.UpdatedAt = time.Now()
	existing.Version++

	if err := s.storage.Set(existing); err != nil {
		return nil, err
	}

	return existing, nil
}

// Verifica que el ID del usuario exista
// func validateUser(userID string) error {
// 	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/users/%s", userID))
// 	//resp almacena la respuesta del servidor
// 	//err almacenara un error en caso de que el servidor no pueda responder
// 	if err != nil {
// 		// return fmt.Errorf("failed to check user: %w", err)
// 		return fmt.Errorf("internal server error")
// 	}
// 	defer resp.Body.Close() //cierra resp.body al final de la funcion

// 	if resp.StatusCode == http.StatusNotFound {
// 		// return fmt.Errorf("user not found")
// 		return fmt.Errorf("bad request")
// 	}
// 	if resp.StatusCode != http.StatusOK {
// 		// return fmt.Errorf("unexpected error checking user: %s", resp.Status)
// 		return fmt.Errorf("bad request")
// 	}

// 	return nil
// }
