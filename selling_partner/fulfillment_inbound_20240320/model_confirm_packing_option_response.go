package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the ConfirmPackingOptionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfirmPackingOptionResponse{}

// ConfirmPackingOptionResponse The `confirmPackingOption` response.
type ConfirmPackingOptionResponse struct {
	// UUID for the given operation.
	OperationId string `json:"operationId" validate:"regexp=^[a-zA-Z0-9-]*$"`
}

type _ConfirmPackingOptionResponse ConfirmPackingOptionResponse

// NewConfirmPackingOptionResponse instantiates a new ConfirmPackingOptionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfirmPackingOptionResponse(operationId string) *ConfirmPackingOptionResponse {
	this := ConfirmPackingOptionResponse{}
	this.OperationId = operationId
	return &this
}

// NewConfirmPackingOptionResponseWithDefaults instantiates a new ConfirmPackingOptionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfirmPackingOptionResponseWithDefaults() *ConfirmPackingOptionResponse {
	this := ConfirmPackingOptionResponse{}
	return &this
}

// GetOperationId returns the OperationId field value
func (o *ConfirmPackingOptionResponse) GetOperationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperationId
}

// GetOperationIdOk returns a tuple with the OperationId field value
// and a boolean to check if the value has been set.
func (o *ConfirmPackingOptionResponse) GetOperationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationId, true
}

// SetOperationId sets field value
func (o *ConfirmPackingOptionResponse) SetOperationId(v string) {
	o.OperationId = v
}

func (o ConfirmPackingOptionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operationId"] = o.OperationId
	return toSerialize, nil
}

type NullableConfirmPackingOptionResponse struct {
	value *ConfirmPackingOptionResponse
	isSet bool
}

func (v NullableConfirmPackingOptionResponse) Get() *ConfirmPackingOptionResponse {
	return v.value
}

func (v *NullableConfirmPackingOptionResponse) Set(val *ConfirmPackingOptionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConfirmPackingOptionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConfirmPackingOptionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfirmPackingOptionResponse(val *ConfirmPackingOptionResponse) *NullableConfirmPackingOptionResponse {
	return &NullableConfirmPackingOptionResponse{value: val, isSet: true}
}

func (v NullableConfirmPackingOptionResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableConfirmPackingOptionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
