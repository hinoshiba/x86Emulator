3.10/README
==========

* [src/golang/emu3.10/main.go](src/golang/emu3.10/main.go)
* [src/sample/emu3.10/main.c](src/sample/emu3.10/main.c)

* result `exec-if-test/test.bin` (not jx)
```
$ go run src/golang/emu3.10/* src/sample/exec-if-test/test.bin
EIP = 7C00, Code = E8
EIP = 7C1F, Code = 55
EIP = 7C20, Code = 89
EIP = 7C22, Code = B8
EIP = 7C27, Code = 5D
EIP = 7C28, Code = C3
EIP = 7C05, Code = E9


end of program.

EAX = 00000003
ECX = 00000000
EDX = 00000000
EBX = 00000000
ESP = 00007c00
EBP = 00000000
ESI = 00000000
EDI = 00000000
EIP = 00000000
```

* result `exec-if-goto/test.bin` (found 7e)
```
$ go run src/golang/emu3.10/* src/sample/exec-if-goto/test.bin
EIP = 7C00, Code = E8
EIP = 7C31, Code = 55
EIP = 7C32, Code = 89
EIP = 7C34, Code = 6A
EIP = 7C36, Code = 6A
EIP = 7C38, Code = E8
EIP = 7C0A, Code = 55
EIP = 7C0B, Code = 89
EIP = 7C0D, Code = 83
EIP = 7C10, Code = C7
EIP = 7C17, Code = EB
EIP = 7C24, Code = 8B
EIP = 7C27, Code = 3B
EIP = 7C2A, Code = 7E
EIP = 7C19, Code = 90


Not Implemented: 90
EAX = 00000001
ECX = 00000000
EDX = 00000000
EBX = 00000000
ESP = 00007bd8
EBP = 00007be8
ESI = 00000000
EDI = 00000000
EIP = 00007c19
```

* result `exec-while-stmt/test.bin`
```
$ go run src/golang/emu3.10/* src/sample/exec-while-stmt/test.bin
EIP = 7C00, Code = E8
EIP = 7C30, Code = 55
EIP = 7C31, Code = 89
EIP = 7C33, Code = 6A
EIP = 7C35, Code = 6A
EIP = 7C37, Code = E8
EIP = 7C0A, Code = 55
EIP = 7C0B, Code = 89
EIP = 7C0D, Code = 83
EIP = 7C10, Code = C7
EIP = 7C17, Code = EB
EIP = 7C23, Code = 8B
EIP = 7C26, Code = 3B
EIP = 7C29, Code = 7E
EIP = 7C19, Code = 8B
EIP = 7C1C, Code = 01
EIP = 7C1F, Code = 83
EIP = 7C23, Code = 8B
EIP = 7C26, Code = 3B
EIP = 7C29, Code = 7E
EIP = 7C19, Code = 8B
EIP = 7C1C, Code = 01
EIP = 7C1F, Code = 83
EIP = 7C23, Code = 8B
EIP = 7C26, Code = 3B
EIP = 7C29, Code = 7E
EIP = 7C19, Code = 8B
EIP = 7C1C, Code = 01
EIP = 7C1F, Code = 83
EIP = 7C23, Code = 8B
EIP = 7C26, Code = 3B
EIP = 7C29, Code = 7E
EIP = 7C19, Code = 8B
EIP = 7C1C, Code = 01
EIP = 7C1F, Code = 83
EIP = 7C23, Code = 8B
EIP = 7C26, Code = 3B
EIP = 7C29, Code = 7E
EIP = 7C19, Code = 8B
EIP = 7C1C, Code = 01
EIP = 7C1F, Code = 83
EIP = 7C23, Code = 8B
EIP = 7C26, Code = 3B
EIP = 7C29, Code = 7E
EIP = 7C19, Code = 8B
EIP = 7C1C, Code = 01
EIP = 7C1F, Code = 83
EIP = 7C23, Code = 8B
EIP = 7C26, Code = 3B
EIP = 7C29, Code = 7E
EIP = 7C19, Code = 8B
EIP = 7C1C, Code = 01
EIP = 7C1F, Code = 83
EIP = 7C23, Code = 8B
EIP = 7C26, Code = 3B
EIP = 7C29, Code = 7E
EIP = 7C19, Code = 8B
EIP = 7C1C, Code = 01
EIP = 7C1F, Code = 83
EIP = 7C23, Code = 8B
EIP = 7C26, Code = 3B
EIP = 7C29, Code = 7E
EIP = 7C19, Code = 8B
EIP = 7C1C, Code = 01
EIP = 7C1F, Code = 83
EIP = 7C23, Code = 8B
EIP = 7C26, Code = 3B
EIP = 7C29, Code = 7E
EIP = 7C19, Code = 8B
EIP = 7C1C, Code = 01
EIP = 7C1F, Code = 83
EIP = 7C23, Code = 8B
EIP = 7C26, Code = 3B
EIP = 7C29, Code = 7E
EIP = 7C2B, Code = 8B
EIP = 7C2E, Code = C9
EIP = 7C2F, Code = C3
EIP = 7C3C, Code = 83
EIP = 7C3F, Code = C9
EIP = 7C40, Code = C3
EIP = 7C05, Code = E9


end of program.

EAX = 00000037
ECX = 00000000
EDX = 00000000
EBX = 00000000
ESP = 00007c00
EBP = 00000000
ESI = 00000000
EDI = 00000000
EIP = 00000000
```
