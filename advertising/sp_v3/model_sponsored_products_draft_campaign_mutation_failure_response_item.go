package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftCampaignMutationFailureResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftCampaignMutationFailureResponseItem{}

// SponsoredProductsDraftCampaignMutationFailureResponseItem struct for SponsoredProductsDraftCampaignMutationFailureResponseItem
type SponsoredProductsDraftCampaignMutationFailureResponseItem struct {
	// the index of the draft in the array from the request body
	Index int32 `json:"index"`
	// A list of validation errors
	Errors []SponsoredProductsDraftCampaignMutationError `json:"errors,omitempty"`
}

type _SponsoredProductsDraftCampaignMutationFailureResponseItem SponsoredProductsDraftCampaignMutationFailureResponseItem

// NewSponsoredProductsDraftCampaignMutationFailureResponseItem instantiates a new SponsoredProductsDraftCampaignMutationFailureResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftCampaignMutationFailureResponseItem(index int32) *SponsoredProductsDraftCampaignMutationFailureResponseItem {
	this := SponsoredProductsDraftCampaignMutationFailureResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsDraftCampaignMutationFailureResponseItemWithDefaults instantiates a new SponsoredProductsDraftCampaignMutationFailureResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftCampaignMutationFailureResponseItemWithDefaults() *SponsoredProductsDraftCampaignMutationFailureResponseItem {
	this := SponsoredProductsDraftCampaignMutationFailureResponseItem{}
	return &this
}

// GetIndex returns the Index field value
func (o *SponsoredProductsDraftCampaignMutationFailureResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignMutationFailureResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsDraftCampaignMutationFailureResponseItem) SetIndex(v int32) {
	o.Index = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsDraftCampaignMutationFailureResponseItem) GetErrors() []SponsoredProductsDraftCampaignMutationError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsDraftCampaignMutationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignMutationFailureResponseItem) GetErrorsOk() ([]SponsoredProductsDraftCampaignMutationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsDraftCampaignMutationFailureResponseItem) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsDraftCampaignMutationError and assigns it to the Errors field.
func (o *SponsoredProductsDraftCampaignMutationFailureResponseItem) SetErrors(v []SponsoredProductsDraftCampaignMutationError) {
	o.Errors = v
}

func (o SponsoredProductsDraftCampaignMutationFailureResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDraftCampaignMutationFailureResponseItem struct {
	value *SponsoredProductsDraftCampaignMutationFailureResponseItem
	isSet bool
}

func (v NullableSponsoredProductsDraftCampaignMutationFailureResponseItem) Get() *SponsoredProductsDraftCampaignMutationFailureResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsDraftCampaignMutationFailureResponseItem) Set(val *SponsoredProductsDraftCampaignMutationFailureResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftCampaignMutationFailureResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftCampaignMutationFailureResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftCampaignMutationFailureResponseItem(val *SponsoredProductsDraftCampaignMutationFailureResponseItem) *NullableSponsoredProductsDraftCampaignMutationFailureResponseItem {
	return &NullableSponsoredProductsDraftCampaignMutationFailureResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftCampaignMutationFailureResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftCampaignMutationFailureResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
