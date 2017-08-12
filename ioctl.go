package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

const XFSIOCFSCOUNTS = 0x80205871

type FsopCounts struct {
	FreeData uint64
	FreeRtX  uint64
	FreeIno  uint64
	AllocIno uint64
}

func decodeIoctl(regs syscall.PtraceRegs, pid int) {
	if uint32(regs.Rsi) != XFSIOCFSCOUNTS {
		// not our ioctl
		return
	}

	data := make([]byte, int(unsafe.Sizeof(FsopCounts{})))
	_, err := syscall.PtracePeekData(pid, uintptr(regs.Rdx), data)
	if err != nil {
		panic(err)
	}

	r := bytes.NewReader(data)
	var f FsopCounts
	err = binary.Read(r, binary.LittleEndian, &f)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stderr, "%+v\n", f)
}
