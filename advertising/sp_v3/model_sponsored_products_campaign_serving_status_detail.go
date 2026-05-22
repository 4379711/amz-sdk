package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCampaignServingStatusDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCampaignServingStatusDetail{}

// SponsoredProductsCampaignServingStatusDetail struct for SponsoredProductsCampaignServingStatusDetail
type SponsoredProductsCampaignServingStatusDetail struct {
	Name *SponsoredProductsCampaignServingStatusReason `json:"name,omitempty"`
	// A URL with additional information about the status identifier.
	HelpUrl *string `json:"helpUrl,omitempty"`
	// A human-readable description of the status identifier specified in the name field.
	Message *string `json:"message,omitempty"`
}

// NewSponsoredProductsCampaignServingStatusDetail instantiates a new SponsoredProductsCampaignServingStatusDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCampaignServingStatusDetail() *SponsoredProductsCampaignServingStatusDetail {
	this := SponsoredProductsCampaignServingStatusDetail{}
	return &this
}

// NewSponsoredProductsCampaignServingStatusDetailWithDefaults instantiates a new SponsoredProductsCampaignServingStatusDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCampaignServingStatusDetailWithDefaults() *SponsoredProductsCampaignServingStatusDetail {
	this := SponsoredProductsCampaignServingStatusDetail{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SponsoredProductsCampaignServingStatusDetail) GetName() SponsoredProductsCampaignServingStatusReason {
	if o == nil || IsNil(o.Name) {
		var ret SponsoredProductsCampaignServingStatusReason
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignServingStatusDetail) GetNameOk() (*SponsoredProductsCampaignServingStatusReason, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SponsoredProductsCampaignServingStatusDetail) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given SponsoredProductsCampaignServingStatusReason and assigns it to the Name field.
func (o *SponsoredProductsCampaignServingStatusDetail) SetName(v SponsoredProductsCampaignServingStatusReason) {
	o.Name = &v
}

// GetHelpUrl returns the HelpUrl field value if set, zero value otherwise.
func (o *SponsoredProductsCampaignServingStatusDetail) GetHelpUrl() string {
	if o == nil || IsNil(o.HelpUrl) {
		var ret string
		return ret
	}
	return *o.HelpUrl
}

// GetHelpUrlOk returns a tuple with the HelpUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignServingStatusDetail) GetHelpUrlOk() (*string, bool) {
	if o == nil || IsNil(o.HelpUrl) {
		return nil, false
	}
	return o.HelpUrl, true
}

// HasHelpUrl returns a boolean if a field has been set.
func (o *SponsoredProductsCampaignServingStatusDetail) HasHelpUrl() bool {
	if o != nil && !IsNil(o.HelpUrl) {
		return true
	}

	return false
}

// SetHelpUrl gets a reference to the given string and assigns it to the HelpUrl field.
func (o *SponsoredProductsCampaignServingStatusDetail) SetHelpUrl(v string) {
	o.HelpUrl = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SponsoredProductsCampaignServingStatusDetail) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignServingStatusDetail) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SponsoredProductsCampaignServingStatusDetail) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *SponsoredProductsCampaignServingStatusDetail) SetMessage(v string) {
	o.Message = &v
}

func (o SponsoredProductsCampaignServingStatusDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.HelpUrl) {
		toSerialize["helpUrl"] = o.HelpUrl
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableSponsoredProductsCampaignServingStatusDetail struct {
	value *SponsoredProductsCampaignServingStatusDetail
	isSet bool
}

func (v NullableSponsoredProductsCampaignServingStatusDetail) Get() *SponsoredProductsCampaignServingStatusDetail {
	return v.value
}

func (v *NullableSponsoredProductsCampaignServingStatusDetail) Set(val *SponsoredProductsCampaignServingStatusDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCampaignServingStatusDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCampaignServingStatusDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCampaignServingStatusDetail(val *SponsoredProductsCampaignServingStatusDetail) *NullableSponsoredProductsCampaignServingStatusDetail {
	return &NullableSponsoredProductsCampaignServingStatusDetail{value: val, isSet: true}
}

func (v NullableSponsoredProductsCampaignServingStatusDetail) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCampaignServingStatusDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
