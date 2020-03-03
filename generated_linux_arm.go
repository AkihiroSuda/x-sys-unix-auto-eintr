// Automatically generated file. DO NOT EDIT MANUALLY.

package unix

import (
	"syscall"
	"time"

	unix "golang.org/x/sys/unix"
)

// SchedGetaffinity is an alias of golang.org/x/sys/unix.SchedGetaffinity, wrapped to automatically retry on EINTR.
func SchedGetaffinity(pid int, set *unix.CPUSet) error {
	var (
		_v0 error
	)
	for {
		_v0 = unix.SchedGetaffinity(pid, set)
		if _v0 != syscall.EINTR {
			break
		}
	}
	return _v0
}

// SchedSetaffinity is an alias of golang.org/x/sys/unix.SchedSetaffinity, wrapped to automatically retry on EINTR.
func SchedSetaffinity(pid int, set *unix.CPUSet) error {
	var (
		_v0 error
	)
	for {
		_v0 = unix.SchedSetaffinity(pid, set)
		if _v0 != syscall.EINTR {
			break
		}
	}
	return _v0
}

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
		_v0 *unix.Winsize
		_v1 error
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

// UnixCredentials is an alias of golang.org/x/sys/unix.UnixCredentials.
func UnixCredentials(ucred *unix.Ucred) []byte {
	return unix.UnixCredentials(ucred)
}

// ParseUnixCredentials is an alias of golang.org/x/sys/unix.ParseUnixCredentials, wrapped to automatically retry on EINTR.
func ParseUnixCredentials(m *unix.SocketControlMessage) (*unix.Ucred, error) {
	var (
		_v0 *unix.Ucred
		_v1 error
	)
	for {
		_v0, _v1 = unix.ParseUnixCredentials(m)
		if _v1 != syscall.EINTR {
			break
		}
	}
	return _v0, _v1
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

// Creat is an alias of golang.org/x/sys/unix.Creat, wrapped to automatically retry on EINTR.
func Creat(path string, mode uint32) (fd int, err error) {
	for {
		fd, err = unix.Creat(path, mode)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// FanotifyMark is an alias of golang.org/x/sys/unix.FanotifyMark, wrapped to automatically retry on EINTR.
func FanotifyMark(fd int, flags uint, mask uint64, dirFd int, pathname string) (err error) {
	for {
		err = unix.FanotifyMark(fd, flags, mask, dirFd, pathname)
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

// IoctlRetInt is an alias of golang.org/x/sys/unix.IoctlRetInt, wrapped to automatically retry on EINTR.
func IoctlRetInt(fd int, req uint) (int, error) {
	var (
		_v0 int
		_v1 error
	)
	for {
		_v0, _v1 = unix.IoctlRetInt(fd, req)
		if _v1 != syscall.EINTR {
			break
		}
	}
	return _v0, _v1
}

// IoctlSetPointerInt is an alias of golang.org/x/sys/unix.IoctlSetPointerInt, wrapped to automatically retry on EINTR.
func IoctlSetPointerInt(fd int, req uint, value int) error {
	var (
		_v0 error
	)
	for {
		_v0 = unix.IoctlSetPointerInt(fd, req, value)
		if _v0 != syscall.EINTR {
			break
		}
	}
	return _v0
}

// IoctlSetRTCTime is an alias of golang.org/x/sys/unix.IoctlSetRTCTime, wrapped to automatically retry on EINTR.
func IoctlSetRTCTime(fd int, value *unix.RTCTime) error {
	var (
		_v0 error
	)
	for {
		_v0 = unix.IoctlSetRTCTime(fd, value)
		if _v0 != syscall.EINTR {
			break
		}
	}
	return _v0
}

// IoctlGetUint32 is an alias of golang.org/x/sys/unix.IoctlGetUint32, wrapped to automatically retry on EINTR.
func IoctlGetUint32(fd int, req uint) (uint32, error) {
	var (
		_v0 uint32
		_v1 error
	)
	for {
		_v0, _v1 = unix.IoctlGetUint32(fd, req)
		if _v1 != syscall.EINTR {
			break
		}
	}
	return _v0, _v1
}

// IoctlGetRTCTime is an alias of golang.org/x/sys/unix.IoctlGetRTCTime, wrapped to automatically retry on EINTR.
func IoctlGetRTCTime(fd int) (*unix.RTCTime, error) {
	var (
		_v0 *unix.RTCTime
		_v1 error
	)
	for {
		_v0, _v1 = unix.IoctlGetRTCTime(fd)
		if _v1 != syscall.EINTR {
			break
		}
	}
	return _v0, _v1
}

// Link is an alias of golang.org/x/sys/unix.Link, wrapped to automatically retry on EINTR.
func Link(oldpath string, newpath string) (err error) {
	for {
		err = unix.Link(oldpath, newpath)
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
func Openat(dirfd int, path string, flags int, mode uint32) (fd int, err error) {
	for {
		fd, err = unix.Openat(dirfd, path, flags, mode)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Ppoll is an alias of golang.org/x/sys/unix.Ppoll, wrapped to automatically retry on EINTR.
func Ppoll(fds []unix.PollFd, timeout *unix.Timespec, sigmask *unix.Sigset_t) (n int, err error) {
	for {
		n, err = unix.Ppoll(fds, timeout, sigmask)
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
func Rename(oldpath string, newpath string) (err error) {
	for {
		err = unix.Rename(oldpath, newpath)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Rmdir is an alias of golang.org/x/sys/unix.Rmdir, wrapped to automatically retry on EINTR.
func Rmdir(path string) error {
	var (
		_v0 error
	)
	for {
		_v0 = unix.Rmdir(path)
		if _v0 != syscall.EINTR {
			break
		}
	}
	return _v0
}

// Symlink is an alias of golang.org/x/sys/unix.Symlink, wrapped to automatically retry on EINTR.
func Symlink(oldpath string, newpath string) (err error) {
	for {
		err = unix.Symlink(oldpath, newpath)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Unlink is an alias of golang.org/x/sys/unix.Unlink, wrapped to automatically retry on EINTR.
func Unlink(path string) error {
	var (
		_v0 error
	)
	for {
		_v0 = unix.Unlink(path)
		if _v0 != syscall.EINTR {
			break
		}
	}
	return _v0
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

// Futimesat is an alias of golang.org/x/sys/unix.Futimesat, wrapped to automatically retry on EINTR.
func Futimesat(dirfd int, path string, tv []unix.Timeval) error {
	var (
		_v0 error
	)
	for {
		_v0 = unix.Futimesat(dirfd, path, tv)
		if _v0 != syscall.EINTR {
			break
		}
	}
	return _v0
}

// Futimes is an alias of golang.org/x/sys/unix.Futimes, wrapped to automatically retry on EINTR.
func Futimes(fd int, tv []unix.Timeval) (err error) {
	for {
		err = unix.Futimes(fd, tv)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Getwd is an alias of golang.org/x/sys/unix.Getwd, wrapped to automatically retry on EINTR.
func Getwd() (wd string, err error) {
	for {
		wd, err = unix.Getwd()
		if err != syscall.EINTR {
			break
		}
	}
	return
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

// Mkfifo is an alias of golang.org/x/sys/unix.Mkfifo, wrapped to automatically retry on EINTR.
func Mkfifo(path string, mode uint32) error {
	var (
		_v0 error
	)
	for {
		_v0 = unix.Mkfifo(path, mode)
		if _v0 != syscall.EINTR {
			break
		}
	}
	return _v0
}

// Mkfifoat is an alias of golang.org/x/sys/unix.Mkfifoat, wrapped to automatically retry on EINTR.
func Mkfifoat(dirfd int, path string, mode uint32) error {
	var (
		_v0 error
	)
	for {
		_v0 = unix.Mkfifoat(dirfd, path, mode)
		if _v0 != syscall.EINTR {
			break
		}
	}
	return _v0
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

// Accept4 is an alias of golang.org/x/sys/unix.Accept4, wrapped to automatically retry on EINTR.
func Accept4(fd int, flags int) (nfd int, sa unix.Sockaddr, err error) {
	for {
		nfd, sa, err = unix.Accept4(fd, flags)
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

// GetsockoptIPMreqn is an alias of golang.org/x/sys/unix.GetsockoptIPMreqn, wrapped to automatically retry on EINTR.
func GetsockoptIPMreqn(fd, level, opt int) (*unix.IPMreqn, error) {
	var (
		_v0 *unix.IPMreqn
		_v1 error
	)
	for {
		_v0, _v1 = unix.GetsockoptIPMreqn(fd, level, opt)
		if _v1 != syscall.EINTR {
			break
		}
	}
	return _v0, _v1
}

// GetsockoptUcred is an alias of golang.org/x/sys/unix.GetsockoptUcred, wrapped to automatically retry on EINTR.
func GetsockoptUcred(fd, level, opt int) (*unix.Ucred, error) {
	var (
		_v0 *unix.Ucred
		_v1 error
	)
	for {
		_v0, _v1 = unix.GetsockoptUcred(fd, level, opt)
		if _v1 != syscall.EINTR {
			break
		}
	}
	return _v0, _v1
}

// GetsockoptTCPInfo is an alias of golang.org/x/sys/unix.GetsockoptTCPInfo, wrapped to automatically retry on EINTR.
func GetsockoptTCPInfo(fd, level, opt int) (*unix.TCPInfo, error) {
	var (
		_v0 *unix.TCPInfo
		_v1 error
	)
	for {
		_v0, _v1 = unix.GetsockoptTCPInfo(fd, level, opt)
		if _v1 != syscall.EINTR {
			break
		}
	}
	return _v0, _v1
}

// GetsockoptString is an alias of golang.org/x/sys/unix.GetsockoptString, wrapped to automatically retry on EINTR.
func GetsockoptString(fd, level, opt int) (string, error) {
	var (
		_v0 string
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

// GetsockoptTpacketStats is an alias of golang.org/x/sys/unix.GetsockoptTpacketStats, wrapped to automatically retry on EINTR.
func GetsockoptTpacketStats(fd, level, opt int) (*unix.TpacketStats, error) {
	var (
		_v0 *unix.TpacketStats
		_v1 error
	)
	for {
		_v0, _v1 = unix.GetsockoptTpacketStats(fd, level, opt)
		if _v1 != syscall.EINTR {
			break
		}
	}
	return _v0, _v1
}

// GetsockoptTpacketStatsV3 is an alias of golang.org/x/sys/unix.GetsockoptTpacketStatsV3, wrapped to automatically retry on EINTR.
func GetsockoptTpacketStatsV3(fd, level, opt int) (*unix.TpacketStatsV3, error) {
	var (
		_v0 *unix.TpacketStatsV3
		_v1 error
	)
	for {
		_v0, _v1 = unix.GetsockoptTpacketStatsV3(fd, level, opt)
		if _v1 != syscall.EINTR {
			break
		}
	}
	return _v0, _v1
}

// SetsockoptIPMreqn is an alias of golang.org/x/sys/unix.SetsockoptIPMreqn, wrapped to automatically retry on EINTR.
func SetsockoptIPMreqn(fd, level, opt int, mreq *unix.IPMreqn) (err error) {
	for {
		err = unix.SetsockoptIPMreqn(fd, level, opt, mreq)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// SetsockoptPacketMreq is an alias of golang.org/x/sys/unix.SetsockoptPacketMreq, wrapped to automatically retry on EINTR.
func SetsockoptPacketMreq(fd, level, opt int, mreq *unix.PacketMreq) error {
	var (
		_v0 error
	)
	for {
		_v0 = unix.SetsockoptPacketMreq(fd, level, opt, mreq)
		if _v0 != syscall.EINTR {
			break
		}
	}
	return _v0
}

// SetsockoptSockFprog is an alias of golang.org/x/sys/unix.SetsockoptSockFprog, wrapped to automatically retry on EINTR.
func SetsockoptSockFprog(fd, level, opt int, fprog *unix.SockFprog) error {
	var (
		_v0 error
	)
	for {
		_v0 = unix.SetsockoptSockFprog(fd, level, opt, fprog)
		if _v0 != syscall.EINTR {
			break
		}
	}
	return _v0
}

// SetsockoptCanRawFilter is an alias of golang.org/x/sys/unix.SetsockoptCanRawFilter, wrapped to automatically retry on EINTR.
func SetsockoptCanRawFilter(fd, level, opt int, filter []unix.CanFilter) error {
	var (
		_v0 error
	)
	for {
		_v0 = unix.SetsockoptCanRawFilter(fd, level, opt, filter)
		if _v0 != syscall.EINTR {
			break
		}
	}
	return _v0
}

// SetsockoptTpacketReq is an alias of golang.org/x/sys/unix.SetsockoptTpacketReq, wrapped to automatically retry on EINTR.
func SetsockoptTpacketReq(fd, level, opt int, tp *unix.TpacketReq) error {
	var (
		_v0 error
	)
	for {
		_v0 = unix.SetsockoptTpacketReq(fd, level, opt, tp)
		if _v0 != syscall.EINTR {
			break
		}
	}
	return _v0
}

// SetsockoptTpacketReq3 is an alias of golang.org/x/sys/unix.SetsockoptTpacketReq3, wrapped to automatically retry on EINTR.
func SetsockoptTpacketReq3(fd, level, opt int, tp *unix.TpacketReq3) error {
	var (
		_v0 error
	)
	for {
		_v0 = unix.SetsockoptTpacketReq3(fd, level, opt, tp)
		if _v0 != syscall.EINTR {
			break
		}
	}
	return _v0
}

// KeyctlString is an alias of golang.org/x/sys/unix.KeyctlString, wrapped to automatically retry on EINTR.
func KeyctlString(cmd int, id int) (string, error) {
	var (
		_v0 string
		_v1 error
	)
	for {
		_v0, _v1 = unix.KeyctlString(cmd, id)
		if _v1 != syscall.EINTR {
			break
		}
	}
	return _v0, _v1
}

// KeyctlGetKeyringID is an alias of golang.org/x/sys/unix.KeyctlGetKeyringID, wrapped to automatically retry on EINTR.
func KeyctlGetKeyringID(id int, create bool) (ringid int, err error) {
	for {
		ringid, err = unix.KeyctlGetKeyringID(id, create)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// KeyctlSetperm is an alias of golang.org/x/sys/unix.KeyctlSetperm, wrapped to automatically retry on EINTR.
func KeyctlSetperm(id int, perm uint32) error {
	var (
		_v0 error
	)
	for {
		_v0 = unix.KeyctlSetperm(id, perm)
		if _v0 != syscall.EINTR {
			break
		}
	}
	return _v0
}

// KeyctlJoinSessionKeyring is an alias of golang.org/x/sys/unix.KeyctlJoinSessionKeyring, wrapped to automatically retry on EINTR.
func KeyctlJoinSessionKeyring(name string) (ringid int, err error) {
	for {
		ringid, err = unix.KeyctlJoinSessionKeyring(name)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// KeyctlSearch is an alias of golang.org/x/sys/unix.KeyctlSearch, wrapped to automatically retry on EINTR.
func KeyctlSearch(ringid int, keyType, description string, destRingid int) (id int, err error) {
	for {
		id, err = unix.KeyctlSearch(ringid, keyType, description, destRingid)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// KeyctlInstantiateIOV is an alias of golang.org/x/sys/unix.KeyctlInstantiateIOV, wrapped to automatically retry on EINTR.
func KeyctlInstantiateIOV(id int, payload []unix.Iovec, ringid int) error {
	var (
		_v0 error
	)
	for {
		_v0 = unix.KeyctlInstantiateIOV(id, payload, ringid)
		if _v0 != syscall.EINTR {
			break
		}
	}
	return _v0
}

// KeyctlDHCompute is an alias of golang.org/x/sys/unix.KeyctlDHCompute, wrapped to automatically retry on EINTR.
func KeyctlDHCompute(params *unix.KeyctlDHParams, buffer []byte) (size int, err error) {
	for {
		size, err = unix.KeyctlDHCompute(params, buffer)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// KeyctlRestrictKeyring is an alias of golang.org/x/sys/unix.KeyctlRestrictKeyring, wrapped to automatically retry on EINTR.
func KeyctlRestrictKeyring(ringid int, keyType string, restriction string) error {
	var (
		_v0 error
	)
	for {
		_v0 = unix.KeyctlRestrictKeyring(ringid, keyType, restriction)
		if _v0 != syscall.EINTR {
			break
		}
	}
	return _v0
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

// BindToDevice is an alias of golang.org/x/sys/unix.BindToDevice, wrapped to automatically retry on EINTR.
func BindToDevice(fd int, device string) (err error) {
	for {
		err = unix.BindToDevice(fd, device)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// PtracePeekText is an alias of golang.org/x/sys/unix.PtracePeekText, wrapped to automatically retry on EINTR.
func PtracePeekText(pid int, addr uintptr, out []byte) (count int, err error) {
	for {
		count, err = unix.PtracePeekText(pid, addr, out)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// PtracePeekData is an alias of golang.org/x/sys/unix.PtracePeekData, wrapped to automatically retry on EINTR.
func PtracePeekData(pid int, addr uintptr, out []byte) (count int, err error) {
	for {
		count, err = unix.PtracePeekData(pid, addr, out)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// PtracePeekUser is an alias of golang.org/x/sys/unix.PtracePeekUser, wrapped to automatically retry on EINTR.
func PtracePeekUser(pid int, addr uintptr, out []byte) (count int, err error) {
	for {
		count, err = unix.PtracePeekUser(pid, addr, out)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// PtracePokeText is an alias of golang.org/x/sys/unix.PtracePokeText, wrapped to automatically retry on EINTR.
func PtracePokeText(pid int, addr uintptr, data []byte) (count int, err error) {
	for {
		count, err = unix.PtracePokeText(pid, addr, data)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// PtracePokeData is an alias of golang.org/x/sys/unix.PtracePokeData, wrapped to automatically retry on EINTR.
func PtracePokeData(pid int, addr uintptr, data []byte) (count int, err error) {
	for {
		count, err = unix.PtracePokeData(pid, addr, data)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// PtracePokeUser is an alias of golang.org/x/sys/unix.PtracePokeUser, wrapped to automatically retry on EINTR.
func PtracePokeUser(pid int, addr uintptr, data []byte) (count int, err error) {
	for {
		count, err = unix.PtracePokeUser(pid, addr, data)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// PtraceGetRegs is an alias of golang.org/x/sys/unix.PtraceGetRegs, wrapped to automatically retry on EINTR.
func PtraceGetRegs(pid int, regsout *unix.PtraceRegs) (err error) {
	for {
		err = unix.PtraceGetRegs(pid, regsout)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// PtraceSetRegs is an alias of golang.org/x/sys/unix.PtraceSetRegs, wrapped to automatically retry on EINTR.
func PtraceSetRegs(pid int, regs *unix.PtraceRegs) (err error) {
	for {
		err = unix.PtraceSetRegs(pid, regs)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// PtraceSetOptions is an alias of golang.org/x/sys/unix.PtraceSetOptions, wrapped to automatically retry on EINTR.
func PtraceSetOptions(pid int, options int) (err error) {
	for {
		err = unix.PtraceSetOptions(pid, options)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// PtraceGetEventMsg is an alias of golang.org/x/sys/unix.PtraceGetEventMsg, wrapped to automatically retry on EINTR.
func PtraceGetEventMsg(pid int) (msg uint, err error) {
	for {
		msg, err = unix.PtraceGetEventMsg(pid)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// PtraceCont is an alias of golang.org/x/sys/unix.PtraceCont, wrapped to automatically retry on EINTR.
func PtraceCont(pid int, signal int) (err error) {
	for {
		err = unix.PtraceCont(pid, signal)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// PtraceSyscall is an alias of golang.org/x/sys/unix.PtraceSyscall, wrapped to automatically retry on EINTR.
func PtraceSyscall(pid int, signal int) (err error) {
	for {
		err = unix.PtraceSyscall(pid, signal)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// PtraceSingleStep is an alias of golang.org/x/sys/unix.PtraceSingleStep, wrapped to automatically retry on EINTR.
func PtraceSingleStep(pid int) (err error) {
	for {
		err = unix.PtraceSingleStep(pid)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// PtraceInterrupt is an alias of golang.org/x/sys/unix.PtraceInterrupt, wrapped to automatically retry on EINTR.
func PtraceInterrupt(pid int) (err error) {
	for {
		err = unix.PtraceInterrupt(pid)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// PtraceAttach is an alias of golang.org/x/sys/unix.PtraceAttach, wrapped to automatically retry on EINTR.
func PtraceAttach(pid int) (err error) {
	for {
		err = unix.PtraceAttach(pid)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// PtraceSeize is an alias of golang.org/x/sys/unix.PtraceSeize, wrapped to automatically retry on EINTR.
func PtraceSeize(pid int) (err error) {
	for {
		err = unix.PtraceSeize(pid)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// PtraceDetach is an alias of golang.org/x/sys/unix.PtraceDetach, wrapped to automatically retry on EINTR.
func PtraceDetach(pid int) (err error) {
	for {
		err = unix.PtraceDetach(pid)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Reboot is an alias of golang.org/x/sys/unix.Reboot, wrapped to automatically retry on EINTR.
func Reboot(cmd int) (err error) {
	for {
		err = unix.Reboot(cmd)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Mount is an alias of golang.org/x/sys/unix.Mount, wrapped to automatically retry on EINTR.
func Mount(source string, target string, fstype string, flags uintptr, data string) (err error) {
	for {
		err = unix.Mount(source, target, fstype, flags, data)
		if err != syscall.EINTR {
			break
		}
	}
	return
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

// Getpgrp is an alias of golang.org/x/sys/unix.Getpgrp.
func Getpgrp() (pid int) {
	return unix.Getpgrp()
}

// PrctlRetInt is an alias of golang.org/x/sys/unix.PrctlRetInt, wrapped to automatically retry on EINTR.
func PrctlRetInt(option int, arg2 uintptr, arg3 uintptr, arg4 uintptr, arg5 uintptr) (int, error) {
	var (
		_v0 int
		_v1 error
	)
	for {
		_v0, _v1 = unix.PrctlRetInt(option, arg2, arg3, arg4, arg5)
		if _v1 != syscall.EINTR {
			break
		}
	}
	return _v0, _v1
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

// Setgid is an alias of golang.org/x/sys/unix.Setgid, wrapped to automatically retry on EINTR.
func Setgid(uid int) (err error) {
	for {
		err = unix.Setgid(uid)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// SetfsgidRetGid is an alias of golang.org/x/sys/unix.SetfsgidRetGid, wrapped to automatically retry on EINTR.
func SetfsgidRetGid(gid int) (int, error) {
	var (
		_v0 int
		_v1 error
	)
	for {
		_v0, _v1 = unix.SetfsgidRetGid(gid)
		if _v1 != syscall.EINTR {
			break
		}
	}
	return _v0, _v1
}

// SetfsuidRetUid is an alias of golang.org/x/sys/unix.SetfsuidRetUid, wrapped to automatically retry on EINTR.
func SetfsuidRetUid(uid int) (int, error) {
	var (
		_v0 int
		_v1 error
	)
	for {
		_v0, _v1 = unix.SetfsuidRetUid(uid)
		if _v1 != syscall.EINTR {
			break
		}
	}
	return _v0, _v1
}

// Setfsgid is an alias of golang.org/x/sys/unix.Setfsgid, wrapped to automatically retry on EINTR.
func Setfsgid(gid int) error {
	var (
		_v0 error
	)
	for {
		_v0 = unix.Setfsgid(gid)
		if _v0 != syscall.EINTR {
			break
		}
	}
	return _v0
}

// Setfsuid is an alias of golang.org/x/sys/unix.Setfsuid, wrapped to automatically retry on EINTR.
func Setfsuid(uid int) error {
	var (
		_v0 error
	)
	for {
		_v0 = unix.Setfsuid(uid)
		if _v0 != syscall.EINTR {
			break
		}
	}
	return _v0
}

// Signalfd is an alias of golang.org/x/sys/unix.Signalfd, wrapped to automatically retry on EINTR.
func Signalfd(fd int, sigmask *unix.Sigset_t, flags int) (newfd int, err error) {
	for {
		newfd, err = unix.Signalfd(fd, sigmask, flags)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Readv is an alias of golang.org/x/sys/unix.Readv, wrapped to automatically retry on EINTR.
func Readv(fd int, iovs [][]byte) (n int, err error) {
	for {
		n, err = unix.Readv(fd, iovs)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Preadv is an alias of golang.org/x/sys/unix.Preadv, wrapped to automatically retry on EINTR.
func Preadv(fd int, iovs [][]byte, offset int64) (n int, err error) {
	for {
		n, err = unix.Preadv(fd, iovs, offset)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Preadv2 is an alias of golang.org/x/sys/unix.Preadv2, wrapped to automatically retry on EINTR.
func Preadv2(fd int, iovs [][]byte, offset int64, flags int) (n int, err error) {
	for {
		n, err = unix.Preadv2(fd, iovs, offset, flags)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Writev is an alias of golang.org/x/sys/unix.Writev, wrapped to automatically retry on EINTR.
func Writev(fd int, iovs [][]byte) (n int, err error) {
	for {
		n, err = unix.Writev(fd, iovs)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Pwritev is an alias of golang.org/x/sys/unix.Pwritev, wrapped to automatically retry on EINTR.
func Pwritev(fd int, iovs [][]byte, offset int64) (n int, err error) {
	for {
		n, err = unix.Pwritev(fd, iovs, offset)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Pwritev2 is an alias of golang.org/x/sys/unix.Pwritev2, wrapped to automatically retry on EINTR.
func Pwritev2(fd int, iovs [][]byte, offset int64, flags int) (n int, err error) {
	for {
		n, err = unix.Pwritev2(fd, iovs, offset, flags)
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

// Vmsplice is an alias of golang.org/x/sys/unix.Vmsplice, wrapped to automatically retry on EINTR.
func Vmsplice(fd int, iovs []unix.Iovec, flags int) (int, error) {
	var (
		_v0 int
		_v1 error
	)
	for {
		_v0, _v1 = unix.Vmsplice(fd, iovs, flags)
		if _v1 != syscall.EINTR {
			break
		}
	}
	return _v0, _v1
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

// NewFileHandle is an alias of golang.org/x/sys/unix.NewFileHandle.
func NewFileHandle(handleType int32, handle []byte) unix.FileHandle {
	return unix.NewFileHandle(handleType, handle)
}

// NameToHandleAt is an alias of golang.org/x/sys/unix.NameToHandleAt, wrapped to automatically retry on EINTR.
func NameToHandleAt(dirfd int, path string, flags int) (handle unix.FileHandle, mountID int, err error) {
	for {
		handle, mountID, err = unix.NameToHandleAt(dirfd, path, flags)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// OpenByHandleAt is an alias of golang.org/x/sys/unix.OpenByHandleAt, wrapped to automatically retry on EINTR.
func OpenByHandleAt(mountFD int, handle unix.FileHandle, flags int) (fd int, err error) {
	for {
		fd, err = unix.OpenByHandleAt(mountFD, handle, flags)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Klogset is an alias of golang.org/x/sys/unix.Klogset, wrapped to automatically retry on EINTR.
func Klogset(typ int, arg int) (err error) {
	for {
		err = unix.Klogset(typ, arg)
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

// Pipe2 is an alias of golang.org/x/sys/unix.Pipe2, wrapped to automatically retry on EINTR.
func Pipe2(p []int, flags int) (err error) {
	for {
		err = unix.Pipe2(p, flags)
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

// Time is an alias of golang.org/x/sys/unix.Time, wrapped to automatically retry on EINTR.
func Time(t *unix.Time_t) (unix.Time_t, error) {
	var (
		_v0 unix.Time_t
		_v1 error
	)
	for {
		_v0, _v1 = unix.Time(t)
		if _v1 != syscall.EINTR {
			break
		}
	}
	return _v0, _v1
}

// Utime is an alias of golang.org/x/sys/unix.Utime, wrapped to automatically retry on EINTR.
func Utime(path string, buf *unix.Utimbuf) error {
	var (
		_v0 error
	)
	for {
		_v0 = unix.Utime(path, buf)
		if _v0 != syscall.EINTR {
			break
		}
	}
	return _v0
}

// Fadvise is an alias of golang.org/x/sys/unix.Fadvise, wrapped to automatically retry on EINTR.
func Fadvise(fd int, offset int64, length int64, advice int) (err error) {
	for {
		err = unix.Fadvise(fd, offset, length, advice)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Fstatfs is an alias of golang.org/x/sys/unix.Fstatfs, wrapped to automatically retry on EINTR.
func Fstatfs(fd int, buf *unix.Statfs_t) (err error) {
	for {
		err = unix.Fstatfs(fd, buf)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Statfs is an alias of golang.org/x/sys/unix.Statfs, wrapped to automatically retry on EINTR.
func Statfs(path string, buf *unix.Statfs_t) (err error) {
	for {
		err = unix.Statfs(path, buf)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Getrlimit is an alias of golang.org/x/sys/unix.Getrlimit, wrapped to automatically retry on EINTR.
func Getrlimit(resource int, rlim *unix.Rlimit) (err error) {
	for {
		err = unix.Getrlimit(resource, rlim)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Setrlimit is an alias of golang.org/x/sys/unix.Setrlimit, wrapped to automatically retry on EINTR.
func Setrlimit(resource int, rlim *unix.Rlimit) (err error) {
	for {
		err = unix.Setrlimit(resource, rlim)
		if err != syscall.EINTR {
			break
		}
	}
	return
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

// SyncFileRange is an alias of golang.org/x/sys/unix.SyncFileRange, wrapped to automatically retry on EINTR.
func SyncFileRange(fd int, off int64, n int64, flags int) error {
	var (
		_v0 error
	)
	for {
		_v0 = unix.SyncFileRange(fd, off, n, flags)
		if _v0 != syscall.EINTR {
			break
		}
	}
	return _v0
}

// KexecFileLoad is an alias of golang.org/x/sys/unix.KexecFileLoad, wrapped to automatically retry on EINTR.
func KexecFileLoad(kernelFd int, initrdFd int, cmdline string, flags int) error {
	var (
		_v0 error
	)
	for {
		_v0 = unix.KexecFileLoad(kernelFd, initrdFd, cmdline, flags)
		if _v0 != syscall.EINTR {
			break
		}
	}
	return _v0
}

// SyscallNoError is an alias of golang.org/x/sys/unix.SyscallNoError.
func SyscallNoError(trap, a1, a2, a3 uintptr) (r1, r2 uintptr) {
	return unix.SyscallNoError(trap, a1, a2, a3)
}

// RawSyscallNoError is an alias of golang.org/x/sys/unix.RawSyscallNoError.
func RawSyscallNoError(trap, a1, a2, a3 uintptr) (r1, r2 uintptr) {
	return unix.RawSyscallNoError(trap, a1, a2, a3)
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

// PtraceGetRegsArm is an alias of golang.org/x/sys/unix.PtraceGetRegsArm, wrapped to automatically retry on EINTR.
func PtraceGetRegsArm(pid int, regsout *unix.PtraceRegsArm) error {
	var (
		_v0 error
	)
	for {
		_v0 = unix.PtraceGetRegsArm(pid, regsout)
		if _v0 != syscall.EINTR {
			break
		}
	}
	return _v0
}

// PtraceSetRegsArm is an alias of golang.org/x/sys/unix.PtraceSetRegsArm, wrapped to automatically retry on EINTR.
func PtraceSetRegsArm(pid int, regs *unix.PtraceRegsArm) error {
	var (
		_v0 error
	)
	for {
		_v0 = unix.PtraceSetRegsArm(pid, regs)
		if _v0 != syscall.EINTR {
			break
		}
	}
	return _v0
}

// PtraceGetRegsArm64 is an alias of golang.org/x/sys/unix.PtraceGetRegsArm64, wrapped to automatically retry on EINTR.
func PtraceGetRegsArm64(pid int, regsout *unix.PtraceRegsArm64) error {
	var (
		_v0 error
	)
	for {
		_v0 = unix.PtraceGetRegsArm64(pid, regsout)
		if _v0 != syscall.EINTR {
			break
		}
	}
	return _v0
}

// PtraceSetRegsArm64 is an alias of golang.org/x/sys/unix.PtraceSetRegsArm64, wrapped to automatically retry on EINTR.
func PtraceSetRegsArm64(pid int, regs *unix.PtraceRegsArm64) error {
	var (
		_v0 error
	)
	for {
		_v0 = unix.PtraceSetRegsArm64(pid, regs)
		if _v0 != syscall.EINTR {
			break
		}
	}
	return _v0
}

// FanotifyInit is an alias of golang.org/x/sys/unix.FanotifyInit, wrapped to automatically retry on EINTR.
func FanotifyInit(flags uint, event_f_flags uint) (fd int, err error) {
	for {
		fd, err = unix.FanotifyInit(flags, event_f_flags)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Linkat is an alias of golang.org/x/sys/unix.Linkat, wrapped to automatically retry on EINTR.
func Linkat(olddirfd int, oldpath string, newdirfd int, newpath string, flags int) (err error) {
	for {
		err = unix.Linkat(olddirfd, oldpath, newdirfd, newpath, flags)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Readlinkat is an alias of golang.org/x/sys/unix.Readlinkat, wrapped to automatically retry on EINTR.
func Readlinkat(dirfd int, path string, buf []byte) (n int, err error) {
	for {
		n, err = unix.Readlinkat(dirfd, path, buf)
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

// KeyctlInt is an alias of golang.org/x/sys/unix.KeyctlInt, wrapped to automatically retry on EINTR.
func KeyctlInt(cmd int, arg2 int, arg3 int, arg4 int, arg5 int) (ret int, err error) {
	for {
		ret, err = unix.KeyctlInt(cmd, arg2, arg3, arg4, arg5)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// KeyctlBuffer is an alias of golang.org/x/sys/unix.KeyctlBuffer, wrapped to automatically retry on EINTR.
func KeyctlBuffer(cmd int, arg2 int, buf []byte, arg5 int) (ret int, err error) {
	for {
		ret, err = unix.KeyctlBuffer(cmd, arg2, buf, arg5)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Acct is an alias of golang.org/x/sys/unix.Acct, wrapped to automatically retry on EINTR.
func Acct(path string) (err error) {
	for {
		err = unix.Acct(path)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// AddKey is an alias of golang.org/x/sys/unix.AddKey, wrapped to automatically retry on EINTR.
func AddKey(keyType string, description string, payload []byte, ringid int) (id int, err error) {
	for {
		id, err = unix.AddKey(keyType, description, payload, ringid)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Adjtimex is an alias of golang.org/x/sys/unix.Adjtimex, wrapped to automatically retry on EINTR.
func Adjtimex(buf *unix.Timex) (state int, err error) {
	for {
		state, err = unix.Adjtimex(buf)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Capget is an alias of golang.org/x/sys/unix.Capget, wrapped to automatically retry on EINTR.
func Capget(hdr *unix.CapUserHeader, data *unix.CapUserData) (err error) {
	for {
		err = unix.Capget(hdr, data)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Capset is an alias of golang.org/x/sys/unix.Capset, wrapped to automatically retry on EINTR.
func Capset(hdr *unix.CapUserHeader, data *unix.CapUserData) (err error) {
	for {
		err = unix.Capset(hdr, data)
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

// ClockGetres is an alias of golang.org/x/sys/unix.ClockGetres, wrapped to automatically retry on EINTR.
func ClockGetres(clockid int32, res *unix.Timespec) (err error) {
	for {
		err = unix.ClockGetres(clockid, res)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// ClockGettime is an alias of golang.org/x/sys/unix.ClockGettime, wrapped to automatically retry on EINTR.
func ClockGettime(clockid int32, time *unix.Timespec) (err error) {
	for {
		err = unix.ClockGettime(clockid, time)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// ClockNanosleep is an alias of golang.org/x/sys/unix.ClockNanosleep, wrapped to automatically retry on EINTR.
func ClockNanosleep(clockid int32, flags int, request *unix.Timespec, remain *unix.Timespec) (err error) {
	for {
		err = unix.ClockNanosleep(clockid, flags, request, remain)
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

// CopyFileRange is an alias of golang.org/x/sys/unix.CopyFileRange, wrapped to automatically retry on EINTR.
func CopyFileRange(rfd int, roff *int64, wfd int, woff *int64, len int, flags int) (n int, err error) {
	for {
		n, err = unix.CopyFileRange(rfd, roff, wfd, woff, len, flags)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// DeleteModule is an alias of golang.org/x/sys/unix.DeleteModule, wrapped to automatically retry on EINTR.
func DeleteModule(name string, flags int) (err error) {
	for {
		err = unix.DeleteModule(name, flags)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Dup is an alias of golang.org/x/sys/unix.Dup, wrapped to automatically retry on EINTR.
func Dup(oldfd int) (fd int, err error) {
	for {
		fd, err = unix.Dup(oldfd)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Dup3 is an alias of golang.org/x/sys/unix.Dup3, wrapped to automatically retry on EINTR.
func Dup3(oldfd int, newfd int, flags int) (err error) {
	for {
		err = unix.Dup3(oldfd, newfd, flags)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// EpollCreate1 is an alias of golang.org/x/sys/unix.EpollCreate1, wrapped to automatically retry on EINTR.
func EpollCreate1(flag int) (fd int, err error) {
	for {
		fd, err = unix.EpollCreate1(flag)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// EpollCtl is an alias of golang.org/x/sys/unix.EpollCtl, wrapped to automatically retry on EINTR.
func EpollCtl(epfd int, op int, fd int, event *unix.EpollEvent) (err error) {
	for {
		err = unix.EpollCtl(epfd, op, fd, event)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Eventfd is an alias of golang.org/x/sys/unix.Eventfd, wrapped to automatically retry on EINTR.
func Eventfd(initval uint, flags int) (fd int, err error) {
	for {
		fd, err = unix.Eventfd(initval, flags)
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

// Fallocate is an alias of golang.org/x/sys/unix.Fallocate, wrapped to automatically retry on EINTR.
func Fallocate(fd int, mode uint32, off int64, len int64) (err error) {
	for {
		err = unix.Fallocate(fd, mode, off, len)
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

// Fdatasync is an alias of golang.org/x/sys/unix.Fdatasync, wrapped to automatically retry on EINTR.
func Fdatasync(fd int) (err error) {
	for {
		err = unix.Fdatasync(fd)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Fgetxattr is an alias of golang.org/x/sys/unix.Fgetxattr, wrapped to automatically retry on EINTR.
func Fgetxattr(fd int, attr string, dest []byte) (sz int, err error) {
	for {
		sz, err = unix.Fgetxattr(fd, attr, dest)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// FinitModule is an alias of golang.org/x/sys/unix.FinitModule, wrapped to automatically retry on EINTR.
func FinitModule(fd int, params string, flags int) (err error) {
	for {
		err = unix.FinitModule(fd, params, flags)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Flistxattr is an alias of golang.org/x/sys/unix.Flistxattr, wrapped to automatically retry on EINTR.
func Flistxattr(fd int, dest []byte) (sz int, err error) {
	for {
		sz, err = unix.Flistxattr(fd, dest)
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

// Fremovexattr is an alias of golang.org/x/sys/unix.Fremovexattr, wrapped to automatically retry on EINTR.
func Fremovexattr(fd int, attr string) (err error) {
	for {
		err = unix.Fremovexattr(fd, attr)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Fsetxattr is an alias of golang.org/x/sys/unix.Fsetxattr, wrapped to automatically retry on EINTR.
func Fsetxattr(fd int, attr string, dest []byte, flags int) (err error) {
	for {
		err = unix.Fsetxattr(fd, attr, dest, flags)
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

// Getrandom is an alias of golang.org/x/sys/unix.Getrandom, wrapped to automatically retry on EINTR.
func Getrandom(buf []byte, flags int) (n int, err error) {
	for {
		n, err = unix.Getrandom(buf, flags)
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

// Gettid is an alias of golang.org/x/sys/unix.Gettid.
func Gettid() (tid int) {
	return unix.Gettid()
}

// Getxattr is an alias of golang.org/x/sys/unix.Getxattr, wrapped to automatically retry on EINTR.
func Getxattr(path string, attr string, dest []byte) (sz int, err error) {
	for {
		sz, err = unix.Getxattr(path, attr, dest)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// InitModule is an alias of golang.org/x/sys/unix.InitModule, wrapped to automatically retry on EINTR.
func InitModule(moduleImage []byte, params string) (err error) {
	for {
		err = unix.InitModule(moduleImage, params)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// InotifyAddWatch is an alias of golang.org/x/sys/unix.InotifyAddWatch, wrapped to automatically retry on EINTR.
func InotifyAddWatch(fd int, pathname string, mask uint32) (watchdesc int, err error) {
	for {
		watchdesc, err = unix.InotifyAddWatch(fd, pathname, mask)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// InotifyInit1 is an alias of golang.org/x/sys/unix.InotifyInit1, wrapped to automatically retry on EINTR.
func InotifyInit1(flags int) (fd int, err error) {
	for {
		fd, err = unix.InotifyInit1(flags)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// InotifyRmWatch is an alias of golang.org/x/sys/unix.InotifyRmWatch, wrapped to automatically retry on EINTR.
func InotifyRmWatch(fd int, watchdesc uint32) (success int, err error) {
	for {
		success, err = unix.InotifyRmWatch(fd, watchdesc)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Kill is an alias of golang.org/x/sys/unix.Kill, wrapped to automatically retry on EINTR.
func Kill(pid int, sig syscall.Signal) (err error) {
	for {
		err = unix.Kill(pid, sig)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Klogctl is an alias of golang.org/x/sys/unix.Klogctl, wrapped to automatically retry on EINTR.
func Klogctl(typ int, buf []byte) (n int, err error) {
	for {
		n, err = unix.Klogctl(typ, buf)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Lgetxattr is an alias of golang.org/x/sys/unix.Lgetxattr, wrapped to automatically retry on EINTR.
func Lgetxattr(path string, attr string, dest []byte) (sz int, err error) {
	for {
		sz, err = unix.Lgetxattr(path, attr, dest)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Listxattr is an alias of golang.org/x/sys/unix.Listxattr, wrapped to automatically retry on EINTR.
func Listxattr(path string, dest []byte) (sz int, err error) {
	for {
		sz, err = unix.Listxattr(path, dest)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Llistxattr is an alias of golang.org/x/sys/unix.Llistxattr, wrapped to automatically retry on EINTR.
func Llistxattr(path string, dest []byte) (sz int, err error) {
	for {
		sz, err = unix.Llistxattr(path, dest)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Lremovexattr is an alias of golang.org/x/sys/unix.Lremovexattr, wrapped to automatically retry on EINTR.
func Lremovexattr(path string, attr string) (err error) {
	for {
		err = unix.Lremovexattr(path, attr)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Lsetxattr is an alias of golang.org/x/sys/unix.Lsetxattr, wrapped to automatically retry on EINTR.
func Lsetxattr(path string, attr string, data []byte, flags int) (err error) {
	for {
		err = unix.Lsetxattr(path, attr, data, flags)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// MemfdCreate is an alias of golang.org/x/sys/unix.MemfdCreate, wrapped to automatically retry on EINTR.
func MemfdCreate(name string, flags int) (fd int, err error) {
	for {
		fd, err = unix.MemfdCreate(name, flags)
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

// Mknodat is an alias of golang.org/x/sys/unix.Mknodat, wrapped to automatically retry on EINTR.
func Mknodat(dirfd int, path string, mode uint32, dev int) (err error) {
	for {
		err = unix.Mknodat(dirfd, path, mode, dev)
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

// PerfEventOpen is an alias of golang.org/x/sys/unix.PerfEventOpen, wrapped to automatically retry on EINTR.
func PerfEventOpen(attr *unix.PerfEventAttr, pid int, cpu int, groupFd int, flags int) (fd int, err error) {
	for {
		fd, err = unix.PerfEventOpen(attr, pid, cpu, groupFd, flags)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// PivotRoot is an alias of golang.org/x/sys/unix.PivotRoot, wrapped to automatically retry on EINTR.
func PivotRoot(newroot string, putold string) (err error) {
	for {
		err = unix.PivotRoot(newroot, putold)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Prctl is an alias of golang.org/x/sys/unix.Prctl, wrapped to automatically retry on EINTR.
func Prctl(option int, arg2 uintptr, arg3 uintptr, arg4 uintptr, arg5 uintptr) (err error) {
	for {
		err = unix.Prctl(option, arg2, arg3, arg4, arg5)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Pselect is an alias of golang.org/x/sys/unix.Pselect, wrapped to automatically retry on EINTR.
func Pselect(nfd int, r *unix.FdSet, w *unix.FdSet, e *unix.FdSet, timeout *unix.Timespec, sigmask *unix.Sigset_t) (n int, err error) {
	for {
		n, err = unix.Pselect(nfd, r, w, e, timeout, sigmask)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Removexattr is an alias of golang.org/x/sys/unix.Removexattr, wrapped to automatically retry on EINTR.
func Removexattr(path string, attr string) (err error) {
	for {
		err = unix.Removexattr(path, attr)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Renameat2 is an alias of golang.org/x/sys/unix.Renameat2, wrapped to automatically retry on EINTR.
func Renameat2(olddirfd int, oldpath string, newdirfd int, newpath string, flags uint) (err error) {
	for {
		err = unix.Renameat2(olddirfd, oldpath, newdirfd, newpath, flags)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// RequestKey is an alias of golang.org/x/sys/unix.RequestKey, wrapped to automatically retry on EINTR.
func RequestKey(keyType string, description string, callback string, destRingid int) (id int, err error) {
	for {
		id, err = unix.RequestKey(keyType, description, callback, destRingid)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Setdomainname is an alias of golang.org/x/sys/unix.Setdomainname, wrapped to automatically retry on EINTR.
func Setdomainname(p []byte) (err error) {
	for {
		err = unix.Setdomainname(p)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Sethostname is an alias of golang.org/x/sys/unix.Sethostname, wrapped to automatically retry on EINTR.
func Sethostname(p []byte) (err error) {
	for {
		err = unix.Sethostname(p)
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
func Settimeofday(tv *unix.Timeval) (err error) {
	for {
		err = unix.Settimeofday(tv)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Setns is an alias of golang.org/x/sys/unix.Setns, wrapped to automatically retry on EINTR.
func Setns(fd int, nstype int) (err error) {
	for {
		err = unix.Setns(fd, nstype)
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

// Setxattr is an alias of golang.org/x/sys/unix.Setxattr, wrapped to automatically retry on EINTR.
func Setxattr(path string, attr string, data []byte, flags int) (err error) {
	for {
		err = unix.Setxattr(path, attr, data, flags)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Statx is an alias of golang.org/x/sys/unix.Statx, wrapped to automatically retry on EINTR.
func Statx(dirfd int, path string, flags int, mask int, stat *unix.Statx_t) (err error) {
	for {
		err = unix.Statx(dirfd, path, flags, mask, stat)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Sync is an alias of golang.org/x/sys/unix.Sync.
func Sync() {
	unix.Sync()
}

// Syncfs is an alias of golang.org/x/sys/unix.Syncfs, wrapped to automatically retry on EINTR.
func Syncfs(fd int) (err error) {
	for {
		err = unix.Syncfs(fd)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Sysinfo is an alias of golang.org/x/sys/unix.Sysinfo, wrapped to automatically retry on EINTR.
func Sysinfo(info *unix.Sysinfo_t) (err error) {
	for {
		err = unix.Sysinfo(info)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Tee is an alias of golang.org/x/sys/unix.Tee, wrapped to automatically retry on EINTR.
func Tee(rfd int, wfd int, len int, flags int) (n int64, err error) {
	for {
		n, err = unix.Tee(rfd, wfd, len, flags)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Tgkill is an alias of golang.org/x/sys/unix.Tgkill, wrapped to automatically retry on EINTR.
func Tgkill(tgid int, tid int, sig syscall.Signal) (err error) {
	for {
		err = unix.Tgkill(tgid, tid, sig)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Times is an alias of golang.org/x/sys/unix.Times, wrapped to automatically retry on EINTR.
func Times(tms *unix.Tms) (ticks uintptr, err error) {
	for {
		ticks, err = unix.Times(tms)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Umask is an alias of golang.org/x/sys/unix.Umask.
func Umask(mask int) (oldmask int) {
	return unix.Umask(mask)
}

// Uname is an alias of golang.org/x/sys/unix.Uname, wrapped to automatically retry on EINTR.
func Uname(buf *unix.Utsname) (err error) {
	for {
		err = unix.Uname(buf)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Unmount is an alias of golang.org/x/sys/unix.Unmount, wrapped to automatically retry on EINTR.
func Unmount(target string, flags int) (err error) {
	for {
		err = unix.Unmount(target, flags)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Unshare is an alias of golang.org/x/sys/unix.Unshare, wrapped to automatically retry on EINTR.
func Unshare(flags int) (err error) {
	for {
		err = unix.Unshare(flags)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Madvise is an alias of golang.org/x/sys/unix.Madvise, wrapped to automatically retry on EINTR.
func Madvise(b []byte, advice int) (err error) {
	for {
		err = unix.Madvise(b, advice)
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

// Dup2 is an alias of golang.org/x/sys/unix.Dup2, wrapped to automatically retry on EINTR.
func Dup2(oldfd int, newfd int) (err error) {
	for {
		err = unix.Dup2(oldfd, newfd)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// EpollCreate is an alias of golang.org/x/sys/unix.EpollCreate, wrapped to automatically retry on EINTR.
func EpollCreate(size int) (fd int, err error) {
	for {
		fd, err = unix.EpollCreate(size)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// EpollWait is an alias of golang.org/x/sys/unix.EpollWait, wrapped to automatically retry on EINTR.
func EpollWait(epfd int, events []unix.EpollEvent, msec int) (n int, err error) {
	for {
		n, err = unix.EpollWait(epfd, events, msec)
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
func Fstatat(dirfd int, path string, stat *unix.Stat_t, flags int) (err error) {
	for {
		err = unix.Fstatat(dirfd, path, stat, flags)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Getegid is an alias of golang.org/x/sys/unix.Getegid.
func Getegid() (egid int) {
	return unix.Getegid()
}

// Geteuid is an alias of golang.org/x/sys/unix.Geteuid.
func Geteuid() (euid int) {
	return unix.Geteuid()
}

// Getgid is an alias of golang.org/x/sys/unix.Getgid.
func Getgid() (gid int) {
	return unix.Getgid()
}

// Getuid is an alias of golang.org/x/sys/unix.Getuid.
func Getuid() (uid int) {
	return unix.Getuid()
}

// InotifyInit is an alias of golang.org/x/sys/unix.InotifyInit, wrapped to automatically retry on EINTR.
func InotifyInit() (fd int, err error) {
	for {
		fd, err = unix.InotifyInit()
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

// Listen is an alias of golang.org/x/sys/unix.Listen, wrapped to automatically retry on EINTR.
func Listen(s int, n int) (err error) {
	for {
		err = unix.Listen(s, n)
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

// Pause is an alias of golang.org/x/sys/unix.Pause, wrapped to automatically retry on EINTR.
func Pause() (err error) {
	for {
		err = unix.Pause()
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Renameat is an alias of golang.org/x/sys/unix.Renameat, wrapped to automatically retry on EINTR.
func Renameat(olddirfd int, oldpath string, newdirfd int, newpath string) (err error) {
	for {
		err = unix.Renameat(olddirfd, oldpath, newdirfd, newpath)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Select is an alias of golang.org/x/sys/unix.Select, wrapped to automatically retry on EINTR.
func Select(nfd int, r *unix.FdSet, w *unix.FdSet, e *unix.FdSet, timeout *unix.Timeval) (n int, err error) {
	for {
		n, err = unix.Select(nfd, r, w, e, timeout)
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

// Shutdown is an alias of golang.org/x/sys/unix.Shutdown, wrapped to automatically retry on EINTR.
func Shutdown(fd int, how int) (err error) {
	for {
		err = unix.Shutdown(fd, how)
		if err != syscall.EINTR {
			break
		}
	}
	return
}

// Splice is an alias of golang.org/x/sys/unix.Splice, wrapped to automatically retry on EINTR.
func Splice(rfd int, roff *int64, wfd int, woff *int64, len int, flags int) (n int, err error) {
	for {
		n, err = unix.Splice(rfd, roff, wfd, woff, len, flags)
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

// Ustat is an alias of golang.org/x/sys/unix.Ustat, wrapped to automatically retry on EINTR.
func Ustat(dev int, ubuf *unix.Ustat_t) (err error) {
	for {
		err = unix.Ustat(dev, ubuf)
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
