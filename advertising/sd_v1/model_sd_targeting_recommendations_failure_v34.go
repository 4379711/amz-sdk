package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDTargetingRecommendationsFailureV34 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDTargetingRecommendationsFailureV34{}

// SDTargetingRecommendationsFailureV34 A targeting recommendation failure record.
type SDTargetingRecommendationsFailureV34 struct {
	// HTTP status code indicating a failure response for targeting recomendations.
	Code *string `json:"code,omitempty"`
	// The theme name specified in the request. If the themes field is not provided in the request, the value of this field will be set to default.
	Name *string `json:"name,omitempty"`
	// A list of expressions that failed to be applied in the product targeting theme.
	Expression []SDProductTargetingThemeExpression `json:"expression,omitempty"`
	// A human friendly error message indicating the failure reasons.
	ErrorMessage *string `json:"errorMessage,omitempty"`
}

// NewSDTargetingRecommendationsFailureV34 instantiates a new SDTargetingRecommendationsFailureV34 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDTargetingRecommendationsFailureV34() *SDTargetingRecommendationsFailureV34 {
	this := SDTargetingRecommendationsFailureV34{}
	return &this
}

// NewSDTargetingRecommendationsFailureV34WithDefaults instantiates a new SDTargetingRecommendationsFailureV34 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDTargetingRecommendationsFailureV34WithDefaults() *SDTargetingRecommendationsFailureV34 {
	this := SDTargetingRecommendationsFailureV34{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SDTargetingRecommendationsFailureV34) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDTargetingRecommendationsFailureV34) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SDTargetingRecommendationsFailureV34) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *SDTargetingRecommendationsFailureV34) SetCode(v string) {
	o.Code = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SDTargetingRecommendationsFailureV34) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDTargetingRecommendationsFailureV34) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SDTargetingRecommendationsFailureV34) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SDTargetingRecommendationsFailureV34) SetName(v string) {
	o.Name = &v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *SDTargetingRecommendationsFailureV34) GetExpression() []SDProductTargetingThemeExpression {
	if o == nil || IsNil(o.Expression) {
		var ret []SDProductTargetingThemeExpression
		return ret
	}
	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDTargetingRecommendationsFailureV34) GetExpressionOk() ([]SDProductTargetingThemeExpression, bool) {
	if o == nil || IsNil(o.Expression) {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *SDTargetingRecommendationsFailureV34) HasExpression() bool {
	if o != nil && !IsNil(o.Expression) {
		return true
	}

	return false
}

// SetExpression gets a reference to the given []SDProductTargetingThemeExpression and assigns it to the Expression field.
func (o *SDTargetingRecommendationsFailureV34) SetExpression(v []SDProductTargetingThemeExpression) {
	o.Expression = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *SDTargetingRecommendationsFailureV34) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDTargetingRecommendationsFailureV34) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *SDTargetingRecommendationsFailureV34) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *SDTargetingRecommendationsFailureV34) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

func (o SDTargetingRecommendationsFailureV34) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Expression) {
		toSerialize["expression"] = o.Expression
	}
	if !IsNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	return toSerialize, nil
}

type NullableSDTargetingRecommendationsFailureV34 struct {
	value *SDTargetingRecommendationsFailureV34
	isSet bool
}

func (v NullableSDTargetingRecommendationsFailureV34) Get() *SDTargetingRecommendationsFailureV34 {
	return v.value
}

func (v *NullableSDTargetingRecommendationsFailureV34) Set(val *SDTargetingRecommendationsFailureV34) {
	v.value = val
	v.isSet = true
}

func (v NullableSDTargetingRecommendationsFailureV34) IsSet() bool {
	return v.isSet
}

func (v *NullableSDTargetingRecommendationsFailureV34) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDTargetingRecommendationsFailureV34(val *SDTargetingRecommendationsFailureV34) *NullableSDTargetingRecommendationsFailureV34 {
	return &NullableSDTargetingRecommendationsFailureV34{value: val, isSet: true}
}

func (v NullableSDTargetingRecommendationsFailureV34) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDTargetingRecommendationsFailureV34) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
