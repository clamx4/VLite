// +build linux

package water

import (
	"log"
	"os"
	"runtime"
	"strings"
	"syscall"
	"unsafe"

)

const (
	cIFF_TUN         = 0x0001
	cIFF_TAP         = 0x0002
	cIFF_NO_PI       = 0x1000
	cIFF_MULTI_QUEUE = 0x0100
)

type ifReq struct {
	Name  [0x10]byte
	Flags uint16
	pad   [0x28 - 0x10 - 2]byte
}

const (
	DevNetTun = "/dev/net/tun"
	DevTun    = "/dev/tun"
)

var tunFile string

func init() {
	if _, err := os.Stat(DevNetTun); err == nil {
		tunFile = DevNetTun
	} else if _, err := os.Stat(DevTun); err == nil {
		tunFile = DevTun
	} else {
		log.Fatal("NO support for", runtime.GOOS)
	}
}

func ioctl(fd uintptr, request uintptr, argp uintptr) error {
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, fd, uintptr(request), argp)
	if errno != 0 {
		return os.NewSyscallError("ioctl", errno)
	}
	return nil
}

func newTAP(config Config) (ifce *Interface, err error) {
	filefd, err := syscall.Open(tunFile, os.O_RDWR, 0)

	if err != nil {
		return nil, err
	}

	var flags uint16
	flags = cIFF_TAP | cIFF_NO_PI
	if config.PlatformSpecificParams.MultiQueue {
		flags |= cIFF_MULTI_QUEUE
	}
	name, err := createInterface(uintptr(filefd), config.Name, flags)
	if err != nil {
		return nil, err
	}

	if err = setDeviceOptions(uintptr(filefd), config); err != nil {
		return nil, err
	}

	syscall.SetNonblock(filefd,true)

	file := os.NewFile(uintptr(filefd),tunFile)


	ifce = &Interface{isTAP: true, ReadWriteCloser: file, name: name}
	ifce.fd = int(file.Fd())
	return ifce, nil
}

func newTUN(config Config) (ifce *Interface, err error) {
	filefd, err := syscall.Open(tunFile, os.O_RDWR, 0)
	if err != nil {
		return nil, err
	}

	var flags uint16
	flags = cIFF_TUN | cIFF_NO_PI
	if config.PlatformSpecificParams.MultiQueue {
		flags |= cIFF_MULTI_QUEUE
	}
	name, err := createInterface(uintptr(filefd), config.Name, flags)
	if err != nil {
		return nil, err
	}

	if err = setDeviceOptions(uintptr(filefd), config); err != nil {
		return nil, err
	}

	syscall.SetNonblock(filefd,true)

	file := os.NewFile(uintptr(filefd),tunFile)

	ifce = &Interface{isTAP: false, ReadWriteCloser: file, name: name}
	ifce.fd = int(file.Fd())
	return ifce, nil
}

func createInterface(fd uintptr, ifName string, flags uint16) (createdIFName string, err error) {
	var req ifReq
	req.Flags = flags
	copy(req.Name[:], ifName)

	err = ioctl(fd, syscall.TUNSETIFF, uintptr(unsafe.Pointer(&req)))
	if err != nil {
		return
	}

	createdIFName = strings.Trim(string(req.Name[:]), "\x00")
	return
}

func setDeviceOptions(fd uintptr, config Config) (err error) {

	// Device Permissions
	if config.Permissions != nil {

		// Set Owner
		if err = ioctl(fd, syscall.TUNSETOWNER, uintptr(config.Permissions.Owner)); err != nil {
			return
		}

		// Set Group
		if err = ioctl(fd, syscall.TUNSETGROUP, uintptr(config.Permissions.Group)); err != nil {
			return
		}
	}

	// Set/Clear Persist Device Flag
	value := 0
	if config.Persist {
		value = 1
	}
	return ioctl(fd, syscall.TUNSETPERSIST, uintptr(value))

}
