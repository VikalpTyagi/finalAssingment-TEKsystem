// * Zerolog logger middleware
package middleware

import "github.com/gin-gonic/gin"


type key string
const TrackerId = "1"
func Logger() gin.HandlerFunc {
	return func(c gin.Context) {
		
	}
}