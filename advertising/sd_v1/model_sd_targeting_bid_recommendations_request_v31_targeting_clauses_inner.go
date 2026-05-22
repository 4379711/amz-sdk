package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDTargetingBidRecommendationsRequestV31TargetingClausesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDTargetingBidRecommendationsRequestV31TargetingClausesInner{}

// SDTargetingBidRecommendationsRequestV31TargetingClausesInner struct for SDTargetingBidRecommendationsRequestV31TargetingClausesInner
type SDTargetingBidRecommendationsRequestV31TargetingClausesInner struct {
	TargetingClause SDTargetingClauseV31 `json:"targetingClause"`
}

type _SDTargetingBidRecommendationsRequestV31TargetingClausesInner SDTargetingBidRecommendationsRequestV31TargetingClausesInner

// NewSDTargetingBidRecommendationsRequestV31TargetingClausesInner instantiates a new SDTargetingBidRecommendationsRequestV31TargetingClausesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDTargetingBidRecommendationsRequestV31TargetingClausesInner(targetingClause SDTargetingClauseV31) *SDTargetingBidRecommendationsRequestV31TargetingClausesInner {
	this := SDTargetingBidRecommendationsRequestV31TargetingClausesInner{}
	this.TargetingClause = targetingClause
	return &this
}

// NewSDTargetingBidRecommendationsRequestV31TargetingClausesInnerWithDefaults instantiates a new SDTargetingBidRecommendationsRequestV31TargetingClausesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDTargetingBidRecommendationsRequestV31TargetingClausesInnerWithDefaults() *SDTargetingBidRecommendationsRequestV31TargetingClausesInner {
	this := SDTargetingBidRecommendationsRequestV31TargetingClausesInner{}
	return &this
}

// GetTargetingClause returns the TargetingClause field value
func (o *SDTargetingBidRecommendationsRequestV31TargetingClausesInner) GetTargetingClause() SDTargetingClauseV31 {
	if o == nil {
		var ret SDTargetingClauseV31
		return ret
	}

	return o.TargetingClause
}

// GetTargetingClauseOk returns a tuple with the TargetingClause field value
// and a boolean to check if the value has been set.
func (o *SDTargetingBidRecommendationsRequestV31TargetingClausesInner) GetTargetingClauseOk() (*SDTargetingClauseV31, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetingClause, true
}

// SetTargetingClause sets field value
func (o *SDTargetingBidRecommendationsRequestV31TargetingClausesInner) SetTargetingClause(v SDTargetingClauseV31) {
	o.TargetingClause = v
}

func (o SDTargetingBidRecommendationsRequestV31TargetingClausesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["targetingClause"] = o.TargetingClause
	return toSerialize, nil
}

type NullableSDTargetingBidRecommendationsRequestV31TargetingClausesInner struct {
	value *SDTargetingBidRecommendationsRequestV31TargetingClausesInner
	isSet bool
}

func (v NullableSDTargetingBidRecommendationsRequestV31TargetingClausesInner) Get() *SDTargetingBidRecommendationsRequestV31TargetingClausesInner {
	return v.value
}

func (v *NullableSDTargetingBidRecommendationsRequestV31TargetingClausesInner) Set(val *SDTargetingBidRecommendationsRequestV31TargetingClausesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSDTargetingBidRecommendationsRequestV31TargetingClausesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSDTargetingBidRecommendationsRequestV31TargetingClausesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDTargetingBidRecommendationsRequestV31TargetingClausesInner(val *SDTargetingBidRecommendationsRequestV31TargetingClausesInner) *NullableSDTargetingBidRecommendationsRequestV31TargetingClausesInner {
	return &NullableSDTargetingBidRecommendationsRequestV31TargetingClausesInner{value: val, isSet: true}
}

func (v NullableSDTargetingBidRecommendationsRequestV31TargetingClausesInner) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDTargetingBidRecommendationsRequestV31TargetingClausesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
