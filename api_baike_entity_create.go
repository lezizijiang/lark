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

// CreateBaikeEntity 通过此接口创建的词条, 无需经过词典管理员审核, 直接写入词库。因此, 调用此接口时, 应当慎重操作。
//
// 为了更好地提升接口文档的的易理解性, 我们对文档进行了升级, 请尽快迁移至[新版本>>](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/lingo-v1/entity/create)
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/baike-v1/entity/create
// new doc: https://open.feishu.cn/document/server-docs/baike-v1/entity/create
//
// Deprecated
func (r *BaikeService) CreateBaikeEntity(ctx context.Context, request *CreateBaikeEntityReq, options ...MethodOptionFunc) (*CreateBaikeEntityResp, *Response, error) {
	if r.cli.mock.mockBaikeCreateBaikeEntity != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Baike#CreateBaikeEntity mock enable")
		return r.cli.mock.mockBaikeCreateBaikeEntity(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Baike",
		API:                   "CreateBaikeEntity",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/baike/v1/entities",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(createBaikeEntityResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockBaikeCreateBaikeEntity mock BaikeCreateBaikeEntity method
func (r *Mock) MockBaikeCreateBaikeEntity(f func(ctx context.Context, request *CreateBaikeEntityReq, options ...MethodOptionFunc) (*CreateBaikeEntityResp, *Response, error)) {
	r.mockBaikeCreateBaikeEntity = f
}

// UnMockBaikeCreateBaikeEntity un-mock BaikeCreateBaikeEntity method
func (r *Mock) UnMockBaikeCreateBaikeEntity() {
	r.mockBaikeCreateBaikeEntity = nil
}

// CreateBaikeEntityReq ...
type CreateBaikeEntityReq struct {
	UserIDType  *IDType                          `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	MainKeys    []*CreateBaikeEntityReqMainKey   `json:"main_keys,omitempty"`    // 词条名, 最大长度: `1`
	Aliases     []*CreateBaikeEntityReqAliase    `json:"aliases,omitempty"`      // 别名, 最大长度: `10`
	Description *string                          `json:"description,omitempty"`  // 纯文本格式词条释义。注: description 和 rich_text 至少有一个, 否则会报错: 1540001, 示例值: "词典是飞书提供的一款知识管理工具, 通过飞书词典可以帮助企业将分散的知识信息进行聚合, 并通过UGC的方式, 促进企业知识的保鲜和流通", 最大长度: `5000` 字符
	RelatedMeta *CreateBaikeEntityReqRelatedMeta `json:"related_meta,omitempty"` // 更多相关信息
	OuterInfo   *CreateBaikeEntityReqOuterInfo   `json:"outer_info,omitempty"`   // 外部系统关联数据
	RichText    *string                          `json:"rich_text,omitempty"`    // 富文本格式（当填写富文本内容时, description字段将会失效可不填写）, 支持的格式参考[飞书词典指南](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/baike-v1/overview)中的释义部分, 注意: 富文本格式至少需要包含一个 `<p>` 标签, 否则请求会报错, 示例值: "<b>加粗</b><i>斜体</i><p><a href=\"https://feishu.cn\">链接</a></p><p><span>词典是飞书提供的一款知识管理工具, 通过飞书词典可以帮助企业将分散的知识信息进行聚合, 并通过UGC的方式, 促进企业知识的保鲜和流通</span></p>", 最大长度: `5000` 字符
}

// CreateBaikeEntityReqAliase ...
type CreateBaikeEntityReqAliase struct {
	Key           string                                   `json:"key,omitempty"`            // 名称的值, 示例值: "飞书词典"
	DisplayStatus *CreateBaikeEntityReqAliaseDisplayStatus `json:"display_status,omitempty"` // 名称展示范围
}

// CreateBaikeEntityReqAliaseDisplayStatus ...
type CreateBaikeEntityReqAliaseDisplayStatus struct {
	AllowHighlight bool `json:"allow_highlight,omitempty"` // 对应名称是否在消息/云文档高亮, 示例值: true
	AllowSearch    bool `json:"allow_search,omitempty"`    // 对应名称是否在搜索结果中展示, 示例值: true
}

// CreateBaikeEntityReqMainKey ...
type CreateBaikeEntityReqMainKey struct {
	Key           string                                    `json:"key,omitempty"`            // 名称的值, 示例值: "飞书词典"
	DisplayStatus *CreateBaikeEntityReqMainKeyDisplayStatus `json:"display_status,omitempty"` // 名称展示范围
}

// CreateBaikeEntityReqMainKeyDisplayStatus ...
type CreateBaikeEntityReqMainKeyDisplayStatus struct {
	AllowHighlight bool `json:"allow_highlight,omitempty"` // 对应名称是否在消息/云文档高亮, 示例值: true
	AllowSearch    bool `json:"allow_search,omitempty"`    // 对应名称是否在搜索结果中展示, 示例值: true
}

// CreateBaikeEntityReqOuterInfo ...
type CreateBaikeEntityReqOuterInfo struct {
	Provider string `json:"provider,omitempty"` // 外部系统（不能包含中横线 "-"）, 示例值: "星云", 长度范围: `2` ～ `32` 字符
	OuterID  string `json:"outer_id,omitempty"` // 词条在外部系统中对应的唯一 ID（不能包含中横线 "-"）, 示例值: "client_653267498d", 长度范围: `1` ～ `64` 字符
}

// CreateBaikeEntityReqRelatedMeta ...
type CreateBaikeEntityReqRelatedMeta struct {
	Users           []*CreateBaikeEntityReqRelatedMetaUser           `json:"users,omitempty"`           // 相关联系人
	Chats           []*CreateBaikeEntityReqRelatedMetaChat           `json:"chats,omitempty"`           // 相关服务中的相关公开群
	Docs            []*CreateBaikeEntityReqRelatedMetaDoc            `json:"docs,omitempty"`            // 相关云文档
	Oncalls         []*CreateBaikeEntityReqRelatedMetaOncall         `json:"oncalls,omitempty"`         // 相关服务中的相关值班号
	Links           []*CreateBaikeEntityReqRelatedMetaLink           `json:"links,omitempty"`           // 相关链接
	Abbreviations   []*CreateBaikeEntityReqRelatedMetaAbbreviation   `json:"abbreviations,omitempty"`   // 相关词条
	Classifications []*CreateBaikeEntityReqRelatedMetaClassification `json:"classifications,omitempty"` // 当前词条所属分类, 词条只能属于二级分类, 且每个一级分类下只能选择一个二级分类。
	Images          []*CreateBaikeEntityReqRelatedMetaImage          `json:"images,omitempty"`          // 上传的图片, 最大长度: `10`
}

// CreateBaikeEntityReqRelatedMetaAbbreviation ...
type CreateBaikeEntityReqRelatedMetaAbbreviation struct {
	ID *string `json:"id,omitempty"` // 相关词条 ID, 示例值: "enterprise_51527260"
}

// CreateBaikeEntityReqRelatedMetaChat ...
type CreateBaikeEntityReqRelatedMetaChat struct {
	ID string `json:"id,omitempty"` // 对应相关信息 ID, 示例值: "格式请看请求体示例"
}

// CreateBaikeEntityReqRelatedMetaClassification ...
type CreateBaikeEntityReqRelatedMetaClassification struct {
	ID       string  `json:"id,omitempty"`        // 二级分类 ID, 示例值: "704960692637761"
	FatherID *string `json:"father_id,omitempty"` // 对应一级分类 ID, 示例值: "704960692637777"
}

// CreateBaikeEntityReqRelatedMetaDoc ...
type CreateBaikeEntityReqRelatedMetaDoc struct {
	Title *string `json:"title,omitempty"` // 对应相关信息的描述, 如相关联系人的描述、相关链接的标题, 示例值: "飞书词典帮助中心"
	URL   *string `json:"url,omitempty"`   // 链接地址, 示例值: "https://www.feishu.cn/hc/zh-CN", 正则校验: `(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:.;]+[-A-Za-z0-9+&@#/%=~_|]`
}

// CreateBaikeEntityReqRelatedMetaImage ...
type CreateBaikeEntityReqRelatedMetaImage struct {
	Token string `json:"token,omitempty"` // 通过文件接口上传图片后, 获得的图片 token, 示例值: "boxbcEcmKiDia3evgqWTpvdc7jc"
}

// CreateBaikeEntityReqRelatedMetaLink ...
type CreateBaikeEntityReqRelatedMetaLink struct {
	Title *string `json:"title,omitempty"` // 对应相关信息的描述, 如相关联系人的描述、相关链接的标题, 示例值: "飞书词典帮助中心"
	URL   *string `json:"url,omitempty"`   // 链接地址, 示例值: "https://www.feishu.cn/hc/zh-CN", 正则校验: `(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:.;]+[-A-Za-z0-9+&@#/%=~_|]`
}

// CreateBaikeEntityReqRelatedMetaOncall ...
type CreateBaikeEntityReqRelatedMetaOncall struct {
	ID string `json:"id,omitempty"` // 对应相关信息 ID, 示例值: "格式请看请求体示例"
}

// CreateBaikeEntityReqRelatedMetaUser ...
type CreateBaikeEntityReqRelatedMetaUser struct {
	ID    string  `json:"id,omitempty"`    // 对应相关信息 ID, 示例值: "格式请看请求体示例"
	Title *string `json:"title,omitempty"` // 对应相关信息的描述, 如相关联系人的描述、相关链接的标题, 示例值: "飞书词典帮助中心"
}

// CreateBaikeEntityResp ...
type CreateBaikeEntityResp struct {
	Entity *CreateBaikeEntityRespEntity `json:"entity,omitempty"` // 词条信息
}

// CreateBaikeEntityRespEntity ...
type CreateBaikeEntityRespEntity struct {
	ID          string                                  `json:"id,omitempty"`           // 词条 ID （需要更新某个词条时填写, 若是创建新词条可不填写）
	MainKeys    []*CreateBaikeEntityRespEntityMainKey   `json:"main_keys,omitempty"`    // 词条名
	Aliases     []*CreateBaikeEntityRespEntityAliase    `json:"aliases,omitempty"`      // 别名
	Description string                                  `json:"description,omitempty"`  // 纯文本格式词条释义。注: description 和 rich_text 至少有一个, 否则会报错: 1540001
	CreateTime  string                                  `json:"create_time,omitempty"`  // 词条创建时间
	UpdateTime  string                                  `json:"update_time,omitempty"`  // 词条最近更新时间
	RelatedMeta *CreateBaikeEntityRespEntityRelatedMeta `json:"related_meta,omitempty"` // 更多相关信息
	Statistics  *CreateBaikeEntityRespEntityStatistics  `json:"statistics,omitempty"`   // 当前词条收到的反馈数据
	OuterInfo   *CreateBaikeEntityRespEntityOuterInfo   `json:"outer_info,omitempty"`   // 外部系统关联数据
	RichText    string                                  `json:"rich_text,omitempty"`    // 富文本格式（当填写富文本内容时, description字段将会失效可不填写）, 支持的格式参考[飞书词典指南](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/baike-v1/overview)中的释义部分
}

// CreateBaikeEntityRespEntityAliase ...
type CreateBaikeEntityRespEntityAliase struct {
	Key           string                                          `json:"key,omitempty"`            // 名称的值
	DisplayStatus *CreateBaikeEntityRespEntityAliaseDisplayStatus `json:"display_status,omitempty"` // 名称展示范围
}

// CreateBaikeEntityRespEntityAliaseDisplayStatus ...
type CreateBaikeEntityRespEntityAliaseDisplayStatus struct {
	AllowHighlight bool `json:"allow_highlight,omitempty"` // 对应名称是否在消息/云文档高亮
	AllowSearch    bool `json:"allow_search,omitempty"`    // 对应名称是否在搜索结果中展示
}

// CreateBaikeEntityRespEntityMainKey ...
type CreateBaikeEntityRespEntityMainKey struct {
	Key           string                                           `json:"key,omitempty"`            // 名称的值
	DisplayStatus *CreateBaikeEntityRespEntityMainKeyDisplayStatus `json:"display_status,omitempty"` // 名称展示范围
}

// CreateBaikeEntityRespEntityMainKeyDisplayStatus ...
type CreateBaikeEntityRespEntityMainKeyDisplayStatus struct {
	AllowHighlight bool `json:"allow_highlight,omitempty"` // 对应名称是否在消息/云文档高亮
	AllowSearch    bool `json:"allow_search,omitempty"`    // 对应名称是否在搜索结果中展示
}

// CreateBaikeEntityRespEntityOuterInfo ...
type CreateBaikeEntityRespEntityOuterInfo struct {
	Provider string `json:"provider,omitempty"` // 外部系统（不能包含中横线 "-"）
	OuterID  string `json:"outer_id,omitempty"` // 词条在外部系统中对应的唯一 ID（不能包含中横线 "-"）
}

// CreateBaikeEntityRespEntityRelatedMeta ...
type CreateBaikeEntityRespEntityRelatedMeta struct {
	Users           []*CreateBaikeEntityRespEntityRelatedMetaUser           `json:"users,omitempty"`           // 相关联系人
	Chats           []*CreateBaikeEntityRespEntityRelatedMetaChat           `json:"chats,omitempty"`           // 相关服务中的相关公开群
	Docs            []*CreateBaikeEntityRespEntityRelatedMetaDoc            `json:"docs,omitempty"`            // 相关云文档
	Oncalls         []*CreateBaikeEntityRespEntityRelatedMetaOncall         `json:"oncalls,omitempty"`         // 相关服务中的相关值班号
	Links           []*CreateBaikeEntityRespEntityRelatedMetaLink           `json:"links,omitempty"`           // 相关链接
	Abbreviations   []*CreateBaikeEntityRespEntityRelatedMetaAbbreviation   `json:"abbreviations,omitempty"`   // 相关词条
	Classifications []*CreateBaikeEntityRespEntityRelatedMetaClassification `json:"classifications,omitempty"` // 当前词条所属分类, 词条只能属于二级分类, 且每个一级分类下只能选择一个二级分类。
	Images          []*CreateBaikeEntityRespEntityRelatedMetaImage          `json:"images,omitempty"`          // 上传的图片
}

// CreateBaikeEntityRespEntityRelatedMetaAbbreviation ...
type CreateBaikeEntityRespEntityRelatedMetaAbbreviation struct {
	ID string `json:"id,omitempty"` // 相关词条 ID
}

// CreateBaikeEntityRespEntityRelatedMetaChat ...
type CreateBaikeEntityRespEntityRelatedMetaChat struct {
	ID    string `json:"id,omitempty"`    // 对应相关信息 ID
	Title string `json:"title,omitempty"` // 对应相关信息的描述, 如相关联系人的描述、相关链接的标题
	URL   string `json:"url,omitempty"`   // 链接地址
}

// CreateBaikeEntityRespEntityRelatedMetaClassification ...
type CreateBaikeEntityRespEntityRelatedMetaClassification struct {
	ID       string `json:"id,omitempty"`        // 二级分类 ID
	Name     string `json:"name,omitempty"`      // 二级分类名称
	FatherID string `json:"father_id,omitempty"` // 对应一级分类 ID
}

// CreateBaikeEntityRespEntityRelatedMetaDoc ...
type CreateBaikeEntityRespEntityRelatedMetaDoc struct {
	Title string `json:"title,omitempty"` // 对应相关信息的描述, 如相关联系人的描述、相关链接的标题
	URL   string `json:"url,omitempty"`   // 链接地址
}

// CreateBaikeEntityRespEntityRelatedMetaImage ...
type CreateBaikeEntityRespEntityRelatedMetaImage struct {
	Token string `json:"token,omitempty"` // 通过文件接口上传图片后, 获得的图片 token
}

// CreateBaikeEntityRespEntityRelatedMetaLink ...
type CreateBaikeEntityRespEntityRelatedMetaLink struct {
	Title string `json:"title,omitempty"` // 对应相关信息的描述, 如相关联系人的描述、相关链接的标题
	URL   string `json:"url,omitempty"`   // 链接地址
}

// CreateBaikeEntityRespEntityRelatedMetaOncall ...
type CreateBaikeEntityRespEntityRelatedMetaOncall struct {
	ID    string `json:"id,omitempty"`    // 对应相关信息 ID
	Title string `json:"title,omitempty"` // 对应相关信息的描述, 如相关联系人的描述、相关链接的标题
	URL   string `json:"url,omitempty"`   // 链接地址
}

// CreateBaikeEntityRespEntityRelatedMetaUser ...
type CreateBaikeEntityRespEntityRelatedMetaUser struct {
	ID    string `json:"id,omitempty"`    // 对应相关信息 ID
	Title string `json:"title,omitempty"` // 对应相关信息的描述, 如相关联系人的描述、相关链接的标题
	URL   string `json:"url,omitempty"`   // 链接地址
}

// CreateBaikeEntityRespEntityStatistics ...
type CreateBaikeEntityRespEntityStatistics struct {
	LikeCount    int64 `json:"like_count,omitempty"`    // 累计点赞
	DislikeCount int64 `json:"dislike_count,omitempty"` // 当前词条版本收到的负反馈数量
}

// createBaikeEntityResp ...
type createBaikeEntityResp struct {
	Code  int64                  `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                 `json:"msg,omitempty"`  // 错误描述
	Data  *CreateBaikeEntityResp `json:"data,omitempty"`
	Error *ErrorDetail           `json:"error,omitempty"`
}
