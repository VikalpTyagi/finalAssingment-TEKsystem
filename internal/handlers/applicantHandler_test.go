package handlers

import (
	"finalAssing/internal/services"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func Test_handler_AcceptApplicant(t *testing.T) {
	type args struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, services.Service)
		expectedStatusCode int
		ExpectedResponse   string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, rr, ms:= tt.args.setup()
			h := &handler{
				s: ms,
			}
			h.AcceptApplicant(c)
			assert.Equal(t,tt.args.expectedStatusCode,rr.Code)
			assert.Equal(t,tt.args.ExpectedResponse,rr.Body.String())

		})
	}
}
