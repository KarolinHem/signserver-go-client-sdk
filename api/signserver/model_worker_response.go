/*
Copyright 2022 Keyfactor
Licensed under the Apache License, Version 2.0 (the "License"); you may
not use this file except in compliance with the License.  You may obtain a
copy of the License at http://www.apache.org/licenses/LICENSE-2.0.  Unless
required by applicable law or agreed to in writing, software distributed
under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES
OR CONDITIONS OF ANY KIND, either express or implied. See the License for
thespecific language governing permissions and limitations under the
License.

SignServer REST Interface

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package signserver

import (
	"encoding/json"
)

// checks if the WorkerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkerResponse{}

// WorkerResponse Represents a worker response.
type WorkerResponse struct {
	ResponseMessage      *string `json:"responseMessage,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkerResponse WorkerResponse

// NewWorkerResponse instantiates a new WorkerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkerResponse() *WorkerResponse {
	this := WorkerResponse{}
	return &this
}

// NewWorkerResponseWithDefaults instantiates a new WorkerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkerResponseWithDefaults() *WorkerResponse {
	this := WorkerResponse{}
	return &this
}

// GetResponseMessage returns the ResponseMessage field value if set, zero value otherwise.
func (o *WorkerResponse) GetResponseMessage() string {
	if o == nil || isNil(o.ResponseMessage) {
		var ret string
		return ret
	}
	return *o.ResponseMessage
}

// GetResponseMessageOk returns a tuple with the ResponseMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkerResponse) GetResponseMessageOk() (*string, bool) {
	if o == nil || isNil(o.ResponseMessage) {
		return nil, false
	}
	return o.ResponseMessage, true
}

// HasResponseMessage returns a boolean if a field has been set.
func (o *WorkerResponse) HasResponseMessage() bool {
	if o != nil && !isNil(o.ResponseMessage) {
		return true
	}

	return false
}

// SetResponseMessage gets a reference to the given string and assigns it to the ResponseMessage field.
func (o *WorkerResponse) SetResponseMessage(v string) {
	o.ResponseMessage = &v
}

func (o WorkerResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ResponseMessage) {
		toSerialize["responseMessage"] = o.ResponseMessage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WorkerResponse) UnmarshalJSON(bytes []byte) (err error) {
	varWorkerResponse := _WorkerResponse{}

	if err = json.Unmarshal(bytes, &varWorkerResponse); err == nil {
		*o = WorkerResponse(varWorkerResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "responseMessage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkerResponse struct {
	value *WorkerResponse
	isSet bool
}

func (v NullableWorkerResponse) Get() *WorkerResponse {
	return v.value
}

func (v *NullableWorkerResponse) Set(val *WorkerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkerResponse(val *WorkerResponse) *NullableWorkerResponse {
	return &NullableWorkerResponse{value: val, isSet: true}
}

func (v NullableWorkerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
