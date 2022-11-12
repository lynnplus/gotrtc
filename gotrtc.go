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

type TrtcCloud interface {
	GetSDKVersion() string
	SetConsoleEnabled(enable bool)
	SetLogCompressEnabled(enable bool)
	SetLogLevel(level LogLevel)
	SetLogDirPath(path string)

	MuteLocalVideo(streamType VideoStreamType, mute bool)
	MuteLocalAudio(mute bool)

	AddCallback(cb Callback)
	RemoveCallback(cb Callback)

	CreateSubCloud() TrtcCloud
	Destroy()
	IsMainCloud() bool

	EnterRoom(params *RoomParams) error
	ExitRoom()

	SetDefaultStreamRecvMode(autoRecvAudio, autoRecvVideo bool)
	EnableCustomVideoCapture(enable bool)
	EnableCustomAudioCapture(enable bool)

	SetVideoEncoderParam(param *VideoEncoderParam)
	GenerateCustomPTS() uint64
	SendCustomVideoData(frame *VideoFrame)
	StartLocalTest()
}

type Callback interface {
	OnError(errCode int, errMsg string, extraInfo any)
	OnWarning(code int, msg string, extraInfo any)
	OnEnterRoom(result int)
	OnExitRoom(reason int)
	OnSendFirstLocalVideoFrame(streamType int)
	OnSendFirstLocalAudioFrame()
	OnRemoteUserEnterRoom(userId string)
	OnRemoteUserLeaveRoom(userId string, reason int)
	OnConnectionLost()
	OnTryToReconnect()
	OnConnectionRecovery()
}
