package main

import (
	"fmt"
)

type instruction_func_t func(*Emulator) error
var instructions = [256]instruction_func_t{}

func init_instructions() {
	instructions[0x01] = add_rm32_r32

    instructions[0x3B] = cmp_r32_rm32
    instructions[0x3C] = cmp_al_imm8
    instructions[0x3D] = cmp_eax_imm32

    for i := 0; i < 8; i++ {
        instructions[0x40 + i] = inc_r32
    }
	for i := 0; i < 8; i++ {
		instructions[0x50 + i] = push_r32
	}
	for i := 0; i < 8; i++ {
		instructions[0x58 + i] = pop_r32
	}

    instructions[0x68] = push_imm32
    instructions[0x6A] = push_imm8

    instructions[0x70] = jo
    instructions[0x71] = jno
    instructions[0x72] = jc
    instructions[0x73] = jnc
    instructions[0x74] = jz
    instructions[0x75] = jnz
    instructions[0x78] = js
    instructions[0x79] = jns
    instructions[0x7C] = jl
    instructions[0x7E] = jle

	instructions[0x83] = code_83
    instructions[0x8A] = mov_r8_rm8
    instructions[0x8B] = mov_r32_rm32
	instructions[0x89] = mov_rm32_r32
	instructions[0x8B] = mov_r32_rm32

    for i := 0; i < 8; i++ {
        instructions[0xB0 + i] = mov_r8_imm8
    }
	for i := 0; i < 8; i++ {
		instructions[0xB8 + i] = mov_r32_imm32
	}
	instructions[0xC3] = ret
	instructions[0xC7] = mov_rm32_imm32
	instructions[0xC9] = leave
    instructions[0xCD] = swi
	instructions[0xE8] = call_rel32
	instructions[0xE9] = near_jump
	instructions[0xEB] = short_jump
    instructions[0xEC] = in_al_dx
    instructions[0xEE] = out_dx_al
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

func mov_r8_rm8(emu *Emulator) error {
	emu.eip += 1
	var modrm ModRM
	parse_modrm(emu, &modrm)

	var rm8 uint8
	rm8, err := get_rm8(emu, &modrm)
	if err != nil {
		return err
	}

	set_r8(emu, &modrm, rm8)
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

func mov_r8_imm8(emu *Emulator) error {
	var reg uint8 = get_code8(emu, 0) - 0xB0
	set_register8(emu, reg, get_code8(emu, 1))

	emu.eip += 2
	return nil
}

func mov_rm8_r8(emu *Emulator) error {
	emu.eip += 1

	var modrm ModRM
	parse_modrm(emu, &modrm)

	var r8 uint8 = get_r8(emu, &modrm)
	set_rm8(emu, &modrm, r8)

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
	var result uint64 = uint64(rm32) - uint64(imm8)

	emu.eip += 1
	if err := set_rm32(emu, modrm, uint32(result)); err != nil {
		return err
	}
	update_eflags_sub(emu, rm32, imm8, result)
	return nil
}

func cmp_r32_rm32(emu *Emulator) error {
	emu.eip += 1

	var modrm ModRM
	parse_modrm(emu, &modrm)

	var r32 uint32
	var rm32 uint32
	var err error

	r32 = get_r32(emu, &modrm)
	rm32, err = get_rm32(emu, &modrm)
	if err != nil {
		return err
	}
	var result uint64 = uint64(r32) - uint64(rm32)

	update_eflags_sub(emu, r32, rm32, result)
	return nil
}

func cmp_rm32_imm8(emu *Emulator, modrm *ModRM) error {
	var rm32 uint32
	rm32, err := get_rm32(emu, modrm)
	if err != nil {
		return err
	}
	var imm8 uint32 = uint32(get_sign_code8(emu, 0))

	emu.eip += 1

	var result uint64 = uint64(rm32) - uint64(imm8)
	update_eflags_sub(emu, rm32, imm8, result)
	return nil
}

func cmp_al_imm8(emu *Emulator) error {
	var value uint8 = get_code8(emu, 1)
	var al uint8 = get_register8(emu, AL)
	var result uint64 = uint64(al) - uint64(value)

	update_eflags_sub(emu, uint32(al), uint32(value), result)
	emu.eip += 2
	return nil
}

func cmp_eax_imm32(emu *Emulator) error {
	var value uint32 = get_code32(emu, 1)
	var eax uint32 = get_register32(emu, uint8(EAX))
	var result uint64 = uint64(eax) - uint64(value)

	update_eflags_sub(emu, uint32(eax), uint32(value), result)
	emu.eip += 5
	return nil
}

func code_83(emu *Emulator) error {
	emu.eip += 1
	var modrm ModRM
	parse_modrm(emu, &modrm)

	switch (*modrm.opecode) {
	case 0:
		if err := add_rm32_imm8(emu, &modrm); err != nil {
			return err
		}
	case 5:
		if err := sub_rm32_imm8(emu, &modrm); err != nil {
			return err
		}
	case 7:
		if err := cmp_rm32_imm8(emu, &modrm); err != nil {
			return err
		}
	default:
		return fmt.Errorf("undefined modrm opecode: 83 %d", modrm.opecode)
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

func inc_r32(emu *Emulator) error {
	var reg uint8 = get_code8(emu, 0) - 0x40;
	set_register32(emu, reg, get_register32(emu, reg) + 1)

	emu.eip += 1
	return nil
}

func push_r32(emu *Emulator) error {
	var reg uint8 = get_code8(emu, 0) - 0x50
	push32(emu, get_register32(emu, reg))
	emu.eip += 1
	return nil
}

func push_imm32(emu *Emulator) error {
	var value uint32 = get_code32(emu, 1)
	push32(emu, value)
	emu.eip += 5
	return nil
}

func push_imm8(emu *Emulator) error {
	var value uint8 = get_code8(emu, 1)
	push32(emu, uint32(value))
	emu.eip += 2
	return nil
}

func add_rm32_imm8(emu *Emulator, modrm *ModRM) error {
	var rm32 uint32
	rm32, err := get_rm32(emu, modrm)
	if err != nil {
		return err
	}

	var imm8 uint32 = uint32(get_sign_code8(emu, 0))
	emu.eip += 1
	set_rm32(emu, modrm, rm32 + imm8)

	return nil
}

func pop_r32(emu *Emulator) error {
	var reg uint8 = get_code8(emu, 0) - 0x58
	set_register32(emu, reg, pop32(emu))
	emu.eip += 1
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
		return fmt.Errorf("not impliment: FF %d", modrm.opecode)
	}
	return nil
}

func call_rel32(emu *Emulator) error {
	var diff int32 = get_sign_code32(emu, 1)
	push32(emu, emu.eip + 5)
	emu.eip = uint32(int32(emu.eip) + diff + 5)
	return nil
}

func ret(emu *Emulator) error {
	emu.eip = pop32(emu)
	return nil
}

func leave(emu *Emulator) error {
	var ebp uint32 = get_register32(emu, uint8(EBP))
	set_register32(emu, uint8(ESP), ebp)
	set_register32(emu, uint8(EBP), pop32(emu))
	emu.eip += 1
	return nil
}

func js(emu *Emulator) error {
	var diff int32 = 0
	if is_sign(emu) {
		diff = int32(get_sign_code8(emu, 1))
	}

	emu.eip += uint32(diff + 2)
	return nil
}

func jns(emu *Emulator) error {
	var diff int32 = int32(get_sign_code8(emu, 1))
	if is_sign(emu) {
		diff = 0
	}

	emu.eip += uint32(diff + 2)
	return nil
}

func jc(emu *Emulator) error {
	var diff int32 = 0
	if is_carry(emu) {
		diff = int32(get_sign_code8(emu, 1))
	}

	emu.eip += uint32(diff + 2)
	return nil
}

func jnc(emu *Emulator) error {
	var diff int32 = int32(get_sign_code8(emu, 1))
	if is_carry(emu) {
		diff = 0
	}

	emu.eip += uint32(diff + 2)
	return nil
}

func jz(emu *Emulator) error {
	var diff int32 = 0
	if is_zero(emu) {
		diff = int32(get_sign_code8(emu, 1))
	}

	emu.eip += uint32(diff + 2)
	return nil
}

func jnz(emu *Emulator) error {
	var diff int32 = int32(get_sign_code8(emu, 1))
	if is_zero(emu) {
		diff = 0
	}

	emu.eip += uint32(diff + 2)
	return nil
}

func jo(emu *Emulator) error {
	var diff int32 = 0
	if is_overflow(emu) {
		diff = int32(get_sign_code8(emu, 1))
	}

	emu.eip += uint32(diff + 2)
	return nil
}

func jno(emu *Emulator) error {
	var diff int32 = int32(get_sign_code8(emu, 1))
	if is_overflow(emu) {
		diff = 0
	}

	emu.eip += uint32(diff + 2)
	return nil
}

func jl(emu *Emulator) error {
	var diff int32 = 0
	if is_sign(emu) != is_overflow(emu) {
		diff = int32(get_sign_code8(emu, 1))
	}

	emu.eip += uint32(diff + 2)
	return nil
}

func jle(emu *Emulator) error {
	var diff int32 = 0
	if is_zero(emu) || is_sign(emu) != is_overflow(emu) {
		diff = int32(get_sign_code8(emu, 1))
	}

	emu.eip += uint32(diff + 2)
	return nil
}

func in_al_dx(emu *Emulator) error {
	var address uint16 = uint16(get_register32(emu, uint8(EDX)) & 0xffff)
	var value uint8 = io_in8(address)

	set_register8(emu, AL, value)
	emu.eip += 1
	return nil
}

func out_dx_al(emu *Emulator) error {
	var address uint16 = uint16(get_register32(emu, uint8(EDX)) & 0xffff)
	var value uint8 = get_register8(emu, AL)

	io_out8(address, value)
	emu.eip += 1
	return nil
}

func swi(emu *Emulator) error {
	var int_index uint8 = get_code8(emu, 1)
	emu.eip += 2

	switch int_index {
	case 0x10 :
		if err := bios_video(emu); err != nil {
			return err
		}
	default :
		return fmt.Errorf("unkwon interruput : 0x%02x\n", int_index)
	}
	return nil
}
