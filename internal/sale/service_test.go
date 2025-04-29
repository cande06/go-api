package sale

import (
	"errors"
	"go-api/internal/user"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestService_Create_Simple(t *testing.T) {
	userStorage := &mockUserStorage{
		mockRead: func(id string) (*user.User, error) {
			return &user.User{
				ID:   id,
				Name: "Test",
			}, nil
		},
	}

	s := NewService(NewLocalStorage(), userStorage, zap.NewNop())

	input := &Sale{
		Status: "pending",
		UserID: "graragjirgjmgpgmoai",
		Amount: 100.30,
	}

	err := s.Create(input)
	require.Nil(t, err)
	require.NotEmpty(t, input.ID)
	require.NotEmpty(t, input.CreatedAt)
	require.NotEmpty(t, input.UpdatedAt)
	require.Equal(t, 1, input.Version)

	s = NewService(&mockStorage{
		mockSet: func(sale *Sale) error {
			return errors.New("fake error trying to set sale")
		},
	}, userStorage, zap.NewNop())

	err = s.Create(input)
	require.NotNil(t, err)
	require.EqualError(t, err, "fake error trying to set sale")
}

func TestService_Create(t *testing.T) {
	type fields struct {
		storage     Storage
		userStorage user.Storage
	}

	type args struct {
		sale *Sale
	}

	tests := []struct {
		name       string
		fields     fields
		args       args
		wantErr    func(t *testing.T, err error)
		wantResult func(t *testing.T, sale *Sale)
	}{
		{
			name: "error",
			fields: fields{
				storage: &mockStorage{
					mockSet: func(sale *Sale) error {
						return errors.New("fake error trying to set sale")
					},
				},
				userStorage: &mockUserStorage{
					mockRead: func(id string) (*user.User, error) {
						return &user.User{ID: id, Name: "Test"}, nil
					},
				},
			},
			args: args{
				sale: &Sale{
					UserID: "123",
					Amount: 100, // evitar validaciones
				},
			},
			wantErr: func(t *testing.T, err error) {
				require.NotNil(t, err)
				require.EqualError(t, err, "fake error trying to set sale")
			},
			wantResult: nil,
		},
		{
			name: "success",
			fields: fields{
				storage: NewLocalStorage(),
				userStorage: &mockUserStorage{
					mockRead: func(id string) (*user.User, error) {
						return &user.User{ID: id, Name: "Test"}, nil
					},
				},
			},
			args: args{
				sale: &Sale{
					Status: "pending",
					UserID: "123",
					Amount: 50, // válido
				},
			},
			wantErr: func(t *testing.T, err error) {
				require.Nil(t, err)
			},
			wantResult: func(t *testing.T, sale *Sale) {
				require.NotEmpty(t, sale.ID)
				require.NotEmpty(t, sale.CreatedAt)
				require.NotEmpty(t, sale.UpdatedAt)
				require.Equal(t, 1, sale.Version)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				storage:     tt.fields.storage,
				userStorage: tt.fields.userStorage,
			}

			err := s.Create(tt.args.sale)

			if tt.wantErr != nil {
				tt.wantErr(t, err)
			}

			if tt.wantResult != nil {
				tt.wantResult(t, tt.args.sale)
			}
		})
	}
}


type mockStorage struct {
	mockSet      func(sale *Sale) error
	mockRead     func(id string) (*Sale, error)
	mockFindSale func(id string, st string) ([]*Sale, error)
	// mockDelete func(id string) error
}

type mockUserStorage struct {
	mockSet    func(user *user.User) error
	mockRead   func(id string) (*user.User, error)
	mockDelete func(id string) error
}

func (m *mockStorage) FindSale(id string, st string) ([]*Sale, error) {
	if m.mockFindSale != nil {
		return m.mockFindSale(id, st)
	}
	return nil, errors.New("find sale not implemented")
}

func (m *mockStorage) Set(sale *Sale) error {
	return m.mockSet(sale)
}

func (m *mockUserStorage) Set(u *user.User) error {
	if m.mockSet != nil {
		return m.mockSet(u)
	}
	return nil
}

func (m *mockStorage) Read(id string) (*Sale, error) {
	if m.mockRead != nil {
		return m.mockRead(id)
	}
	return nil, errors.New("sale not found")
}

func (m *mockUserStorage) Read(id string) (*user.User, error) {
	if m.mockRead != nil {
		return m.mockRead(id)
	}
	return nil, errors.New("user not found")
}

func (m *mockUserStorage) Delete(id string) error {
	if m.mockDelete != nil {
		return m.mockDelete(id)
	}
	return nil
}
