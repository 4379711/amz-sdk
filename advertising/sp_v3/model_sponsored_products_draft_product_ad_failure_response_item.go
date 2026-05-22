package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftProductAdFailureResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftProductAdFailureResponseItem{}

// SponsoredProductsDraftProductAdFailureResponseItem struct for SponsoredProductsDraftProductAdFailureResponseItem
type SponsoredProductsDraftProductAdFailureResponseItem struct {
	// the index of the campaign in the array from the request body
	Index int32 `json:"index"`
	// A list of validation errors
	Errors []SponsoredProductsDraftProductAdMutationError `json:"errors,omitempty"`
}

type _SponsoredProductsDraftProductAdFailureResponseItem SponsoredProductsDraftProductAdFailureResponseItem

// NewSponsoredProductsDraftProductAdFailureResponseItem instantiates a new SponsoredProductsDraftProductAdFailureResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftProductAdFailureResponseItem(index int32) *SponsoredProductsDraftProductAdFailureResponseItem {
	this := SponsoredProductsDraftProductAdFailureResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsDraftProductAdFailureResponseItemWithDefaults instantiates a new SponsoredProductsDraftProductAdFailureResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftProductAdFailureResponseItemWithDefaults() *SponsoredProductsDraftProductAdFailureResponseItem {
	this := SponsoredProductsDraftProductAdFailureResponseItem{}
	return &this
}

// GetIndex returns the Index field value
func (o *SponsoredProductsDraftProductAdFailureResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftProductAdFailureResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsDraftProductAdFailureResponseItem) SetIndex(v int32) {
	o.Index = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsDraftProductAdFailureResponseItem) GetErrors() []SponsoredProductsDraftProductAdMutationError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsDraftProductAdMutationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftProductAdFailureResponseItem) GetErrorsOk() ([]SponsoredProductsDraftProductAdMutationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsDraftProductAdFailureResponseItem) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsDraftProductAdMutationError and assigns it to the Errors field.
func (o *SponsoredProductsDraftProductAdFailureResponseItem) SetErrors(v []SponsoredProductsDraftProductAdMutationError) {
	o.Errors = v
}

func (o SponsoredProductsDraftProductAdFailureResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDraftProductAdFailureResponseItem struct {
	value *SponsoredProductsDraftProductAdFailureResponseItem
	isSet bool
}

func (v NullableSponsoredProductsDraftProductAdFailureResponseItem) Get() *SponsoredProductsDraftProductAdFailureResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsDraftProductAdFailureResponseItem) Set(val *SponsoredProductsDraftProductAdFailureResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftProductAdFailureResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftProductAdFailureResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftProductAdFailureResponseItem(val *SponsoredProductsDraftProductAdFailureResponseItem) *NullableSponsoredProductsDraftProductAdFailureResponseItem {
	return &NullableSponsoredProductsDraftProductAdFailureResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftProductAdFailureResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftProductAdFailureResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
