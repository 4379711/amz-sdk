package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGlobalProductAdFailureResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGlobalProductAdFailureResponseItem{}

// SponsoredProductsGlobalProductAdFailureResponseItem struct for SponsoredProductsGlobalProductAdFailureResponseItem
type SponsoredProductsGlobalProductAdFailureResponseItem struct {
	// the index of the product ad in the array from the request body
	Index int32 `json:"index"`
	// A list of validation errors
	Errors []SponsoredProductsProductAdMutationError `json:"errors,omitempty"`
}

type _SponsoredProductsGlobalProductAdFailureResponseItem SponsoredProductsGlobalProductAdFailureResponseItem

// NewSponsoredProductsGlobalProductAdFailureResponseItem instantiates a new SponsoredProductsGlobalProductAdFailureResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGlobalProductAdFailureResponseItem(index int32) *SponsoredProductsGlobalProductAdFailureResponseItem {
	this := SponsoredProductsGlobalProductAdFailureResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsGlobalProductAdFailureResponseItemWithDefaults instantiates a new SponsoredProductsGlobalProductAdFailureResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGlobalProductAdFailureResponseItemWithDefaults() *SponsoredProductsGlobalProductAdFailureResponseItem {
	this := SponsoredProductsGlobalProductAdFailureResponseItem{}
	return &this
}

// GetIndex returns the Index field value
func (o *SponsoredProductsGlobalProductAdFailureResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalProductAdFailureResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsGlobalProductAdFailureResponseItem) SetIndex(v int32) {
	o.Index = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalProductAdFailureResponseItem) GetErrors() []SponsoredProductsProductAdMutationError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsProductAdMutationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalProductAdFailureResponseItem) GetErrorsOk() ([]SponsoredProductsProductAdMutationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalProductAdFailureResponseItem) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsProductAdMutationError and assigns it to the Errors field.
func (o *SponsoredProductsGlobalProductAdFailureResponseItem) SetErrors(v []SponsoredProductsProductAdMutationError) {
	o.Errors = v
}

func (o SponsoredProductsGlobalProductAdFailureResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsGlobalProductAdFailureResponseItem struct {
	value *SponsoredProductsGlobalProductAdFailureResponseItem
	isSet bool
}

func (v NullableSponsoredProductsGlobalProductAdFailureResponseItem) Get() *SponsoredProductsGlobalProductAdFailureResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsGlobalProductAdFailureResponseItem) Set(val *SponsoredProductsGlobalProductAdFailureResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalProductAdFailureResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalProductAdFailureResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalProductAdFailureResponseItem(val *SponsoredProductsGlobalProductAdFailureResponseItem) *NullableSponsoredProductsGlobalProductAdFailureResponseItem {
	return &NullableSponsoredProductsGlobalProductAdFailureResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalProductAdFailureResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalProductAdFailureResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
