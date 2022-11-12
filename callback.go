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
void cgoCallOnTrtcExitRoom(void* ctx,int reason);
void cgoCallOnTrtcSendFirstLocalVideoFrame (void* ctx,int streamType);
void cgoCallOnTrtcSendFirstLocalAudioFrame(void* ctx);
void cgoCallOnTrtcRemoteUserEnterRoom(void* ctx,const char *userId);
void cgoCallOnTrtcRemoteUserLeaveRoom(void* ctx,const char *userId,int reason);
void cgoCallOnTrtcConnectionLost(void* ctx);
void cgoCallOnTrtcTryToReconnect(void* ctx);
void cgoCallOnTrtcConnectionRecovery(void* ctx);


*/
import "C"
import (
	"unsafe"
)

//export cgoCallOnTrtcError
func cgoCallOnTrtcError(ctx unsafe.Pointer, errCode C.int, errMsg *C.const_char, extraInfo unsafe.Pointer) {
	if ctx == nil {
		return
	}
	cb := *(*Callback)(ctx)
	cb.OnError(int(errCode), C.GoString(errMsg), extraInfo)
}

//export cgoCallOnTrtcWarn
func cgoCallOnTrtcWarn(ctx unsafe.Pointer, code C.int, msg *C.const_char, extraInfo unsafe.Pointer) {
	if ctx == nil {
		return
	}
	cb := *(*Callback)(ctx)
	cb.OnWarning(int(code), C.GoString(msg), extraInfo)
}

//export cgoCallOnTrtcEnterRoom
func cgoCallOnTrtcEnterRoom(ctx unsafe.Pointer, result C.int) {
	if ctx == nil {
		return
	}
	cb := *(*Callback)(ctx)
	cb.OnEnterRoom(int(result))
}

//export cgoCallOnTrtcExitRoom
func cgoCallOnTrtcExitRoom(ctx unsafe.Pointer, reason C.int) {
	if ctx == nil {
		return
	}
	cb := *(*Callback)(ctx)
	cb.OnExitRoom(int(reason))
}

//export cgoCallOnTrtcSendFirstLocalVideoFrame
func cgoCallOnTrtcSendFirstLocalVideoFrame(ctx unsafe.Pointer, streamType C.int) {
	if ctx == nil {
		return
	}
	cb := *(*Callback)(ctx)
	cb.OnSendFirstLocalVideoFrame(int(streamType))
}

//export cgoCallOnTrtcSendFirstLocalAudioFrame
func cgoCallOnTrtcSendFirstLocalAudioFrame(ctx unsafe.Pointer) {
	if ctx == nil {
		return
	}
	cb := *(*Callback)(ctx)
	cb.OnSendFirstLocalAudioFrame()
}

//export cgoCallOnTrtcRemoteUserEnterRoom
func cgoCallOnTrtcRemoteUserEnterRoom(ctx unsafe.Pointer, userId *C.const_char) {
	if ctx == nil || userId == nil {
		return
	}
	cb := *(*Callback)(ctx)
	cb.OnRemoteUserEnterRoom(C.GoString(userId))
}

//export cgoCallOnTrtcRemoteUserLeaveRoom
func cgoCallOnTrtcRemoteUserLeaveRoom(ctx unsafe.Pointer, userId *C.const_char, reason C.int) {
	if ctx == nil || userId == nil {
		return
	}
	cb := *(*Callback)(ctx)
	cb.OnRemoteUserLeaveRoom(C.GoString(userId), int(reason))
}

//export cgoCallOnTrtcConnectionLost
func cgoCallOnTrtcConnectionLost(ctx unsafe.Pointer) {
	if ctx == nil {
		return
	}
	cb := *(*Callback)(ctx)
	cb.OnConnectionLost()
}

//export cgoCallOnTrtcTryToReconnect
func cgoCallOnTrtcTryToReconnect(ctx unsafe.Pointer) {
	if ctx == nil {
		return
	}
	cb := *(*Callback)(ctx)
	cb.OnTryToReconnect()
}

//export cgoCallOnTrtcConnectionRecovery
func cgoCallOnTrtcConnectionRecovery(ctx unsafe.Pointer) {
	if ctx == nil {
		return
	}
	cb := *(*Callback)(ctx)
	cb.OnConnectionRecovery()
}

func createCallback(callback Callback) *C.CTrtcCloudCallback {
	cb := C.createTrtcCallback(nil)
	cb.ctx = unsafe.Pointer(&callback)
	cb.onError = C.TrtcCBErrorFunc(C.cgoCallOnTrtcError)
	cb.onWarning = C.TrtcCBWarnFunc(C.cgoCallOnTrtcWarn)
	cb.onEnterRoom = C.TrtcCBEnterRoomFunc(C.cgoCallOnTrtcEnterRoom)
	cb.onExitRoom = C.TrtcCBExitRoomFunc(C.cgoCallOnTrtcExitRoom)
	cb.onRemoteUserEnterRoom = C.TrtcCBRemoteUserEnterRoomFunc(C.cgoCallOnTrtcRemoteUserEnterRoom)
	cb.onRemoteUserLeaveRoom = C.TrtcCBRemoteUserLeaveRoomFunc(C.cgoCallOnTrtcRemoteUserLeaveRoom)
	cb.onSendFirstLocalVideoFrame = C.TrtcCBSendFirstLocalVideoFrameFunc(C.cgoCallOnTrtcSendFirstLocalVideoFrame)
	cb.onSendFirstLocalAudioFrame = C.TrtcCBSendFirstLocalAudioFrameFunc(C.cgoCallOnTrtcSendFirstLocalAudioFrame)
	cb.onConnectionLost = C.TrtcCBConnectionLostFunc(C.cgoCallOnTrtcConnectionLost)
	cb.onTryToReconnect = C.TrtcCBTryToReconnectFunc(C.cgoCallOnTrtcTryToReconnect)
	cb.onConnectionRecovery = C.TrtcCBConnectionRecoveryFunc(C.cgoCallOnTrtcConnectionRecovery)
	return cb
}

func destroyCallback(src *C.CTrtcCloudCallback) {
	C.destroyTrtcCallback(src)
}
