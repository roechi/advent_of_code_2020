package main

import (
	. "./joltage"
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestAdapterChain(t *testing.T) {
	adapterRatings := []int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4}

	oneCount, twoCount, threeCount := AdapterChain(adapterRatings)

	assert.Equal(t, oneCount, 7)
	assert.Equal(t, threeCount, 5)
	fmt.Println("Distribution: ", oneCount, twoCount, threeCount)
}

func TestAdapterChain2(t *testing.T) {
	adapterRatings := []int{28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45, 19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3}

	oneCount, twoCount, threeCount := AdapterChain(adapterRatings)

	assert.Equal(t, oneCount, 22)
	assert.Equal(t, threeCount, 10)
	fmt.Println("Two count: ", twoCount)
	fmt.Println("Distribution: ", oneCount, twoCount, threeCount)
}
