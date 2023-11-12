package handlers

import (
	"bytes"
	"context"
	"finalAssing/internal/auth"
	"finalAssing/internal/middleware"
	"finalAssing/internal/services"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Test_handler_AcceptApplicant(t *testing.T) {
	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, services.Service)
		expectedStatusCode int
		ExpectedResponse   string
	}{
		{
			name: "Missing tracker Id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httptest, _ := http.NewRequest(http.MethodPost, "http://test.com", nil)
				c.Request = httptest
				return c, rr, nil
			},
			expectedStatusCode: 500,
			ExpectedResponse:   `{"msg":"Internal Server Error"}`,
		},
		{
			name: "Unccessful validation",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodPost, "http://test.com", bytes.NewBufferString(``))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TrackerIdKey,"12")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				return c, rr, nil
			},
			expectedStatusCode: 400,
			ExpectedResponse: `{"Error": "All fields are mandatory"}`,
		},
		{
			name: "Invalid Request Body",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, services.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodPost, "http://test.com", bytes.NewBufferString(`{bksdbv`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TrackerIdKey,"123")
				ctx = context.WithValue(ctx, auth.AuthKey, jwt.RegisteredClaims{})
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest
				return c, rr, nil
			},
			expectedStatusCode: 400,
			ExpectedResponse: `{"msg2":"Bad Request"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()
			h := &handler{
				s: ms,
			}
			h.AcceptApplicant(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.ExpectedResponse, rr.Body.String())

		})
	}
}
