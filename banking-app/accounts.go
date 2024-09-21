package main

import (
	"math/rand"
	"time"
)

// var arrayName [size]elementType

var accountNumbers [20]uint64
var accountBalances [20]float64

const charSet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generateAccountNumbers() {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	for i := range accountNumbers {
		accountNumbers[i] = r.Uint64()
	}
}

