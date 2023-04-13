// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package win

import (
	"unsafe"
)

type SCODE int32

type EXCEPINFO struct {
	wCode             uint16
	wReserved         uint16
	bstrSource        *uint16 /*BSTR*/
	bstrDescription   *uint16 /*BSTR*/
	bstrHelpFile      *uint16 /*BSTR*/
	dwHelpContext     uint32
	pvReserved        uintptr
	pfnDeferredFillIn uintptr
	scode             SCODE
}

var (
	IID_ITypeInfo = IID{0x00020401, 0x0000, 0x0000, [8]byte{0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}
)

type ITypeInfoVtbl struct {
	IUnknownVtbl
	GetTypeAttr          uintptr
	GetTypeComp          uintptr
	GetFuncDesc          uintptr
	GetVarDesc           uintptr
	GetNames             uintptr
	GetRefTypeOfImplType uintptr
	GetImplTypeFlags     uintptr
	GetIDsOfNames        uintptr
	Invoke               uintptr
	GetDocumentation     uintptr
	GetDllEntry          uintptr
	GetRefTypeInfo       uintptr
	AddressOfMember      uintptr
	CreateInstance       uintptr
	GetMops              uintptr
	GetContainingTypeLib uintptr
	ReleaseTypeAttr      uintptr
	ReleaseFuncDesc      uintptr
	ReleaseVarDesc       uintptr
}

type ITypeInfo struct {
	LpVtbl *ITypeInfoVtbl
}

func (obj *ITypeInfo) QueryInterface(riid REFIID, ppvObject *unsafe.Pointer) HRESULT {
	return (*IUnknown)(unsafe.Pointer(obj)).QueryInterface(riid, ppvObject)
}

func (obj *ITypeInfo) AddRef() uint32 {
	return (*IUnknown)(unsafe.Pointer(obj)).AddRef()
}

func (obj *ITypeInfo) Release() uint32 {
	return (*IUnknown)(unsafe.Pointer(obj)).Release()
}
