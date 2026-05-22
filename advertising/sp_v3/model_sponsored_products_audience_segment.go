package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsAudienceSegment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsAudienceSegment{}

// SponsoredProductsAudienceSegment struct for SponsoredProductsAudienceSegment
type SponsoredProductsAudienceSegment struct {
	// `audienceId` is specified based on the `audienceSegmentType` used.
	AudienceId          string                               `json:"audienceId"`
	AudienceSegmentType SponsoredProductsAudienceSegmentType `json:"audienceSegmentType"`
}

type _SponsoredProductsAudienceSegment SponsoredProductsAudienceSegment

// NewSponsoredProductsAudienceSegment instantiates a new SponsoredProductsAudienceSegment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsAudienceSegment(audienceId string, audienceSegmentType SponsoredProductsAudienceSegmentType) *SponsoredProductsAudienceSegment {
	this := SponsoredProductsAudienceSegment{}
	this.AudienceId = audienceId
	this.AudienceSegmentType = audienceSegmentType
	return &this
}

// NewSponsoredProductsAudienceSegmentWithDefaults instantiates a new SponsoredProductsAudienceSegment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsAudienceSegmentWithDefaults() *SponsoredProductsAudienceSegment {
	this := SponsoredProductsAudienceSegment{}
	return &this
}

// GetAudienceId returns the AudienceId field value
func (o *SponsoredProductsAudienceSegment) GetAudienceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AudienceId
}

// GetAudienceIdOk returns a tuple with the AudienceId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAudienceSegment) GetAudienceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AudienceId, true
}

// SetAudienceId sets field value
func (o *SponsoredProductsAudienceSegment) SetAudienceId(v string) {
	o.AudienceId = v
}

// GetAudienceSegmentType returns the AudienceSegmentType field value
func (o *SponsoredProductsAudienceSegment) GetAudienceSegmentType() SponsoredProductsAudienceSegmentType {
	if o == nil {
		var ret SponsoredProductsAudienceSegmentType
		return ret
	}

	return o.AudienceSegmentType
}

// GetAudienceSegmentTypeOk returns a tuple with the AudienceSegmentType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAudienceSegment) GetAudienceSegmentTypeOk() (*SponsoredProductsAudienceSegmentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AudienceSegmentType, true
}

// SetAudienceSegmentType sets field value
func (o *SponsoredProductsAudienceSegment) SetAudienceSegmentType(v SponsoredProductsAudienceSegmentType) {
	o.AudienceSegmentType = v
}

func (o SponsoredProductsAudienceSegment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["audienceId"] = o.AudienceId
	toSerialize["audienceSegmentType"] = o.AudienceSegmentType
	return toSerialize, nil
}

type NullableSponsoredProductsAudienceSegment struct {
	value *SponsoredProductsAudienceSegment
	isSet bool
}

func (v NullableSponsoredProductsAudienceSegment) Get() *SponsoredProductsAudienceSegment {
	return v.value
}

func (v *NullableSponsoredProductsAudienceSegment) Set(val *SponsoredProductsAudienceSegment) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsAudienceSegment) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsAudienceSegment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsAudienceSegment(val *SponsoredProductsAudienceSegment) *NullableSponsoredProductsAudienceSegment {
	return &NullableSponsoredProductsAudienceSegment{value: val, isSet: true}
}

func (v NullableSponsoredProductsAudienceSegment) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsAudienceSegment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
