package promotions_20251201

import (
	"github.com/bytedance/sonic"
)

// checks if the BrandSegmentDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BrandSegmentDetails{}

// BrandSegmentDetails Additional details specific to the `BRAND` segments type.
type BrandSegmentDetails struct {
	// Brand identifier.
	BrandId string `json:"brandId"`
}

// NewBrandSegmentDetails instantiates a new BrandSegmentDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrandSegmentDetails(brandId string) *BrandSegmentDetails {
	this := BrandSegmentDetails{}
	this.BrandId = brandId
	return &this
}

// NewBrandSegmentDetailsWithDefaults instantiates a new BrandSegmentDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandSegmentDetailsWithDefaults() *BrandSegmentDetails {
	this := BrandSegmentDetails{}
	return &this
}

// GetBrandId returns the BrandId field value
func (o *BrandSegmentDetails) GetBrandId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BrandId
}

// GetBrandIdOk returns a tuple with the BrandId field value
// and a boolean to check if the value has been set.
func (o *BrandSegmentDetails) GetBrandIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrandId, true
}

// SetBrandId sets field value
func (o *BrandSegmentDetails) SetBrandId(v string) {
	o.BrandId = v
}

func (o BrandSegmentDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["brandId"] = o.BrandId
	return toSerialize, nil
}

type NullableBrandSegmentDetails struct {
	value *BrandSegmentDetails
	isSet bool
}

func (v NullableBrandSegmentDetails) Get() *BrandSegmentDetails {
	return v.value
}

func (v *NullableBrandSegmentDetails) Set(val *BrandSegmentDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandSegmentDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandSegmentDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandSegmentDetails(val *BrandSegmentDetails) *NullableBrandSegmentDetails {
	return &NullableBrandSegmentDetails{value: val, isSet: true}
}

func (v NullableBrandSegmentDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBrandSegmentDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
