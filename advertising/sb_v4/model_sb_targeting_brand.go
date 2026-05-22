package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBTargetingBrand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBTargetingBrand{}

// SBTargetingBrand struct for SBTargetingBrand
type SBTargetingBrand struct {
	// Id of brand. Use /sb/targets/categories/{categoryRefinementId}/refinements to retrieve Brand Refinement IDs.
	BrandRefinementId string `json:"brandRefinementId"`
	// Name of brand.
	Name *string `json:"name,omitempty"`
}

type _SBTargetingBrand SBTargetingBrand

// NewSBTargetingBrand instantiates a new SBTargetingBrand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBTargetingBrand(brandRefinementId string) *SBTargetingBrand {
	this := SBTargetingBrand{}
	this.BrandRefinementId = brandRefinementId
	return &this
}

// NewSBTargetingBrandWithDefaults instantiates a new SBTargetingBrand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBTargetingBrandWithDefaults() *SBTargetingBrand {
	this := SBTargetingBrand{}
	return &this
}

// GetBrandRefinementId returns the BrandRefinementId field value
func (o *SBTargetingBrand) GetBrandRefinementId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BrandRefinementId
}

// GetBrandRefinementIdOk returns a tuple with the BrandRefinementId field value
// and a boolean to check if the value has been set.
func (o *SBTargetingBrand) GetBrandRefinementIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrandRefinementId, true
}

// SetBrandRefinementId sets field value
func (o *SBTargetingBrand) SetBrandRefinementId(v string) {
	o.BrandRefinementId = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SBTargetingBrand) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBTargetingBrand) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SBTargetingBrand) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SBTargetingBrand) SetName(v string) {
	o.Name = &v
}

func (o SBTargetingBrand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["brandRefinementId"] = o.BrandRefinementId
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableSBTargetingBrand struct {
	value *SBTargetingBrand
	isSet bool
}

func (v NullableSBTargetingBrand) Get() *SBTargetingBrand {
	return v.value
}

func (v *NullableSBTargetingBrand) Set(val *SBTargetingBrand) {
	v.value = val
	v.isSet = true
}

func (v NullableSBTargetingBrand) IsSet() bool {
	return v.isSet
}

func (v *NullableSBTargetingBrand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBTargetingBrand(val *SBTargetingBrand) *NullableSBTargetingBrand {
	return &NullableSBTargetingBrand{value: val, isSet: true}
}

func (v NullableSBTargetingBrand) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBTargetingBrand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
