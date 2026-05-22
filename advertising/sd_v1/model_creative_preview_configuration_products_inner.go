package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the CreativePreviewConfigurationProductsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreativePreviewConfigurationProductsInner{}

// CreativePreviewConfigurationProductsInner struct for CreativePreviewConfigurationProductsInner
type CreativePreviewConfigurationProductsInner struct {
	// The ASIN of the product.
	Asin *string `json:"asin,omitempty"`
}

// NewCreativePreviewConfigurationProductsInner instantiates a new CreativePreviewConfigurationProductsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreativePreviewConfigurationProductsInner() *CreativePreviewConfigurationProductsInner {
	this := CreativePreviewConfigurationProductsInner{}
	return &this
}

// NewCreativePreviewConfigurationProductsInnerWithDefaults instantiates a new CreativePreviewConfigurationProductsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreativePreviewConfigurationProductsInnerWithDefaults() *CreativePreviewConfigurationProductsInner {
	this := CreativePreviewConfigurationProductsInner{}
	return &this
}

// GetAsin returns the Asin field value if set, zero value otherwise.
func (o *CreativePreviewConfigurationProductsInner) GetAsin() string {
	if o == nil || IsNil(o.Asin) {
		var ret string
		return ret
	}
	return *o.Asin
}

// GetAsinOk returns a tuple with the Asin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativePreviewConfigurationProductsInner) GetAsinOk() (*string, bool) {
	if o == nil || IsNil(o.Asin) {
		return nil, false
	}
	return o.Asin, true
}

// HasAsin returns a boolean if a field has been set.
func (o *CreativePreviewConfigurationProductsInner) HasAsin() bool {
	if o != nil && !IsNil(o.Asin) {
		return true
	}

	return false
}

// SetAsin gets a reference to the given string and assigns it to the Asin field.
func (o *CreativePreviewConfigurationProductsInner) SetAsin(v string) {
	o.Asin = &v
}

func (o CreativePreviewConfigurationProductsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Asin) {
		toSerialize["asin"] = o.Asin
	}
	return toSerialize, nil
}

type NullableCreativePreviewConfigurationProductsInner struct {
	value *CreativePreviewConfigurationProductsInner
	isSet bool
}

func (v NullableCreativePreviewConfigurationProductsInner) Get() *CreativePreviewConfigurationProductsInner {
	return v.value
}

func (v *NullableCreativePreviewConfigurationProductsInner) Set(val *CreativePreviewConfigurationProductsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreativePreviewConfigurationProductsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreativePreviewConfigurationProductsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreativePreviewConfigurationProductsInner(val *CreativePreviewConfigurationProductsInner) *NullableCreativePreviewConfigurationProductsInner {
	return &NullableCreativePreviewConfigurationProductsInner{value: val, isSet: true}
}

func (v NullableCreativePreviewConfigurationProductsInner) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreativePreviewConfigurationProductsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
