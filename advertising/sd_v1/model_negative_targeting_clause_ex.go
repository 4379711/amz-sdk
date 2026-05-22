package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the NegativeTargetingClauseEx type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NegativeTargetingClauseEx{}

// NegativeTargetingClauseEx struct for NegativeTargetingClauseEx
type NegativeTargetingClauseEx struct {
	TargetId       *float32 `json:"targetId,omitempty"`
	AdGroupId      *float32 `json:"adGroupId,omitempty"`
	State          *string  `json:"state,omitempty"`
	ExpressionType *string  `json:"expressionType,omitempty"`
	// The expression to negatively match against. * Only one brand may be specified per targeting expression. * Only one asin may be specified per targeting expression. * To exclude a brand from a targeting expression, you must create a negative targeting expression in the same ad group as the positive targeting expression.
	Expression []NegativeTargetingClauseExExpressionInner `json:"expression,omitempty"`
	// The status of the target.
	ServingStatus *string `json:"servingStatus,omitempty"`
	// Epoch date the target was created.
	CreationDate *int64 `json:"creationDate,omitempty"`
	// Epoch date of the last update to any property associated with the target.
	LastUpdatedDate *int64 `json:"lastUpdatedDate,omitempty"`
}

// NewNegativeTargetingClauseEx instantiates a new NegativeTargetingClauseEx object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNegativeTargetingClauseEx() *NegativeTargetingClauseEx {
	this := NegativeTargetingClauseEx{}
	return &this
}

// NewNegativeTargetingClauseExWithDefaults instantiates a new NegativeTargetingClauseEx object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNegativeTargetingClauseExWithDefaults() *NegativeTargetingClauseEx {
	this := NegativeTargetingClauseEx{}
	return &this
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *NegativeTargetingClauseEx) GetTargetId() float32 {
	if o == nil || IsNil(o.TargetId) {
		var ret float32
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NegativeTargetingClauseEx) GetTargetIdOk() (*float32, bool) {
	if o == nil || IsNil(o.TargetId) {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *NegativeTargetingClauseEx) HasTargetId() bool {
	if o != nil && !IsNil(o.TargetId) {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given float32 and assigns it to the TargetId field.
func (o *NegativeTargetingClauseEx) SetTargetId(v float32) {
	o.TargetId = &v
}

// GetAdGroupId returns the AdGroupId field value if set, zero value otherwise.
func (o *NegativeTargetingClauseEx) GetAdGroupId() float32 {
	if o == nil || IsNil(o.AdGroupId) {
		var ret float32
		return ret
	}
	return *o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NegativeTargetingClauseEx) GetAdGroupIdOk() (*float32, bool) {
	if o == nil || IsNil(o.AdGroupId) {
		return nil, false
	}
	return o.AdGroupId, true
}

// HasAdGroupId returns a boolean if a field has been set.
func (o *NegativeTargetingClauseEx) HasAdGroupId() bool {
	if o != nil && !IsNil(o.AdGroupId) {
		return true
	}

	return false
}

// SetAdGroupId gets a reference to the given float32 and assigns it to the AdGroupId field.
func (o *NegativeTargetingClauseEx) SetAdGroupId(v float32) {
	o.AdGroupId = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *NegativeTargetingClauseEx) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NegativeTargetingClauseEx) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *NegativeTargetingClauseEx) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *NegativeTargetingClauseEx) SetState(v string) {
	o.State = &v
}

// GetExpressionType returns the ExpressionType field value if set, zero value otherwise.
func (o *NegativeTargetingClauseEx) GetExpressionType() string {
	if o == nil || IsNil(o.ExpressionType) {
		var ret string
		return ret
	}
	return *o.ExpressionType
}

// GetExpressionTypeOk returns a tuple with the ExpressionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NegativeTargetingClauseEx) GetExpressionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ExpressionType) {
		return nil, false
	}
	return o.ExpressionType, true
}

// HasExpressionType returns a boolean if a field has been set.
func (o *NegativeTargetingClauseEx) HasExpressionType() bool {
	if o != nil && !IsNil(o.ExpressionType) {
		return true
	}

	return false
}

// SetExpressionType gets a reference to the given string and assigns it to the ExpressionType field.
func (o *NegativeTargetingClauseEx) SetExpressionType(v string) {
	o.ExpressionType = &v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *NegativeTargetingClauseEx) GetExpression() []NegativeTargetingClauseExExpressionInner {
	if o == nil || IsNil(o.Expression) {
		var ret []NegativeTargetingClauseExExpressionInner
		return ret
	}
	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NegativeTargetingClauseEx) GetExpressionOk() ([]NegativeTargetingClauseExExpressionInner, bool) {
	if o == nil || IsNil(o.Expression) {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *NegativeTargetingClauseEx) HasExpression() bool {
	if o != nil && !IsNil(o.Expression) {
		return true
	}

	return false
}

// SetExpression gets a reference to the given []NegativeTargetingClauseExExpressionInner and assigns it to the Expression field.
func (o *NegativeTargetingClauseEx) SetExpression(v []NegativeTargetingClauseExExpressionInner) {
	o.Expression = v
}

// GetServingStatus returns the ServingStatus field value if set, zero value otherwise.
func (o *NegativeTargetingClauseEx) GetServingStatus() string {
	if o == nil || IsNil(o.ServingStatus) {
		var ret string
		return ret
	}
	return *o.ServingStatus
}

// GetServingStatusOk returns a tuple with the ServingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NegativeTargetingClauseEx) GetServingStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ServingStatus) {
		return nil, false
	}
	return o.ServingStatus, true
}

// HasServingStatus returns a boolean if a field has been set.
func (o *NegativeTargetingClauseEx) HasServingStatus() bool {
	if o != nil && !IsNil(o.ServingStatus) {
		return true
	}

	return false
}

// SetServingStatus gets a reference to the given string and assigns it to the ServingStatus field.
func (o *NegativeTargetingClauseEx) SetServingStatus(v string) {
	o.ServingStatus = &v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *NegativeTargetingClauseEx) GetCreationDate() int64 {
	if o == nil || IsNil(o.CreationDate) {
		var ret int64
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NegativeTargetingClauseEx) GetCreationDateOk() (*int64, bool) {
	if o == nil || IsNil(o.CreationDate) {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *NegativeTargetingClauseEx) HasCreationDate() bool {
	if o != nil && !IsNil(o.CreationDate) {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given int64 and assigns it to the CreationDate field.
func (o *NegativeTargetingClauseEx) SetCreationDate(v int64) {
	o.CreationDate = &v
}

// GetLastUpdatedDate returns the LastUpdatedDate field value if set, zero value otherwise.
func (o *NegativeTargetingClauseEx) GetLastUpdatedDate() int64 {
	if o == nil || IsNil(o.LastUpdatedDate) {
		var ret int64
		return ret
	}
	return *o.LastUpdatedDate
}

// GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NegativeTargetingClauseEx) GetLastUpdatedDateOk() (*int64, bool) {
	if o == nil || IsNil(o.LastUpdatedDate) {
		return nil, false
	}
	return o.LastUpdatedDate, true
}

// HasLastUpdatedDate returns a boolean if a field has been set.
func (o *NegativeTargetingClauseEx) HasLastUpdatedDate() bool {
	if o != nil && !IsNil(o.LastUpdatedDate) {
		return true
	}

	return false
}

// SetLastUpdatedDate gets a reference to the given int64 and assigns it to the LastUpdatedDate field.
func (o *NegativeTargetingClauseEx) SetLastUpdatedDate(v int64) {
	o.LastUpdatedDate = &v
}

func (o NegativeTargetingClauseEx) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TargetId) {
		toSerialize["targetId"] = o.TargetId
	}
	if !IsNil(o.AdGroupId) {
		toSerialize["adGroupId"] = o.AdGroupId
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.ExpressionType) {
		toSerialize["expressionType"] = o.ExpressionType
	}
	if !IsNil(o.Expression) {
		toSerialize["expression"] = o.Expression
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

type NullableNegativeTargetingClauseEx struct {
	value *NegativeTargetingClauseEx
	isSet bool
}

func (v NullableNegativeTargetingClauseEx) Get() *NegativeTargetingClauseEx {
	return v.value
}

func (v *NullableNegativeTargetingClauseEx) Set(val *NegativeTargetingClauseEx) {
	v.value = val
	v.isSet = true
}

func (v NullableNegativeTargetingClauseEx) IsSet() bool {
	return v.isSet
}

func (v *NullableNegativeTargetingClauseEx) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNegativeTargetingClauseEx(val *NegativeTargetingClauseEx) *NullableNegativeTargetingClauseEx {
	return &NullableNegativeTargetingClauseEx{value: val, isSet: true}
}

func (v NullableNegativeTargetingClauseEx) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableNegativeTargetingClauseEx) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
