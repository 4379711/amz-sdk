package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGlobalNegativeKeywordFailureResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGlobalNegativeKeywordFailureResponseItem{}

// SponsoredProductsGlobalNegativeKeywordFailureResponseItem struct for SponsoredProductsGlobalNegativeKeywordFailureResponseItem
type SponsoredProductsGlobalNegativeKeywordFailureResponseItem struct {
	// the index of the negativeKeyword in the array from the request body
	Index int32 `json:"index"`
	// A list of validation errors
	Errors []SponsoredProductsNegativeKeywordMutationError `json:"errors,omitempty"`
}

type _SponsoredProductsGlobalNegativeKeywordFailureResponseItem SponsoredProductsGlobalNegativeKeywordFailureResponseItem

// NewSponsoredProductsGlobalNegativeKeywordFailureResponseItem instantiates a new SponsoredProductsGlobalNegativeKeywordFailureResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGlobalNegativeKeywordFailureResponseItem(index int32) *SponsoredProductsGlobalNegativeKeywordFailureResponseItem {
	this := SponsoredProductsGlobalNegativeKeywordFailureResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsGlobalNegativeKeywordFailureResponseItemWithDefaults instantiates a new SponsoredProductsGlobalNegativeKeywordFailureResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGlobalNegativeKeywordFailureResponseItemWithDefaults() *SponsoredProductsGlobalNegativeKeywordFailureResponseItem {
	this := SponsoredProductsGlobalNegativeKeywordFailureResponseItem{}
	return &this
}

// GetIndex returns the Index field value
func (o *SponsoredProductsGlobalNegativeKeywordFailureResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalNegativeKeywordFailureResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsGlobalNegativeKeywordFailureResponseItem) SetIndex(v int32) {
	o.Index = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalNegativeKeywordFailureResponseItem) GetErrors() []SponsoredProductsNegativeKeywordMutationError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsNegativeKeywordMutationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalNegativeKeywordFailureResponseItem) GetErrorsOk() ([]SponsoredProductsNegativeKeywordMutationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalNegativeKeywordFailureResponseItem) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsNegativeKeywordMutationError and assigns it to the Errors field.
func (o *SponsoredProductsGlobalNegativeKeywordFailureResponseItem) SetErrors(v []SponsoredProductsNegativeKeywordMutationError) {
	o.Errors = v
}

func (o SponsoredProductsGlobalNegativeKeywordFailureResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsGlobalNegativeKeywordFailureResponseItem struct {
	value *SponsoredProductsGlobalNegativeKeywordFailureResponseItem
	isSet bool
}

func (v NullableSponsoredProductsGlobalNegativeKeywordFailureResponseItem) Get() *SponsoredProductsGlobalNegativeKeywordFailureResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsGlobalNegativeKeywordFailureResponseItem) Set(val *SponsoredProductsGlobalNegativeKeywordFailureResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalNegativeKeywordFailureResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalNegativeKeywordFailureResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalNegativeKeywordFailureResponseItem(val *SponsoredProductsGlobalNegativeKeywordFailureResponseItem) *NullableSponsoredProductsGlobalNegativeKeywordFailureResponseItem {
	return &NullableSponsoredProductsGlobalNegativeKeywordFailureResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalNegativeKeywordFailureResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalNegativeKeywordFailureResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
