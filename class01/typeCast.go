package main

import (
	"fmt"
	"unsafe"
)

func main() {
	value := 2000
	ptr := &value
	transformToByte(ptr)
	fmt.Printf("%b\n", value)
	fmt.Printf("Endereço como uintptr: %d\n", uintptr(unsafe.Pointer(ptr)))

}

func transformToByte(i *int) {
	*i = int(uint8(*i))
}
