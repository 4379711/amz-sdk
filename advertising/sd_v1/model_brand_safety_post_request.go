package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the BrandSafetyPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BrandSafetyPostRequest{}

// BrandSafetyPostRequest POST Request for Brand Safety
type BrandSafetyPostRequest struct {
	Domains []BrandSafetyDenyListDomain `json:"domains"`
}

type _BrandSafetyPostRequest BrandSafetyPostRequest

// NewBrandSafetyPostRequest instantiates a new BrandSafetyPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrandSafetyPostRequest(domains []BrandSafetyDenyListDomain) *BrandSafetyPostRequest {
	this := BrandSafetyPostRequest{}
	this.Domains = domains
	return &this
}

// NewBrandSafetyPostRequestWithDefaults instantiates a new BrandSafetyPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandSafetyPostRequestWithDefaults() *BrandSafetyPostRequest {
	this := BrandSafetyPostRequest{}
	return &this
}

// GetDomains returns the Domains field value
func (o *BrandSafetyPostRequest) GetDomains() []BrandSafetyDenyListDomain {
	if o == nil {
		var ret []BrandSafetyDenyListDomain
		return ret
	}

	return o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value
// and a boolean to check if the value has been set.
func (o *BrandSafetyPostRequest) GetDomainsOk() ([]BrandSafetyDenyListDomain, bool) {
	if o == nil {
		return nil, false
	}
	return o.Domains, true
}

// SetDomains sets field value
func (o *BrandSafetyPostRequest) SetDomains(v []BrandSafetyDenyListDomain) {
	o.Domains = v
}

func (o BrandSafetyPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["domains"] = o.Domains
	return toSerialize, nil
}

type NullableBrandSafetyPostRequest struct {
	value *BrandSafetyPostRequest
	isSet bool
}

func (v NullableBrandSafetyPostRequest) Get() *BrandSafetyPostRequest {
	return v.value
}

func (v *NullableBrandSafetyPostRequest) Set(val *BrandSafetyPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandSafetyPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandSafetyPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandSafetyPostRequest(val *BrandSafetyPostRequest) *NullableBrandSafetyPostRequest {
	return &NullableBrandSafetyPostRequest{value: val, isSet: true}
}

func (v NullableBrandSafetyPostRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBrandSafetyPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
