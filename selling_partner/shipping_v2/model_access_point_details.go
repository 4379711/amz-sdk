package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the AccessPointDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessPointDetails{}

// AccessPointDetails AccessPointDetails object
type AccessPointDetails struct {
	// Unique identifier for the access point
	AccessPointId *string `json:"accessPointId,omitempty"`
}

// NewAccessPointDetails instantiates a new AccessPointDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessPointDetails() *AccessPointDetails {
	this := AccessPointDetails{}
	return &this
}

// NewAccessPointDetailsWithDefaults instantiates a new AccessPointDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessPointDetailsWithDefaults() *AccessPointDetails {
	this := AccessPointDetails{}
	return &this
}

// GetAccessPointId returns the AccessPointId field value if set, zero value otherwise.
func (o *AccessPointDetails) GetAccessPointId() string {
	if o == nil || IsNil(o.AccessPointId) {
		var ret string
		return ret
	}
	return *o.AccessPointId
}

// GetAccessPointIdOk returns a tuple with the AccessPointId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPointDetails) GetAccessPointIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccessPointId) {
		return nil, false
	}
	return o.AccessPointId, true
}

// HasAccessPointId returns a boolean if a field has been set.
func (o *AccessPointDetails) HasAccessPointId() bool {
	if o != nil && !IsNil(o.AccessPointId) {
		return true
	}

	return false
}

// SetAccessPointId gets a reference to the given string and assigns it to the AccessPointId field.
func (o *AccessPointDetails) SetAccessPointId(v string) {
	o.AccessPointId = &v
}

func (o AccessPointDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessPointId) {
		toSerialize["accessPointId"] = o.AccessPointId
	}
	return toSerialize, nil
}

type NullableAccessPointDetails struct {
	value *AccessPointDetails
	isSet bool
}

func (v NullableAccessPointDetails) Get() *AccessPointDetails {
	return v.value
}

func (v *NullableAccessPointDetails) Set(val *AccessPointDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessPointDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessPointDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessPointDetails(val *AccessPointDetails) *NullableAccessPointDetails {
	return &NullableAccessPointDetails{value: val, isSet: true}
}

func (v NullableAccessPointDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAccessPointDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
