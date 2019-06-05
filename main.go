package main

import (

	"github.com/imorph/string-unpacker/pkg/unpkr"
)

func main() {
	fmt.Println("lets UNPACK!")
	fmt.Println("Press any time CTRL-C to exit")

	var pkgString PackedString

	for {
		fmt.Print("Please enter a correct packed string: ")

		_, err := fmt.Scanf("%d", &pkgString)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println("Unpacked string: ", pkgString.Unpack())
		}
		}
	}

}