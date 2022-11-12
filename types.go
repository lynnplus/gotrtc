//go:build windows || darwin

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

//#include "CTypeDef.h"
import "C"

type RoleType C.CTRTCRoleType

const (
	RoleTypeAnchor   RoleType = C.TRTCRoleAnchor
	RoleTypeAudience RoleType = C.TRTCRoleAudience
)

type AppScene C.CTRTCAppScene

const (
	AppSceneVideoCall     AppScene = C.TRTCAppSceneVideoCall
	AppSceneLive          AppScene = C.TRTCAppSceneLIVE
	AppSceneAudioCall     AppScene = C.TRTCAppSceneAudioCall
	AppSceneVoiceChatRoom AppScene = C.TRTCAppSceneVoiceChatRoom
)

type VideoStreamType C.CTRTCVideoStreamType

const (
	VideoStreamTypeBig   VideoStreamType = C.TRTCVideoStreamTypeBig
	VideoStreamTypeSmall VideoStreamType = C.TRTCVideoStreamTypeSmall
	VideoStreamTypeSub   VideoStreamType = C.TRTCVideoStreamTypeSub
)

type VideoRotation C.CTRTCVideoRotation

const (
	VideoRotation0   VideoRotation = C.TRTCVideoRotation0
	VideoRotation90  VideoRotation = C.TRTCVideoRotation90
	VideoRotation180 VideoRotation = C.TRTCVideoRotation180
	VideoRotation270 VideoRotation = C.TRTCVideoRotation270
)

type VideoResolution C.int

const (
	// VideoResolution_480_360 宽高比 4:3；分辨率 480x360；建议码率（VideoCall）400kbps; 建议码率（LIVE）600kbps。
	VideoResolution_480_360 VideoResolution = 60
	// VideoResolution_640_480 宽高比 4:3；分辨率 640x480；建议码率（VideoCall）600kbps; 建议码率（LIVE）900kbps。
	VideoResolution_640_480 VideoResolution = 62
	// VideoResolution_960_720 宽高比 4:3；分辨率 960x720；建议码率（VideoCall）1000kbps; 建议码率（LIVE）1500kbps。
	VideoResolution_960_720 VideoResolution = 64
	// VideoResolution_480_270 宽高比 16:9；分辨率 480x270；建议码率（VideoCall）350kbps; 建议码率（LIVE）550kbps。
	VideoResolution_480_270 VideoResolution = 106
	// VideoResolution_640_360 宽高比 16:9；分辨率 640x360；建议码率（VideoCall）500kbps; 建议码率（LIVE）900kbps。
	VideoResolution_640_360 VideoResolution = 108
	// VideoResolution_960_540 宽高比 16:9；分辨率 960x540；建议码率（VideoCall）850kbps; 建议码率（LIVE）1300kbps。
	VideoResolution_960_540 VideoResolution = 110
	// VideoResolution_1280_720 宽高比 16:9；分辨率 1280x720；建议码率（VideoCall）1200kbps; 建议码率（LIVE）1800kbps。
	VideoResolution_1280_720 VideoResolution = 112
	// VideoResolution_1920_1080 宽高比 16:9；分辨率 1920x1080；建议码率（VideoCall）2000kbps; 建议码率（LIVE）3000kbps。
	VideoResolution_1920_1080 VideoResolution = 114
)

type VideoResolutionMode C.CTRTCVideoResolutionMode

const (
	VideoResolutionModeLandscape VideoResolutionMode = C.TRTCVideoResolutionModeLandscape
	VideoResolutionModePortrait  VideoResolutionMode = C.TRTCVideoResolutionModePortrait
)

type LogLevel int

const (
	LogLevelVerbose LogLevel = 0
	LogLevelDebug   LogLevel = 1
	LogLevelInfo    LogLevel = 2
	LogLevelWarn    LogLevel = 3
	LogLevelError   LogLevel = 4
	LogLevelFatal   LogLevel = 5
	LogLevelNone    LogLevel = 6
)

type VideoEncoderParam struct {
	Resolution      VideoResolution
	ResolutionMode  VideoResolutionMode
	Fps             uint
	Bitrate         uint
	MinBitrate      uint
	EnableAdjustRes bool
}

type VideoFrame struct {
	StreamType VideoStreamType
	Rotation   VideoRotation
	Width      int
	Height     int
	Buffer     []byte
	BufferLen  int
	Timestamp  int64
}

type RoomParams struct {
	AppId         uint64
	UserId        string
	UserSignature string
	StrRoomId     string
	RoomId        uint
	PrivateMapKey string
	Role          RoleType
}
