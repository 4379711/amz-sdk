package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the BrandSafetyDenyListDomain type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BrandSafetyDenyListDomain{}

// BrandSafetyDenyListDomain struct for BrandSafetyDenyListDomain
type BrandSafetyDenyListDomain struct {
	// The website or app identifier. This can be in the form of full domain (eg. 'example.com' or 'example.net'), or mobile app identifier (eg. 'com.example.app' for Android apps or '1234567890' for iOS apps)
	Name string                        `json:"name"`
	Type BrandSafetyDenyListDomainType `json:"type"`
}

type _BrandSafetyDenyListDomain BrandSafetyDenyListDomain

// NewBrandSafetyDenyListDomain instantiates a new BrandSafetyDenyListDomain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrandSafetyDenyListDomain(name string, type_ BrandSafetyDenyListDomainType) *BrandSafetyDenyListDomain {
	this := BrandSafetyDenyListDomain{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewBrandSafetyDenyListDomainWithDefaults instantiates a new BrandSafetyDenyListDomain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandSafetyDenyListDomainWithDefaults() *BrandSafetyDenyListDomain {
	this := BrandSafetyDenyListDomain{}
	return &this
}

// GetName returns the Name field value
func (o *BrandSafetyDenyListDomain) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BrandSafetyDenyListDomain) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BrandSafetyDenyListDomain) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *BrandSafetyDenyListDomain) GetType() BrandSafetyDenyListDomainType {
	if o == nil {
		var ret BrandSafetyDenyListDomainType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BrandSafetyDenyListDomain) GetTypeOk() (*BrandSafetyDenyListDomainType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BrandSafetyDenyListDomain) SetType(v BrandSafetyDenyListDomainType) {
	o.Type = v
}

func (o BrandSafetyDenyListDomain) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableBrandSafetyDenyListDomain struct {
	value *BrandSafetyDenyListDomain
	isSet bool
}

func (v NullableBrandSafetyDenyListDomain) Get() *BrandSafetyDenyListDomain {
	return v.value
}

func (v *NullableBrandSafetyDenyListDomain) Set(val *BrandSafetyDenyListDomain) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandSafetyDenyListDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandSafetyDenyListDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandSafetyDenyListDomain(val *BrandSafetyDenyListDomain) *NullableBrandSafetyDenyListDomain {
	return &NullableBrandSafetyDenyListDomain{value: val, isSet: true}
}

func (v NullableBrandSafetyDenyListDomain) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBrandSafetyDenyListDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
