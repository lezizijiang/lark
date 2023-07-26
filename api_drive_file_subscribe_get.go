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

package lark

import (
	"context"
)

// GetSubscribeDriveFile 该接口仅支持文档拥有者查询自己文档的订阅状态, 可订阅的文档类型为旧版文档、新版文档、电子表格和多维表格。在调用该接口之前请确保正确[配置事件回调网址和订阅事件类型](https://open.feishu.cn/document/ukTMukTMukTM/uUTNz4SN1MjL1UzM#2eb3504a), 事件类型参考[事件列表](https://open.feishu.cn/document/ukTMukTMukTM/uYDNxYjL2QTM24iN0EjN/event-list)。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/get_subscribe
func (r *DriveService) GetSubscribeDriveFile(ctx context.Context, request *GetSubscribeDriveFileReq, options ...MethodOptionFunc) (*GetSubscribeDriveFileResp, *Response, error) {
	if r.cli.mock.mockDriveGetSubscribeDriveFile != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#GetSubscribeDriveFile mock enable")
		return r.cli.mock.mockDriveGetSubscribeDriveFile(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "GetSubscribeDriveFile",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/drive/v1/files/:file_token/get_subscribe",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getSubscribeDriveFileResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveGetSubscribeDriveFile mock DriveGetSubscribeDriveFile method
func (r *Mock) MockDriveGetSubscribeDriveFile(f func(ctx context.Context, request *GetSubscribeDriveFileReq, options ...MethodOptionFunc) (*GetSubscribeDriveFileResp, *Response, error)) {
	r.mockDriveGetSubscribeDriveFile = f
}

// UnMockDriveGetSubscribeDriveFile un-mock DriveGetSubscribeDriveFile method
func (r *Mock) UnMockDriveGetSubscribeDriveFile() {
	r.mockDriveGetSubscribeDriveFile = nil
}

// GetSubscribeDriveFileReq ...
type GetSubscribeDriveFileReq struct {
	FileToken string   `path:"file_token" json:"-"` // 文档token, 示例值: "doccnxxxxxxxxxxxxxxxxxxxxxx"
	FileType  FileType `query:"file_type" json:"-"` // 文档类型, 示例值: doc, 可选值有: doc: 文档, docx: docx文档, sheet: 表格, bitable: 多维表格, file: 文件
}

// GetSubscribeDriveFileResp ...
type GetSubscribeDriveFileResp struct {
}

// getSubscribeDriveFileResp ...
type getSubscribeDriveFileResp struct {
	Code int64                      `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 错误描述
	Data *GetSubscribeDriveFileResp `json:"data,omitempty"`
}
