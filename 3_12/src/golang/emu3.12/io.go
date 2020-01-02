package main

import (
	"os"
	"fmt"
	"bufio"
)

func io_in8(address uint16) uint8 {
	switch address {
	case 0x03f8:
		return getchar()
	default :
		return 0x0
	}
}

func io_out8(address uint16, value uint8) {
	switch address {
	case 0x03f8:
		fmt.Printf("%s", string(value))
	default:
	}
}

func getchar() uint8 {
	r := bufio.NewReader(os.Stdin)
	in, err := r.ReadString('\n')
	if err != nil {
		return 0x00
	}

	in_b := []byte(in)
	if len(in_b) < 1 {
		return 0x00
	}
	if '\n' == in_b[0] {
		return 0x00
	}
	return in_b[0]
}
