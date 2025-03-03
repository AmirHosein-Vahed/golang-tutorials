package model

type Customer struct { // exported struct type
	Id      int        // exported field
	Name    string     // exported field
	addr    Address    // unexported field
	married bool       // unexported field (only accessible inside package `model`)
}
