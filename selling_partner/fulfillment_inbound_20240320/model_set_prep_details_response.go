package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the SetPrepDetailsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetPrepDetailsResponse{}

// SetPrepDetailsResponse The `setPrepDetails` response.
type SetPrepDetailsResponse struct {
	// UUID for the given operation.
	OperationId string `json:"operationId" validate:"regexp=^[a-zA-Z0-9-]*$"`
}

type _SetPrepDetailsResponse SetPrepDetailsResponse

// NewSetPrepDetailsResponse instantiates a new SetPrepDetailsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetPrepDetailsResponse(operationId string) *SetPrepDetailsResponse {
	this := SetPrepDetailsResponse{}
	this.OperationId = operationId
	return &this
}

// NewSetPrepDetailsResponseWithDefaults instantiates a new SetPrepDetailsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetPrepDetailsResponseWithDefaults() *SetPrepDetailsResponse {
	this := SetPrepDetailsResponse{}
	return &this
}

// GetOperationId returns the OperationId field value
func (o *SetPrepDetailsResponse) GetOperationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperationId
}

// GetOperationIdOk returns a tuple with the OperationId field value
// and a boolean to check if the value has been set.
func (o *SetPrepDetailsResponse) GetOperationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationId, true
}

// SetOperationId sets field value
func (o *SetPrepDetailsResponse) SetOperationId(v string) {
	o.OperationId = v
}

func (o SetPrepDetailsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operationId"] = o.OperationId
	return toSerialize, nil
}

type NullableSetPrepDetailsResponse struct {
	value *SetPrepDetailsResponse
	isSet bool
}

func (v NullableSetPrepDetailsResponse) Get() *SetPrepDetailsResponse {
	return v.value
}

func (v *NullableSetPrepDetailsResponse) Set(val *SetPrepDetailsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSetPrepDetailsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSetPrepDetailsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetPrepDetailsResponse(val *SetPrepDetailsResponse) *NullableSetPrepDetailsResponse {
	return &NullableSetPrepDetailsResponse{value: val, isSet: true}
}

func (v NullableSetPrepDetailsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSetPrepDetailsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
