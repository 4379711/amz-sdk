package services

import (
	"github.com/bytedance/sonic"
)

// checks if the ServiceJobProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceJobProvider{}

// ServiceJobProvider Information about the service job provider.
type ServiceJobProvider struct {
	// The identifier of the service job provider.
	ServiceJobProviderId *string `json:"serviceJobProviderId,omitempty" validate:"regexp=^[A-Z0-9]*$"`
}

// NewServiceJobProvider instantiates a new ServiceJobProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceJobProvider() *ServiceJobProvider {
	this := ServiceJobProvider{}
	return &this
}

// NewServiceJobProviderWithDefaults instantiates a new ServiceJobProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceJobProviderWithDefaults() *ServiceJobProvider {
	this := ServiceJobProvider{}
	return &this
}

// GetServiceJobProviderId returns the ServiceJobProviderId field value if set, zero value otherwise.
func (o *ServiceJobProvider) GetServiceJobProviderId() string {
	if o == nil || IsNil(o.ServiceJobProviderId) {
		var ret string
		return ret
	}
	return *o.ServiceJobProviderId
}

// GetServiceJobProviderIdOk returns a tuple with the ServiceJobProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceJobProvider) GetServiceJobProviderIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceJobProviderId) {
		return nil, false
	}
	return o.ServiceJobProviderId, true
}

// HasServiceJobProviderId returns a boolean if a field has been set.
func (o *ServiceJobProvider) HasServiceJobProviderId() bool {
	if o != nil && !IsNil(o.ServiceJobProviderId) {
		return true
	}

	return false
}

// SetServiceJobProviderId gets a reference to the given string and assigns it to the ServiceJobProviderId field.
func (o *ServiceJobProvider) SetServiceJobProviderId(v string) {
	o.ServiceJobProviderId = &v
}

func (o ServiceJobProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServiceJobProviderId) {
		toSerialize["serviceJobProviderId"] = o.ServiceJobProviderId
	}
	return toSerialize, nil
}

type NullableServiceJobProvider struct {
	value *ServiceJobProvider
	isSet bool
}

func (v NullableServiceJobProvider) Get() *ServiceJobProvider {
	return v.value
}

func (v *NullableServiceJobProvider) Set(val *ServiceJobProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceJobProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceJobProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceJobProvider(val *ServiceJobProvider) *NullableServiceJobProvider {
	return &NullableServiceJobProvider{value: val, isSet: true}
}

func (v NullableServiceJobProvider) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableServiceJobProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
