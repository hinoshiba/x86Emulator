4.2/README
==========

* [src/golang/emu4.2/main.go](src/golang/emu4.2/main.go)
* [src/sample/emu4.2/main.c](src/sample/emu4.2/main.c)

---

* result subroutine32.bin
```
$ go run src/golang/emu4.2/* src/sample/bios-subroutine32/subroutine32.bin -q
hello, world


end of program.

EAX = 00000e00
ECX = 00000000
EDX = 00000000
EBX = 0000000a
ESP = 00007c00
EBP = 00000000
ESI = 00007c31
EDI = 00000000
EIP = 00000000
```
