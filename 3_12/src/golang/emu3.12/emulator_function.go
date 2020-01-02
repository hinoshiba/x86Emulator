package main

const (
	CARRY_FLAG uint32 = 1
	ZERO_FLAG uint32 = 1 << 6
	SIGN_FLAG uint32 = 1 << 7
	OVERFLOW_FLAG uint32 = 1 << 11
)

func get_code8(emu *Emulator, index uint8) uint8 {
	return emu.memory[emu.eip + uint32(index)]
}

func get_sign_code8(emu *Emulator, index uint8) int8 {
	return int8(emu.memory[emu.eip + uint32(index)])
}

func get_sign_code32(emu *Emulator, index uint8) int32 {
	return int32(get_code32(emu, index))
}

func get_code32(emu *Emulator, index uint8) uint32 {
	var i uint8
	var ret uint32

	for i = 0; i < 4; i++ {
		ret |= uint32(get_code8(emu, index + i)) << (i * 8)
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

func get_register8(emu *Emulator, index uint8) uint8 {
	if index < 4 {
		return uint8(emu.registers[index] & 0xff)
	}
	return uint8((emu.registers[index - 4] >> 8) & 0xff)
}

func get_register32(emu *Emulator, index uint8) uint32 {
	//fmt.Println("GET regid :", index, "value :", emu.registers[index])
	return emu.registers[index]
}

func set_register8(emu *Emulator, index uint8, value uint8) {
	if index < 4 {
		var r uint32 = uint32(emu.registers[index] & 0xffffff00)
		emu.registers[index] = r | uint32(value)
		return
	}
	var r uint32 = uint32(emu.registers[index - 4] & 0xffffff00)
	emu.registers[index - 4] = r | uint32(value << 8)
}

func set_register32(emu *Emulator, index uint8, value uint32) {
	//fmt.Println("SET regid :", index, "value :", value)
	emu.registers[index] = value
}

func push32(emu *Emulator, value uint32) {
	var address uint32 = get_register32(emu, uint8(ESP)) - 4
	set_register32(emu, uint8(ESP), address)
	set_memory32(emu, address, value)
	//fmt.Println("push :", value, "address :", address)
}

func pop32(emu *Emulator) uint32 {
	var address uint32 = get_register32(emu, uint8(ESP))
	var ret uint32 = get_memory32(emu, address)
	set_register32(emu, uint8(ESP), address + 4)
	//fmt.Println("pop :", ret, "address :", address)
	return ret
}

func update_eflags_sub(emu *Emulator, v1 uint32, v2 uint32, result uint64) {
	var sign1 int32 = int32(v1 >> 31)
	var sign2 int32 = int32(v2 >> 31)
	var signr int32 = int32(result >> 31) & 1

	set_carry(emu, int32(result >> 32))
	set_zero(emu, int32(result))
	set_sign(emu, signr)
	set_overflow(emu, sign1 != sign2 && sign1 != signr)
}

func set_carry(emu *Emulator, carry_value int32) {
	if (carry_value != 0) {
		emu.eflags |= CARRY_FLAG
		return
	}
	emu.eflags &= ^CARRY_FLAG
}

func set_zero(emu *Emulator, zero_value int32) {
	if (zero_value == 0) {
		emu.eflags |= ZERO_FLAG
		return
	}
	emu.eflags &= ^ZERO_FLAG
}

func set_sign(emu *Emulator, sign_value int32) {
	if (sign_value != 0) {
		emu.eflags |= SIGN_FLAG
		return
	}
	emu.eflags &= ^SIGN_FLAG
}

func is_carry(emu *Emulator) bool {
	return (emu.eflags & CARRY_FLAG) != 0
}

func is_zero(emu *Emulator) bool {
	return (emu.eflags & ZERO_FLAG) != 0
}

func is_sign(emu *Emulator) bool {
	return (emu.eflags & SIGN_FLAG) != 0
}

func is_overflow(emu *Emulator) bool {
	return (emu.eflags & OVERFLOW_FLAG) != 0
}

func set_overflow(emu *Emulator, is_overflow bool) {
	if is_overflow {
		emu.eflags |= OVERFLOW_FLAG
		return
	}
	emu.eflags &= ^OVERFLOW_FLAG
}
