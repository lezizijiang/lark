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

// CreateDriveFileShortcut 创建指定文件的快捷方式到云空间的其它文件夹中。此接口不支持在同一个文件夹下并发创建多个快捷方式。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/create_shortcut
func (r *DriveService) CreateDriveFileShortcut(ctx context.Context, request *CreateDriveFileShortcutReq, options ...MethodOptionFunc) (*CreateDriveFileShortcutResp, *Response, error) {
	if r.cli.mock.mockDriveCreateDriveFileShortcut != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#CreateDriveFileShortcut mock enable")
		return r.cli.mock.mockDriveCreateDriveFileShortcut(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "CreateDriveFileShortcut",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/drive/v1/files/create_shortcut",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(createDriveFileShortcutResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveCreateDriveFileShortcut mock DriveCreateDriveFileShortcut method
func (r *Mock) MockDriveCreateDriveFileShortcut(f func(ctx context.Context, request *CreateDriveFileShortcutReq, options ...MethodOptionFunc) (*CreateDriveFileShortcutResp, *Response, error)) {
	r.mockDriveCreateDriveFileShortcut = f
}

// UnMockDriveCreateDriveFileShortcut un-mock DriveCreateDriveFileShortcut method
func (r *Mock) UnMockDriveCreateDriveFileShortcut() {
	r.mockDriveCreateDriveFileShortcut = nil
}

// CreateDriveFileShortcutReq ...
type CreateDriveFileShortcutReq struct {
	ParentToken string                                 `json:"parent_token,omitempty"` // 目标父文件夹`Token`, 获取方式见[如何获取云文档资源相关 token](https://open.feishu.cn/document/ukTMukTMukTM/uczNzUjL3czM14yN3MTN#08bb5df6), 示例值: "fldbc5qgwyQnO0uedNllWuF3fAd"
	ReferEntity *CreateDriveFileShortcutReqReferEntity `json:"refer_entity,omitempty"` // 源文档的信息
}

// CreateDriveFileShortcutReqReferEntity ...
type CreateDriveFileShortcutReqReferEntity struct {
	ReferToken string `json:"refer_token,omitempty"` // 源文档的`Token`, 示例值: "doxbcGvhSVN0R6octqPwAEYNfFb"
	ReferType  string `json:"refer_type,omitempty"`  // 源文档的类型, 示例值: "docx", 可选值有: file: 文件, docx: 新版文档, bitable: 多维表格, doc: 旧版文档, sheet: 电子表格, mindnote: 思维笔记
}

// CreateDriveFileShortcutResp ...
type CreateDriveFileShortcutResp struct {
	SuccShortcutNode *CreateDriveFileShortcutRespSuccShortcutNode `json:"succ_shortcut_node,omitempty"` // 快捷方式
}

// CreateDriveFileShortcutRespSuccShortcutNode ...
type CreateDriveFileShortcutRespSuccShortcutNode struct {
	Token        string                                                   `json:"token,omitempty"`         // 文件`Token`
	Name         string                                                   `json:"name,omitempty"`          // 文件名
	Type         string                                                   `json:"type,omitempty"`          // 文件类型, 可选值参照请求体的`refer_type`
	ParentToken  string                                                   `json:"parent_token,omitempty"`  // 父文件夹`Token`
	URL          string                                                   `json:"url,omitempty"`           // 访问链接
	ShortcutInfo *CreateDriveFileShortcutRespSuccShortcutNodeShortcutInfo `json:"shortcut_info,omitempty"` // 快捷方式源文件信息
}

// CreateDriveFileShortcutRespSuccShortcutNodeShortcutInfo ...
type CreateDriveFileShortcutRespSuccShortcutNodeShortcutInfo struct {
	TargetType  string `json:"target_type,omitempty"`  // 快捷方式指向的源文件类型, 可选值参照请求体的`refer_type`
	TargetToken string `json:"target_token,omitempty"` // 快捷方式指向的源文件`Token`
}

// createDriveFileShortcutResp ...
type createDriveFileShortcutResp struct {
	Code int64                        `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                       `json:"msg,omitempty"`  // 错误描述
	Data *CreateDriveFileShortcutResp `json:"data,omitempty"`
}
