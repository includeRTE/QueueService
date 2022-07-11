package ringbuf

import (
	"fmt"
	"math/rand"
	"testing"
)

var randArray = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

func TestWrite(t *testing.T) {

	buf := NewRingBuf()
	randArrayLen := len(randArray)
	for i := 0; i < 10000; i++ {
		randSize := rand.Intn(512)
		tempBuf := make([]byte, 0, randSize)
		for j := 0; j < randSize; j++ {
			randIndex := rand.Intn(randArrayLen)
			tempBuf = append(tempBuf, randArray[randIndex])
		}

		buf.Write(tempBuf)
		readBuf, _ := buf.ReadAll()
		fmt.Println(readBuf)
	}
}
