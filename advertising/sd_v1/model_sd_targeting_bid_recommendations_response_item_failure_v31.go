package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDTargetingBidRecommendationsResponseItemFailureV31 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDTargetingBidRecommendationsResponseItemFailureV31{}

// SDTargetingBidRecommendationsResponseItemFailureV31 Failed bid recommendation response.
type SDTargetingBidRecommendationsResponseItemFailureV31 struct {
	// The HTTP status code of this item.
	Code string `json:"code"`
	// A human-readable description of this item on error.
	Details string `json:"details"`
}

type _SDTargetingBidRecommendationsResponseItemFailureV31 SDTargetingBidRecommendationsResponseItemFailureV31

// NewSDTargetingBidRecommendationsResponseItemFailureV31 instantiates a new SDTargetingBidRecommendationsResponseItemFailureV31 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDTargetingBidRecommendationsResponseItemFailureV31(code string, details string) *SDTargetingBidRecommendationsResponseItemFailureV31 {
	this := SDTargetingBidRecommendationsResponseItemFailureV31{}
	this.Code = code
	this.Details = details
	return &this
}

// NewSDTargetingBidRecommendationsResponseItemFailureV31WithDefaults instantiates a new SDTargetingBidRecommendationsResponseItemFailureV31 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDTargetingBidRecommendationsResponseItemFailureV31WithDefaults() *SDTargetingBidRecommendationsResponseItemFailureV31 {
	this := SDTargetingBidRecommendationsResponseItemFailureV31{}
	return &this
}

// GetCode returns the Code field value
func (o *SDTargetingBidRecommendationsResponseItemFailureV31) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SDTargetingBidRecommendationsResponseItemFailureV31) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SDTargetingBidRecommendationsResponseItemFailureV31) SetCode(v string) {
	o.Code = v
}

// GetDetails returns the Details field value
func (o *SDTargetingBidRecommendationsResponseItemFailureV31) GetDetails() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *SDTargetingBidRecommendationsResponseItemFailureV31) GetDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *SDTargetingBidRecommendationsResponseItemFailureV31) SetDetails(v string) {
	o.Details = v
}

func (o SDTargetingBidRecommendationsResponseItemFailureV31) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["details"] = o.Details
	return toSerialize, nil
}

type NullableSDTargetingBidRecommendationsResponseItemFailureV31 struct {
	value *SDTargetingBidRecommendationsResponseItemFailureV31
	isSet bool
}

func (v NullableSDTargetingBidRecommendationsResponseItemFailureV31) Get() *SDTargetingBidRecommendationsResponseItemFailureV31 {
	return v.value
}

func (v *NullableSDTargetingBidRecommendationsResponseItemFailureV31) Set(val *SDTargetingBidRecommendationsResponseItemFailureV31) {
	v.value = val
	v.isSet = true
}

func (v NullableSDTargetingBidRecommendationsResponseItemFailureV31) IsSet() bool {
	return v.isSet
}

func (v *NullableSDTargetingBidRecommendationsResponseItemFailureV31) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDTargetingBidRecommendationsResponseItemFailureV31(val *SDTargetingBidRecommendationsResponseItemFailureV31) *NullableSDTargetingBidRecommendationsResponseItemFailureV31 {
	return &NullableSDTargetingBidRecommendationsResponseItemFailureV31{value: val, isSet: true}
}

func (v NullableSDTargetingBidRecommendationsResponseItemFailureV31) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDTargetingBidRecommendationsResponseItemFailureV31) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
