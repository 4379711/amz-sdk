package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftCampaignNegativeKeywordFailureResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftCampaignNegativeKeywordFailureResponseItem{}

// SponsoredProductsDraftCampaignNegativeKeywordFailureResponseItem struct for SponsoredProductsDraftCampaignNegativeKeywordFailureResponseItem
type SponsoredProductsDraftCampaignNegativeKeywordFailureResponseItem struct {
	// the index of the draft in the array from the request body
	Index int32 `json:"index"`
	// A list of validation errors
	Errors []SponsoredProductsDraftCampaignNegativeKeywordMutationError `json:"errors,omitempty"`
}

type _SponsoredProductsDraftCampaignNegativeKeywordFailureResponseItem SponsoredProductsDraftCampaignNegativeKeywordFailureResponseItem

// NewSponsoredProductsDraftCampaignNegativeKeywordFailureResponseItem instantiates a new SponsoredProductsDraftCampaignNegativeKeywordFailureResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftCampaignNegativeKeywordFailureResponseItem(index int32) *SponsoredProductsDraftCampaignNegativeKeywordFailureResponseItem {
	this := SponsoredProductsDraftCampaignNegativeKeywordFailureResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsDraftCampaignNegativeKeywordFailureResponseItemWithDefaults instantiates a new SponsoredProductsDraftCampaignNegativeKeywordFailureResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftCampaignNegativeKeywordFailureResponseItemWithDefaults() *SponsoredProductsDraftCampaignNegativeKeywordFailureResponseItem {
	this := SponsoredProductsDraftCampaignNegativeKeywordFailureResponseItem{}
	return &this
}

// GetIndex returns the Index field value
func (o *SponsoredProductsDraftCampaignNegativeKeywordFailureResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignNegativeKeywordFailureResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsDraftCampaignNegativeKeywordFailureResponseItem) SetIndex(v int32) {
	o.Index = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsDraftCampaignNegativeKeywordFailureResponseItem) GetErrors() []SponsoredProductsDraftCampaignNegativeKeywordMutationError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsDraftCampaignNegativeKeywordMutationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignNegativeKeywordFailureResponseItem) GetErrorsOk() ([]SponsoredProductsDraftCampaignNegativeKeywordMutationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsDraftCampaignNegativeKeywordFailureResponseItem) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsDraftCampaignNegativeKeywordMutationError and assigns it to the Errors field.
func (o *SponsoredProductsDraftCampaignNegativeKeywordFailureResponseItem) SetErrors(v []SponsoredProductsDraftCampaignNegativeKeywordMutationError) {
	o.Errors = v
}

func (o SponsoredProductsDraftCampaignNegativeKeywordFailureResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDraftCampaignNegativeKeywordFailureResponseItem struct {
	value *SponsoredProductsDraftCampaignNegativeKeywordFailureResponseItem
	isSet bool
}

func (v NullableSponsoredProductsDraftCampaignNegativeKeywordFailureResponseItem) Get() *SponsoredProductsDraftCampaignNegativeKeywordFailureResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsDraftCampaignNegativeKeywordFailureResponseItem) Set(val *SponsoredProductsDraftCampaignNegativeKeywordFailureResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftCampaignNegativeKeywordFailureResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftCampaignNegativeKeywordFailureResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftCampaignNegativeKeywordFailureResponseItem(val *SponsoredProductsDraftCampaignNegativeKeywordFailureResponseItem) *NullableSponsoredProductsDraftCampaignNegativeKeywordFailureResponseItem {
	return &NullableSponsoredProductsDraftCampaignNegativeKeywordFailureResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftCampaignNegativeKeywordFailureResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftCampaignNegativeKeywordFailureResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
