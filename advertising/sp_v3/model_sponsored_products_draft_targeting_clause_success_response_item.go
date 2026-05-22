package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftTargetingClauseSuccessResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftTargetingClauseSuccessResponseItem{}

// SponsoredProductsDraftTargetingClauseSuccessResponseItem struct for SponsoredProductsDraftTargetingClauseSuccessResponseItem
type SponsoredProductsDraftTargetingClauseSuccessResponseItem struct {
	TargetingClause *SponsoredProductsDraftTargetingClause `json:"targetingClause,omitempty"`
	// the draftTargetingClause ID
	TargetId *string `json:"targetId,omitempty"`
	// the index of the draftTargetingClause in the array from the request body
	Index int32 `json:"index"`
}

type _SponsoredProductsDraftTargetingClauseSuccessResponseItem SponsoredProductsDraftTargetingClauseSuccessResponseItem

// NewSponsoredProductsDraftTargetingClauseSuccessResponseItem instantiates a new SponsoredProductsDraftTargetingClauseSuccessResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftTargetingClauseSuccessResponseItem(index int32) *SponsoredProductsDraftTargetingClauseSuccessResponseItem {
	this := SponsoredProductsDraftTargetingClauseSuccessResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsDraftTargetingClauseSuccessResponseItemWithDefaults instantiates a new SponsoredProductsDraftTargetingClauseSuccessResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftTargetingClauseSuccessResponseItemWithDefaults() *SponsoredProductsDraftTargetingClauseSuccessResponseItem {
	this := SponsoredProductsDraftTargetingClauseSuccessResponseItem{}
	return &this
}

// GetTargetingClause returns the TargetingClause field value if set, zero value otherwise.
func (o *SponsoredProductsDraftTargetingClauseSuccessResponseItem) GetTargetingClause() SponsoredProductsDraftTargetingClause {
	if o == nil || IsNil(o.TargetingClause) {
		var ret SponsoredProductsDraftTargetingClause
		return ret
	}
	return *o.TargetingClause
}

// GetTargetingClauseOk returns a tuple with the TargetingClause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftTargetingClauseSuccessResponseItem) GetTargetingClauseOk() (*SponsoredProductsDraftTargetingClause, bool) {
	if o == nil || IsNil(o.TargetingClause) {
		return nil, false
	}
	return o.TargetingClause, true
}

// HasTargetingClause returns a boolean if a field has been set.
func (o *SponsoredProductsDraftTargetingClauseSuccessResponseItem) HasTargetingClause() bool {
	if o != nil && !IsNil(o.TargetingClause) {
		return true
	}

	return false
}

// SetTargetingClause gets a reference to the given SponsoredProductsDraftTargetingClause and assigns it to the TargetingClause field.
func (o *SponsoredProductsDraftTargetingClauseSuccessResponseItem) SetTargetingClause(v SponsoredProductsDraftTargetingClause) {
	o.TargetingClause = &v
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *SponsoredProductsDraftTargetingClauseSuccessResponseItem) GetTargetId() string {
	if o == nil || IsNil(o.TargetId) {
		var ret string
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftTargetingClauseSuccessResponseItem) GetTargetIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetId) {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *SponsoredProductsDraftTargetingClauseSuccessResponseItem) HasTargetId() bool {
	if o != nil && !IsNil(o.TargetId) {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given string and assigns it to the TargetId field.
func (o *SponsoredProductsDraftTargetingClauseSuccessResponseItem) SetTargetId(v string) {
	o.TargetId = &v
}

// GetIndex returns the Index field value
func (o *SponsoredProductsDraftTargetingClauseSuccessResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftTargetingClauseSuccessResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsDraftTargetingClauseSuccessResponseItem) SetIndex(v int32) {
	o.Index = v
}

func (o SponsoredProductsDraftTargetingClauseSuccessResponseItem) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsDraftTargetingClauseSuccessResponseItem struct {
	value *SponsoredProductsDraftTargetingClauseSuccessResponseItem
	isSet bool
}

func (v NullableSponsoredProductsDraftTargetingClauseSuccessResponseItem) Get() *SponsoredProductsDraftTargetingClauseSuccessResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsDraftTargetingClauseSuccessResponseItem) Set(val *SponsoredProductsDraftTargetingClauseSuccessResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftTargetingClauseSuccessResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftTargetingClauseSuccessResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftTargetingClauseSuccessResponseItem(val *SponsoredProductsDraftTargetingClauseSuccessResponseItem) *NullableSponsoredProductsDraftTargetingClauseSuccessResponseItem {
	return &NullableSponsoredProductsDraftTargetingClauseSuccessResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftTargetingClauseSuccessResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftTargetingClauseSuccessResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
