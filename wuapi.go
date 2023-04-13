// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows
// +build windows

package win

import (
	"syscall"
	"unsafe"
)

var IID_IUpdateIdentity = IID{0x46297823, 0x9940, 0x4c09, [8]byte{0xae, 0xd9, 0xcd, 0x3e, 0xa6, 0xd0, 0x59, 0x68}}

type IUpdateIdentityVtbl struct {
	IDispatchVtbl
	GetRevisionNumber uintptr
	GetUpdateID       uintptr
}

type IUpdateIdentity struct {
	LpVtbl *IUpdateIdentityVtbl
}

func (obj *IUpdateIdentity) QueryInterface(riid REFIID, ppvObject *unsafe.Pointer) HRESULT {
	return (*IUnknown)(unsafe.Pointer(obj)).QueryInterface(riid, ppvObject)
}

func (obj *IUpdateIdentity) AddRef() uint32 {
	return (*IUnknown)(unsafe.Pointer(obj)).AddRef()
}

func (obj *IUpdateIdentity) Release() uint32 {
	return (*IUnknown)(unsafe.Pointer(obj)).Release()
}

func (obj *IUpdateIdentity) RevisionNumber() (int32, HRESULT) {
	var retval int32
	ret, _, _ := syscall.SyscallN(obj.LpVtbl.GetRevisionNumber,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&retval)))
	return retval, HRESULT(ret)
}

func (obj *IUpdateIdentity) UpdateID() (*uint16 /*BSTR*/, HRESULT) {
	var retval *uint16 /*BSTR*/
	ret, _, _ := syscall.SyscallN(obj.LpVtbl.GetUpdateID,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&retval)))
	return retval, HRESULT(ret)
}

var IID_IUpdateHistoryEntry = IID{0xbe56a644, 0xaf0e, 0x4e0e, [8]byte{0xa3, 0x11, 0xc1, 0xd8, 0xe6, 0x95, 0xcb, 0xff}}

type IUpdateHistoryEntryVtbl struct {
	IDispatchVtbl
	GetOperation           uintptr
	GetResultCode          uintptr
	GetHResult             uintptr
	GetDate                uintptr
	GetUpdateIdentity      uintptr
	GetTitle               uintptr
	GetDescription         uintptr
	GetUnmappedResultCode  uintptr
	GetClientApplicationID uintptr
	GetServerSelection     uintptr
	GetServiceID           uintptr
	GetUninstallationSteps uintptr
	GetUninstallationNotes uintptr
	GetSupportUrl          uintptr
}

type IUpdateHistoryEntry struct {
	LpVtbl *IUpdateHistoryEntryVtbl
}

func (obj *IUpdateHistoryEntry) QueryInterface(riid REFIID, ppvObject *unsafe.Pointer) HRESULT {
	return (*IUnknown)(unsafe.Pointer(obj)).QueryInterface(riid, ppvObject)
}

func (obj *IUpdateHistoryEntry) AddRef() uint32 {
	return (*IUnknown)(unsafe.Pointer(obj)).AddRef()
}

func (obj *IUpdateHistoryEntry) Release() uint32 {
	return (*IUnknown)(unsafe.Pointer(obj)).Release()
}

type UpdateOperation int32

const (
	UOInstallation   UpdateOperation = 1
	UOUninstallation UpdateOperation = 2
)

func (obj *IUpdateHistoryEntry) Operation() (UpdateOperation, HRESULT) {
	var retval UpdateOperation
	ret, _, _ := syscall.SyscallN(obj.LpVtbl.GetOperation,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&retval)))
	return retval, HRESULT(ret)
}

type OperationResultCode int32

const (
	ORCNotStarted          OperationResultCode = 0
	ORCInProgress          OperationResultCode = 1
	ORCSucceeded           OperationResultCode = 2
	ORCSucceededWithErrors OperationResultCode = 3
	ORCFailed              OperationResultCode = 4
	ORCAborted             OperationResultCode = 5
)

func (obj *IUpdateHistoryEntry) ResultCode() (OperationResultCode, HRESULT) {
	var retval OperationResultCode
	ret, _, _ := syscall.SyscallN(obj.LpVtbl.GetResultCode,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&retval)))
	return retval, HRESULT(ret)
}

func (obj *IUpdateHistoryEntry) HResult() (HRESULT, HRESULT) {
	var retval HRESULT
	ret, _, _ := syscall.SyscallN(obj.LpVtbl.GetHResult,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&retval)))
	return retval, HRESULT(ret)
}

func (obj *IUpdateHistoryEntry) Date() (float64 /*DATE*/, HRESULT) {
	var retval float64 /*DATE*/
	ret, _, _ := syscall.SyscallN(obj.LpVtbl.GetDate,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&retval)))
	return retval, HRESULT(ret)
}

func (obj *IUpdateHistoryEntry) UpdateIdentity() (*IUpdateIdentity, HRESULT) {
	var retval *IUpdateIdentity
	ret, _, _ := syscall.SyscallN(obj.LpVtbl.GetUpdateIdentity,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&retval)))
	return retval, HRESULT(ret)
}

func (obj *IUpdateHistoryEntry) Title() (*uint16 /*BSTR*/, HRESULT) {
	var retval *uint16 /*BSTR*/
	ret, _, _ := syscall.SyscallN(obj.LpVtbl.GetTitle,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&retval)))
	return retval, HRESULT(ret)
}

func (obj *IUpdateHistoryEntry) Description() (*uint16 /*BSTR*/, HRESULT) {
	var retval *uint16 /*BSTR*/
	ret, _, _ := syscall.SyscallN(obj.LpVtbl.GetDescription,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&retval)))
	return retval, HRESULT(ret)
}

var IID_IUpdateHistoryEntryCollection = IID{0xa7f04f3c, 0xa290, 0x435b, [8]byte{0xaa, 0xdf, 0xa1, 0x16, 0xc3, 0x35, 0x7a, 0x5c}}

type IUpdateHistoryEntryCollectionVtbl struct {
	IDispatchVtbl
	GetItem     uintptr
	Get_NewEnum uintptr
	GetCount    uintptr
}

type IUpdateHistoryEntryCollection struct {
	LpVtbl *IUpdateHistoryEntryCollectionVtbl
}

func (obj *IUpdateHistoryEntryCollection) QueryInterface(riid REFIID, ppvObject *unsafe.Pointer) HRESULT {
	return (*IUnknown)(unsafe.Pointer(obj)).QueryInterface(riid, ppvObject)
}

func (obj *IUpdateHistoryEntryCollection) AddRef() uint32 {
	return (*IUnknown)(unsafe.Pointer(obj)).AddRef()
}

func (obj *IUpdateHistoryEntryCollection) Release() uint32 {
	return (*IUnknown)(unsafe.Pointer(obj)).Release()
}

func (obj *IUpdateHistoryEntryCollection) Item(index int32) (*IUpdateHistoryEntry, HRESULT) {
	var retval *IUpdateHistoryEntry
	ret, _, _ := syscall.SyscallN(obj.LpVtbl.GetItem,
		uintptr(unsafe.Pointer(obj)),
		uintptr(index),
		uintptr(unsafe.Pointer(&retval)))
	return retval, HRESULT(ret)
}

func (obj *IUpdateHistoryEntryCollection) NewEnum() (*IUnknown, HRESULT) {
	var retval *IUnknown
	ret, _, _ := syscall.SyscallN(obj.LpVtbl.Get_NewEnum,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&retval)))
	return retval, HRESULT(ret)
}

func (obj *IUpdateHistoryEntryCollection) Count() (int32, HRESULT) {
	var retval int32
	ret, _, _ := syscall.SyscallN(obj.LpVtbl.GetCount,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&retval)))
	return retval, HRESULT(ret)
}

var IID_IUpdateServiceManager = IID{0x23857e3c, 0x02ba, 0x44a3, [8]byte{0x94, 0x23, 0xb1, 0xc9, 0x00, 0x80, 0x5f, 0x37}}

type IUpdateServiceManagerVtbl struct {
	IDispatch
	GetServices             uintptr
	AddService              uintptr
	RegisterServiceWithAU   uintptr
	RemoveService           uintptr
	UnregisterServiceWithAU uintptr
	AddScanPackageService   uintptr
	SetOption               uintptr
}

type IUpdateServiceManager struct {
	LpVtbl *IUpdateServiceManagerVtbl
}

var IID_IUpdateServiceManager2 = IID{0x0bb8531d, 0x7e8d, 0x424f, [8]byte{0x98, 0x6c, 0xa0, 0xb8, 0xf6, 0x0a, 0x3e, 0x7b}}

type IUpdateServiceManager2Vtbl struct {
	IUpdateServiceManagerVtbl
	GetClientApplicationID   uintptr
	PutClientApplicationID   uintptr
	QueryServiceRegistration uintptr
	AddService2              uintptr
}

type IUpdateServiceManager2 struct {
	LpVtbl *IUpdateServiceManager2Vtbl
}

var (
	IID_IUpdateSession  = IID{0x816858a4, 0x260d, 0x4260, [8]byte{0x93, 0x3a, 0x25, 0x85, 0xf1, 0xab, 0xc7, 0x6b}}
	CLSID_UpdateSession = CLSID{0x4CB43D7F, 0x7EEE, 0x4906, [8]byte{0x86, 0x98, 0x60, 0xDA, 0x1C, 0x38, 0xF2, 0xFE}}
)

type IUpdateSessionVtbl struct {
	IDispatchVtbl
	GetClientApplicationID uintptr
	PutClientApplicationID uintptr
	GetReadOnly            uintptr
	GetWebProxy            uintptr
	PutWebProxy            uintptr
	CreateUpdateSearcher   uintptr
	CreateUpdateDownloader uintptr
	CreateUpdateInstaller  uintptr
}

type IUpdateSession struct {
	LpVtbl *IUpdateSessionVtbl
}

func (obj *IUpdateSession) ClientApplicationID() (*uint16 /*BSTR*/, HRESULT) {
	var retval *uint16 /*BSTR*/
	ret, _, _ := syscall.SyscallN(obj.LpVtbl.GetClientApplicationID,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&retval)))
	return retval, HRESULT(ret)
}

func (obj *IUpdateSession) SetClientApplicationID(value *uint16 /*BSTR*/) HRESULT {
	ret, _, _ := syscall.SyscallN(obj.LpVtbl.PutClientApplicationID,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(value)))
	return HRESULT(ret)
}

func (obj *IUpdateSession) ReadOnly() (VARIANT_BOOL, HRESULT) {
	var retval VARIANT_BOOL
	ret, _, _ := syscall.SyscallN(obj.LpVtbl.GetReadOnly,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&retval)))
	return retval, HRESULT(ret)
}

var IID_IUpdateSession2 = IID{0x918EFD1E, 0xB5D8, 0x4c90, [8]byte{0x85, 0x40, 0xAE, 0xB9, 0xBD, 0xC5, 0x6F, 0x9D}}

type IUpdateSession2Vtbl struct {
	IUpdateSessionVtbl
	GetUserLocale uintptr
	PutUserLocale uintptr
}

type IUpdateSession2 struct {
	LpVtbl *IUpdateSession2Vtbl
}

func (obj *IUpdateSession2) UserLocale() (LCID, HRESULT) {
	var retval LCID
	ret, _, _ := syscall.SyscallN(obj.LpVtbl.GetUserLocale,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&retval)))
	return retval, HRESULT(ret)
}

func (obj *IUpdateSession2) PutUserLocale(lcid LCID) HRESULT {
	ret, _, _ := syscall.SyscallN(obj.LpVtbl.PutUserLocale,
		uintptr(unsafe.Pointer(obj)),
		uintptr(lcid))
	return HRESULT(ret)
}

var IID_IUpdateSession3 = IID{0x918EFD1E, 0xB5D8, 0x4c90, [8]byte{0x85, 0x40, 0xAE, 0xB9, 0xBD, 0xC5, 0x6F, 0x9D}}

type IUpdateSession3Vtbl struct {
	IUpdateSession2Vtbl
	CreateUpdateServiceManager uintptr
	QueryHistory               uintptr
}

type IUpdateSession3 struct {
	LpVtbl *IUpdateSession3Vtbl
}

func (obj *IUpdateSession3) QueryInterface(riid REFIID, ppvObject *unsafe.Pointer) HRESULT {
	return (*IUnknown)(unsafe.Pointer(obj)).QueryInterface(riid, ppvObject)
}

func (obj *IUpdateSession3) AddRef() uint32 {
	return (*IUnknown)(unsafe.Pointer(obj)).AddRef()
}

func (obj *IUpdateSession3) Release() uint32 {
	return (*IUnknown)(unsafe.Pointer(obj)).Release()
}

func (obj *IUpdateSession3) CreateUpdateServiceManager() (*IUpdateServiceManager2, HRESULT) {
	var retval *IUpdateServiceManager2
	ret, _, _ := syscall.SyscallN(obj.LpVtbl.CreateUpdateServiceManager,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&retval)))
	return retval, HRESULT(ret)
}

func (obj *IUpdateSession3) QueryHistory(criteria *uint16 /*BSTR*/, startIndex int32, count int32) (*IUpdateHistoryEntryCollection, HRESULT) {
	var retval *IUpdateHistoryEntryCollection
	ret, _, _ := syscall.SyscallN(obj.LpVtbl.QueryHistory,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(criteria)),
		uintptr(startIndex),
		uintptr(count),
		uintptr(unsafe.Pointer(&retval)))
	return retval, HRESULT(ret)
}
