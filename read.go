package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	var filePath string
	fmt.Println("Enter file path: ")
	fmt.Scan(&filePath)

	slice := make([]Name, 0, 3)
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		var name Name
		name.fname, name.lname = s[0], s[1]
		slice = append(slice, name)
	}

	file.Close()

	for _, v := range slice {
		fmt.Println(v.fname, v.lname)
	}

}