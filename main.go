package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
)

func getRand() (int64, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(15))
	if err != nil {
		return 0, err
	}

	return n.Int64() + 1, nil
}

func main() {
	nums := make(map[string]int64)
	keys := []string{"Morning", "Matinee", "Midday", "Evening", "Late Night"}
	for _, key := range keys {
		n, err := getRand()
		if err != nil {
			log.Fatal(err)
		}
		nums[key] = n
	}

	for key, val := range nums {
		fmt.Printf("%10s: %2d\n", key, val)
	}
}
