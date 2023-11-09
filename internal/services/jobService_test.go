// // * Job Service

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

func TestStore_JobByCompanyId(t *testing.T) {
	type args struct {
		jobs   []models.JobReq
		compId string
	}
	tests := []struct {
		name string
		// s       *Store
		args             args
		want             []models.JobRespo
		wantErr          bool
		mockRepoResponse func() ([]models.JobRespo, error)
	}{
		{
			name: "sucess case",
			args: args{
				jobs: []models.JobReq{
					{
	
						Name:       "Go Dev",
						Field:      "It",
						Experience: 2,
					},
				},
				compId: "1",
			},
			want: []models.JobRespo{
				{
					Id: 1,
				},
			},
			wantErr: false,
			mockRepoResponse: func() ([]models.JobRespo, error) {
				return []models.JobRespo{
					{
						Id: 1,
					},
				}, nil
			},
		},

		{
			name: "repository error",
			args: args{
				jobs: []models.JobReq{
					{
						Name:       "DevOps Engineer",
						Field:      "IT",
						Experience: 4,
					},
				},
				compId: "2",
			},
			want:    nil,
			wantErr: true,
			mockRepoResponse: func() ([]models.JobRespo, error) {
				return nil, errors.New("repository error")
			},
		},
		{
			name: "error case - empty company ID",
			args: args{
				jobs: []models.JobReq{
					{
						Name:       "Sales Manager",
						Field:      "Marketing",
						Experience: 5,
					},
				},
				compId: "abc",
			},
			want:    nil,
			wantErr: true,
			mockRepoResponse: func() ([]models.JobRespo, error) {
				return nil, errors.New("test error")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockInterface := repository.NewMockRepoInterface(mc)
			mockInterface.EXPECT().SaveJobsByCompanyId(tt.args.jobs, tt.args.compId).Return(tt.mockRepoResponse()).AnyTimes()
			s := NewStore(mockInterface)
			got, err := s.JobByCompanyId(tt.args.jobs, tt.args.compId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Store.JobByCompanyId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Store.JobByCompanyId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_FetchJobByCompanyId(t *testing.T) {
	type args struct {
		ctx       context.Context
		companyId string
	}
	tests := []struct {
		name string
		// s       *Store
		args             args
		want             []models.Job
		wantErr          bool
		mockRepoResponse func() ([]models.Job, error)
	}{
		{
			name: "Successful case",
			args: args{
				ctx:       context.Background(),
				companyId: "123",
			},
			want: []models.Job{
				{
					Name:       "Software Engineer",
					Field:      "IT",
					Experience: 2,
					CompanyId:  123,
				},
				{
					Name:       "Data Analyst",
					Field:      "IT",
					Experience: 1,
					CompanyId:  123,
				},
			},
			wantErr: false,
			mockRepoResponse: func() ([]models.Job, error) {
				return []models.Job{
					{
						Name:       "Software Engineer",
						Field:      "IT",
						Experience: 2,
						CompanyId:  123,
					},
					{
						Name:       "Data Analyst",
						Field:      "IT",
						Experience: 1,
						CompanyId:  123,
					},
				}, nil
			},
		},
		{
			name: "Error :repository error",
			args: args{
				ctx:       context.Background(),
				companyId: "456",
			},
			want:    nil,
			wantErr: true,
			mockRepoResponse: func() ([]models.Job, error) {
				return nil, errors.New("repository error")
			},
		},
		{
			name: "Error :empty company ID",
			args: args{
				ctx:       context.Background(),
				companyId: "",
			},
			want:    nil,
			wantErr: true,
			mockRepoResponse: func() ([]models.Job, error) {
				return nil, errors.New("Company Id invalid")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockInterface := repository.NewMockRepoInterface(mc)
			mockInterface.EXPECT().GetJobsByCId(tt.args.ctx, tt.args.companyId).Return(tt.mockRepoResponse()).AnyTimes()
			s := NewStore(mockInterface)
			got, err := s.FetchJobByCompanyId(tt.args.ctx, tt.args.companyId)

			if (err != nil) != tt.wantErr {
				t.Errorf("Store.FetchJobByCompanyId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Store.FetchJobByCompanyId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_GetJobById(t *testing.T) {
	type args struct {
		ctx   context.Context
		jobId string
	}
	tests := []struct {
		name string
		// s       *Store
		args             args
		want             models.Job
		wantErr          bool
		mockRepoResponse func() (models.Job, error)
	}{
		{
			name: "Successful case",
			args: args{
				ctx:   context.Background(),
				jobId: "123",
			},
			want: models.Job{
				Name:       "Software Engineer",
				Field:      "IT",
				Experience: 2,
				CompanyId:  456,
			},
			wantErr: false,
			mockRepoResponse: func() (models.Job, error) {
				return models.Job{
					Name:       "Software Engineer",
					Field:      "IT",
					Experience: 2,
					CompanyId:  456,
				}, nil
			},
		},
		{
			name: "Error:repository error",
			args: args{
				ctx:   context.Background(),
				jobId: "786",
			},
			want:    models.Job{},
			wantErr: true,
			mockRepoResponse: func() (models.Job, error) {
				return models.Job{}, errors.New("repository error")
			},
		},
		{
			name: "Error:empty job ID",
			args: args{
				ctx:   context.Background(),
				jobId: "",
			},
			want:    models.Job{},
			wantErr: true,
			mockRepoResponse: func() (models.Job, error) {
				return models.Job{}, errors.New("JobId Invalid")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockInterface := repository.NewMockRepoInterface(mc)
			mockInterface.EXPECT().FetchByJobId(tt.args.ctx, tt.args.jobId).Return(tt.mockRepoResponse()).AnyTimes()
			s := NewStore(mockInterface)
			got, err := s.GetJobById(tt.args.ctx, tt.args.jobId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Store.GetJobById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Store.GetJobById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_GetAllJobs(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		// s       *Store
		args    args
		want    []models.Job
		wantErr bool
		mockRepoResponse func() ([]models.Job,error)
	}{
		{
            name: "Successful case",
            args: args{
                ctx: context.Background(),
            },
            want: []models.Job{
                {
                    Name:       "Software Engineer",
                    Field:      "IT",
                    Experience: 2,
                    CompanyId:  456,
                },
                {
                    Name:       "Data Analyst",
                    Field:      "IT",
                    Experience: 1,
                    CompanyId:  456,
                },
            },
            wantErr: false,
            mockRepoResponse: func() ([]models.Job, error) {
                return []models.Job{
                    {
                        Name:       "Software Engineer",
                        Field:      "IT",
                        Experience: 2,
                        CompanyId:  456,
                    },
                    {
                        Name:       "Data Analyst",
                        Field:      "IT",
                        Experience: 1,
                        CompanyId:  456,
                    },
                }, nil
            },
        },
		{
            name: "Error:repository error",
            args: args{
                ctx: context.Background(),
            },
            want:    []models.Job{},
            wantErr: true,
            mockRepoResponse: func() ([]models.Job, error) {
                return []models.Job{}, errors.New("repository error")
            },
        },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockInterface := repository.NewMockRepoInterface(mc)
			mockInterface.EXPECT().FetchAllJobs(tt.args.ctx).Return(tt.mockRepoResponse()).AnyTimes()
			s := NewStore(mockInterface)
			got, err := s.GetAllJobs(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Store.GetAllJobs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Store.GetAllJobs() = %v, want %v", got, tt.want)
			}
		})
	}
}
