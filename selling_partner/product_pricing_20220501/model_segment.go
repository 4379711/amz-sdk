package product_pricing_20220501

import (
	"github.com/bytedance/sonic"
)

// checks if the Segment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Segment{}

// Segment Input segment for featured offer expected price. The segment contains the location information for which featured offer expected price is requested.
type Segment struct {
	SegmentDetails *SegmentDetails `json:"segmentDetails,omitempty"`
}

// NewSegment instantiates a new Segment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSegment() *Segment {
	this := Segment{}
	return &this
}

// NewSegmentWithDefaults instantiates a new Segment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSegmentWithDefaults() *Segment {
	this := Segment{}
	return &this
}

// GetSegmentDetails returns the SegmentDetails field value if set, zero value otherwise.
func (o *Segment) GetSegmentDetails() SegmentDetails {
	if o == nil || IsNil(o.SegmentDetails) {
		var ret SegmentDetails
		return ret
	}
	return *o.SegmentDetails
}

// GetSegmentDetailsOk returns a tuple with the SegmentDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Segment) GetSegmentDetailsOk() (*SegmentDetails, bool) {
	if o == nil || IsNil(o.SegmentDetails) {
		return nil, false
	}
	return o.SegmentDetails, true
}

// HasSegmentDetails returns a boolean if a field has been set.
func (o *Segment) HasSegmentDetails() bool {
	if o != nil && !IsNil(o.SegmentDetails) {
		return true
	}

	return false
}

// SetSegmentDetails gets a reference to the given SegmentDetails and assigns it to the SegmentDetails field.
func (o *Segment) SetSegmentDetails(v SegmentDetails) {
	o.SegmentDetails = &v
}

func (o Segment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SegmentDetails) {
		toSerialize["segmentDetails"] = o.SegmentDetails
	}
	return toSerialize, nil
}

type NullableSegment struct {
	value *Segment
	isSet bool
}

func (v NullableSegment) Get() *Segment {
	return v.value
}

func (v *NullableSegment) Set(val *Segment) {
	v.value = val
	v.isSet = true
}

func (v NullableSegment) IsSet() bool {
	return v.isSet
}

func (v *NullableSegment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSegment(val *Segment) *NullableSegment {
	return &NullableSegment{value: val, isSet: true}
}

func (v NullableSegment) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSegment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
