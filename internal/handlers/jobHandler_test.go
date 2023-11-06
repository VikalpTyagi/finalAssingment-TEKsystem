// * job related handlers endpoint

package handlers

import (
	"bytes"
	"context"
	"errors"
	"finalAssing/internal/middleware"
	"finalAssing/internal/models"
	"finalAssing/internal/services"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"go.uber.org/mock/gomock"
)

func Test_handler_fetchJobById(t *testing.T) {
	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, services.Service)
		expectedStatusCode int
		ExpectedResponse   string
	}{
		{
			name: "Tracker Id missing",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				c.Request = httptest.NewRequest("GET", "http://test.com", nil)
				return c, rr, nil
			},
			expectedStatusCode: 500,
			ExpectedResponse:   `{"msg":"Internal Server Error"}`,
		},
		{
			name: "Invalid Job Id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080/12", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TrackerIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest
				c.Params = append(c.Params, gin.Param{Key: "ID", Value: "ab2"})

				mc := gomock.NewController(t)
				ms := services.NewMockService(mc)
				ms.EXPECT().GetJobById(gomock.Any(), gomock.Any()).Return(models.Job{}, errors.New("test error")).Times(1)
				return c, rr, ms
			},
			expectedStatusCode: 400,
			ExpectedResponse:   `{"msg":"Bad Request"}`,
		},
		{
			name: "Sucessful",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080/12", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TrackerIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest
				c.Params = append(c.Params, gin.Param{Key: "ID", Value: "786"})

				mc := gomock.NewController(t)
				ms := services.NewMockService(mc)
				ms.EXPECT().GetJobById(gomock.Any(), gomock.Any()).Return(models.Job{}, nil).Times(1)
				return c, rr, ms
			},
			expectedStatusCode: 200,
			ExpectedResponse:   `{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"title":"","field":"","experience":0,"companyId":0}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()
			h := &handler{
				s: ms,
			}
			h.fetchJobById(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.ExpectedResponse, rr.Body.String())
		})
	}
}

func Test_handler_addJobsById(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, services.Service)
		expectedStatusCode int
		ExpectedResponse   string
	}{
		{
			name: "Tracker Id missing",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				c.Request = httptest.NewRequest("GET", "http://test.com", nil)
				return c, rr, nil
			},
			expectedStatusCode: 500,
			ExpectedResponse:   `{"msg":"Internal Server Error"}`,
		},
		{
			name: "body data empty",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				c.Request = httptest.NewRequest("GET", "http://test.com", bytes.NewBufferString(" "))
				return c, rr, nil
			},
			expectedStatusCode: 500,
			ExpectedResponse:   `{"msg":"Internal Server Error"}`,
		},
		{
			name: "Invalid company Id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080/12", bytes.NewBufferString(`[]`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TrackerIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest
				c.Params = append(c.Params, gin.Param{Key: "ID", Value: "12"})

				mc := gomock.NewController(t)
				ms := services.NewMockService(mc)
				ms.EXPECT().JobByCompanyId(gomock.Any(), gomock.Any()).Return(nil, errors.New("test error")).Times(1)
				return c, rr, ms
			},
			expectedStatusCode: 400,
			ExpectedResponse:   `{"msg":"Job creation failed"}`,
		},
		{
			name: "Successful",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080/12", bytes.NewBufferString(`[{
					"title":  "GO dev",
					"field": "IT",
				},]`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TrackerIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest
				c.Params = append(c.Params, gin.Param{Key: "ID", Value: "12"})
				mc := gomock.NewController(t)
				ms := services.NewMockService(mc)
				ms.EXPECT().JobByCompanyId(gomock.Any(), gomock.Any()).Return([]models.Job{{
					Name:  "GO dev",
					Field: "IT",
				},}, nil).AnyTimes()
				return c, rr, ms
			},
			expectedStatusCode: 201,
			ExpectedResponse:   `[{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"title":"GO dev","field":"IT","experience":0,"companyId":0}]`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()
			h := &handler{
				s: ms,
			}
			h.addJobsById(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.ExpectedResponse, rr.Body.String())
		})
	}
}
