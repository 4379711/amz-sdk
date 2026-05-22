package sellers

import (
	"github.com/bytedance/sonic"
)

// checks if the PrimaryContact type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrimaryContact{}

// PrimaryContact Information about the seller's primary contact.
type PrimaryContact struct {
	// The full name of the seller's primary contact.
	Name    string  `json:"name"`
	Address Address `json:"address"`
	// The non-Latin script version of the primary contact's name, if applicable.
	NonLatinName *string `json:"nonLatinName,omitempty"`
}

type _PrimaryContact PrimaryContact

// NewPrimaryContact instantiates a new PrimaryContact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrimaryContact(name string, address Address) *PrimaryContact {
	this := PrimaryContact{}
	this.Name = name
	this.Address = address
	return &this
}

// NewPrimaryContactWithDefaults instantiates a new PrimaryContact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrimaryContactWithDefaults() *PrimaryContact {
	this := PrimaryContact{}
	return &this
}

// GetName returns the Name field value
func (o *PrimaryContact) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PrimaryContact) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PrimaryContact) SetName(v string) {
	o.Name = v
}

// GetAddress returns the Address field value
func (o *PrimaryContact) GetAddress() Address {
	if o == nil {
		var ret Address
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *PrimaryContact) GetAddressOk() (*Address, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *PrimaryContact) SetAddress(v Address) {
	o.Address = v
}

// GetNonLatinName returns the NonLatinName field value if set, zero value otherwise.
func (o *PrimaryContact) GetNonLatinName() string {
	if o == nil || IsNil(o.NonLatinName) {
		var ret string
		return ret
	}
	return *o.NonLatinName
}

// GetNonLatinNameOk returns a tuple with the NonLatinName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrimaryContact) GetNonLatinNameOk() (*string, bool) {
	if o == nil || IsNil(o.NonLatinName) {
		return nil, false
	}
	return o.NonLatinName, true
}

// HasNonLatinName returns a boolean if a field has been set.
func (o *PrimaryContact) HasNonLatinName() bool {
	if o != nil && !IsNil(o.NonLatinName) {
		return true
	}

	return false
}

// SetNonLatinName gets a reference to the given string and assigns it to the NonLatinName field.
func (o *PrimaryContact) SetNonLatinName(v string) {
	o.NonLatinName = &v
}

func (o PrimaryContact) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["address"] = o.Address
	if !IsNil(o.NonLatinName) {
		toSerialize["nonLatinName"] = o.NonLatinName
	}
	return toSerialize, nil
}

type NullablePrimaryContact struct {
	value *PrimaryContact
	isSet bool
}

func (v NullablePrimaryContact) Get() *PrimaryContact {
	return v.value
}

func (v *NullablePrimaryContact) Set(val *PrimaryContact) {
	v.value = val
	v.isSet = true
}

func (v NullablePrimaryContact) IsSet() bool {
	return v.isSet
}

func (v *NullablePrimaryContact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrimaryContact(val *PrimaryContact) *NullablePrimaryContact {
	return &NullablePrimaryContact{value: val, isSet: true}
}

func (v NullablePrimaryContact) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePrimaryContact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
