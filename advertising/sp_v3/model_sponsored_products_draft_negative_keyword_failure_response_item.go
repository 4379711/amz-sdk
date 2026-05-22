package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftNegativeKeywordFailureResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftNegativeKeywordFailureResponseItem{}

// SponsoredProductsDraftNegativeKeywordFailureResponseItem struct for SponsoredProductsDraftNegativeKeywordFailureResponseItem
type SponsoredProductsDraftNegativeKeywordFailureResponseItem struct {
	// the index of the draftNegativeKeyword in the array from the request body
	Index int32 `json:"index"`
	// A list of validation errors
	Errors []SponsoredProductsDraftNegativeKeywordMutationError `json:"errors,omitempty"`
}

type _SponsoredProductsDraftNegativeKeywordFailureResponseItem SponsoredProductsDraftNegativeKeywordFailureResponseItem

// NewSponsoredProductsDraftNegativeKeywordFailureResponseItem instantiates a new SponsoredProductsDraftNegativeKeywordFailureResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftNegativeKeywordFailureResponseItem(index int32) *SponsoredProductsDraftNegativeKeywordFailureResponseItem {
	this := SponsoredProductsDraftNegativeKeywordFailureResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsDraftNegativeKeywordFailureResponseItemWithDefaults instantiates a new SponsoredProductsDraftNegativeKeywordFailureResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftNegativeKeywordFailureResponseItemWithDefaults() *SponsoredProductsDraftNegativeKeywordFailureResponseItem {
	this := SponsoredProductsDraftNegativeKeywordFailureResponseItem{}
	return &this
}

// GetIndex returns the Index field value
func (o *SponsoredProductsDraftNegativeKeywordFailureResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftNegativeKeywordFailureResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsDraftNegativeKeywordFailureResponseItem) SetIndex(v int32) {
	o.Index = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsDraftNegativeKeywordFailureResponseItem) GetErrors() []SponsoredProductsDraftNegativeKeywordMutationError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsDraftNegativeKeywordMutationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftNegativeKeywordFailureResponseItem) GetErrorsOk() ([]SponsoredProductsDraftNegativeKeywordMutationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsDraftNegativeKeywordFailureResponseItem) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsDraftNegativeKeywordMutationError and assigns it to the Errors field.
func (o *SponsoredProductsDraftNegativeKeywordFailureResponseItem) SetErrors(v []SponsoredProductsDraftNegativeKeywordMutationError) {
	o.Errors = v
}

func (o SponsoredProductsDraftNegativeKeywordFailureResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDraftNegativeKeywordFailureResponseItem struct {
	value *SponsoredProductsDraftNegativeKeywordFailureResponseItem
	isSet bool
}

func (v NullableSponsoredProductsDraftNegativeKeywordFailureResponseItem) Get() *SponsoredProductsDraftNegativeKeywordFailureResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsDraftNegativeKeywordFailureResponseItem) Set(val *SponsoredProductsDraftNegativeKeywordFailureResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftNegativeKeywordFailureResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftNegativeKeywordFailureResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftNegativeKeywordFailureResponseItem(val *SponsoredProductsDraftNegativeKeywordFailureResponseItem) *NullableSponsoredProductsDraftNegativeKeywordFailureResponseItem {
	return &NullableSponsoredProductsDraftNegativeKeywordFailureResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftNegativeKeywordFailureResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftNegativeKeywordFailureResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
