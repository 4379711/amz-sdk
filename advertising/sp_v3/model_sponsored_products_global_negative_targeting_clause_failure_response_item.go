package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGlobalNegativeTargetingClauseFailureResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGlobalNegativeTargetingClauseFailureResponseItem{}

// SponsoredProductsGlobalNegativeTargetingClauseFailureResponseItem struct for SponsoredProductsGlobalNegativeTargetingClauseFailureResponseItem
type SponsoredProductsGlobalNegativeTargetingClauseFailureResponseItem struct {
	// the index of the NegativeTargetingClause in the array from the request body
	Index int32 `json:"index"`
	// A list of validation errors
	Errors []SponsoredProductsNegativeTargetMutationError `json:"errors,omitempty"`
}

type _SponsoredProductsGlobalNegativeTargetingClauseFailureResponseItem SponsoredProductsGlobalNegativeTargetingClauseFailureResponseItem

// NewSponsoredProductsGlobalNegativeTargetingClauseFailureResponseItem instantiates a new SponsoredProductsGlobalNegativeTargetingClauseFailureResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGlobalNegativeTargetingClauseFailureResponseItem(index int32) *SponsoredProductsGlobalNegativeTargetingClauseFailureResponseItem {
	this := SponsoredProductsGlobalNegativeTargetingClauseFailureResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsGlobalNegativeTargetingClauseFailureResponseItemWithDefaults instantiates a new SponsoredProductsGlobalNegativeTargetingClauseFailureResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGlobalNegativeTargetingClauseFailureResponseItemWithDefaults() *SponsoredProductsGlobalNegativeTargetingClauseFailureResponseItem {
	this := SponsoredProductsGlobalNegativeTargetingClauseFailureResponseItem{}
	return &this
}

// GetIndex returns the Index field value
func (o *SponsoredProductsGlobalNegativeTargetingClauseFailureResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalNegativeTargetingClauseFailureResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsGlobalNegativeTargetingClauseFailureResponseItem) SetIndex(v int32) {
	o.Index = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalNegativeTargetingClauseFailureResponseItem) GetErrors() []SponsoredProductsNegativeTargetMutationError {
	if o == nil || IsNil(o.Errors) {
		var ret []SponsoredProductsNegativeTargetMutationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalNegativeTargetingClauseFailureResponseItem) GetErrorsOk() ([]SponsoredProductsNegativeTargetMutationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalNegativeTargetingClauseFailureResponseItem) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SponsoredProductsNegativeTargetMutationError and assigns it to the Errors field.
func (o *SponsoredProductsGlobalNegativeTargetingClauseFailureResponseItem) SetErrors(v []SponsoredProductsNegativeTargetMutationError) {
	o.Errors = v
}

func (o SponsoredProductsGlobalNegativeTargetingClauseFailureResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableSponsoredProductsGlobalNegativeTargetingClauseFailureResponseItem struct {
	value *SponsoredProductsGlobalNegativeTargetingClauseFailureResponseItem
	isSet bool
}

func (v NullableSponsoredProductsGlobalNegativeTargetingClauseFailureResponseItem) Get() *SponsoredProductsGlobalNegativeTargetingClauseFailureResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsGlobalNegativeTargetingClauseFailureResponseItem) Set(val *SponsoredProductsGlobalNegativeTargetingClauseFailureResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalNegativeTargetingClauseFailureResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalNegativeTargetingClauseFailureResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalNegativeTargetingClauseFailureResponseItem(val *SponsoredProductsGlobalNegativeTargetingClauseFailureResponseItem) *NullableSponsoredProductsGlobalNegativeTargetingClauseFailureResponseItem {
	return &NullableSponsoredProductsGlobalNegativeTargetingClauseFailureResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalNegativeTargetingClauseFailureResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalNegativeTargetingClauseFailureResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
