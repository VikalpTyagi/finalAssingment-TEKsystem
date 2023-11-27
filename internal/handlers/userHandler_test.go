// * user related handlers endpoints

package handlers

import (
	"bytes"
	"context"
	"errors"
	"finalAssing/internal/middleware"
	"finalAssing/internal/models"
	"finalAssing/internal/services"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/mock/gomock"
)

func init() {
	os.Setenv("APP_PORT", "8080")
}

func Test_handler_Signup(t *testing.T) {
	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, services.Service)
		expectedStatusCode int
		ExpectedResponse   string
	}{
		{
			name: "Error: Tracker Id missing",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				c.Request = httptest.NewRequest("POST", "http://test.com:8080", nil)
				return c, rr, nil
			},
			expectedStatusCode: 500,
			ExpectedResponse:   `{"msg":"Internal Server Error"}`,
		},
		{
			name: "Error: Invalid json",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest := httptest.NewRequest("POST", "http://test.com:8080", bytes.NewBufferString("{Invalid Body}"))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TrackerIdKey, "12")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest
				return c, rr, nil
			},
			expectedStatusCode: 500,
			ExpectedResponse:   `{"msg in decoder":"Internal Server Error"}`,
		},
		{
			name: "Error: Validation Failed",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest := httptest.NewRequest("Post", "https://test.com:8080", bytes.NewBufferString(`{
					"Name":"Vikalp Tyagi"
				}`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TrackerIdKey, "12")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest
				return c, rr, nil
			},
			expectedStatusCode: 400,
			ExpectedResponse:   `{"msg":"please provide Name, Email and Password"}`,
		},
		{
			name: "Error: Mocked method fail",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest := httptest.NewRequest("Post", "https://test.com:8080", bytes.NewBufferString(`{
					"name":"Vikalp Tyagi",
					"email":"vikalp@gmail.com",
					"dateOfBirth":"15-05-1999",
					"password": "vikalp123"
				}`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TrackerIdKey, "12")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				mc := gomock.NewController(t)
				ms := services.NewMockService(mc)
				ms.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(models.User{}, errors.New("test error")).Times(1)
				return c, rr, ms
			},
			expectedStatusCode: 400,
			ExpectedResponse:   `{"msg":"user signup failed"}`,
		},
		{
			name: "Successful",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest := httptest.NewRequest("Post", "https://test.com:8080", bytes.NewBufferString(`{
					"name":"Vikalp Tyagi",
					"email":"vikalp@gmail.com",
					"dateOfBirth":"15-05-1999",
					"password": "vikalp123"
				}`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TrackerIdKey, "12")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				mc := gomock.NewController(t)
				ms := services.NewMockService(mc)
				ms.EXPECT().CreateUser(ctx, gomock.Any()).Return(models.User{
					Name:     "Vikalp Tyagi",
					Email:    "vikalp@gmail.com",
					DOB:      "15-05-1999",
					PassHash: "vikalp123",
				}, nil).Times(1)
				return c, rr, ms
			},
			expectedStatusCode: 200,
			ExpectedResponse:   `{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"name":"Vikalp Tyagi","email":"vikalp@gmail.com","dateOfBirth":"15-05-1999"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()
			h := &handler{
				s: ms,
			}
			h.Signup(c)

			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.ExpectedResponse, rr.Body.String())
		})
	}
}

func Test_handler_Login(t *testing.T) {
	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, services.Service)
		expectedStatusCode int
		ExpectedResponse   string
	}{
		{
			name: "Error: Tracker Id missing",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				c.Request = httptest.NewRequest("POST", "http://test.com:8080", nil)
				return c, rr, nil
			},
			expectedStatusCode: 500,
			ExpectedResponse:   `{"msg":"Internal Server Error"}`,
		},
		{
			name: "Error: Invalid json",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest := httptest.NewRequest("POST", "http://test.com:8080", bytes.NewBufferString("{Invalid Body}"))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TrackerIdKey, "12")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest
				return c, rr, nil
			},
			expectedStatusCode: 500,
			ExpectedResponse:   `{"msg":"Internal Server Error"}`,
		},
		{
			name: "Error: Validation Failed",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest := httptest.NewRequest("Post", "https://test.com:8080", bytes.NewBufferString(`{
					"Name":"Vikalp Tyagi"
				}`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TrackerIdKey, "12")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest
				return c, rr, nil
			},
			expectedStatusCode: 400,
			ExpectedResponse:   `{"msg":"please provide Email and Password"}`,
		},
		{
			name: "Error: Mocked method fail",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest := httptest.NewRequest("Post", "https://test.com:8080", bytes.NewBufferString(`{
					"email":"vikalp@gmail.com",
					"password": "vikalp123"
				}`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TrackerIdKey, "12")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				mc := gomock.NewController(t)
				ms := services.NewMockService(mc)
				ms.EXPECT().Authenticate(gomock.Any(), gomock.Any(), gomock.Any()).Return(jwt.RegisteredClaims{}, errors.New("test error")).Times(1)
				return c, rr, ms
			},
			expectedStatusCode: 401,
			ExpectedResponse:   `{"msg":"login failed"}`,
		},
		{
			name: "Success Case",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest := httptest.NewRequest("Post", "https://test.com:8080", bytes.NewBufferString(`{
					"email":"vikalp@gmail.com",
					"password": "vikalp123"
				}`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TrackerIdKey, "12")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				mc := gomock.NewController(t)
				ms := services.NewMockService(mc)
				ms.EXPECT().Authenticate(gomock.Any(), gomock.Any(), gomock.Any()).Return(jwt.RegisteredClaims{}, nil).Times(1)
				return c, rr, ms
			},
			expectedStatusCode: 200,
			ExpectedResponse:   `{"msg":"login failed"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()
			h := &handler{
				s: ms,
			}
			h.Login(c)
			assert.Equal(t, tt.ExpectedResponse, rr.Body.String())
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
		})
	}
}

func Test_handler_ForgetPassword(t *testing.T) {
	
	tests := []struct {
		name string
		setup func ()  (*gin.Context, *httptest.ResponseRecorder, services.Service)
		ExpectedResponse string
		expectedStatusCode int
		
	}{
		{
			name: "Error: Tracker Id missing",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				c.Request = httptest.NewRequest("POST", "http://test.com:8080", nil)
				return c, rr, nil
			},
			expectedStatusCode: 500,
			ExpectedResponse:   `{"msg":"Internal Server Error"}`,
		},
		{
			name: "Error: Invalid json",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest := httptest.NewRequest("POST", "http://test.com:8080", bytes.NewBufferString("{Invalid Body}"))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TrackerIdKey, "12")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest
				return c, rr, nil
			},
			expectedStatusCode: 500,
			ExpectedResponse:   `{"msg":"Internal Server Error"}`,
		},
		{
			name: "Error: Validation Failed",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest := httptest.NewRequest("POST", "https://test.com:8080", bytes.NewBufferString(`{
					"dateOfBirth":"15-05-1999",
					"email":"vikalptyagi15@gmail.com"
				}`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TrackerIdKey, "12")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest
				return c, rr, nil
			},
			expectedStatusCode: 400,
			ExpectedResponse:   `{"Error":"provided Invalid data"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms :=tt.setup()
			h := &handler{
				s:ms,
			}
			h.ForgetPassword(c)
			assert.Equal(t, tt.ExpectedResponse,rr.Body.String())
			assert.Equal(t, tt.expectedStatusCode,rr.Code)
		})
	}
}
