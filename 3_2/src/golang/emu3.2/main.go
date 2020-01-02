package main

import (
	"os"
	"fmt"
	"errors"
)

const MEMORY_SIZE uint32 = 1024 * 1024
const (
	EAX int32 = iota
	ECX
	EDX
	EBX
	ESP
	EBP
	ESI
	EDI
	REGISTERS_COUNT
)

var registers_name = []string{ "EAX", "ECX", "EDX",
									"EBX", "ESP", "EBP", "ESI", "EDI"}

type Emulator struct {
	registers  [REGISTERS_COUNT]uint32
	eflags     uint32
	memory     []byte
	eip        uint32
}

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

func get_code8(emu *Emulator, index int32) uint8 {
	return emu.memory[int32(emu.eip) + index]
}

func get_sign_code8(emu *Emulator, index int32) int8 {
	return int8(emu.memory[int32(emu.eip) + index])
}

func get_sign_code32(emu *Emulator, index int32) int32 {
	return int32(get_code32(emu, index))
}

func get_code32(emu *Emulator, index int32) uint32 {
	var i uint32
	var ret uint32

	for i = 0; i < 4; i++ {
		ret |= uint32(get_code8(emu, index + int32(i))) << (i * 8)
	}

	return ret
}

func mov_r32_imm32(emu *Emulator) {
	var reg uint8 = get_code8(emu, 0) - 0xB8
	var value uint32 = get_code32(emu, 1)
	emu.registers[reg] = value
	emu.eip += 5
}

func short_jump(emu *Emulator) {
	var diff int8
	diff = get_sign_code8(emu, 1)
	emu.eip += uint32(diff + 2)
}

func near_jump(emu *Emulator) {
	var diff int32
	diff = get_sign_code32(emu, 1)
	emu.eip += uint32(diff + 5)
}

type instruction_func_t func(*Emulator)
var instructions = [256]instruction_func_t{}

func init_instructions() {
	for i := 0; i < 8; i++ {
		instructions[0xB8 + i] = mov_r32_imm32
	}
	instructions[0xE9] = near_jump
	instructions[0xEB] = short_jump
}

func run() error {
	if len(os.Args) < 2 {
		return errors.New("usage: px86 <filename>\n")
	}

	emu := create_emu(MEMORY_SIZE, 0x7c00, 0x7c00)

	f, err := os.OpenFile(os.Args[1], os.O_RDONLY, 0666)
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
		fmt.Printf("EIP = %X, Code = %02X\n", emu.eip, code)

		if instructions[code] == nil {
			fmt.Printf("\n\nNot Implemented: %x\n", code)
			break
		}

		instructions[code](emu)

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
