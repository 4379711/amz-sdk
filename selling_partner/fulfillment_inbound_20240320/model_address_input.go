package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the AddressInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressInput{}

// AddressInput Specific details to identify a place.
type AddressInput struct {
	// Street address information.
	AddressLine1 string `json:"addressLine1"`
	// Additional street address information.
	AddressLine2 *string `json:"addressLine2,omitempty"`
	// The city.
	City string `json:"city"`
	// The name of the business.
	CompanyName *string `json:"companyName,omitempty"`
	// The country code in two-character ISO 3166-1 alpha-2 format.
	CountryCode string `json:"countryCode" validate:"regexp=^[A-Z]{2}$"`
	// The email address.
	Email *string `json:"email,omitempty"`
	// The name of the individual who is the primary contact.
	Name string `json:"name"`
	// The phone number.
	PhoneNumber string `json:"phoneNumber"`
	// The postal code.
	PostalCode string `json:"postalCode"`
	// The state or province code.
	StateOrProvinceCode *string `json:"stateOrProvinceCode,omitempty"`
}

type _AddressInput AddressInput

// NewAddressInput instantiates a new AddressInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressInput(addressLine1 string, city string, countryCode string, name string, phoneNumber string, postalCode string) *AddressInput {
	this := AddressInput{}
	this.AddressLine1 = addressLine1
	this.City = city
	this.CountryCode = countryCode
	this.Name = name
	this.PhoneNumber = phoneNumber
	this.PostalCode = postalCode
	return &this
}

// NewAddressInputWithDefaults instantiates a new AddressInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressInputWithDefaults() *AddressInput {
	this := AddressInput{}
	return &this
}

// GetAddressLine1 returns the AddressLine1 field value
func (o *AddressInput) GetAddressLine1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddressLine1
}

// GetAddressLine1Ok returns a tuple with the AddressLine1 field value
// and a boolean to check if the value has been set.
func (o *AddressInput) GetAddressLine1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressLine1, true
}

// SetAddressLine1 sets field value
func (o *AddressInput) SetAddressLine1(v string) {
	o.AddressLine1 = v
}

// GetAddressLine2 returns the AddressLine2 field value if set, zero value otherwise.
func (o *AddressInput) GetAddressLine2() string {
	if o == nil || IsNil(o.AddressLine2) {
		var ret string
		return ret
	}
	return *o.AddressLine2
}

// GetAddressLine2Ok returns a tuple with the AddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressInput) GetAddressLine2Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine2) {
		return nil, false
	}
	return o.AddressLine2, true
}

// HasAddressLine2 returns a boolean if a field has been set.
func (o *AddressInput) HasAddressLine2() bool {
	if o != nil && !IsNil(o.AddressLine2) {
		return true
	}

	return false
}

// SetAddressLine2 gets a reference to the given string and assigns it to the AddressLine2 field.
func (o *AddressInput) SetAddressLine2(v string) {
	o.AddressLine2 = &v
}

// GetCity returns the City field value
func (o *AddressInput) GetCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *AddressInput) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *AddressInput) SetCity(v string) {
	o.City = v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *AddressInput) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressInput) GetCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyName) {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *AddressInput) HasCompanyName() bool {
	if o != nil && !IsNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *AddressInput) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetCountryCode returns the CountryCode field value
func (o *AddressInput) GetCountryCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value
// and a boolean to check if the value has been set.
func (o *AddressInput) GetCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryCode, true
}

// SetCountryCode sets field value
func (o *AddressInput) SetCountryCode(v string) {
	o.CountryCode = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *AddressInput) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressInput) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *AddressInput) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *AddressInput) SetEmail(v string) {
	o.Email = &v
}

// GetName returns the Name field value
func (o *AddressInput) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AddressInput) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AddressInput) SetName(v string) {
	o.Name = v
}

// GetPhoneNumber returns the PhoneNumber field value
func (o *AddressInput) GetPhoneNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value
// and a boolean to check if the value has been set.
func (o *AddressInput) GetPhoneNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PhoneNumber, true
}

// SetPhoneNumber sets field value
func (o *AddressInput) SetPhoneNumber(v string) {
	o.PhoneNumber = v
}

// GetPostalCode returns the PostalCode field value
func (o *AddressInput) GetPostalCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value
// and a boolean to check if the value has been set.
func (o *AddressInput) GetPostalCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostalCode, true
}

// SetPostalCode sets field value
func (o *AddressInput) SetPostalCode(v string) {
	o.PostalCode = v
}

// GetStateOrProvinceCode returns the StateOrProvinceCode field value if set, zero value otherwise.
func (o *AddressInput) GetStateOrProvinceCode() string {
	if o == nil || IsNil(o.StateOrProvinceCode) {
		var ret string
		return ret
	}
	return *o.StateOrProvinceCode
}

// GetStateOrProvinceCodeOk returns a tuple with the StateOrProvinceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressInput) GetStateOrProvinceCodeOk() (*string, bool) {
	if o == nil || IsNil(o.StateOrProvinceCode) {
		return nil, false
	}
	return o.StateOrProvinceCode, true
}

// HasStateOrProvinceCode returns a boolean if a field has been set.
func (o *AddressInput) HasStateOrProvinceCode() bool {
	if o != nil && !IsNil(o.StateOrProvinceCode) {
		return true
	}

	return false
}

// SetStateOrProvinceCode gets a reference to the given string and assigns it to the StateOrProvinceCode field.
func (o *AddressInput) SetStateOrProvinceCode(v string) {
	o.StateOrProvinceCode = &v
}

func (o AddressInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["addressLine1"] = o.AddressLine1
	if !IsNil(o.AddressLine2) {
		toSerialize["addressLine2"] = o.AddressLine2
	}
	toSerialize["city"] = o.City
	if !IsNil(o.CompanyName) {
		toSerialize["companyName"] = o.CompanyName
	}
	toSerialize["countryCode"] = o.CountryCode
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	toSerialize["name"] = o.Name
	toSerialize["phoneNumber"] = o.PhoneNumber
	toSerialize["postalCode"] = o.PostalCode
	if !IsNil(o.StateOrProvinceCode) {
		toSerialize["stateOrProvinceCode"] = o.StateOrProvinceCode
	}
	return toSerialize, nil
}

type NullableAddressInput struct {
	value *AddressInput
	isSet bool
}

func (v NullableAddressInput) Get() *AddressInput {
	return v.value
}

func (v *NullableAddressInput) Set(val *AddressInput) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressInput) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressInput(val *AddressInput) *NullableAddressInput {
	return &NullableAddressInput{value: val, isSet: true}
}

func (v NullableAddressInput) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAddressInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
