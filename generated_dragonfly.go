// Automatically generated file. DO NOT EDIT MANUALLY.

package unix

import (
	"syscall"
	"time"

	unix "golang.org/x/sys/unix"
)

// Major is an alias of golang.org/x/sys/unix.Major.
func Major(dev uint64) uint32 {
	return unix.Major(dev)
}

// Minor is an alias of golang.org/x/sys/unix.Minor.
func Minor(dev uint64) uint32 {
	return unix.Minor(dev)
}

// Mkdev is an alias of golang.org/x/sys/unix.Mkdev.
func Mkdev(major, minor uint32) uint64 {
	return unix.Mkdev(major, minor)
}

// ParseDirent is an alias of golang.org/x/sys/unix.ParseDirent.
func ParseDirent(buf []byte, max int, names []string) (consumed int, count int, newnames []string) {
	return unix.ParseDirent(buf, max, names)
}

// Getenv is an alias of golang.org/x/sys/unix.Getenv.
func Getenv(key string) (value string, found bool) {
	return unix.Getenv(key)
}

// Setenv is an alias of golang.org/x/sys/unix.Setenv, wrapped to automatically retry on EINTR.
func Setenv(key, value string) error {
	var (
		_v0 error
	)
	for {
		_v0 = unix.Setenv(key, value)
		if _v0 != syscall.EINTR {
			break
		}
	}
	return _v0
}

// Clearenv is an alias of golang.org/x/sys/unix.Clearenv.
func Clearenv() {
	unix.Clearenv()
}

// Environ is an alias of golang.org/x/sys/unix.Environ.
func Environ() []string {
	return unix.Environ()
}

// Unsetenv is an alias of golang.org/x/sys/unix.Unsetenv, wrapped to automatically retry on EINTR.
func Unsetenv(key string) error {
	var (
		_v0 error
	)
	for {
		_v0 = unix.Unsetenv(key)
		if _v0 != syscall.EINTR {
			break
		}
	}
	return _v0
}

// FcntlInt is an alias of golang.org/x/sys/unix.FcntlInt, wrapped to automatically retry on EINTR.
func FcntlInt(fd uintptr, cmd, arg int) (int, error) {
	var (
		_v0 int
	)
	var (
		_v1 error
	)
	for {
		_v0, _v1 = unix.FcntlInt(fd, cmd, arg)
		if _v1 != syscall.EINTR {
			break
		}
	}
	return _v0, _v1
}

// FcntlFlock is an alias of golang.org/x/sys/unix.FcntlFlock, wrapped to automatically retry on EINTR.
func FcntlFlock(fd uintptr, cmd int, lk *unix.Flock_t) error {
	var (
		_v0 error
	)
	for {
		_v0 = unix.FcntlFlock(fd, cmd, lk)
		if _v0 != syscall.EINTR {
			break
		}
	}
	return _v0
}

// IoctlSetInt is an alias of golang.org/x/sys/unix.IoctlSetInt, wrapped to automatically retry on EINTR.
func IoctlSetInt(fd int, req uint, value int) error {
	var (
		_v0 error
	)
	for {
		_v0 = unix.IoctlSetInt(fd, req, value)
		if _v0 != syscall.EINTR {
			break
		}
	}
	return _v0
}

// IoctlSetWinsize is an alias of golang.org/x/sys/unix.IoctlSetWinsize, wrapped to automatically retry on EINTR.
func IoctlSetWinsize(fd int, req uint, value *unix.Winsize) error {
	var (
		_v0 error
	)
	for {
		_v0 = unix.IoctlSetWinsize(fd, req, value)
		if _v0 != syscall.EINTR {
			break
		}
	}
	return _v0
}

// IoctlSetTermios is an alias of golang.org/x/sys/unix.IoctlSetTermios, wrapped to automatically retry on EINTR.
func IoctlSetTermios(fd int, req uint, value *unix.Termios) error {
	var (
		_v0 error
	)
	for {
		_v0 = unix.IoctlSetTermios(fd, req, value)
		if _v0 != syscall.EINTR {
			break
		}
	}
	return _v0
}

// IoctlGetInt is an alias of golang.org/x/sys/unix.IoctlGetInt, wrapped to automatically retry on EINTR.
func IoctlGetInt(fd int, req uint) (int, error) {
	var (
		_v0 int
	)
	var (
		_v1 error
	)
	for {
		_v0, _v1 = unix.IoctlGetInt(fd, req)
		if _v1 != syscall.EINTR {
			break
		}
	}
	return _v0, _v1
}

// IoctlGetWinsize is an alias of golang.org/x/sys/unix.IoctlGetWinsize, wrapped to automatically retry on EINTR.
func IoctlGetWinsize(fd int, req uint) (*unix.Winsize, error) {
	var (
		_v1 error
	)
	var (
		_v0 *unix.Winsize
	)
	for {
		_v0, _v1 = unix.IoctlGetWinsize(fd, req)
		if _v1 != syscall.EINTR {
			break
		}
	}
	return _v0, _v1
}

// IoctlGetTermios is an alias of golang.org/x/sys/unix.IoctlGetTermios, wrapped to automatically retry on EINTR.
func IoctlGetTermios(fd int, req uint) (*unix.Termios, error) {
	var (
		_v0 *unix.Termios
	)
	var (
		_v1 error
	)
	for {
		_v0, _v1 = unix.IoctlGetTermios(fd, req)
		if _v1 != syscall.EINTR {
			break
		}
	}
	return _v0, _v1
}

// Getpagesize is an alias of golang.org/x/sys/unix.Getpagesize.
func Getpagesize() int {
	return unix.Getpagesize()
}

// ReadDirent is an alias of golang.org/x/sys/unix.ReadDirent, wrapped to automatically retry on EINTR.
func ReadDirent(fd int, buf []byte) (n int, err error) {
	for {
		n, err = unix.ReadDirent(fd, buf)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// CmsgLen is an alias of golang.org/x/sys/unix.CmsgLen.
func CmsgLen(datalen int) int {
	return unix.CmsgLen(datalen)
}

// CmsgSpace is an alias of golang.org/x/sys/unix.CmsgSpace.
func CmsgSpace(datalen int) int {
	return unix.CmsgSpace(datalen)
}

// ParseSocketControlMessage is an alias of golang.org/x/sys/unix.ParseSocketControlMessage, wrapped to automatically retry on EINTR.
func ParseSocketControlMessage(b []byte) ([]unix.SocketControlMessage, error) {
	var (
		_v0 []unix.SocketControlMessage
	)
	var (
		_v1 error
	)
	for {
		_v0, _v1 = unix.ParseSocketControlMessage(b)
		if _v1 != syscall.EINTR {
			break
		}
	}
	return _v0, _v1
}

// UnixRights is an alias of golang.org/x/sys/unix.UnixRights.
func UnixRights(fds int) []byte {
	return unix.UnixRights(fds)
}

// ParseUnixRights is an alias of golang.org/x/sys/unix.ParseUnixRights, wrapped to automatically retry on EINTR.
func ParseUnixRights(m *unix.SocketControlMessage) ([]int, error) {
	var (
		_v0 []int
	)
	var (
		_v1 error
	)
	for {
		_v0, _v1 = unix.ParseUnixRights(m)
		if _v1 != syscall.EINTR {
			break
		}
	}
	return _v0, _v1
}

// ByteSliceFromString is an alias of golang.org/x/sys/unix.ByteSliceFromString, wrapped to automatically retry on EINTR.
func ByteSliceFromString(s string) ([]byte, error) {
	var (
		_v0 []byte
	)
	var (
		_v1 error
	)
	for {
		_v0, _v1 = unix.ByteSliceFromString(s)
		if _v1 != syscall.EINTR {
			break
		}
	}
	return _v0, _v1
}

// BytePtrFromString is an alias of golang.org/x/sys/unix.BytePtrFromString, wrapped to automatically retry on EINTR.
func BytePtrFromString(s string) (*byte, error) {
	var (
		_v0 *byte
	)
	var (
		_v1 error
	)
	for {
		_v0, _v1 = unix.BytePtrFromString(s)
		if _v1 != syscall.EINTR {
			break
		}
	}
	return _v0, _v1
}

// Getgroups is an alias of golang.org/x/sys/unix.Getgroups, wrapped to automatically retry on EINTR.
func Getgroups() (gids []int, err error) {
	for {
		gids, err = unix.Getgroups()
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Setgroups is an alias of golang.org/x/sys/unix.Setgroups, wrapped to automatically retry on EINTR.
func Setgroups(gids []int) (err error) {
	for {
		err = unix.Setgroups(gids)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Wait4 is an alias of golang.org/x/sys/unix.Wait4, wrapped to automatically retry on EINTR.
func Wait4(pid int, wstatus *unix.WaitStatus, options int, rusage *unix.Rusage) (wpid int, err error) {
	for {
		wpid, err = unix.Wait4(pid, wstatus, options, rusage)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Accept is an alias of golang.org/x/sys/unix.Accept, wrapped to automatically retry on EINTR.
func Accept(fd int) (nfd int, sa unix.Sockaddr, err error) {
	for {
		nfd, sa, err = unix.Accept(fd)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Getsockname is an alias of golang.org/x/sys/unix.Getsockname, wrapped to automatically retry on EINTR.
func Getsockname(fd int) (sa unix.Sockaddr, err error) {
	for {
		sa, err = unix.Getsockname(fd)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// GetsockoptString is an alias of golang.org/x/sys/unix.GetsockoptString, wrapped to automatically retry on EINTR.
func GetsockoptString(fd, level, opt int) (string, error) {
	var (
		_v0 string
	)
	var (
		_v1 error
	)
	for {
		_v0, _v1 = unix.GetsockoptString(fd, level, opt)
		if _v1 != syscall.EINTR {
			break
		}
	}
	return _v0, _v1
}

// Recvmsg is an alias of golang.org/x/sys/unix.Recvmsg, wrapped to automatically retry on EINTR.
func Recvmsg(fd int, p, oob []byte, flags int) (n, oobn int, recvflags int, from unix.Sockaddr, err error) {
	for {
		n, oobn, recvflags, from, err = unix.Recvmsg(fd, p, oob, flags)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Sendmsg is an alias of golang.org/x/sys/unix.Sendmsg, wrapped to automatically retry on EINTR.
func Sendmsg(fd int, p, oob []byte, to unix.Sockaddr, flags int) (err error) {
	for {
		err = unix.Sendmsg(fd, p, oob, to, flags)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// SendmsgN is an alias of golang.org/x/sys/unix.SendmsgN, wrapped to automatically retry on EINTR.
func SendmsgN(fd int, p, oob []byte, to unix.Sockaddr, flags int) (n int, err error) {
	for {
		n, err = unix.SendmsgN(fd, p, oob, to, flags)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Kevent is an alias of golang.org/x/sys/unix.Kevent, wrapped to automatically retry on EINTR.
func Kevent(kq int, changes, events []unix.Kevent_t, timeout *unix.Timespec) (n int, err error) {
	for {
		n, err = unix.Kevent(kq, changes, events, timeout)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Sysctl is an alias of golang.org/x/sys/unix.Sysctl, wrapped to automatically retry on EINTR.
func Sysctl(name string) (string, error) {
	var (
		_v0 string
	)
	var (
		_v1 error
	)
	for {
		_v0, _v1 = unix.Sysctl(name)
		if _v1 != syscall.EINTR {
			break
		}
	}
	return _v0, _v1
}

// SysctlArgs is an alias of golang.org/x/sys/unix.SysctlArgs, wrapped to automatically retry on EINTR.
func SysctlArgs(name string, args int) (string, error) {
	var (
		_v0 string
	)
	var (
		_v1 error
	)
	for {
		_v0, _v1 = unix.SysctlArgs(name, args)
		if _v1 != syscall.EINTR {
			break
		}
	}
	return _v0, _v1
}

// SysctlUint32 is an alias of golang.org/x/sys/unix.SysctlUint32, wrapped to automatically retry on EINTR.
func SysctlUint32(name string) (uint32, error) {
	var (
		_v0 uint32
	)
	var (
		_v1 error
	)
	for {
		_v0, _v1 = unix.SysctlUint32(name)
		if _v1 != syscall.EINTR {
			break
		}
	}
	return _v0, _v1
}

// SysctlUint32Args is an alias of golang.org/x/sys/unix.SysctlUint32Args, wrapped to automatically retry on EINTR.
func SysctlUint32Args(name string, args int) (uint32, error) {
	var (
		_v0 uint32
	)
	var (
		_v1 error
	)
	for {
		_v0, _v1 = unix.SysctlUint32Args(name, args)
		if _v1 != syscall.EINTR {
			break
		}
	}
	return _v0, _v1
}

// SysctlUint64 is an alias of golang.org/x/sys/unix.SysctlUint64, wrapped to automatically retry on EINTR.
func SysctlUint64(name string, args int) (uint64, error) {
	var (
		_v0 uint64
	)
	var (
		_v1 error
	)
	for {
		_v0, _v1 = unix.SysctlUint64(name, args)
		if _v1 != syscall.EINTR {
			break
		}
	}
	return _v0, _v1
}

// SysctlRaw is an alias of golang.org/x/sys/unix.SysctlRaw, wrapped to automatically retry on EINTR.
func SysctlRaw(name string, args int) ([]byte, error) {
	var (
		_v0 []byte
	)
	var (
		_v1 error
	)
	for {
		_v0, _v1 = unix.SysctlRaw(name, args)
		if _v1 != syscall.EINTR {
			break
		}
	}
	return _v0, _v1
}

// Utimes is an alias of golang.org/x/sys/unix.Utimes, wrapped to automatically retry on EINTR.
func Utimes(path string, tv []unix.Timeval) error {
	var (
		_v0 error
	)
	for {
		_v0 = unix.Utimes(path, tv)
		if _v0 != syscall.EINTR {
			break
		}
	}
	return _v0
}

// UtimesNano is an alias of golang.org/x/sys/unix.UtimesNano, wrapped to automatically retry on EINTR.
func UtimesNano(path string, ts []unix.Timespec) error {
	var (
		_v0 error
	)
	for {
		_v0 = unix.UtimesNano(path, ts)
		if _v0 != syscall.EINTR {
			break
		}
	}
	return _v0
}

// UtimesNanoAt is an alias of golang.org/x/sys/unix.UtimesNanoAt, wrapped to automatically retry on EINTR.
func UtimesNanoAt(dirfd int, path string, ts []unix.Timespec, flags int) error {
	var (
		_v0 error
	)
	for {
		_v0 = unix.UtimesNanoAt(dirfd, path, ts, flags)
		if _v0 != syscall.EINTR {
			break
		}
	}
	return _v0
}

// Futimes is an alias of golang.org/x/sys/unix.Futimes, wrapped to automatically retry on EINTR.
func Futimes(fd int, tv []unix.Timeval) error {
	var (
		_v0 error
	)
	for {
		_v0 = unix.Futimes(fd, tv)
		if _v0 != syscall.EINTR {
			break
		}
	}
	return _v0
}

// Poll is an alias of golang.org/x/sys/unix.Poll, wrapped to automatically retry on EINTR.
func Poll(fds []unix.PollFd, timeout int) (n int, err error) {
	for {
		n, err = unix.Poll(fds, timeout)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Mmap is an alias of golang.org/x/sys/unix.Mmap, wrapped to automatically retry on EINTR.
func Mmap(fd int, offset int64, length int, prot int, flags int) (data []byte, err error) {
	for {
		data, err = unix.Mmap(fd, offset, length, prot, flags)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Munmap is an alias of golang.org/x/sys/unix.Munmap, wrapped to automatically retry on EINTR.
func Munmap(b []byte) (err error) {
	for {
		err = unix.Munmap(b)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Pipe is an alias of golang.org/x/sys/unix.Pipe, wrapped to automatically retry on EINTR.
func Pipe(p []int) (err error) {
	for {
		err = unix.Pipe(p)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Pread is an alias of golang.org/x/sys/unix.Pread, wrapped to automatically retry on EINTR.
func Pread(fd int, p []byte, offset int64) (n int, err error) {
	for {
		n, err = unix.Pread(fd, p, offset)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Pwrite is an alias of golang.org/x/sys/unix.Pwrite, wrapped to automatically retry on EINTR.
func Pwrite(fd int, p []byte, offset int64) (n int, err error) {
	for {
		n, err = unix.Pwrite(fd, p, offset)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Accept4 is an alias of golang.org/x/sys/unix.Accept4, wrapped to automatically retry on EINTR.
func Accept4(fd, flags int) (nfd int, sa unix.Sockaddr, err error) {
	for {
		nfd, sa, err = unix.Accept4(fd, flags)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Getwd is an alias of golang.org/x/sys/unix.Getwd, wrapped to automatically retry on EINTR.
func Getwd() (string, error) {
	var (
		_v0 string
	)
	var (
		_v1 error
	)
	for {
		_v0, _v1 = unix.Getwd()
		if _v1 != syscall.EINTR {
			break
		}
	}
	return _v0, _v1
}

// Getfsstat is an alias of golang.org/x/sys/unix.Getfsstat, wrapped to automatically retry on EINTR.
func Getfsstat(buf []unix.Statfs_t, flags int) (n int, err error) {
	for {
		n, err = unix.Getfsstat(buf, flags)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Uname is an alias of golang.org/x/sys/unix.Uname, wrapped to automatically retry on EINTR.
func Uname(uname *unix.Utsname) error {
	var (
		_v0 error
	)
	for {
		_v0 = unix.Uname(uname)
		if _v0 != syscall.EINTR {
			break
		}
	}
	return _v0
}

// Sendfile is an alias of golang.org/x/sys/unix.Sendfile, wrapped to automatically retry on EINTR.
func Sendfile(outfd int, infd int, offset *int64, count int) (written int, err error) {
	for {
		written, err = unix.Sendfile(outfd, infd, offset, count)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// SetKevent is an alias of golang.org/x/sys/unix.SetKevent.
func SetKevent(k *unix.Kevent_t, fd, mode, flags int) {
	unix.SetKevent(k, fd, mode, flags)
}

// Syscall9 is an alias of golang.org/x/sys/unix.Syscall9, wrapped to automatically retry on EINTR.
func Syscall9(num, a1, a2, a3, a4, a5, a6, a7, a8, a9 uintptr) (r1, r2 uintptr, err syscall.Errno) {
	for {
		r1, r2, err = unix.Syscall9(num, a1, a2, a3, a4, a5, a6, a7, a8, a9)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// ErrnoName is an alias of golang.org/x/sys/unix.ErrnoName.
func ErrnoName(e syscall.Errno) string {
	return unix.ErrnoName(e)
}

// SignalName is an alias of golang.org/x/sys/unix.SignalName.
func SignalName(s syscall.Signal) string {
	return unix.SignalName(s)
}

// SignalNum is an alias of golang.org/x/sys/unix.SignalNum.
func SignalNum(s string) syscall.Signal {
	return unix.SignalNum(s)
}

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

// Bind is an alias of golang.org/x/sys/unix.Bind, wrapped to automatically retry on EINTR.
func Bind(fd int, sa unix.Sockaddr) (err error) {
	for {
		err = unix.Bind(fd, sa)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Connect is an alias of golang.org/x/sys/unix.Connect, wrapped to automatically retry on EINTR.
func Connect(fd int, sa unix.Sockaddr) (err error) {
	for {
		err = unix.Connect(fd, sa)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Getpeername is an alias of golang.org/x/sys/unix.Getpeername, wrapped to automatically retry on EINTR.
func Getpeername(fd int) (sa unix.Sockaddr, err error) {
	for {
		sa, err = unix.Getpeername(fd)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// GetsockoptByte is an alias of golang.org/x/sys/unix.GetsockoptByte, wrapped to automatically retry on EINTR.
func GetsockoptByte(fd, level, opt int) (value byte, err error) {
	for {
		value, err = unix.GetsockoptByte(fd, level, opt)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// GetsockoptInt is an alias of golang.org/x/sys/unix.GetsockoptInt, wrapped to automatically retry on EINTR.
func GetsockoptInt(fd, level, opt int) (value int, err error) {
	for {
		value, err = unix.GetsockoptInt(fd, level, opt)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// GetsockoptInet4Addr is an alias of golang.org/x/sys/unix.GetsockoptInet4Addr, wrapped to automatically retry on EINTR.
func GetsockoptInet4Addr(fd, level, opt int) (value [4]byte, err error) {
	for {
		value, err = unix.GetsockoptInet4Addr(fd, level, opt)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// GetsockoptIPMreq is an alias of golang.org/x/sys/unix.GetsockoptIPMreq, wrapped to automatically retry on EINTR.
func GetsockoptIPMreq(fd, level, opt int) (*unix.IPMreq, error) {
	var (
		_v0 *unix.IPMreq
	)
	var (
		_v1 error
	)
	for {
		_v0, _v1 = unix.GetsockoptIPMreq(fd, level, opt)
		if _v1 != syscall.EINTR {
			break
		}
	}
	return _v0, _v1
}

// GetsockoptIPv6Mreq is an alias of golang.org/x/sys/unix.GetsockoptIPv6Mreq, wrapped to automatically retry on EINTR.
func GetsockoptIPv6Mreq(fd, level, opt int) (*unix.IPv6Mreq, error) {
	var (
		_v0 *unix.IPv6Mreq
	)
	var (
		_v1 error
	)
	for {
		_v0, _v1 = unix.GetsockoptIPv6Mreq(fd, level, opt)
		if _v1 != syscall.EINTR {
			break
		}
	}
	return _v0, _v1
}

// GetsockoptIPv6MTUInfo is an alias of golang.org/x/sys/unix.GetsockoptIPv6MTUInfo, wrapped to automatically retry on EINTR.
func GetsockoptIPv6MTUInfo(fd, level, opt int) (*unix.IPv6MTUInfo, error) {
	var (
		_v0 *unix.IPv6MTUInfo
	)
	var (
		_v1 error
	)
	for {
		_v0, _v1 = unix.GetsockoptIPv6MTUInfo(fd, level, opt)
		if _v1 != syscall.EINTR {
			break
		}
	}
	return _v0, _v1
}

// GetsockoptICMPv6Filter is an alias of golang.org/x/sys/unix.GetsockoptICMPv6Filter, wrapped to automatically retry on EINTR.
func GetsockoptICMPv6Filter(fd, level, opt int) (*unix.ICMPv6Filter, error) {
	var (
		_v0 *unix.ICMPv6Filter
	)
	var (
		_v1 error
	)
	for {
		_v0, _v1 = unix.GetsockoptICMPv6Filter(fd, level, opt)
		if _v1 != syscall.EINTR {
			break
		}
	}
	return _v0, _v1
}

// GetsockoptLinger is an alias of golang.org/x/sys/unix.GetsockoptLinger, wrapped to automatically retry on EINTR.
func GetsockoptLinger(fd, level, opt int) (*unix.Linger, error) {
	var (
		_v0 *unix.Linger
	)
	var (
		_v1 error
	)
	for {
		_v0, _v1 = unix.GetsockoptLinger(fd, level, opt)
		if _v1 != syscall.EINTR {
			break
		}
	}
	return _v0, _v1
}

// GetsockoptTimeval is an alias of golang.org/x/sys/unix.GetsockoptTimeval, wrapped to automatically retry on EINTR.
func GetsockoptTimeval(fd, level, opt int) (*unix.Timeval, error) {
	var (
		_v0 *unix.Timeval
	)
	var (
		_v1 error
	)
	for {
		_v0, _v1 = unix.GetsockoptTimeval(fd, level, opt)
		if _v1 != syscall.EINTR {
			break
		}
	}
	return _v0, _v1
}

// GetsockoptUint64 is an alias of golang.org/x/sys/unix.GetsockoptUint64, wrapped to automatically retry on EINTR.
func GetsockoptUint64(fd, level, opt int) (value uint64, err error) {
	for {
		value, err = unix.GetsockoptUint64(fd, level, opt)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Recvfrom is an alias of golang.org/x/sys/unix.Recvfrom, wrapped to automatically retry on EINTR.
func Recvfrom(fd int, p []byte, flags int) (n int, from unix.Sockaddr, err error) {
	for {
		n, from, err = unix.Recvfrom(fd, p, flags)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Sendto is an alias of golang.org/x/sys/unix.Sendto, wrapped to automatically retry on EINTR.
func Sendto(fd int, p []byte, flags int, to unix.Sockaddr) (err error) {
	for {
		err = unix.Sendto(fd, p, flags, to)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// SetsockoptByte is an alias of golang.org/x/sys/unix.SetsockoptByte, wrapped to automatically retry on EINTR.
func SetsockoptByte(fd, level, opt int, value byte) (err error) {
	for {
		err = unix.SetsockoptByte(fd, level, opt, value)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// SetsockoptInt is an alias of golang.org/x/sys/unix.SetsockoptInt, wrapped to automatically retry on EINTR.
func SetsockoptInt(fd, level, opt int, value int) (err error) {
	for {
		err = unix.SetsockoptInt(fd, level, opt, value)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// SetsockoptInet4Addr is an alias of golang.org/x/sys/unix.SetsockoptInet4Addr, wrapped to automatically retry on EINTR.
func SetsockoptInet4Addr(fd, level, opt int, value [4]byte) (err error) {
	for {
		err = unix.SetsockoptInet4Addr(fd, level, opt, value)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// SetsockoptIPMreq is an alias of golang.org/x/sys/unix.SetsockoptIPMreq, wrapped to automatically retry on EINTR.
func SetsockoptIPMreq(fd, level, opt int, mreq *unix.IPMreq) (err error) {
	for {
		err = unix.SetsockoptIPMreq(fd, level, opt, mreq)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// SetsockoptIPv6Mreq is an alias of golang.org/x/sys/unix.SetsockoptIPv6Mreq, wrapped to automatically retry on EINTR.
func SetsockoptIPv6Mreq(fd, level, opt int, mreq *unix.IPv6Mreq) (err error) {
	for {
		err = unix.SetsockoptIPv6Mreq(fd, level, opt, mreq)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// SetsockoptICMPv6Filter is an alias of golang.org/x/sys/unix.SetsockoptICMPv6Filter, wrapped to automatically retry on EINTR.
func SetsockoptICMPv6Filter(fd, level, opt int, filter *unix.ICMPv6Filter) error {
	var (
		_v0 error
	)
	for {
		_v0 = unix.SetsockoptICMPv6Filter(fd, level, opt, filter)
		if _v0 != syscall.EINTR {
			break
		}
	}
	return _v0
}

// SetsockoptLinger is an alias of golang.org/x/sys/unix.SetsockoptLinger, wrapped to automatically retry on EINTR.
func SetsockoptLinger(fd, level, opt int, l *unix.Linger) (err error) {
	for {
		err = unix.SetsockoptLinger(fd, level, opt, l)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// SetsockoptString is an alias of golang.org/x/sys/unix.SetsockoptString, wrapped to automatically retry on EINTR.
func SetsockoptString(fd, level, opt int, s string) (err error) {
	for {
		err = unix.SetsockoptString(fd, level, opt, s)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// SetsockoptTimeval is an alias of golang.org/x/sys/unix.SetsockoptTimeval, wrapped to automatically retry on EINTR.
func SetsockoptTimeval(fd, level, opt int, tv *unix.Timeval) (err error) {
	for {
		err = unix.SetsockoptTimeval(fd, level, opt, tv)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// SetsockoptUint64 is an alias of golang.org/x/sys/unix.SetsockoptUint64, wrapped to automatically retry on EINTR.
func SetsockoptUint64(fd, level, opt int, value uint64) (err error) {
	for {
		err = unix.SetsockoptUint64(fd, level, opt, value)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Socket is an alias of golang.org/x/sys/unix.Socket, wrapped to automatically retry on EINTR.
func Socket(domain, typ, proto int) (fd int, err error) {
	for {
		fd, err = unix.Socket(domain, typ, proto)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Socketpair is an alias of golang.org/x/sys/unix.Socketpair, wrapped to automatically retry on EINTR.
func Socketpair(domain, typ, proto int) (fd [2]int, err error) {
	for {
		fd, err = unix.Socketpair(domain, typ, proto)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// CloseOnExec is an alias of golang.org/x/sys/unix.CloseOnExec.
func CloseOnExec(fd int) {
	unix.CloseOnExec(fd)
}

// SetNonblock is an alias of golang.org/x/sys/unix.SetNonblock, wrapped to automatically retry on EINTR.
func SetNonblock(fd int, nonblocking bool) (err error) {
	for {
		err = unix.SetNonblock(fd, nonblocking)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Exec is an alias of golang.org/x/sys/unix.Exec, wrapped to automatically retry on EINTR.
func Exec(argv0 string, argv []string, envv []string) error {
	var (
		_v0 error
	)
	for {
		_v0 = unix.Exec(argv0, argv, envv)
		if _v0 != syscall.EINTR {
			break
		}
	}
	return _v0
}

// Lutimes is an alias of golang.org/x/sys/unix.Lutimes, wrapped to automatically retry on EINTR.
func Lutimes(path string, tv []unix.Timeval) error {
	var (
		_v0 error
	)
	for {
		_v0 = unix.Lutimes(path, tv)
		if _v0 != syscall.EINTR {
			break
		}
	}
	return _v0
}

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

// TimespecToNsec is an alias of golang.org/x/sys/unix.TimespecToNsec.
func TimespecToNsec(ts unix.Timespec) int64 {
	return unix.TimespecToNsec(ts)
}

// NsecToTimespec is an alias of golang.org/x/sys/unix.NsecToTimespec.
func NsecToTimespec(nsec int64) unix.Timespec {
	return unix.NsecToTimespec(nsec)
}

// TimeToTimespec is an alias of golang.org/x/sys/unix.TimeToTimespec, wrapped to automatically retry on EINTR.
func TimeToTimespec(t time.Time) (unix.Timespec, error) {
	var (
		_v0 unix.Timespec
	)
	var (
		_v1 error
	)
	for {
		_v0, _v1 = unix.TimeToTimespec(t)
		if _v1 != syscall.EINTR {
			break
		}
	}
	return _v0, _v1
}

// TimevalToNsec is an alias of golang.org/x/sys/unix.TimevalToNsec.
func TimevalToNsec(tv unix.Timeval) int64 {
	return unix.TimevalToNsec(tv)
}

// NsecToTimeval is an alias of golang.org/x/sys/unix.NsecToTimeval.
func NsecToTimeval(nsec int64) unix.Timeval {
	return unix.NsecToTimeval(nsec)
}

// Shutdown is an alias of golang.org/x/sys/unix.Shutdown, wrapped to automatically retry on EINTR.
func Shutdown(s int, how int) (err error) {
	for {
		err = unix.Shutdown(s, how)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Madvise is an alias of golang.org/x/sys/unix.Madvise, wrapped to automatically retry on EINTR.
func Madvise(b []byte, behav int) (err error) {
	for {
		err = unix.Madvise(b, behav)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Mlock is an alias of golang.org/x/sys/unix.Mlock, wrapped to automatically retry on EINTR.
func Mlock(b []byte) (err error) {
	for {
		err = unix.Mlock(b)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Mlockall is an alias of golang.org/x/sys/unix.Mlockall, wrapped to automatically retry on EINTR.
func Mlockall(flags int) (err error) {
	for {
		err = unix.Mlockall(flags)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Mprotect is an alias of golang.org/x/sys/unix.Mprotect, wrapped to automatically retry on EINTR.
func Mprotect(b []byte, prot int) (err error) {
	for {
		err = unix.Mprotect(b, prot)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Msync is an alias of golang.org/x/sys/unix.Msync, wrapped to automatically retry on EINTR.
func Msync(b []byte, flags int) (err error) {
	for {
		err = unix.Msync(b, flags)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Munlock is an alias of golang.org/x/sys/unix.Munlock, wrapped to automatically retry on EINTR.
func Munlock(b []byte) (err error) {
	for {
		err = unix.Munlock(b)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Munlockall is an alias of golang.org/x/sys/unix.Munlockall, wrapped to automatically retry on EINTR.
func Munlockall() (err error) {
	for {
		err = unix.Munlockall()
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Getcwd is an alias of golang.org/x/sys/unix.Getcwd, wrapped to automatically retry on EINTR.
func Getcwd(buf []byte) (n int, err error) {
	for {
		n, err = unix.Getcwd(buf)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Access is an alias of golang.org/x/sys/unix.Access, wrapped to automatically retry on EINTR.
func Access(path string, mode uint32) (err error) {
	for {
		err = unix.Access(path, mode)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Adjtime is an alias of golang.org/x/sys/unix.Adjtime, wrapped to automatically retry on EINTR.
func Adjtime(delta *unix.Timeval, olddelta *unix.Timeval) (err error) {
	for {
		err = unix.Adjtime(delta, olddelta)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Chdir is an alias of golang.org/x/sys/unix.Chdir, wrapped to automatically retry on EINTR.
func Chdir(path string) (err error) {
	for {
		err = unix.Chdir(path)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Chflags is an alias of golang.org/x/sys/unix.Chflags, wrapped to automatically retry on EINTR.
func Chflags(path string, flags int) (err error) {
	for {
		err = unix.Chflags(path, flags)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Chmod is an alias of golang.org/x/sys/unix.Chmod, wrapped to automatically retry on EINTR.
func Chmod(path string, mode uint32) (err error) {
	for {
		err = unix.Chmod(path, mode)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Chown is an alias of golang.org/x/sys/unix.Chown, wrapped to automatically retry on EINTR.
func Chown(path string, uid int, gid int) (err error) {
	for {
		err = unix.Chown(path, uid, gid)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Chroot is an alias of golang.org/x/sys/unix.Chroot, wrapped to automatically retry on EINTR.
func Chroot(path string) (err error) {
	for {
		err = unix.Chroot(path)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Close is an alias of golang.org/x/sys/unix.Close. Does NOT retry on EINTR. See close(2) for the reason.
func Close(fd int) (err error) {
	return unix.Close(fd)
}

// Dup is an alias of golang.org/x/sys/unix.Dup, wrapped to automatically retry on EINTR.
func Dup(fd int) (nfd int, err error) {
	for {
		nfd, err = unix.Dup(fd)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Dup2 is an alias of golang.org/x/sys/unix.Dup2, wrapped to automatically retry on EINTR.
func Dup2(from int, to int) (err error) {
	for {
		err = unix.Dup2(from, to)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Exit is an alias of golang.org/x/sys/unix.Exit.
func Exit(code int) {
	unix.Exit(code)
}

// Faccessat is an alias of golang.org/x/sys/unix.Faccessat, wrapped to automatically retry on EINTR.
func Faccessat(dirfd int, path string, mode uint32, flags int) (err error) {
	for {
		err = unix.Faccessat(dirfd, path, mode, flags)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Fchdir is an alias of golang.org/x/sys/unix.Fchdir, wrapped to automatically retry on EINTR.
func Fchdir(fd int) (err error) {
	for {
		err = unix.Fchdir(fd)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Fchflags is an alias of golang.org/x/sys/unix.Fchflags, wrapped to automatically retry on EINTR.
func Fchflags(fd int, flags int) (err error) {
	for {
		err = unix.Fchflags(fd, flags)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Fchmod is an alias of golang.org/x/sys/unix.Fchmod, wrapped to automatically retry on EINTR.
func Fchmod(fd int, mode uint32) (err error) {
	for {
		err = unix.Fchmod(fd, mode)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Fchmodat is an alias of golang.org/x/sys/unix.Fchmodat, wrapped to automatically retry on EINTR.
func Fchmodat(dirfd int, path string, mode uint32, flags int) (err error) {
	for {
		err = unix.Fchmodat(dirfd, path, mode, flags)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Fchown is an alias of golang.org/x/sys/unix.Fchown, wrapped to automatically retry on EINTR.
func Fchown(fd int, uid int, gid int) (err error) {
	for {
		err = unix.Fchown(fd, uid, gid)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Fchownat is an alias of golang.org/x/sys/unix.Fchownat, wrapped to automatically retry on EINTR.
func Fchownat(dirfd int, path string, uid int, gid int, flags int) (err error) {
	for {
		err = unix.Fchownat(dirfd, path, uid, gid, flags)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Flock is an alias of golang.org/x/sys/unix.Flock, wrapped to automatically retry on EINTR.
func Flock(fd int, how int) (err error) {
	for {
		err = unix.Flock(fd, how)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Fpathconf is an alias of golang.org/x/sys/unix.Fpathconf, wrapped to automatically retry on EINTR.
func Fpathconf(fd int, name int) (val int, err error) {
	for {
		val, err = unix.Fpathconf(fd, name)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Fstat is an alias of golang.org/x/sys/unix.Fstat, wrapped to automatically retry on EINTR.
func Fstat(fd int, stat *unix.Stat_t) (err error) {
	for {
		err = unix.Fstat(fd, stat)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Fstatat is an alias of golang.org/x/sys/unix.Fstatat, wrapped to automatically retry on EINTR.
func Fstatat(fd int, path string, stat *unix.Stat_t, flags int) (err error) {
	for {
		err = unix.Fstatat(fd, path, stat, flags)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Fstatfs is an alias of golang.org/x/sys/unix.Fstatfs, wrapped to automatically retry on EINTR.
func Fstatfs(fd int, stat *unix.Statfs_t) (err error) {
	for {
		err = unix.Fstatfs(fd, stat)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Fsync is an alias of golang.org/x/sys/unix.Fsync, wrapped to automatically retry on EINTR.
func Fsync(fd int) (err error) {
	for {
		err = unix.Fsync(fd)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Ftruncate is an alias of golang.org/x/sys/unix.Ftruncate, wrapped to automatically retry on EINTR.
func Ftruncate(fd int, length int64) (err error) {
	for {
		err = unix.Ftruncate(fd, length)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Getdents is an alias of golang.org/x/sys/unix.Getdents, wrapped to automatically retry on EINTR.
func Getdents(fd int, buf []byte) (n int, err error) {
	for {
		n, err = unix.Getdents(fd, buf)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Getdirentries is an alias of golang.org/x/sys/unix.Getdirentries, wrapped to automatically retry on EINTR.
func Getdirentries(fd int, buf []byte, basep *uintptr) (n int, err error) {
	for {
		n, err = unix.Getdirentries(fd, buf, basep)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Getdtablesize is an alias of golang.org/x/sys/unix.Getdtablesize.
func Getdtablesize() (size int) {
	return unix.Getdtablesize()
}

// Getegid is an alias of golang.org/x/sys/unix.Getegid.
func Getegid() (egid int) {
	return unix.Getegid()
}

// Geteuid is an alias of golang.org/x/sys/unix.Geteuid.
func Geteuid() (uid int) {
	return unix.Geteuid()
}

// Getgid is an alias of golang.org/x/sys/unix.Getgid.
func Getgid() (gid int) {
	return unix.Getgid()
}

// Getpgid is an alias of golang.org/x/sys/unix.Getpgid, wrapped to automatically retry on EINTR.
func Getpgid(pid int) (pgid int, err error) {
	for {
		pgid, err = unix.Getpgid(pid)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Getpgrp is an alias of golang.org/x/sys/unix.Getpgrp.
func Getpgrp() (pgrp int) {
	return unix.Getpgrp()
}

// Getpid is an alias of golang.org/x/sys/unix.Getpid.
func Getpid() (pid int) {
	return unix.Getpid()
}

// Getppid is an alias of golang.org/x/sys/unix.Getppid.
func Getppid() (ppid int) {
	return unix.Getppid()
}

// Getpriority is an alias of golang.org/x/sys/unix.Getpriority, wrapped to automatically retry on EINTR.
func Getpriority(which int, who int) (prio int, err error) {
	for {
		prio, err = unix.Getpriority(which, who)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Getrlimit is an alias of golang.org/x/sys/unix.Getrlimit, wrapped to automatically retry on EINTR.
func Getrlimit(which int, lim *unix.Rlimit) (err error) {
	for {
		err = unix.Getrlimit(which, lim)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Getrusage is an alias of golang.org/x/sys/unix.Getrusage, wrapped to automatically retry on EINTR.
func Getrusage(who int, rusage *unix.Rusage) (err error) {
	for {
		err = unix.Getrusage(who, rusage)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Getsid is an alias of golang.org/x/sys/unix.Getsid, wrapped to automatically retry on EINTR.
func Getsid(pid int) (sid int, err error) {
	for {
		sid, err = unix.Getsid(pid)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Gettimeofday is an alias of golang.org/x/sys/unix.Gettimeofday, wrapped to automatically retry on EINTR.
func Gettimeofday(tv *unix.Timeval) (err error) {
	for {
		err = unix.Gettimeofday(tv)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Getuid is an alias of golang.org/x/sys/unix.Getuid.
func Getuid() (uid int) {
	return unix.Getuid()
}

// Issetugid is an alias of golang.org/x/sys/unix.Issetugid.
func Issetugid() (tainted bool) {
	return unix.Issetugid()
}

// Kill is an alias of golang.org/x/sys/unix.Kill, wrapped to automatically retry on EINTR.
func Kill(pid int, signum syscall.Signal) (err error) {
	for {
		err = unix.Kill(pid, signum)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Kqueue is an alias of golang.org/x/sys/unix.Kqueue, wrapped to automatically retry on EINTR.
func Kqueue() (fd int, err error) {
	for {
		fd, err = unix.Kqueue()
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Lchown is an alias of golang.org/x/sys/unix.Lchown, wrapped to automatically retry on EINTR.
func Lchown(path string, uid int, gid int) (err error) {
	for {
		err = unix.Lchown(path, uid, gid)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Link is an alias of golang.org/x/sys/unix.Link, wrapped to automatically retry on EINTR.
func Link(path string, link string) (err error) {
	for {
		err = unix.Link(path, link)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Linkat is an alias of golang.org/x/sys/unix.Linkat, wrapped to automatically retry on EINTR.
func Linkat(pathfd int, path string, linkfd int, link string, flags int) (err error) {
	for {
		err = unix.Linkat(pathfd, path, linkfd, link, flags)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Listen is an alias of golang.org/x/sys/unix.Listen, wrapped to automatically retry on EINTR.
func Listen(s int, backlog int) (err error) {
	for {
		err = unix.Listen(s, backlog)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Lstat is an alias of golang.org/x/sys/unix.Lstat, wrapped to automatically retry on EINTR.
func Lstat(path string, stat *unix.Stat_t) (err error) {
	for {
		err = unix.Lstat(path, stat)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Mkdir is an alias of golang.org/x/sys/unix.Mkdir, wrapped to automatically retry on EINTR.
func Mkdir(path string, mode uint32) (err error) {
	for {
		err = unix.Mkdir(path, mode)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Mkdirat is an alias of golang.org/x/sys/unix.Mkdirat, wrapped to automatically retry on EINTR.
func Mkdirat(dirfd int, path string, mode uint32) (err error) {
	for {
		err = unix.Mkdirat(dirfd, path, mode)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Mkfifo is an alias of golang.org/x/sys/unix.Mkfifo, wrapped to automatically retry on EINTR.
func Mkfifo(path string, mode uint32) (err error) {
	for {
		err = unix.Mkfifo(path, mode)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Mknod is an alias of golang.org/x/sys/unix.Mknod, wrapped to automatically retry on EINTR.
func Mknod(path string, mode uint32, dev int) (err error) {
	for {
		err = unix.Mknod(path, mode, dev)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Mknodat is an alias of golang.org/x/sys/unix.Mknodat, wrapped to automatically retry on EINTR.
func Mknodat(fd int, path string, mode uint32, dev int) (err error) {
	for {
		err = unix.Mknodat(fd, path, mode, dev)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Nanosleep is an alias of golang.org/x/sys/unix.Nanosleep, wrapped to automatically retry on EINTR.
func Nanosleep(time *unix.Timespec, leftover *unix.Timespec) (err error) {
	for {
		err = unix.Nanosleep(time, leftover)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Open is an alias of golang.org/x/sys/unix.Open, wrapped to automatically retry on EINTR.
func Open(path string, mode int, perm uint32) (fd int, err error) {
	for {
		fd, err = unix.Open(path, mode, perm)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Openat is an alias of golang.org/x/sys/unix.Openat, wrapped to automatically retry on EINTR.
func Openat(dirfd int, path string, mode int, perm uint32) (fd int, err error) {
	for {
		fd, err = unix.Openat(dirfd, path, mode, perm)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Pathconf is an alias of golang.org/x/sys/unix.Pathconf, wrapped to automatically retry on EINTR.
func Pathconf(path string, name int) (val int, err error) {
	for {
		val, err = unix.Pathconf(path, name)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Readlink is an alias of golang.org/x/sys/unix.Readlink, wrapped to automatically retry on EINTR.
func Readlink(path string, buf []byte) (n int, err error) {
	for {
		n, err = unix.Readlink(path, buf)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Rename is an alias of golang.org/x/sys/unix.Rename, wrapped to automatically retry on EINTR.
func Rename(from string, to string) (err error) {
	for {
		err = unix.Rename(from, to)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Renameat is an alias of golang.org/x/sys/unix.Renameat, wrapped to automatically retry on EINTR.
func Renameat(fromfd int, from string, tofd int, to string) (err error) {
	for {
		err = unix.Renameat(fromfd, from, tofd, to)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Revoke is an alias of golang.org/x/sys/unix.Revoke, wrapped to automatically retry on EINTR.
func Revoke(path string) (err error) {
	for {
		err = unix.Revoke(path)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Rmdir is an alias of golang.org/x/sys/unix.Rmdir, wrapped to automatically retry on EINTR.
func Rmdir(path string) (err error) {
	for {
		err = unix.Rmdir(path)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Seek is an alias of golang.org/x/sys/unix.Seek, wrapped to automatically retry on EINTR.
func Seek(fd int, offset int64, whence int) (newoffset int64, err error) {
	for {
		newoffset, err = unix.Seek(fd, offset, whence)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Select is an alias of golang.org/x/sys/unix.Select, wrapped to automatically retry on EINTR.
func Select(n int, r *unix.FdSet, w *unix.FdSet, e *unix.FdSet, timeout *unix.Timeval) (err error) {
	for {
		err = unix.Select(n, r, w, e, timeout)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Setegid is an alias of golang.org/x/sys/unix.Setegid, wrapped to automatically retry on EINTR.
func Setegid(egid int) (err error) {
	for {
		err = unix.Setegid(egid)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Seteuid is an alias of golang.org/x/sys/unix.Seteuid, wrapped to automatically retry on EINTR.
func Seteuid(euid int) (err error) {
	for {
		err = unix.Seteuid(euid)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Setgid is an alias of golang.org/x/sys/unix.Setgid, wrapped to automatically retry on EINTR.
func Setgid(gid int) (err error) {
	for {
		err = unix.Setgid(gid)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Setlogin is an alias of golang.org/x/sys/unix.Setlogin, wrapped to automatically retry on EINTR.
func Setlogin(name string) (err error) {
	for {
		err = unix.Setlogin(name)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Setpgid is an alias of golang.org/x/sys/unix.Setpgid, wrapped to automatically retry on EINTR.
func Setpgid(pid int, pgid int) (err error) {
	for {
		err = unix.Setpgid(pid, pgid)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Setpriority is an alias of golang.org/x/sys/unix.Setpriority, wrapped to automatically retry on EINTR.
func Setpriority(which int, who int, prio int) (err error) {
	for {
		err = unix.Setpriority(which, who, prio)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Setregid is an alias of golang.org/x/sys/unix.Setregid, wrapped to automatically retry on EINTR.
func Setregid(rgid int, egid int) (err error) {
	for {
		err = unix.Setregid(rgid, egid)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Setreuid is an alias of golang.org/x/sys/unix.Setreuid, wrapped to automatically retry on EINTR.
func Setreuid(ruid int, euid int) (err error) {
	for {
		err = unix.Setreuid(ruid, euid)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Setresgid is an alias of golang.org/x/sys/unix.Setresgid, wrapped to automatically retry on EINTR.
func Setresgid(rgid int, egid int, sgid int) (err error) {
	for {
		err = unix.Setresgid(rgid, egid, sgid)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Setresuid is an alias of golang.org/x/sys/unix.Setresuid, wrapped to automatically retry on EINTR.
func Setresuid(ruid int, euid int, suid int) (err error) {
	for {
		err = unix.Setresuid(ruid, euid, suid)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Setrlimit is an alias of golang.org/x/sys/unix.Setrlimit, wrapped to automatically retry on EINTR.
func Setrlimit(which int, lim *unix.Rlimit) (err error) {
	for {
		err = unix.Setrlimit(which, lim)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Setsid is an alias of golang.org/x/sys/unix.Setsid, wrapped to automatically retry on EINTR.
func Setsid() (pid int, err error) {
	for {
		pid, err = unix.Setsid()
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Settimeofday is an alias of golang.org/x/sys/unix.Settimeofday, wrapped to automatically retry on EINTR.
func Settimeofday(tp *unix.Timeval) (err error) {
	for {
		err = unix.Settimeofday(tp)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Setuid is an alias of golang.org/x/sys/unix.Setuid, wrapped to automatically retry on EINTR.
func Setuid(uid int) (err error) {
	for {
		err = unix.Setuid(uid)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Stat is an alias of golang.org/x/sys/unix.Stat, wrapped to automatically retry on EINTR.
func Stat(path string, stat *unix.Stat_t) (err error) {
	for {
		err = unix.Stat(path, stat)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Statfs is an alias of golang.org/x/sys/unix.Statfs, wrapped to automatically retry on EINTR.
func Statfs(path string, stat *unix.Statfs_t) (err error) {
	for {
		err = unix.Statfs(path, stat)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Symlink is an alias of golang.org/x/sys/unix.Symlink, wrapped to automatically retry on EINTR.
func Symlink(path string, link string) (err error) {
	for {
		err = unix.Symlink(path, link)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Symlinkat is an alias of golang.org/x/sys/unix.Symlinkat, wrapped to automatically retry on EINTR.
func Symlinkat(oldpath string, newdirfd int, newpath string) (err error) {
	for {
		err = unix.Symlinkat(oldpath, newdirfd, newpath)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Sync is an alias of golang.org/x/sys/unix.Sync, wrapped to automatically retry on EINTR.
func Sync() (err error) {
	for {
		err = unix.Sync()
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Truncate is an alias of golang.org/x/sys/unix.Truncate, wrapped to automatically retry on EINTR.
func Truncate(path string, length int64) (err error) {
	for {
		err = unix.Truncate(path, length)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Umask is an alias of golang.org/x/sys/unix.Umask.
func Umask(newmask int) (oldmask int) {
	return unix.Umask(newmask)
}

// Undelete is an alias of golang.org/x/sys/unix.Undelete, wrapped to automatically retry on EINTR.
func Undelete(path string) (err error) {
	for {
		err = unix.Undelete(path)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Unlink is an alias of golang.org/x/sys/unix.Unlink, wrapped to automatically retry on EINTR.
func Unlink(path string) (err error) {
	for {
		err = unix.Unlink(path)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Unlinkat is an alias of golang.org/x/sys/unix.Unlinkat, wrapped to automatically retry on EINTR.
func Unlinkat(dirfd int, path string, flags int) (err error) {
	for {
		err = unix.Unlinkat(dirfd, path, flags)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Unmount is an alias of golang.org/x/sys/unix.Unmount, wrapped to automatically retry on EINTR.
func Unmount(path string, flags int) (err error) {
	for {
		err = unix.Unmount(path, flags)
		if err != syscall.EINTR {
			break
		}
	}
	return
}
