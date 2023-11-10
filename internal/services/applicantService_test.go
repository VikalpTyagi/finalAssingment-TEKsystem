package services

import (
	"context"
	"finalAssing/internal/models"
	"finalAssing/internal/repository"
	"testing"

	"go.uber.org/mock/gomock"
	"gopkg.in/go-playground/assert.v1"
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
		mockRepoResponse func (repoMock *repository.MockRepoInterface)
	}{

		{
			name: "multiple job id",
			args: args{
				ctx: context.Background(),
				applicantList: []*models.ApplicantReq{
					{ //* problem NP
						Name:           "Vikalp Tyagi",
						JobId:          1,
						Experience:     3,
						Max_NP:         8,
						Budget:         50000,
						Locations:      []uint{1, 2},
						Stack:          []uint{1, 2, 3},
						WorkMode:       "Full-Time",
						Qualifications: []uint{1},
						Shift:          "Day",
					},
					{ //* problem : Shift
						Name:           "Vikalp Tyagi",
						JobId:          2,
						Experience:     3,
						Max_NP:         2,
						Budget:         50000,
						Locations:      []uint{1, 2},
						Stack:          []uint{1, 2, 3},
						WorkMode:       "Full-Time",
						Qualifications: []uint{1},
						Shift:          "Night",
					},
					{ //* Problem: Workmode
						Name:           "Vikalp Tyagi",
						JobId:          3,
						Experience:     3,
						Max_NP:         2,
						Budget:         50000,
						Locations:      []uint{1, 2},
						Stack:          []uint{1, 2, 3},
						WorkMode:       "Full-Time",
						Qualifications: []uint{1},
						Shift:          "Day",
					},
				},
			},
			want:    nil,
			wantErr: false,
			mockRepoResponse: func (repoMock *repository.MockRepoInterface) {
				
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			mockInterface := repository.NewMockRepoInterface(mc)
			mockInterface.EXPECT().ApplicantsFilter(gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()
			s := NewStore(mockInterface)

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

// func TestStore_FIlterApplication(t *testing.T) {
// 	type args struct {
// 		ctx           context.Context
// 		applicantList []*models.ApplicantReq
// 	}
// 	tests := []struct {
// 		name             string
// 		args             args
// 		want             []*models.ApplicantRespo
// 		wantErr          bool
// 		mockRepoResponse func() (*models.Job, error)
// 	}{
// 		{
// 			name: "Success",
// 			args: args{
// 				ctx: context.Background(),
// 				applicantList: []*models.ApplicantReq{
// 					{
// 						Name:           "Vikalp Tyagi",
// 						JobId:          1,
// 						Experience:     3,
// 						Max_NP:         2,
// 						Budget:         50000,
// 						Locations:      []uint{1, 2},
// 						Stack:          []uint{1, 2, 3},
// 						WorkMode:       "Full-Time",
// 						Qualifications: []uint{1},
// 						Shift:          "Day",
// 					},
// 					{
// 						Name:           "Satyam",
// 						JobId:          1,
// 						Experience:     3,
// 						Max_NP:         2,
// 						Budget:         6000,
// 						Locations:      []uint{1, 2},
// 						Stack:          []uint{1, 2, 3},
// 						WorkMode:       "Full-Time",
// 						Qualifications: []uint{1},
// 						Shift:          "Day",
// 					},
// 					{
// 						Name:           "ajay",
// 						JobId:          1,
// 						Experience:     3,
// 						Max_NP:         2,
// 						Budget:         600000,
// 						Locations:      []uint{1, 2},
// 						Stack:          []uint{1, 2, 3},
// 						WorkMode:       "Full-Time",
// 						Qualifications: []uint{1},
// 						Shift:          "Day",
// 					},
// 					{
// 						Name:           "Ramesh",
// 						JobId:          1,
// 						Experience:     3,
// 						Max_NP:         2,
// 						Budget:         40000,
// 						Locations:      []uint{1, 2},
// 						Stack:          []uint{1, 2, 3},
// 						WorkMode:       "Full-Time",
// 						Qualifications: []uint{1},
// 						Shift:          "Day",
// 					},
// 				},
// 			},
// 			want: []*models.ApplicantRespo{
// 				{
// 					Name:  "Vikalp Tyagi",
// 					JobId: 1,
// 				},
// 				{
// 					Name:  "Satyam",
// 					JobId: 1,
// 				},
// 				{
// 					Name:  "Ramesh",
// 					JobId: 1,
// 				},
// 			},
// 			wantErr: false,
// 			mockRepoResponse: func() (*models.Job, error) {
// 				return &models.Job{
// 					Experience: 4,
// 					MinExp:     1,
// 					Min_NP:     1,
// 					Max_NP:     4,
// 					Budget:     80000,
// 					Stack: []models.Skill{
// 						{
// 							Model: gorm.Model{ID: 1},
// 						},
// 						{
// 							Model: gorm.Model{ID: 2},
// 						},
// 						{
// 							Model: gorm.Model{ID: 3},
// 						},
// 					},
// 					Locations: []models.Location{
// 						{
// 							Model: gorm.Model{ID: 1},
// 						},
// 						{
// 							Model: gorm.Model{ID: 2},
// 						},
// 						{
// 							Model: gorm.Model{ID: 3},
// 						},
// 					},
// 					Qualifications: []models.Qualification{
// 						{
// 							Model: gorm.Model{ID: 1},
// 						},
// 					},
// 					WorkMode: "Full-Time",
// 					Shift:    "Day",
// 				}, nil
// 			},
// 		},
// 		{
// 			name: "Invalid job Id",
// 			args: args{
// 				ctx: context.Background(),
// 				applicantList: []*models.ApplicantReq{
// 					{
// 						Name:           "Vikalp Tyagi",
// 						JobId:          2,
// 						Experience:     3,
// 						Max_NP:         2,
// 						Budget:         50000,
// 						Locations:      []uint{1, 2},
// 						Stack:          []uint{1, 2, 3},
// 						WorkMode:       "Full-Time",
// 						Qualifications: []uint{1},
// 						Shift:          "Day",
// 					},
// 				},
// 			},
// 			want:    nil,
// 			wantErr: false,
// 			mockRepoResponse: func() (*models.Job, error) {
// 				return nil, errors.New("tets error")
// 			},
// 		},

// 		{
// 			name: "budget requirments not met",
// 			args: args{
// 				ctx: context.Background(),
// 				applicantList: []*models.ApplicantReq{
// 					{
// 						Name:           "Vikalp Tyagi",
// 						JobId:          1,
// 						Experience:     3,
// 						Max_NP:         2,
// 						Budget:         50000,
// 						Locations:      []uint{1, 2},
// 						Stack:          []uint{1, 2, 3},
// 						WorkMode:       "Full-Time",
// 						Qualifications: []uint{1},
// 						Shift:          "Day",
// 					},
// 				},
// 			},
// 			want:    nil,
// 			wantErr: false,
// 			mockRepoResponse: func() (*models.Job, error) {
// 				return nil, errors.New("hvjygfc")
// 			},
// 		},
// 		{
// 			name: "Experience requirments not met",
// 			args: args{
// 				ctx: context.Background(),
// 				applicantList: []*models.ApplicantReq{
// 					{
// 						Name:           "Vikalp Tyagi",
// 						JobId:          2,
// 						Experience:     3,
// 						Max_NP:         2,
// 						Budget:         50000,
// 						Locations:      []uint{1, 2},
// 						Stack:          []uint{1, 2, 3},
// 						WorkMode:       "Full-Time",
// 						Qualifications: []uint{1},
// 						Shift:          "Day",
// 					},
// 				},
// 			},
// 			want:    nil,
// 			wantErr: false,
// 			mockRepoResponse: func() (*models.Job, error) {
// 				return &models.Job{}, errors.New("test error")
// 			},
// 		},
// 		{
// 			name: "multiple job id",
// 			args: args{
// 				ctx: context.Background(),
// 				applicantList: []*models.ApplicantReq{
// 					{ //* problem NP
// 						Name:           "Vikalp Tyagi",
// 						JobId:          1,
// 						Experience:     3,
// 						Max_NP:         8,
// 						Budget:         50000,
// 						Locations:      []uint{1, 2},
// 						Stack:          []uint{1, 2, 3},
// 						WorkMode:       "Full-Time",
// 						Qualifications: []uint{1},
// 						Shift:          "Day",
// 					},
// 					{//* problem : Shift
// 						Name:           "Vikalp Tyagi",
// 						JobId:          1,
// 						Experience:     3,
// 						Max_NP:         2,
// 						Budget:         50000,
// 						Locations:      []uint{1, 2},
// 						Stack:          []uint{1, 2, 3},
// 						WorkMode:       "Full-Time",
// 						Qualifications: []uint{1},
// 						Shift:          "Night",
// 					},
// 					{//* Problem: Workmode
// 						Name:           "Vikalp Tyagi",
// 						JobId:          1,
// 						Experience:     3,
// 						Max_NP:         2,
// 						Budget:         50000,
// 						Locations:      []uint{1, 2},
// 						Stack:          []uint{1, 2, 3},
// 						WorkMode:       "Full-Time",
// 						Qualifications: []uint{1},
// 						Shift:          "Day",
// 					},
// 				},
// 			},
// 			want:    nil,
// 			wantErr: false,
// 			mockRepoResponse: func() () {

// 			},
// 		},

// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mc := gomock.NewController(t)
// 			mockInterface := repository.NewMockRepoInterface(mc)
// 			mockInterface.EXPECT().ApplicantsFilter(gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()
// 			s := NewStore(mockInterface)

// 			got, err := s.FIlterApplication(tt.args.ctx, tt.args.applicantList)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Store.FIlterApplication() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			assert.Equal(t, got, tt.want)
// 			// if !reflect.DeepEqual(got, tt.want) {
// 			// 	t.Errorf("Store.FIlterApplication() = %v, want %v", got, tt.want)
// 			// }
// 		})
// 	}
// }
