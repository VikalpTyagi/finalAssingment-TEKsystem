// * User service

package services

import (
	"context"
	"errors"
	"finalAssing/internal/models"
	"finalAssing/internal/repository"
	"reflect"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/mock/gomock"
)

func TestStore_CreateUser(t *testing.T) {
	type args struct {
		ctx context.Context
		nu  models.NewUser
	}
	tests := []struct {
		name string
		// s       *Store
		args             args
		want             models.User
		wantErr          bool
		mockRepoResponse func() (models.User, error)
	}{
		{
			name: "Successful case",
			args: args{
				ctx: context.Background(),
				nu: models.NewUser{
					Name:  "Vikalp",
					Email: "Vikalp@gmai.com",
				},
			},
			want: models.User{
				Name:     "Vikalp",
				Email:    "Vikalp@gmail.com",
				PassHash: "hashed-password",
			},
			wantErr: false,
			mockRepoResponse: func() (models.User, error) {
				return models.User{
					Name:     "Vikalp",
					Email:    "Vikalp@gmail.com",
					PassHash: "hashed-password",
				}, nil
			},
		},
		{
			name: "Error:Password hashing failure",
			args: args{
				ctx: context.Background(),
				nu: models.NewUser{
					Name:     "Abhishek",
					Email:    "abhishek@gmail.com",
					Password: "123",
				},
			},
			want:    models.User{},
			wantErr: true,
			mockRepoResponse: func() (models.User, error) {
				return models.User{}, errors.New("Hashing failed")
			},
		},
		{
			name: "Error: Repository error",
			args: args{
				ctx: context.Background(),
				nu: models.NewUser{
					Name:     "Priya",
					Email:    "Priya@gmail.com",
					Password: "Priya@1356",
				},
			},
			want:    models.User{},
			wantErr: true,
			mockRepoResponse: func() (models.User, error) {
				return models.User{}, errors.New("repository error")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockInterface := repository.NewMockRepoInterface(mc)
			mockInterface.EXPECT().SaveUser(tt.args.ctx, tt.args.nu).Return(tt.mockRepoResponse()).AnyTimes()
			s := NewStore(mockInterface,nil)
			got, err := s.CreateUser(tt.args.ctx, tt.args.nu)
			if (err != nil) != tt.wantErr {
				t.Errorf("Store.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Store.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_Authenticate(t *testing.T) {
	type args struct {
		ctx      context.Context
		email    string
		password string
	}
	tests := []struct {
		name string
		// s       *Store
		args             args
		want             jwt.RegisteredClaims
		wantErr          bool
		mockRepoResponse func() (models.User, error)
	}{
		{
			name: "Error :Repository error",
			args: args{
				ctx:      context.Background(),
				email:    "piyush@gmail.com",
				password: "password",
			},
			want:    jwt.RegisteredClaims{},
			wantErr: true,
			mockRepoResponse: func() (models.User, error) {
				return models.User{}, errors.New("repository error")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockInterface := repository.NewMockRepoInterface(mc)
			mockInterface.EXPECT().CheckEmail(tt.args.email, tt.args.password).Return(tt.mockRepoResponse()).AnyTimes()
			s := NewStore(mockInterface,nil)
			got, err := s.Authenticate(tt.args.ctx, tt.args.email, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Store.Authenticate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Store.Authenticate() = %v, want %v", got, tt.want)
			}
		})
	}
}
