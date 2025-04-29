package sale

import (
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestService_Create_Simple(t *testing.T) {
	s := NewService(NewLocalStorage(), nil)

	input := &Sale{
		Status: "pending",
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
	}, nil)

	err = s.Create(input)
	require.NotNil(t, err)
	require.EqualError(t, err, "fake error trying to set user")
}

func TestService_Create(t *testing.T) {
	type fields struct {
		storage Storage
	}

	type args struct {
		sale *Sale
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		wantErr  func(t *testing.T, err error)
		wantSale func(t *testing.T, sale *Sale)
	}{
		{
			name: "error",
			fields: fields{
				storage: &mockStorage{
					mockSet: func(sale *Sale) error {
						return errors.New("fake error trying to set sale")
					},
				},
			},
			args: args{
				sale: &Sale{},
			},
			wantErr: func(t *testing.T, err error) {
				require.NotNil(t, err)
				require.EqualError(t, err, "fake error trying to set sale")
			},
			wantSale: nil,
		},
		{
			name: "success",
			fields: fields{
				storage: NewLocalStorage(),
			},
			args: args{
				sale: &Sale{
					Status: "pending",
				},
			},
			wantErr: func(t *testing.T, err error) {
				require.Nil(t, err)
			},
			wantSale: func(t *testing.T, input *Sale) {
				require.NotEmpty(t, input.ID)
				require.NotEmpty(t, input.CreatedAt)
				require.NotEmpty(t, input.UpdatedAt)
				require.Equal(t, 1, input.Version)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				storage: tt.fields.storage,
			}

			err := s.Create(tt.args.sale)
			if tt.wantErr != nil {
				tt.wantErr(t, err)
			}

			if tt.wantSale != nil {
				tt.wantSale(t, tt.args.sale)
			}
		})
	}
}

type mockStorage struct {
	mockSet    func(sale *Sale) error
	mockRead   func(id string) (*Sale, error)
	mockFindSale func(id string, st string) ([]*Sale, error)
	// mockDelete func(id string) error
}

func (m *mockStorage) Set(sale *Sale) error {
	return m.mockSet(sale)
}

func (m *mockStorage) Read(id string) (*Sale, error) {
	return m.mockRead(id)
}

func (m *mockStorage) FindSale(id string, st string) ([]*Sale, error) {
	return m.mockFindSale(id, st)
}

// func (m *mockStorage) Delete(id string) error {
// 	return m.mockDelete(id)
// }