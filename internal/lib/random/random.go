package random

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomString(n int) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	str := ""
	for i := 0; i < n; i++ {
		num := rnd.Intn(26) + 'a'
		r := rune(num)

		str += fmt.Sprintf("%c", r)
	}

	return str

}
