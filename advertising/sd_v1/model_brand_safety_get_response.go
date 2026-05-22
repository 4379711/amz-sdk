package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the BrandSafetyGetResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BrandSafetyGetResponse{}

// BrandSafetyGetResponse Response for Brand Safety Deny List GET requests
type BrandSafetyGetResponse struct {
	Pagination *BrandSafetyGetResponsePagination `json:"pagination,omitempty"`
	// List of Brand Safety Deny List Domains
	Domains []BrandSafetyDenyListProcessedDomain `json:"domains,omitempty"`
}

// NewBrandSafetyGetResponse instantiates a new BrandSafetyGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrandSafetyGetResponse() *BrandSafetyGetResponse {
	this := BrandSafetyGetResponse{}
	return &this
}

// NewBrandSafetyGetResponseWithDefaults instantiates a new BrandSafetyGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandSafetyGetResponseWithDefaults() *BrandSafetyGetResponse {
	this := BrandSafetyGetResponse{}
	return &this
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *BrandSafetyGetResponse) GetPagination() BrandSafetyGetResponsePagination {
	if o == nil || IsNil(o.Pagination) {
		var ret BrandSafetyGetResponsePagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandSafetyGetResponse) GetPaginationOk() (*BrandSafetyGetResponsePagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *BrandSafetyGetResponse) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given BrandSafetyGetResponsePagination and assigns it to the Pagination field.
func (o *BrandSafetyGetResponse) SetPagination(v BrandSafetyGetResponsePagination) {
	o.Pagination = &v
}

// GetDomains returns the Domains field value if set, zero value otherwise.
func (o *BrandSafetyGetResponse) GetDomains() []BrandSafetyDenyListProcessedDomain {
	if o == nil || IsNil(o.Domains) {
		var ret []BrandSafetyDenyListProcessedDomain
		return ret
	}
	return o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandSafetyGetResponse) GetDomainsOk() ([]BrandSafetyDenyListProcessedDomain, bool) {
	if o == nil || IsNil(o.Domains) {
		return nil, false
	}
	return o.Domains, true
}

// HasDomains returns a boolean if a field has been set.
func (o *BrandSafetyGetResponse) HasDomains() bool {
	if o != nil && !IsNil(o.Domains) {
		return true
	}

	return false
}

// SetDomains gets a reference to the given []BrandSafetyDenyListProcessedDomain and assigns it to the Domains field.
func (o *BrandSafetyGetResponse) SetDomains(v []BrandSafetyDenyListProcessedDomain) {
	o.Domains = v
}

func (o BrandSafetyGetResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	if !IsNil(o.Domains) {
		toSerialize["domains"] = o.Domains
	}
	return toSerialize, nil
}

type NullableBrandSafetyGetResponse struct {
	value *BrandSafetyGetResponse
	isSet bool
}

func (v NullableBrandSafetyGetResponse) Get() *BrandSafetyGetResponse {
	return v.value
}

func (v *NullableBrandSafetyGetResponse) Set(val *BrandSafetyGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandSafetyGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandSafetyGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandSafetyGetResponse(val *BrandSafetyGetResponse) *NullableBrandSafetyGetResponse {
	return &NullableBrandSafetyGetResponse{value: val, isSet: true}
}

func (v NullableBrandSafetyGetResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBrandSafetyGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
