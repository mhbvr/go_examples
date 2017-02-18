package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"flag"
)


func hrsize(size int64) string {
	var value int64 = size
	var postfix string = ""

	switch {
	case size > 1 << 40:
		postfix = "TB"
		value = size / (1 << 40)
	case size > 1 << 30:
		postfix = "GB"
		value = size / (1 << 30)
	case size > 1 << 20:
		postfix = "MB"
		value = size / (1 << 20)
	case size > 1 << 10:
		postfix = "KB"
		value = size / (1 << 10)
	}
	return fmt.Sprintf("%v%v", value, postfix)
}


func main() {

    //var rsize = flag.Bool("rsize", false, "Calculate directory size recursively")
	var hr = flag.Bool("hr", false, "Print human readable size (ie M/K/G bytes)")

	flag.Parse()
	args := flag.Args()

	var dir string

	if len(args) > 0 {
		dir = args[0]
	} else {
		dir = "."
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {

		log.Fatal(err)
	}

	for _, file := range files {
		s := file.ModTime().Format("_2 Jan 2006 15:04:05")
		if *hr {
			fmt.Printf("%11s  %s %7s %s\n", file.Mode(), s, hrsize(file.Size()), file.Name())
		} else {
			fmt.Printf("%11s  %s %-15d %s\n", file.Mode(), s, file.Size(), file.Name())
		}
	}
}
