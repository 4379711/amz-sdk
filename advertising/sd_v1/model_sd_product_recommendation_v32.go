package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDProductRecommendationV32 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDProductRecommendationV32{}

// SDProductRecommendationV32 A recommended product to target ads on
type SDProductRecommendationV32 struct {
	// Amazon Standard Identification Number
	Asin *string `json:"asin,omitempty" validate:"regexp=[a-zA-Z0-9]{10}"`
	// A rank to signify which recommendations are weighed more heavily, with a lower rank signifying a stronger recommendation
	Rank *int32 `json:"rank,omitempty"`
	// The top advertised products this recommendation is made for.
	AdvertisedAsins []string `json:"advertisedAsins,omitempty"`
}

// NewSDProductRecommendationV32 instantiates a new SDProductRecommendationV32 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDProductRecommendationV32() *SDProductRecommendationV32 {
	this := SDProductRecommendationV32{}
	return &this
}

// NewSDProductRecommendationV32WithDefaults instantiates a new SDProductRecommendationV32 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDProductRecommendationV32WithDefaults() *SDProductRecommendationV32 {
	this := SDProductRecommendationV32{}
	return &this
}

// GetAsin returns the Asin field value if set, zero value otherwise.
func (o *SDProductRecommendationV32) GetAsin() string {
	if o == nil || IsNil(o.Asin) {
		var ret string
		return ret
	}
	return *o.Asin
}

// GetAsinOk returns a tuple with the Asin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDProductRecommendationV32) GetAsinOk() (*string, bool) {
	if o == nil || IsNil(o.Asin) {
		return nil, false
	}
	return o.Asin, true
}

// HasAsin returns a boolean if a field has been set.
func (o *SDProductRecommendationV32) HasAsin() bool {
	if o != nil && !IsNil(o.Asin) {
		return true
	}

	return false
}

// SetAsin gets a reference to the given string and assigns it to the Asin field.
func (o *SDProductRecommendationV32) SetAsin(v string) {
	o.Asin = &v
}

// GetRank returns the Rank field value if set, zero value otherwise.
func (o *SDProductRecommendationV32) GetRank() int32 {
	if o == nil || IsNil(o.Rank) {
		var ret int32
		return ret
	}
	return *o.Rank
}

// GetRankOk returns a tuple with the Rank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDProductRecommendationV32) GetRankOk() (*int32, bool) {
	if o == nil || IsNil(o.Rank) {
		return nil, false
	}
	return o.Rank, true
}

// HasRank returns a boolean if a field has been set.
func (o *SDProductRecommendationV32) HasRank() bool {
	if o != nil && !IsNil(o.Rank) {
		return true
	}

	return false
}

// SetRank gets a reference to the given int32 and assigns it to the Rank field.
func (o *SDProductRecommendationV32) SetRank(v int32) {
	o.Rank = &v
}

// GetAdvertisedAsins returns the AdvertisedAsins field value if set, zero value otherwise.
func (o *SDProductRecommendationV32) GetAdvertisedAsins() []string {
	if o == nil || IsNil(o.AdvertisedAsins) {
		var ret []string
		return ret
	}
	return o.AdvertisedAsins
}

// GetAdvertisedAsinsOk returns a tuple with the AdvertisedAsins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDProductRecommendationV32) GetAdvertisedAsinsOk() ([]string, bool) {
	if o == nil || IsNil(o.AdvertisedAsins) {
		return nil, false
	}
	return o.AdvertisedAsins, true
}

// HasAdvertisedAsins returns a boolean if a field has been set.
func (o *SDProductRecommendationV32) HasAdvertisedAsins() bool {
	if o != nil && !IsNil(o.AdvertisedAsins) {
		return true
	}

	return false
}

// SetAdvertisedAsins gets a reference to the given []string and assigns it to the AdvertisedAsins field.
func (o *SDProductRecommendationV32) SetAdvertisedAsins(v []string) {
	o.AdvertisedAsins = v
}

func (o SDProductRecommendationV32) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Asin) {
		toSerialize["asin"] = o.Asin
	}
	if !IsNil(o.Rank) {
		toSerialize["rank"] = o.Rank
	}
	if !IsNil(o.AdvertisedAsins) {
		toSerialize["advertisedAsins"] = o.AdvertisedAsins
	}
	return toSerialize, nil
}

type NullableSDProductRecommendationV32 struct {
	value *SDProductRecommendationV32
	isSet bool
}

func (v NullableSDProductRecommendationV32) Get() *SDProductRecommendationV32 {
	return v.value
}

func (v *NullableSDProductRecommendationV32) Set(val *SDProductRecommendationV32) {
	v.value = val
	v.isSet = true
}

func (v NullableSDProductRecommendationV32) IsSet() bool {
	return v.isSet
}

func (v *NullableSDProductRecommendationV32) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDProductRecommendationV32(val *SDProductRecommendationV32) *NullableSDProductRecommendationV32 {
	return &NullableSDProductRecommendationV32{value: val, isSet: true}
}

func (v NullableSDProductRecommendationV32) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDProductRecommendationV32) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
