package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the OrderTaxRegistration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderTaxRegistration{}

// OrderTaxRegistration Tax registration information for an entity associated with the order.
type OrderTaxRegistration struct {
	// The type of entity that the tax registration belongs to.  **Possible values**: - `BUYER` (Indicates that this is the buyer's tax registration information) - `MERCHANT` (Indicates that this is the merchant's tax registration information) - `MARKETPLACE` (Indicates that this is the marketplace's tax registration information)
	EntityType *string `json:"entityType,omitempty"`
	// The legal name associated with the tax registration.
	LegalName *string `json:"legalName,omitempty"`
	// The type of the tax registration number.  **Possible values**: `BUSINESS`, `VAT`, `CST`, `CPF`, `CNPJ`
	TaxRegistrationType *string `json:"taxRegistrationType,omitempty"`
	// The tax registration number that identifies the entity for tax purposes.
	TaxRegistrationNumber  *string          `json:"taxRegistrationNumber,omitempty"`
	TaxRegistrationAddress *CustomerAddress `json:"taxRegistrationAddress,omitempty"`
	// Additional attributes related to the tax registration.
	TaxRegistrationAttributes []TaxRegistrationAttribute `json:"taxRegistrationAttributes,omitempty"`
}

// NewOrderTaxRegistration instantiates a new OrderTaxRegistration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderTaxRegistration() *OrderTaxRegistration {
	this := OrderTaxRegistration{}
	return &this
}

// NewOrderTaxRegistrationWithDefaults instantiates a new OrderTaxRegistration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderTaxRegistrationWithDefaults() *OrderTaxRegistration {
	this := OrderTaxRegistration{}
	return &this
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *OrderTaxRegistration) GetEntityType() string {
	if o == nil || IsNil(o.EntityType) {
		var ret string
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTaxRegistration) GetEntityTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EntityType) {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *OrderTaxRegistration) HasEntityType() bool {
	if o != nil && !IsNil(o.EntityType) {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given string and assigns it to the EntityType field.
func (o *OrderTaxRegistration) SetEntityType(v string) {
	o.EntityType = &v
}

// GetLegalName returns the LegalName field value if set, zero value otherwise.
func (o *OrderTaxRegistration) GetLegalName() string {
	if o == nil || IsNil(o.LegalName) {
		var ret string
		return ret
	}
	return *o.LegalName
}

// GetLegalNameOk returns a tuple with the LegalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTaxRegistration) GetLegalNameOk() (*string, bool) {
	if o == nil || IsNil(o.LegalName) {
		return nil, false
	}
	return o.LegalName, true
}

// HasLegalName returns a boolean if a field has been set.
func (o *OrderTaxRegistration) HasLegalName() bool {
	if o != nil && !IsNil(o.LegalName) {
		return true
	}

	return false
}

// SetLegalName gets a reference to the given string and assigns it to the LegalName field.
func (o *OrderTaxRegistration) SetLegalName(v string) {
	o.LegalName = &v
}

// GetTaxRegistrationType returns the TaxRegistrationType field value if set, zero value otherwise.
func (o *OrderTaxRegistration) GetTaxRegistrationType() string {
	if o == nil || IsNil(o.TaxRegistrationType) {
		var ret string
		return ret
	}
	return *o.TaxRegistrationType
}

// GetTaxRegistrationTypeOk returns a tuple with the TaxRegistrationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTaxRegistration) GetTaxRegistrationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TaxRegistrationType) {
		return nil, false
	}
	return o.TaxRegistrationType, true
}

// HasTaxRegistrationType returns a boolean if a field has been set.
func (o *OrderTaxRegistration) HasTaxRegistrationType() bool {
	if o != nil && !IsNil(o.TaxRegistrationType) {
		return true
	}

	return false
}

// SetTaxRegistrationType gets a reference to the given string and assigns it to the TaxRegistrationType field.
func (o *OrderTaxRegistration) SetTaxRegistrationType(v string) {
	o.TaxRegistrationType = &v
}

// GetTaxRegistrationNumber returns the TaxRegistrationNumber field value if set, zero value otherwise.
func (o *OrderTaxRegistration) GetTaxRegistrationNumber() string {
	if o == nil || IsNil(o.TaxRegistrationNumber) {
		var ret string
		return ret
	}
	return *o.TaxRegistrationNumber
}

// GetTaxRegistrationNumberOk returns a tuple with the TaxRegistrationNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTaxRegistration) GetTaxRegistrationNumberOk() (*string, bool) {
	if o == nil || IsNil(o.TaxRegistrationNumber) {
		return nil, false
	}
	return o.TaxRegistrationNumber, true
}

// HasTaxRegistrationNumber returns a boolean if a field has been set.
func (o *OrderTaxRegistration) HasTaxRegistrationNumber() bool {
	if o != nil && !IsNil(o.TaxRegistrationNumber) {
		return true
	}

	return false
}

// SetTaxRegistrationNumber gets a reference to the given string and assigns it to the TaxRegistrationNumber field.
func (o *OrderTaxRegistration) SetTaxRegistrationNumber(v string) {
	o.TaxRegistrationNumber = &v
}

// GetTaxRegistrationAddress returns the TaxRegistrationAddress field value if set, zero value otherwise.
func (o *OrderTaxRegistration) GetTaxRegistrationAddress() CustomerAddress {
	if o == nil || IsNil(o.TaxRegistrationAddress) {
		var ret CustomerAddress
		return ret
	}
	return *o.TaxRegistrationAddress
}

// GetTaxRegistrationAddressOk returns a tuple with the TaxRegistrationAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTaxRegistration) GetTaxRegistrationAddressOk() (*CustomerAddress, bool) {
	if o == nil || IsNil(o.TaxRegistrationAddress) {
		return nil, false
	}
	return o.TaxRegistrationAddress, true
}

// HasTaxRegistrationAddress returns a boolean if a field has been set.
func (o *OrderTaxRegistration) HasTaxRegistrationAddress() bool {
	if o != nil && !IsNil(o.TaxRegistrationAddress) {
		return true
	}

	return false
}

// SetTaxRegistrationAddress gets a reference to the given CustomerAddress and assigns it to the TaxRegistrationAddress field.
func (o *OrderTaxRegistration) SetTaxRegistrationAddress(v CustomerAddress) {
	o.TaxRegistrationAddress = &v
}

// GetTaxRegistrationAttributes returns the TaxRegistrationAttributes field value if set, zero value otherwise.
func (o *OrderTaxRegistration) GetTaxRegistrationAttributes() []TaxRegistrationAttribute {
	if o == nil || IsNil(o.TaxRegistrationAttributes) {
		var ret []TaxRegistrationAttribute
		return ret
	}
	return o.TaxRegistrationAttributes
}

// GetTaxRegistrationAttributesOk returns a tuple with the TaxRegistrationAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTaxRegistration) GetTaxRegistrationAttributesOk() ([]TaxRegistrationAttribute, bool) {
	if o == nil || IsNil(o.TaxRegistrationAttributes) {
		return nil, false
	}
	return o.TaxRegistrationAttributes, true
}

// HasTaxRegistrationAttributes returns a boolean if a field has been set.
func (o *OrderTaxRegistration) HasTaxRegistrationAttributes() bool {
	if o != nil && !IsNil(o.TaxRegistrationAttributes) {
		return true
	}

	return false
}

// SetTaxRegistrationAttributes gets a reference to the given []TaxRegistrationAttribute and assigns it to the TaxRegistrationAttributes field.
func (o *OrderTaxRegistration) SetTaxRegistrationAttributes(v []TaxRegistrationAttribute) {
	o.TaxRegistrationAttributes = v
}

func (o OrderTaxRegistration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EntityType) {
		toSerialize["entityType"] = o.EntityType
	}
	if !IsNil(o.LegalName) {
		toSerialize["legalName"] = o.LegalName
	}
	if !IsNil(o.TaxRegistrationType) {
		toSerialize["taxRegistrationType"] = o.TaxRegistrationType
	}
	if !IsNil(o.TaxRegistrationNumber) {
		toSerialize["taxRegistrationNumber"] = o.TaxRegistrationNumber
	}
	if !IsNil(o.TaxRegistrationAddress) {
		toSerialize["taxRegistrationAddress"] = o.TaxRegistrationAddress
	}
	if !IsNil(o.TaxRegistrationAttributes) {
		toSerialize["taxRegistrationAttributes"] = o.TaxRegistrationAttributes
	}
	return toSerialize, nil
}

type NullableOrderTaxRegistration struct {
	value *OrderTaxRegistration
	isSet bool
}

func (v NullableOrderTaxRegistration) Get() *OrderTaxRegistration {
	return v.value
}

func (v *NullableOrderTaxRegistration) Set(val *OrderTaxRegistration) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderTaxRegistration) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderTaxRegistration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderTaxRegistration(val *OrderTaxRegistration) *NullableOrderTaxRegistration {
	return &NullableOrderTaxRegistration{value: val, isSet: true}
}

func (v NullableOrderTaxRegistration) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOrderTaxRegistration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
