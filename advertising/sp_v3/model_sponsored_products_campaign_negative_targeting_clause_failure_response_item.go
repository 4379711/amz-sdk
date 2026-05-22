package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCampaignNegativeTargetingClauseFailureResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCampaignNegativeTargetingClauseFailureResponseItem{}

// SponsoredProductsCampaignNegativeTargetingClauseFailureResponseItem struct for SponsoredProductsCampaignNegativeTargetingClauseFailureResponseItem
type SponsoredProductsCampaignNegativeTargetingClauseFailureResponseItem struct {
	// the index of the CampaignNegativeTargets in the array from the request body
	Index int32 `json:"index"`
	// A list of validation errors
	Errors []SponsoredProductsCampaignNegativeTargetsMutationError `json:"errors,omitempty"`
}

type _SponsoredProductsCampaignNegativeTargetingClauseFailureResponseItem SponsoredProductsCampaignNegativeTargetingClauseFailureResponseItem

// NewSponsoredProductsCampaignNegativeTargetingClauseFailureResponseItem instantiates a new SponsoredProductsCampaignNegativeTargetingClauseFailureResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCampaignNegativeTargetingClauseFailureResponseItem(index int32) *SponsoredProductsCampaignNegativeTargetingClauseFailureResponseItem {
	this := SponsoredProductsCampaignNegativeTargetingClauseFailureResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsCampaignNegativeTargetingClauseFailureResponseItemWithDefaults instantiates a new SponsoredProductsCampaignNegativeTargetingClauseFailureResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCampaignNegativeTargetingClauseFailureResponseItemWithDefaults() *SponsoredProductsCampaignNegativeTargetingClauseFailureResponseItem {
	this := SponsoredProductsCampaignNegativeTargetingClauseFailureResponseItem{}
	return &this
}

// GetIndex returns the Index field value
func (o *SponsoredProductsCampaignNegativeTargetingClauseFailureResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignNegativeTargetingClauseFailureResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsCampaignNegativeTargetingClauseFailureResponseItem) SetIndex(v int32) {
	o.Index = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsCampaignNegativeTargetingClauseFailureResponseItem) GetErrors() []SponsoredProductsCampaignNegativeTargetsMutationError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsCampaignNegativeTargetsMutationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCampaignNegativeTargetingClauseFailureResponseItem) GetErrorsOk() ([]SponsoredProductsCampaignNegativeTargetsMutationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsCampaignNegativeTargetingClauseFailureResponseItem) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsCampaignNegativeTargetsMutationError and assigns it to the Errors field.
func (o *SponsoredProductsCampaignNegativeTargetingClauseFailureResponseItem) SetErrors(v []SponsoredProductsCampaignNegativeTargetsMutationError) {
	o.Errors = v
}

func (o SponsoredProductsCampaignNegativeTargetingClauseFailureResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsCampaignNegativeTargetingClauseFailureResponseItem struct {
	value *SponsoredProductsCampaignNegativeTargetingClauseFailureResponseItem
	isSet bool
}

func (v NullableSponsoredProductsCampaignNegativeTargetingClauseFailureResponseItem) Get() *SponsoredProductsCampaignNegativeTargetingClauseFailureResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsCampaignNegativeTargetingClauseFailureResponseItem) Set(val *SponsoredProductsCampaignNegativeTargetingClauseFailureResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCampaignNegativeTargetingClauseFailureResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCampaignNegativeTargetingClauseFailureResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCampaignNegativeTargetingClauseFailureResponseItem(val *SponsoredProductsCampaignNegativeTargetingClauseFailureResponseItem) *NullableSponsoredProductsCampaignNegativeTargetingClauseFailureResponseItem {
	return &NullableSponsoredProductsCampaignNegativeTargetingClauseFailureResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsCampaignNegativeTargetingClauseFailureResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCampaignNegativeTargetingClauseFailureResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
