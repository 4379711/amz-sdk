package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGlobalTargetingClauseSuccessResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGlobalTargetingClauseSuccessResponseItem{}

// SponsoredProductsGlobalTargetingClauseSuccessResponseItem struct for SponsoredProductsGlobalTargetingClauseSuccessResponseItem
type SponsoredProductsGlobalTargetingClauseSuccessResponseItem struct {
	TargetingClause *SponsoredProductsGlobalTargetingClause `json:"targetingClause,omitempty"`
	// the targetingClause ID
	TargetId *string `json:"targetId,omitempty"`
	// the index of the targetingClause in the array from the request body
	Index int32 `json:"index"`
}

type _SponsoredProductsGlobalTargetingClauseSuccessResponseItem SponsoredProductsGlobalTargetingClauseSuccessResponseItem

// NewSponsoredProductsGlobalTargetingClauseSuccessResponseItem instantiates a new SponsoredProductsGlobalTargetingClauseSuccessResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGlobalTargetingClauseSuccessResponseItem(index int32) *SponsoredProductsGlobalTargetingClauseSuccessResponseItem {
	this := SponsoredProductsGlobalTargetingClauseSuccessResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsGlobalTargetingClauseSuccessResponseItemWithDefaults instantiates a new SponsoredProductsGlobalTargetingClauseSuccessResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGlobalTargetingClauseSuccessResponseItemWithDefaults() *SponsoredProductsGlobalTargetingClauseSuccessResponseItem {
	this := SponsoredProductsGlobalTargetingClauseSuccessResponseItem{}
	return &this
}

// GetTargetingClause returns the TargetingClause field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalTargetingClauseSuccessResponseItem) GetTargetingClause() SponsoredProductsGlobalTargetingClause {
	if o == nil || IsNil(o.TargetingClause) {
		var ret SponsoredProductsGlobalTargetingClause
		return ret
	}
	return *o.TargetingClause
}

// GetTargetingClauseOk returns a tuple with the TargetingClause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalTargetingClauseSuccessResponseItem) GetTargetingClauseOk() (*SponsoredProductsGlobalTargetingClause, bool) {
	if o == nil || IsNil(o.TargetingClause) {
		return nil, false
	}
	return o.TargetingClause, true
}

// HasTargetingClause returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalTargetingClauseSuccessResponseItem) HasTargetingClause() bool {
	if o != nil && !IsNil(o.TargetingClause) {
		return true
	}

	return false
}

// SetTargetingClause gets a reference to the given SponsoredProductsGlobalTargetingClause and assigns it to the TargetingClause field.
func (o *SponsoredProductsGlobalTargetingClauseSuccessResponseItem) SetTargetingClause(v SponsoredProductsGlobalTargetingClause) {
	o.TargetingClause = &v
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalTargetingClauseSuccessResponseItem) GetTargetId() string {
	if o == nil || IsNil(o.TargetId) {
		var ret string
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalTargetingClauseSuccessResponseItem) GetTargetIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetId) {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalTargetingClauseSuccessResponseItem) HasTargetId() bool {
	if o != nil && !IsNil(o.TargetId) {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given string and assigns it to the TargetId field.
func (o *SponsoredProductsGlobalTargetingClauseSuccessResponseItem) SetTargetId(v string) {
	o.TargetId = &v
}

// GetIndex returns the Index field value
func (o *SponsoredProductsGlobalTargetingClauseSuccessResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalTargetingClauseSuccessResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsGlobalTargetingClauseSuccessResponseItem) SetIndex(v int32) {
	o.Index = v
}

func (o SponsoredProductsGlobalTargetingClauseSuccessResponseItem) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsGlobalTargetingClauseSuccessResponseItem struct {
	value *SponsoredProductsGlobalTargetingClauseSuccessResponseItem
	isSet bool
}

func (v NullableSponsoredProductsGlobalTargetingClauseSuccessResponseItem) Get() *SponsoredProductsGlobalTargetingClauseSuccessResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsGlobalTargetingClauseSuccessResponseItem) Set(val *SponsoredProductsGlobalTargetingClauseSuccessResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalTargetingClauseSuccessResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalTargetingClauseSuccessResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalTargetingClauseSuccessResponseItem(val *SponsoredProductsGlobalTargetingClauseSuccessResponseItem) *NullableSponsoredProductsGlobalTargetingClauseSuccessResponseItem {
	return &NullableSponsoredProductsGlobalTargetingClauseSuccessResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalTargetingClauseSuccessResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalTargetingClauseSuccessResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
