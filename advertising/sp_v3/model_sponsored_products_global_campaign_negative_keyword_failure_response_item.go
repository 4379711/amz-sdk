package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGlobalCampaignNegativeKeywordFailureResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGlobalCampaignNegativeKeywordFailureResponseItem{}

// SponsoredProductsGlobalCampaignNegativeKeywordFailureResponseItem struct for SponsoredProductsGlobalCampaignNegativeKeywordFailureResponseItem
type SponsoredProductsGlobalCampaignNegativeKeywordFailureResponseItem struct {
	// the index of the campaign in the array from the request body
	Index int32 `json:"index"`
	// A list of validation errors
	Errors []SponsoredProductsCampaignNegativeKeywordMutationError `json:"errors,omitempty"`
}

type _SponsoredProductsGlobalCampaignNegativeKeywordFailureResponseItem SponsoredProductsGlobalCampaignNegativeKeywordFailureResponseItem

// NewSponsoredProductsGlobalCampaignNegativeKeywordFailureResponseItem instantiates a new SponsoredProductsGlobalCampaignNegativeKeywordFailureResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGlobalCampaignNegativeKeywordFailureResponseItem(index int32) *SponsoredProductsGlobalCampaignNegativeKeywordFailureResponseItem {
	this := SponsoredProductsGlobalCampaignNegativeKeywordFailureResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsGlobalCampaignNegativeKeywordFailureResponseItemWithDefaults instantiates a new SponsoredProductsGlobalCampaignNegativeKeywordFailureResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGlobalCampaignNegativeKeywordFailureResponseItemWithDefaults() *SponsoredProductsGlobalCampaignNegativeKeywordFailureResponseItem {
	this := SponsoredProductsGlobalCampaignNegativeKeywordFailureResponseItem{}
	return &this
}

// GetIndex returns the Index field value
func (o *SponsoredProductsGlobalCampaignNegativeKeywordFailureResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalCampaignNegativeKeywordFailureResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsGlobalCampaignNegativeKeywordFailureResponseItem) SetIndex(v int32) {
	o.Index = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalCampaignNegativeKeywordFailureResponseItem) GetErrors() []SponsoredProductsCampaignNegativeKeywordMutationError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsCampaignNegativeKeywordMutationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalCampaignNegativeKeywordFailureResponseItem) GetErrorsOk() ([]SponsoredProductsCampaignNegativeKeywordMutationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalCampaignNegativeKeywordFailureResponseItem) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsCampaignNegativeKeywordMutationError and assigns it to the Errors field.
func (o *SponsoredProductsGlobalCampaignNegativeKeywordFailureResponseItem) SetErrors(v []SponsoredProductsCampaignNegativeKeywordMutationError) {
	o.Errors = v
}

func (o SponsoredProductsGlobalCampaignNegativeKeywordFailureResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsGlobalCampaignNegativeKeywordFailureResponseItem struct {
	value *SponsoredProductsGlobalCampaignNegativeKeywordFailureResponseItem
	isSet bool
}

func (v NullableSponsoredProductsGlobalCampaignNegativeKeywordFailureResponseItem) Get() *SponsoredProductsGlobalCampaignNegativeKeywordFailureResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsGlobalCampaignNegativeKeywordFailureResponseItem) Set(val *SponsoredProductsGlobalCampaignNegativeKeywordFailureResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalCampaignNegativeKeywordFailureResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalCampaignNegativeKeywordFailureResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalCampaignNegativeKeywordFailureResponseItem(val *SponsoredProductsGlobalCampaignNegativeKeywordFailureResponseItem) *NullableSponsoredProductsGlobalCampaignNegativeKeywordFailureResponseItem {
	return &NullableSponsoredProductsGlobalCampaignNegativeKeywordFailureResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalCampaignNegativeKeywordFailureResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalCampaignNegativeKeywordFailureResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
