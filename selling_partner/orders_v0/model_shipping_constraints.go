package orders_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the ShippingConstraints type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShippingConstraints{}

// ShippingConstraints Delivery constraints applicable to this order.
type ShippingConstraints struct {
	PalletDelivery                *ConstraintType `json:"PalletDelivery,omitempty"`
	SignatureConfirmation         *ConstraintType `json:"SignatureConfirmation,omitempty"`
	RecipientIdentityVerification *ConstraintType `json:"RecipientIdentityVerification,omitempty"`
	RecipientAgeVerification      *ConstraintType `json:"RecipientAgeVerification,omitempty"`
}

// NewShippingConstraints instantiates a new ShippingConstraints object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingConstraints() *ShippingConstraints {
	this := ShippingConstraints{}
	return &this
}

// NewShippingConstraintsWithDefaults instantiates a new ShippingConstraints object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingConstraintsWithDefaults() *ShippingConstraints {
	this := ShippingConstraints{}
	return &this
}

// GetPalletDelivery returns the PalletDelivery field value if set, zero value otherwise.
func (o *ShippingConstraints) GetPalletDelivery() ConstraintType {
	if o == nil || IsNil(o.PalletDelivery) {
		var ret ConstraintType
		return ret
	}
	return *o.PalletDelivery
}

// GetPalletDeliveryOk returns a tuple with the PalletDelivery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingConstraints) GetPalletDeliveryOk() (*ConstraintType, bool) {
	if o == nil || IsNil(o.PalletDelivery) {
		return nil, false
	}
	return o.PalletDelivery, true
}

// HasPalletDelivery returns a boolean if a field has been set.
func (o *ShippingConstraints) HasPalletDelivery() bool {
	if o != nil && !IsNil(o.PalletDelivery) {
		return true
	}

	return false
}

// SetPalletDelivery gets a reference to the given ConstraintType and assigns it to the PalletDelivery field.
func (o *ShippingConstraints) SetPalletDelivery(v ConstraintType) {
	o.PalletDelivery = &v
}

// GetSignatureConfirmation returns the SignatureConfirmation field value if set, zero value otherwise.
func (o *ShippingConstraints) GetSignatureConfirmation() ConstraintType {
	if o == nil || IsNil(o.SignatureConfirmation) {
		var ret ConstraintType
		return ret
	}
	return *o.SignatureConfirmation
}

// GetSignatureConfirmationOk returns a tuple with the SignatureConfirmation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingConstraints) GetSignatureConfirmationOk() (*ConstraintType, bool) {
	if o == nil || IsNil(o.SignatureConfirmation) {
		return nil, false
	}
	return o.SignatureConfirmation, true
}

// HasSignatureConfirmation returns a boolean if a field has been set.
func (o *ShippingConstraints) HasSignatureConfirmation() bool {
	if o != nil && !IsNil(o.SignatureConfirmation) {
		return true
	}

	return false
}

// SetSignatureConfirmation gets a reference to the given ConstraintType and assigns it to the SignatureConfirmation field.
func (o *ShippingConstraints) SetSignatureConfirmation(v ConstraintType) {
	o.SignatureConfirmation = &v
}

// GetRecipientIdentityVerification returns the RecipientIdentityVerification field value if set, zero value otherwise.
func (o *ShippingConstraints) GetRecipientIdentityVerification() ConstraintType {
	if o == nil || IsNil(o.RecipientIdentityVerification) {
		var ret ConstraintType
		return ret
	}
	return *o.RecipientIdentityVerification
}

// GetRecipientIdentityVerificationOk returns a tuple with the RecipientIdentityVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingConstraints) GetRecipientIdentityVerificationOk() (*ConstraintType, bool) {
	if o == nil || IsNil(o.RecipientIdentityVerification) {
		return nil, false
	}
	return o.RecipientIdentityVerification, true
}

// HasRecipientIdentityVerification returns a boolean if a field has been set.
func (o *ShippingConstraints) HasRecipientIdentityVerification() bool {
	if o != nil && !IsNil(o.RecipientIdentityVerification) {
		return true
	}

	return false
}

// SetRecipientIdentityVerification gets a reference to the given ConstraintType and assigns it to the RecipientIdentityVerification field.
func (o *ShippingConstraints) SetRecipientIdentityVerification(v ConstraintType) {
	o.RecipientIdentityVerification = &v
}

// GetRecipientAgeVerification returns the RecipientAgeVerification field value if set, zero value otherwise.
func (o *ShippingConstraints) GetRecipientAgeVerification() ConstraintType {
	if o == nil || IsNil(o.RecipientAgeVerification) {
		var ret ConstraintType
		return ret
	}
	return *o.RecipientAgeVerification
}

// GetRecipientAgeVerificationOk returns a tuple with the RecipientAgeVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingConstraints) GetRecipientAgeVerificationOk() (*ConstraintType, bool) {
	if o == nil || IsNil(o.RecipientAgeVerification) {
		return nil, false
	}
	return o.RecipientAgeVerification, true
}

// HasRecipientAgeVerification returns a boolean if a field has been set.
func (o *ShippingConstraints) HasRecipientAgeVerification() bool {
	if o != nil && !IsNil(o.RecipientAgeVerification) {
		return true
	}

	return false
}

// SetRecipientAgeVerification gets a reference to the given ConstraintType and assigns it to the RecipientAgeVerification field.
func (o *ShippingConstraints) SetRecipientAgeVerification(v ConstraintType) {
	o.RecipientAgeVerification = &v
}

func (o ShippingConstraints) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PalletDelivery) {
		toSerialize["PalletDelivery"] = o.PalletDelivery
	}
	if !IsNil(o.SignatureConfirmation) {
		toSerialize["SignatureConfirmation"] = o.SignatureConfirmation
	}
	if !IsNil(o.RecipientIdentityVerification) {
		toSerialize["RecipientIdentityVerification"] = o.RecipientIdentityVerification
	}
	if !IsNil(o.RecipientAgeVerification) {
		toSerialize["RecipientAgeVerification"] = o.RecipientAgeVerification
	}
	return toSerialize, nil
}

type NullableShippingConstraints struct {
	value *ShippingConstraints
	isSet bool
}

func (v NullableShippingConstraints) Get() *ShippingConstraints {
	return v.value
}

func (v *NullableShippingConstraints) Set(val *ShippingConstraints) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingConstraints) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingConstraints) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingConstraints(val *ShippingConstraints) *NullableShippingConstraints {
	return &NullableShippingConstraints{value: val, isSet: true}
}

func (v NullableShippingConstraints) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableShippingConstraints) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
