package seller_wallet_20240301

import (
	"github.com/bytedance/sonic"
)

// checks if the AccountHolderAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountHolderAddress{}

// AccountHolderAddress The Address used to verify the bank account of the payee. This can be a person or business mailing address.
type AccountHolderAddress struct {
	// Address Line 1 of the public address.
	AddressLine1 string `json:"addressLine1"`
	// Address Line 2 of the public address.
	AddressLine2 *string `json:"addressLine2,omitempty"`
	// City name of the public address.
	City string `json:"city"`
	// State name of the public address. This will be state or region for CN (China) based addresses.
	State string `json:"state"`
	// Postal code of the public address.
	PostalCode string `json:"postalCode"`
	// Country name of the public address.
	Country *string `json:"country,omitempty"`
	// The two digit country code, in ISO 3166 format.
	CountryCode string `json:"countryCode"`
}

type _AccountHolderAddress AccountHolderAddress

// NewAccountHolderAddress instantiates a new AccountHolderAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountHolderAddress(addressLine1 string, city string, state string, postalCode string, countryCode string) *AccountHolderAddress {
	this := AccountHolderAddress{}
	this.AddressLine1 = addressLine1
	this.City = city
	this.State = state
	this.PostalCode = postalCode
	this.CountryCode = countryCode
	return &this
}

// NewAccountHolderAddressWithDefaults instantiates a new AccountHolderAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountHolderAddressWithDefaults() *AccountHolderAddress {
	this := AccountHolderAddress{}
	return &this
}

// GetAddressLine1 returns the AddressLine1 field value
func (o *AccountHolderAddress) GetAddressLine1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddressLine1
}

// GetAddressLine1Ok returns a tuple with the AddressLine1 field value
// and a boolean to check if the value has been set.
func (o *AccountHolderAddress) GetAddressLine1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressLine1, true
}

// SetAddressLine1 sets field value
func (o *AccountHolderAddress) SetAddressLine1(v string) {
	o.AddressLine1 = v
}

// GetAddressLine2 returns the AddressLine2 field value if set, zero value otherwise.
func (o *AccountHolderAddress) GetAddressLine2() string {
	if o == nil || IsNil(o.AddressLine2) {
		var ret string
		return ret
	}
	return *o.AddressLine2
}

// GetAddressLine2Ok returns a tuple with the AddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderAddress) GetAddressLine2Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine2) {
		return nil, false
	}
	return o.AddressLine2, true
}

// HasAddressLine2 returns a boolean if a field has been set.
func (o *AccountHolderAddress) HasAddressLine2() bool {
	if o != nil && !IsNil(o.AddressLine2) {
		return true
	}

	return false
}

// SetAddressLine2 gets a reference to the given string and assigns it to the AddressLine2 field.
func (o *AccountHolderAddress) SetAddressLine2(v string) {
	o.AddressLine2 = &v
}

// GetCity returns the City field value
func (o *AccountHolderAddress) GetCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *AccountHolderAddress) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *AccountHolderAddress) SetCity(v string) {
	o.City = v
}

// GetState returns the State field value
func (o *AccountHolderAddress) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *AccountHolderAddress) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *AccountHolderAddress) SetState(v string) {
	o.State = v
}

// GetPostalCode returns the PostalCode field value
func (o *AccountHolderAddress) GetPostalCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value
// and a boolean to check if the value has been set.
func (o *AccountHolderAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostalCode, true
}

// SetPostalCode sets field value
func (o *AccountHolderAddress) SetPostalCode(v string) {
	o.PostalCode = v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *AccountHolderAddress) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderAddress) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *AccountHolderAddress) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *AccountHolderAddress) SetCountry(v string) {
	o.Country = &v
}

// GetCountryCode returns the CountryCode field value
func (o *AccountHolderAddress) GetCountryCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value
// and a boolean to check if the value has been set.
func (o *AccountHolderAddress) GetCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryCode, true
}

// SetCountryCode sets field value
func (o *AccountHolderAddress) SetCountryCode(v string) {
	o.CountryCode = v
}

func (o AccountHolderAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["addressLine1"] = o.AddressLine1
	if !IsNil(o.AddressLine2) {
		toSerialize["addressLine2"] = o.AddressLine2
	}
	toSerialize["city"] = o.City
	toSerialize["state"] = o.State
	toSerialize["postalCode"] = o.PostalCode
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	toSerialize["countryCode"] = o.CountryCode
	return toSerialize, nil
}

type NullableAccountHolderAddress struct {
	value *AccountHolderAddress
	isSet bool
}

func (v NullableAccountHolderAddress) Get() *AccountHolderAddress {
	return v.value
}

func (v *NullableAccountHolderAddress) Set(val *AccountHolderAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountHolderAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountHolderAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountHolderAddress(val *AccountHolderAddress) *NullableAccountHolderAddress {
	return &NullableAccountHolderAddress{value: val, isSet: true}
}

func (v NullableAccountHolderAddress) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAccountHolderAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
