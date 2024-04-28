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

// CreateLingoEntity 通过此接口创建的词条, 无需经过词典管理员审核, 直接写入词库。因此, 调用此接口时, 应当慎重操作。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/lingo-v1/entity/create
func (r *LingoService) CreateLingoEntity(ctx context.Context, request *CreateLingoEntityReq, options ...MethodOptionFunc) (*CreateLingoEntityResp, *Response, error) {
	if r.cli.mock.mockLingoCreateLingoEntity != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Lingo#CreateLingoEntity mock enable")
		return r.cli.mock.mockLingoCreateLingoEntity(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Lingo",
		API:                   "CreateLingoEntity",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/lingo/v1/entities",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createLingoEntityResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockLingoCreateLingoEntity mock LingoCreateLingoEntity method
func (r *Mock) MockLingoCreateLingoEntity(f func(ctx context.Context, request *CreateLingoEntityReq, options ...MethodOptionFunc) (*CreateLingoEntityResp, *Response, error)) {
	r.mockLingoCreateLingoEntity = f
}

// UnMockLingoCreateLingoEntity un-mock LingoCreateLingoEntity method
func (r *Mock) UnMockLingoCreateLingoEntity() {
	r.mockLingoCreateLingoEntity = nil
}

// CreateLingoEntityReq ...
type CreateLingoEntityReq struct {
	RepoID      *string                          `query:"repo_id" json:"-"`      // 词库 ID（需要在指定词库创建词条时传入, 不传时默认创建至全员词库）, 如以应用身份创建词条到非全员词库, 需要在“词库设置”页面添加应用；若以用户身份创建词条到非全员词库, 该用户需要拥有对应词库的可见权限, 示例值: 71527909274113
	UserIDType  *IDType                          `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: open_id, 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	MainKeys    []*CreateLingoEntityReqMainKey   `json:"main_keys,omitempty"`    // 词条名, 最大长度: `1`
	Aliases     []*CreateLingoEntityReqAliase    `json:"aliases,omitempty"`      // 别名, 最大长度: `10`
	Description *string                          `json:"description,omitempty"`  // 纯文本格式词条释义。注: description 和 rich_text 至少有一个, 否则会报错: 1540001, 示例值: "词典是飞书提供的一款知识管理工具, 通过飞书词典可以帮助企业将分散的知识信息进行聚合, 并通过UGC的方式, 促进企业知识的保鲜和流通", 最大长度: `5000` 字符
	RelatedMeta *CreateLingoEntityReqRelatedMeta `json:"related_meta,omitempty"` // 词条相关信息
	OuterInfo   *CreateLingoEntityReqOuterInfo   `json:"outer_info,omitempty"`   // 外部系统关联数据
	RichText    *string                          `json:"rich_text,omitempty"`    // 富文本格式（当填写富文本内容时, description字段将会失效可不填写）, 支持的格式参考[飞书词典指南](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/baike-v1/overview)中的释义部分, 注意: 富文本格式至少需要包含一个 `<p>` 标签, 否则请求会报错, 示例值: "<b>加粗</b><i>斜体</i><p><a href=\"https://feishu.cn\">链接</a></p><p><span>词典是飞书提供的一款知识管理工具, 通过飞书词典可以帮助企业将分散的知识信息进行聚合, 并通过UGC的方式, 促进企业知识的保鲜和流通</span></p>", 最大长度: `5000` 字符
	I18nDescs   []*CreateLingoEntityReqI18nDesc  `json:"i18n_descs,omitempty"`   // 国际化的词条释义, 最大长度: `3`
}

// CreateLingoEntityReqAliase ...
type CreateLingoEntityReqAliase struct {
	Key           string                                   `json:"key,omitempty"`            // 名称, 示例值: "词典"
	DisplayStatus *CreateLingoEntityReqAliaseDisplayStatus `json:"display_status,omitempty"` // 展示状态
}

// CreateLingoEntityReqAliaseDisplayStatus ...
type CreateLingoEntityReqAliaseDisplayStatus struct {
	AllowHighlight bool `json:"allow_highlight,omitempty"` // 是否允许在 IM 和 Doc 等场景进行高亮提示, 示例值: true
	AllowSearch    bool `json:"allow_search,omitempty"`    // 是否允许在飞书中被搜索到, 示例值: true
}

// CreateLingoEntityReqI18nDesc ...
type CreateLingoEntityReqI18nDesc struct {
	Language    int64   `json:"language,omitempty"`    // 语言类型, 示例值: 1, 可选值有: 1: 中文, 2: 英文, 3: 日文
	Description *string `json:"description,omitempty"` // 纯文本释义, 示例值: "词典是飞书提供的一款知识管理工具, 通过飞书词典可以帮助企业将分散的知识信息进行聚合, 并通过UGC的方式, 促进企业知识的保鲜和流通", 长度范围: `1` ～ `5000` 字符
	RichText    *string `json:"rich_text,omitempty"`   // 富文本描述, 示例值: "<b>加粗</b><i>斜体</i><p><a href=\"https://feishu.cn\">链接</a></p><p><span>词典是飞书提供的一款知识管理工具, 通过飞书词典可以帮助企业将分散的知识信息进行聚合, 并通过UGC的方式, 促进企业知识的保鲜和流通</span></p>", 长度范围: `1` ～ `5000` 字符
}

// CreateLingoEntityReqMainKey ...
type CreateLingoEntityReqMainKey struct {
	Key           string                                    `json:"key,omitempty"`            // 名称, 示例值: "飞书词典"
	DisplayStatus *CreateLingoEntityReqMainKeyDisplayStatus `json:"display_status,omitempty"` // 展示状态
}

// CreateLingoEntityReqMainKeyDisplayStatus ...
type CreateLingoEntityReqMainKeyDisplayStatus struct {
	AllowHighlight bool `json:"allow_highlight,omitempty"` // 是否允许在 IM 和 Doc 等场景进行高亮提示, 示例值: true
	AllowSearch    bool `json:"allow_search,omitempty"`    // 是否允许在飞书中被搜索到, 示例值: true
}

// CreateLingoEntityReqOuterInfo ...
type CreateLingoEntityReqOuterInfo struct {
	Provider string `json:"provider,omitempty"` // 数据提供方（不能包含中横线 "-"）, 示例值: "星云", 长度范围: `2` ～ `32` 字符
	OuterID  string `json:"outer_id,omitempty"` // 词条在外部系统中对应的唯一 ID（不能包含中横线 "-"）, 示例值: "client_65326***7498d", 长度范围: `1` ～ `64` 字符
}

// CreateLingoEntityReqRelatedMeta ...
type CreateLingoEntityReqRelatedMeta struct {
	Users           []*CreateLingoEntityReqRelatedMetaUser           `json:"users,omitempty"`           // 相关联系人
	Chats           []*CreateLingoEntityReqRelatedMetaChat           `json:"chats,omitempty"`           // 有关的公开群
	Docs            []*CreateLingoEntityReqRelatedMetaDoc            `json:"docs,omitempty"`            // 飞书文档或飞书 wiki
	Oncalls         []*CreateLingoEntityReqRelatedMetaOncall         `json:"oncalls,omitempty"`         // 飞书值班号
	Links           []*CreateLingoEntityReqRelatedMetaLink           `json:"links,omitempty"`           // 其他网页链接
	Abbreviations   []*CreateLingoEntityReqRelatedMetaAbbreviation   `json:"abbreviations,omitempty"`   // 相关词条
	Classifications []*CreateLingoEntityReqRelatedMetaClassification `json:"classifications,omitempty"` // 当前词条所属分类, 词条只能属于二级分类, 且每个一级分类下只能选择一个二级分类。
	Images          []*CreateLingoEntityReqRelatedMetaImage          `json:"images,omitempty"`          // 上传的相关图片, 最大长度: `10`
}

// CreateLingoEntityReqRelatedMetaAbbreviation ...
type CreateLingoEntityReqRelatedMetaAbbreviation struct {
	ID *string `json:"id,omitempty"` // 其他相关词条 id, 示例值: "enterprise_5158***7960"
}

// CreateLingoEntityReqRelatedMetaChat ...
type CreateLingoEntityReqRelatedMetaChat struct {
	ID string `json:"id,omitempty"` // 公开群 id, 示例值: "oc_c13831833eaa8c92***cfa759ea4806"
}

// CreateLingoEntityReqRelatedMetaClassification ...
type CreateLingoEntityReqRelatedMetaClassification struct {
	ID       string  `json:"id,omitempty"`        // 二级分类 ID, 示例值: "704960692637761"
	FatherID *string `json:"father_id,omitempty"` // 对应一级分类 ID, 示例值: "704960692637777"
}

// CreateLingoEntityReqRelatedMetaDoc ...
type CreateLingoEntityReqRelatedMetaDoc struct {
	Title *string `json:"title,omitempty"` // 文档标题, 示例值: "快速了解飞书文档"
	URL   *string `json:"url,omitempty"`   // 文档 url, 示例值: "https://example.feishu.cn/docs/doccnxlVCs***sJE15I7PLAjIWc", 正则校验: `(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:.;]+[-A-Za-z0-9+&@#/%=~_|]`
}

// CreateLingoEntityReqRelatedMetaImage ...
type CreateLingoEntityReqRelatedMetaImage struct {
	Token string `json:"token,omitempty"` // 通过文件接口上传后的图片 token, 示例值: "boxbcEcmKiDia3e***WTpvdc7jc"
}

// CreateLingoEntityReqRelatedMetaLink ...
type CreateLingoEntityReqRelatedMetaLink struct {
	Title *string `json:"title,omitempty"` // 标题, 示例值: "飞书官网"
	URL   *string `json:"url,omitempty"`   // 网页链接, 示例值: "https://feishu.cn", 正则校验: `(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:.;]+[-A-Za-z0-9+&@#/%=~_|]`
}

// CreateLingoEntityReqRelatedMetaOncall ...
type CreateLingoEntityReqRelatedMetaOncall struct {
	ID string `json:"id,omitempty"` // 值班号 id, 示例值: "70236890***45548034"
}

// CreateLingoEntityReqRelatedMetaUser ...
type CreateLingoEntityReqRelatedMetaUser struct {
	ID    string  `json:"id,omitempty"`    // 格式根据 user_id_type 不同需要符合 open_id、user_id、union_id 格式的有效 id, 示例值: "ou_30b07b63089e***18789914dac63d36"
	Title *string `json:"title,omitempty"` // 备注, 示例值: "xxx 负责人"
}

// CreateLingoEntityResp ...
type CreateLingoEntityResp struct {
	Entity *CreateLingoEntityRespEntity `json:"entity,omitempty"` // 词条信息
}

// CreateLingoEntityRespEntity ...
type CreateLingoEntityRespEntity struct {
	ID          string                                  `json:"id,omitempty"`           // 词条 ID （需要更新某个词条时填写, 若是创建新词条可不填写）
	MainKeys    []*CreateLingoEntityRespEntityMainKey   `json:"main_keys,omitempty"`    // 词条名
	Aliases     []*CreateLingoEntityRespEntityAliase    `json:"aliases,omitempty"`      // 别名
	Description string                                  `json:"description,omitempty"`  // 纯文本格式词条释义。注: description 和 rich_text 至少有一个, 否则会报错: 1540001
	Creator     string                                  `json:"creator,omitempty"`      // 创建者
	CreateTime  string                                  `json:"create_time,omitempty"`  // 词条创建时间（秒级时间戳）
	Updater     string                                  `json:"updater,omitempty"`      // 最近一次更新者
	UpdateTime  string                                  `json:"update_time,omitempty"`  // 最近一次更新词条时间（秒级时间戳）
	RelatedMeta *CreateLingoEntityRespEntityRelatedMeta `json:"related_meta,omitempty"` // 相关数据
	Statistics  *CreateLingoEntityRespEntityStatistics  `json:"statistics,omitempty"`   // 当前词条收到的反馈数据
	OuterInfo   *CreateLingoEntityRespEntityOuterInfo   `json:"outer_info,omitempty"`   // 外部 id 关联数据
	RichText    string                                  `json:"rich_text,omitempty"`    // 富文本格式（当填写富文本内容时, description字段将会失效可不填写）, 支持的格式参考[飞书词典指南](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/baike-v1/overview)中的释义部分
	Source      int64                                   `json:"source,omitempty"`       // 词条的创建来源, 1: 用户主动创建, 2: 批量导入, 3: 官方词, 4: OpenAPI 创建
	I18nDescs   []*CreateLingoEntityRespEntityI18nDesc  `json:"i18n_descs,omitempty"`   // 国际化的词条释义
}

// CreateLingoEntityRespEntityAliase ...
type CreateLingoEntityRespEntityAliase struct {
	Key           string                                          `json:"key,omitempty"`            // 名称
	DisplayStatus *CreateLingoEntityRespEntityAliaseDisplayStatus `json:"display_status,omitempty"` // 展示状态
}

// CreateLingoEntityRespEntityAliaseDisplayStatus ...
type CreateLingoEntityRespEntityAliaseDisplayStatus struct {
	AllowHighlight bool `json:"allow_highlight,omitempty"` // 是否允许在 IM 和 Doc 等场景进行高亮提示
	AllowSearch    bool `json:"allow_search,omitempty"`    // 是否允许在飞书中被搜索到
}

// CreateLingoEntityRespEntityI18nDesc ...
type CreateLingoEntityRespEntityI18nDesc struct {
	Language    int64  `json:"language,omitempty"`    // 语言类型, 可选值有: 1: 中文, 2: 英文, 3: 日文
	Description string `json:"description,omitempty"` // 纯文本释义
	RichText    string `json:"rich_text,omitempty"`   // 富文本描述
}

// CreateLingoEntityRespEntityMainKey ...
type CreateLingoEntityRespEntityMainKey struct {
	Key           string                                           `json:"key,omitempty"`            // 名称
	DisplayStatus *CreateLingoEntityRespEntityMainKeyDisplayStatus `json:"display_status,omitempty"` // 展示状态
}

// CreateLingoEntityRespEntityMainKeyDisplayStatus ...
type CreateLingoEntityRespEntityMainKeyDisplayStatus struct {
	AllowHighlight bool `json:"allow_highlight,omitempty"` // 是否允许在 IM 和 Doc 等场景进行高亮提示
	AllowSearch    bool `json:"allow_search,omitempty"`    // 是否允许在飞书中被搜索到
}

// CreateLingoEntityRespEntityOuterInfo ...
type CreateLingoEntityRespEntityOuterInfo struct {
	Provider string `json:"provider,omitempty"` // 外部系统（不能包含中横线 "-"）
	OuterID  string `json:"outer_id,omitempty"` // 词条在外部系统中对应的唯一 ID（不能包含中横线 "-"）
}

// CreateLingoEntityRespEntityRelatedMeta ...
type CreateLingoEntityRespEntityRelatedMeta struct {
	Users           []*CreateLingoEntityRespEntityRelatedMetaUser           `json:"users,omitempty"`           // 关联用户信息
	Chats           []*CreateLingoEntityRespEntityRelatedMetaChat           `json:"chats,omitempty"`           // 关联群组信息
	Docs            []*CreateLingoEntityRespEntityRelatedMetaDoc            `json:"docs,omitempty"`            // 关联文档信息
	Oncalls         []*CreateLingoEntityRespEntityRelatedMetaOncall         `json:"oncalls,omitempty"`         // 相关服务中的相关值班号
	Links           []*CreateLingoEntityRespEntityRelatedMetaLink           `json:"links,omitempty"`           // 关联链接信息
	Abbreviations   []*CreateLingoEntityRespEntityRelatedMetaAbbreviation   `json:"abbreviations,omitempty"`   // 相关词条信息
	Classifications []*CreateLingoEntityRespEntityRelatedMetaClassification `json:"classifications,omitempty"` // 当前词条所属分类, 词条只能属于二级分类, 且每个一级分类下只能选择一个二级分类。
	Images          []*CreateLingoEntityRespEntityRelatedMetaImage          `json:"images,omitempty"`          // 上传的相关图片
}

// CreateLingoEntityRespEntityRelatedMetaAbbreviation ...
type CreateLingoEntityRespEntityRelatedMetaAbbreviation struct {
	ID string `json:"id,omitempty"` // 相关其他词条 id
}

// CreateLingoEntityRespEntityRelatedMetaChat ...
type CreateLingoEntityRespEntityRelatedMetaChat struct {
	ID    string `json:"id,omitempty"`    // 数据 id
	Title string `json:"title,omitempty"` // 标题
	URL   string `json:"url,omitempty"`   // 链接地址
}

// CreateLingoEntityRespEntityRelatedMetaClassification ...
type CreateLingoEntityRespEntityRelatedMetaClassification struct {
	ID       string `json:"id,omitempty"`        // 二级分类 ID
	FatherID string `json:"father_id,omitempty"` // 对应一级分类 ID
}

// CreateLingoEntityRespEntityRelatedMetaDoc ...
type CreateLingoEntityRespEntityRelatedMetaDoc struct {
	Title string `json:"title,omitempty"` // 标题
	URL   string `json:"url,omitempty"`   // 链接地址
}

// CreateLingoEntityRespEntityRelatedMetaImage ...
type CreateLingoEntityRespEntityRelatedMetaImage struct {
	Token string `json:"token,omitempty"` // 通过文件接口上传图片后, 获得的图片 token
}

// CreateLingoEntityRespEntityRelatedMetaLink ...
type CreateLingoEntityRespEntityRelatedMetaLink struct {
	Title string `json:"title,omitempty"` // 标题
	URL   string `json:"url,omitempty"`   // 链接地址
}

// CreateLingoEntityRespEntityRelatedMetaOncall ...
type CreateLingoEntityRespEntityRelatedMetaOncall struct {
	ID    string `json:"id,omitempty"`    // 数据 id
	Title string `json:"title,omitempty"` // 标题
	URL   string `json:"url,omitempty"`   // 链接地址
}

// CreateLingoEntityRespEntityRelatedMetaUser ...
type CreateLingoEntityRespEntityRelatedMetaUser struct {
	ID    string `json:"id,omitempty"`    // 数据 id
	Title string `json:"title,omitempty"` // 标题
	URL   string `json:"url,omitempty"`   // 链接地址
}

// CreateLingoEntityRespEntityStatistics ...
type CreateLingoEntityRespEntityStatistics struct {
	LikeCount    int64 `json:"like_count,omitempty"`    // 点赞数量
	DislikeCount int64 `json:"dislike_count,omitempty"` // 点踩数量
}

// createLingoEntityResp ...
type createLingoEntityResp struct {
	Code  int64                  `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                 `json:"msg,omitempty"`  // 错误描述
	Data  *CreateLingoEntityResp `json:"data,omitempty"`
	Error *ErrorDetail           `json:"error,omitempty"`
}
