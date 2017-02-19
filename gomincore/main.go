package main

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

func FileMincore(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	finfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	size := int(finfo.Size())
	data, err := syscall.Mmap(int(file.Fd()), 0, size, syscall.PROT_READ, syscall.MAP_SHARED)
	if err != nil {
		return nil, err
	}
	defer syscall.Munmap(data)

	pages := (size + os.Getpagesize() - 1)/os.Getpagesize() // Number of pages in file. 
	result := make([]byte, pages)
	ret, _, _ := syscall.Syscall(syscall.SYS_MINCORE, uintptr(unsafe.Pointer(&data[0])), uintptr(size), uintptr(unsafe.Pointer(&result[0])))
	if ret != 0 {
		return nil, err
	}

	return result, nil
}

func main() {
	res, err := FileMincore(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
