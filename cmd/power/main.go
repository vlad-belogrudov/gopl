// Declare power numbers, KB, MB ... which are multiples of 1000
package main

import "fmt"

const (
	KB = 1000
	MB = KB * 1000
	GB = MB * 1000
	TB = GB * 1000
)

func main() {
	fmt.Println("KB = ", KB)
	fmt.Println("MB = ", MB)
	fmt.Println("GB = ", GB)
	fmt.Println("TB = ", TB)

}
