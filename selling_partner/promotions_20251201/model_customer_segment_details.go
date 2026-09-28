package promotions_20251201

import (
	"github.com/bytedance/sonic"
)

// checks if the CustomerSegmentDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerSegmentDetails{}

// CustomerSegmentDetails Additional segment type-specific details. Use the appropriate property based on `segmentType`.
type CustomerSegmentDetails struct {
	BrandSegmentDetails *BrandSegmentDetails `json:"brandSegmentDetails,omitempty"`
}

// NewCustomerSegmentDetails instantiates a new CustomerSegmentDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerSegmentDetails() *CustomerSegmentDetails {
	this := CustomerSegmentDetails{}
	return &this
}

// NewCustomerSegmentDetailsWithDefaults instantiates a new CustomerSegmentDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerSegmentDetailsWithDefaults() *CustomerSegmentDetails {
	this := CustomerSegmentDetails{}
	return &this
}

// GetBrandSegmentDetails returns the BrandSegmentDetails field value if set, zero value otherwise.
func (o *CustomerSegmentDetails) GetBrandSegmentDetails() BrandSegmentDetails {
	if o == nil || IsNil(o.BrandSegmentDetails) {
		var ret BrandSegmentDetails
		return ret
	}
	return *o.BrandSegmentDetails
}

// GetBrandSegmentDetailsOk returns a tuple with the BrandSegmentDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerSegmentDetails) GetBrandSegmentDetailsOk() (*BrandSegmentDetails, bool) {
	if o == nil || IsNil(o.BrandSegmentDetails) {
		return nil, false
	}
	return o.BrandSegmentDetails, true
}

// HasBrandSegmentDetails returns a boolean if a field has been set.
func (o *CustomerSegmentDetails) HasBrandSegmentDetails() bool {
	if o != nil && !IsNil(o.BrandSegmentDetails) {
		return true
	}

	return false
}

// SetBrandSegmentDetails gets a reference to the given BrandSegmentDetails and assigns it to the BrandSegmentDetails field.
func (o *CustomerSegmentDetails) SetBrandSegmentDetails(v BrandSegmentDetails) {
	o.BrandSegmentDetails = &v
}

func (o CustomerSegmentDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BrandSegmentDetails) {
		toSerialize["brandSegmentDetails"] = o.BrandSegmentDetails
	}
	return toSerialize, nil
}

type NullableCustomerSegmentDetails struct {
	value *CustomerSegmentDetails
	isSet bool
}

func (v NullableCustomerSegmentDetails) Get() *CustomerSegmentDetails {
	return v.value
}

func (v *NullableCustomerSegmentDetails) Set(val *CustomerSegmentDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerSegmentDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerSegmentDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerSegmentDetails(val *CustomerSegmentDetails) *NullableCustomerSegmentDetails {
	return &NullableCustomerSegmentDetails{value: val, isSet: true}
}

func (v NullableCustomerSegmentDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCustomerSegmentDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
