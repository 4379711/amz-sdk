package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the GeneratePackingOptionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GeneratePackingOptionsResponse{}

// GeneratePackingOptionsResponse The `generatePackingOptions` response.
type GeneratePackingOptionsResponse struct {
	// UUID for the given operation.
	OperationId string `json:"operationId" validate:"regexp=^[a-zA-Z0-9-]*$"`
}

type _GeneratePackingOptionsResponse GeneratePackingOptionsResponse

// NewGeneratePackingOptionsResponse instantiates a new GeneratePackingOptionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeneratePackingOptionsResponse(operationId string) *GeneratePackingOptionsResponse {
	this := GeneratePackingOptionsResponse{}
	this.OperationId = operationId
	return &this
}

// NewGeneratePackingOptionsResponseWithDefaults instantiates a new GeneratePackingOptionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeneratePackingOptionsResponseWithDefaults() *GeneratePackingOptionsResponse {
	this := GeneratePackingOptionsResponse{}
	return &this
}

// GetOperationId returns the OperationId field value
func (o *GeneratePackingOptionsResponse) GetOperationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperationId
}

// GetOperationIdOk returns a tuple with the OperationId field value
// and a boolean to check if the value has been set.
func (o *GeneratePackingOptionsResponse) GetOperationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationId, true
}

// SetOperationId sets field value
func (o *GeneratePackingOptionsResponse) SetOperationId(v string) {
	o.OperationId = v
}

func (o GeneratePackingOptionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operationId"] = o.OperationId
	return toSerialize, nil
}

type NullableGeneratePackingOptionsResponse struct {
	value *GeneratePackingOptionsResponse
	isSet bool
}

func (v NullableGeneratePackingOptionsResponse) Get() *GeneratePackingOptionsResponse {
	return v.value
}

func (v *NullableGeneratePackingOptionsResponse) Set(val *GeneratePackingOptionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGeneratePackingOptionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGeneratePackingOptionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeneratePackingOptionsResponse(val *GeneratePackingOptionsResponse) *NullableGeneratePackingOptionsResponse {
	return &NullableGeneratePackingOptionsResponse{value: val, isSet: true}
}

func (v NullableGeneratePackingOptionsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGeneratePackingOptionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
