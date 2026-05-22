package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the ConfirmShipmentContentUpdatePreviewResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfirmShipmentContentUpdatePreviewResponse{}

// ConfirmShipmentContentUpdatePreviewResponse The `confirmShipmentContentUpdatePreview` response.
type ConfirmShipmentContentUpdatePreviewResponse struct {
	// UUID for the given operation.
	OperationId string `json:"operationId" validate:"regexp=^[a-zA-Z0-9-]*$"`
}

type _ConfirmShipmentContentUpdatePreviewResponse ConfirmShipmentContentUpdatePreviewResponse

// NewConfirmShipmentContentUpdatePreviewResponse instantiates a new ConfirmShipmentContentUpdatePreviewResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfirmShipmentContentUpdatePreviewResponse(operationId string) *ConfirmShipmentContentUpdatePreviewResponse {
	this := ConfirmShipmentContentUpdatePreviewResponse{}
	this.OperationId = operationId
	return &this
}

// NewConfirmShipmentContentUpdatePreviewResponseWithDefaults instantiates a new ConfirmShipmentContentUpdatePreviewResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfirmShipmentContentUpdatePreviewResponseWithDefaults() *ConfirmShipmentContentUpdatePreviewResponse {
	this := ConfirmShipmentContentUpdatePreviewResponse{}
	return &this
}

// GetOperationId returns the OperationId field value
func (o *ConfirmShipmentContentUpdatePreviewResponse) GetOperationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperationId
}

// GetOperationIdOk returns a tuple with the OperationId field value
// and a boolean to check if the value has been set.
func (o *ConfirmShipmentContentUpdatePreviewResponse) GetOperationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationId, true
}

// SetOperationId sets field value
func (o *ConfirmShipmentContentUpdatePreviewResponse) SetOperationId(v string) {
	o.OperationId = v
}

func (o ConfirmShipmentContentUpdatePreviewResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operationId"] = o.OperationId
	return toSerialize, nil
}

type NullableConfirmShipmentContentUpdatePreviewResponse struct {
	value *ConfirmShipmentContentUpdatePreviewResponse
	isSet bool
}

func (v NullableConfirmShipmentContentUpdatePreviewResponse) Get() *ConfirmShipmentContentUpdatePreviewResponse {
	return v.value
}

func (v *NullableConfirmShipmentContentUpdatePreviewResponse) Set(val *ConfirmShipmentContentUpdatePreviewResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConfirmShipmentContentUpdatePreviewResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConfirmShipmentContentUpdatePreviewResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfirmShipmentContentUpdatePreviewResponse(val *ConfirmShipmentContentUpdatePreviewResponse) *NullableConfirmShipmentContentUpdatePreviewResponse {
	return &NullableConfirmShipmentContentUpdatePreviewResponse{value: val, isSet: true}
}

func (v NullableConfirmShipmentContentUpdatePreviewResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableConfirmShipmentContentUpdatePreviewResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
