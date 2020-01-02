3.12/README
==========

* [src/golang/emu3.12/main.go](src/golang/emu3.12/main.go)
* [src/sample/emu3.12/main.c](src/sample/emu3.12/main.c)

---

* result in.bin
```
$ go run src/golang/emu3.12/* -q src/sample/exec-io-test/in.bin
a


end of program.

EAX = 00000061
ECX = 00000000
EDX = 000003f8
EBX = 00000000
ESP = 00007c00
EBP = 00000000
ESI = 00000000
EDI = 00000000
EIP = 00000000
```

* result out.bin
```
$ go run src/golang/emu3.12/* -q src/sample/exec-io-test/out.bin
A

end of program.

EAX = 00000041
ECX = 00000000
EDX = 000003f8
EBX = 00000000
ESP = 00007c00
EBP = 00000000
ESI = 00000000
EDI = 00000000
EIP = 00000000
```

* result select.bin
```
$ go run src/golang/emu3.12/* src/sample/exec-io-test/select.bin -q
>w
world
>h
hello
>q


end of program.

EAX = 00000071
ECX = 00000000
EDX = 000003f8
EBX = 00000000
ESP = 00007c00
EBP = 00000000
ESI = 00007c47
EDI = 00000000
EIP = 00000000
```
