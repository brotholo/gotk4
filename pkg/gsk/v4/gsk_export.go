// Code generated by girgen. DO NOT EDIT.

package gsk

import (
	"unsafe"

	"github.com/brotholo/gotk4/pkg/core/gbox"
	"github.com/brotholo/gotk4/pkg/core/gerror"
	"github.com/brotholo/gotk4/pkg/core/gextras"
)

// #include <stdlib.h>
// #include <gsk/gsk.h>
import "C"

//export _gotk4_gsk4_ParseErrorFunc
func _gotk4_gsk4_ParseErrorFunc(arg1 *C.GskParseLocation, arg2 *C.GskParseLocation, arg3 *C.GError, arg4 C.gpointer) {
	var fn ParseErrorFunc
	{
		v := gbox.Get(uintptr(arg4))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(ParseErrorFunc)
	}

	var _start *ParseLocation // out
	var _end *ParseLocation   // out
	var _err error            // out

	_start = (*ParseLocation)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	_end = (*ParseLocation)(gextras.NewStructNative(unsafe.Pointer(arg2)))
	_err = gerror.Take(unsafe.Pointer(arg3))

	fn(_start, _end, _err)
}
