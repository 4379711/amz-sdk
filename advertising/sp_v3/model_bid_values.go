package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the BidValues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BidValues{}

// BidValues Suggested bid range
type BidValues struct {
	// The suggested bid
	Suggested *float64 `json:"suggested,omitempty"`
	// The bid range start
	RangeStart *float64 `json:"rangeStart,omitempty"`
	// The bid range end
	RangeEnd *float64 `json:"rangeEnd,omitempty"`
}

// NewBidValues instantiates a new BidValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBidValues() *BidValues {
	this := BidValues{}
	return &this
}

// NewBidValuesWithDefaults instantiates a new BidValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBidValuesWithDefaults() *BidValues {
	this := BidValues{}
	return &this
}

// GetSuggested returns the Suggested field value if set, zero value otherwise.
func (o *BidValues) GetSuggested() float64 {
	if o == nil || IsNil(o.Suggested) {
		var ret float64
		return ret
	}
	return *o.Suggested
}

// GetSuggestedOk returns a tuple with the Suggested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BidValues) GetSuggestedOk() (*float64, bool) {
	if o == nil || IsNil(o.Suggested) {
		return nil, false
	}
	return o.Suggested, true
}

// HasSuggested returns a boolean if a field has been set.
func (o *BidValues) HasSuggested() bool {
	if o != nil && !IsNil(o.Suggested) {
		return true
	}

	return false
}

// SetSuggested gets a reference to the given float64 and assigns it to the Suggested field.
func (o *BidValues) SetSuggested(v float64) {
	o.Suggested = &v
}

// GetRangeStart returns the RangeStart field value if set, zero value otherwise.
func (o *BidValues) GetRangeStart() float64 {
	if o == nil || IsNil(o.RangeStart) {
		var ret float64
		return ret
	}
	return *o.RangeStart
}

// GetRangeStartOk returns a tuple with the RangeStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BidValues) GetRangeStartOk() (*float64, bool) {
	if o == nil || IsNil(o.RangeStart) {
		return nil, false
	}
	return o.RangeStart, true
}

// HasRangeStart returns a boolean if a field has been set.
func (o *BidValues) HasRangeStart() bool {
	if o != nil && !IsNil(o.RangeStart) {
		return true
	}

	return false
}

// SetRangeStart gets a reference to the given float64 and assigns it to the RangeStart field.
func (o *BidValues) SetRangeStart(v float64) {
	o.RangeStart = &v
}

// GetRangeEnd returns the RangeEnd field value if set, zero value otherwise.
func (o *BidValues) GetRangeEnd() float64 {
	if o == nil || IsNil(o.RangeEnd) {
		var ret float64
		return ret
	}
	return *o.RangeEnd
}

// GetRangeEndOk returns a tuple with the RangeEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BidValues) GetRangeEndOk() (*float64, bool) {
	if o == nil || IsNil(o.RangeEnd) {
		return nil, false
	}
	return o.RangeEnd, true
}

// HasRangeEnd returns a boolean if a field has been set.
func (o *BidValues) HasRangeEnd() bool {
	if o != nil && !IsNil(o.RangeEnd) {
		return true
	}

	return false
}

// SetRangeEnd gets a reference to the given float64 and assigns it to the RangeEnd field.
func (o *BidValues) SetRangeEnd(v float64) {
	o.RangeEnd = &v
}

func (o BidValues) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Suggested) {
		toSerialize["suggested"] = o.Suggested
	}
	if !IsNil(o.RangeStart) {
		toSerialize["rangeStart"] = o.RangeStart
	}
	if !IsNil(o.RangeEnd) {
		toSerialize["rangeEnd"] = o.RangeEnd
	}
	return toSerialize, nil
}

type NullableBidValues struct {
	value *BidValues
	isSet bool
}

func (v NullableBidValues) Get() *BidValues {
	return v.value
}

func (v *NullableBidValues) Set(val *BidValues) {
	v.value = val
	v.isSet = true
}

func (v NullableBidValues) IsSet() bool {
	return v.isSet
}

func (v *NullableBidValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBidValues(val *BidValues) *NullableBidValues {
	return &NullableBidValues{value: val, isSet: true}
}

func (v NullableBidValues) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBidValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
