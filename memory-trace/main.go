package main

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
)

func generateRandBytes(length int) ([]byte, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	return b, err
}

func main() {
	buf := &bytes.Buffer{}

	for i := 0; i < 22; i++ {
		b, err := generateRandBytes(int(math.Pow(2, float64(i))))
		if err != nil {
			panic(err)
		}
		buf.Write(b)
		fmt.Printf("bytes buffer capacity %d\n", buf.Cap())
		buf.Reset()
		fmt.Printf("bytes buffer capacity after reset %d\n", buf.Cap())
	}

	for i := 22; i > 0; i-- {
		b, err := generateRandBytes(int(math.Pow(2, float64(i))))
		if err != nil {
			panic(err)
		}
		buf.Write(b)
		fmt.Printf("bytes buffer capacity %d\n", buf.Cap())
		buf.Reset()
		fmt.Printf("bytes buffer capacity after reset %d\n", buf.Cap())
	}
}
