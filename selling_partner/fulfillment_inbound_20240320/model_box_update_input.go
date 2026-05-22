package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the BoxUpdateInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BoxUpdateInput{}

// BoxUpdateInput Input information for updating a box
type BoxUpdateInput struct {
	ContentInformationSource BoxContentInformationSource `json:"contentInformationSource"`
	Dimensions               Dimensions                  `json:"dimensions"`
	// The items and their quantity in the box. This must be empty if the box `contentInformationSource` is `BARCODE_2D` or `MANUAL_PROCESS`.
	Items []ItemInput `json:"items,omitempty"`
	// Primary key to uniquely identify a Box Package. PackageId must be provided if the intent is to update an existing box. Adding a new box will not require providing this value. Any existing PackageIds not provided will be treated as to-be-removed
	PackageId *string `json:"packageId,omitempty" validate:"regexp=^[a-zA-Z0-9-]*$"`
	// The number of containers where all other properties like weight or dimensions are identical.
	Quantity int32  `json:"quantity"`
	Weight   Weight `json:"weight"`
}

type _BoxUpdateInput BoxUpdateInput

// NewBoxUpdateInput instantiates a new BoxUpdateInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBoxUpdateInput(contentInformationSource BoxContentInformationSource, dimensions Dimensions, quantity int32, weight Weight) *BoxUpdateInput {
	this := BoxUpdateInput{}
	this.ContentInformationSource = contentInformationSource
	this.Dimensions = dimensions
	this.Quantity = quantity
	this.Weight = weight
	return &this
}

// NewBoxUpdateInputWithDefaults instantiates a new BoxUpdateInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBoxUpdateInputWithDefaults() *BoxUpdateInput {
	this := BoxUpdateInput{}
	return &this
}

// GetContentInformationSource returns the ContentInformationSource field value
func (o *BoxUpdateInput) GetContentInformationSource() BoxContentInformationSource {
	if o == nil {
		var ret BoxContentInformationSource
		return ret
	}

	return o.ContentInformationSource
}

// GetContentInformationSourceOk returns a tuple with the ContentInformationSource field value
// and a boolean to check if the value has been set.
func (o *BoxUpdateInput) GetContentInformationSourceOk() (*BoxContentInformationSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentInformationSource, true
}

// SetContentInformationSource sets field value
func (o *BoxUpdateInput) SetContentInformationSource(v BoxContentInformationSource) {
	o.ContentInformationSource = v
}

// GetDimensions returns the Dimensions field value
func (o *BoxUpdateInput) GetDimensions() Dimensions {
	if o == nil {
		var ret Dimensions
		return ret
	}

	return o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value
// and a boolean to check if the value has been set.
func (o *BoxUpdateInput) GetDimensionsOk() (*Dimensions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dimensions, true
}

// SetDimensions sets field value
func (o *BoxUpdateInput) SetDimensions(v Dimensions) {
	o.Dimensions = v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *BoxUpdateInput) GetItems() []ItemInput {
	if o == nil || IsNil(o.Items) {
		var ret []ItemInput
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoxUpdateInput) GetItemsOk() ([]ItemInput, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *BoxUpdateInput) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ItemInput and assigns it to the Items field.
func (o *BoxUpdateInput) SetItems(v []ItemInput) {
	o.Items = v
}

// GetPackageId returns the PackageId field value if set, zero value otherwise.
func (o *BoxUpdateInput) GetPackageId() string {
	if o == nil || IsNil(o.PackageId) {
		var ret string
		return ret
	}
	return *o.PackageId
}

// GetPackageIdOk returns a tuple with the PackageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoxUpdateInput) GetPackageIdOk() (*string, bool) {
	if o == nil || IsNil(o.PackageId) {
		return nil, false
	}
	return o.PackageId, true
}

// HasPackageId returns a boolean if a field has been set.
func (o *BoxUpdateInput) HasPackageId() bool {
	if o != nil && !IsNil(o.PackageId) {
		return true
	}

	return false
}

// SetPackageId gets a reference to the given string and assigns it to the PackageId field.
func (o *BoxUpdateInput) SetPackageId(v string) {
	o.PackageId = &v
}

// GetQuantity returns the Quantity field value
func (o *BoxUpdateInput) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *BoxUpdateInput) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *BoxUpdateInput) SetQuantity(v int32) {
	o.Quantity = v
}

// GetWeight returns the Weight field value
func (o *BoxUpdateInput) GetWeight() Weight {
	if o == nil {
		var ret Weight
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *BoxUpdateInput) GetWeightOk() (*Weight, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *BoxUpdateInput) SetWeight(v Weight) {
	o.Weight = v
}

func (o BoxUpdateInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["contentInformationSource"] = o.ContentInformationSource
	toSerialize["dimensions"] = o.Dimensions
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.PackageId) {
		toSerialize["packageId"] = o.PackageId
	}
	toSerialize["quantity"] = o.Quantity
	toSerialize["weight"] = o.Weight
	return toSerialize, nil
}

type NullableBoxUpdateInput struct {
	value *BoxUpdateInput
	isSet bool
}

func (v NullableBoxUpdateInput) Get() *BoxUpdateInput {
	return v.value
}

func (v *NullableBoxUpdateInput) Set(val *BoxUpdateInput) {
	v.value = val
	v.isSet = true
}

func (v NullableBoxUpdateInput) IsSet() bool {
	return v.isSet
}

func (v *NullableBoxUpdateInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBoxUpdateInput(val *BoxUpdateInput) *NullableBoxUpdateInput {
	return &NullableBoxUpdateInput{value: val, isSet: true}
}

func (v NullableBoxUpdateInput) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBoxUpdateInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
