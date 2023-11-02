// * User service

package services

import (
	"context"
	"finalAssing/internal/models"
	"reflect"
	"testing"
)

func TestDbConnStruct_CreateUser(t *testing.T) {
	type args struct {
		ctx context.Context
		nu  models.NewUser
	}
	tests := []struct {
		name    string
		s       *DbConnStruct
		args    args
		want    models.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.CreateUser(tt.args.ctx, tt.args.nu)
			if (err != nil) != tt.wantErr {
				t.Errorf("DbConnStruct.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DbConnStruct.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
