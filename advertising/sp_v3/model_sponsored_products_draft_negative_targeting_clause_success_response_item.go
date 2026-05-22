package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem{}

// SponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem struct for SponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem
type SponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem struct {
	// the DraftNegativeTargetingClause ID
	TargetId *string `json:"targetId,omitempty"`
	// the index of the DraftNegativeTargetingClause in the array from the request body
	Index                   int32                                          `json:"index"`
	NegativeTargetingClause *SponsoredProductsDraftNegativeTargetingClause `json:"negativeTargetingClause,omitempty"`
}

type _SponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem SponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem

// NewSponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem instantiates a new SponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem(index int32) *SponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem {
	this := SponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsDraftNegativeTargetingClauseSuccessResponseItemWithDefaults instantiates a new SponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftNegativeTargetingClauseSuccessResponseItemWithDefaults() *SponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem {
	this := SponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem{}
	return &this
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *SponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem) GetTargetId() string {
	if o == nil || IsNil(o.TargetId) {
		var ret string
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem) GetTargetIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetId) {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *SponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem) HasTargetId() bool {
	if o != nil && !IsNil(o.TargetId) {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given string and assigns it to the TargetId field.
func (o *SponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem) SetTargetId(v string) {
	o.TargetId = &v
}

// GetIndex returns the Index field value
func (o *SponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem) SetIndex(v int32) {
	o.Index = v
}

// GetNegativeTargetingClause returns the NegativeTargetingClause field value if set, zero value otherwise.
func (o *SponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem) GetNegativeTargetingClause() SponsoredProductsDraftNegativeTargetingClause {
	if o == nil || IsNil(o.NegativeTargetingClause) {
		var ret SponsoredProductsDraftNegativeTargetingClause
		return ret
	}
	return *o.NegativeTargetingClause
}

// GetNegativeTargetingClauseOk returns a tuple with the NegativeTargetingClause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem) GetNegativeTargetingClauseOk() (*SponsoredProductsDraftNegativeTargetingClause, bool) {
	if o == nil || IsNil(o.NegativeTargetingClause) {
		return nil, false
	}
	return o.NegativeTargetingClause, true
}

// HasNegativeTargetingClause returns a boolean if a field has been set.
func (o *SponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem) HasNegativeTargetingClause() bool {
	if o != nil && !IsNil(o.NegativeTargetingClause) {
		return true
	}

	return false
}

// SetNegativeTargetingClause gets a reference to the given SponsoredProductsDraftNegativeTargetingClause and assigns it to the NegativeTargetingClause field.
func (o *SponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem) SetNegativeTargetingClause(v SponsoredProductsDraftNegativeTargetingClause) {
	o.NegativeTargetingClause = &v
}

func (o SponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem struct {
	value *SponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem
	isSet bool
}

func (v NullableSponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem) Get() *SponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem) Set(val *SponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem(val *SponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem) *NullableSponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem {
	return &NullableSponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftNegativeTargetingClauseSuccessResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
