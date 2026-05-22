package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the BidRecommendationError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BidRecommendationError{}

// BidRecommendationError struct for BidRecommendationError
type BidRecommendationError struct {
	// A machine-readable error code.
	Code string `json:"code"`
	// A human-readable description of the response.
	Details string `json:"details"`
}

type _BidRecommendationError BidRecommendationError

// NewBidRecommendationError instantiates a new BidRecommendationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBidRecommendationError(code string, details string) *BidRecommendationError {
	this := BidRecommendationError{}
	this.Code = code
	this.Details = details
	return &this
}

// NewBidRecommendationErrorWithDefaults instantiates a new BidRecommendationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBidRecommendationErrorWithDefaults() *BidRecommendationError {
	this := BidRecommendationError{}
	return &this
}

// GetCode returns the Code field value
func (o *BidRecommendationError) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *BidRecommendationError) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *BidRecommendationError) SetCode(v string) {
	o.Code = v
}

// GetDetails returns the Details field value
func (o *BidRecommendationError) GetDetails() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *BidRecommendationError) GetDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *BidRecommendationError) SetDetails(v string) {
	o.Details = v
}

func (o BidRecommendationError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["details"] = o.Details
	return toSerialize, nil
}

type NullableBidRecommendationError struct {
	value *BidRecommendationError
	isSet bool
}

func (v NullableBidRecommendationError) Get() *BidRecommendationError {
	return v.value
}

func (v *NullableBidRecommendationError) Set(val *BidRecommendationError) {
	v.value = val
	v.isSet = true
}

func (v NullableBidRecommendationError) IsSet() bool {
	return v.isSet
}

func (v *NullableBidRecommendationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBidRecommendationError(val *BidRecommendationError) *NullableBidRecommendationError {
	return &NullableBidRecommendationError{value: val, isSet: true}
}

func (v NullableBidRecommendationError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBidRecommendationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
