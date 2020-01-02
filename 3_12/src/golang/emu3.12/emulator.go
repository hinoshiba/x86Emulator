package main

const (
	EAX uint8 = iota
	ECX
	EDX
	EBX
	ESP
	EBP
	ESI
	EDI
	REGISTERS_COUNT
	AL = EAX
	CL = ECX
	DL = EDX
	BL = EBX
	AH = AL + 4
	CH = CL + 4
	DH = DL + 4
	BH = BL + 4
)

type Emulator struct {
	registers  [REGISTERS_COUNT]uint32
	eflags     uint32
	memory     []byte
	eip        uint32
}
