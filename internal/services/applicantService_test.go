package services

import (
	"context"
	"errors"
	"finalAssing/internal/cacheier"
	"finalAssing/internal/models"
	"finalAssing/internal/repository"
	"testing"

	"go.uber.org/mock/gomock"
	"gopkg.in/go-playground/assert.v1"
	"gorm.io/gorm"
)

func TestStore_FIlterApplication(t *testing.T) {
	type args struct {
		ctx           context.Context
		applicantList []*models.ApplicantReq
	}
	tests := []struct {
		name             string
		args             args
		want             []*models.ApplicantRespo
		wantErr          bool
		mockRedisRes func(redisMock *cacheier.MockRedInterface)
		mockRepoResponse func(repoMock *repository.MockRepoInterface)
		
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				applicantList: []*models.ApplicantReq{
					{
						Name:           "Vikalp Tyagi",
						JobId:          1,
						Experience:     3,
						Max_NP:         2,
						Budget:         50000,
						Locations:      []uint{1, 2, 3},
						Stack:          []uint{1, 2, 3},
						WorkMode:       "Full-Time",
						Qualifications: []uint{1},
						Shift:          "Day",
					},
					{
						Name:           "Satyam",
						JobId:          2,
						Experience:     3,
						Max_NP:         2,
						Budget:         50000,
						Locations:      []uint{1, 2, 3},
						Stack:          []uint{1, 2, 3},
						WorkMode:       "Full-Time",
						Qualifications: []uint{1},
						Shift:          "Day",
					},
					{ //* Problem: Budget
						Name:           "Ajay",
						JobId:          3,
						Experience:     3,
						Max_NP:         2,
						Budget:         90000,
						Locations:      []uint{1, 2},
						Stack:          []uint{1, 2, 3},
						WorkMode:       "Full-Time",
						Qualifications: []uint{1},
						Shift:          "Day",
					},
				},
			},
			want: []*models.ApplicantRespo{
				{
					Name:  "Vikalp Tyagi",
					JobId: 1,
				},
				{
					Name:  "Satyam",
					JobId: 2,
				},
			},
			wantErr: false,
			mockRedisRes: func(redisMock *cacheier.MockRedInterface) {
				redisMock.EXPECT().FetchJobData(gomock.Any(),uint(1)).Return(nil,errors.New("test error: faliure in fetching from redis")).Times(1)
				redisMock.EXPECT().AddJobData(gomock.Any(),uint(1),gomock.Any()).Return(nil).Times(1)

				redisMock.EXPECT().FetchJobData(gomock.Any(),uint(2)).Return(&models.Job{
					Experience: 4,
					MinExp:     1,
					Min_NP:     1,
					Max_NP:     4,
					Budget:     80000,
					Stack: []models.Skill{
						{
							Model: gorm.Model{ID: 1},
						},
						{
							Model: gorm.Model{ID: 2},
						},
						{
							Model: gorm.Model{ID: 3},
						},
					},
					Locations: []models.Location{
						{
							Model: gorm.Model{ID: 1},
						},
						{
							Model: gorm.Model{ID: 2},
						},
						{
							Model: gorm.Model{ID: 3},
						},
					},
					Qualifications: []models.Qualification{
						{
							Model: gorm.Model{ID: 1},
						},
					},
					WorkMode: "Full-Time",
					Shift:    "Day",
				},nil).Times(1)
				redisMock.EXPECT().FetchJobData(gomock.Any(),uint(3)).Return(nil,errors.New("test error: faliure case")).Times(1)

			},
			mockRepoResponse: func(repoMock *repository.MockRepoInterface) {

				repoMock.EXPECT().GetJobRequirment(uint(1)).Return(&models.Job{
					Experience: 4,
					MinExp:     1,
					Min_NP:     1,
					Max_NP:     4,
					Budget:     80000,
					Stack: []models.Skill{
						{
							Model: gorm.Model{ID: 1},
						},
						{
							Model: gorm.Model{ID: 2},
						},
						{
							Model: gorm.Model{ID: 3},
						},
					},
					Locations: []models.Location{
						{
							Model: gorm.Model{ID: 1},
						},
						{
							Model: gorm.Model{ID: 2},
						},
						{
							Model: gorm.Model{ID: 3},
						},
					},
					Qualifications: []models.Qualification{
						{
							Model: gorm.Model{ID: 1},
						},
					},
					WorkMode: "Full-Time",
					Shift:    "Day",
				}, nil).Times(1)
				// repoMock.EXPECT().GetJobRequirment(uint(2)).Return(&models.Job{
					// Experience: 4,
					// MinExp:     1,
					// Min_NP:     1,
					// Max_NP:     4,
					// Budget:     80000,
					// Stack: []models.Skill{
					// 	{
					// 		Model: gorm.Model{ID: 1},
					// 	},
					// 	{
					// 		Model: gorm.Model{ID: 2},
					// 	},
					// 	{
					// 		Model: gorm.Model{ID: 3},
					// 	},
					// },
					// Locations: []models.Location{
					// 	{
					// 		Model: gorm.Model{ID: 1},
					// 	},
					// 	{
					// 		Model: gorm.Model{ID: 2},
					// 	},
					// 	{
					// 		Model: gorm.Model{ID: 3},
					// 	},
					// },
					// Qualifications: []models.Qualification{
					// 	{
					// 		Model: gorm.Model{ID: 1},
					// 	},
					// },
					// WorkMode: "Full-Time",
					// Shift:    "Day",
				// }, nil).Times(1)
				repoMock.EXPECT().GetJobRequirment(uint(3)).Return(nil, errors.New("test error")).Times(1)
			},
		},

		// {
		// 	name: "Faliure cases",
		// 	args: args{
		// 		ctx: context.Background(),
		// 		applicantList: []*models.ApplicantReq{
		// 			{ //* problem Budget
		// 				Name:           "Vijay",
		// 				JobId:          1,
		// 				Experience:     3,
		// 				Max_NP:         1,
		// 				Budget:         90000,
		// 				Locations:      []uint{1, 2},
		// 				Stack:          []uint{1, 2, 3},
		// 				WorkMode:       "Full-Time",
		// 				Qualifications: []uint{1},
		// 				Shift:          "Day",
		// 			},
		// 			{ //* problem Exp
		// 				Name:           "Ajay",
		// 				JobId:          2,
		// 				Experience:     0,
		// 				Max_NP:         1,
		// 				Budget:         50000,
		// 				Locations:      []uint{1, 2},
		// 				Stack:          []uint{1, 2, 3},
		// 				WorkMode:       "Full-Time",
		// 				Qualifications: []uint{1},
		// 				Shift:          "Day",
		// 			},
		// 			{ //* problem NP
		// 				Name:           "Ram",
		// 				JobId:          3,
		// 				Experience:     3,
		// 				Max_NP:         8,
		// 				Budget:         50000,
		// 				Locations:      []uint{1, 2},
		// 				Stack:          []uint{1, 2, 3},
		// 				WorkMode:       "Full-Time",
		// 				Qualifications: []uint{1},
		// 				Shift:          "Day",
		// 			},
		// 			{ //* problem : Shift
		// 				Name:           "Mohommad",
		// 				JobId:          4,
		// 				Experience:     3,
		// 				Max_NP:         2,
		// 				Budget:         50000,
		// 				Locations:      []uint{1, 2},
		// 				Stack:          []uint{1, 2, 3},
		// 				WorkMode:       "Full-Time",
		// 				Qualifications: []uint{1},
		// 				Shift:          "Night",
		// 			},
		// 			{ //* Problem: Workmode
		// 				Name:           "Ashish",
		// 				JobId:          5,
		// 				Experience:     3,
		// 				Max_NP:         2,
		// 				Budget:         50000,
		// 				Locations:      []uint{1, 2},
		// 				Stack:          []uint{1, 2, 3},
		// 				WorkMode:       "Part-Time",
		// 				Qualifications: []uint{1},
		// 				Shift:          "Day",
		// 			},
		// 			{ //* Problem: Qualification
		// 				Name:           "Akbar",
		// 				JobId:          6,
		// 				Experience:     3,
		// 				Max_NP:         2,
		// 				Budget:         50000,
		// 				Locations:      []uint{1, 2},
		// 				Stack:          []uint{1, 2, 3},
		// 				WorkMode:       "Full-Time",
		// 				Qualifications: []uint{8, 9},
		// 				Shift:          "Day",
		// 			},
		// 			{ //* Problem: Location
		// 				Name:           "Anthony",
		// 				JobId:          7,
		// 				Experience:     3,
		// 				Max_NP:         2,
		// 				Budget:         50000,
		// 				Locations:      []uint{5, 6},
		// 				Stack:          []uint{1, 2, 3},
		// 				WorkMode:       "Full-Time",
		// 				Qualifications: []uint{1},
		// 				Shift:          "Day",
		// 			},
		// 			{ //* Problem: skills
		// 				Name:           "Pooja",
		// 				JobId:          8,
		// 				Experience:     3,
		// 				Max_NP:         2,
		// 				Budget:         50000,
		// 				Locations:      []uint{1, 2},
		// 				Stack:          []uint{1},
		// 				WorkMode:       "Full-Time",
		// 				Qualifications: []uint{1},
		// 				Shift:          "Day",
		// 			},
		// 		},
		// 	},
		// 	want:    nil,
		// 	wantErr: false,
		// 	mockRepoResponse: func(repoMock *repository.MockRepoInterface) {
		// 		repoMock.EXPECT().GetJobRequirment(uint(1)).Return(&models.Job{
		// 			Experience: 4,
		// 			MinExp:     1,
		// 			Min_NP:     1,
		// 			Max_NP:     4,
		// 			Budget:     80000,
		// 			Stack: []models.Skill{
		// 				{
		// 					Model: gorm.Model{ID: 1},
		// 				},
		// 				{
		// 					Model: gorm.Model{ID: 2},
		// 				},
		// 				{
		// 					Model: gorm.Model{ID: 3},
		// 				},
		// 			},
		// 			Locations: []models.Location{
		// 				{
		// 					Model: gorm.Model{ID: 1},
		// 				},
		// 				{
		// 					Model: gorm.Model{ID: 2},
		// 				},
		// 				{
		// 					Model: gorm.Model{ID: 3},
		// 				},
		// 			},
		// 			Qualifications: []models.Qualification{
		// 				{
		// 					Model: gorm.Model{ID: 1},
		// 				},
		// 			},
		// 			WorkMode: "Full-Time",
		// 			Shift:    "Day",
		// 		}, nil).AnyTimes()
		// 		repoMock.EXPECT().GetJobRequirment(uint(2)).Return(&models.Job{
		// 			Experience: 4,
		// 			MinExp:     1,
		// 			Min_NP:     1,
		// 			Max_NP:     4,
		// 			Budget:     80000,
		// 			Stack: []models.Skill{
		// 				{
		// 					Model: gorm.Model{ID: 1},
		// 				},
		// 				{
		// 					Model: gorm.Model{ID: 2},
		// 				},
		// 				{
		// 					Model: gorm.Model{ID: 3},
		// 				},
		// 			},
		// 			Locations: []models.Location{
		// 				{
		// 					Model: gorm.Model{ID: 1},
		// 				},
		// 				{
		// 					Model: gorm.Model{ID: 2},
		// 				},
		// 				{
		// 					Model: gorm.Model{ID: 3},
		// 				},
		// 			},
		// 			Qualifications: []models.Qualification{
		// 				{
		// 					Model: gorm.Model{ID: 1},
		// 				},
		// 			},
		// 			WorkMode: "Full-Time",
		// 			Shift:    "Day",
		// 		}, nil).AnyTimes()
		// 		repoMock.EXPECT().GetJobRequirment(uint(3)).Return(&models.Job{
		// 			Experience: 4,
		// 			MinExp:     1,
		// 			Min_NP:     1,
		// 			Max_NP:     4,
		// 			Budget:     80000,
		// 			Stack: []models.Skill{
		// 				{
		// 					Model: gorm.Model{ID: 1},
		// 				},
		// 				{
		// 					Model: gorm.Model{ID: 2},
		// 				},
		// 				{
		// 					Model: gorm.Model{ID: 3},
		// 				},
		// 			},
		// 			Locations: []models.Location{
		// 				{
		// 					Model: gorm.Model{ID: 1},
		// 				},
		// 				{
		// 					Model: gorm.Model{ID: 2},
		// 				},
		// 				{
		// 					Model: gorm.Model{ID: 3},
		// 				},
		// 			},
		// 			Qualifications: []models.Qualification{
		// 				{
		// 					Model: gorm.Model{ID: 1},
		// 				},
		// 			},
		// 			WorkMode: "Full-Time",
		// 			Shift:    "Day",
		// 		}, nil).AnyTimes()
		// 		repoMock.EXPECT().GetJobRequirment(uint(4)).Return(&models.Job{
		// 			Experience: 4,
		// 			MinExp:     1,
		// 			Min_NP:     1,
		// 			Max_NP:     4,
		// 			Budget:     80000,
		// 			Stack: []models.Skill{
		// 				{
		// 					Model: gorm.Model{ID: 1},
		// 				},
		// 				{
		// 					Model: gorm.Model{ID: 2},
		// 				},
		// 				{
		// 					Model: gorm.Model{ID: 3},
		// 				},
		// 			},
		// 			Locations: []models.Location{
		// 				{
		// 					Model: gorm.Model{ID: 1},
		// 				},
		// 				{
		// 					Model: gorm.Model{ID: 2},
		// 				},
		// 				{
		// 					Model: gorm.Model{ID: 3},
		// 				},
		// 			},
		// 			Qualifications: []models.Qualification{
		// 				{
		// 					Model: gorm.Model{ID: 1},
		// 				},
		// 			},
		// 			WorkMode: "Full-Time",
		// 			Shift:    "Day",
		// 		}, nil).AnyTimes()
		// 		repoMock.EXPECT().GetJobRequirment(uint(5)).Return(&models.Job{
		// 			Experience: 4,
		// 			MinExp:     1,
		// 			Min_NP:     1,
		// 			Max_NP:     4,
		// 			Budget:     80000,
		// 			Stack: []models.Skill{
		// 				{
		// 					Model: gorm.Model{ID: 1},
		// 				},
		// 				{
		// 					Model: gorm.Model{ID: 2},
		// 				},
		// 				{
		// 					Model: gorm.Model{ID: 3},
		// 				},
		// 			},
		// 			Locations: []models.Location{
		// 				{
		// 					Model: gorm.Model{ID: 1},
		// 				},
		// 				{
		// 					Model: gorm.Model{ID: 2},
		// 				},
		// 				{
		// 					Model: gorm.Model{ID: 3},
		// 				},
		// 			},
		// 			Qualifications: []models.Qualification{
		// 				{
		// 					Model: gorm.Model{ID: 1},
		// 				},
		// 			},
		// 			WorkMode: "Full-Time",
		// 			Shift:    "Day",
		// 		}, nil).AnyTimes()
		// 		repoMock.EXPECT().GetJobRequirment(uint(6)).Return(&models.Job{
		// 			Experience: 4,
		// 			MinExp:     1,
		// 			Min_NP:     1,
		// 			Max_NP:     4,
		// 			Budget:     80000,
		// 			Stack: []models.Skill{
		// 				{
		// 					Model: gorm.Model{ID: 1},
		// 				},
		// 				{
		// 					Model: gorm.Model{ID: 2},
		// 				},
		// 				{
		// 					Model: gorm.Model{ID: 3},
		// 				},
		// 			},
		// 			Locations: []models.Location{
		// 				{
		// 					Model: gorm.Model{ID: 1},
		// 				},
		// 				{
		// 					Model: gorm.Model{ID: 2},
		// 				},
		// 				{
		// 					Model: gorm.Model{ID: 3},
		// 				},
		// 			},
		// 			Qualifications: []models.Qualification{
		// 				{
		// 					Model: gorm.Model{ID: 1},
		// 				},
		// 			},
		// 			WorkMode: "Full-Time",
		// 			Shift:    "Day",
		// 		}, nil).AnyTimes()
		// 		repoMock.EXPECT().GetJobRequirment(uint(7)).Return(&models.Job{
		// 			Experience: 4,
		// 			MinExp:     1,
		// 			Min_NP:     1,
		// 			Max_NP:     4,
		// 			Budget:     80000,
		// 			Stack: []models.Skill{
		// 				{
		// 					Model: gorm.Model{ID: 1},
		// 				},
		// 				{
		// 					Model: gorm.Model{ID: 2},
		// 				},
		// 				{
		// 					Model: gorm.Model{ID: 3},
		// 				},
		// 			},
		// 			Locations: []models.Location{
		// 				{
		// 					Model: gorm.Model{ID: 1},
		// 				},
		// 				{
		// 					Model: gorm.Model{ID: 2},
		// 				},
		// 				{
		// 					Model: gorm.Model{ID: 3},
		// 				},
		// 			},
		// 			Qualifications: []models.Qualification{
		// 				{
		// 					Model: gorm.Model{ID: 1},
		// 				},
		// 			},
		// 			WorkMode: "Full-Time",
		// 			Shift:    "Day",
		// 		}, nil).AnyTimes()
		// 		repoMock.EXPECT().GetJobRequirment(uint(8)).Return(&models.Job{
		// 			Experience: 4,
		// 			MinExp:     1,
		// 			Min_NP:     1,
		// 			Max_NP:     4,
		// 			Budget:     80000,
		// 			Stack: []models.Skill{
		// 				{
		// 					Model: gorm.Model{ID: 1},
		// 				},
		// 				{
		// 					Model: gorm.Model{ID: 2},
		// 				},
		// 				{
		// 					Model: gorm.Model{ID: 3},
		// 				},
		// 				{
		// 					Model: gorm.Model{ID: 4},
		// 				},
		// 			},
		// 			Locations: []models.Location{
		// 				{
		// 					Model: gorm.Model{ID: 1},
		// 				},
		// 				{
		// 					Model: gorm.Model{ID: 2},
		// 				},
		// 				{
		// 					Model: gorm.Model{ID: 3},
		// 				},
		// 			},
		// 			Qualifications: []models.Qualification{
		// 				{
		// 					Model: gorm.Model{ID: 1},
		// 				},
		// 			},
		// 			WorkMode: "Full-Time",
		// 			Shift:    "Day",
		// 		}, nil).AnyTimes()
		// 	},
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockInterface := repository.NewMockRepoInterface(mc)
			redisInterface := cacheier.NewMockRedInterface(mc)
			tt.mockRepoResponse(mockInterface)
			tt.mockRedisRes(redisInterface)

			s := NewStore(mockInterface, redisInterface)

			got, err := s.FIlterApplication(tt.args.ctx, tt.args.applicantList)
			if (err != nil) != tt.wantErr {
				t.Errorf("Store.FIlterApplication() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, got, tt.want)
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("Store.FIlterApplication() = %v, want %v", got, tt.want)
			// }
		})
	}
}
