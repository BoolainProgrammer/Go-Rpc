package util

import (
	"crypto/rand"
	"io"
)

var randowCodes = [...]byte{
	'1','2','3','4','5','6','7','8','9','0',
}

func GenValidateCode(max int) string {
	b := make([]byte, max) // [100,53,10,200,12,34]

	io.ReadAtLeast(rand.Reader, b, max)

	for i := 0; i < max; i++ {
		b[i] = randowCodes[int(b[i]) % len(randowCodes)]
	}

	return string(b)
}
