package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the SetPackingInformationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetPackingInformationResponse{}

// SetPackingInformationResponse The `setPackingInformation` response.
type SetPackingInformationResponse struct {
	// UUID for the given operation.
	OperationId string `json:"operationId" validate:"regexp=^[a-zA-Z0-9-]*$"`
}

type _SetPackingInformationResponse SetPackingInformationResponse

// NewSetPackingInformationResponse instantiates a new SetPackingInformationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetPackingInformationResponse(operationId string) *SetPackingInformationResponse {
	this := SetPackingInformationResponse{}
	this.OperationId = operationId
	return &this
}

// NewSetPackingInformationResponseWithDefaults instantiates a new SetPackingInformationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetPackingInformationResponseWithDefaults() *SetPackingInformationResponse {
	this := SetPackingInformationResponse{}
	return &this
}

// GetOperationId returns the OperationId field value
func (o *SetPackingInformationResponse) GetOperationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperationId
}

// GetOperationIdOk returns a tuple with the OperationId field value
// and a boolean to check if the value has been set.
func (o *SetPackingInformationResponse) GetOperationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationId, true
}

// SetOperationId sets field value
func (o *SetPackingInformationResponse) SetOperationId(v string) {
	o.OperationId = v
}

func (o SetPackingInformationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operationId"] = o.OperationId
	return toSerialize, nil
}

type NullableSetPackingInformationResponse struct {
	value *SetPackingInformationResponse
	isSet bool
}

func (v NullableSetPackingInformationResponse) Get() *SetPackingInformationResponse {
	return v.value
}

func (v *NullableSetPackingInformationResponse) Set(val *SetPackingInformationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSetPackingInformationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSetPackingInformationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetPackingInformationResponse(val *SetPackingInformationResponse) *NullableSetPackingInformationResponse {
	return &NullableSetPackingInformationResponse{value: val, isSet: true}
}

func (v NullableSetPackingInformationResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSetPackingInformationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
