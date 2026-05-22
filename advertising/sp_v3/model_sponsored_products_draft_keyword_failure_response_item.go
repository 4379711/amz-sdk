package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftKeywordFailureResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftKeywordFailureResponseItem{}

// SponsoredProductsDraftKeywordFailureResponseItem struct for SponsoredProductsDraftKeywordFailureResponseItem
type SponsoredProductsDraftKeywordFailureResponseItem struct {
	// the index of the draft keyword in the array from the request body
	Index int32 `json:"index"`
	// A list of validation errors
	Errors []SponsoredProductsDraftKeywordMutationError `json:"errors,omitempty"`
}

type _SponsoredProductsDraftKeywordFailureResponseItem SponsoredProductsDraftKeywordFailureResponseItem

// NewSponsoredProductsDraftKeywordFailureResponseItem instantiates a new SponsoredProductsDraftKeywordFailureResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftKeywordFailureResponseItem(index int32) *SponsoredProductsDraftKeywordFailureResponseItem {
	this := SponsoredProductsDraftKeywordFailureResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsDraftKeywordFailureResponseItemWithDefaults instantiates a new SponsoredProductsDraftKeywordFailureResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftKeywordFailureResponseItemWithDefaults() *SponsoredProductsDraftKeywordFailureResponseItem {
	this := SponsoredProductsDraftKeywordFailureResponseItem{}
	return &this
}

// GetIndex returns the Index field value
func (o *SponsoredProductsDraftKeywordFailureResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftKeywordFailureResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsDraftKeywordFailureResponseItem) SetIndex(v int32) {
	o.Index = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsDraftKeywordFailureResponseItem) GetErrors() []SponsoredProductsDraftKeywordMutationError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsDraftKeywordMutationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftKeywordFailureResponseItem) GetErrorsOk() ([]SponsoredProductsDraftKeywordMutationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsDraftKeywordFailureResponseItem) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsDraftKeywordMutationError and assigns it to the Errors field.
func (o *SponsoredProductsDraftKeywordFailureResponseItem) SetErrors(v []SponsoredProductsDraftKeywordMutationError) {
	o.Errors = v
}

func (o SponsoredProductsDraftKeywordFailureResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDraftKeywordFailureResponseItem struct {
	value *SponsoredProductsDraftKeywordFailureResponseItem
	isSet bool
}

func (v NullableSponsoredProductsDraftKeywordFailureResponseItem) Get() *SponsoredProductsDraftKeywordFailureResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsDraftKeywordFailureResponseItem) Set(val *SponsoredProductsDraftKeywordFailureResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftKeywordFailureResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftKeywordFailureResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftKeywordFailureResponseItem(val *SponsoredProductsDraftKeywordFailureResponseItem) *NullableSponsoredProductsDraftKeywordFailureResponseItem {
	return &NullableSponsoredProductsDraftKeywordFailureResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftKeywordFailureResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftKeywordFailureResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
