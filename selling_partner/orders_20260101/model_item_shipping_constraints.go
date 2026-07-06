package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemShippingConstraints type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemShippingConstraints{}

// ItemShippingConstraints Special shipping requirements and restrictions that must be observed when shipping an order item.
type ItemShippingConstraints struct {
	PalletDelivery                *ConstraintType `json:"palletDelivery,omitempty"`
	CashOnDelivery                *ConstraintType `json:"cashOnDelivery,omitempty"`
	SignatureConfirmation         *ConstraintType `json:"signatureConfirmation,omitempty"`
	RecipientIdentityVerification *ConstraintType `json:"recipientIdentityVerification,omitempty"`
	RecipientAgeVerification      *ConstraintType `json:"recipientAgeVerification,omitempty"`
}

// NewItemShippingConstraints instantiates a new ItemShippingConstraints object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemShippingConstraints() *ItemShippingConstraints {
	this := ItemShippingConstraints{}
	return &this
}

// NewItemShippingConstraintsWithDefaults instantiates a new ItemShippingConstraints object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemShippingConstraintsWithDefaults() *ItemShippingConstraints {
	this := ItemShippingConstraints{}
	return &this
}

// GetPalletDelivery returns the PalletDelivery field value if set, zero value otherwise.
func (o *ItemShippingConstraints) GetPalletDelivery() ConstraintType {
	if o == nil || IsNil(o.PalletDelivery) {
		var ret ConstraintType
		return ret
	}
	return *o.PalletDelivery
}

// GetPalletDeliveryOk returns a tuple with the PalletDelivery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemShippingConstraints) GetPalletDeliveryOk() (*ConstraintType, bool) {
	if o == nil || IsNil(o.PalletDelivery) {
		return nil, false
	}
	return o.PalletDelivery, true
}

// HasPalletDelivery returns a boolean if a field has been set.
func (o *ItemShippingConstraints) HasPalletDelivery() bool {
	if o != nil && !IsNil(o.PalletDelivery) {
		return true
	}

	return false
}

// SetPalletDelivery gets a reference to the given ConstraintType and assigns it to the PalletDelivery field.
func (o *ItemShippingConstraints) SetPalletDelivery(v ConstraintType) {
	o.PalletDelivery = &v
}

// GetCashOnDelivery returns the CashOnDelivery field value if set, zero value otherwise.
func (o *ItemShippingConstraints) GetCashOnDelivery() ConstraintType {
	if o == nil || IsNil(o.CashOnDelivery) {
		var ret ConstraintType
		return ret
	}
	return *o.CashOnDelivery
}

// GetCashOnDeliveryOk returns a tuple with the CashOnDelivery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemShippingConstraints) GetCashOnDeliveryOk() (*ConstraintType, bool) {
	if o == nil || IsNil(o.CashOnDelivery) {
		return nil, false
	}
	return o.CashOnDelivery, true
}

// HasCashOnDelivery returns a boolean if a field has been set.
func (o *ItemShippingConstraints) HasCashOnDelivery() bool {
	if o != nil && !IsNil(o.CashOnDelivery) {
		return true
	}

	return false
}

// SetCashOnDelivery gets a reference to the given ConstraintType and assigns it to the CashOnDelivery field.
func (o *ItemShippingConstraints) SetCashOnDelivery(v ConstraintType) {
	o.CashOnDelivery = &v
}

// GetSignatureConfirmation returns the SignatureConfirmation field value if set, zero value otherwise.
func (o *ItemShippingConstraints) GetSignatureConfirmation() ConstraintType {
	if o == nil || IsNil(o.SignatureConfirmation) {
		var ret ConstraintType
		return ret
	}
	return *o.SignatureConfirmation
}

// GetSignatureConfirmationOk returns a tuple with the SignatureConfirmation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemShippingConstraints) GetSignatureConfirmationOk() (*ConstraintType, bool) {
	if o == nil || IsNil(o.SignatureConfirmation) {
		return nil, false
	}
	return o.SignatureConfirmation, true
}

// HasSignatureConfirmation returns a boolean if a field has been set.
func (o *ItemShippingConstraints) HasSignatureConfirmation() bool {
	if o != nil && !IsNil(o.SignatureConfirmation) {
		return true
	}

	return false
}

// SetSignatureConfirmation gets a reference to the given ConstraintType and assigns it to the SignatureConfirmation field.
func (o *ItemShippingConstraints) SetSignatureConfirmation(v ConstraintType) {
	o.SignatureConfirmation = &v
}

// GetRecipientIdentityVerification returns the RecipientIdentityVerification field value if set, zero value otherwise.
func (o *ItemShippingConstraints) GetRecipientIdentityVerification() ConstraintType {
	if o == nil || IsNil(o.RecipientIdentityVerification) {
		var ret ConstraintType
		return ret
	}
	return *o.RecipientIdentityVerification
}

// GetRecipientIdentityVerificationOk returns a tuple with the RecipientIdentityVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemShippingConstraints) GetRecipientIdentityVerificationOk() (*ConstraintType, bool) {
	if o == nil || IsNil(o.RecipientIdentityVerification) {
		return nil, false
	}
	return o.RecipientIdentityVerification, true
}

// HasRecipientIdentityVerification returns a boolean if a field has been set.
func (o *ItemShippingConstraints) HasRecipientIdentityVerification() bool {
	if o != nil && !IsNil(o.RecipientIdentityVerification) {
		return true
	}

	return false
}

// SetRecipientIdentityVerification gets a reference to the given ConstraintType and assigns it to the RecipientIdentityVerification field.
func (o *ItemShippingConstraints) SetRecipientIdentityVerification(v ConstraintType) {
	o.RecipientIdentityVerification = &v
}

// GetRecipientAgeVerification returns the RecipientAgeVerification field value if set, zero value otherwise.
func (o *ItemShippingConstraints) GetRecipientAgeVerification() ConstraintType {
	if o == nil || IsNil(o.RecipientAgeVerification) {
		var ret ConstraintType
		return ret
	}
	return *o.RecipientAgeVerification
}

// GetRecipientAgeVerificationOk returns a tuple with the RecipientAgeVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemShippingConstraints) GetRecipientAgeVerificationOk() (*ConstraintType, bool) {
	if o == nil || IsNil(o.RecipientAgeVerification) {
		return nil, false
	}
	return o.RecipientAgeVerification, true
}

// HasRecipientAgeVerification returns a boolean if a field has been set.
func (o *ItemShippingConstraints) HasRecipientAgeVerification() bool {
	if o != nil && !IsNil(o.RecipientAgeVerification) {
		return true
	}

	return false
}

// SetRecipientAgeVerification gets a reference to the given ConstraintType and assigns it to the RecipientAgeVerification field.
func (o *ItemShippingConstraints) SetRecipientAgeVerification(v ConstraintType) {
	o.RecipientAgeVerification = &v
}

func (o ItemShippingConstraints) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PalletDelivery) {
		toSerialize["palletDelivery"] = o.PalletDelivery
	}
	if !IsNil(o.CashOnDelivery) {
		toSerialize["cashOnDelivery"] = o.CashOnDelivery
	}
	if !IsNil(o.SignatureConfirmation) {
		toSerialize["signatureConfirmation"] = o.SignatureConfirmation
	}
	if !IsNil(o.RecipientIdentityVerification) {
		toSerialize["recipientIdentityVerification"] = o.RecipientIdentityVerification
	}
	if !IsNil(o.RecipientAgeVerification) {
		toSerialize["recipientAgeVerification"] = o.RecipientAgeVerification
	}
	return toSerialize, nil
}

type NullableItemShippingConstraints struct {
	value *ItemShippingConstraints
	isSet bool
}

func (v NullableItemShippingConstraints) Get() *ItemShippingConstraints {
	return v.value
}

func (v *NullableItemShippingConstraints) Set(val *ItemShippingConstraints) {
	v.value = val
	v.isSet = true
}

func (v NullableItemShippingConstraints) IsSet() bool {
	return v.isSet
}

func (v *NullableItemShippingConstraints) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemShippingConstraints(val *ItemShippingConstraints) *NullableItemShippingConstraints {
	return &NullableItemShippingConstraints{value: val, isSet: true}
}

func (v NullableItemShippingConstraints) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemShippingConstraints) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
