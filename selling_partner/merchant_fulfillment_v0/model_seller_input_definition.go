package merchant_fulfillment_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the SellerInputDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SellerInputDefinition{}

// SellerInputDefinition Specifies characteristics that apply to a seller input.
type SellerInputDefinition struct {
	// When true, the additional input field is required.
	IsRequired bool `json:"IsRequired"`
	// The data type of the additional input field.
	DataType string `json:"DataType"`
	// List of constraints.
	Constraints []Constraint `json:"Constraints"`
	// The display text for the additional input field.
	InputDisplayText string                `json:"InputDisplayText"`
	InputTarget      *InputTargetType      `json:"InputTarget,omitempty"`
	StoredValue      AdditionalSellerInput `json:"StoredValue"`
	// The set of fixed values in an additional seller input.
	RestrictedSetValues []string `json:"RestrictedSetValues,omitempty"`
}

type _SellerInputDefinition SellerInputDefinition

// NewSellerInputDefinition instantiates a new SellerInputDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSellerInputDefinition(isRequired bool, dataType string, constraints []Constraint, inputDisplayText string, storedValue AdditionalSellerInput) *SellerInputDefinition {
	this := SellerInputDefinition{}
	this.IsRequired = isRequired
	this.DataType = dataType
	this.Constraints = constraints
	this.InputDisplayText = inputDisplayText
	this.StoredValue = storedValue
	return &this
}

// NewSellerInputDefinitionWithDefaults instantiates a new SellerInputDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSellerInputDefinitionWithDefaults() *SellerInputDefinition {
	this := SellerInputDefinition{}
	return &this
}

// GetIsRequired returns the IsRequired field value
func (o *SellerInputDefinition) GetIsRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsRequired
}

// GetIsRequiredOk returns a tuple with the IsRequired field value
// and a boolean to check if the value has been set.
func (o *SellerInputDefinition) GetIsRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsRequired, true
}

// SetIsRequired sets field value
func (o *SellerInputDefinition) SetIsRequired(v bool) {
	o.IsRequired = v
}

// GetDataType returns the DataType field value
func (o *SellerInputDefinition) GetDataType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value
// and a boolean to check if the value has been set.
func (o *SellerInputDefinition) GetDataTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataType, true
}

// SetDataType sets field value
func (o *SellerInputDefinition) SetDataType(v string) {
	o.DataType = v
}

// GetConstraints returns the Constraints field value
func (o *SellerInputDefinition) GetConstraints() []Constraint {
	if o == nil {
		var ret []Constraint
		return ret
	}

	return o.Constraints
}

// GetConstraintsOk returns a tuple with the Constraints field value
// and a boolean to check if the value has been set.
func (o *SellerInputDefinition) GetConstraintsOk() ([]Constraint, bool) {
	if o == nil {
		return nil, false
	}
	return o.Constraints, true
}

// SetConstraints sets field value
func (o *SellerInputDefinition) SetConstraints(v []Constraint) {
	o.Constraints = v
}

// GetInputDisplayText returns the InputDisplayText field value
func (o *SellerInputDefinition) GetInputDisplayText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InputDisplayText
}

// GetInputDisplayTextOk returns a tuple with the InputDisplayText field value
// and a boolean to check if the value has been set.
func (o *SellerInputDefinition) GetInputDisplayTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InputDisplayText, true
}

// SetInputDisplayText sets field value
func (o *SellerInputDefinition) SetInputDisplayText(v string) {
	o.InputDisplayText = v
}

// GetInputTarget returns the InputTarget field value if set, zero value otherwise.
func (o *SellerInputDefinition) GetInputTarget() InputTargetType {
	if o == nil || IsNil(o.InputTarget) {
		var ret InputTargetType
		return ret
	}
	return *o.InputTarget
}

// GetInputTargetOk returns a tuple with the InputTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SellerInputDefinition) GetInputTargetOk() (*InputTargetType, bool) {
	if o == nil || IsNil(o.InputTarget) {
		return nil, false
	}
	return o.InputTarget, true
}

// HasInputTarget returns a boolean if a field has been set.
func (o *SellerInputDefinition) HasInputTarget() bool {
	if o != nil && !IsNil(o.InputTarget) {
		return true
	}

	return false
}

// SetInputTarget gets a reference to the given InputTargetType and assigns it to the InputTarget field.
func (o *SellerInputDefinition) SetInputTarget(v InputTargetType) {
	o.InputTarget = &v
}

// GetStoredValue returns the StoredValue field value
func (o *SellerInputDefinition) GetStoredValue() AdditionalSellerInput {
	if o == nil {
		var ret AdditionalSellerInput
		return ret
	}

	return o.StoredValue
}

// GetStoredValueOk returns a tuple with the StoredValue field value
// and a boolean to check if the value has been set.
func (o *SellerInputDefinition) GetStoredValueOk() (*AdditionalSellerInput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StoredValue, true
}

// SetStoredValue sets field value
func (o *SellerInputDefinition) SetStoredValue(v AdditionalSellerInput) {
	o.StoredValue = v
}

// GetRestrictedSetValues returns the RestrictedSetValues field value if set, zero value otherwise.
func (o *SellerInputDefinition) GetRestrictedSetValues() []string {
	if o == nil || IsNil(o.RestrictedSetValues) {
		var ret []string
		return ret
	}
	return o.RestrictedSetValues
}

// GetRestrictedSetValuesOk returns a tuple with the RestrictedSetValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SellerInputDefinition) GetRestrictedSetValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.RestrictedSetValues) {
		return nil, false
	}
	return o.RestrictedSetValues, true
}

// HasRestrictedSetValues returns a boolean if a field has been set.
func (o *SellerInputDefinition) HasRestrictedSetValues() bool {
	if o != nil && !IsNil(o.RestrictedSetValues) {
		return true
	}

	return false
}

// SetRestrictedSetValues gets a reference to the given []string and assigns it to the RestrictedSetValues field.
func (o *SellerInputDefinition) SetRestrictedSetValues(v []string) {
	o.RestrictedSetValues = v
}

func (o SellerInputDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["IsRequired"] = o.IsRequired
	toSerialize["DataType"] = o.DataType
	toSerialize["Constraints"] = o.Constraints
	toSerialize["InputDisplayText"] = o.InputDisplayText
	if !IsNil(o.InputTarget) {
		toSerialize["InputTarget"] = o.InputTarget
	}
	toSerialize["StoredValue"] = o.StoredValue
	if !IsNil(o.RestrictedSetValues) {
		toSerialize["RestrictedSetValues"] = o.RestrictedSetValues
	}
	return toSerialize, nil
}

type NullableSellerInputDefinition struct {
	value *SellerInputDefinition
	isSet bool
}

func (v NullableSellerInputDefinition) Get() *SellerInputDefinition {
	return v.value
}

func (v *NullableSellerInputDefinition) Set(val *SellerInputDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableSellerInputDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableSellerInputDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSellerInputDefinition(val *SellerInputDefinition) *NullableSellerInputDefinition {
	return &NullableSellerInputDefinition{value: val, isSet: true}
}

func (v NullableSellerInputDefinition) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSellerInputDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
