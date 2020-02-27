# Wrapper package for [`golang.org/x/sys/unix`](https://pkg.go.dev/golang.org/x/sys/unix) with automatic `EINTR` handler

Package `github.com/AkihiroSuda/x-sys-unix-auto-eintr` contains wrapper functions for [`golang.org/x/sys/unix`](https://pkg.go.dev/golang.org/x/sys/unix) with automatic `EINTR` handler as follows:

```go
// Read is an alias of golang.org/x/sys/unix.Read, wrapped to automatically retry on EINTR.
func Read(fd int, p []byte) (n int, err error) {
        for {
                n, err = unix.Read(fd, p)
                if err != syscall.EINTR {
                        break
                }
        }
        return
}
...
// Write is an alias of golang.org/x/sys/unix.Write, wrapped to automatically retry on EINTR.
func Write(fd int, p []byte) (n int, err error) {
        for {
                n, err = unix.Write(fd, p)
                if err != syscall.EINTR {
                        break
                }
        }
        return
}
...
// Close is an alias of golang.org/x/sys/unix.Close. Does NOT retry on EINTR. See close(2) for the reason.
func Close(fd int) (err error) {
        return unix.Close(fd)
}
...
// Syscall is an alias of golang.org/x/sys/unix.Syscall, wrapped to automatically retry on EINTR.
func Syscall(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err syscall.Errno) {
        for {
                r1, r2, err = unix.Syscall(trap, a1, a2, a3)
                if err != syscall.EINTR {
                        break
                }
        }
        return
}

// Syscall6 is an alias of golang.org/x/sys/unix.Syscall6, wrapped to automatically retry on EINTR.
func Syscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err syscall.Errno) {
        for {
                r1, r2, err = unix.Syscall6(trap, a1, a2, a3, a4, a5, a6)
                if err != syscall.EINTR {
                        break
                }
        }
        return
}

// RawSyscall is an alias of golang.org/x/sys/unix.RawSyscall, wrapped to automatically retry on EINTR.
func RawSyscall(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err syscall.Errno) {
        for {
                r1, r2, err = unix.RawSyscall(trap, a1, a2, a3)
                if err != syscall.EINTR {
                        break
                }
        }
        return
}

// RawSyscall6 is an alias of golang.org/x/sys/unix.RawSyscall6, wrapped to automatically retry on EINTR.
func RawSyscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err syscall.Errno) {
        for {
                r1, r2, err = unix.RawSyscall6(trap, a1, a2, a3, a4, a5, a6)
                if err != syscall.EINTR {
                        break
                }
        }
        return
}
```

## Background

https://golang.org/doc/go1.14#runtime

> A consequence of the implementation of preemption is that on Unix systems,
> including Linux and macOS systems, programs built with Go 1.14 will receive
> more signals than programs built with earlier releases. This means that
> programs that use packages like `syscall` or `golang.org/x/sys/unix` will see more
> slow system calls fail with EINTR errors. Those programs will have to handle
> those errors in some way, most likely looping to try the system call again. 

## Usage

Just substitute `import "golang.org/x/sys/unix"` in your project with `import unix "github.com/AkihiroSuda/x-sys-unix-auto-eintr"`.

## GoDoc

GoDoc is [here](https://pkg.go.dev/github.com/AkihiroSuda/x-sys-unix-auto-eintr), but you don't need to read this because the usage is same as [`golang.org/x/sys/unix`](https://pkg.go.dev/golang.org/x/sys/unix).

## Contributing
Pull requests are welcome.

Note that `generated_*.go` must not be modified manually.
Instead modify the code generator (`_codegen`) and run `make.sh` to generate the codes.
