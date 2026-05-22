package supply_sources_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the ReturnLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReturnLocation{}

// ReturnLocation The address or reference to another `supplySourceId` to act as a return location.
type ReturnLocation struct {
	// The Amazon provided `supplySourceId` where orders can be returned to.
	SupplySourceId     *string             `json:"supplySourceId,omitempty"`
	AddressWithContact *AddressWithContact `json:"addressWithContact,omitempty"`
}

// NewReturnLocation instantiates a new ReturnLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnLocation() *ReturnLocation {
	this := ReturnLocation{}
	return &this
}

// NewReturnLocationWithDefaults instantiates a new ReturnLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnLocationWithDefaults() *ReturnLocation {
	this := ReturnLocation{}
	return &this
}

// GetSupplySourceId returns the SupplySourceId field value if set, zero value otherwise.
func (o *ReturnLocation) GetSupplySourceId() string {
	if o == nil || IsNil(o.SupplySourceId) {
		var ret string
		return ret
	}
	return *o.SupplySourceId
}

// GetSupplySourceIdOk returns a tuple with the SupplySourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnLocation) GetSupplySourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.SupplySourceId) {
		return nil, false
	}
	return o.SupplySourceId, true
}

// HasSupplySourceId returns a boolean if a field has been set.
func (o *ReturnLocation) HasSupplySourceId() bool {
	if o != nil && !IsNil(o.SupplySourceId) {
		return true
	}

	return false
}

// SetSupplySourceId gets a reference to the given string and assigns it to the SupplySourceId field.
func (o *ReturnLocation) SetSupplySourceId(v string) {
	o.SupplySourceId = &v
}

// GetAddressWithContact returns the AddressWithContact field value if set, zero value otherwise.
func (o *ReturnLocation) GetAddressWithContact() AddressWithContact {
	if o == nil || IsNil(o.AddressWithContact) {
		var ret AddressWithContact
		return ret
	}
	return *o.AddressWithContact
}

// GetAddressWithContactOk returns a tuple with the AddressWithContact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnLocation) GetAddressWithContactOk() (*AddressWithContact, bool) {
	if o == nil || IsNil(o.AddressWithContact) {
		return nil, false
	}
	return o.AddressWithContact, true
}

// HasAddressWithContact returns a boolean if a field has been set.
func (o *ReturnLocation) HasAddressWithContact() bool {
	if o != nil && !IsNil(o.AddressWithContact) {
		return true
	}

	return false
}

// SetAddressWithContact gets a reference to the given AddressWithContact and assigns it to the AddressWithContact field.
func (o *ReturnLocation) SetAddressWithContact(v AddressWithContact) {
	o.AddressWithContact = &v
}

func (o ReturnLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SupplySourceId) {
		toSerialize["supplySourceId"] = o.SupplySourceId
	}
	if !IsNil(o.AddressWithContact) {
		toSerialize["addressWithContact"] = o.AddressWithContact
	}
	return toSerialize, nil
}

type NullableReturnLocation struct {
	value *ReturnLocation
	isSet bool
}

func (v NullableReturnLocation) Get() *ReturnLocation {
	return v.value
}

func (v *NullableReturnLocation) Set(val *ReturnLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnLocation(val *ReturnLocation) *NullableReturnLocation {
	return &NullableReturnLocation{value: val, isSet: true}
}

func (v NullableReturnLocation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableReturnLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
