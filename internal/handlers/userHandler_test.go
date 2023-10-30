// * user related handlers endpoints

package handlers

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_handler_Signup(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		h    *handler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Signup(tt.args.c)
		})
	}
}
