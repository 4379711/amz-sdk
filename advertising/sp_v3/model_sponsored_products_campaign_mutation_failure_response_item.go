package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCampaignMutationFailureResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCampaignMutationFailureResponseItem{}

// SponsoredProductsCampaignMutationFailureResponseItem struct for SponsoredProductsCampaignMutationFailureResponseItem
type SponsoredProductsCampaignMutationFailureResponseItem struct {
	// the index of the campaign in the array from the request body
	Index int32 `json:"index"`
	// A list of validation errors
	Errors []SponsoredProductsCampaignMutationError `json:"errors,omitempty"`
}

type _SponsoredProductsCampaignMutationFailureResponseItem SponsoredProductsCampaignMutationFailureResponseItem

// NewSponsoredProductsCampaignMutationFailureResponseItem instantiates a new SponsoredProductsCampaignMutationFailureResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCampaignMutationFailureResponseItem(index int32) *SponsoredProductsCampaignMutationFailureResponseItem {
	this := SponsoredProductsCampaignMutationFailureResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsCampaignMutationFailureResponseItemWithDefaults instantiates a new SponsoredProductsCampaignMutationFailureResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCampaignMutationFailureResponseItemWithDefaults() *SponsoredProductsCampaignMutationFailureResponseItem {
	this := SponsoredProductsCampaignMutationFailureResponseItem{}
	return &this
}

// GetIndex returns the Index field value
func (o *SponsoredProductsCampaignMutationFailureResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignMutationFailureResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsCampaignMutationFailureResponseItem) SetIndex(v int32) {
	o.Index = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsCampaignMutationFailureResponseItem) GetErrors() []SponsoredProductsCampaignMutationError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsCampaignMutationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignMutationFailureResponseItem) GetErrorsOk() ([]SponsoredProductsCampaignMutationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsCampaignMutationFailureResponseItem) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsCampaignMutationError and assigns it to the Errors field.
func (o *SponsoredProductsCampaignMutationFailureResponseItem) SetErrors(v []SponsoredProductsCampaignMutationError) {
	o.Errors = v
}

func (o SponsoredProductsCampaignMutationFailureResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsCampaignMutationFailureResponseItem struct {
	value *SponsoredProductsCampaignMutationFailureResponseItem
	isSet bool
}

func (v NullableSponsoredProductsCampaignMutationFailureResponseItem) Get() *SponsoredProductsCampaignMutationFailureResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsCampaignMutationFailureResponseItem) Set(val *SponsoredProductsCampaignMutationFailureResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCampaignMutationFailureResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCampaignMutationFailureResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCampaignMutationFailureResponseItem(val *SponsoredProductsCampaignMutationFailureResponseItem) *NullableSponsoredProductsCampaignMutationFailureResponseItem {
	return &NullableSponsoredProductsCampaignMutationFailureResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsCampaignMutationFailureResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCampaignMutationFailureResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
