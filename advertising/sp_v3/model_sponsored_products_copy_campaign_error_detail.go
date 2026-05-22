package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCopyCampaignErrorDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCopyCampaignErrorDetail{}

// SponsoredProductsCopyCampaignErrorDetail struct for SponsoredProductsCopyCampaignErrorDetail
type SponsoredProductsCopyCampaignErrorDetail struct {
	// An enumerated success or error code for machine use
	Code *string `json:"code,omitempty"`
	// A human-readable description of the error, if unsuccessful
	Details *string `json:"details,omitempty"`
}

// NewSponsoredProductsCopyCampaignErrorDetail instantiates a new SponsoredProductsCopyCampaignErrorDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCopyCampaignErrorDetail() *SponsoredProductsCopyCampaignErrorDetail {
	this := SponsoredProductsCopyCampaignErrorDetail{}
	return &this
}

// NewSponsoredProductsCopyCampaignErrorDetailWithDefaults instantiates a new SponsoredProductsCopyCampaignErrorDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCopyCampaignErrorDetailWithDefaults() *SponsoredProductsCopyCampaignErrorDetail {
	this := SponsoredProductsCopyCampaignErrorDetail{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SponsoredProductsCopyCampaignErrorDetail) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCopyCampaignErrorDetail) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SponsoredProductsCopyCampaignErrorDetail) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *SponsoredProductsCopyCampaignErrorDetail) SetCode(v string) {
	o.Code = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *SponsoredProductsCopyCampaignErrorDetail) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCopyCampaignErrorDetail) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *SponsoredProductsCopyCampaignErrorDetail) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *SponsoredProductsCopyCampaignErrorDetail) SetDetails(v string) {
	o.Details = &v
}

func (o SponsoredProductsCopyCampaignErrorDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableSponsoredProductsCopyCampaignErrorDetail struct {
	value *SponsoredProductsCopyCampaignErrorDetail
	isSet bool
}

func (v NullableSponsoredProductsCopyCampaignErrorDetail) Get() *SponsoredProductsCopyCampaignErrorDetail {
	return v.value
}

func (v *NullableSponsoredProductsCopyCampaignErrorDetail) Set(val *SponsoredProductsCopyCampaignErrorDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCopyCampaignErrorDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCopyCampaignErrorDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCopyCampaignErrorDetail(val *SponsoredProductsCopyCampaignErrorDetail) *NullableSponsoredProductsCopyCampaignErrorDetail {
	return &NullableSponsoredProductsCopyCampaignErrorDetail{value: val, isSet: true}
}

func (v NullableSponsoredProductsCopyCampaignErrorDetail) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCopyCampaignErrorDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
