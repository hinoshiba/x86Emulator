package main

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

type Emulator struct {
	registers  [REGISTERS_COUNT]uint32
	eflags     uint32
	memory     []byte
	eip        uint32
}
