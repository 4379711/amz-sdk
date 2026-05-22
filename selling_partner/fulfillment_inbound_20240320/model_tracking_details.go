package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the TrackingDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrackingDetails{}

// TrackingDetails Tracking information for Less-Than-Truckload (LTL) and Small Parcel Delivery (SPD) shipments.
type TrackingDetails struct {
	LtlTrackingDetail *LtlTrackingDetail `json:"ltlTrackingDetail,omitempty"`
	SpdTrackingDetail *SpdTrackingDetail `json:"spdTrackingDetail,omitempty"`
}

// NewTrackingDetails instantiates a new TrackingDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrackingDetails() *TrackingDetails {
	this := TrackingDetails{}
	return &this
}

// NewTrackingDetailsWithDefaults instantiates a new TrackingDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrackingDetailsWithDefaults() *TrackingDetails {
	this := TrackingDetails{}
	return &this
}

// GetLtlTrackingDetail returns the LtlTrackingDetail field value if set, zero value otherwise.
func (o *TrackingDetails) GetLtlTrackingDetail() LtlTrackingDetail {
	if o == nil || IsNil(o.LtlTrackingDetail) {
		var ret LtlTrackingDetail
		return ret
	}
	return *o.LtlTrackingDetail
}

// GetLtlTrackingDetailOk returns a tuple with the LtlTrackingDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingDetails) GetLtlTrackingDetailOk() (*LtlTrackingDetail, bool) {
	if o == nil || IsNil(o.LtlTrackingDetail) {
		return nil, false
	}
	return o.LtlTrackingDetail, true
}

// HasLtlTrackingDetail returns a boolean if a field has been set.
func (o *TrackingDetails) HasLtlTrackingDetail() bool {
	if o != nil && !IsNil(o.LtlTrackingDetail) {
		return true
	}

	return false
}

// SetLtlTrackingDetail gets a reference to the given LtlTrackingDetail and assigns it to the LtlTrackingDetail field.
func (o *TrackingDetails) SetLtlTrackingDetail(v LtlTrackingDetail) {
	o.LtlTrackingDetail = &v
}

// GetSpdTrackingDetail returns the SpdTrackingDetail field value if set, zero value otherwise.
func (o *TrackingDetails) GetSpdTrackingDetail() SpdTrackingDetail {
	if o == nil || IsNil(o.SpdTrackingDetail) {
		var ret SpdTrackingDetail
		return ret
	}
	return *o.SpdTrackingDetail
}

// GetSpdTrackingDetailOk returns a tuple with the SpdTrackingDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingDetails) GetSpdTrackingDetailOk() (*SpdTrackingDetail, bool) {
	if o == nil || IsNil(o.SpdTrackingDetail) {
		return nil, false
	}
	return o.SpdTrackingDetail, true
}

// HasSpdTrackingDetail returns a boolean if a field has been set.
func (o *TrackingDetails) HasSpdTrackingDetail() bool {
	if o != nil && !IsNil(o.SpdTrackingDetail) {
		return true
	}

	return false
}

// SetSpdTrackingDetail gets a reference to the given SpdTrackingDetail and assigns it to the SpdTrackingDetail field.
func (o *TrackingDetails) SetSpdTrackingDetail(v SpdTrackingDetail) {
	o.SpdTrackingDetail = &v
}

func (o TrackingDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LtlTrackingDetail) {
		toSerialize["ltlTrackingDetail"] = o.LtlTrackingDetail
	}
	if !IsNil(o.SpdTrackingDetail) {
		toSerialize["spdTrackingDetail"] = o.SpdTrackingDetail
	}
	return toSerialize, nil
}

type NullableTrackingDetails struct {
	value *TrackingDetails
	isSet bool
}

func (v NullableTrackingDetails) Get() *TrackingDetails {
	return v.value
}

func (v *NullableTrackingDetails) Set(val *TrackingDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableTrackingDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableTrackingDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrackingDetails(val *TrackingDetails) *NullableTrackingDetails {
	return &NullableTrackingDetails{value: val, isSet: true}
}

func (v NullableTrackingDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTrackingDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
