package promotions_20251201

import (
	"github.com/bytedance/sonic"
)

// checks if the PurchaseRequirements type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PurchaseRequirements{}

// PurchaseRequirements Purchase requirements that customers must meet to qualify for a basket building promotion. Contains eligibility selection (what to buy), purchase condition (how much to buy), and an optional claim code. This is the 'Buy X' portion of 'Buy X Get Y' promotions.
type PurchaseRequirements struct {
	Selection Selection         `json:"selection"`
	Condition PurchaseCondition `json:"condition"`
	ClaimCode *ClaimCode        `json:"claimCode,omitempty"`
}

// NewPurchaseRequirements instantiates a new PurchaseRequirements object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPurchaseRequirements(selection Selection, condition PurchaseCondition) *PurchaseRequirements {
	this := PurchaseRequirements{}
	this.Selection = selection
	this.Condition = condition
	return &this
}

// NewPurchaseRequirementsWithDefaults instantiates a new PurchaseRequirements object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPurchaseRequirementsWithDefaults() *PurchaseRequirements {
	this := PurchaseRequirements{}
	return &this
}

// GetSelection returns the Selection field value
func (o *PurchaseRequirements) GetSelection() Selection {
	if o == nil {
		var ret Selection
		return ret
	}

	return o.Selection
}

// GetSelectionOk returns a tuple with the Selection field value
// and a boolean to check if the value has been set.
func (o *PurchaseRequirements) GetSelectionOk() (*Selection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Selection, true
}

// SetSelection sets field value
func (o *PurchaseRequirements) SetSelection(v Selection) {
	o.Selection = v
}

// GetCondition returns the Condition field value
func (o *PurchaseRequirements) GetCondition() PurchaseCondition {
	if o == nil {
		var ret PurchaseCondition
		return ret
	}

	return o.Condition
}

// GetConditionOk returns a tuple with the Condition field value
// and a boolean to check if the value has been set.
func (o *PurchaseRequirements) GetConditionOk() (*PurchaseCondition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Condition, true
}

// SetCondition sets field value
func (o *PurchaseRequirements) SetCondition(v PurchaseCondition) {
	o.Condition = v
}

// GetClaimCode returns the ClaimCode field value if set, zero value otherwise.
func (o *PurchaseRequirements) GetClaimCode() ClaimCode {
	if o == nil || IsNil(o.ClaimCode) {
		var ret ClaimCode
		return ret
	}
	return *o.ClaimCode
}

// GetClaimCodeOk returns a tuple with the ClaimCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PurchaseRequirements) GetClaimCodeOk() (*ClaimCode, bool) {
	if o == nil || IsNil(o.ClaimCode) {
		return nil, false
	}
	return o.ClaimCode, true
}

// HasClaimCode returns a boolean if a field has been set.
func (o *PurchaseRequirements) HasClaimCode() bool {
	if o != nil && !IsNil(o.ClaimCode) {
		return true
	}

	return false
}

// SetClaimCode gets a reference to the given ClaimCode and assigns it to the ClaimCode field.
func (o *PurchaseRequirements) SetClaimCode(v ClaimCode) {
	o.ClaimCode = &v
}

func (o PurchaseRequirements) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["selection"] = o.Selection
	toSerialize["condition"] = o.Condition
	if !IsNil(o.ClaimCode) {
		toSerialize["claimCode"] = o.ClaimCode
	}
	return toSerialize, nil
}

type NullablePurchaseRequirements struct {
	value *PurchaseRequirements
	isSet bool
}

func (v NullablePurchaseRequirements) Get() *PurchaseRequirements {
	return v.value
}

func (v *NullablePurchaseRequirements) Set(val *PurchaseRequirements) {
	v.value = val
	v.isSet = true
}

func (v NullablePurchaseRequirements) IsSet() bool {
	return v.isSet
}

func (v *NullablePurchaseRequirements) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePurchaseRequirements(val *PurchaseRequirements) *NullablePurchaseRequirements {
	return &NullablePurchaseRequirements{value: val, isSet: true}
}

func (v NullablePurchaseRequirements) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePurchaseRequirements) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
