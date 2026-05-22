package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsNegativeTargetingClauseSuccessResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsNegativeTargetingClauseSuccessResponseItem{}

// SponsoredProductsNegativeTargetingClauseSuccessResponseItem struct for SponsoredProductsNegativeTargetingClauseSuccessResponseItem
type SponsoredProductsNegativeTargetingClauseSuccessResponseItem struct {
	// the NegativeTargetingClause ID
	TargetId *string `json:"targetId,omitempty"`
	// the index of the NegativeTargetingClause in the array from the request body
	Index                   int32                                     `json:"index"`
	NegativeTargetingClause *SponsoredProductsNegativeTargetingClause `json:"negativeTargetingClause,omitempty"`
}

type _SponsoredProductsNegativeTargetingClauseSuccessResponseItem SponsoredProductsNegativeTargetingClauseSuccessResponseItem

// NewSponsoredProductsNegativeTargetingClauseSuccessResponseItem instantiates a new SponsoredProductsNegativeTargetingClauseSuccessResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsNegativeTargetingClauseSuccessResponseItem(index int32) *SponsoredProductsNegativeTargetingClauseSuccessResponseItem {
	this := SponsoredProductsNegativeTargetingClauseSuccessResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsNegativeTargetingClauseSuccessResponseItemWithDefaults instantiates a new SponsoredProductsNegativeTargetingClauseSuccessResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsNegativeTargetingClauseSuccessResponseItemWithDefaults() *SponsoredProductsNegativeTargetingClauseSuccessResponseItem {
	this := SponsoredProductsNegativeTargetingClauseSuccessResponseItem{}
	return &this
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *SponsoredProductsNegativeTargetingClauseSuccessResponseItem) GetTargetId() string {
	if o == nil || IsNil(o.TargetId) {
		var ret string
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeTargetingClauseSuccessResponseItem) GetTargetIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetId) {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *SponsoredProductsNegativeTargetingClauseSuccessResponseItem) HasTargetId() bool {
	if o != nil && !IsNil(o.TargetId) {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given string and assigns it to the TargetId field.
func (o *SponsoredProductsNegativeTargetingClauseSuccessResponseItem) SetTargetId(v string) {
	o.TargetId = &v
}

// GetIndex returns the Index field value
func (o *SponsoredProductsNegativeTargetingClauseSuccessResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeTargetingClauseSuccessResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsNegativeTargetingClauseSuccessResponseItem) SetIndex(v int32) {
	o.Index = v
}

// GetNegativeTargetingClause returns the NegativeTargetingClause field value if set, zero value otherwise.
func (o *SponsoredProductsNegativeTargetingClauseSuccessResponseItem) GetNegativeTargetingClause() SponsoredProductsNegativeTargetingClause {
	if o == nil || IsNil(o.NegativeTargetingClause) {
		var ret SponsoredProductsNegativeTargetingClause
		return ret
	}
	return *o.NegativeTargetingClause
}

// GetNegativeTargetingClauseOk returns a tuple with the NegativeTargetingClause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeTargetingClauseSuccessResponseItem) GetNegativeTargetingClauseOk() (*SponsoredProductsNegativeTargetingClause, bool) {
	if o == nil || IsNil(o.NegativeTargetingClause) {
		return nil, false
	}
	return o.NegativeTargetingClause, true
}

// HasNegativeTargetingClause returns a boolean if a field has been set.
func (o *SponsoredProductsNegativeTargetingClauseSuccessResponseItem) HasNegativeTargetingClause() bool {
	if o != nil && !IsNil(o.NegativeTargetingClause) {
		return true
	}

	return false
}

// SetNegativeTargetingClause gets a reference to the given SponsoredProductsNegativeTargetingClause and assigns it to the NegativeTargetingClause field.
func (o *SponsoredProductsNegativeTargetingClauseSuccessResponseItem) SetNegativeTargetingClause(v SponsoredProductsNegativeTargetingClause) {
	o.NegativeTargetingClause = &v
}

func (o SponsoredProductsNegativeTargetingClauseSuccessResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TargetId) {
		toSerialize["targetId"] = o.TargetId
	}
	toSerialize["index"] = o.Index
	if !IsNil(o.NegativeTargetingClause) {
		toSerialize["negativeTargetingClause"] = o.NegativeTargetingClause
	}
	return toSerialize, nil
}

type NullableSponsoredProductsNegativeTargetingClauseSuccessResponseItem struct {
	value *SponsoredProductsNegativeTargetingClauseSuccessResponseItem
	isSet bool
}

func (v NullableSponsoredProductsNegativeTargetingClauseSuccessResponseItem) Get() *SponsoredProductsNegativeTargetingClauseSuccessResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsNegativeTargetingClauseSuccessResponseItem) Set(val *SponsoredProductsNegativeTargetingClauseSuccessResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsNegativeTargetingClauseSuccessResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsNegativeTargetingClauseSuccessResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsNegativeTargetingClauseSuccessResponseItem(val *SponsoredProductsNegativeTargetingClauseSuccessResponseItem) *NullableSponsoredProductsNegativeTargetingClauseSuccessResponseItem {
	return &NullableSponsoredProductsNegativeTargetingClauseSuccessResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsNegativeTargetingClauseSuccessResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsNegativeTargetingClauseSuccessResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
