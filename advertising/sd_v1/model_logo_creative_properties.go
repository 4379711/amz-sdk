package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the LogoCreativeProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogoCreativeProperties{}

// LogoCreativeProperties User-customizable properties of a creative with a logo.
type LogoCreativeProperties struct {
	BrandLogo *Image `json:"brandLogo,omitempty"`
}

// NewLogoCreativeProperties instantiates a new LogoCreativeProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogoCreativeProperties() *LogoCreativeProperties {
	this := LogoCreativeProperties{}
	return &this
}

// NewLogoCreativePropertiesWithDefaults instantiates a new LogoCreativeProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogoCreativePropertiesWithDefaults() *LogoCreativeProperties {
	this := LogoCreativeProperties{}
	return &this
}

// GetBrandLogo returns the BrandLogo field value if set, zero value otherwise.
func (o *LogoCreativeProperties) GetBrandLogo() Image {
	if o == nil || IsNil(o.BrandLogo) {
		var ret Image
		return ret
	}
	return *o.BrandLogo
}

// GetBrandLogoOk returns a tuple with the BrandLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogoCreativeProperties) GetBrandLogoOk() (*Image, bool) {
	if o == nil || IsNil(o.BrandLogo) {
		return nil, false
	}
	return o.BrandLogo, true
}

// HasBrandLogo returns a boolean if a field has been set.
func (o *LogoCreativeProperties) HasBrandLogo() bool {
	if o != nil && !IsNil(o.BrandLogo) {
		return true
	}

	return false
}

// SetBrandLogo gets a reference to the given Image and assigns it to the BrandLogo field.
func (o *LogoCreativeProperties) SetBrandLogo(v Image) {
	o.BrandLogo = &v
}

func (o LogoCreativeProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BrandLogo) {
		toSerialize["brandLogo"] = o.BrandLogo
	}
	return toSerialize, nil
}

type NullableLogoCreativeProperties struct {
	value *LogoCreativeProperties
	isSet bool
}

func (v NullableLogoCreativeProperties) Get() *LogoCreativeProperties {
	return v.value
}

func (v *NullableLogoCreativeProperties) Set(val *LogoCreativeProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableLogoCreativeProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableLogoCreativeProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogoCreativeProperties(val *LogoCreativeProperties) *NullableLogoCreativeProperties {
	return &NullableLogoCreativeProperties{value: val, isSet: true}
}

func (v NullableLogoCreativeProperties) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableLogoCreativeProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
