package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftAdGroupFailureResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftAdGroupFailureResponseItem{}

// SponsoredProductsDraftAdGroupFailureResponseItem struct for SponsoredProductsDraftAdGroupFailureResponseItem
type SponsoredProductsDraftAdGroupFailureResponseItem struct {
	// the index of the adGroup in the array from the request body
	Index int32 `json:"index"`
	// A list of validation errors
	Errors []SponsoredProductsDraftAdGroupMutationError `json:"errors,omitempty"`
}

type _SponsoredProductsDraftAdGroupFailureResponseItem SponsoredProductsDraftAdGroupFailureResponseItem

// NewSponsoredProductsDraftAdGroupFailureResponseItem instantiates a new SponsoredProductsDraftAdGroupFailureResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftAdGroupFailureResponseItem(index int32) *SponsoredProductsDraftAdGroupFailureResponseItem {
	this := SponsoredProductsDraftAdGroupFailureResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsDraftAdGroupFailureResponseItemWithDefaults instantiates a new SponsoredProductsDraftAdGroupFailureResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftAdGroupFailureResponseItemWithDefaults() *SponsoredProductsDraftAdGroupFailureResponseItem {
	this := SponsoredProductsDraftAdGroupFailureResponseItem{}
	return &this
}

// GetIndex returns the Index field value
func (o *SponsoredProductsDraftAdGroupFailureResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftAdGroupFailureResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsDraftAdGroupFailureResponseItem) SetIndex(v int32) {
	o.Index = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsDraftAdGroupFailureResponseItem) GetErrors() []SponsoredProductsDraftAdGroupMutationError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsDraftAdGroupMutationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftAdGroupFailureResponseItem) GetErrorsOk() ([]SponsoredProductsDraftAdGroupMutationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsDraftAdGroupFailureResponseItem) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsDraftAdGroupMutationError and assigns it to the Errors field.
func (o *SponsoredProductsDraftAdGroupFailureResponseItem) SetErrors(v []SponsoredProductsDraftAdGroupMutationError) {
	o.Errors = v
}

func (o SponsoredProductsDraftAdGroupFailureResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDraftAdGroupFailureResponseItem struct {
	value *SponsoredProductsDraftAdGroupFailureResponseItem
	isSet bool
}

func (v NullableSponsoredProductsDraftAdGroupFailureResponseItem) Get() *SponsoredProductsDraftAdGroupFailureResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsDraftAdGroupFailureResponseItem) Set(val *SponsoredProductsDraftAdGroupFailureResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftAdGroupFailureResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftAdGroupFailureResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftAdGroupFailureResponseItem(val *SponsoredProductsDraftAdGroupFailureResponseItem) *NullableSponsoredProductsDraftAdGroupFailureResponseItem {
	return &NullableSponsoredProductsDraftAdGroupFailureResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftAdGroupFailureResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftAdGroupFailureResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
