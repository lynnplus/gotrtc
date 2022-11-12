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

//#cgo LDFLAGS: -lctrtc
//#include "CTrtcCloud.h"
//#include <stdlib.h>
import "C"
import (
	"errors"
	"sync"
	"unsafe"
)

var mainCloud *TrtcCloud
var initLock sync.Mutex

func GetShareInstance() *TrtcCloud {
	if mainCloud != nil {
		return mainCloud
	}
	initLock.Lock()
	if mainCloud == nil {
		p := unsafe.Pointer(C.getTrtcGlobalShareInstance())
		mainCloud = newTrtcCloud(true, p)
	}
	initLock.Unlock()
	return mainCloud
}

func DestroyShareInstance() {
	if mainCloud == nil {
		return
	}
	initLock.Lock()
	if mainCloud != nil {
		C.destroyTrtcGlobalShareInstance()
		mainCloud.p = nil
		mainCloud.callbackCache = nil
		mainCloud = nil
	}
	initLock.Unlock()
}

func newTrtcCloud(main bool, pointer unsafe.Pointer) *TrtcCloud {
	return &TrtcCloud{
		p:             pointer,
		isMain:        main,
		callbackCache: map[Callback]*C.CTrtcCloudCallback{},
	}
}

type TrtcCloud struct {
	p             unsafe.Pointer
	isMain        bool
	callbackCache map[Callback]*C.CTrtcCloudCallback
}

func (tc *TrtcCloud) GetSDKVersion() string {
	return C.GoString(C.getTrtcSDKVersion((C.CTrtcCloud)(tc.p)))
}

func (tc *TrtcCloud) AddCallback(cb Callback) {
	temp := createCallback(cb)
	C.addTrtcCallback((C.CTrtcCloud)(tc.p), temp)
	tc.callbackCache[cb] = temp
}

func (tc *TrtcCloud) RemoveCallback(cb Callback) {
	temp, ok := tc.callbackCache[cb]
	if ok {
		C.removeTrtcCallback((C.CTrtcCloud)(tc.p), temp)
		destroyCallback(temp)
	}
}

func (tc *TrtcCloud) CreateSubCloud() *TrtcCloud {
	p := C.createTrtcSubCloud()
	return newTrtcCloud(false, unsafe.Pointer(p))
}

func (tc *TrtcCloud) DestroySubCloud(sub *TrtcCloud) {
	if sub.isMain {
		panic("must be use main destroy")
	}
	C.destroyTrtcSubCloud((C.CTrtcCloud)(sub.p))
}

func (tc *TrtcCloud) EnterRoom(params *RoomParams) error {
	if params == nil || params.UserId == "" || params.UserSignature == "" {
		return errors.New("enter room params err")
	}
	if params.RoomId == 0 && params.StrRoomId == "" {
		return errors.New("enter room params err:room_id nil")
	}

	if params.Role == 0 {
		params.Role = RoleTypeAnchor
	}

	cUserId := C.CString(params.UserId)
	cUserSig := C.CString(params.UserSignature)
	cPk := C.CString(params.PrivateMapKey)
	cRoomId := C.CString(params.StrRoomId)
	defer func() {
		C.free(unsafe.Pointer(cUserId))
		C.free(unsafe.Pointer(cUserSig))
		C.free(unsafe.Pointer(cPk))
		C.free(unsafe.Pointer(cRoomId))
	}()

	var param C.CTRTCParams
	param.sdkAppId = C.uint32_t(params.AppId)
	param.roomId = C.uint32_t(params.RoomId)
	param.userId = cUserId
	param.userSig = cUserSig
	param.strRoomId = cRoomId
	param.privateMapKey = cPk
	param.role = params.Role
	C.enterTrtcRoom((C.CTrtcCloud)(tc.p), &param, C.TRTCAppSceneVideoCall)

	return nil
}

func (tc *TrtcCloud) ExitRoom() {
	C.exitTrtcRoom((C.CTrtcCloud)(tc.p))
}

func (tc *TrtcCloud) SetDefaultStreamRecvMode(autoRecvAudio, autoRecvVideo bool) {
	C.setTrtcDefaultStreamRecvMode((C.CTrtcCloud)(tc.p), C.bool(autoRecvAudio), C.bool(autoRecvVideo))
}

func (tc *TrtcCloud) EnableCustomVideoCapture(enable bool) {
	C.enableTrtcCustomVideoCapture((C.CTrtcCloud)(tc.p), C.TRTCVideoStreamTypeBig, C.bool(enable))
}

func (tc *TrtcCloud) EnableCustomAudioCapture(enable bool) {
	C.enableTrtcCustomAudioCapture((C.CTrtcCloud)(tc.p), C.bool(enable))
}

func (tc *TrtcCloud) SetVideoEncoderParam(param *VideoEncoderParam) {
	var data C.CTRTCVideoEncParam
	data.videoResolution = param.Resolution
	data.resMode = param.ResolutionMode
	data.videoFps = C.uint32_t(param.Fps)
	data.videoBitrate = C.uint32_t(param.Bitrate)
	data.minVideoBitrateC.uint32_t(param.MinBitrate)
	data.enableAdjustRes = C.bool(param.EnableAdjustRes)
	C.setTrtcVideoEncoderParam((C.CTrtcCloud)(tc.p), &data)
}

func (tc *TrtcCloud) GenerateCustomPTS() uint64 {
	return uint64(C.generateTrtcCustomPTS((C.CTrtcCloud)(tc.p)))
}

func (tc *TrtcCloud) SendCustomVideoData(frame *VideoFrame) {
	C.sendTrtcCustomVideoData((C.CTrtcCloud)(tc.p), frame.StreamType,
		C.int(frame.Width),
		C.int(frame.Height),
		(*C.char)(unsafe.Pointer(&frame.Buffer[0])),
		C.int(frame.BufferLen),
		C.uint64_t(frame.Timestamp),
		frame.Rotation)
}

func (tc *TrtcCloud) StartLocalTest() {
	C.startTrtcLocalTest((C.CTrtcCloud)(tc.p))
}
