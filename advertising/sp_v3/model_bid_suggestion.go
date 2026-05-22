package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the BidSuggestion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BidSuggestion{}

// BidSuggestion Suggested bid range in major and minor currency units (example: dollars and cents).
type BidSuggestion struct {
	// The suggested bid
	Suggested *float64 `json:"suggested,omitempty"`
	// The bid range start
	RangeStart *float64 `json:"rangeStart,omitempty"`
	// The bid recommendation id
	BidRecId *string `json:"bidRecId,omitempty"`
	// The bid range end
	RangeEnd *float64 `json:"rangeEnd,omitempty"`
}

// NewBidSuggestion instantiates a new BidSuggestion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBidSuggestion() *BidSuggestion {
	this := BidSuggestion{}
	return &this
}

// NewBidSuggestionWithDefaults instantiates a new BidSuggestion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBidSuggestionWithDefaults() *BidSuggestion {
	this := BidSuggestion{}
	return &this
}

// GetSuggested returns the Suggested field value if set, zero value otherwise.
func (o *BidSuggestion) GetSuggested() float64 {
	if o == nil || IsNil(o.Suggested) {
		var ret float64
		return ret
	}
	return *o.Suggested
}

// GetSuggestedOk returns a tuple with the Suggested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BidSuggestion) GetSuggestedOk() (*float64, bool) {
	if o == nil || IsNil(o.Suggested) {
		return nil, false
	}
	return o.Suggested, true
}

// HasSuggested returns a boolean if a field has been set.
func (o *BidSuggestion) HasSuggested() bool {
	if o != nil && !IsNil(o.Suggested) {
		return true
	}

	return false
}

// SetSuggested gets a reference to the given float64 and assigns it to the Suggested field.
func (o *BidSuggestion) SetSuggested(v float64) {
	o.Suggested = &v
}

// GetRangeStart returns the RangeStart field value if set, zero value otherwise.
func (o *BidSuggestion) GetRangeStart() float64 {
	if o == nil || IsNil(o.RangeStart) {
		var ret float64
		return ret
	}
	return *o.RangeStart
}

// GetRangeStartOk returns a tuple with the RangeStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BidSuggestion) GetRangeStartOk() (*float64, bool) {
	if o == nil || IsNil(o.RangeStart) {
		return nil, false
	}
	return o.RangeStart, true
}

// HasRangeStart returns a boolean if a field has been set.
func (o *BidSuggestion) HasRangeStart() bool {
	if o != nil && !IsNil(o.RangeStart) {
		return true
	}

	return false
}

// SetRangeStart gets a reference to the given float64 and assigns it to the RangeStart field.
func (o *BidSuggestion) SetRangeStart(v float64) {
	o.RangeStart = &v
}

// GetBidRecId returns the BidRecId field value if set, zero value otherwise.
func (o *BidSuggestion) GetBidRecId() string {
	if o == nil || IsNil(o.BidRecId) {
		var ret string
		return ret
	}
	return *o.BidRecId
}

// GetBidRecIdOk returns a tuple with the BidRecId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BidSuggestion) GetBidRecIdOk() (*string, bool) {
	if o == nil || IsNil(o.BidRecId) {
		return nil, false
	}
	return o.BidRecId, true
}

// HasBidRecId returns a boolean if a field has been set.
func (o *BidSuggestion) HasBidRecId() bool {
	if o != nil && !IsNil(o.BidRecId) {
		return true
	}

	return false
}

// SetBidRecId gets a reference to the given string and assigns it to the BidRecId field.
func (o *BidSuggestion) SetBidRecId(v string) {
	o.BidRecId = &v
}

// GetRangeEnd returns the RangeEnd field value if set, zero value otherwise.
func (o *BidSuggestion) GetRangeEnd() float64 {
	if o == nil || IsNil(o.RangeEnd) {
		var ret float64
		return ret
	}
	return *o.RangeEnd
}

// GetRangeEndOk returns a tuple with the RangeEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BidSuggestion) GetRangeEndOk() (*float64, bool) {
	if o == nil || IsNil(o.RangeEnd) {
		return nil, false
	}
	return o.RangeEnd, true
}

// HasRangeEnd returns a boolean if a field has been set.
func (o *BidSuggestion) HasRangeEnd() bool {
	if o != nil && !IsNil(o.RangeEnd) {
		return true
	}

	return false
}

// SetRangeEnd gets a reference to the given float64 and assigns it to the RangeEnd field.
func (o *BidSuggestion) SetRangeEnd(v float64) {
	o.RangeEnd = &v
}

func (o BidSuggestion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Suggested) {
		toSerialize["suggested"] = o.Suggested
	}
	if !IsNil(o.RangeStart) {
		toSerialize["rangeStart"] = o.RangeStart
	}
	if !IsNil(o.BidRecId) {
		toSerialize["bidRecId"] = o.BidRecId
	}
	if !IsNil(o.RangeEnd) {
		toSerialize["rangeEnd"] = o.RangeEnd
	}
	return toSerialize, nil
}

type NullableBidSuggestion struct {
	value *BidSuggestion
	isSet bool
}

func (v NullableBidSuggestion) Get() *BidSuggestion {
	return v.value
}

func (v *NullableBidSuggestion) Set(val *BidSuggestion) {
	v.value = val
	v.isSet = true
}

func (v NullableBidSuggestion) IsSet() bool {
	return v.isSet
}

func (v *NullableBidSuggestion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBidSuggestion(val *BidSuggestion) *NullableBidSuggestion {
	return &NullableBidSuggestion{value: val, isSet: true}
}

func (v NullableBidSuggestion) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBidSuggestion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
