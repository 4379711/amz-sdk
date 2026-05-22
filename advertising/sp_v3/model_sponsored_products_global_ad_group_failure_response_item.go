package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGlobalAdGroupFailureResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGlobalAdGroupFailureResponseItem{}

// SponsoredProductsGlobalAdGroupFailureResponseItem struct for SponsoredProductsGlobalAdGroupFailureResponseItem
type SponsoredProductsGlobalAdGroupFailureResponseItem struct {
	// the index of the adGroup in the array from the request body
	Index int32 `json:"index"`
	// A list of validation errors
	Errors []SponsoredProductsAdGroupMutationError `json:"errors,omitempty"`
}

type _SponsoredProductsGlobalAdGroupFailureResponseItem SponsoredProductsGlobalAdGroupFailureResponseItem

// NewSponsoredProductsGlobalAdGroupFailureResponseItem instantiates a new SponsoredProductsGlobalAdGroupFailureResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGlobalAdGroupFailureResponseItem(index int32) *SponsoredProductsGlobalAdGroupFailureResponseItem {
	this := SponsoredProductsGlobalAdGroupFailureResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsGlobalAdGroupFailureResponseItemWithDefaults instantiates a new SponsoredProductsGlobalAdGroupFailureResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGlobalAdGroupFailureResponseItemWithDefaults() *SponsoredProductsGlobalAdGroupFailureResponseItem {
	this := SponsoredProductsGlobalAdGroupFailureResponseItem{}
	return &this
}

// GetIndex returns the Index field value
func (o *SponsoredProductsGlobalAdGroupFailureResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalAdGroupFailureResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsGlobalAdGroupFailureResponseItem) SetIndex(v int32) {
	o.Index = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalAdGroupFailureResponseItem) GetErrors() []SponsoredProductsAdGroupMutationError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsAdGroupMutationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalAdGroupFailureResponseItem) GetErrorsOk() ([]SponsoredProductsAdGroupMutationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalAdGroupFailureResponseItem) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsAdGroupMutationError and assigns it to the Errors field.
func (o *SponsoredProductsGlobalAdGroupFailureResponseItem) SetErrors(v []SponsoredProductsAdGroupMutationError) {
	o.Errors = v
}

func (o SponsoredProductsGlobalAdGroupFailureResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsGlobalAdGroupFailureResponseItem struct {
	value *SponsoredProductsGlobalAdGroupFailureResponseItem
	isSet bool
}

func (v NullableSponsoredProductsGlobalAdGroupFailureResponseItem) Get() *SponsoredProductsGlobalAdGroupFailureResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsGlobalAdGroupFailureResponseItem) Set(val *SponsoredProductsGlobalAdGroupFailureResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalAdGroupFailureResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalAdGroupFailureResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalAdGroupFailureResponseItem(val *SponsoredProductsGlobalAdGroupFailureResponseItem) *NullableSponsoredProductsGlobalAdGroupFailureResponseItem {
	return &NullableSponsoredProductsGlobalAdGroupFailureResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalAdGroupFailureResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalAdGroupFailureResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
