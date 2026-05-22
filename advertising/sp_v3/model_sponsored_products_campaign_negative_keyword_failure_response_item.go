package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCampaignNegativeKeywordFailureResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCampaignNegativeKeywordFailureResponseItem{}

// SponsoredProductsCampaignNegativeKeywordFailureResponseItem struct for SponsoredProductsCampaignNegativeKeywordFailureResponseItem
type SponsoredProductsCampaignNegativeKeywordFailureResponseItem struct {
	// the index of the campaign in the array from the request body
	Index int32 `json:"index"`
	// A list of validation errors
	Errors []SponsoredProductsCampaignNegativeKeywordMutationError `json:"errors,omitempty"`
}

type _SponsoredProductsCampaignNegativeKeywordFailureResponseItem SponsoredProductsCampaignNegativeKeywordFailureResponseItem

// NewSponsoredProductsCampaignNegativeKeywordFailureResponseItem instantiates a new SponsoredProductsCampaignNegativeKeywordFailureResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCampaignNegativeKeywordFailureResponseItem(index int32) *SponsoredProductsCampaignNegativeKeywordFailureResponseItem {
	this := SponsoredProductsCampaignNegativeKeywordFailureResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsCampaignNegativeKeywordFailureResponseItemWithDefaults instantiates a new SponsoredProductsCampaignNegativeKeywordFailureResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCampaignNegativeKeywordFailureResponseItemWithDefaults() *SponsoredProductsCampaignNegativeKeywordFailureResponseItem {
	this := SponsoredProductsCampaignNegativeKeywordFailureResponseItem{}
	return &this
}

// GetIndex returns the Index field value
func (o *SponsoredProductsCampaignNegativeKeywordFailureResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignNegativeKeywordFailureResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsCampaignNegativeKeywordFailureResponseItem) SetIndex(v int32) {
	o.Index = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsCampaignNegativeKeywordFailureResponseItem) GetErrors() []SponsoredProductsCampaignNegativeKeywordMutationError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsCampaignNegativeKeywordMutationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignNegativeKeywordFailureResponseItem) GetErrorsOk() ([]SponsoredProductsCampaignNegativeKeywordMutationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsCampaignNegativeKeywordFailureResponseItem) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsCampaignNegativeKeywordMutationError and assigns it to the Errors field.
func (o *SponsoredProductsCampaignNegativeKeywordFailureResponseItem) SetErrors(v []SponsoredProductsCampaignNegativeKeywordMutationError) {
	o.Errors = v
}

func (o SponsoredProductsCampaignNegativeKeywordFailureResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsCampaignNegativeKeywordFailureResponseItem struct {
	value *SponsoredProductsCampaignNegativeKeywordFailureResponseItem
	isSet bool
}

func (v NullableSponsoredProductsCampaignNegativeKeywordFailureResponseItem) Get() *SponsoredProductsCampaignNegativeKeywordFailureResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsCampaignNegativeKeywordFailureResponseItem) Set(val *SponsoredProductsCampaignNegativeKeywordFailureResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCampaignNegativeKeywordFailureResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCampaignNegativeKeywordFailureResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCampaignNegativeKeywordFailureResponseItem(val *SponsoredProductsCampaignNegativeKeywordFailureResponseItem) *NullableSponsoredProductsCampaignNegativeKeywordFailureResponseItem {
	return &NullableSponsoredProductsCampaignNegativeKeywordFailureResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsCampaignNegativeKeywordFailureResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCampaignNegativeKeywordFailureResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
