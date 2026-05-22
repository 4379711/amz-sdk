package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGlobalTargetingClauseFailureResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGlobalTargetingClauseFailureResponseItem{}

// SponsoredProductsGlobalTargetingClauseFailureResponseItem struct for SponsoredProductsGlobalTargetingClauseFailureResponseItem
type SponsoredProductsGlobalTargetingClauseFailureResponseItem struct {
	// the index of the targetingClause in the array from the request body
	Index int32 `json:"index"`
	// A list of validation errors
	Errors []SponsoredProductsTargetMutationError `json:"errors,omitempty"`
}

type _SponsoredProductsGlobalTargetingClauseFailureResponseItem SponsoredProductsGlobalTargetingClauseFailureResponseItem

// NewSponsoredProductsGlobalTargetingClauseFailureResponseItem instantiates a new SponsoredProductsGlobalTargetingClauseFailureResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGlobalTargetingClauseFailureResponseItem(index int32) *SponsoredProductsGlobalTargetingClauseFailureResponseItem {
	this := SponsoredProductsGlobalTargetingClauseFailureResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsGlobalTargetingClauseFailureResponseItemWithDefaults instantiates a new SponsoredProductsGlobalTargetingClauseFailureResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGlobalTargetingClauseFailureResponseItemWithDefaults() *SponsoredProductsGlobalTargetingClauseFailureResponseItem {
	this := SponsoredProductsGlobalTargetingClauseFailureResponseItem{}
	return &this
}

// GetIndex returns the Index field value
func (o *SponsoredProductsGlobalTargetingClauseFailureResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalTargetingClauseFailureResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsGlobalTargetingClauseFailureResponseItem) SetIndex(v int32) {
	o.Index = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalTargetingClauseFailureResponseItem) GetErrors() []SponsoredProductsTargetMutationError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsTargetMutationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalTargetingClauseFailureResponseItem) GetErrorsOk() ([]SponsoredProductsTargetMutationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalTargetingClauseFailureResponseItem) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsTargetMutationError and assigns it to the Errors field.
func (o *SponsoredProductsGlobalTargetingClauseFailureResponseItem) SetErrors(v []SponsoredProductsTargetMutationError) {
	o.Errors = v
}

func (o SponsoredProductsGlobalTargetingClauseFailureResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsGlobalTargetingClauseFailureResponseItem struct {
	value *SponsoredProductsGlobalTargetingClauseFailureResponseItem
	isSet bool
}

func (v NullableSponsoredProductsGlobalTargetingClauseFailureResponseItem) Get() *SponsoredProductsGlobalTargetingClauseFailureResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsGlobalTargetingClauseFailureResponseItem) Set(val *SponsoredProductsGlobalTargetingClauseFailureResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalTargetingClauseFailureResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalTargetingClauseFailureResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalTargetingClauseFailureResponseItem(val *SponsoredProductsGlobalTargetingClauseFailureResponseItem) *NullableSponsoredProductsGlobalTargetingClauseFailureResponseItem {
	return &NullableSponsoredProductsGlobalTargetingClauseFailureResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalTargetingClauseFailureResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalTargetingClauseFailureResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
