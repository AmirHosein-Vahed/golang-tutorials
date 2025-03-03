package model

// Unexported struct (only accessible inside package `model`)
type Address struct {
	houseNo, street, city, state, country string
	zipCode                               int
}
