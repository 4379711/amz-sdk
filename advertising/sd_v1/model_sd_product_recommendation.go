package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDProductRecommendation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDProductRecommendation{}

// SDProductRecommendation A recommended product to target ads on
type SDProductRecommendation struct {
	// Amazon Standard Identification Number
	Asin *string `json:"asin,omitempty" validate:"regexp=[a-zA-Z0-9]{10}"`
	// A rank to signify which recommendations are weighed more heavily, with a lower rank signifying a stronger recommendation
	Rank *int32 `json:"rank,omitempty"`
}

// NewSDProductRecommendation instantiates a new SDProductRecommendation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDProductRecommendation() *SDProductRecommendation {
	this := SDProductRecommendation{}
	return &this
}

// NewSDProductRecommendationWithDefaults instantiates a new SDProductRecommendation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDProductRecommendationWithDefaults() *SDProductRecommendation {
	this := SDProductRecommendation{}
	return &this
}

// GetAsin returns the Asin field value if set, zero value otherwise.
func (o *SDProductRecommendation) GetAsin() string {
	if o == nil || IsNil(o.Asin) {
		var ret string
		return ret
	}
	return *o.Asin
}

// GetAsinOk returns a tuple with the Asin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDProductRecommendation) GetAsinOk() (*string, bool) {
	if o == nil || IsNil(o.Asin) {
		return nil, false
	}
	return o.Asin, true
}

// HasAsin returns a boolean if a field has been set.
func (o *SDProductRecommendation) HasAsin() bool {
	if o != nil && !IsNil(o.Asin) {
		return true
	}

	return false
}

// SetAsin gets a reference to the given string and assigns it to the Asin field.
func (o *SDProductRecommendation) SetAsin(v string) {
	o.Asin = &v
}

// GetRank returns the Rank field value if set, zero value otherwise.
func (o *SDProductRecommendation) GetRank() int32 {
	if o == nil || IsNil(o.Rank) {
		var ret int32
		return ret
	}
	return *o.Rank
}

// GetRankOk returns a tuple with the Rank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDProductRecommendation) GetRankOk() (*int32, bool) {
	if o == nil || IsNil(o.Rank) {
		return nil, false
	}
	return o.Rank, true
}

// HasRank returns a boolean if a field has been set.
func (o *SDProductRecommendation) HasRank() bool {
	if o != nil && !IsNil(o.Rank) {
		return true
	}

	return false
}

// SetRank gets a reference to the given int32 and assigns it to the Rank field.
func (o *SDProductRecommendation) SetRank(v int32) {
	o.Rank = &v
}

func (o SDProductRecommendation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Asin) {
		toSerialize["asin"] = o.Asin
	}
	if !IsNil(o.Rank) {
		toSerialize["rank"] = o.Rank
	}
	return toSerialize, nil
}

type NullableSDProductRecommendation struct {
	value *SDProductRecommendation
	isSet bool
}

func (v NullableSDProductRecommendation) Get() *SDProductRecommendation {
	return v.value
}

func (v *NullableSDProductRecommendation) Set(val *SDProductRecommendation) {
	v.value = val
	v.isSet = true
}

func (v NullableSDProductRecommendation) IsSet() bool {
	return v.isSet
}

func (v *NullableSDProductRecommendation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDProductRecommendation(val *SDProductRecommendation) *NullableSDProductRecommendation {
	return &NullableSDProductRecommendation{value: val, isSet: true}
}

func (v NullableSDProductRecommendation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDProductRecommendation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
