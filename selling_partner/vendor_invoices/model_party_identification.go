package vendor_invoices

import (
	"github.com/bytedance/sonic"
)

// checks if the PartyIdentification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PartyIdentification{}

// PartyIdentification Name, address and tax details of a party.
type PartyIdentification struct {
	// Assigned identification for the party.
	PartyId string   `json:"partyId"`
	Address *Address `json:"address,omitempty"`
	// Tax registration details of the party.
	TaxRegistrationDetails []TaxRegistrationDetails `json:"taxRegistrationDetails,omitempty"`
}

type _PartyIdentification PartyIdentification

// NewPartyIdentification instantiates a new PartyIdentification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartyIdentification(partyId string) *PartyIdentification {
	this := PartyIdentification{}
	this.PartyId = partyId
	return &this
}

// NewPartyIdentificationWithDefaults instantiates a new PartyIdentification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartyIdentificationWithDefaults() *PartyIdentification {
	this := PartyIdentification{}
	return &this
}

// GetPartyId returns the PartyId field value
func (o *PartyIdentification) GetPartyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PartyId
}

// GetPartyIdOk returns a tuple with the PartyId field value
// and a boolean to check if the value has been set.
func (o *PartyIdentification) GetPartyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PartyId, true
}

// SetPartyId sets field value
func (o *PartyIdentification) SetPartyId(v string) {
	o.PartyId = v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *PartyIdentification) GetAddress() Address {
	if o == nil || IsNil(o.Address) {
		var ret Address
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyIdentification) GetAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *PartyIdentification) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address and assigns it to the Address field.
func (o *PartyIdentification) SetAddress(v Address) {
	o.Address = &v
}

// GetTaxRegistrationDetails returns the TaxRegistrationDetails field value if set, zero value otherwise.
func (o *PartyIdentification) GetTaxRegistrationDetails() []TaxRegistrationDetails {
	if o == nil || IsNil(o.TaxRegistrationDetails) {
		var ret []TaxRegistrationDetails
		return ret
	}
	return o.TaxRegistrationDetails
}

// GetTaxRegistrationDetailsOk returns a tuple with the TaxRegistrationDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyIdentification) GetTaxRegistrationDetailsOk() ([]TaxRegistrationDetails, bool) {
	if o == nil || IsNil(o.TaxRegistrationDetails) {
		return nil, false
	}
	return o.TaxRegistrationDetails, true
}

// HasTaxRegistrationDetails returns a boolean if a field has been set.
func (o *PartyIdentification) HasTaxRegistrationDetails() bool {
	if o != nil && !IsNil(o.TaxRegistrationDetails) {
		return true
	}

	return false
}

// SetTaxRegistrationDetails gets a reference to the given []TaxRegistrationDetails and assigns it to the TaxRegistrationDetails field.
func (o *PartyIdentification) SetTaxRegistrationDetails(v []TaxRegistrationDetails) {
	o.TaxRegistrationDetails = v
}

func (o PartyIdentification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["partyId"] = o.PartyId
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.TaxRegistrationDetails) {
		toSerialize["taxRegistrationDetails"] = o.TaxRegistrationDetails
	}
	return toSerialize, nil
}

type NullablePartyIdentification struct {
	value *PartyIdentification
	isSet bool
}

func (v NullablePartyIdentification) Get() *PartyIdentification {
	return v.value
}

func (v *NullablePartyIdentification) Set(val *PartyIdentification) {
	v.value = val
	v.isSet = true
}

func (v NullablePartyIdentification) IsSet() bool {
	return v.isSet
}

func (v *NullablePartyIdentification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartyIdentification(val *PartyIdentification) *NullablePartyIdentification {
	return &NullablePartyIdentification{value: val, isSet: true}
}

func (v NullablePartyIdentification) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePartyIdentification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
