package errmanagment

import (
	"fmt"
	"os"
)

func OpenFile() {

	var fileName string
	fmt.Println("Name of file: ")
	fmt.Scanln(&fileName)

	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755)

	defer func() {
		fmt.Print("Closing file")
		file.Close()
	}()

	file.Write([]byte("Yes"))
	if err != nil {
		panic(err.Error())
	}

}
