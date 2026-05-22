package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftTargetingClauseFailureResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftTargetingClauseFailureResponseItem{}

// SponsoredProductsDraftTargetingClauseFailureResponseItem struct for SponsoredProductsDraftTargetingClauseFailureResponseItem
type SponsoredProductsDraftTargetingClauseFailureResponseItem struct {
	// the index of the draftTargetingClause in the array from the request body
	Index int32 `json:"index"`
	// A list of validation errors
	Errors []SponsoredProductsDraftTargetMutationError `json:"errors,omitempty"`
}

type _SponsoredProductsDraftTargetingClauseFailureResponseItem SponsoredProductsDraftTargetingClauseFailureResponseItem

// NewSponsoredProductsDraftTargetingClauseFailureResponseItem instantiates a new SponsoredProductsDraftTargetingClauseFailureResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftTargetingClauseFailureResponseItem(index int32) *SponsoredProductsDraftTargetingClauseFailureResponseItem {
	this := SponsoredProductsDraftTargetingClauseFailureResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsDraftTargetingClauseFailureResponseItemWithDefaults instantiates a new SponsoredProductsDraftTargetingClauseFailureResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftTargetingClauseFailureResponseItemWithDefaults() *SponsoredProductsDraftTargetingClauseFailureResponseItem {
	this := SponsoredProductsDraftTargetingClauseFailureResponseItem{}
	return &this
}

// GetIndex returns the Index field value
func (o *SponsoredProductsDraftTargetingClauseFailureResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftTargetingClauseFailureResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsDraftTargetingClauseFailureResponseItem) SetIndex(v int32) {
	o.Index = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsDraftTargetingClauseFailureResponseItem) GetErrors() []SponsoredProductsDraftTargetMutationError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsDraftTargetMutationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftTargetingClauseFailureResponseItem) GetErrorsOk() ([]SponsoredProductsDraftTargetMutationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsDraftTargetingClauseFailureResponseItem) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsDraftTargetMutationError and assigns it to the Errors field.
func (o *SponsoredProductsDraftTargetingClauseFailureResponseItem) SetErrors(v []SponsoredProductsDraftTargetMutationError) {
	o.Errors = v
}

func (o SponsoredProductsDraftTargetingClauseFailureResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDraftTargetingClauseFailureResponseItem struct {
	value *SponsoredProductsDraftTargetingClauseFailureResponseItem
	isSet bool
}

func (v NullableSponsoredProductsDraftTargetingClauseFailureResponseItem) Get() *SponsoredProductsDraftTargetingClauseFailureResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsDraftTargetingClauseFailureResponseItem) Set(val *SponsoredProductsDraftTargetingClauseFailureResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftTargetingClauseFailureResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftTargetingClauseFailureResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftTargetingClauseFailureResponseItem(val *SponsoredProductsDraftTargetingClauseFailureResponseItem) *NullableSponsoredProductsDraftTargetingClauseFailureResponseItem {
	return &NullableSponsoredProductsDraftTargetingClauseFailureResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftTargetingClauseFailureResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftTargetingClauseFailureResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
