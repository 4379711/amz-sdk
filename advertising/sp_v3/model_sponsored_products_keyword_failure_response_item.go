package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsKeywordFailureResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsKeywordFailureResponseItem{}

// SponsoredProductsKeywordFailureResponseItem struct for SponsoredProductsKeywordFailureResponseItem
type SponsoredProductsKeywordFailureResponseItem struct {
	// the index of the keyword in the array from the request body
	Index int32 `json:"index"`
	// A list of validation errors
	Errors []SponsoredProductsKeywordMutationError `json:"errors,omitempty"`
}

type _SponsoredProductsKeywordFailureResponseItem SponsoredProductsKeywordFailureResponseItem

// NewSponsoredProductsKeywordFailureResponseItem instantiates a new SponsoredProductsKeywordFailureResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsKeywordFailureResponseItem(index int32) *SponsoredProductsKeywordFailureResponseItem {
	this := SponsoredProductsKeywordFailureResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsKeywordFailureResponseItemWithDefaults instantiates a new SponsoredProductsKeywordFailureResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsKeywordFailureResponseItemWithDefaults() *SponsoredProductsKeywordFailureResponseItem {
	this := SponsoredProductsKeywordFailureResponseItem{}
	return &this
}

// GetIndex returns the Index field value
func (o *SponsoredProductsKeywordFailureResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsKeywordFailureResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsKeywordFailureResponseItem) SetIndex(v int32) {
	o.Index = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsKeywordFailureResponseItem) GetErrors() []SponsoredProductsKeywordMutationError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsKeywordMutationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsKeywordFailureResponseItem) GetErrorsOk() ([]SponsoredProductsKeywordMutationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsKeywordFailureResponseItem) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsKeywordMutationError and assigns it to the Errors field.
func (o *SponsoredProductsKeywordFailureResponseItem) SetErrors(v []SponsoredProductsKeywordMutationError) {
	o.Errors = v
}

func (o SponsoredProductsKeywordFailureResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsKeywordFailureResponseItem struct {
	value *SponsoredProductsKeywordFailureResponseItem
	isSet bool
}

func (v NullableSponsoredProductsKeywordFailureResponseItem) Get() *SponsoredProductsKeywordFailureResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsKeywordFailureResponseItem) Set(val *SponsoredProductsKeywordFailureResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsKeywordFailureResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsKeywordFailureResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsKeywordFailureResponseItem(val *SponsoredProductsKeywordFailureResponseItem) *NullableSponsoredProductsKeywordFailureResponseItem {
	return &NullableSponsoredProductsKeywordFailureResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsKeywordFailureResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsKeywordFailureResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
