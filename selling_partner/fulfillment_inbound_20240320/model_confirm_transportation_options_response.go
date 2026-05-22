package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the ConfirmTransportationOptionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfirmTransportationOptionsResponse{}

// ConfirmTransportationOptionsResponse The `confirmTransportationOptions` response.
type ConfirmTransportationOptionsResponse struct {
	// UUID for the given operation.
	OperationId string `json:"operationId" validate:"regexp=^[a-zA-Z0-9-]*$"`
}

type _ConfirmTransportationOptionsResponse ConfirmTransportationOptionsResponse

// NewConfirmTransportationOptionsResponse instantiates a new ConfirmTransportationOptionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfirmTransportationOptionsResponse(operationId string) *ConfirmTransportationOptionsResponse {
	this := ConfirmTransportationOptionsResponse{}
	this.OperationId = operationId
	return &this
}

// NewConfirmTransportationOptionsResponseWithDefaults instantiates a new ConfirmTransportationOptionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfirmTransportationOptionsResponseWithDefaults() *ConfirmTransportationOptionsResponse {
	this := ConfirmTransportationOptionsResponse{}
	return &this
}

// GetOperationId returns the OperationId field value
func (o *ConfirmTransportationOptionsResponse) GetOperationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperationId
}

// GetOperationIdOk returns a tuple with the OperationId field value
// and a boolean to check if the value has been set.
func (o *ConfirmTransportationOptionsResponse) GetOperationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationId, true
}

// SetOperationId sets field value
func (o *ConfirmTransportationOptionsResponse) SetOperationId(v string) {
	o.OperationId = v
}

func (o ConfirmTransportationOptionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operationId"] = o.OperationId
	return toSerialize, nil
}

type NullableConfirmTransportationOptionsResponse struct {
	value *ConfirmTransportationOptionsResponse
	isSet bool
}

func (v NullableConfirmTransportationOptionsResponse) Get() *ConfirmTransportationOptionsResponse {
	return v.value
}

func (v *NullableConfirmTransportationOptionsResponse) Set(val *ConfirmTransportationOptionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConfirmTransportationOptionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConfirmTransportationOptionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfirmTransportationOptionsResponse(val *ConfirmTransportationOptionsResponse) *NullableConfirmTransportationOptionsResponse {
	return &NullableConfirmTransportationOptionsResponse{value: val, isSet: true}
}

func (v NullableConfirmTransportationOptionsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableConfirmTransportationOptionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
