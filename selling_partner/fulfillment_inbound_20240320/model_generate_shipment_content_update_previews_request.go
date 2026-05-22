package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the GenerateShipmentContentUpdatePreviewsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenerateShipmentContentUpdatePreviewsRequest{}

// GenerateShipmentContentUpdatePreviewsRequest The `GenerateShipmentContentUpdatePreviews` request.
type GenerateShipmentContentUpdatePreviewsRequest struct {
	// A list of boxes that will be present in the shipment after the update.
	Boxes []BoxUpdateInput `json:"boxes"`
	// A list of all items that will be present in the shipment after the update.
	Items []ItemInput `json:"items"`
}

type _GenerateShipmentContentUpdatePreviewsRequest GenerateShipmentContentUpdatePreviewsRequest

// NewGenerateShipmentContentUpdatePreviewsRequest instantiates a new GenerateShipmentContentUpdatePreviewsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateShipmentContentUpdatePreviewsRequest(boxes []BoxUpdateInput, items []ItemInput) *GenerateShipmentContentUpdatePreviewsRequest {
	this := GenerateShipmentContentUpdatePreviewsRequest{}
	this.Boxes = boxes
	this.Items = items
	return &this
}

// NewGenerateShipmentContentUpdatePreviewsRequestWithDefaults instantiates a new GenerateShipmentContentUpdatePreviewsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateShipmentContentUpdatePreviewsRequestWithDefaults() *GenerateShipmentContentUpdatePreviewsRequest {
	this := GenerateShipmentContentUpdatePreviewsRequest{}
	return &this
}

// GetBoxes returns the Boxes field value
func (o *GenerateShipmentContentUpdatePreviewsRequest) GetBoxes() []BoxUpdateInput {
	if o == nil {
		var ret []BoxUpdateInput
		return ret
	}

	return o.Boxes
}

// GetBoxesOk returns a tuple with the Boxes field value
// and a boolean to check if the value has been set.
func (o *GenerateShipmentContentUpdatePreviewsRequest) GetBoxesOk() ([]BoxUpdateInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Boxes, true
}

// SetBoxes sets field value
func (o *GenerateShipmentContentUpdatePreviewsRequest) SetBoxes(v []BoxUpdateInput) {
	o.Boxes = v
}

// GetItems returns the Items field value
func (o *GenerateShipmentContentUpdatePreviewsRequest) GetItems() []ItemInput {
	if o == nil {
		var ret []ItemInput
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *GenerateShipmentContentUpdatePreviewsRequest) GetItemsOk() ([]ItemInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *GenerateShipmentContentUpdatePreviewsRequest) SetItems(v []ItemInput) {
	o.Items = v
}

func (o GenerateShipmentContentUpdatePreviewsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["boxes"] = o.Boxes
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableGenerateShipmentContentUpdatePreviewsRequest struct {
	value *GenerateShipmentContentUpdatePreviewsRequest
	isSet bool
}

func (v NullableGenerateShipmentContentUpdatePreviewsRequest) Get() *GenerateShipmentContentUpdatePreviewsRequest {
	return v.value
}

func (v *NullableGenerateShipmentContentUpdatePreviewsRequest) Set(val *GenerateShipmentContentUpdatePreviewsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateShipmentContentUpdatePreviewsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateShipmentContentUpdatePreviewsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateShipmentContentUpdatePreviewsRequest(val *GenerateShipmentContentUpdatePreviewsRequest) *NullableGenerateShipmentContentUpdatePreviewsRequest {
	return &NullableGenerateShipmentContentUpdatePreviewsRequest{value: val, isSet: true}
}

func (v NullableGenerateShipmentContentUpdatePreviewsRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGenerateShipmentContentUpdatePreviewsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
