package KeyboardManager

import (
	"syscall"
	"time"
	"unsafe"
)

var (
	user32           = syscall.NewLazyDLL("user32.dll")
	GetAsyncKeyState = user32.NewProc("GetAsyncKeyState")
	sendInputProc    = user32.NewProc("SendInput")
)

type keyboardInput struct {
	wVk         uint16
	wScan       uint16
	dwFlags     uint32
	time        uint32
	dwExtraInfo uint64
}

type input struct {
	itype   uint32
	ki      keyboardInput
	padding uint64
}

func PressKey(message string) {
	if message == "Presiona A" {

		var i input
		i.itype = 1     //INPUT_KEYBOARD
		i.ki.wVk = 0x41 // virtual key code for a
		sendInputProc.Call(
			uintptr(1),
			uintptr(unsafe.Pointer(&i)),
			uintptr(unsafe.Sizeof(i)))

		time.Sleep(time.Second * 2)
	}
}
