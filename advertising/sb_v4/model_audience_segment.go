package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the AudienceSegment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AudienceSegment{}

// AudienceSegment struct for AudienceSegment
type AudienceSegment struct {
	AudienceId          *string              `json:"audienceId,omitempty"`
	AudienceSegmentType *AudienceSegmentType `json:"audienceSegmentType,omitempty"`
}

// NewAudienceSegment instantiates a new AudienceSegment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAudienceSegment() *AudienceSegment {
	this := AudienceSegment{}
	return &this
}

// NewAudienceSegmentWithDefaults instantiates a new AudienceSegment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAudienceSegmentWithDefaults() *AudienceSegment {
	this := AudienceSegment{}
	return &this
}

// GetAudienceId returns the AudienceId field value if set, zero value otherwise.
func (o *AudienceSegment) GetAudienceId() string {
	if o == nil || IsNil(o.AudienceId) {
		var ret string
		return ret
	}
	return *o.AudienceId
}

// GetAudienceIdOk returns a tuple with the AudienceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AudienceSegment) GetAudienceIdOk() (*string, bool) {
	if o == nil || IsNil(o.AudienceId) {
		return nil, false
	}
	return o.AudienceId, true
}

// HasAudienceId returns a boolean if a field has been set.
func (o *AudienceSegment) HasAudienceId() bool {
	if o != nil && !IsNil(o.AudienceId) {
		return true
	}

	return false
}

// SetAudienceId gets a reference to the given string and assigns it to the AudienceId field.
func (o *AudienceSegment) SetAudienceId(v string) {
	o.AudienceId = &v
}

// GetAudienceSegmentType returns the AudienceSegmentType field value if set, zero value otherwise.
func (o *AudienceSegment) GetAudienceSegmentType() AudienceSegmentType {
	if o == nil || IsNil(o.AudienceSegmentType) {
		var ret AudienceSegmentType
		return ret
	}
	return *o.AudienceSegmentType
}

// GetAudienceSegmentTypeOk returns a tuple with the AudienceSegmentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AudienceSegment) GetAudienceSegmentTypeOk() (*AudienceSegmentType, bool) {
	if o == nil || IsNil(o.AudienceSegmentType) {
		return nil, false
	}
	return o.AudienceSegmentType, true
}

// HasAudienceSegmentType returns a boolean if a field has been set.
func (o *AudienceSegment) HasAudienceSegmentType() bool {
	if o != nil && !IsNil(o.AudienceSegmentType) {
		return true
	}

	return false
}

// SetAudienceSegmentType gets a reference to the given AudienceSegmentType and assigns it to the AudienceSegmentType field.
func (o *AudienceSegment) SetAudienceSegmentType(v AudienceSegmentType) {
	o.AudienceSegmentType = &v
}

func (o AudienceSegment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AudienceId) {
		toSerialize["audienceId"] = o.AudienceId
	}
	if !IsNil(o.AudienceSegmentType) {
		toSerialize["audienceSegmentType"] = o.AudienceSegmentType
	}
	return toSerialize, nil
}

type NullableAudienceSegment struct {
	value *AudienceSegment
	isSet bool
}

func (v NullableAudienceSegment) Get() *AudienceSegment {
	return v.value
}

func (v *NullableAudienceSegment) Set(val *AudienceSegment) {
	v.value = val
	v.isSet = true
}

func (v NullableAudienceSegment) IsSet() bool {
	return v.isSet
}

func (v *NullableAudienceSegment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAudienceSegment(val *AudienceSegment) *NullableAudienceSegment {
	return &NullableAudienceSegment{value: val, isSet: true}
}

func (v NullableAudienceSegment) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAudienceSegment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
