package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the BoxInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BoxInput{}

// BoxInput Input information for a given box.
type BoxInput struct {
	ContentInformationSource BoxContentInformationSource `json:"contentInformationSource"`
	Dimensions               Dimensions                  `json:"dimensions"`
	// The items and their quantity in the box. This must be empty if the box `contentInformationSource` is `BARCODE_2D` or `MANUAL_PROCESS`.
	Items []ItemInput `json:"items,omitempty"`
	// The number of containers where all other properties like weight or dimensions are identical.
	Quantity int32  `json:"quantity"`
	Weight   Weight `json:"weight"`
}

type _BoxInput BoxInput

// NewBoxInput instantiates a new BoxInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBoxInput(contentInformationSource BoxContentInformationSource, dimensions Dimensions, quantity int32, weight Weight) *BoxInput {
	this := BoxInput{}
	this.ContentInformationSource = contentInformationSource
	this.Dimensions = dimensions
	this.Quantity = quantity
	this.Weight = weight
	return &this
}

// NewBoxInputWithDefaults instantiates a new BoxInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBoxInputWithDefaults() *BoxInput {
	this := BoxInput{}
	return &this
}

// GetContentInformationSource returns the ContentInformationSource field value
func (o *BoxInput) GetContentInformationSource() BoxContentInformationSource {
	if o == nil {
		var ret BoxContentInformationSource
		return ret
	}

	return o.ContentInformationSource
}

// GetContentInformationSourceOk returns a tuple with the ContentInformationSource field value
// and a boolean to check if the value has been set.
func (o *BoxInput) GetContentInformationSourceOk() (*BoxContentInformationSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentInformationSource, true
}

// SetContentInformationSource sets field value
func (o *BoxInput) SetContentInformationSource(v BoxContentInformationSource) {
	o.ContentInformationSource = v
}

// GetDimensions returns the Dimensions field value
func (o *BoxInput) GetDimensions() Dimensions {
	if o == nil {
		var ret Dimensions
		return ret
	}

	return o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value
// and a boolean to check if the value has been set.
func (o *BoxInput) GetDimensionsOk() (*Dimensions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dimensions, true
}

// SetDimensions sets field value
func (o *BoxInput) SetDimensions(v Dimensions) {
	o.Dimensions = v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *BoxInput) GetItems() []ItemInput {
	if o == nil || IsNil(o.Items) {
		var ret []ItemInput
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoxInput) GetItemsOk() ([]ItemInput, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *BoxInput) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ItemInput and assigns it to the Items field.
func (o *BoxInput) SetItems(v []ItemInput) {
	o.Items = v
}

// GetQuantity returns the Quantity field value
func (o *BoxInput) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *BoxInput) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *BoxInput) SetQuantity(v int32) {
	o.Quantity = v
}

// GetWeight returns the Weight field value
func (o *BoxInput) GetWeight() Weight {
	if o == nil {
		var ret Weight
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *BoxInput) GetWeightOk() (*Weight, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *BoxInput) SetWeight(v Weight) {
	o.Weight = v
}

func (o BoxInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["contentInformationSource"] = o.ContentInformationSource
	toSerialize["dimensions"] = o.Dimensions
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	toSerialize["quantity"] = o.Quantity
	toSerialize["weight"] = o.Weight
	return toSerialize, nil
}

type NullableBoxInput struct {
	value *BoxInput
	isSet bool
}

func (v NullableBoxInput) Get() *BoxInput {
	return v.value
}

func (v *NullableBoxInput) Set(val *BoxInput) {
	v.value = val
	v.isSet = true
}

func (v NullableBoxInput) IsSet() bool {
	return v.isSet
}

func (v *NullableBoxInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBoxInput(val *BoxInput) *NullableBoxInput {
	return &NullableBoxInput{value: val, isSet: true}
}

func (v NullableBoxInput) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBoxInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
