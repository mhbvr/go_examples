package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
)

func hrsize(size int64) string {
	var value float64 = float64(size)
	var postfix string = ""

	switch {
	case size > 1<<40:
		postfix = "TB"
		value = value / (1 << 40)
	case size > 1<<30:
		postfix = "GB"
		value = value / (1 << 30)
	case size > 1<<20:
		postfix = "MB"
		value = value / (1 << 20)
	case size > 1<<10:
		postfix = "KB"
		value = value / (1 << 10)
	}
	return fmt.Sprintf("%.1f%v", value, postfix)
}

type walker struct {
	totSize int64
}

func (w *walker) walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		log.Printf("Error visiting %v: %v", path, err)
	} else {
		w.totSize = w.totSize + info.Size()
	}
	return nil
}

func main() {

	var rsize = flag.Bool("rsize", false, "Calculate directory size recursively")
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
		var size int64 = 0
		s := file.ModTime().Format("_2 Jan 2006 15:04:05")

		if *rsize {
			w := walker{0}
			err := filepath.Walk(path.Join(dir, file.Name()), w.walkFunc)
			if err != nil {
				log.Fatal(err)
			}
			size = w.totSize
		} else {
			size = file.Size()
		}
		if *hr {
			fmt.Printf("%11s  %s %7s %s\n", file.Mode(), s, hrsize(size), file.Name())
		} else {
			fmt.Printf("%11s  %s %15d %s\n", file.Mode(), s, size, file.Name())
		}
	}
}
