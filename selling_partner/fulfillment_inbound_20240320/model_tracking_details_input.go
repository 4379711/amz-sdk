package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the TrackingDetailsInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrackingDetailsInput{}

// TrackingDetailsInput Tracking information input for Less-Than-Truckload (LTL) and Small Parcel Delivery (SPD) shipments.
type TrackingDetailsInput struct {
	LtlTrackingDetail *LtlTrackingDetailInput `json:"ltlTrackingDetail,omitempty"`
	SpdTrackingDetail *SpdTrackingDetailInput `json:"spdTrackingDetail,omitempty"`
}

// NewTrackingDetailsInput instantiates a new TrackingDetailsInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrackingDetailsInput() *TrackingDetailsInput {
	this := TrackingDetailsInput{}
	return &this
}

// NewTrackingDetailsInputWithDefaults instantiates a new TrackingDetailsInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrackingDetailsInputWithDefaults() *TrackingDetailsInput {
	this := TrackingDetailsInput{}
	return &this
}

// GetLtlTrackingDetail returns the LtlTrackingDetail field value if set, zero value otherwise.
func (o *TrackingDetailsInput) GetLtlTrackingDetail() LtlTrackingDetailInput {
	if o == nil || IsNil(o.LtlTrackingDetail) {
		var ret LtlTrackingDetailInput
		return ret
	}
	return *o.LtlTrackingDetail
}

// GetLtlTrackingDetailOk returns a tuple with the LtlTrackingDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingDetailsInput) GetLtlTrackingDetailOk() (*LtlTrackingDetailInput, bool) {
	if o == nil || IsNil(o.LtlTrackingDetail) {
		return nil, false
	}
	return o.LtlTrackingDetail, true
}

// HasLtlTrackingDetail returns a boolean if a field has been set.
func (o *TrackingDetailsInput) HasLtlTrackingDetail() bool {
	if o != nil && !IsNil(o.LtlTrackingDetail) {
		return true
	}

	return false
}

// SetLtlTrackingDetail gets a reference to the given LtlTrackingDetailInput and assigns it to the LtlTrackingDetail field.
func (o *TrackingDetailsInput) SetLtlTrackingDetail(v LtlTrackingDetailInput) {
	o.LtlTrackingDetail = &v
}

// GetSpdTrackingDetail returns the SpdTrackingDetail field value if set, zero value otherwise.
func (o *TrackingDetailsInput) GetSpdTrackingDetail() SpdTrackingDetailInput {
	if o == nil || IsNil(o.SpdTrackingDetail) {
		var ret SpdTrackingDetailInput
		return ret
	}
	return *o.SpdTrackingDetail
}

// GetSpdTrackingDetailOk returns a tuple with the SpdTrackingDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingDetailsInput) GetSpdTrackingDetailOk() (*SpdTrackingDetailInput, bool) {
	if o == nil || IsNil(o.SpdTrackingDetail) {
		return nil, false
	}
	return o.SpdTrackingDetail, true
}

// HasSpdTrackingDetail returns a boolean if a field has been set.
func (o *TrackingDetailsInput) HasSpdTrackingDetail() bool {
	if o != nil && !IsNil(o.SpdTrackingDetail) {
		return true
	}

	return false
}

// SetSpdTrackingDetail gets a reference to the given SpdTrackingDetailInput and assigns it to the SpdTrackingDetail field.
func (o *TrackingDetailsInput) SetSpdTrackingDetail(v SpdTrackingDetailInput) {
	o.SpdTrackingDetail = &v
}

func (o TrackingDetailsInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LtlTrackingDetail) {
		toSerialize["ltlTrackingDetail"] = o.LtlTrackingDetail
	}
	if !IsNil(o.SpdTrackingDetail) {
		toSerialize["spdTrackingDetail"] = o.SpdTrackingDetail
	}
	return toSerialize, nil
}

type NullableTrackingDetailsInput struct {
	value *TrackingDetailsInput
	isSet bool
}

func (v NullableTrackingDetailsInput) Get() *TrackingDetailsInput {
	return v.value
}

func (v *NullableTrackingDetailsInput) Set(val *TrackingDetailsInput) {
	v.value = val
	v.isSet = true
}

func (v NullableTrackingDetailsInput) IsSet() bool {
	return v.isSet
}

func (v *NullableTrackingDetailsInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrackingDetailsInput(val *TrackingDetailsInput) *NullableTrackingDetailsInput {
	return &NullableTrackingDetailsInput{value: val, isSet: true}
}

func (v NullableTrackingDetailsInput) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTrackingDetailsInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
