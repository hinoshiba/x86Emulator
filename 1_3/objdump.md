1.3 objdump
===========

* object file
```
$ gcc -c -g -o casm-c-sample.o casm-c-sample.c
```

* Disassemble in 64bit mode
```
$ objdump -d -S -M intel casm-c-sample.o

casm-c-sample.o:     ファイル形式 elf64-x86-64


セクション .text の逆アセンブル:

0000000000000000 <func>:
void func(void) {
   0:   55                      push   rbp
   1:   48 89 e5                mov    rbp,rsp
  int val = 0;
   4:   c7 45 fc 00 00 00 00    mov    DWORD PTR [rbp-0x4],0x0
  val++;
   b:   83 45 fc 01             add    DWORD PTR [rbp-0x4],0x1
}
   f:   90                      nop
  10:   5d                      pop    rbp
  11:   c3                      ret
```

* Disassemble in 32bit mode
```
$ objdump -d -S -M intel -m i386 casm-c-sample.o

casm-c-sample.o:     ファイル形式 elf64-x86-64


セクション .text の逆アセンブル:

0000000000000000 <func>:
void func(void) {
   0:   55                      push   ebp
   1:   48                      dec    eax
   2:   89 e5                   mov    ebp,esp
  int val = 0;
   4:   c7 45 fc 00 00 00 00    mov    DWORD PTR [ebp-0x4],0x0
  val++;
   b:   83 45 fc 01             add    DWORD PTR [ebp-0x4],0x1
}
   f:   90                      nop
  10:   5d                      pop    ebp
  11:   c3                      ret
```
