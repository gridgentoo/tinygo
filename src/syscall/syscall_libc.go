// +build darwin nintendoswitch wasi

package syscall

import (
	"unsafe"
)

type sliceHeader struct {
	buf *byte
	len uintptr
	cap uintptr
}

func Close(fd int) (err error) {
	if libc_close(int32(fd)) < 0 {
		err = getErrno()
	}
	return
}

func Write(fd int, p []byte) (n int, err error) {
	buf, count := splitSlice(p)
	n = libc_write(int32(fd), buf, uint(count))
	if n < 0 {
		err = getErrno()
	}
	return
}

func Read(fd int, p []byte) (n int, err error) {
	buf, count := splitSlice(p)
	n = libc_read(int32(fd), buf, uint(count))
	if n < 0 {
		err = getErrno()
	}
	return
}

func Seek(fd int, offset int64, whence int) (off int64, err error) {
	return 0, ENOSYS // TODO
}

func Open(path string, flag int, mode uint32) (fd int, err error) {
	data := append([]byte(path), 0)
	fd = int(libc_open(&data[0], int32(flag), mode))
	if fd < 0 {
		err = getErrno()
	}
	return
}

func Mkdir(path string, mode uint32) (err error) {
	return ENOSYS // TODO
}

func Unlink(path string) (err error) {
	return ENOSYS // TODO
}

func Kill(pid int, sig Signal) (err error) {
	return ENOSYS // TODO
}

type SysProcAttr struct{}

func Pipe2(p []int, flags int) (err error) {
	return ENOSYS // TODO
}

func Getenv(key string) (value string, found bool) {
	data := append([]byte(key), 0)
	raw := libc_getenv(&data[0])
	if raw == nil {
		return "", false
	}

	ptr := uintptr(unsafe.Pointer(raw))
	for size := uintptr(0); ; size++ {
		v := *(*byte)(unsafe.Pointer(ptr))
		if v == 0 {
			src := *(*[]byte)(unsafe.Pointer(&sliceHeader{buf: raw, len: size, cap: size}))
			return string(src), true
		}
		ptr += unsafe.Sizeof(byte(0))
	}
}

func Environ() []string {
	environ := libc_environ
	var envs []string
	for *environ != nil {
		// Convert the C string to a Go string.
		length := libc_strlen(*environ)
		var envVar string
		rawEnvVar := (*struct {
			ptr    unsafe.Pointer
			length uintptr
		})(unsafe.Pointer(&envVar))
		rawEnvVar.ptr = *environ
		rawEnvVar.length = length
		envs = append(envs, envVar)
		// This is the Go equivalent of "environ++" in C.
		environ = (*unsafe.Pointer)(unsafe.Pointer(uintptr(unsafe.Pointer(environ)) + unsafe.Sizeof(environ)))
	}
	return envs
}

func splitSlice(p []byte) (buf *byte, len uintptr) {
	slice := (*sliceHeader)(unsafe.Pointer(&p))
	return slice.buf, slice.len
}

//export strlen
func libc_strlen(ptr unsafe.Pointer) uintptr

// ssize_t write(int fd, const void *buf, size_t count)
//export write
func libc_write(fd int32, buf *byte, count uint) int

// char *getenv(const char *name);
//export getenv
func libc_getenv(name *byte) *byte

// ssize_t read(int fd, void *buf, size_t count);
//export read
func libc_read(fd int32, buf *byte, count uint) int

// int open(const char *pathname, int flags, mode_t mode);
//export open
func libc_open(pathname *byte, flags int32, mode uint32) int32

// int close(int fd)
//export close
func libc_close(fd int32) int32

//go:extern environ
var libc_environ *unsafe.Pointer
