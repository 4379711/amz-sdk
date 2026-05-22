package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsAllTargetsFailureResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsAllTargetsFailureResponseItem{}

// SponsoredProductsAllTargetsFailureResponseItem struct for SponsoredProductsAllTargetsFailureResponseItem
type SponsoredProductsAllTargetsFailureResponseItem struct {
	// the index of the targetingClause in the array from the request body
	Index int32 `json:"index"`
	// A list of validation errors
	Errors []SponsoredProductsBatchResponseError `json:"errors,omitempty"`
}

type _SponsoredProductsAllTargetsFailureResponseItem SponsoredProductsAllTargetsFailureResponseItem

// NewSponsoredProductsAllTargetsFailureResponseItem instantiates a new SponsoredProductsAllTargetsFailureResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsAllTargetsFailureResponseItem(index int32) *SponsoredProductsAllTargetsFailureResponseItem {
	this := SponsoredProductsAllTargetsFailureResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsAllTargetsFailureResponseItemWithDefaults instantiates a new SponsoredProductsAllTargetsFailureResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsAllTargetsFailureResponseItemWithDefaults() *SponsoredProductsAllTargetsFailureResponseItem {
	this := SponsoredProductsAllTargetsFailureResponseItem{}
	return &this
}

// GetIndex returns the Index field value
func (o *SponsoredProductsAllTargetsFailureResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAllTargetsFailureResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsAllTargetsFailureResponseItem) SetIndex(v int32) {
	o.Index = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsAllTargetsFailureResponseItem) GetErrors() []SponsoredProductsBatchResponseError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsBatchResponseError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsAllTargetsFailureResponseItem) GetErrorsOk() ([]SponsoredProductsBatchResponseError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsAllTargetsFailureResponseItem) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsBatchResponseError and assigns it to the Errors field.
func (o *SponsoredProductsAllTargetsFailureResponseItem) SetErrors(v []SponsoredProductsBatchResponseError) {
	o.Errors = v
}

func (o SponsoredProductsAllTargetsFailureResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsAllTargetsFailureResponseItem struct {
	value *SponsoredProductsAllTargetsFailureResponseItem
	isSet bool
}

func (v NullableSponsoredProductsAllTargetsFailureResponseItem) Get() *SponsoredProductsAllTargetsFailureResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsAllTargetsFailureResponseItem) Set(val *SponsoredProductsAllTargetsFailureResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsAllTargetsFailureResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsAllTargetsFailureResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsAllTargetsFailureResponseItem(val *SponsoredProductsAllTargetsFailureResponseItem) *NullableSponsoredProductsAllTargetsFailureResponseItem {
	return &NullableSponsoredProductsAllTargetsFailureResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsAllTargetsFailureResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsAllTargetsFailureResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
