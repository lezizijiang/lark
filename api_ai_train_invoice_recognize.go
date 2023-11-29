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

// RecognizeAiTrainInvoice 火车票识别接口, 支持JPG/JPEG/PNG/PDF/OFD五种文件类型的一次性的识别。
//
// 单租户限流: 10QPS, 同租户下的应用没有限流, 共享本租户的 10QPS 限流
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/ai/document_ai-v1/train_invoice/recognize
func (r *AIService) RecognizeAiTrainInvoice(ctx context.Context, request *RecognizeAiTrainInvoiceReq, options ...MethodOptionFunc) (*RecognizeAiTrainInvoiceResp, *Response, error) {
	if r.cli.mock.mockAIRecognizeAiTrainInvoice != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] AI#RecognizeAiTrainInvoice mock enable")
		return r.cli.mock.mockAIRecognizeAiTrainInvoice(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "AI",
		API:                   "RecognizeAiTrainInvoice",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/document_ai/v1/train_invoice/recognize",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(recognizeAiTrainInvoiceResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockAIRecognizeAiTrainInvoice mock AIRecognizeAiTrainInvoice method
func (r *Mock) MockAIRecognizeAiTrainInvoice(f func(ctx context.Context, request *RecognizeAiTrainInvoiceReq, options ...MethodOptionFunc) (*RecognizeAiTrainInvoiceResp, *Response, error)) {
	r.mockAIRecognizeAiTrainInvoice = f
}

// UnMockAIRecognizeAiTrainInvoice un-mock AIRecognizeAiTrainInvoice method
func (r *Mock) UnMockAIRecognizeAiTrainInvoice() {
	r.mockAIRecognizeAiTrainInvoice = nil
}

// RecognizeAiTrainInvoiceReq ...
type RecognizeAiTrainInvoiceReq struct {
	File *RecognizeAiTrainInvoiceReqFile `json:"file,omitempty"` // 识别的火车票源文件, 示例值: file binary
}

// RecognizeAiTrainInvoiceResp ...
type RecognizeAiTrainInvoiceResp struct {
	TrainInvoices []*RecognizeAiTrainInvoiceRespTrainInvoice `json:"train_invoices,omitempty"` // 火车票信息
}

// RecognizeAiTrainInvoiceRespTrainInvoice ...
type RecognizeAiTrainInvoiceRespTrainInvoice struct {
	Entities []*RecognizeAiTrainInvoiceRespTrainInvoiceEntity `json:"entities,omitempty"` // 识别出的实体列表
}

// RecognizeAiTrainInvoiceRespTrainInvoiceEntity ...
type RecognizeAiTrainInvoiceRespTrainInvoiceEntity struct {
	Type  string `json:"type,omitempty"`  // 识别的字段种类, 可选值有: start_station: 出发站, end_station: 到达站, train_num: 车次编号, name: 火车票姓名, seat_num: 座位号, ticket_num: 车票编号, total_amount: 价格, time: 出发时间
	Value string `json:"value,omitempty"` // 识别出字段的文本信息
}

// recognizeAiTrainInvoiceResp ...
type recognizeAiTrainInvoiceResp struct {
	Code int64                        `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                       `json:"msg,omitempty"`  // 错误描述
	Data *RecognizeAiTrainInvoiceResp `json:"data,omitempty"`
}
