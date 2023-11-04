package services

import (
	"context"
	"errors"
	"finalAssing/internal/models"
	"finalAssing/internal/repository"
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestStore_CreateCompany(t *testing.T) {
	type args struct {
		ctx     context.Context
		newComp models.Company
	}
	tests := []struct {
		name string
		// s       *Store
		args             args
		want             models.Company
		wantErr          bool
		mockRepoResponse func() (models.Company, error)
	}{
		{
			name: "Successful case",
			args: args{
				ctx: context.Background(),
				newComp: models.Company{
					Name: "Google",
					City: "Bamglore",
					Jobs: []models.Job{
						{
							Name:       "Go dev",
							Field:      "IT",
							Experience: 1,
						},
						{
							Name:       "Sales Manager",
							Field:      "Marketing",
							Experience: 6,
						},
					},
				},
			},
			want: models.Company{
				Name: "Google",
				City: "Banglore",
				Jobs: []models.Job{
					{
						Name:       "Go dev",
						Field:      "IT",
						Experience: 1,
					},
					{
						Name:       "Sales Manager",
						Field:      "Marketing",
						Experience: 6,
					},
				},
			},
			wantErr: false,
			mockRepoResponse: func() (models.Company, error) {
				return models.Company{
					Name: "Google",
					City: "Banglore", Jobs: []models.Job{
						{
							Name:       "Go dev",
							Field:      "IT",
							Experience: 1,
						},
						{
							Name:       "Sales Manager",
							Field:      "Marketing",
							Experience: 6,
						},
					},
				}, nil
			},
		},
		{
			name: "Error: Repository error",
			args: args{
				ctx: context.Background(),
				newComp: models.Company{
					Name: "Invalid Company",
					City: "Invalid City",
				},
			},
			want:    models.Company{},
			wantErr: true,
			mockRepoResponse: func() (models.Company, error) {
				return models.Company{}, errors.New("repository error")
			},
		},
		{
			name: "Error: Invalid Name",
			args: args{
				ctx: context.Background(),
				newComp: models.Company{
					Name: "",
					City: "Banglore",
				},
			},
			want:    models.Company{},
			wantErr: true,
			mockRepoResponse: func() (models.Company, error) {
				return models.Company{}, errors.New("Inavlid Name")
			},
		},
		{
			name: "Error: Invalid City",
			args: args{
				ctx: context.Background(),
				newComp: models.Company{
					Name: "Google",
					City: "",
				},
			},
			want:    models.Company{},
			wantErr: true,
			mockRepoResponse: func() (models.Company, error) {
				return models.Company{}, errors.New("Inavlid City")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockInterface := repository.NewMockRepoInterface(mc)
			mockInterface.EXPECT().SaveCompany(tt.args.newComp).Return(tt.mockRepoResponse()).AnyTimes()
			s := NewStore(mockInterface)

			got, err := s.CreateCompany(tt.args.ctx, tt.args.newComp)
			if (err != nil) != tt.wantErr {
				t.Errorf("Store.CreateCompany() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Store.CreateCompany() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_ViewCompanies(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		// s       *Store
		args             args
		want             []models.Company
		wantErr          bool
		mockRepoResponse func() ([]models.Company, error)
	}{
		{
			name: "Error: Repository error",
			args: args{
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: true,
			mockRepoResponse: func() ([]models.Company, error) {
				return nil, errors.New("repository error")
			},
		},
		{
			name: "Successful case",
			args: args{
				ctx: context.Background(),
			},
			want: []models.Company{
				{
					Name: "Google",
					City: "Gurugram",
				},
				{
					Name: "TEKsysytem",
					City: "Banglore",
				},
			},
			wantErr: false,
			mockRepoResponse: func() ([]models.Company, error) {
				return []models.Company{
					{
						Name: "Google",
						City: "Gurugram",
					},
					{
						Name: "TEKsysytem",
						City: "Banglore",
					},
				}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockInterface := repository.NewMockRepoInterface(mc)
			mockInterface.EXPECT().FetchAllCompanies(tt.args.ctx).Return(tt.mockRepoResponse()).AnyTimes()
			s := NewStore(mockInterface)

			got, err := s.ViewCompanies(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Store.ViewCompanies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Store.ViewCompanies() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_FetchCompanyByID(t *testing.T) {
	type args struct {
		ctx       context.Context
		companyId string
	}
	tests := []struct {
		name    string
		// s       *Store
		args    args
		want    models.Company
		wantErr bool
		mockRepoResponse func() (models.Company, error)
	}{
		{
            name: "Successful case",
            args: args{
                ctx:       context.Background(),
                companyId: "1",
            },
            want: models.Company{
                Name: "Google",
                City: "Gurugram",
            },
            wantErr: false,
            mockRepoResponse: func() (models.Company, error) {
                return models.Company{
                    Name: "Google",
                    City: "Gurugram",
                }, nil
            },
        },
        {
            name: "Error: Company not found",
            args: args{
                ctx:       context.Background(),
                companyId: "999",
            },
            want:    models.Company{},
            wantErr: true,
            mockRepoResponse: func() (models.Company, error) {
                return models.Company{}, errors.New("Company Not found in db")
            },
        },
        {
            name: "Error: Repository error",
            args: args{
                ctx:       context.Background(),
                companyId: "1",
            },
            want:    models.Company{},
            wantErr: true,
            mockRepoResponse: func() (models.Company, error) {
                return models.Company{}, errors.New("repository error")
            },
        },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockInterface := repository.NewMockRepoInterface(mc)
			mockInterface.EXPECT().GetCompaniesById(tt.args.ctx,tt.args.companyId).Return(tt.mockRepoResponse()).AnyTimes()
			s := NewStore(mockInterface)
			got, err := s.FetchCompanyByID(tt.args.ctx, tt.args.companyId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Store.FetchCompanyByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Store.FetchCompanyByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
