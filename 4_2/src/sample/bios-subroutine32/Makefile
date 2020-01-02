TARGET = subroutine32.bin
Z_TOOLS = ../../../dls/tolset_p86/z_tools

CC = gcc
LD = ld
AS = nasm
CFLAGS += -m32 -nostdlib -fno-asynchronous-unwind-tables \
	-I$(Z_TOOLS)/i386-elf-gcc/include -g -fno-stack-protector
LDFLAGS += -m elf_i386 --entry=func --oformat=binary

.PHONY: all
all :
	make $(TARGET)

%.o : %.c Makefile
	$(CC) $(CFLAGS) -c $<

%.bin : %.o Makefile
	$(LD) $(LDFLAGS) -o $@ $<

%.bin : %.asm Makefile
	$(AS) -f bin -o $@ $<
