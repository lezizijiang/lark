// Code generated by lark_sdk_gen. DO NOT EDIT.
/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_Baike_Sample_Failed(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		cli.Mock().MockGetAppAccessToken(mockGetTenantAccessTokenFailed)
		moduleCli := cli.Baike

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateBaikeDraft(ctx, &lark.CreateBaikeDraftReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

	})

	t.Run("request mock failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		moduleCli := cli.Baike

		t.Run("", func(t *testing.T) {

			cli.Mock().MockBaikeCreateBaikeDraft(func(ctx context.Context, request *lark.CreateBaikeDraftReq, options ...lark.MethodOptionFunc) (*lark.CreateBaikeDraftResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockBaikeCreateBaikeDraft()

			_, _, err := moduleCli.CreateBaikeDraft(ctx, &lark.CreateBaikeDraftReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockBaikeCreateBaikeUpdate(func(ctx context.Context, request *lark.CreateBaikeUpdateReq, options ...lark.MethodOptionFunc) (*lark.CreateBaikeUpdateResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockBaikeCreateBaikeUpdate()

			_, _, err := moduleCli.CreateBaikeUpdate(ctx, &lark.CreateBaikeUpdateReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockBaikeCreateBaikeEntity(func(ctx context.Context, request *lark.CreateBaikeEntityReq, options ...lark.MethodOptionFunc) (*lark.CreateBaikeEntityResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockBaikeCreateBaikeEntity()

			_, _, err := moduleCli.CreateBaikeEntity(ctx, &lark.CreateBaikeEntityReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockBaikeUpdateBaikeEntity(func(ctx context.Context, request *lark.UpdateBaikeEntityReq, options ...lark.MethodOptionFunc) (*lark.UpdateBaikeEntityResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockBaikeUpdateBaikeEntity()

			_, _, err := moduleCli.UpdateBaikeEntity(ctx, &lark.UpdateBaikeEntityReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockBaikeGetBaikeEntity(func(ctx context.Context, request *lark.GetBaikeEntityReq, options ...lark.MethodOptionFunc) (*lark.GetBaikeEntityResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockBaikeGetBaikeEntity()

			_, _, err := moduleCli.GetBaikeEntity(ctx, &lark.GetBaikeEntityReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockBaikeGetBaikeEntityList(func(ctx context.Context, request *lark.GetBaikeEntityListReq, options ...lark.MethodOptionFunc) (*lark.GetBaikeEntityListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockBaikeGetBaikeEntityList()

			_, _, err := moduleCli.GetBaikeEntityList(ctx, &lark.GetBaikeEntityListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockBaikeMatchBaikeEntity(func(ctx context.Context, request *lark.MatchBaikeEntityReq, options ...lark.MethodOptionFunc) (*lark.MatchBaikeEntityResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockBaikeMatchBaikeEntity()

			_, _, err := moduleCli.MatchBaikeEntity(ctx, &lark.MatchBaikeEntityReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockBaikeSearchBaikeEntity(func(ctx context.Context, request *lark.SearchBaikeEntityReq, options ...lark.MethodOptionFunc) (*lark.SearchBaikeEntityResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockBaikeSearchBaikeEntity()

			_, _, err := moduleCli.SearchBaikeEntity(ctx, &lark.SearchBaikeEntityReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockBaikeHighlightBaikeEntity(func(ctx context.Context, request *lark.HighlightBaikeEntityReq, options ...lark.MethodOptionFunc) (*lark.HighlightBaikeEntityResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockBaikeHighlightBaikeEntity()

			_, _, err := moduleCli.HighlightBaikeEntity(ctx, &lark.HighlightBaikeEntityReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockBaikeExtractBaikeEntity(func(ctx context.Context, request *lark.ExtractBaikeEntityReq, options ...lark.MethodOptionFunc) (*lark.ExtractBaikeEntityResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockBaikeExtractBaikeEntity()

			_, _, err := moduleCli.ExtractBaikeEntity(ctx, &lark.ExtractBaikeEntityReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockBaikeGetBaikeClassificationList(func(ctx context.Context, request *lark.GetBaikeClassificationListReq, options ...lark.MethodOptionFunc) (*lark.GetBaikeClassificationListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockBaikeGetBaikeClassificationList()

			_, _, err := moduleCli.GetBaikeClassificationList(ctx, &lark.GetBaikeClassificationListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockBaikeUploadBaikeImage(func(ctx context.Context, request *lark.UploadBaikeImageReq, options ...lark.MethodOptionFunc) (*lark.UploadBaikeImageResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockBaikeUploadBaikeImage()

			_, _, err := moduleCli.UploadBaikeImage(ctx, &lark.UploadBaikeImageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {

			cli.Mock().MockBaikeDownloadBaikeImage(func(ctx context.Context, request *lark.DownloadBaikeImageReq, options ...lark.MethodOptionFunc) (*lark.DownloadBaikeImageResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockBaikeDownloadBaikeImage()

			_, _, err := moduleCli.DownloadBaikeImage(ctx, &lark.DownloadBaikeImageReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

	})

	t.Run("response is failed", func(t *testing.T) {
		cli := AppNoPermission.Ins()
		moduleCli := cli.Baike

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateBaikeDraft(ctx, &lark.CreateBaikeDraftReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateBaikeUpdate(ctx, &lark.CreateBaikeUpdateReq{
				DraftID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateBaikeEntity(ctx, &lark.CreateBaikeEntityReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateBaikeEntity(ctx, &lark.UpdateBaikeEntityReq{
				EntityID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetBaikeEntity(ctx, &lark.GetBaikeEntityReq{
				EntityID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetBaikeEntityList(ctx, &lark.GetBaikeEntityListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.MatchBaikeEntity(ctx, &lark.MatchBaikeEntityReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.SearchBaikeEntity(ctx, &lark.SearchBaikeEntityReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.HighlightBaikeEntity(ctx, &lark.HighlightBaikeEntityReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.ExtractBaikeEntity(ctx, &lark.ExtractBaikeEntityReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetBaikeClassificationList(ctx, &lark.GetBaikeClassificationListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UploadBaikeImage(ctx, &lark.UploadBaikeImageReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DownloadBaikeImage(ctx, &lark.DownloadBaikeImageReq{
				FileToken: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

	})

	t.Run("fake request is failed", func(t *testing.T) {
		cli := AppAllPermission.Ins()
		moduleCli := cli.Baike
		cli.Mock().MockRawRequest(func(ctx context.Context, req *lark.RawRequestReq, resp interface{}) (response *lark.Response, err error) {
			return nil, fmt.Errorf("fake raw request")
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateBaikeDraft(ctx, &lark.CreateBaikeDraftReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateBaikeUpdate(ctx, &lark.CreateBaikeUpdateReq{
				DraftID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.CreateBaikeEntity(ctx, &lark.CreateBaikeEntityReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UpdateBaikeEntity(ctx, &lark.UpdateBaikeEntityReq{
				EntityID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetBaikeEntity(ctx, &lark.GetBaikeEntityReq{
				EntityID: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetBaikeEntityList(ctx, &lark.GetBaikeEntityListReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.MatchBaikeEntity(ctx, &lark.MatchBaikeEntityReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.SearchBaikeEntity(ctx, &lark.SearchBaikeEntityReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.HighlightBaikeEntity(ctx, &lark.HighlightBaikeEntityReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.ExtractBaikeEntity(ctx, &lark.ExtractBaikeEntityReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.GetBaikeClassificationList(ctx, &lark.GetBaikeClassificationListReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.UploadBaikeImage(ctx, &lark.UploadBaikeImageReq{})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

		t.Run("", func(t *testing.T) {

			_, _, err := moduleCli.DownloadBaikeImage(ctx, &lark.DownloadBaikeImageReq{
				FileToken: "x",
			})
			as.NotNil(err)
			as.Equal("fake raw request", err.Error())
		})

	})
}
