package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the ConfirmPlacementOptionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfirmPlacementOptionResponse{}

// ConfirmPlacementOptionResponse The `confirmPlacementOption` response.
type ConfirmPlacementOptionResponse struct {
	// UUID for the given operation.
	OperationId string `json:"operationId" validate:"regexp=^[a-zA-Z0-9-]*$"`
}

type _ConfirmPlacementOptionResponse ConfirmPlacementOptionResponse

// NewConfirmPlacementOptionResponse instantiates a new ConfirmPlacementOptionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfirmPlacementOptionResponse(operationId string) *ConfirmPlacementOptionResponse {
	this := ConfirmPlacementOptionResponse{}
	this.OperationId = operationId
	return &this
}

// NewConfirmPlacementOptionResponseWithDefaults instantiates a new ConfirmPlacementOptionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfirmPlacementOptionResponseWithDefaults() *ConfirmPlacementOptionResponse {
	this := ConfirmPlacementOptionResponse{}
	return &this
}

// GetOperationId returns the OperationId field value
func (o *ConfirmPlacementOptionResponse) GetOperationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperationId
}

// GetOperationIdOk returns a tuple with the OperationId field value
// and a boolean to check if the value has been set.
func (o *ConfirmPlacementOptionResponse) GetOperationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationId, true
}

// SetOperationId sets field value
func (o *ConfirmPlacementOptionResponse) SetOperationId(v string) {
	o.OperationId = v
}

func (o ConfirmPlacementOptionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operationId"] = o.OperationId
	return toSerialize, nil
}

type NullableConfirmPlacementOptionResponse struct {
	value *ConfirmPlacementOptionResponse
	isSet bool
}

func (v NullableConfirmPlacementOptionResponse) Get() *ConfirmPlacementOptionResponse {
	return v.value
}

func (v *NullableConfirmPlacementOptionResponse) Set(val *ConfirmPlacementOptionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConfirmPlacementOptionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConfirmPlacementOptionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfirmPlacementOptionResponse(val *ConfirmPlacementOptionResponse) *NullableConfirmPlacementOptionResponse {
	return &NullableConfirmPlacementOptionResponse{value: val, isSet: true}
}

func (v NullableConfirmPlacementOptionResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableConfirmPlacementOptionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
