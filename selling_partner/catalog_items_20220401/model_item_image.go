package catalog_items_20220401

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemImage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemImage{}

// ItemImage Image for an item in the Amazon catalog.
type ItemImage struct {
	// Variant of the image, such as `MAIN` or `PT01`.
	Variant string `json:"variant"`
	// URL for the image.
	Link string `json:"link"`
	// Height of the image in pixels.
	Height int32 `json:"height"`
	// Width of the image in pixels.
	Width int32 `json:"width"`
}

type _ItemImage ItemImage

// NewItemImage instantiates a new ItemImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemImage(variant string, link string, height int32, width int32) *ItemImage {
	this := ItemImage{}
	this.Variant = variant
	this.Link = link
	this.Height = height
	this.Width = width
	return &this
}

// NewItemImageWithDefaults instantiates a new ItemImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemImageWithDefaults() *ItemImage {
	this := ItemImage{}
	return &this
}

// GetVariant returns the Variant field value
func (o *ItemImage) GetVariant() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Variant
}

// GetVariantOk returns a tuple with the Variant field value
// and a boolean to check if the value has been set.
func (o *ItemImage) GetVariantOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Variant, true
}

// SetVariant sets field value
func (o *ItemImage) SetVariant(v string) {
	o.Variant = v
}

// GetLink returns the Link field value
func (o *ItemImage) GetLink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Link
}

// GetLinkOk returns a tuple with the Link field value
// and a boolean to check if the value has been set.
func (o *ItemImage) GetLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Link, true
}

// SetLink sets field value
func (o *ItemImage) SetLink(v string) {
	o.Link = v
}

// GetHeight returns the Height field value
func (o *ItemImage) GetHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *ItemImage) GetHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *ItemImage) SetHeight(v int32) {
	o.Height = v
}

// GetWidth returns the Width field value
func (o *ItemImage) GetWidth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Width
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
func (o *ItemImage) GetWidthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Width, true
}

// SetWidth sets field value
func (o *ItemImage) SetWidth(v int32) {
	o.Width = v
}

func (o ItemImage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["variant"] = o.Variant
	toSerialize["link"] = o.Link
	toSerialize["height"] = o.Height
	toSerialize["width"] = o.Width
	return toSerialize, nil
}

type NullableItemImage struct {
	value *ItemImage
	isSet bool
}

func (v NullableItemImage) Get() *ItemImage {
	return v.value
}

func (v *NullableItemImage) Set(val *ItemImage) {
	v.value = val
	v.isSet = true
}

func (v NullableItemImage) IsSet() bool {
	return v.isSet
}

func (v *NullableItemImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemImage(val *ItemImage) *NullableItemImage {
	return &NullableItemImage{value: val, isSet: true}
}

func (v NullableItemImage) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
