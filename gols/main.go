package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var dir string

	if len(os.Args) > 1 {
		dir = os.Args[1]
	} else {
		dir = "."
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {

		log.Fatal(err)
	}

	for _, file := range files {
		s := file.ModTime().Format("_2 Jan 2006 15:04:05")
		fmt.Printf("%s %s %-15d %s\n", file.Mode(), s, file.Size(), file.Name())
	}
}
