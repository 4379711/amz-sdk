package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsTargetingClauseFailureResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsTargetingClauseFailureResponseItem{}

// SponsoredProductsTargetingClauseFailureResponseItem struct for SponsoredProductsTargetingClauseFailureResponseItem
type SponsoredProductsTargetingClauseFailureResponseItem struct {
	// the index of the targetingClause in the array from the request body
	Index int32 `json:"index"`
	// A list of validation errors
	Errors []SponsoredProductsTargetMutationError `json:"errors,omitempty"`
}

type _SponsoredProductsTargetingClauseFailureResponseItem SponsoredProductsTargetingClauseFailureResponseItem

// NewSponsoredProductsTargetingClauseFailureResponseItem instantiates a new SponsoredProductsTargetingClauseFailureResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsTargetingClauseFailureResponseItem(index int32) *SponsoredProductsTargetingClauseFailureResponseItem {
	this := SponsoredProductsTargetingClauseFailureResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsTargetingClauseFailureResponseItemWithDefaults instantiates a new SponsoredProductsTargetingClauseFailureResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsTargetingClauseFailureResponseItemWithDefaults() *SponsoredProductsTargetingClauseFailureResponseItem {
	this := SponsoredProductsTargetingClauseFailureResponseItem{}
	return &this
}

// GetIndex returns the Index field value
func (o *SponsoredProductsTargetingClauseFailureResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetingClauseFailureResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsTargetingClauseFailureResponseItem) SetIndex(v int32) {
	o.Index = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsTargetingClauseFailureResponseItem) GetErrors() []SponsoredProductsTargetMutationError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsTargetMutationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetingClauseFailureResponseItem) GetErrorsOk() ([]SponsoredProductsTargetMutationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsTargetingClauseFailureResponseItem) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsTargetMutationError and assigns it to the Errors field.
func (o *SponsoredProductsTargetingClauseFailureResponseItem) SetErrors(v []SponsoredProductsTargetMutationError) {
	o.Errors = v
}

func (o SponsoredProductsTargetingClauseFailureResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsTargetingClauseFailureResponseItem struct {
	value *SponsoredProductsTargetingClauseFailureResponseItem
	isSet bool
}

func (v NullableSponsoredProductsTargetingClauseFailureResponseItem) Get() *SponsoredProductsTargetingClauseFailureResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsTargetingClauseFailureResponseItem) Set(val *SponsoredProductsTargetingClauseFailureResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsTargetingClauseFailureResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsTargetingClauseFailureResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsTargetingClauseFailureResponseItem(val *SponsoredProductsTargetingClauseFailureResponseItem) *NullableSponsoredProductsTargetingClauseFailureResponseItem {
	return &NullableSponsoredProductsTargetingClauseFailureResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsTargetingClauseFailureResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsTargetingClauseFailureResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
