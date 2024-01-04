package main

import "C"

import (
	"math/big"

	"poseidon-go/poseidon"
)

//export Hash
func Hash(xStr *C.char, yStr *C.char) *C.char {
	// convert string to bigInt
	xBigInt, success := new(big.Int).SetString(C.GoString(xStr), 10)
	if !success {
		panic("Fail to convert poseidon hash input X into integer")
	}
	yBigInt, success := new(big.Int).SetString(C.GoString(yStr), 10)
	if !success {
		panic("Fail to convert poseidon hash input Y into integer")
	}

	// calculate poseidon and return
	result, err := poseidon.HashFixed([]*big.Int{xBigInt, yBigInt})
	if err != nil {
		panic(err)
	}
	return C.CString(result.String())
}

func main() {}
