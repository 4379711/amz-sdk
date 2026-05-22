package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the AvailableValueAddedServiceGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AvailableValueAddedServiceGroup{}

// AvailableValueAddedServiceGroup The value-added services available for purchase with a shipping service offering.
type AvailableValueAddedServiceGroup struct {
	// The type of the value-added service group.
	GroupId string `json:"groupId"`
	// The name of the value-added service group.
	GroupDescription string `json:"groupDescription"`
	// When true, one or more of the value-added services listed must be specified.
	IsRequired bool `json:"isRequired"`
	// A list of optional value-added services available for purchase with a shipping service offering.
	ValueAddedServices []ValueAddedService `json:"valueAddedServices,omitempty"`
}

type _AvailableValueAddedServiceGroup AvailableValueAddedServiceGroup

// NewAvailableValueAddedServiceGroup instantiates a new AvailableValueAddedServiceGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailableValueAddedServiceGroup(groupId string, groupDescription string, isRequired bool) *AvailableValueAddedServiceGroup {
	this := AvailableValueAddedServiceGroup{}
	this.GroupId = groupId
	this.GroupDescription = groupDescription
	this.IsRequired = isRequired
	return &this
}

// NewAvailableValueAddedServiceGroupWithDefaults instantiates a new AvailableValueAddedServiceGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailableValueAddedServiceGroupWithDefaults() *AvailableValueAddedServiceGroup {
	this := AvailableValueAddedServiceGroup{}
	return &this
}

// GetGroupId returns the GroupId field value
func (o *AvailableValueAddedServiceGroup) GetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *AvailableValueAddedServiceGroup) GetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *AvailableValueAddedServiceGroup) SetGroupId(v string) {
	o.GroupId = v
}

// GetGroupDescription returns the GroupDescription field value
func (o *AvailableValueAddedServiceGroup) GetGroupDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupDescription
}

// GetGroupDescriptionOk returns a tuple with the GroupDescription field value
// and a boolean to check if the value has been set.
func (o *AvailableValueAddedServiceGroup) GetGroupDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupDescription, true
}

// SetGroupDescription sets field value
func (o *AvailableValueAddedServiceGroup) SetGroupDescription(v string) {
	o.GroupDescription = v
}

// GetIsRequired returns the IsRequired field value
func (o *AvailableValueAddedServiceGroup) GetIsRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsRequired
}

// GetIsRequiredOk returns a tuple with the IsRequired field value
// and a boolean to check if the value has been set.
func (o *AvailableValueAddedServiceGroup) GetIsRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsRequired, true
}

// SetIsRequired sets field value
func (o *AvailableValueAddedServiceGroup) SetIsRequired(v bool) {
	o.IsRequired = v
}

// GetValueAddedServices returns the ValueAddedServices field value if set, zero value otherwise.
func (o *AvailableValueAddedServiceGroup) GetValueAddedServices() []ValueAddedService {
	if o == nil || IsNil(o.ValueAddedServices) {
		var ret []ValueAddedService
		return ret
	}
	return o.ValueAddedServices
}

// GetValueAddedServicesOk returns a tuple with the ValueAddedServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableValueAddedServiceGroup) GetValueAddedServicesOk() ([]ValueAddedService, bool) {
	if o == nil || IsNil(o.ValueAddedServices) {
		return nil, false
	}
	return o.ValueAddedServices, true
}

// HasValueAddedServices returns a boolean if a field has been set.
func (o *AvailableValueAddedServiceGroup) HasValueAddedServices() bool {
	if o != nil && !IsNil(o.ValueAddedServices) {
		return true
	}

	return false
}

// SetValueAddedServices gets a reference to the given []ValueAddedService and assigns it to the ValueAddedServices field.
func (o *AvailableValueAddedServiceGroup) SetValueAddedServices(v []ValueAddedService) {
	o.ValueAddedServices = v
}

func (o AvailableValueAddedServiceGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["groupId"] = o.GroupId
	toSerialize["groupDescription"] = o.GroupDescription
	toSerialize["isRequired"] = o.IsRequired
	if !IsNil(o.ValueAddedServices) {
		toSerialize["valueAddedServices"] = o.ValueAddedServices
	}
	return toSerialize, nil
}

type NullableAvailableValueAddedServiceGroup struct {
	value *AvailableValueAddedServiceGroup
	isSet bool
}

func (v NullableAvailableValueAddedServiceGroup) Get() *AvailableValueAddedServiceGroup {
	return v.value
}

func (v *NullableAvailableValueAddedServiceGroup) Set(val *AvailableValueAddedServiceGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailableValueAddedServiceGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailableValueAddedServiceGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailableValueAddedServiceGroup(val *AvailableValueAddedServiceGroup) *NullableAvailableValueAddedServiceGroup {
	return &NullableAvailableValueAddedServiceGroup{value: val, isSet: true}
}

func (v NullableAvailableValueAddedServiceGroup) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAvailableValueAddedServiceGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
