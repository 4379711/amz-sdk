package vendor_invoices

import (
	"github.com/bytedance/sonic"
)

// checks if the AllowanceDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AllowanceDetails{}

// AllowanceDetails Monetary and tax details of the allowance.
type AllowanceDetails struct {
	// Type of the allowance applied.
	Type string `json:"type"`
	// Description of the allowance.
	Description     *string `json:"description,omitempty"`
	AllowanceAmount Money   `json:"allowanceAmount"`
	// Tax amount details applied on this allowance.
	TaxDetails []TaxDetails `json:"taxDetails,omitempty"`
}

type _AllowanceDetails AllowanceDetails

// NewAllowanceDetails instantiates a new AllowanceDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllowanceDetails(type_ string, allowanceAmount Money) *AllowanceDetails {
	this := AllowanceDetails{}
	this.Type = type_
	this.AllowanceAmount = allowanceAmount
	return &this
}

// NewAllowanceDetailsWithDefaults instantiates a new AllowanceDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllowanceDetailsWithDefaults() *AllowanceDetails {
	this := AllowanceDetails{}
	return &this
}

// GetType returns the Type field value
func (o *AllowanceDetails) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AllowanceDetails) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AllowanceDetails) SetType(v string) {
	o.Type = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AllowanceDetails) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowanceDetails) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AllowanceDetails) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AllowanceDetails) SetDescription(v string) {
	o.Description = &v
}

// GetAllowanceAmount returns the AllowanceAmount field value
func (o *AllowanceDetails) GetAllowanceAmount() Money {
	if o == nil {
		var ret Money
		return ret
	}

	return o.AllowanceAmount
}

// GetAllowanceAmountOk returns a tuple with the AllowanceAmount field value
// and a boolean to check if the value has been set.
func (o *AllowanceDetails) GetAllowanceAmountOk() (*Money, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowanceAmount, true
}

// SetAllowanceAmount sets field value
func (o *AllowanceDetails) SetAllowanceAmount(v Money) {
	o.AllowanceAmount = v
}

// GetTaxDetails returns the TaxDetails field value if set, zero value otherwise.
func (o *AllowanceDetails) GetTaxDetails() []TaxDetails {
	if o == nil || IsNil(o.TaxDetails) {
		var ret []TaxDetails
		return ret
	}
	return o.TaxDetails
}

// GetTaxDetailsOk returns a tuple with the TaxDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowanceDetails) GetTaxDetailsOk() ([]TaxDetails, bool) {
	if o == nil || IsNil(o.TaxDetails) {
		return nil, false
	}
	return o.TaxDetails, true
}

// HasTaxDetails returns a boolean if a field has been set.
func (o *AllowanceDetails) HasTaxDetails() bool {
	if o != nil && !IsNil(o.TaxDetails) {
		return true
	}

	return false
}

// SetTaxDetails gets a reference to the given []TaxDetails and assigns it to the TaxDetails field.
func (o *AllowanceDetails) SetTaxDetails(v []TaxDetails) {
	o.TaxDetails = v
}

func (o AllowanceDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["allowanceAmount"] = o.AllowanceAmount
	if !IsNil(o.TaxDetails) {
		toSerialize["taxDetails"] = o.TaxDetails
	}
	return toSerialize, nil
}

type NullableAllowanceDetails struct {
	value *AllowanceDetails
	isSet bool
}

func (v NullableAllowanceDetails) Get() *AllowanceDetails {
	return v.value
}

func (v *NullableAllowanceDetails) Set(val *AllowanceDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableAllowanceDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableAllowanceDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllowanceDetails(val *AllowanceDetails) *NullableAllowanceDetails {
	return &NullableAllowanceDetails{value: val, isSet: true}
}

func (v NullableAllowanceDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAllowanceDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
