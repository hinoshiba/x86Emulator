package main

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

func set_memory8(emu *Emulator, address uint32, value uint8) {
	emu.memory[address] = value
}

func set_memory32(emu *Emulator, address uint32, value uint32) {
	var i uint32
	for i = 0; i < 4; i++ {
		set_memory8(emu, address + i, uint8(value >> (i * 8)))
	}
}

func get_memory8(emu *Emulator, address uint32) uint8 {
	return emu.memory[address]
}

func get_memory32(emu *Emulator, address uint32) uint32 {
	var i int32
	var ret uint32

	for i = 0; i < 4; i++ {
		ret |= uint32(get_memory8(emu, address + uint32(i))) << uint32(8 * i)
	}
	return ret
}

func get_register32(emu *Emulator, index uint8) uint32 {
	return emu.registers[index]
}

func set_register32(emu *Emulator, index uint8, value uint32) {
	emu.registers[index] = value
}
