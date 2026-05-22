package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsAdGroupServingStatusDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsAdGroupServingStatusDetail{}

// SponsoredProductsAdGroupServingStatusDetail struct for SponsoredProductsAdGroupServingStatusDetail
type SponsoredProductsAdGroupServingStatusDetail struct {
	Name *SponsoredProductsAdGroupServingStatusReason `json:"name,omitempty"`
	// A URL with additional information about the status identifier.
	HelpUrl *string `json:"helpUrl,omitempty"`
	// A human-readable description of the status identifier specified in the name field.
	Message *string `json:"message,omitempty"`
}

// NewSponsoredProductsAdGroupServingStatusDetail instantiates a new SponsoredProductsAdGroupServingStatusDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsAdGroupServingStatusDetail() *SponsoredProductsAdGroupServingStatusDetail {
	this := SponsoredProductsAdGroupServingStatusDetail{}
	return &this
}

// NewSponsoredProductsAdGroupServingStatusDetailWithDefaults instantiates a new SponsoredProductsAdGroupServingStatusDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsAdGroupServingStatusDetailWithDefaults() *SponsoredProductsAdGroupServingStatusDetail {
	this := SponsoredProductsAdGroupServingStatusDetail{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SponsoredProductsAdGroupServingStatusDetail) GetName() SponsoredProductsAdGroupServingStatusReason {
	if o == nil || IsNil(o.Name) {
		var ret SponsoredProductsAdGroupServingStatusReason
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdGroupServingStatusDetail) GetNameOk() (*SponsoredProductsAdGroupServingStatusReason, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SponsoredProductsAdGroupServingStatusDetail) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given SponsoredProductsAdGroupServingStatusReason and assigns it to the Name field.
func (o *SponsoredProductsAdGroupServingStatusDetail) SetName(v SponsoredProductsAdGroupServingStatusReason) {
	o.Name = &v
}

// GetHelpUrl returns the HelpUrl field value if set, zero value otherwise.
func (o *SponsoredProductsAdGroupServingStatusDetail) GetHelpUrl() string {
	if o == nil || IsNil(o.HelpUrl) {
		var ret string
		return ret
	}
	return *o.HelpUrl
}

// GetHelpUrlOk returns a tuple with the HelpUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdGroupServingStatusDetail) GetHelpUrlOk() (*string, bool) {
	if o == nil || IsNil(o.HelpUrl) {
		return nil, false
	}
	return o.HelpUrl, true
}

// HasHelpUrl returns a boolean if a field has been set.
func (o *SponsoredProductsAdGroupServingStatusDetail) HasHelpUrl() bool {
	if o != nil && !IsNil(o.HelpUrl) {
		return true
	}

	return false
}

// SetHelpUrl gets a reference to the given string and assigns it to the HelpUrl field.
func (o *SponsoredProductsAdGroupServingStatusDetail) SetHelpUrl(v string) {
	o.HelpUrl = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SponsoredProductsAdGroupServingStatusDetail) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAdGroupServingStatusDetail) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SponsoredProductsAdGroupServingStatusDetail) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *SponsoredProductsAdGroupServingStatusDetail) SetMessage(v string) {
	o.Message = &v
}

func (o SponsoredProductsAdGroupServingStatusDetail) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsAdGroupServingStatusDetail struct {
	value *SponsoredProductsAdGroupServingStatusDetail
	isSet bool
}

func (v NullableSponsoredProductsAdGroupServingStatusDetail) Get() *SponsoredProductsAdGroupServingStatusDetail {
	return v.value
}

func (v *NullableSponsoredProductsAdGroupServingStatusDetail) Set(val *SponsoredProductsAdGroupServingStatusDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsAdGroupServingStatusDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsAdGroupServingStatusDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsAdGroupServingStatusDetail(val *SponsoredProductsAdGroupServingStatusDetail) *NullableSponsoredProductsAdGroupServingStatusDetail {
	return &NullableSponsoredProductsAdGroupServingStatusDetail{value: val, isSet: true}
}

func (v NullableSponsoredProductsAdGroupServingStatusDetail) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsAdGroupServingStatusDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
