package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the GenerateTransportationOptionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenerateTransportationOptionsRequest{}

// GenerateTransportationOptionsRequest The `generateTransportationOptions` request.
type GenerateTransportationOptionsRequest struct {
	// The placement option to generate transportation options for.
	PlacementOptionId string `json:"placementOptionId" validate:"regexp=^[a-zA-Z0-9-]*$"`
	// List of shipment transportation configurations.
	ShipmentTransportationConfigurations []ShipmentTransportationConfiguration `json:"shipmentTransportationConfigurations"`
}

type _GenerateTransportationOptionsRequest GenerateTransportationOptionsRequest

// NewGenerateTransportationOptionsRequest instantiates a new GenerateTransportationOptionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateTransportationOptionsRequest(placementOptionId string, shipmentTransportationConfigurations []ShipmentTransportationConfiguration) *GenerateTransportationOptionsRequest {
	this := GenerateTransportationOptionsRequest{}
	this.PlacementOptionId = placementOptionId
	this.ShipmentTransportationConfigurations = shipmentTransportationConfigurations
	return &this
}

// NewGenerateTransportationOptionsRequestWithDefaults instantiates a new GenerateTransportationOptionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateTransportationOptionsRequestWithDefaults() *GenerateTransportationOptionsRequest {
	this := GenerateTransportationOptionsRequest{}
	return &this
}

// GetPlacementOptionId returns the PlacementOptionId field value
func (o *GenerateTransportationOptionsRequest) GetPlacementOptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlacementOptionId
}

// GetPlacementOptionIdOk returns a tuple with the PlacementOptionId field value
// and a boolean to check if the value has been set.
func (o *GenerateTransportationOptionsRequest) GetPlacementOptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlacementOptionId, true
}

// SetPlacementOptionId sets field value
func (o *GenerateTransportationOptionsRequest) SetPlacementOptionId(v string) {
	o.PlacementOptionId = v
}

// GetShipmentTransportationConfigurations returns the ShipmentTransportationConfigurations field value
func (o *GenerateTransportationOptionsRequest) GetShipmentTransportationConfigurations() []ShipmentTransportationConfiguration {
	if o == nil {
		var ret []ShipmentTransportationConfiguration
		return ret
	}

	return o.ShipmentTransportationConfigurations
}

// GetShipmentTransportationConfigurationsOk returns a tuple with the ShipmentTransportationConfigurations field value
// and a boolean to check if the value has been set.
func (o *GenerateTransportationOptionsRequest) GetShipmentTransportationConfigurationsOk() ([]ShipmentTransportationConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShipmentTransportationConfigurations, true
}

// SetShipmentTransportationConfigurations sets field value
func (o *GenerateTransportationOptionsRequest) SetShipmentTransportationConfigurations(v []ShipmentTransportationConfiguration) {
	o.ShipmentTransportationConfigurations = v
}

func (o GenerateTransportationOptionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["placementOptionId"] = o.PlacementOptionId
	toSerialize["shipmentTransportationConfigurations"] = o.ShipmentTransportationConfigurations
	return toSerialize, nil
}

type NullableGenerateTransportationOptionsRequest struct {
	value *GenerateTransportationOptionsRequest
	isSet bool
}

func (v NullableGenerateTransportationOptionsRequest) Get() *GenerateTransportationOptionsRequest {
	return v.value
}

func (v *NullableGenerateTransportationOptionsRequest) Set(val *GenerateTransportationOptionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateTransportationOptionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateTransportationOptionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateTransportationOptionsRequest(val *GenerateTransportationOptionsRequest) *NullableGenerateTransportationOptionsRequest {
	return &NullableGenerateTransportationOptionsRequest{value: val, isSet: true}
}

func (v NullableGenerateTransportationOptionsRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGenerateTransportationOptionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
