package supply_sources_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the SupplySourceCapabilities type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SupplySourceCapabilities{}

// SupplySourceCapabilities The capabilities of a supply source.
type SupplySourceCapabilities struct {
	Outbound *OutboundCapability `json:"outbound,omitempty"`
	Services *ServicesCapability `json:"services,omitempty"`
}

// NewSupplySourceCapabilities instantiates a new SupplySourceCapabilities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupplySourceCapabilities() *SupplySourceCapabilities {
	this := SupplySourceCapabilities{}
	return &this
}

// NewSupplySourceCapabilitiesWithDefaults instantiates a new SupplySourceCapabilities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupplySourceCapabilitiesWithDefaults() *SupplySourceCapabilities {
	this := SupplySourceCapabilities{}
	return &this
}

// GetOutbound returns the Outbound field value if set, zero value otherwise.
func (o *SupplySourceCapabilities) GetOutbound() OutboundCapability {
	if o == nil || IsNil(o.Outbound) {
		var ret OutboundCapability
		return ret
	}
	return *o.Outbound
}

// GetOutboundOk returns a tuple with the Outbound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplySourceCapabilities) GetOutboundOk() (*OutboundCapability, bool) {
	if o == nil || IsNil(o.Outbound) {
		return nil, false
	}
	return o.Outbound, true
}

// HasOutbound returns a boolean if a field has been set.
func (o *SupplySourceCapabilities) HasOutbound() bool {
	if o != nil && !IsNil(o.Outbound) {
		return true
	}

	return false
}

// SetOutbound gets a reference to the given OutboundCapability and assigns it to the Outbound field.
func (o *SupplySourceCapabilities) SetOutbound(v OutboundCapability) {
	o.Outbound = &v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *SupplySourceCapabilities) GetServices() ServicesCapability {
	if o == nil || IsNil(o.Services) {
		var ret ServicesCapability
		return ret
	}
	return *o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupplySourceCapabilities) GetServicesOk() (*ServicesCapability, bool) {
	if o == nil || IsNil(o.Services) {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *SupplySourceCapabilities) HasServices() bool {
	if o != nil && !IsNil(o.Services) {
		return true
	}

	return false
}

// SetServices gets a reference to the given ServicesCapability and assigns it to the Services field.
func (o *SupplySourceCapabilities) SetServices(v ServicesCapability) {
	o.Services = &v
}

func (o SupplySourceCapabilities) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Outbound) {
		toSerialize["outbound"] = o.Outbound
	}
	if !IsNil(o.Services) {
		toSerialize["services"] = o.Services
	}
	return toSerialize, nil
}

type NullableSupplySourceCapabilities struct {
	value *SupplySourceCapabilities
	isSet bool
}

func (v NullableSupplySourceCapabilities) Get() *SupplySourceCapabilities {
	return v.value
}

func (v *NullableSupplySourceCapabilities) Set(val *SupplySourceCapabilities) {
	v.value = val
	v.isSet = true
}

func (v NullableSupplySourceCapabilities) IsSet() bool {
	return v.isSet
}

func (v *NullableSupplySourceCapabilities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupplySourceCapabilities(val *SupplySourceCapabilities) *NullableSupplySourceCapabilities {
	return &NullableSupplySourceCapabilities{value: val, isSet: true}
}

func (v NullableSupplySourceCapabilities) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSupplySourceCapabilities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
