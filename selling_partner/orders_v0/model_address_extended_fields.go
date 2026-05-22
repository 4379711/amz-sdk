package orders_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the AddressExtendedFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressExtendedFields{}

// AddressExtendedFields The container for address extended fields (such as `street name` and `street number`). Currently only available with Brazil shipping addresses.
type AddressExtendedFields struct {
	// The street name.
	StreetName *string `json:"StreetName,omitempty"`
	// The house, building, or property number associated with the location's street address.
	StreetNumber *string `json:"StreetNumber,omitempty"`
	// The floor number/unit number in the building/private house number.
	Complement *string `json:"Complement,omitempty"`
	// The neighborhood. This value is only used in some countries (such as Brazil).
	Neighborhood *string `json:"Neighborhood,omitempty"`
}

// NewAddressExtendedFields instantiates a new AddressExtendedFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressExtendedFields() *AddressExtendedFields {
	this := AddressExtendedFields{}
	return &this
}

// NewAddressExtendedFieldsWithDefaults instantiates a new AddressExtendedFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressExtendedFieldsWithDefaults() *AddressExtendedFields {
	this := AddressExtendedFields{}
	return &this
}

// GetStreetName returns the StreetName field value if set, zero value otherwise.
func (o *AddressExtendedFields) GetStreetName() string {
	if o == nil || IsNil(o.StreetName) {
		var ret string
		return ret
	}
	return *o.StreetName
}

// GetStreetNameOk returns a tuple with the StreetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressExtendedFields) GetStreetNameOk() (*string, bool) {
	if o == nil || IsNil(o.StreetName) {
		return nil, false
	}
	return o.StreetName, true
}

// HasStreetName returns a boolean if a field has been set.
func (o *AddressExtendedFields) HasStreetName() bool {
	if o != nil && !IsNil(o.StreetName) {
		return true
	}

	return false
}

// SetStreetName gets a reference to the given string and assigns it to the StreetName field.
func (o *AddressExtendedFields) SetStreetName(v string) {
	o.StreetName = &v
}

// GetStreetNumber returns the StreetNumber field value if set, zero value otherwise.
func (o *AddressExtendedFields) GetStreetNumber() string {
	if o == nil || IsNil(o.StreetNumber) {
		var ret string
		return ret
	}
	return *o.StreetNumber
}

// GetStreetNumberOk returns a tuple with the StreetNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressExtendedFields) GetStreetNumberOk() (*string, bool) {
	if o == nil || IsNil(o.StreetNumber) {
		return nil, false
	}
	return o.StreetNumber, true
}

// HasStreetNumber returns a boolean if a field has been set.
func (o *AddressExtendedFields) HasStreetNumber() bool {
	if o != nil && !IsNil(o.StreetNumber) {
		return true
	}

	return false
}

// SetStreetNumber gets a reference to the given string and assigns it to the StreetNumber field.
func (o *AddressExtendedFields) SetStreetNumber(v string) {
	o.StreetNumber = &v
}

// GetComplement returns the Complement field value if set, zero value otherwise.
func (o *AddressExtendedFields) GetComplement() string {
	if o == nil || IsNil(o.Complement) {
		var ret string
		return ret
	}
	return *o.Complement
}

// GetComplementOk returns a tuple with the Complement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressExtendedFields) GetComplementOk() (*string, bool) {
	if o == nil || IsNil(o.Complement) {
		return nil, false
	}
	return o.Complement, true
}

// HasComplement returns a boolean if a field has been set.
func (o *AddressExtendedFields) HasComplement() bool {
	if o != nil && !IsNil(o.Complement) {
		return true
	}

	return false
}

// SetComplement gets a reference to the given string and assigns it to the Complement field.
func (o *AddressExtendedFields) SetComplement(v string) {
	o.Complement = &v
}

// GetNeighborhood returns the Neighborhood field value if set, zero value otherwise.
func (o *AddressExtendedFields) GetNeighborhood() string {
	if o == nil || IsNil(o.Neighborhood) {
		var ret string
		return ret
	}
	return *o.Neighborhood
}

// GetNeighborhoodOk returns a tuple with the Neighborhood field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressExtendedFields) GetNeighborhoodOk() (*string, bool) {
	if o == nil || IsNil(o.Neighborhood) {
		return nil, false
	}
	return o.Neighborhood, true
}

// HasNeighborhood returns a boolean if a field has been set.
func (o *AddressExtendedFields) HasNeighborhood() bool {
	if o != nil && !IsNil(o.Neighborhood) {
		return true
	}

	return false
}

// SetNeighborhood gets a reference to the given string and assigns it to the Neighborhood field.
func (o *AddressExtendedFields) SetNeighborhood(v string) {
	o.Neighborhood = &v
}

func (o AddressExtendedFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StreetName) {
		toSerialize["StreetName"] = o.StreetName
	}
	if !IsNil(o.StreetNumber) {
		toSerialize["StreetNumber"] = o.StreetNumber
	}
	if !IsNil(o.Complement) {
		toSerialize["Complement"] = o.Complement
	}
	if !IsNil(o.Neighborhood) {
		toSerialize["Neighborhood"] = o.Neighborhood
	}
	return toSerialize, nil
}

type NullableAddressExtendedFields struct {
	value *AddressExtendedFields
	isSet bool
}

func (v NullableAddressExtendedFields) Get() *AddressExtendedFields {
	return v.value
}

func (v *NullableAddressExtendedFields) Set(val *AddressExtendedFields) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressExtendedFields) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressExtendedFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressExtendedFields(val *AddressExtendedFields) *NullableAddressExtendedFields {
	return &NullableAddressExtendedFields{value: val, isSet: true}
}

func (v NullableAddressExtendedFields) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAddressExtendedFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
