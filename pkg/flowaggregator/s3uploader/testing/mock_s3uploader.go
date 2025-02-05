// Copyright 2024 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: antrea.io/antrea/pkg/flowaggregator/s3uploader (interfaces: S3UploaderAPI)
//
// Generated by this command:
//
//	mockgen -copyright_file hack/boilerplate/license_header.raw.txt -destination pkg/flowaggregator/s3uploader/testing/mock_s3uploader.go -package testing antrea.io/antrea/pkg/flowaggregator/s3uploader S3UploaderAPI
//

// Package testing is a generated GoMock package.
package testing

import (
	context "context"
	reflect "reflect"

	manager "github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	s3 "github.com/aws/aws-sdk-go-v2/service/s3"
	gomock "go.uber.org/mock/gomock"
)

// MockS3UploaderAPI is a mock of S3UploaderAPI interface.
type MockS3UploaderAPI struct {
	ctrl     *gomock.Controller
	recorder *MockS3UploaderAPIMockRecorder
	isgomock struct{}
}

// MockS3UploaderAPIMockRecorder is the mock recorder for MockS3UploaderAPI.
type MockS3UploaderAPIMockRecorder struct {
	mock *MockS3UploaderAPI
}

// NewMockS3UploaderAPI creates a new mock instance.
func NewMockS3UploaderAPI(ctrl *gomock.Controller) *MockS3UploaderAPI {
	mock := &MockS3UploaderAPI{ctrl: ctrl}
	mock.recorder = &MockS3UploaderAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockS3UploaderAPI) EXPECT() *MockS3UploaderAPIMockRecorder {
	return m.recorder
}

// Upload mocks base method.
func (m *MockS3UploaderAPI) Upload(ctx context.Context, input *s3.PutObjectInput, awsS3Uploader *manager.Uploader, opts ...func(*manager.Uploader)) (*manager.UploadOutput, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, input, awsS3Uploader}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Upload", varargs...)
	ret0, _ := ret[0].(*manager.UploadOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upload indicates an expected call of Upload.
func (mr *MockS3UploaderAPIMockRecorder) Upload(ctx, input, awsS3Uploader any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, input, awsS3Uploader}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockS3UploaderAPI)(nil).Upload), varargs...)
}
