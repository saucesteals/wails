package edge

import (
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)

type ICoreWebView2EnvironmentOptionsVtbl struct {
	_IUnknownVtbl
	GetAdditionalBrowserArguments             ComProc
	PutAdditionalBrowserArguments             ComProc
	GetAllowSingleSignOnUsingOSPrimaryAccount ComProc
	PutAllowSingleSignOnUsingOSPrimaryAccount ComProc
	GetLanguage                               ComProc
	PutLanguage                               ComProc
	GetTargetCompatibleBrowserVersion         ComProc
	PutTargetCompatibleBrowserVersion         ComProc
}

type ICoreWebView2EnvironmentOptions struct {
	vtbl *ICoreWebView2EnvironmentOptionsVtbl
}

func (i *ICoreWebView2EnvironmentOptions) AddRef() uintptr {
	return i.AddRef()
}

func (i *ICoreWebView2EnvironmentOptions) PutAdditionalBrowserArguments(arguments string) error {
	var err error

	// Cast to a uint16 as that's what the call is expecting
	col := *(*uint16)(unsafe.Pointer(&arguments))

	_, _, err = i.vtbl.PutAdditionalBrowserArguments.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(col),
	)

	if err != windows.ERROR_SUCCESS {
		return err
	}
	return nil
}

func (i *ICoreWebView2EnvironmentOptions) GetAdditionalBrowserArguments() (string, error) {
	// Create *uint16 to hold result
	var _arguments *uint16

	res, _, err := i.vtbl.GetAdditionalBrowserArguments.Call(
		uintptr(unsafe.Pointer(i)),
		uintptr(unsafe.Pointer(&_arguments)),
	)

	if err != windows.ERROR_SUCCESS {
		return "", err
	}
	if windows.Handle(res) != windows.S_OK {
		return "", syscall.Errno(res)
	}

	// Get result and cleanup
	arguments := windows.UTF16PtrToString(_arguments)
	windows.CoTaskMemFree(unsafe.Pointer(_arguments))
	return arguments, nil
}
