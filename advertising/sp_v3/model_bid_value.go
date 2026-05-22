package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the BidValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BidValue{}

// BidValue Bid value of the bid recommendations.
type BidValue struct {
	// The suggested bid.
	SuggestedBid float64 `json:"suggestedBid"`
}

type _BidValue BidValue

// NewBidValue instantiates a new BidValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBidValue(suggestedBid float64) *BidValue {
	this := BidValue{}
	this.SuggestedBid = suggestedBid
	return &this
}

// NewBidValueWithDefaults instantiates a new BidValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBidValueWithDefaults() *BidValue {
	this := BidValue{}
	return &this
}

// GetSuggestedBid returns the SuggestedBid field value
func (o *BidValue) GetSuggestedBid() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.SuggestedBid
}

// GetSuggestedBidOk returns a tuple with the SuggestedBid field value
// and a boolean to check if the value has been set.
func (o *BidValue) GetSuggestedBidOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SuggestedBid, true
}

// SetSuggestedBid sets field value
func (o *BidValue) SetSuggestedBid(v float64) {
	o.SuggestedBid = v
}

func (o BidValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["suggestedBid"] = o.SuggestedBid
	return toSerialize, nil
}

type NullableBidValue struct {
	value *BidValue
	isSet bool
}

func (v NullableBidValue) Get() *BidValue {
	return v.value
}

func (v *NullableBidValue) Set(val *BidValue) {
	v.value = val
	v.isSet = true
}

func (v NullableBidValue) IsSet() bool {
	return v.isSet
}

func (v *NullableBidValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBidValue(val *BidValue) *NullableBidValue {
	return &NullableBidValue{value: val, isSet: true}
}

func (v NullableBidValue) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBidValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
