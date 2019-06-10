package main

import (
	"fmt"
	"log"

	u "github.com/string-unpacker/pkg/unpkr"
)


func main() {
	fmt.Println("\nlet's UNPACK!\nPress any time CTRL-C to exit")
	fmt.Println("")

	var pkgString PackedString

	for {
		fmt.Print("Please enter a correct packed string: ")

		_, err := fmt.Scanf("%s", &pkgString)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println("Unpacked string: ", pkgString.Unpack())
			fmt.Println("")
		}
	}
}
