package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem{}

// SponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem struct for SponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem
type SponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem struct {
	// the NegativeTargetingClause ID
	TargetId *string `json:"targetId,omitempty"`
	// the index of the NegativeTargetingClause in the array from the request body
	Index                   int32                                           `json:"index"`
	NegativeTargetingClause *SponsoredProductsGlobalNegativeTargetingClause `json:"negativeTargetingClause,omitempty"`
}

type _SponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem SponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem

// NewSponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem instantiates a new SponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem(index int32) *SponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem {
	this := SponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItemWithDefaults instantiates a new SponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItemWithDefaults() *SponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem {
	this := SponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem{}
	return &this
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem) GetTargetId() string {
	if o == nil || IsNil(o.TargetId) {
		var ret string
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem) GetTargetIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetId) {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem) HasTargetId() bool {
	if o != nil && !IsNil(o.TargetId) {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given string and assigns it to the TargetId field.
func (o *SponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem) SetTargetId(v string) {
	o.TargetId = &v
}

// GetIndex returns the Index field value
func (o *SponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem) SetIndex(v int32) {
	o.Index = v
}

// GetNegativeTargetingClause returns the NegativeTargetingClause field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem) GetNegativeTargetingClause() SponsoredProductsGlobalNegativeTargetingClause {
	if o == nil || IsNil(o.NegativeTargetingClause) {
		var ret SponsoredProductsGlobalNegativeTargetingClause
		return ret
	}
	return *o.NegativeTargetingClause
}

// GetNegativeTargetingClauseOk returns a tuple with the NegativeTargetingClause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem) GetNegativeTargetingClauseOk() (*SponsoredProductsGlobalNegativeTargetingClause, bool) {
	if o == nil || IsNil(o.NegativeTargetingClause) {
		return nil, false
	}
	return o.NegativeTargetingClause, true
}

// HasNegativeTargetingClause returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem) HasNegativeTargetingClause() bool {
	if o != nil && !IsNil(o.NegativeTargetingClause) {
		return true
	}

	return false
}

// SetNegativeTargetingClause gets a reference to the given SponsoredProductsGlobalNegativeTargetingClause and assigns it to the NegativeTargetingClause field.
func (o *SponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem) SetNegativeTargetingClause(v SponsoredProductsGlobalNegativeTargetingClause) {
	o.NegativeTargetingClause = &v
}

func (o SponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem struct {
	value *SponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem
	isSet bool
}

func (v NullableSponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem) Get() *SponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem) Set(val *SponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem(val *SponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem) *NullableSponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem {
	return &NullableSponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalNegativeTargetingClauseSuccessResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
