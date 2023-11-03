// * User service

package services

import (
	"context"
	"finalAssing/internal/models"
	"finalAssing/internal/repository"
	"finalAssing/internal/repository/mockRepo"
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestStore_CreateUser(t *testing.T) {
	type args struct {
		ctx context.Context
		nu  models.NewUser
	}
	tests := []struct {
		name    string
		s       *Store
		args    args
		want    models.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mc:=gomock.NewController(t)
		     s:=mockRepo.NewMockRepoInterface(mc)
			got, err := tt.s.CreateUser(tt.args.ctx, tt.args.nu)
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
