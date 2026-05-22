package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftNegativeTargetingClauseFailureResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftNegativeTargetingClauseFailureResponseItem{}

// SponsoredProductsDraftNegativeTargetingClauseFailureResponseItem struct for SponsoredProductsDraftNegativeTargetingClauseFailureResponseItem
type SponsoredProductsDraftNegativeTargetingClauseFailureResponseItem struct {
	// the index of the DraftNegativeTargetingClause in the array from the request body
	Index int32 `json:"index"`
	// A list of validation errors
	Errors []SponsoredProductsDraftNegativeTargetMutationError `json:"errors,omitempty"`
}

type _SponsoredProductsDraftNegativeTargetingClauseFailureResponseItem SponsoredProductsDraftNegativeTargetingClauseFailureResponseItem

// NewSponsoredProductsDraftNegativeTargetingClauseFailureResponseItem instantiates a new SponsoredProductsDraftNegativeTargetingClauseFailureResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftNegativeTargetingClauseFailureResponseItem(index int32) *SponsoredProductsDraftNegativeTargetingClauseFailureResponseItem {
	this := SponsoredProductsDraftNegativeTargetingClauseFailureResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsDraftNegativeTargetingClauseFailureResponseItemWithDefaults instantiates a new SponsoredProductsDraftNegativeTargetingClauseFailureResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftNegativeTargetingClauseFailureResponseItemWithDefaults() *SponsoredProductsDraftNegativeTargetingClauseFailureResponseItem {
	this := SponsoredProductsDraftNegativeTargetingClauseFailureResponseItem{}
	return &this
}

// GetIndex returns the Index field value
func (o *SponsoredProductsDraftNegativeTargetingClauseFailureResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftNegativeTargetingClauseFailureResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsDraftNegativeTargetingClauseFailureResponseItem) SetIndex(v int32) {
	o.Index = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsDraftNegativeTargetingClauseFailureResponseItem) GetErrors() []SponsoredProductsDraftNegativeTargetMutationError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsDraftNegativeTargetMutationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftNegativeTargetingClauseFailureResponseItem) GetErrorsOk() ([]SponsoredProductsDraftNegativeTargetMutationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsDraftNegativeTargetingClauseFailureResponseItem) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsDraftNegativeTargetMutationError and assigns it to the Errors field.
func (o *SponsoredProductsDraftNegativeTargetingClauseFailureResponseItem) SetErrors(v []SponsoredProductsDraftNegativeTargetMutationError) {
	o.Errors = v
}

func (o SponsoredProductsDraftNegativeTargetingClauseFailureResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDraftNegativeTargetingClauseFailureResponseItem struct {
	value *SponsoredProductsDraftNegativeTargetingClauseFailureResponseItem
	isSet bool
}

func (v NullableSponsoredProductsDraftNegativeTargetingClauseFailureResponseItem) Get() *SponsoredProductsDraftNegativeTargetingClauseFailureResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsDraftNegativeTargetingClauseFailureResponseItem) Set(val *SponsoredProductsDraftNegativeTargetingClauseFailureResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftNegativeTargetingClauseFailureResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftNegativeTargetingClauseFailureResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftNegativeTargetingClauseFailureResponseItem(val *SponsoredProductsDraftNegativeTargetingClauseFailureResponseItem) *NullableSponsoredProductsDraftNegativeTargetingClauseFailureResponseItem {
	return &NullableSponsoredProductsDraftNegativeTargetingClauseFailureResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftNegativeTargetingClauseFailureResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftNegativeTargetingClauseFailureResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
