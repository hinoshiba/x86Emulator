3.4/README
==========

* [src/golang/emu3.4/main.go](src/golang/emu3.4/main.go)
* [src/sample/emu3.4/main.c](src/sample/emu3.4/main.c)

---

* result helloworld.bin
```
$ go run ./golang/emu3.4/*.go ./sample/pasm-helloworld/helloworld.bin
EIP = 7C00, Code = B8
EIP = 7C05, Code = E9


end of program.

EAX = 00000029
ECX = 00000000
EDX = 00000000
EBX = 00000000
ESP = 00007c00
EBP = 00000000
ESI = 00000000
EDI = 00000000
EIP = 00000000
```
* result modrm-test.bin
```
$ go run ./golang/emu3.4/*.go ./sample/exec-modrm-test/modrm-test.bin
EIP = 7C00, Code = 83
EIP = 7C03, Code = 89
EIP = 7C05, Code = B8
EIP = 7C0A, Code = C7
EIP = 7C11, Code = 01
EIP = 7C14, Code = 8B
EIP = 7C17, Code = FF
EIP = 7C1A, Code = 8B
EIP = 7C1D, Code = E9


end of program.

EAX = 00000002
ECX = 00000000
EDX = 00000000
EBX = 00000000
ESP = 00007bf0
EBP = 00007bf0
ESI = 00000007
EDI = 00000008
EIP = 00000000
```
