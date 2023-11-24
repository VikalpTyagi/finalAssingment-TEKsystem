package pkg

import "math/rand"

func OtpGenerator() int{
	return rand.Intn(10000)
}
