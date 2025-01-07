// Copyright 2021 Layotto Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sms

import (
	ref "mosn.io/layotto/components/ref"
)

// Config is the component's configuration
type Config struct {
	ref.Config
	Type     string            `json:"type"`
	Metadata map[string]string `json:"metadata"`
}

// SendSmsRequest is the request of the `SendSms` method.
type SendSmsWithTemplateRequest struct {
	// Required. The saas service name
	//  If your system uses multiple SMS services at the same time,
	//  you can specify which service to use with this field.
	ComponentName string `json:"component_name,omitempty"`
	// Required. The SMS receive phone numbers.
	PhoneNumbers []string `json:"phone_numbers,omitempty"`
	// Required.
	Template *Template `json:"template,omitempty"`
	// The registered sign name
	SignName string `json:"sign_name,omitempty"`
	// The SMS sender tag.
	SenderId string `json:"sender_id,omitempty"`
	// The metadata which will be sent to SMS components.
	Metadata map[string]string `json:"metadata,omitempty"`
}

// Sms template
type Template struct {
	// Required
	TemplateId string `json:"template_id,omitempty"`
	// Required
	TemplateParams map[string]string `json:"template_params,omitempty"`
}

// SendSmsResponse is the reponse of the `SendSms` method.
type SendSmsWithTemplateResponse struct {
	// The unique requestId.
	RequestId string `json:"request_id,omitempty"`
	// The status set of SMS
	Results []*SendStatus `json:"results,omitempty"`
}

// Status contains more information about the response
type SendStatus struct {
	// "OK" represents success.
	Code string `json:"code,omitempty"`
	// The error message.
	Message string `json:"message,omitempty"`
	// The send status metadata returned from SMS service.
	// Includes `PhoneNumber`.
	// `PhoneNumber`, is the phone number SMS send to. Supported by tencentcloud.
	Metadata map[string]string `json:"metadata,omitempty"`
}
