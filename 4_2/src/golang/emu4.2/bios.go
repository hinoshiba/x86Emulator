package main

import (
	"fmt"
)

var (
	bios_to_terminal [8]int8 = [8]int8{30, 34, 32, 36, 31, 35, 33, 37}
)

func put_string(s []byte, n uint32) {
	var i uint32
	for i = 0; i < n; i++ {
		io_out8(0x03f8, s[i])
	}
}

func bios_video_teletype(emu *Emulator) {
	var color uint8 = get_register8(emu, BL) & 0x0f
	var ch uint8 = get_register8(emu, AL)

	var terminal_color int8 = bios_to_terminal[color & 0x07]
	var bright int8 = 0
	if (color & 0x08) != 0 {
		bright = 1
	}
	var buf []byte = []byte(fmt.Sprintf("\x1b[%d;%dm%c\x1b[0m", bright, terminal_color, ch))
	var length uint32 = uint32(len(buf))

	put_string(buf, length)
}

func bios_video(emu *Emulator) error {
	var fn uint8 = get_register8(emu, AH)

	switch fn {
	case 0x0e :
		bios_video_teletype(emu)
	default :
		return fmt.Errorf("not implemented BIOS video function 0x%02x\n", fn)
	}
	return nil
}
