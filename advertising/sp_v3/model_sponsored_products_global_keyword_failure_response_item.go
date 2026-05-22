package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGlobalKeywordFailureResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGlobalKeywordFailureResponseItem{}

// SponsoredProductsGlobalKeywordFailureResponseItem struct for SponsoredProductsGlobalKeywordFailureResponseItem
type SponsoredProductsGlobalKeywordFailureResponseItem struct {
	// the index of the keyword in the array from the request body
	Index int32 `json:"index"`
	// A list of validation errors
	Errors []SponsoredProductsKeywordMutationError `json:"errors,omitempty"`
}

type _SponsoredProductsGlobalKeywordFailureResponseItem SponsoredProductsGlobalKeywordFailureResponseItem

// NewSponsoredProductsGlobalKeywordFailureResponseItem instantiates a new SponsoredProductsGlobalKeywordFailureResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGlobalKeywordFailureResponseItem(index int32) *SponsoredProductsGlobalKeywordFailureResponseItem {
	this := SponsoredProductsGlobalKeywordFailureResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsGlobalKeywordFailureResponseItemWithDefaults instantiates a new SponsoredProductsGlobalKeywordFailureResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGlobalKeywordFailureResponseItemWithDefaults() *SponsoredProductsGlobalKeywordFailureResponseItem {
	this := SponsoredProductsGlobalKeywordFailureResponseItem{}
	return &this
}

// GetIndex returns the Index field value
func (o *SponsoredProductsGlobalKeywordFailureResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalKeywordFailureResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsGlobalKeywordFailureResponseItem) SetIndex(v int32) {
	o.Index = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalKeywordFailureResponseItem) GetErrors() []SponsoredProductsKeywordMutationError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsKeywordMutationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalKeywordFailureResponseItem) GetErrorsOk() ([]SponsoredProductsKeywordMutationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalKeywordFailureResponseItem) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsKeywordMutationError and assigns it to the Errors field.
func (o *SponsoredProductsGlobalKeywordFailureResponseItem) SetErrors(v []SponsoredProductsKeywordMutationError) {
	o.Errors = v
}

func (o SponsoredProductsGlobalKeywordFailureResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsGlobalKeywordFailureResponseItem struct {
	value *SponsoredProductsGlobalKeywordFailureResponseItem
	isSet bool
}

func (v NullableSponsoredProductsGlobalKeywordFailureResponseItem) Get() *SponsoredProductsGlobalKeywordFailureResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsGlobalKeywordFailureResponseItem) Set(val *SponsoredProductsGlobalKeywordFailureResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalKeywordFailureResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalKeywordFailureResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalKeywordFailureResponseItem(val *SponsoredProductsGlobalKeywordFailureResponseItem) *NullableSponsoredProductsGlobalKeywordFailureResponseItem {
	return &NullableSponsoredProductsGlobalKeywordFailureResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalKeywordFailureResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalKeywordFailureResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
