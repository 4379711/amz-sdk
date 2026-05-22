package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGlobalCampaignMutationFailureResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGlobalCampaignMutationFailureResponseItem{}

// SponsoredProductsGlobalCampaignMutationFailureResponseItem struct for SponsoredProductsGlobalCampaignMutationFailureResponseItem
type SponsoredProductsGlobalCampaignMutationFailureResponseItem struct {
	// the index of the campaign in the array from the request body
	Index int32 `json:"index"`
	// A list of validation errors
	Errors []SponsoredProductsCampaignMutationError `json:"errors,omitempty"`
}

type _SponsoredProductsGlobalCampaignMutationFailureResponseItem SponsoredProductsGlobalCampaignMutationFailureResponseItem

// NewSponsoredProductsGlobalCampaignMutationFailureResponseItem instantiates a new SponsoredProductsGlobalCampaignMutationFailureResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGlobalCampaignMutationFailureResponseItem(index int32) *SponsoredProductsGlobalCampaignMutationFailureResponseItem {
	this := SponsoredProductsGlobalCampaignMutationFailureResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsGlobalCampaignMutationFailureResponseItemWithDefaults instantiates a new SponsoredProductsGlobalCampaignMutationFailureResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGlobalCampaignMutationFailureResponseItemWithDefaults() *SponsoredProductsGlobalCampaignMutationFailureResponseItem {
	this := SponsoredProductsGlobalCampaignMutationFailureResponseItem{}
	return &this
}

// GetIndex returns the Index field value
func (o *SponsoredProductsGlobalCampaignMutationFailureResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalCampaignMutationFailureResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsGlobalCampaignMutationFailureResponseItem) SetIndex(v int32) {
	o.Index = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalCampaignMutationFailureResponseItem) GetErrors() []SponsoredProductsCampaignMutationError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsCampaignMutationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalCampaignMutationFailureResponseItem) GetErrorsOk() ([]SponsoredProductsCampaignMutationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalCampaignMutationFailureResponseItem) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsCampaignMutationError and assigns it to the Errors field.
func (o *SponsoredProductsGlobalCampaignMutationFailureResponseItem) SetErrors(v []SponsoredProductsCampaignMutationError) {
	o.Errors = v
}

func (o SponsoredProductsGlobalCampaignMutationFailureResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsGlobalCampaignMutationFailureResponseItem struct {
	value *SponsoredProductsGlobalCampaignMutationFailureResponseItem
	isSet bool
}

func (v NullableSponsoredProductsGlobalCampaignMutationFailureResponseItem) Get() *SponsoredProductsGlobalCampaignMutationFailureResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsGlobalCampaignMutationFailureResponseItem) Set(val *SponsoredProductsGlobalCampaignMutationFailureResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalCampaignMutationFailureResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalCampaignMutationFailureResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalCampaignMutationFailureResponseItem(val *SponsoredProductsGlobalCampaignMutationFailureResponseItem) *NullableSponsoredProductsGlobalCampaignMutationFailureResponseItem {
	return &NullableSponsoredProductsGlobalCampaignMutationFailureResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalCampaignMutationFailureResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalCampaignMutationFailureResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
