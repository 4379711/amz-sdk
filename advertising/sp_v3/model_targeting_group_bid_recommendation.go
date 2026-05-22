package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the TargetingGroupBidRecommendation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TargetingGroupBidRecommendation{}

// TargetingGroupBidRecommendation Contains suggested recommendation for the auto targeting group.
type TargetingGroupBidRecommendation struct {
	// The suggested bid value associated with this targeting.
	SuggestedBid *float64 `json:"suggestedBid,omitempty"`
	// The target identifier.
	TargetId *string `json:"targetId,omitempty"`
	// The type of targeting group expression. | Value | Description | | --- | --- | | `LOOSE_MATCH` | This will show your ad to shoppers who use search terms loosely related to your products.| | `CLOSE_MATCH` | This will show your ad to shoppers who use search terms closely related to your products.| | `COMPLEMENTS` | This will show your ad to shoppers who view the detail pages of products that complement your product.| | `SUBSTITUTES` | This will show your ad to shoppers who use detail pages of products similar to yours.|
	TargetingGroupExpression *string `json:"targetingGroupExpression,omitempty"`
	// Type of suggested action.
	Action *string `json:"action,omitempty"`
	// The ad group identifier.
	AdGroupId *string `json:"adGroupId,omitempty"`
}

// NewTargetingGroupBidRecommendation instantiates a new TargetingGroupBidRecommendation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetingGroupBidRecommendation() *TargetingGroupBidRecommendation {
	this := TargetingGroupBidRecommendation{}
	return &this
}

// NewTargetingGroupBidRecommendationWithDefaults instantiates a new TargetingGroupBidRecommendation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetingGroupBidRecommendationWithDefaults() *TargetingGroupBidRecommendation {
	this := TargetingGroupBidRecommendation{}
	return &this
}

// GetSuggestedBid returns the SuggestedBid field value if set, zero value otherwise.
func (o *TargetingGroupBidRecommendation) GetSuggestedBid() float64 {
	if o == nil || IsNil(o.SuggestedBid) {
		var ret float64
		return ret
	}
	return *o.SuggestedBid
}

// GetSuggestedBidOk returns a tuple with the SuggestedBid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetingGroupBidRecommendation) GetSuggestedBidOk() (*float64, bool) {
	if o == nil || IsNil(o.SuggestedBid) {
		return nil, false
	}
	return o.SuggestedBid, true
}

// HasSuggestedBid returns a boolean if a field has been set.
func (o *TargetingGroupBidRecommendation) HasSuggestedBid() bool {
	if o != nil && !IsNil(o.SuggestedBid) {
		return true
	}

	return false
}

// SetSuggestedBid gets a reference to the given float64 and assigns it to the SuggestedBid field.
func (o *TargetingGroupBidRecommendation) SetSuggestedBid(v float64) {
	o.SuggestedBid = &v
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *TargetingGroupBidRecommendation) GetTargetId() string {
	if o == nil || IsNil(o.TargetId) {
		var ret string
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetingGroupBidRecommendation) GetTargetIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetId) {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *TargetingGroupBidRecommendation) HasTargetId() bool {
	if o != nil && !IsNil(o.TargetId) {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given string and assigns it to the TargetId field.
func (o *TargetingGroupBidRecommendation) SetTargetId(v string) {
	o.TargetId = &v
}

// GetTargetingGroupExpression returns the TargetingGroupExpression field value if set, zero value otherwise.
func (o *TargetingGroupBidRecommendation) GetTargetingGroupExpression() string {
	if o == nil || IsNil(o.TargetingGroupExpression) {
		var ret string
		return ret
	}
	return *o.TargetingGroupExpression
}

// GetTargetingGroupExpressionOk returns a tuple with the TargetingGroupExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetingGroupBidRecommendation) GetTargetingGroupExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.TargetingGroupExpression) {
		return nil, false
	}
	return o.TargetingGroupExpression, true
}

// HasTargetingGroupExpression returns a boolean if a field has been set.
func (o *TargetingGroupBidRecommendation) HasTargetingGroupExpression() bool {
	if o != nil && !IsNil(o.TargetingGroupExpression) {
		return true
	}

	return false
}

// SetTargetingGroupExpression gets a reference to the given string and assigns it to the TargetingGroupExpression field.
func (o *TargetingGroupBidRecommendation) SetTargetingGroupExpression(v string) {
	o.TargetingGroupExpression = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *TargetingGroupBidRecommendation) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetingGroupBidRecommendation) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *TargetingGroupBidRecommendation) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *TargetingGroupBidRecommendation) SetAction(v string) {
	o.Action = &v
}

// GetAdGroupId returns the AdGroupId field value if set, zero value otherwise.
func (o *TargetingGroupBidRecommendation) GetAdGroupId() string {
	if o == nil || IsNil(o.AdGroupId) {
		var ret string
		return ret
	}
	return *o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetingGroupBidRecommendation) GetAdGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.AdGroupId) {
		return nil, false
	}
	return o.AdGroupId, true
}

// HasAdGroupId returns a boolean if a field has been set.
func (o *TargetingGroupBidRecommendation) HasAdGroupId() bool {
	if o != nil && !IsNil(o.AdGroupId) {
		return true
	}

	return false
}

// SetAdGroupId gets a reference to the given string and assigns it to the AdGroupId field.
func (o *TargetingGroupBidRecommendation) SetAdGroupId(v string) {
	o.AdGroupId = &v
}

func (o TargetingGroupBidRecommendation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SuggestedBid) {
		toSerialize["suggestedBid"] = o.SuggestedBid
	}
	if !IsNil(o.TargetId) {
		toSerialize["targetId"] = o.TargetId
	}
	if !IsNil(o.TargetingGroupExpression) {
		toSerialize["targetingGroupExpression"] = o.TargetingGroupExpression
	}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.AdGroupId) {
		toSerialize["adGroupId"] = o.AdGroupId
	}
	return toSerialize, nil
}

type NullableTargetingGroupBidRecommendation struct {
	value *TargetingGroupBidRecommendation
	isSet bool
}

func (v NullableTargetingGroupBidRecommendation) Get() *TargetingGroupBidRecommendation {
	return v.value
}

func (v *NullableTargetingGroupBidRecommendation) Set(val *TargetingGroupBidRecommendation) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetingGroupBidRecommendation) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetingGroupBidRecommendation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetingGroupBidRecommendation(val *TargetingGroupBidRecommendation) *NullableTargetingGroupBidRecommendation {
	return &NullableTargetingGroupBidRecommendation{value: val, isSet: true}
}

func (v NullableTargetingGroupBidRecommendation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTargetingGroupBidRecommendation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
