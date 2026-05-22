package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the ContactInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContactInformation{}

// ContactInformation The seller's contact information.
type ContactInformation struct {
	// The email address.
	Email *string `json:"email,omitempty"`
	// The contact's name.
	Name string `json:"name"`
	// The phone number.
	PhoneNumber string `json:"phoneNumber"`
}

type _ContactInformation ContactInformation

// NewContactInformation instantiates a new ContactInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactInformation(name string, phoneNumber string) *ContactInformation {
	this := ContactInformation{}
	this.Name = name
	this.PhoneNumber = phoneNumber
	return &this
}

// NewContactInformationWithDefaults instantiates a new ContactInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactInformationWithDefaults() *ContactInformation {
	this := ContactInformation{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ContactInformation) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactInformation) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ContactInformation) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *ContactInformation) SetEmail(v string) {
	o.Email = &v
}

// GetName returns the Name field value
func (o *ContactInformation) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContactInformation) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContactInformation) SetName(v string) {
	o.Name = v
}

// GetPhoneNumber returns the PhoneNumber field value
func (o *ContactInformation) GetPhoneNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value
// and a boolean to check if the value has been set.
func (o *ContactInformation) GetPhoneNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PhoneNumber, true
}

// SetPhoneNumber sets field value
func (o *ContactInformation) SetPhoneNumber(v string) {
	o.PhoneNumber = v
}

func (o ContactInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	toSerialize["name"] = o.Name
	toSerialize["phoneNumber"] = o.PhoneNumber
	return toSerialize, nil
}

type NullableContactInformation struct {
	value *ContactInformation
	isSet bool
}

func (v NullableContactInformation) Get() *ContactInformation {
	return v.value
}

func (v *NullableContactInformation) Set(val *ContactInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableContactInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableContactInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactInformation(val *ContactInformation) *NullableContactInformation {
	return &NullableContactInformation{value: val, isSet: true}
}

func (v NullableContactInformation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableContactInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
