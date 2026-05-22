package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the TargetingClauseEx type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TargetingClauseEx{}

// TargetingClauseEx struct for TargetingClauseEx
type TargetingClauseEx struct {
	TargetId       *float32 `json:"targetId,omitempty"`
	AdGroupId      *float32 `json:"adGroupId,omitempty"`
	CampaignId     *float32 `json:"campaignId,omitempty"`
	State          *string  `json:"state,omitempty"`
	ExpressionType *string  `json:"expressionType,omitempty"`
	// If a value for `bid` is specified, it overrides the current adGroup bid. When using vcpm costType. $1 is the minimum bid for vCPM. Note that this field is ignored for negative targeting clauses.
	Bid *float32 `json:"bid,omitempty"`
	// The targeting expression to match against.  ------- Applicable to contextual or content targeting (T00020) ------- * A 'TargetingExpression' in a contextual targeting campaign can contain 'TargetingPredicate' or 'ContentTargetingPredicate' components. * Contextual expressions must specify either a category predicate or an ASIN predicate, but never both. * Only one category may be specified per targeting expression. * Only one brand may be specified per targeting expression. * Only one asin may be specified per targeting expression. * To exclude a brand from a targeting expression you must create a negative targeting expression in the same ad group as the positive targeting expression.  ------- Applicable to audiences or contextual targeting (T00030) ------- * A 'TargetingExpression' in a audiences or contextual campaign can contain any target, including 'TargetingPredicate', 'ContentTargetingPredicate', or 'TargetingPredicateNested'.
	Expression []TargetingExpressionInner `json:"expression,omitempty"`
	// The targeting expression to match against.  ------- Applicable to contextual or content targeting (T00020) ------- * A 'TargetingExpression' in a contextual targeting campaign can contain 'TargetingPredicate' or 'ContentTargetingPredicate' components. * Contextual expressions must specify either a category predicate or an ASIN predicate, but never both. * Only one category may be specified per targeting expression. * Only one brand may be specified per targeting expression. * Only one asin may be specified per targeting expression. * To exclude a brand from a targeting expression you must create a negative targeting expression in the same ad group as the positive targeting expression.  ------- Applicable to audiences or contextual targeting (T00030) ------- * A 'TargetingExpression' in a audiences or contextual campaign can contain any target, including 'TargetingPredicate', 'ContentTargetingPredicate', or 'TargetingPredicateNested'.
	ResolvedExpression []TargetingExpressionInner `json:"resolvedExpression,omitempty"`
	// The status of the target.
	ServingStatus *string `json:"servingStatus,omitempty"`
	// Epoch date the target was created.
	CreationDate *int64 `json:"creationDate,omitempty"`
	// Epoch date of the last update to any property associated with the target.
	LastUpdatedDate *int64 `json:"lastUpdatedDate,omitempty"`
}

// NewTargetingClauseEx instantiates a new TargetingClauseEx object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetingClauseEx() *TargetingClauseEx {
	this := TargetingClauseEx{}
	return &this
}

// NewTargetingClauseExWithDefaults instantiates a new TargetingClauseEx object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetingClauseExWithDefaults() *TargetingClauseEx {
	this := TargetingClauseEx{}
	return &this
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *TargetingClauseEx) GetTargetId() float32 {
	if o == nil || IsNil(o.TargetId) {
		var ret float32
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetingClauseEx) GetTargetIdOk() (*float32, bool) {
	if o == nil || IsNil(o.TargetId) {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *TargetingClauseEx) HasTargetId() bool {
	if o != nil && !IsNil(o.TargetId) {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given float32 and assigns it to the TargetId field.
func (o *TargetingClauseEx) SetTargetId(v float32) {
	o.TargetId = &v
}

// GetAdGroupId returns the AdGroupId field value if set, zero value otherwise.
func (o *TargetingClauseEx) GetAdGroupId() float32 {
	if o == nil || IsNil(o.AdGroupId) {
		var ret float32
		return ret
	}
	return *o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetingClauseEx) GetAdGroupIdOk() (*float32, bool) {
	if o == nil || IsNil(o.AdGroupId) {
		return nil, false
	}
	return o.AdGroupId, true
}

// HasAdGroupId returns a boolean if a field has been set.
func (o *TargetingClauseEx) HasAdGroupId() bool {
	if o != nil && !IsNil(o.AdGroupId) {
		return true
	}

	return false
}

// SetAdGroupId gets a reference to the given float32 and assigns it to the AdGroupId field.
func (o *TargetingClauseEx) SetAdGroupId(v float32) {
	o.AdGroupId = &v
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *TargetingClauseEx) GetCampaignId() float32 {
	if o == nil || IsNil(o.CampaignId) {
		var ret float32
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetingClauseEx) GetCampaignIdOk() (*float32, bool) {
	if o == nil || IsNil(o.CampaignId) {
		return nil, false
	}
	return o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *TargetingClauseEx) HasCampaignId() bool {
	if o != nil && !IsNil(o.CampaignId) {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given float32 and assigns it to the CampaignId field.
func (o *TargetingClauseEx) SetCampaignId(v float32) {
	o.CampaignId = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *TargetingClauseEx) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetingClauseEx) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *TargetingClauseEx) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *TargetingClauseEx) SetState(v string) {
	o.State = &v
}

// GetExpressionType returns the ExpressionType field value if set, zero value otherwise.
func (o *TargetingClauseEx) GetExpressionType() string {
	if o == nil || IsNil(o.ExpressionType) {
		var ret string
		return ret
	}
	return *o.ExpressionType
}

// GetExpressionTypeOk returns a tuple with the ExpressionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetingClauseEx) GetExpressionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ExpressionType) {
		return nil, false
	}
	return o.ExpressionType, true
}

// HasExpressionType returns a boolean if a field has been set.
func (o *TargetingClauseEx) HasExpressionType() bool {
	if o != nil && !IsNil(o.ExpressionType) {
		return true
	}

	return false
}

// SetExpressionType gets a reference to the given string and assigns it to the ExpressionType field.
func (o *TargetingClauseEx) SetExpressionType(v string) {
	o.ExpressionType = &v
}

// GetBid returns the Bid field value if set, zero value otherwise.
func (o *TargetingClauseEx) GetBid() float32 {
	if o == nil || IsNil(o.Bid) {
		var ret float32
		return ret
	}
	return *o.Bid
}

// GetBidOk returns a tuple with the Bid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetingClauseEx) GetBidOk() (*float32, bool) {
	if o == nil || IsNil(o.Bid) {
		return nil, false
	}
	return o.Bid, true
}

// HasBid returns a boolean if a field has been set.
func (o *TargetingClauseEx) HasBid() bool {
	if o != nil && !IsNil(o.Bid) {
		return true
	}

	return false
}

// SetBid gets a reference to the given float32 and assigns it to the Bid field.
func (o *TargetingClauseEx) SetBid(v float32) {
	o.Bid = &v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *TargetingClauseEx) GetExpression() []TargetingExpressionInner {
	if o == nil || IsNil(o.Expression) {
		var ret []TargetingExpressionInner
		return ret
	}
	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetingClauseEx) GetExpressionOk() ([]TargetingExpressionInner, bool) {
	if o == nil || IsNil(o.Expression) {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *TargetingClauseEx) HasExpression() bool {
	if o != nil && !IsNil(o.Expression) {
		return true
	}

	return false
}

// SetExpression gets a reference to the given []TargetingExpressionInner and assigns it to the Expression field.
func (o *TargetingClauseEx) SetExpression(v []TargetingExpressionInner) {
	o.Expression = v
}

// GetResolvedExpression returns the ResolvedExpression field value if set, zero value otherwise.
func (o *TargetingClauseEx) GetResolvedExpression() []TargetingExpressionInner {
	if o == nil || IsNil(o.ResolvedExpression) {
		var ret []TargetingExpressionInner
		return ret
	}
	return o.ResolvedExpression
}

// GetResolvedExpressionOk returns a tuple with the ResolvedExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetingClauseEx) GetResolvedExpressionOk() ([]TargetingExpressionInner, bool) {
	if o == nil || IsNil(o.ResolvedExpression) {
		return nil, false
	}
	return o.ResolvedExpression, true
}

// HasResolvedExpression returns a boolean if a field has been set.
func (o *TargetingClauseEx) HasResolvedExpression() bool {
	if o != nil && !IsNil(o.ResolvedExpression) {
		return true
	}

	return false
}

// SetResolvedExpression gets a reference to the given []TargetingExpressionInner and assigns it to the ResolvedExpression field.
func (o *TargetingClauseEx) SetResolvedExpression(v []TargetingExpressionInner) {
	o.ResolvedExpression = v
}

// GetServingStatus returns the ServingStatus field value if set, zero value otherwise.
func (o *TargetingClauseEx) GetServingStatus() string {
	if o == nil || IsNil(o.ServingStatus) {
		var ret string
		return ret
	}
	return *o.ServingStatus
}

// GetServingStatusOk returns a tuple with the ServingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetingClauseEx) GetServingStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ServingStatus) {
		return nil, false
	}
	return o.ServingStatus, true
}

// HasServingStatus returns a boolean if a field has been set.
func (o *TargetingClauseEx) HasServingStatus() bool {
	if o != nil && !IsNil(o.ServingStatus) {
		return true
	}

	return false
}

// SetServingStatus gets a reference to the given string and assigns it to the ServingStatus field.
func (o *TargetingClauseEx) SetServingStatus(v string) {
	o.ServingStatus = &v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *TargetingClauseEx) GetCreationDate() int64 {
	if o == nil || IsNil(o.CreationDate) {
		var ret int64
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetingClauseEx) GetCreationDateOk() (*int64, bool) {
	if o == nil || IsNil(o.CreationDate) {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *TargetingClauseEx) HasCreationDate() bool {
	if o != nil && !IsNil(o.CreationDate) {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given int64 and assigns it to the CreationDate field.
func (o *TargetingClauseEx) SetCreationDate(v int64) {
	o.CreationDate = &v
}

// GetLastUpdatedDate returns the LastUpdatedDate field value if set, zero value otherwise.
func (o *TargetingClauseEx) GetLastUpdatedDate() int64 {
	if o == nil || IsNil(o.LastUpdatedDate) {
		var ret int64
		return ret
	}
	return *o.LastUpdatedDate
}

// GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetingClauseEx) GetLastUpdatedDateOk() (*int64, bool) {
	if o == nil || IsNil(o.LastUpdatedDate) {
		return nil, false
	}
	return o.LastUpdatedDate, true
}

// HasLastUpdatedDate returns a boolean if a field has been set.
func (o *TargetingClauseEx) HasLastUpdatedDate() bool {
	if o != nil && !IsNil(o.LastUpdatedDate) {
		return true
	}

	return false
}

// SetLastUpdatedDate gets a reference to the given int64 and assigns it to the LastUpdatedDate field.
func (o *TargetingClauseEx) SetLastUpdatedDate(v int64) {
	o.LastUpdatedDate = &v
}

func (o TargetingClauseEx) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TargetId) {
		toSerialize["targetId"] = o.TargetId
	}
	if !IsNil(o.AdGroupId) {
		toSerialize["adGroupId"] = o.AdGroupId
	}
	if !IsNil(o.CampaignId) {
		toSerialize["campaignId"] = o.CampaignId
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.ExpressionType) {
		toSerialize["expressionType"] = o.ExpressionType
	}
	if !IsNil(o.Bid) {
		toSerialize["bid"] = o.Bid
	}
	if !IsNil(o.Expression) {
		toSerialize["expression"] = o.Expression
	}
	if !IsNil(o.ResolvedExpression) {
		toSerialize["resolvedExpression"] = o.ResolvedExpression
	}
	if !IsNil(o.ServingStatus) {
		toSerialize["servingStatus"] = o.ServingStatus
	}
	if !IsNil(o.CreationDate) {
		toSerialize["creationDate"] = o.CreationDate
	}
	if !IsNil(o.LastUpdatedDate) {
		toSerialize["lastUpdatedDate"] = o.LastUpdatedDate
	}
	return toSerialize, nil
}

type NullableTargetingClauseEx struct {
	value *TargetingClauseEx
	isSet bool
}

func (v NullableTargetingClauseEx) Get() *TargetingClauseEx {
	return v.value
}

func (v *NullableTargetingClauseEx) Set(val *TargetingClauseEx) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetingClauseEx) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetingClauseEx) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetingClauseEx(val *TargetingClauseEx) *NullableTargetingClauseEx {
	return &NullableTargetingClauseEx{value: val, isSet: true}
}

func (v NullableTargetingClauseEx) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTargetingClauseEx) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
