package merchant_fulfillment_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the Constraint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Constraint{}

// Constraint A validation constraint.
type Constraint struct {
	// A regular expression.
	ValidationRegEx *string `json:"ValidationRegEx,omitempty"`
	// A validation string.
	ValidationString string `json:"ValidationString"`
}

type _Constraint Constraint

// NewConstraint instantiates a new Constraint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConstraint(validationString string) *Constraint {
	this := Constraint{}
	this.ValidationString = validationString
	return &this
}

// NewConstraintWithDefaults instantiates a new Constraint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConstraintWithDefaults() *Constraint {
	this := Constraint{}
	return &this
}

// GetValidationRegEx returns the ValidationRegEx field value if set, zero value otherwise.
func (o *Constraint) GetValidationRegEx() string {
	if o == nil || IsNil(o.ValidationRegEx) {
		var ret string
		return ret
	}
	return *o.ValidationRegEx
}

// GetValidationRegExOk returns a tuple with the ValidationRegEx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Constraint) GetValidationRegExOk() (*string, bool) {
	if o == nil || IsNil(o.ValidationRegEx) {
		return nil, false
	}
	return o.ValidationRegEx, true
}

// HasValidationRegEx returns a boolean if a field has been set.
func (o *Constraint) HasValidationRegEx() bool {
	if o != nil && !IsNil(o.ValidationRegEx) {
		return true
	}

	return false
}

// SetValidationRegEx gets a reference to the given string and assigns it to the ValidationRegEx field.
func (o *Constraint) SetValidationRegEx(v string) {
	o.ValidationRegEx = &v
}

// GetValidationString returns the ValidationString field value
func (o *Constraint) GetValidationString() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ValidationString
}

// GetValidationStringOk returns a tuple with the ValidationString field value
// and a boolean to check if the value has been set.
func (o *Constraint) GetValidationStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidationString, true
}

// SetValidationString sets field value
func (o *Constraint) SetValidationString(v string) {
	o.ValidationString = v
}

func (o Constraint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ValidationRegEx) {
		toSerialize["ValidationRegEx"] = o.ValidationRegEx
	}
	toSerialize["ValidationString"] = o.ValidationString
	return toSerialize, nil
}

type NullableConstraint struct {
	value *Constraint
	isSet bool
}

func (v NullableConstraint) Get() *Constraint {
	return v.value
}

func (v *NullableConstraint) Set(val *Constraint) {
	v.value = val
	v.isSet = true
}

func (v NullableConstraint) IsSet() bool {
	return v.isSet
}

func (v *NullableConstraint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConstraint(val *Constraint) *NullableConstraint {
	return &NullableConstraint{value: val, isSet: true}
}

func (v NullableConstraint) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableConstraint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
