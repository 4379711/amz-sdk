package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the GenerateShipmentContentUpdatePreviewsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenerateShipmentContentUpdatePreviewsResponse{}

// GenerateShipmentContentUpdatePreviewsResponse The `GenerateShipmentContentUpdatePreviews` response.
type GenerateShipmentContentUpdatePreviewsResponse struct {
	// UUID for the given operation.
	OperationId string `json:"operationId" validate:"regexp=^[a-zA-Z0-9-]*$"`
}

type _GenerateShipmentContentUpdatePreviewsResponse GenerateShipmentContentUpdatePreviewsResponse

// NewGenerateShipmentContentUpdatePreviewsResponse instantiates a new GenerateShipmentContentUpdatePreviewsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateShipmentContentUpdatePreviewsResponse(operationId string) *GenerateShipmentContentUpdatePreviewsResponse {
	this := GenerateShipmentContentUpdatePreviewsResponse{}
	this.OperationId = operationId
	return &this
}

// NewGenerateShipmentContentUpdatePreviewsResponseWithDefaults instantiates a new GenerateShipmentContentUpdatePreviewsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateShipmentContentUpdatePreviewsResponseWithDefaults() *GenerateShipmentContentUpdatePreviewsResponse {
	this := GenerateShipmentContentUpdatePreviewsResponse{}
	return &this
}

// GetOperationId returns the OperationId field value
func (o *GenerateShipmentContentUpdatePreviewsResponse) GetOperationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperationId
}

// GetOperationIdOk returns a tuple with the OperationId field value
// and a boolean to check if the value has been set.
func (o *GenerateShipmentContentUpdatePreviewsResponse) GetOperationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationId, true
}

// SetOperationId sets field value
func (o *GenerateShipmentContentUpdatePreviewsResponse) SetOperationId(v string) {
	o.OperationId = v
}

func (o GenerateShipmentContentUpdatePreviewsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operationId"] = o.OperationId
	return toSerialize, nil
}

type NullableGenerateShipmentContentUpdatePreviewsResponse struct {
	value *GenerateShipmentContentUpdatePreviewsResponse
	isSet bool
}

func (v NullableGenerateShipmentContentUpdatePreviewsResponse) Get() *GenerateShipmentContentUpdatePreviewsResponse {
	return v.value
}

func (v *NullableGenerateShipmentContentUpdatePreviewsResponse) Set(val *GenerateShipmentContentUpdatePreviewsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateShipmentContentUpdatePreviewsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateShipmentContentUpdatePreviewsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateShipmentContentUpdatePreviewsResponse(val *GenerateShipmentContentUpdatePreviewsResponse) *NullableGenerateShipmentContentUpdatePreviewsResponse {
	return &NullableGenerateShipmentContentUpdatePreviewsResponse{value: val, isSet: true}
}

func (v NullableGenerateShipmentContentUpdatePreviewsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGenerateShipmentContentUpdatePreviewsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
