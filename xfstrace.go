package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"syscall"
)

// boilerplate main() from
// https://github.com/lizrice/strace-from-scratch
// with minor modifications and a call to decodeIoctl() added

func main() {
	runtime.LockOSThread()

	fmt.Printf("Run %v\n", os.Args[1:])

	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Ptrace: true,
	}

	cmd.Start()
	err := cmd.Wait()
	if err != nil {
		fmt.Printf("Wait returned: %v\n", err)
	}

	pid := cmd.Process.Pid
	exit := true

	var regs syscall.PtraceRegs

	for {
		if exit {
			err = syscall.PtraceGetRegs(pid, &regs)
			if err != nil {
				break
			}

			if regs.Orig_rax == syscall.SYS_IOCTL {
				decodeIoctl(regs, pid)
			}
		}

		err = syscall.PtraceSyscall(pid, 0)
		if err != nil {
			panic(err)
		}

		_, err = syscall.Wait4(pid, nil, 0, nil)
		if err != nil {
			panic(err)
		}

		exit = !exit
	}
}
