package main

import (
	"fmt"
	"errors"
)

type instruction_func_t func(*Emulator) error
var instructions = [256]instruction_func_t{}

func init_instructions() {
	instructions[0x01] = add_rm32_r32
	instructions[0x83] = code_83
	instructions[0x89] = mov_rm32_r32
	instructions[0x8B] = mov_r32_rm32
	for i := 0; i < 8; i++ {
		instructions[0xB8 + i] = mov_r32_imm32
	}
	instructions[0xC7] = mov_rm32_imm32
	instructions[0xE9] = near_jump
	instructions[0xEB] = short_jump
	instructions[0xFF] = code_ff
}

func mov_r32_imm32(emu *Emulator) error {
	var reg uint8 = get_code8(emu, 0) - 0xB8
	var value uint32 = get_code32(emu, 1)
	emu.registers[reg] = value
	emu.eip += 5
	return nil
}

func mov_rm32_imm32(emu *Emulator) error {
	emu.eip += 1
	var modrm ModRM
	parse_modrm(emu, &modrm)
	var value uint32 = get_code32(emu, 0)
	emu.eip += 4
	set_rm32(emu, &modrm, value)
	return nil
}

func mov_rm32_r32(emu *Emulator) error {
	emu.eip += 1
	var modrm ModRM
	parse_modrm(emu, &modrm)

	var r32 uint32 = get_r32(emu, &modrm)
	set_rm32(emu, &modrm, r32)
	return nil
}

func mov_r32_rm32(emu *Emulator) error {
	emu.eip += 1
	var modrm ModRM
	parse_modrm(emu, &modrm)

	var rm32 uint32
	rm32, err := get_rm32(emu, &modrm);
	if err != nil {
		return err
	}
	set_r32(emu, &modrm, rm32)
	return nil
}

func short_jump(emu *Emulator) error {
	var diff int8
	diff = get_sign_code8(emu, 1)
	emu.eip += uint32(diff + 2)
	return nil
}

func near_jump(emu *Emulator) error {
	var diff int32
	diff = get_sign_code32(emu, 1)
	emu.eip += uint32(diff + 5)
	return nil
}

func add_rm32_r32(emu *Emulator) error {
	emu.eip += 1
	var modrm ModRM
	parse_modrm(emu, &modrm)

	var r32 uint32 = get_r32(emu, &modrm)
	var rm32 uint32
	rm32, err := get_rm32(emu, &modrm)
	if err != nil {
		return err
	}

	if err := set_rm32(emu, &modrm, rm32 + r32); err != nil {
		return err
	}
	return nil
}

func sub_rm32_imm8(emu *Emulator, modrm *ModRM) error {
	var rm32 uint32
	rm32, err := get_rm32(emu, modrm)
	if err != nil {
		return err
	}
	var imm8 uint32 = uint32(get_sign_code8(emu, 0))

	emu.eip += 1
	if err := set_rm32(emu, modrm, rm32 - imm8); err != nil {
		return err
	}
	return nil
}

func code_83(emu *Emulator) error {
	emu.eip += 1
	var modrm ModRM
	parse_modrm(emu, &modrm)

	switch (*modrm.opecode) {
	case 5:
		if err := sub_rm32_imm8(emu, &modrm); err != nil {
			return err
		}
	default:
		return errors.New(fmt.Sprintf("not impliment: 83 %d", modrm.opecode))
	}
	return nil
}

func inc_rm32(emu *Emulator, modrm *ModRM) error {
	var value uint32
	value, err := get_rm32(emu, modrm)
	if err != nil {
		return err
	}
	if err := set_rm32(emu, modrm, value + 1); err != nil {
		return err
	}
	return nil
}

func code_ff(emu *Emulator) error {
	emu.eip += 1
	var modrm ModRM
	parse_modrm(emu, &modrm)

	switch (*modrm.opecode) {
	case 0:
		if err := inc_rm32(emu, &modrm); err != nil {
			return err
		}
	default:
		return errors.New(fmt.Sprintf("not impliment: FF %d", modrm.opecode))
	}
	return nil
}
