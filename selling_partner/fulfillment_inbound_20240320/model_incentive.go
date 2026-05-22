package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the Incentive type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Incentive{}

// Incentive Contains details about cost related modifications to the placement cost.
type Incentive struct {
	// Description of the incentive.
	Description string `json:"description"`
	// Target of the incentive. Possible values: 'Placement Services', 'Fulfillment Fee Discount'.
	Target string `json:"target"`
	// Type of incentive. Possible values: `FEE`, `DISCOUNT`.
	Type  string   `json:"type"`
	Value Currency `json:"value"`
}

type _Incentive Incentive

// NewIncentive instantiates a new Incentive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncentive(description string, target string, type_ string, value Currency) *Incentive {
	this := Incentive{}
	this.Description = description
	this.Target = target
	this.Type = type_
	this.Value = value
	return &this
}

// NewIncentiveWithDefaults instantiates a new Incentive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncentiveWithDefaults() *Incentive {
	this := Incentive{}
	return &this
}

// GetDescription returns the Description field value
func (o *Incentive) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Incentive) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Incentive) SetDescription(v string) {
	o.Description = v
}

// GetTarget returns the Target field value
func (o *Incentive) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *Incentive) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *Incentive) SetTarget(v string) {
	o.Target = v
}

// GetType returns the Type field value
func (o *Incentive) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Incentive) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Incentive) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *Incentive) GetValue() Currency {
	if o == nil {
		var ret Currency
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Incentive) GetValueOk() (*Currency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Incentive) SetValue(v Currency) {
	o.Value = v
}

func (o Incentive) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["target"] = o.Target
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableIncentive struct {
	value *Incentive
	isSet bool
}

func (v NullableIncentive) Get() *Incentive {
	return v.value
}

func (v *NullableIncentive) Set(val *Incentive) {
	v.value = val
	v.isSet = true
}

func (v NullableIncentive) IsSet() bool {
	return v.isSet
}

func (v *NullableIncentive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncentive(val *Incentive) *NullableIncentive {
	return &NullableIncentive{value: val, isSet: true}
}

func (v NullableIncentive) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableIncentive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
