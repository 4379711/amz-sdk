package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsKeywordServingStatusDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsKeywordServingStatusDetail{}

// SponsoredProductsKeywordServingStatusDetail struct for SponsoredProductsKeywordServingStatusDetail
type SponsoredProductsKeywordServingStatusDetail struct {
	Name *SponsoredProductsKeywordServingStatusReason `json:"name,omitempty"`
	// A URL with additional information about the status identifier.
	HelpUrl *string `json:"helpUrl,omitempty"`
	// A human-readable description of the status identifier specified in the name field.
	Message *string `json:"message,omitempty"`
}

// NewSponsoredProductsKeywordServingStatusDetail instantiates a new SponsoredProductsKeywordServingStatusDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsKeywordServingStatusDetail() *SponsoredProductsKeywordServingStatusDetail {
	this := SponsoredProductsKeywordServingStatusDetail{}
	return &this
}

// NewSponsoredProductsKeywordServingStatusDetailWithDefaults instantiates a new SponsoredProductsKeywordServingStatusDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsKeywordServingStatusDetailWithDefaults() *SponsoredProductsKeywordServingStatusDetail {
	this := SponsoredProductsKeywordServingStatusDetail{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SponsoredProductsKeywordServingStatusDetail) GetName() SponsoredProductsKeywordServingStatusReason {
	if o == nil || IsNil(o.Name) {
		var ret SponsoredProductsKeywordServingStatusReason
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsKeywordServingStatusDetail) GetNameOk() (*SponsoredProductsKeywordServingStatusReason, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SponsoredProductsKeywordServingStatusDetail) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given SponsoredProductsKeywordServingStatusReason and assigns it to the Name field.
func (o *SponsoredProductsKeywordServingStatusDetail) SetName(v SponsoredProductsKeywordServingStatusReason) {
	o.Name = &v
}

// GetHelpUrl returns the HelpUrl field value if set, zero value otherwise.
func (o *SponsoredProductsKeywordServingStatusDetail) GetHelpUrl() string {
	if o == nil || IsNil(o.HelpUrl) {
		var ret string
		return ret
	}
	return *o.HelpUrl
}

// GetHelpUrlOk returns a tuple with the HelpUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsKeywordServingStatusDetail) GetHelpUrlOk() (*string, bool) {
	if o == nil || IsNil(o.HelpUrl) {
		return nil, false
	}
	return o.HelpUrl, true
}

// HasHelpUrl returns a boolean if a field has been set.
func (o *SponsoredProductsKeywordServingStatusDetail) HasHelpUrl() bool {
	if o != nil && !IsNil(o.HelpUrl) {
		return true
	}

	return false
}

// SetHelpUrl gets a reference to the given string and assigns it to the HelpUrl field.
func (o *SponsoredProductsKeywordServingStatusDetail) SetHelpUrl(v string) {
	o.HelpUrl = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SponsoredProductsKeywordServingStatusDetail) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsKeywordServingStatusDetail) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SponsoredProductsKeywordServingStatusDetail) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *SponsoredProductsKeywordServingStatusDetail) SetMessage(v string) {
	o.Message = &v
}

func (o SponsoredProductsKeywordServingStatusDetail) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsKeywordServingStatusDetail struct {
	value *SponsoredProductsKeywordServingStatusDetail
	isSet bool
}

func (v NullableSponsoredProductsKeywordServingStatusDetail) Get() *SponsoredProductsKeywordServingStatusDetail {
	return v.value
}

func (v *NullableSponsoredProductsKeywordServingStatusDetail) Set(val *SponsoredProductsKeywordServingStatusDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsKeywordServingStatusDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsKeywordServingStatusDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsKeywordServingStatusDetail(val *SponsoredProductsKeywordServingStatusDetail) *NullableSponsoredProductsKeywordServingStatusDetail {
	return &NullableSponsoredProductsKeywordServingStatusDetail{value: val, isSet: true}
}

func (v NullableSponsoredProductsKeywordServingStatusDetail) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsKeywordServingStatusDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
