package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the ServiceSelection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceSelection{}

// ServiceSelection Service Selection Criteria.
type ServiceSelection struct {
	// A list of ServiceId.
	ServiceId []string `json:"serviceId"`
}

type _ServiceSelection ServiceSelection

// NewServiceSelection instantiates a new ServiceSelection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceSelection(serviceId []string) *ServiceSelection {
	this := ServiceSelection{}
	this.ServiceId = serviceId
	return &this
}

// NewServiceSelectionWithDefaults instantiates a new ServiceSelection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceSelectionWithDefaults() *ServiceSelection {
	this := ServiceSelection{}
	return &this
}

// GetServiceId returns the ServiceId field value
func (o *ServiceSelection) GetServiceId() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *ServiceSelection) GetServiceIdOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceId, true
}

// SetServiceId sets field value
func (o *ServiceSelection) SetServiceId(v []string) {
	o.ServiceId = v
}

func (o ServiceSelection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["serviceId"] = o.ServiceId
	return toSerialize, nil
}

type NullableServiceSelection struct {
	value *ServiceSelection
	isSet bool
}

func (v NullableServiceSelection) Get() *ServiceSelection {
	return v.value
}

func (v *NullableServiceSelection) Set(val *ServiceSelection) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceSelection) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceSelection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceSelection(val *ServiceSelection) *NullableServiceSelection {
	return &NullableServiceSelection{value: val, isSet: true}
}

func (v NullableServiceSelection) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableServiceSelection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
