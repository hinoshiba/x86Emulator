3.7/README
==========

* [src/golang/emu3.7/main.go](src/golang/emu3.7/main.go)
* [src/sample/emu3.7/main.c](src/sample/emu3.7/main.c)

* result call-test.asm
```
$ go run golang/emu3.7/*  sample/exec-call-test/call-test.bin
EIP = 7C00, Code = B8
EIP = 7C05, Code = BB
EIP = 7C0A, Code = E8
EIP = 7C14, Code = 89
EIP = 7C16, Code = 01
EIP = 7C18, Code = C3
EIP = 7C0F, Code = E9


end of program.

EAX = 000000f1
ECX = 0000011a
EDX = 00000000
EBX = 00000029
ESP = 00007c00
EBP = 00000000
ESI = 00000000
EDI = 00000000
EIP = 00000000
```

* check call-c-test
```
$ objdump -b binary -m i8086 -D sample/exec-c-test/test.bin

sample/exec-c-test/test.bin:     ファイル形式 binary


セクション .data の逆アセンブル:

00000000 <.data>:
   0:   e8 05 00                call   0x8
   3:   00 00                   add    %al,(%bx,%si)
   5:   e9 f6 83                jmp    0x83fe
   8:   ff                      (bad)
   9:   ff 55 89                call   *-0x77(%di)
   c:   e5 83                   in     $0x83,%ax
   e:   ec                      in     (%dx),%al
   f:   10 c7                   adc    %al,%bh
  11:   45                      inc    %bp
  12:   fc                      cld
  13:   28 00                   sub    %al,(%bx,%si)
  15:   00 00                   add    %al,(%bx,%si)
  17:   83 45 fc 01             addw   $0x1,-0x4(%di)
  1b:   8b 45 fc                mov    -0x4(%di),%ax
  1e:   c9                      leave
  1f:   c3                      ret
```

* result call-c-test
```
$ go run golang/emu3.7/*  sample/exec-c-test/test.bin
EIP = 7C00, Code = E8
EIP = 7C0A, Code = 55
EIP = 7C0B, Code = 89
EIP = 7C0D, Code = 83
EIP = 7C10, Code = C7
EIP = 7C17, Code = 83
EIP = 7C1B, Code = 8B
EIP = 7C1E, Code = C9
EIP = 7C1F, Code = C3
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

* result call-args-test
```
$ go run golang/emu3.7/*  sample/exec-arg-test/test.bin
EIP = 7C00, Code = E8
EIP = 7C17, Code = 55
EIP = 7C18, Code = 89
EIP = 7C1A, Code = 6A
EIP = 7C1C, Code = 6A
EIP = 7C1E, Code = E8
EIP = 7C0A, Code = 55
EIP = 7C0B, Code = 89
EIP = 7C0D, Code = 8B
EIP = 7C10, Code = 8B
EIP = 7C13, Code = 01
EIP = 7C15, Code = 5D
EIP = 7C16, Code = C3
EIP = 7C23, Code = 83
EIP = 7C26, Code = C9
EIP = 7C27, Code = C3
EIP = 7C05, Code = E9


end of program.

EAX = 00000007
ECX = 00000000
EDX = 00000002
EBX = 00000000
ESP = 00007c00
EBP = 00000000
ESI = 00000000
EDI = 00000000
EIP = 00000000
```
