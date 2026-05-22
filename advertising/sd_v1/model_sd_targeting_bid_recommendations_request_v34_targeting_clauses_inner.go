package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDTargetingBidRecommendationsRequestV34TargetingClausesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDTargetingBidRecommendationsRequestV34TargetingClausesInner{}

// SDTargetingBidRecommendationsRequestV34TargetingClausesInner struct for SDTargetingBidRecommendationsRequestV34TargetingClausesInner
type SDTargetingBidRecommendationsRequestV34TargetingClausesInner struct {
	TargetingClause SDTargetingClauseV32 `json:"targetingClause"`
}

type _SDTargetingBidRecommendationsRequestV34TargetingClausesInner SDTargetingBidRecommendationsRequestV34TargetingClausesInner

// NewSDTargetingBidRecommendationsRequestV34TargetingClausesInner instantiates a new SDTargetingBidRecommendationsRequestV34TargetingClausesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDTargetingBidRecommendationsRequestV34TargetingClausesInner(targetingClause SDTargetingClauseV32) *SDTargetingBidRecommendationsRequestV34TargetingClausesInner {
	this := SDTargetingBidRecommendationsRequestV34TargetingClausesInner{}
	this.TargetingClause = targetingClause
	return &this
}

// NewSDTargetingBidRecommendationsRequestV34TargetingClausesInnerWithDefaults instantiates a new SDTargetingBidRecommendationsRequestV34TargetingClausesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDTargetingBidRecommendationsRequestV34TargetingClausesInnerWithDefaults() *SDTargetingBidRecommendationsRequestV34TargetingClausesInner {
	this := SDTargetingBidRecommendationsRequestV34TargetingClausesInner{}
	return &this
}

// GetTargetingClause returns the TargetingClause field value
func (o *SDTargetingBidRecommendationsRequestV34TargetingClausesInner) GetTargetingClause() SDTargetingClauseV32 {
	if o == nil {
		var ret SDTargetingClauseV32
		return ret
	}

	return o.TargetingClause
}

// GetTargetingClauseOk returns a tuple with the TargetingClause field value
// and a boolean to check if the value has been set.
func (o *SDTargetingBidRecommendationsRequestV34TargetingClausesInner) GetTargetingClauseOk() (*SDTargetingClauseV32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetingClause, true
}

// SetTargetingClause sets field value
func (o *SDTargetingBidRecommendationsRequestV34TargetingClausesInner) SetTargetingClause(v SDTargetingClauseV32) {
	o.TargetingClause = v
}

func (o SDTargetingBidRecommendationsRequestV34TargetingClausesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["targetingClause"] = o.TargetingClause
	return toSerialize, nil
}

type NullableSDTargetingBidRecommendationsRequestV34TargetingClausesInner struct {
	value *SDTargetingBidRecommendationsRequestV34TargetingClausesInner
	isSet bool
}

func (v NullableSDTargetingBidRecommendationsRequestV34TargetingClausesInner) Get() *SDTargetingBidRecommendationsRequestV34TargetingClausesInner {
	return v.value
}

func (v *NullableSDTargetingBidRecommendationsRequestV34TargetingClausesInner) Set(val *SDTargetingBidRecommendationsRequestV34TargetingClausesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSDTargetingBidRecommendationsRequestV34TargetingClausesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSDTargetingBidRecommendationsRequestV34TargetingClausesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDTargetingBidRecommendationsRequestV34TargetingClausesInner(val *SDTargetingBidRecommendationsRequestV34TargetingClausesInner) *NullableSDTargetingBidRecommendationsRequestV34TargetingClausesInner {
	return &NullableSDTargetingBidRecommendationsRequestV34TargetingClausesInner{value: val, isSet: true}
}

func (v NullableSDTargetingBidRecommendationsRequestV34TargetingClausesInner) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDTargetingBidRecommendationsRequestV34TargetingClausesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
