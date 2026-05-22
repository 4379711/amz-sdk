package aplus_content_20201101

import (
	"github.com/bytedance/sonic"
)

// checks if the StandardCompanyLogoModule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StandardCompanyLogoModule{}

// StandardCompanyLogoModule The standard company logo image.
type StandardCompanyLogoModule struct {
	CompanyLogo ImageComponent `json:"companyLogo"`
}

type _StandardCompanyLogoModule StandardCompanyLogoModule

// NewStandardCompanyLogoModule instantiates a new StandardCompanyLogoModule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandardCompanyLogoModule(companyLogo ImageComponent) *StandardCompanyLogoModule {
	this := StandardCompanyLogoModule{}
	this.CompanyLogo = companyLogo
	return &this
}

// NewStandardCompanyLogoModuleWithDefaults instantiates a new StandardCompanyLogoModule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandardCompanyLogoModuleWithDefaults() *StandardCompanyLogoModule {
	this := StandardCompanyLogoModule{}
	return &this
}

// GetCompanyLogo returns the CompanyLogo field value
func (o *StandardCompanyLogoModule) GetCompanyLogo() ImageComponent {
	if o == nil {
		var ret ImageComponent
		return ret
	}

	return o.CompanyLogo
}

// GetCompanyLogoOk returns a tuple with the CompanyLogo field value
// and a boolean to check if the value has been set.
func (o *StandardCompanyLogoModule) GetCompanyLogoOk() (*ImageComponent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompanyLogo, true
}

// SetCompanyLogo sets field value
func (o *StandardCompanyLogoModule) SetCompanyLogo(v ImageComponent) {
	o.CompanyLogo = v
}

func (o StandardCompanyLogoModule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["companyLogo"] = o.CompanyLogo
	return toSerialize, nil
}

type NullableStandardCompanyLogoModule struct {
	value *StandardCompanyLogoModule
	isSet bool
}

func (v NullableStandardCompanyLogoModule) Get() *StandardCompanyLogoModule {
	return v.value
}

func (v *NullableStandardCompanyLogoModule) Set(val *StandardCompanyLogoModule) {
	v.value = val
	v.isSet = true
}

func (v NullableStandardCompanyLogoModule) IsSet() bool {
	return v.isSet
}

func (v *NullableStandardCompanyLogoModule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandardCompanyLogoModule(val *StandardCompanyLogoModule) *NullableStandardCompanyLogoModule {
	return &NullableStandardCompanyLogoModule{value: val, isSet: true}
}

func (v NullableStandardCompanyLogoModule) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableStandardCompanyLogoModule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
