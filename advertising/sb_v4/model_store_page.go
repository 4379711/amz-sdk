package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the StorePage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorePage{}

// StorePage struct for StorePage
type StorePage struct {
	// Display Name of the store page shown on a store spotlight campaign.
	DisplayName *string `json:"displayName,omitempty"`
	// Selected asin from the store page which is displayed on the store spotlight campaign.
	PrimaryAsin *string `json:"primaryAsin,omitempty"`
}

// NewStorePage instantiates a new StorePage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorePage() *StorePage {
	this := StorePage{}
	return &this
}

// NewStorePageWithDefaults instantiates a new StorePage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorePageWithDefaults() *StorePage {
	this := StorePage{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *StorePage) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorePage) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *StorePage) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *StorePage) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetPrimaryAsin returns the PrimaryAsin field value if set, zero value otherwise.
func (o *StorePage) GetPrimaryAsin() string {
	if o == nil || IsNil(o.PrimaryAsin) {
		var ret string
		return ret
	}
	return *o.PrimaryAsin
}

// GetPrimaryAsinOk returns a tuple with the PrimaryAsin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorePage) GetPrimaryAsinOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryAsin) {
		return nil, false
	}
	return o.PrimaryAsin, true
}

// HasPrimaryAsin returns a boolean if a field has been set.
func (o *StorePage) HasPrimaryAsin() bool {
	if o != nil && !IsNil(o.PrimaryAsin) {
		return true
	}

	return false
}

// SetPrimaryAsin gets a reference to the given string and assigns it to the PrimaryAsin field.
func (o *StorePage) SetPrimaryAsin(v string) {
	o.PrimaryAsin = &v
}

func (o StorePage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.PrimaryAsin) {
		toSerialize["primaryAsin"] = o.PrimaryAsin
	}
	return toSerialize, nil
}

type NullableStorePage struct {
	value *StorePage
	isSet bool
}

func (v NullableStorePage) Get() *StorePage {
	return v.value
}

func (v *NullableStorePage) Set(val *StorePage) {
	v.value = val
	v.isSet = true
}

func (v NullableStorePage) IsSet() bool {
	return v.isSet
}

func (v *NullableStorePage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorePage(val *StorePage) *NullableStorePage {
	return &NullableStorePage{value: val, isSet: true}
}

func (v NullableStorePage) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableStorePage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
