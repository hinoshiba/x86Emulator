package main

import (
	"errors"
)

type ModRM struct {
	mod       uint8
	opecode   *uint8
	reg_index *uint8
	rm        uint8

	sib       uint8

	disp8     int8
	disp32    int32
}

func parse_modrm(emu *Emulator, modrm *ModRM) {
	var code uint8

	var box uint8
	modrm.opecode = &box
	modrm.reg_index = &box

	code = get_code8(emu, 0)
	modrm.mod = ((code & 0xC0) >> 6)
	buf_ope := ((code & 0x38) >> 3)
	*modrm.opecode = buf_ope
	modrm.rm = code & 0x07

	emu.eip += 1

	if ((modrm.mod == 0 && modrm.rm == 5) || modrm.mod == 2) {
		modrm.disp32 = get_sign_code32(emu, 0)
		emu.eip += 4
	} else if (modrm.mod == 1) {
		modrm.disp8 = get_sign_code8(emu, 0)
		emu.eip += 1
	}
}

func calc_memory_address(emu *Emulator, modrm *ModRM) (uint32, error) {
	switch (modrm.mod) {
	case 0:
		if (modrm.rm == 4) {
			return 0, errors.New("not implimented modRM mod = 0, rm = 4")
		} else if (modrm.rm == 5) {
			return uint32(modrm.disp32), nil
		} else {
			return get_register32(emu, modrm.rm), nil
		}
	case 1:
		if (modrm.rm == 4) {
			return 0, errors.New("not implimented modRM mod = 1, rm = 4")
		}
		return get_register32(emu, modrm.rm) + uint32(modrm.disp8), nil
	case 2:
		if (modrm.rm == 4) {
			return 0, errors.New("not implimented modRM mod = 2, rm = 4")
		}
		return get_register32(emu, modrm.rm) + uint32(modrm.disp32), nil
	default:
		return 0, errors.New("not implimented ModRM mod = 3")
	}
}

func set_rm32(emu *Emulator, modrm *ModRM, value uint32) error {
	if (modrm.mod == 3) {
		set_register32(emu, modrm.rm, value)
	} else {
		var address uint32
		address, err := calc_memory_address(emu, modrm)
		if err != nil {
			return err
		}
		set_memory32(emu, address, value)
	}
	return nil
}

func get_rm32(emu *Emulator, modrm *ModRM) (uint32, error) {
	if (modrm.mod == 3) {
		return get_register32(emu, modrm.rm), nil
	} else {
		var address uint32
		address, err := calc_memory_address(emu, modrm)
		if err != nil {
			return 0, err
		}
		return get_memory32(emu, address), nil
	}
	return 0, nil
}

func set_r32(emu *Emulator, modrm *ModRM, value uint32) {
	set_register32(emu, *modrm.reg_index, value)
}

func get_r32(emu *Emulator, modrm *ModRM) uint32 {
	return get_register32(emu, *modrm.reg_index)
}
