//go:build windows

/*
 * Copyright (c) 2022 Lynn <lynnplus90@gmail.com>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package gotrtc

/*
#include "CTrtcCloudCallback.h"
typedef const char const_char;

void cgoCallOnTrtcError(void* ctx,int errCode, const char *errMsg, void *extraInfo);
void cgoCallOnTrtcWarn(void* ctx,int warningCode, const char *warningMsg, void *extraInfo);
void cgoCallOnTrtcEnterRoom(void* ctx,int result);
void cgoCallOnTrtcSendFirstLocalVideoFrame (void* ctx,int streamType);
*/
import "C"
import "unsafe"

type Callback interface {
	OnError(errCode int, errMsg string, extraInfo any)
	OnWarning(code int, msg string, extraInfo any)
	OnEnterRoom(result int)
	OnSendFirstLocalVideoFrame(streamType int)
}

//export cgoCallOnTrtcError
func cgoCallOnTrtcError(ctx unsafe.Pointer, errCode C.int, errMsg *C.const_char, extraInfo unsafe.Pointer) {
	cb := *(*Callback)(ctx)
	cb.OnError(int(errCode), C.GoString(errMsg), extraInfo)
}

//export cgoCallOnTrtcWarn
func cgoCallOnTrtcWarn(ctx unsafe.Pointer, code C.int, msg *C.const_char, extraInfo unsafe.Pointer) {
	cb := *(*Callback)(ctx)
	cb.OnWarning(int(code), C.GoString(msg), extraInfo)
}

//export cgoCallOnTrtcEnterRoom
func cgoCallOnTrtcEnterRoom(ctx unsafe.Pointer, result C.int) {
	cb := *(*Callback)(ctx)
	cb.OnEnterRoom(int(result))
}

//export cgoCallOnTrtcSendFirstLocalVideoFrame
func cgoCallOnTrtcSendFirstLocalVideoFrame(ctx unsafe.Pointer, streamType C.int) {
	cb := *(*Callback)(ctx)
	cb.OnSendFirstLocalVideoFrame(int(streamType))
}

func createCallback(callback Callback) *C.CTrtcCloudCallback {
	cb := C.createTrtcCallback(nil)
	cb.ctx = unsafe.Pointer(&callback)
	cb.onError = C.TrtcCBErrorFunc(C.cgoCallOnTrtcError)
	cb.onWarning = C.TrtcCBWarnFunc(C.cgoCallOnTrtcWarn)
	cb.onEnterRoom = C.TrtcCBEnterRoomFunc(C.cgoCallOnTrtcEnterRoom)
	cb.onSendFirstLocalVideoFrame = C.TrtcCBSendFirstLocalVideoFrameFunc(C.cgoCallOnTrtcSendFirstLocalVideoFrame)
	return cb
}

func destroyCallback(src *C.CTrtcCloudCallback) {
	C.destroyTrtcCallback(src)
}
