package shipping

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the GetRatesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRatesRequest{}

// GetRatesRequest The payload schema for the getRates operation.
type GetRatesRequest struct {
	ShipTo   Address `json:"shipTo"`
	ShipFrom Address `json:"shipFrom"`
	// A list of service types that can be used to send the shipment.
	ServiceTypes []ServiceType `json:"serviceTypes"`
	// The start date and time. This defaults to the current date and time.
	ShipDate *time.Time `json:"shipDate,omitempty"`
	// A list of container specifications.
	ContainerSpecifications []ContainerSpecification `json:"containerSpecifications"`
}

type _GetRatesRequest GetRatesRequest

// NewGetRatesRequest instantiates a new GetRatesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRatesRequest(shipTo Address, shipFrom Address, serviceTypes []ServiceType, containerSpecifications []ContainerSpecification) *GetRatesRequest {
	this := GetRatesRequest{}
	this.ShipTo = shipTo
	this.ShipFrom = shipFrom
	this.ServiceTypes = serviceTypes
	this.ContainerSpecifications = containerSpecifications
	return &this
}

// NewGetRatesRequestWithDefaults instantiates a new GetRatesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRatesRequestWithDefaults() *GetRatesRequest {
	this := GetRatesRequest{}
	return &this
}

// GetShipTo returns the ShipTo field value
func (o *GetRatesRequest) GetShipTo() Address {
	if o == nil {
		var ret Address
		return ret
	}

	return o.ShipTo
}

// GetShipToOk returns a tuple with the ShipTo field value
// and a boolean to check if the value has been set.
func (o *GetRatesRequest) GetShipToOk() (*Address, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipTo, true
}

// SetShipTo sets field value
func (o *GetRatesRequest) SetShipTo(v Address) {
	o.ShipTo = v
}

// GetShipFrom returns the ShipFrom field value
func (o *GetRatesRequest) GetShipFrom() Address {
	if o == nil {
		var ret Address
		return ret
	}

	return o.ShipFrom
}

// GetShipFromOk returns a tuple with the ShipFrom field value
// and a boolean to check if the value has been set.
func (o *GetRatesRequest) GetShipFromOk() (*Address, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipFrom, true
}

// SetShipFrom sets field value
func (o *GetRatesRequest) SetShipFrom(v Address) {
	o.ShipFrom = v
}

// GetServiceTypes returns the ServiceTypes field value
func (o *GetRatesRequest) GetServiceTypes() []ServiceType {
	if o == nil {
		var ret []ServiceType
		return ret
	}

	return o.ServiceTypes
}

// GetServiceTypesOk returns a tuple with the ServiceTypes field value
// and a boolean to check if the value has been set.
func (o *GetRatesRequest) GetServiceTypesOk() ([]ServiceType, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceTypes, true
}

// SetServiceTypes sets field value
func (o *GetRatesRequest) SetServiceTypes(v []ServiceType) {
	o.ServiceTypes = v
}

// GetShipDate returns the ShipDate field value if set, zero value otherwise.
func (o *GetRatesRequest) GetShipDate() time.Time {
	if o == nil || IsNil(o.ShipDate) {
		var ret time.Time
		return ret
	}
	return *o.ShipDate
}

// GetShipDateOk returns a tuple with the ShipDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRatesRequest) GetShipDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ShipDate) {
		return nil, false
	}
	return o.ShipDate, true
}

// HasShipDate returns a boolean if a field has been set.
func (o *GetRatesRequest) HasShipDate() bool {
	if o != nil && !IsNil(o.ShipDate) {
		return true
	}

	return false
}

// SetShipDate gets a reference to the given time.Time and assigns it to the ShipDate field.
func (o *GetRatesRequest) SetShipDate(v time.Time) {
	o.ShipDate = &v
}

// GetContainerSpecifications returns the ContainerSpecifications field value
func (o *GetRatesRequest) GetContainerSpecifications() []ContainerSpecification {
	if o == nil {
		var ret []ContainerSpecification
		return ret
	}

	return o.ContainerSpecifications
}

// GetContainerSpecificationsOk returns a tuple with the ContainerSpecifications field value
// and a boolean to check if the value has been set.
func (o *GetRatesRequest) GetContainerSpecificationsOk() ([]ContainerSpecification, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContainerSpecifications, true
}

// SetContainerSpecifications sets field value
func (o *GetRatesRequest) SetContainerSpecifications(v []ContainerSpecification) {
	o.ContainerSpecifications = v
}

func (o GetRatesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["shipTo"] = o.ShipTo
	toSerialize["shipFrom"] = o.ShipFrom
	toSerialize["serviceTypes"] = o.ServiceTypes
	if !IsNil(o.ShipDate) {
		toSerialize["shipDate"] = o.ShipDate
	}
	toSerialize["containerSpecifications"] = o.ContainerSpecifications
	return toSerialize, nil
}

type NullableGetRatesRequest struct {
	value *GetRatesRequest
	isSet bool
}

func (v NullableGetRatesRequest) Get() *GetRatesRequest {
	return v.value
}

func (v *NullableGetRatesRequest) Set(val *GetRatesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRatesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRatesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRatesRequest(val *GetRatesRequest) *NullableGetRatesRequest {
	return &NullableGetRatesRequest{value: val, isSet: true}
}

func (v NullableGetRatesRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetRatesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
