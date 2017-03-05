package main

import (
	"os"
	"syscall"
	"fmt"
	"strconv"
	"runtime"
)

func main() {
	pid, _ := strconv.Atoi(os.Args[1])

    /* Lock thread execution for ptrace calls */
	runtime.LockOSThread()

	var wstatus syscall.WaitStatus
	var rusage syscall.Rusage
    var in_syscall bool = false
    var registers syscall.PtraceRegs

    /* Attach to process */
	syscall.PtraceAttach(pid)
	_, err := syscall.Wait4(pid, &wstatus, 0, &rusage)
    if err != nil {
        panic(err)
    }

    /* Set option for syscall tracing return code */
	syscall.PtraceSetOptions(pid, syscall.PTRACE_O_TRACESYSGOOD)

    /* Tracing loop */
    for true {
		syscall.PtraceSyscall(pid, 0)

        _, err = syscall.Wait4(pid, &wstatus, 0, &rusage)
        if err != nil {
            panic(err)
        }
        if wstatus.StopSignal() != 0x85 {
            continue
        }

        /* Get registers at current state */
        err = syscall.PtraceGetRegs(pid, &registers)
        if  err != nil {
            panic(err)
        }

        if in_syscall {
            fmt.Printf("exit_code=%d\n", registers.Rax)
            in_syscall = false
        } else {
            fmt.Printf("SYSCALL #%d ", registers.Orig_rax)
            in_syscall = true
        }
	}
}
