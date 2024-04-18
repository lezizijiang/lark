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

// InstantReminderAppFeedCard 即时提醒能力是飞书在消息列表中提供的强提醒能力, 当有重要通知或任务需要及时触达用户, 可将群组或机器人对话在消息列表中置顶展示, 打开飞书首页即可处理重要任务。
//
// doc: https://open.larkoffice.com/document/uAjLw4CM/ukTMukTMukTM/group/im-v2/feed_card/patch
func (r *MessageService) InstantReminderAppFeedCard(ctx context.Context, request *InstantReminderAppFeedCardReq, options ...MethodOptionFunc) (*InstantReminderAppFeedCardResp, *Response, error) {
	if r.cli.mock.mockMessageInstantReminderAppFeedCard != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Message#InstantReminderAppFeedCard mock enable")
		return r.cli.mock.mockMessageInstantReminderAppFeedCard(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Message",
		API:                   "InstantReminderAppFeedCard",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v2/feed_cards/:feed_card_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(instantReminderAppFeedCardResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockMessageInstantReminderAppFeedCard mock MessageInstantReminderAppFeedCard method
func (r *Mock) MockMessageInstantReminderAppFeedCard(f func(ctx context.Context, request *InstantReminderAppFeedCardReq, options ...MethodOptionFunc) (*InstantReminderAppFeedCardResp, *Response, error)) {
	r.mockMessageInstantReminderAppFeedCard = f
}

// UnMockMessageInstantReminderAppFeedCard un-mock MessageInstantReminderAppFeedCard method
func (r *Mock) UnMockMessageInstantReminderAppFeedCard() {
	r.mockMessageInstantReminderAppFeedCard = nil
}

// InstantReminderAppFeedCardReq ...
type InstantReminderAppFeedCardReq struct {
	FeedCardID    string   `path:"feed_card_id" json:"-"`    // 群id, 当前只支持群聊, 示例值: "oc_679eaeb583654bff73fefcc6e6371370"
	UserIDType    IDType   `query:"user_id_type" json:"-"`   // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	TimeSensitive bool     `json:"time_sensitive,omitempty"` // 即时提醒状态, true-打开, false-关闭, 示例值: true
	UserIDs       []string `json:"user_ids,omitempty"`       // 用户id 列表, 支持OpenID, UnionID, UserID, 示例值: ["ou_9d2beb613c85a2412862a49a924558c5"]
}

// InstantReminderAppFeedCardResp ...
type InstantReminderAppFeedCardResp struct {
	FailedUserReasons []*InstantReminderAppFeedCardRespFailedUserReason `json:"failed_user_reasons,omitempty"` // 失败原因
}

// InstantReminderAppFeedCardRespFailedUserReason ...
type InstantReminderAppFeedCardRespFailedUserReason struct {
	ErrorCode    int64  `json:"error_code,omitempty"`    // 错误码
	ErrorMessage string `json:"error_message,omitempty"` // 错误信息
	UserID       string `json:"user_id,omitempty"`       // 用户id
}

// instantReminderAppFeedCardResp ...
type instantReminderAppFeedCardResp struct {
	Code  int64                           `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                          `json:"msg,omitempty"`  // 错误描述
	Data  *InstantReminderAppFeedCardResp `json:"data,omitempty"`
	Error *ErrorDetail                    `json:"error,omitempty"`
}
