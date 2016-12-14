// Copyright 2015-2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/aws/amazon-ecs-cli/ecs-cli/modules/compose/ecs (interfaces: ProjectEntity)

package mock_ecs

import (
	ecs "github.com/aws/amazon-ecs-cli/ecs-cli/modules/compose/ecs"
	utils "github.com/aws/amazon-ecs-cli/ecs-cli/utils"
	cache "github.com/aws/amazon-ecs-cli/ecs-cli/utils/cache"
	ecs0 "github.com/aws/aws-sdk-go/service/ecs"
	project "github.com/docker/libcompose/project"
	gomock "github.com/golang/mock/gomock"
)

// Mock of ProjectEntity interface
type MockProjectEntity struct {
	ctrl     *gomock.Controller
	recorder *_MockProjectEntityRecorder
}

// Recorder for MockProjectEntity (not exported)
type _MockProjectEntityRecorder struct {
	mock *MockProjectEntity
}

func NewMockProjectEntity(ctrl *gomock.Controller) *MockProjectEntity {
	mock := &MockProjectEntity{ctrl: ctrl}
	mock.recorder = &_MockProjectEntityRecorder{mock}
	return mock
}

func (_m *MockProjectEntity) EXPECT() *_MockProjectEntityRecorder {
	return _m.recorder
}

func (_m *MockProjectEntity) Context() *ecs.Context {
	ret := _m.ctrl.Call(_m, "Context")
	ret0, _ := ret[0].(*ecs.Context)
	return ret0
}

func (_mr *_MockProjectEntityRecorder) Context() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Context")
}

func (_m *MockProjectEntity) Create() error {
	ret := _m.ctrl.Call(_m, "Create")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockProjectEntityRecorder) Create() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Create")
}

func (_m *MockProjectEntity) Down() error {
	ret := _m.ctrl.Call(_m, "Down")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockProjectEntityRecorder) Down() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Down")
}

func (_m *MockProjectEntity) Info(_param0 bool) (project.InfoSet, error) {
	ret := _m.ctrl.Call(_m, "Info", _param0)
	ret0, _ := ret[0].(project.InfoSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockProjectEntityRecorder) Info(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Info", arg0)
}

func (_m *MockProjectEntity) LoadContext() error {
	ret := _m.ctrl.Call(_m, "LoadContext")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockProjectEntityRecorder) LoadContext() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LoadContext")
}

func (_m *MockProjectEntity) Run(_param0 map[string]string) error {
	ret := _m.ctrl.Call(_m, "Run", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockProjectEntityRecorder) Run(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Run", arg0)
}

func (_m *MockProjectEntity) Scale(_param0 int) error {
	ret := _m.ctrl.Call(_m, "Scale", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockProjectEntityRecorder) Scale(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Scale", arg0)
}

func (_m *MockProjectEntity) SetTaskDefinition(_param0 *ecs0.TaskDefinition) {
	_m.ctrl.Call(_m, "SetTaskDefinition", _param0)
}

func (_mr *_MockProjectEntityRecorder) SetTaskDefinition(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetTaskDefinition", arg0)
}

func (_m *MockProjectEntity) Sleeper() *utils.TimeSleeper {
	ret := _m.ctrl.Call(_m, "Sleeper")
	ret0, _ := ret[0].(*utils.TimeSleeper)
	return ret0
}

func (_mr *_MockProjectEntityRecorder) Sleeper() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Sleeper")
}

func (_m *MockProjectEntity) Start() error {
	ret := _m.ctrl.Call(_m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockProjectEntityRecorder) Start() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Start")
}

func (_m *MockProjectEntity) Stop() error {
	ret := _m.ctrl.Call(_m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockProjectEntityRecorder) Stop() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Stop")
}

func (_m *MockProjectEntity) TaskDefinition() *ecs0.TaskDefinition {
	ret := _m.ctrl.Call(_m, "TaskDefinition")
	ret0, _ := ret[0].(*ecs0.TaskDefinition)
	return ret0
}

func (_mr *_MockProjectEntityRecorder) TaskDefinition() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TaskDefinition")
}

func (_m *MockProjectEntity) TaskDefinitionCache() cache.Cache {
	ret := _m.ctrl.Call(_m, "TaskDefinitionCache")
	ret0, _ := ret[0].(cache.Cache)
	return ret0
}

func (_mr *_MockProjectEntityRecorder) TaskDefinitionCache() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TaskDefinitionCache")
}

func (_m *MockProjectEntity) Up() error {
	ret := _m.ctrl.Call(_m, "Up")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockProjectEntityRecorder) Up() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Up")
}
