package main

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func getTermSize() (uint, uint) {
	ws := &winsize{}
	ret, _, err := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if int(ret) != 0 {
		panic(err)
	}
	return uint(ws.Col), uint(ws.Row)
}

func sum(b []byte) int {
	s := 0
	for _, v := range b {
		s = s + int(v)
	}
	return s
}

/* Scale vector of 0/1 bytes to smaller size */
func scale(vec []byte, size int) []float64 {
	var result []float64

	origSize := len(vec)
	scale := origSize/size + 1
	max := 0
	for i := 0; (i+1)*scale < origSize; i++ {
		result = append(result, float64(sum(vec[i*scale:(i+1)*scale]))/float64(scale))
		max = i
	}
	if (max+1)*scale < origSize {
		result = append(result, float64(sum(vec[(max+1)*scale:]))/float64(origSize-(max+1)*scale))
	}
	return result
}

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

	pages := (size + os.Getpagesize() - 1) / os.Getpagesize() // Number of pages in file.
	result := make([]byte, pages)
	ret, _, _ := syscall.Syscall(syscall.SYS_MINCORE, uintptr(unsafe.Pointer(&data[0])), uintptr(size), uintptr(unsafe.Pointer(&result[0])))
	if ret != 0 {
		return nil, err
	}

	return result, nil
}

func progressBar(b []float64) []rune {
	result := make([]rune, len(b))
	var sym rune
	for i, v := range b {
		switch {
		case v == 0:
			sym = ' '
		case v < 0.3:
			sym = 0x2591
		case v < 0.6:
			sym = 0x2592
		case v < 1:
			sym = 0x2593
		default:
			sym = 0x2588
		}
		result[i] = sym
	}
	return result
}

func main() {
	res, err := FileMincore(os.Args[1])
	if err != nil {
		panic(err)
	}

	cols, _ := getTermSize()

	fmt.Printf("Resident pages: %v/%v (%v%%)\n", sum(res), len(res), 100*sum(res)/len(res))
	fmt.Printf("[%v]\n", string(progressBar(scale(res, int(cols-2)))))
}
