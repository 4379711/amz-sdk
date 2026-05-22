package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsProductAdFailureResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsProductAdFailureResponseItem{}

// SponsoredProductsProductAdFailureResponseItem struct for SponsoredProductsProductAdFailureResponseItem
type SponsoredProductsProductAdFailureResponseItem struct {
	// the index of the product ad in the array from the request body
	Index int32 `json:"index"`
	// A list of validation errors
	Errors []SponsoredProductsProductAdMutationError `json:"errors,omitempty"`
}

type _SponsoredProductsProductAdFailureResponseItem SponsoredProductsProductAdFailureResponseItem

// NewSponsoredProductsProductAdFailureResponseItem instantiates a new SponsoredProductsProductAdFailureResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsProductAdFailureResponseItem(index int32) *SponsoredProductsProductAdFailureResponseItem {
	this := SponsoredProductsProductAdFailureResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsProductAdFailureResponseItemWithDefaults instantiates a new SponsoredProductsProductAdFailureResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsProductAdFailureResponseItemWithDefaults() *SponsoredProductsProductAdFailureResponseItem {
	this := SponsoredProductsProductAdFailureResponseItem{}
	return &this
}

// GetIndex returns the Index field value
func (o *SponsoredProductsProductAdFailureResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsProductAdFailureResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsProductAdFailureResponseItem) SetIndex(v int32) {
	o.Index = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsProductAdFailureResponseItem) GetErrors() []SponsoredProductsProductAdMutationError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsProductAdMutationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsProductAdFailureResponseItem) GetErrorsOk() ([]SponsoredProductsProductAdMutationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsProductAdFailureResponseItem) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsProductAdMutationError and assigns it to the Errors field.
func (o *SponsoredProductsProductAdFailureResponseItem) SetErrors(v []SponsoredProductsProductAdMutationError) {
	o.Errors = v
}

func (o SponsoredProductsProductAdFailureResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsProductAdFailureResponseItem struct {
	value *SponsoredProductsProductAdFailureResponseItem
	isSet bool
}

func (v NullableSponsoredProductsProductAdFailureResponseItem) Get() *SponsoredProductsProductAdFailureResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsProductAdFailureResponseItem) Set(val *SponsoredProductsProductAdFailureResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsProductAdFailureResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsProductAdFailureResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsProductAdFailureResponseItem(val *SponsoredProductsProductAdFailureResponseItem) *NullableSponsoredProductsProductAdFailureResponseItem {
	return &NullableSponsoredProductsProductAdFailureResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsProductAdFailureResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsProductAdFailureResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
