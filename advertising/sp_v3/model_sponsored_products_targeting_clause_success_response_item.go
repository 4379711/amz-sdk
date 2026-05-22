package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsTargetingClauseSuccessResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsTargetingClauseSuccessResponseItem{}

// SponsoredProductsTargetingClauseSuccessResponseItem struct for SponsoredProductsTargetingClauseSuccessResponseItem
type SponsoredProductsTargetingClauseSuccessResponseItem struct {
	TargetingClause *SponsoredProductsTargetingClause `json:"targetingClause,omitempty"`
	// the targetingClause ID
	TargetId *string `json:"targetId,omitempty"`
	// the index of the targetingClause in the array from the request body
	Index int32 `json:"index"`
}

type _SponsoredProductsTargetingClauseSuccessResponseItem SponsoredProductsTargetingClauseSuccessResponseItem

// NewSponsoredProductsTargetingClauseSuccessResponseItem instantiates a new SponsoredProductsTargetingClauseSuccessResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsTargetingClauseSuccessResponseItem(index int32) *SponsoredProductsTargetingClauseSuccessResponseItem {
	this := SponsoredProductsTargetingClauseSuccessResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsTargetingClauseSuccessResponseItemWithDefaults instantiates a new SponsoredProductsTargetingClauseSuccessResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsTargetingClauseSuccessResponseItemWithDefaults() *SponsoredProductsTargetingClauseSuccessResponseItem {
	this := SponsoredProductsTargetingClauseSuccessResponseItem{}
	return &this
}

// GetTargetingClause returns the TargetingClause field value if set, zero value otherwise.
func (o *SponsoredProductsTargetingClauseSuccessResponseItem) GetTargetingClause() SponsoredProductsTargetingClause {
	if o == nil || IsNil(o.TargetingClause) {
		var ret SponsoredProductsTargetingClause
		return ret
	}
	return *o.TargetingClause
}

// GetTargetingClauseOk returns a tuple with the TargetingClause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetingClauseSuccessResponseItem) GetTargetingClauseOk() (*SponsoredProductsTargetingClause, bool) {
	if o == nil || IsNil(o.TargetingClause) {
		return nil, false
	}
	return o.TargetingClause, true
}

// HasTargetingClause returns a boolean if a field has been set.
func (o *SponsoredProductsTargetingClauseSuccessResponseItem) HasTargetingClause() bool {
	if o != nil && !IsNil(o.TargetingClause) {
		return true
	}

	return false
}

// SetTargetingClause gets a reference to the given SponsoredProductsTargetingClause and assigns it to the TargetingClause field.
func (o *SponsoredProductsTargetingClauseSuccessResponseItem) SetTargetingClause(v SponsoredProductsTargetingClause) {
	o.TargetingClause = &v
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *SponsoredProductsTargetingClauseSuccessResponseItem) GetTargetId() string {
	if o == nil || IsNil(o.TargetId) {
		var ret string
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetingClauseSuccessResponseItem) GetTargetIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetId) {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *SponsoredProductsTargetingClauseSuccessResponseItem) HasTargetId() bool {
	if o != nil && !IsNil(o.TargetId) {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given string and assigns it to the TargetId field.
func (o *SponsoredProductsTargetingClauseSuccessResponseItem) SetTargetId(v string) {
	o.TargetId = &v
}

// GetIndex returns the Index field value
func (o *SponsoredProductsTargetingClauseSuccessResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsTargetingClauseSuccessResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsTargetingClauseSuccessResponseItem) SetIndex(v int32) {
	o.Index = v
}

func (o SponsoredProductsTargetingClauseSuccessResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TargetingClause) {
		toSerialize["targetingClause"] = o.TargetingClause
	}
	if !IsNil(o.TargetId) {
		toSerialize["targetId"] = o.TargetId
	}
	toSerialize["index"] = o.Index
	return toSerialize, nil
}

type NullableSponsoredProductsTargetingClauseSuccessResponseItem struct {
	value *SponsoredProductsTargetingClauseSuccessResponseItem
	isSet bool
}

func (v NullableSponsoredProductsTargetingClauseSuccessResponseItem) Get() *SponsoredProductsTargetingClauseSuccessResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsTargetingClauseSuccessResponseItem) Set(val *SponsoredProductsTargetingClauseSuccessResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsTargetingClauseSuccessResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsTargetingClauseSuccessResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsTargetingClauseSuccessResponseItem(val *SponsoredProductsTargetingClauseSuccessResponseItem) *NullableSponsoredProductsTargetingClauseSuccessResponseItem {
	return &NullableSponsoredProductsTargetingClauseSuccessResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsTargetingClauseSuccessResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsTargetingClauseSuccessResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
