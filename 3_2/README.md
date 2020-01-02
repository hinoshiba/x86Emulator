2uu/README
==========

* [src/golang/emu3.2/main.go](src/golang/emu3.2/main.go)
* [src/sample/emu3.2/main.c](src/sample/emu3.2/main.c)


---

* result
```
$ go run ./golang/emu3.2/main.go ./sample/pasm-helloworld/helloworld.bin
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
