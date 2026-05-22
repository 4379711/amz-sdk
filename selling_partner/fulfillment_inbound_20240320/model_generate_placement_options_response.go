package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the GeneratePlacementOptionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GeneratePlacementOptionsResponse{}

// GeneratePlacementOptionsResponse The `generatePlacementOptions` response.
type GeneratePlacementOptionsResponse struct {
	// UUID for the given operation.
	OperationId string `json:"operationId" validate:"regexp=^[a-zA-Z0-9-]*$"`
}

type _GeneratePlacementOptionsResponse GeneratePlacementOptionsResponse

// NewGeneratePlacementOptionsResponse instantiates a new GeneratePlacementOptionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeneratePlacementOptionsResponse(operationId string) *GeneratePlacementOptionsResponse {
	this := GeneratePlacementOptionsResponse{}
	this.OperationId = operationId
	return &this
}

// NewGeneratePlacementOptionsResponseWithDefaults instantiates a new GeneratePlacementOptionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeneratePlacementOptionsResponseWithDefaults() *GeneratePlacementOptionsResponse {
	this := GeneratePlacementOptionsResponse{}
	return &this
}

// GetOperationId returns the OperationId field value
func (o *GeneratePlacementOptionsResponse) GetOperationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperationId
}

// GetOperationIdOk returns a tuple with the OperationId field value
// and a boolean to check if the value has been set.
func (o *GeneratePlacementOptionsResponse) GetOperationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationId, true
}

// SetOperationId sets field value
func (o *GeneratePlacementOptionsResponse) SetOperationId(v string) {
	o.OperationId = v
}

func (o GeneratePlacementOptionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operationId"] = o.OperationId
	return toSerialize, nil
}

type NullableGeneratePlacementOptionsResponse struct {
	value *GeneratePlacementOptionsResponse
	isSet bool
}

func (v NullableGeneratePlacementOptionsResponse) Get() *GeneratePlacementOptionsResponse {
	return v.value
}

func (v *NullableGeneratePlacementOptionsResponse) Set(val *GeneratePlacementOptionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGeneratePlacementOptionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGeneratePlacementOptionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeneratePlacementOptionsResponse(val *GeneratePlacementOptionsResponse) *NullableGeneratePlacementOptionsResponse {
	return &NullableGeneratePlacementOptionsResponse{value: val, isSet: true}
}

func (v NullableGeneratePlacementOptionsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGeneratePlacementOptionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
