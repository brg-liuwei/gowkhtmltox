package gowkhtmltoimage

/*
#include "wkhtmltox/image.h"
#include <stdio.h>
#include <stdlib.h>

#cgo CFLAGS: -I/usr/local/include
#cgo LDFLAGS: -L/usr/local/lib -lwkhtmltox

*/
import "C"

import (
	"errors"
	"runtime"
	"unsafe"
)

func Init(use_graphics bool) {
	var ok C.int
	if use_graphics {
		ok = C.wkhtmltoimage_init(C.int(1))
	} else {
		ok = C.wkhtmltoimage_init(C.int(0))
	}
	if ok != C.int(1) {
		panic("wkhtmltoimage init fail")
	}
}

func DeInit() {
	C.wkhtmltoimage_deinit()
}

type Convertor struct {
	ready       bool
	html        string
	wkSettings  *C.wkhtmltoimage_global_settings
	wkConvertor *C.wkhtmltoimage_converter
	img         string
}

func NewConvertor() *Convertor {
	settings := C.wkhtmltoimage_create_global_settings()
	if unsafe.Pointer(settings) == unsafe.Pointer(nil) {
		return nil
	}
	convertor := &Convertor{
		ready:       false,
		wkSettings:  settings,
		wkConvertor: nil,
	}

	// XXX: Warning!
	// Memory leak would be cause if func Ready being not called
	runtime.SetFinalizer(convertor, func(c *Convertor) {
		if c.wkConvertor != nil {
			C.wkhtmltoimage_destroy_converter(c.wkConvertor)
		}
	})
	return convertor
}

func (conv *Convertor) SetProperty(key, value string) {
	k, v := C.CString(key), C.CString(value)
	defer func() {
		C.free(unsafe.Pointer(k))
		C.free(unsafe.Pointer(v))
	}()
	C.wkhtmltoimage_set_global_setting(conv.wkSettings, k, v)
}

func (conv *Convertor) SetHtml(html string) {
	conv.html = html
}

func (conv *Convertor) Ready() error {
	if conv.ready {
		return nil
	}
	conv.ready = true
	if len(conv.html) == 0 {
		conv.wkConvertor = C.wkhtmltoimage_create_converter(
			conv.wkSettings, (*C.char)(unsafe.Pointer(nil)))
	} else {
		data := C.CString(conv.html)
		defer C.free(unsafe.Pointer(data))
		conv.wkConvertor = C.wkhtmltoimage_create_converter(
			conv.wkSettings, data)
	}
	if unsafe.Pointer(conv.wkConvertor) == unsafe.Pointer(nil) {
		return errors.New("wkhtmltoimage_create_converter error")
	}
	return nil
}

func (conv *Convertor) Run() error {
	if !conv.ready {
		return errors.New("this convertor has not ready")
	}
	rc := C.wkhtmltoimage_convert(conv.wkConvertor)
	if int(rc) != 1 {
		// get error
		return errors.New("some error happened")
	}

	var img *C.uchar
	size := C.wkhtmltoimage_get_output(conv.wkConvertor, &img)
	conv.img = C.GoStringN((*C.char)(unsafe.Pointer(img)), C.int(size))
	return nil
}

func (conv *Convertor) GetImage() []byte {
	return []byte(conv.img)
}
