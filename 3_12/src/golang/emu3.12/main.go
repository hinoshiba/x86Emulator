package main

import (
	"os"
	"fmt"
	"errors"
)

const MEMORY_SIZE uint32 = 1024 * 1024

var registers_name = []string{ "EAX", "ECX", "EDX",
									"EBX", "ESP", "EBP", "ESI", "EDI"}

func create_emu(size uint32, eip uint32, esp uint32) *Emulator {
	var emu *Emulator = &Emulator{}
	emu.memory = make([]byte, size)

	emu.eip = eip
	emu.registers[ESP] = esp

	return emu
}

func destroy_emu(emu *Emulator) {
	emu = nil
}

func dump_registers(emu *Emulator) {
	for i, v := range emu.registers {
		fmt.Printf("%s = %08x\n", registers_name[i], v);
	}

	fmt.Printf("EIP = %08x\n", emu.eip)
}

func run() error {

	var quiet bool = true
	var args []string
	for _, a := range os.Args {
		if a == "-q" {
			quiet = false
			continue
		}
		args = append(args, a)
	}


	if len(args) < 2 {
		return errors.New("usage: px86 <filename>\n")
	}

	emu := create_emu(MEMORY_SIZE, 0x7c00, 0x7c00)

	f, err := os.OpenFile(args[1], os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	buf := make([]byte, 0x200)
	n, err := f.Read(buf)
	if err != nil {
		return err
	}
	for i := 0; i < n; i++ {
		emu.memory[i + 0x7c00] = buf[i]
	}

	init_instructions()

	for emu.eip < MEMORY_SIZE {
		var code uint8 = get_code8(emu, 0)
		if quiet {
			fmt.Printf("EIP = %X, Code = %02X\n", emu.eip, code)
		}

		if instructions[code] == nil {
			fmt.Printf("\n\nNot Implemented: %x\n", code)
			break
		}

		if err := instructions[code](emu); err != nil {
			return err
		}

		if emu.eip == 0x00 {
			fmt.Printf("\n\nend of program.\n\n")
			break
		}
	}

	dump_registers(emu)
	destroy_emu(emu)
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
}
