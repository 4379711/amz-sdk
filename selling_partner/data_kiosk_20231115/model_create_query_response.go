package data_kiosk_20231115

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateQueryResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateQueryResponse{}

// CreateQueryResponse The response for the `createQuery` operation.
type CreateQueryResponse struct {
	// The identifier for the query. This identifier is unique only in combination with a selling partner account ID.
	QueryId string `json:"queryId"`
}

type _CreateQueryResponse CreateQueryResponse

// NewCreateQueryResponse instantiates a new CreateQueryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateQueryResponse(queryId string) *CreateQueryResponse {
	this := CreateQueryResponse{}
	this.QueryId = queryId
	return &this
}

// NewCreateQueryResponseWithDefaults instantiates a new CreateQueryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateQueryResponseWithDefaults() *CreateQueryResponse {
	this := CreateQueryResponse{}
	return &this
}

// GetQueryId returns the QueryId field value
func (o *CreateQueryResponse) GetQueryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryId
}

// GetQueryIdOk returns a tuple with the QueryId field value
// and a boolean to check if the value has been set.
func (o *CreateQueryResponse) GetQueryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryId, true
}

// SetQueryId sets field value
func (o *CreateQueryResponse) SetQueryId(v string) {
	o.QueryId = v
}

func (o CreateQueryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["queryId"] = o.QueryId
	return toSerialize, nil
}

type NullableCreateQueryResponse struct {
	value *CreateQueryResponse
	isSet bool
}

func (v NullableCreateQueryResponse) Get() *CreateQueryResponse {
	return v.value
}

func (v *NullableCreateQueryResponse) Set(val *CreateQueryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateQueryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateQueryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateQueryResponse(val *CreateQueryResponse) *NullableCreateQueryResponse {
	return &NullableCreateQueryResponse{value: val, isSet: true}
}

func (v NullableCreateQueryResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateQueryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
