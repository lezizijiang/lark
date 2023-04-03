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

// UpdateVCReserveConfigAdmin 更新会议室预定管理员。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/vc-v1/reserve_config-admin/patch
func (r *VCService) UpdateVCReserveConfigAdmin(ctx context.Context, request *UpdateVCReserveConfigAdminReq, options ...MethodOptionFunc) (*UpdateVCReserveConfigAdminResp, *Response, error) {
	if r.cli.mock.mockVCUpdateVCReserveConfigAdmin != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] VC#UpdateVCReserveConfigAdmin mock enable")
		return r.cli.mock.mockVCUpdateVCReserveConfigAdmin(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "VC",
		API:                   "UpdateVCReserveConfigAdmin",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/vc/v1/reserve_configs/:reserve_config_id/admin",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(updateVCReserveConfigAdminResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockVCUpdateVCReserveConfigAdmin mock VCUpdateVCReserveConfigAdmin method
func (r *Mock) MockVCUpdateVCReserveConfigAdmin(f func(ctx context.Context, request *UpdateVCReserveConfigAdminReq, options ...MethodOptionFunc) (*UpdateVCReserveConfigAdminResp, *Response, error)) {
	r.mockVCUpdateVCReserveConfigAdmin = f
}

// UnMockVCUpdateVCReserveConfigAdmin un-mock VCUpdateVCReserveConfigAdmin method
func (r *Mock) UnMockVCUpdateVCReserveConfigAdmin() {
	r.mockVCUpdateVCReserveConfigAdmin = nil
}

// UpdateVCReserveConfigAdminReq ...
type UpdateVCReserveConfigAdminReq struct {
	ReserveConfigID    string                                           `path:"reserve_config_id" json:"-"`     // 会议室或层级id, 示例值: "omm_3c5dd7e09bac0c1758fcf9511bd1a771"
	UserIDType         *IDType                                          `query:"user_id_type" json:"-"`         // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	ScopeType          int64                                            `json:"scope_type,omitempty"`           // 1代表层级, 2代表会议室, 示例值: 2, 取值范围: `1` ～ `2`
	ReserveAdminConfig *UpdateVCReserveConfigAdminReqReserveAdminConfig `json:"reserve_admin_config,omitempty"` // 预定管理员或部门
}

// UpdateVCReserveConfigAdminReqReserveAdminConfig ...
type UpdateVCReserveConfigAdminReqReserveAdminConfig struct {
	Depts []*UpdateVCReserveConfigAdminReqReserveAdminConfigDept `json:"depts,omitempty"` // 预定管理部门
	Users []*UpdateVCReserveConfigAdminReqReserveAdminConfigUser `json:"users,omitempty"` // 预定管理员
}

// UpdateVCReserveConfigAdminReqReserveAdminConfigDept ...
type UpdateVCReserveConfigAdminReqReserveAdminConfigDept struct {
	DepartmentID string `json:"department_id,omitempty"` // 预定管理部门ID, 使用open_department_id, 示例值: "od-47d8b570b0a011e9679a755efcc5f61a"
}

// UpdateVCReserveConfigAdminReqReserveAdminConfigUser ...
type UpdateVCReserveConfigAdminReqReserveAdminConfigUser struct {
	UserID string `json:"user_id,omitempty"` // 预定管理员ID, 示例值: "ou_a27b07a9071d90577c0177bcec98f856"
}

// UpdateVCReserveConfigAdminResp ...
type UpdateVCReserveConfigAdminResp struct {
}

// updateVCReserveConfigAdminResp ...
type updateVCReserveConfigAdminResp struct {
	Code int64                           `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                          `json:"msg,omitempty"`  // 错误描述
	Data *UpdateVCReserveConfigAdminResp `json:"data,omitempty"`
}
