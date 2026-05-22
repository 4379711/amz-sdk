package supply_sources_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateSupplySourceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSupplySourceRequest{}

// UpdateSupplySourceRequest A request to update the configuration and capabilities of a supply source.
type UpdateSupplySourceRequest struct {
	// The custom alias for this supply source
	Alias         *string                    `json:"alias,omitempty"`
	Configuration *SupplySourceConfiguration `json:"configuration,omitempty"`
	Capabilities  *SupplySourceCapabilities  `json:"capabilities,omitempty"`
}

// NewUpdateSupplySourceRequest instantiates a new UpdateSupplySourceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSupplySourceRequest() *UpdateSupplySourceRequest {
	this := UpdateSupplySourceRequest{}
	return &this
}

// NewUpdateSupplySourceRequestWithDefaults instantiates a new UpdateSupplySourceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSupplySourceRequestWithDefaults() *UpdateSupplySourceRequest {
	this := UpdateSupplySourceRequest{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *UpdateSupplySourceRequest) GetAlias() string {
	if o == nil || IsNil(o.Alias) {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSupplySourceRequest) GetAliasOk() (*string, bool) {
	if o == nil || IsNil(o.Alias) {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *UpdateSupplySourceRequest) HasAlias() bool {
	if o != nil && !IsNil(o.Alias) {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *UpdateSupplySourceRequest) SetAlias(v string) {
	o.Alias = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *UpdateSupplySourceRequest) GetConfiguration() SupplySourceConfiguration {
	if o == nil || IsNil(o.Configuration) {
		var ret SupplySourceConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSupplySourceRequest) GetConfigurationOk() (*SupplySourceConfiguration, bool) {
	if o == nil || IsNil(o.Configuration) {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *UpdateSupplySourceRequest) HasConfiguration() bool {
	if o != nil && !IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given SupplySourceConfiguration and assigns it to the Configuration field.
func (o *UpdateSupplySourceRequest) SetConfiguration(v SupplySourceConfiguration) {
	o.Configuration = &v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *UpdateSupplySourceRequest) GetCapabilities() SupplySourceCapabilities {
	if o == nil || IsNil(o.Capabilities) {
		var ret SupplySourceCapabilities
		return ret
	}
	return *o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSupplySourceRequest) GetCapabilitiesOk() (*SupplySourceCapabilities, bool) {
	if o == nil || IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *UpdateSupplySourceRequest) HasCapabilities() bool {
	if o != nil && !IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given SupplySourceCapabilities and assigns it to the Capabilities field.
func (o *UpdateSupplySourceRequest) SetCapabilities(v SupplySourceCapabilities) {
	o.Capabilities = &v
}

func (o UpdateSupplySourceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Alias) {
		toSerialize["alias"] = o.Alias
	}
	if !IsNil(o.Configuration) {
		toSerialize["configuration"] = o.Configuration
	}
	if !IsNil(o.Capabilities) {
		toSerialize["capabilities"] = o.Capabilities
	}
	return toSerialize, nil
}

type NullableUpdateSupplySourceRequest struct {
	value *UpdateSupplySourceRequest
	isSet bool
}

func (v NullableUpdateSupplySourceRequest) Get() *UpdateSupplySourceRequest {
	return v.value
}

func (v *NullableUpdateSupplySourceRequest) Set(val *UpdateSupplySourceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSupplySourceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSupplySourceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSupplySourceRequest(val *UpdateSupplySourceRequest) *NullableUpdateSupplySourceRequest {
	return &NullableUpdateSupplySourceRequest{value: val, isSet: true}
}

func (v NullableUpdateSupplySourceRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateSupplySourceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
