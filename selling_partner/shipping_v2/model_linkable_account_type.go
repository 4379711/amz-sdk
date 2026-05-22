package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the LinkableAccountType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinkableAccountType{}

// LinkableAccountType Info About Linkable Account Type
type LinkableAccountType struct {
	AccountType *AccountType `json:"accountType,omitempty"`
	// A list of CarrierAccountInput
	CarrierAccountInputs []CarrierAccountInput `json:"carrierAccountInputs,omitempty"`
}

// NewLinkableAccountType instantiates a new LinkableAccountType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkableAccountType() *LinkableAccountType {
	this := LinkableAccountType{}
	return &this
}

// NewLinkableAccountTypeWithDefaults instantiates a new LinkableAccountType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkableAccountTypeWithDefaults() *LinkableAccountType {
	this := LinkableAccountType{}
	return &this
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *LinkableAccountType) GetAccountType() AccountType {
	if o == nil || IsNil(o.AccountType) {
		var ret AccountType
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkableAccountType) GetAccountTypeOk() (*AccountType, bool) {
	if o == nil || IsNil(o.AccountType) {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *LinkableAccountType) HasAccountType() bool {
	if o != nil && !IsNil(o.AccountType) {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given AccountType and assigns it to the AccountType field.
func (o *LinkableAccountType) SetAccountType(v AccountType) {
	o.AccountType = &v
}

// GetCarrierAccountInputs returns the CarrierAccountInputs field value if set, zero value otherwise.
func (o *LinkableAccountType) GetCarrierAccountInputs() []CarrierAccountInput {
	if o == nil || IsNil(o.CarrierAccountInputs) {
		var ret []CarrierAccountInput
		return ret
	}
	return o.CarrierAccountInputs
}

// GetCarrierAccountInputsOk returns a tuple with the CarrierAccountInputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkableAccountType) GetCarrierAccountInputsOk() ([]CarrierAccountInput, bool) {
	if o == nil || IsNil(o.CarrierAccountInputs) {
		return nil, false
	}
	return o.CarrierAccountInputs, true
}

// HasCarrierAccountInputs returns a boolean if a field has been set.
func (o *LinkableAccountType) HasCarrierAccountInputs() bool {
	if o != nil && !IsNil(o.CarrierAccountInputs) {
		return true
	}

	return false
}

// SetCarrierAccountInputs gets a reference to the given []CarrierAccountInput and assigns it to the CarrierAccountInputs field.
func (o *LinkableAccountType) SetCarrierAccountInputs(v []CarrierAccountInput) {
	o.CarrierAccountInputs = v
}

func (o LinkableAccountType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountType) {
		toSerialize["accountType"] = o.AccountType
	}
	if !IsNil(o.CarrierAccountInputs) {
		toSerialize["carrierAccountInputs"] = o.CarrierAccountInputs
	}
	return toSerialize, nil
}

type NullableLinkableAccountType struct {
	value *LinkableAccountType
	isSet bool
}

func (v NullableLinkableAccountType) Get() *LinkableAccountType {
	return v.value
}

func (v *NullableLinkableAccountType) Set(val *LinkableAccountType) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkableAccountType) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkableAccountType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkableAccountType(val *LinkableAccountType) *NullableLinkableAccountType {
	return &NullableLinkableAccountType{value: val, isSet: true}
}

func (v NullableLinkableAccountType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableLinkableAccountType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
