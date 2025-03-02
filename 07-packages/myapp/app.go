package main

import (
	"fmt"
	str "strings" // Package Alias

	"github.com/AmirHosein-Vahed/golang-tutorials/07-packages/numbers"
	"github.com/AmirHosein-Vahed/golang-tutorials/07-packages/strings"
	"github.com/AmirHosein-Vahed/golang-tutorials/07-packages/strings/greeting" // Importing a nested package
)

func main() {
	fmt.Println(numbers.IsPrime(19))

	fmt.Println(greeting.WelcomeText)

	fmt.Println(strings.Reverse("callicoder"))

	fmt.Println(str.Count("Go is Awesome. I love Go", "Go"))
}
