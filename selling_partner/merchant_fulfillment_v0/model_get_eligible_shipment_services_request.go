package merchant_fulfillment_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the GetEligibleShipmentServicesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEligibleShipmentServicesRequest{}

// GetEligibleShipmentServicesRequest Request schema.
type GetEligibleShipmentServicesRequest struct {
	ShipmentRequestDetails ShipmentRequestDetails  `json:"ShipmentRequestDetails"`
	ShippingOfferingFilter *ShippingOfferingFilter `json:"ShippingOfferingFilter,omitempty"`
}

type _GetEligibleShipmentServicesRequest GetEligibleShipmentServicesRequest

// NewGetEligibleShipmentServicesRequest instantiates a new GetEligibleShipmentServicesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEligibleShipmentServicesRequest(shipmentRequestDetails ShipmentRequestDetails) *GetEligibleShipmentServicesRequest {
	this := GetEligibleShipmentServicesRequest{}
	this.ShipmentRequestDetails = shipmentRequestDetails
	return &this
}

// NewGetEligibleShipmentServicesRequestWithDefaults instantiates a new GetEligibleShipmentServicesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEligibleShipmentServicesRequestWithDefaults() *GetEligibleShipmentServicesRequest {
	this := GetEligibleShipmentServicesRequest{}
	return &this
}

// GetShipmentRequestDetails returns the ShipmentRequestDetails field value
func (o *GetEligibleShipmentServicesRequest) GetShipmentRequestDetails() ShipmentRequestDetails {
	if o == nil {
		var ret ShipmentRequestDetails
		return ret
	}

	return o.ShipmentRequestDetails
}

// GetShipmentRequestDetailsOk returns a tuple with the ShipmentRequestDetails field value
// and a boolean to check if the value has been set.
func (o *GetEligibleShipmentServicesRequest) GetShipmentRequestDetailsOk() (*ShipmentRequestDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipmentRequestDetails, true
}

// SetShipmentRequestDetails sets field value
func (o *GetEligibleShipmentServicesRequest) SetShipmentRequestDetails(v ShipmentRequestDetails) {
	o.ShipmentRequestDetails = v
}

// GetShippingOfferingFilter returns the ShippingOfferingFilter field value if set, zero value otherwise.
func (o *GetEligibleShipmentServicesRequest) GetShippingOfferingFilter() ShippingOfferingFilter {
	if o == nil || IsNil(o.ShippingOfferingFilter) {
		var ret ShippingOfferingFilter
		return ret
	}
	return *o.ShippingOfferingFilter
}

// GetShippingOfferingFilterOk returns a tuple with the ShippingOfferingFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEligibleShipmentServicesRequest) GetShippingOfferingFilterOk() (*ShippingOfferingFilter, bool) {
	if o == nil || IsNil(o.ShippingOfferingFilter) {
		return nil, false
	}
	return o.ShippingOfferingFilter, true
}

// HasShippingOfferingFilter returns a boolean if a field has been set.
func (o *GetEligibleShipmentServicesRequest) HasShippingOfferingFilter() bool {
	if o != nil && !IsNil(o.ShippingOfferingFilter) {
		return true
	}

	return false
}

// SetShippingOfferingFilter gets a reference to the given ShippingOfferingFilter and assigns it to the ShippingOfferingFilter field.
func (o *GetEligibleShipmentServicesRequest) SetShippingOfferingFilter(v ShippingOfferingFilter) {
	o.ShippingOfferingFilter = &v
}

func (o GetEligibleShipmentServicesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ShipmentRequestDetails"] = o.ShipmentRequestDetails
	if !IsNil(o.ShippingOfferingFilter) {
		toSerialize["ShippingOfferingFilter"] = o.ShippingOfferingFilter
	}
	return toSerialize, nil
}

type NullableGetEligibleShipmentServicesRequest struct {
	value *GetEligibleShipmentServicesRequest
	isSet bool
}

func (v NullableGetEligibleShipmentServicesRequest) Get() *GetEligibleShipmentServicesRequest {
	return v.value
}

func (v *NullableGetEligibleShipmentServicesRequest) Set(val *GetEligibleShipmentServicesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEligibleShipmentServicesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEligibleShipmentServicesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEligibleShipmentServicesRequest(val *GetEligibleShipmentServicesRequest) *NullableGetEligibleShipmentServicesRequest {
	return &NullableGetEligibleShipmentServicesRequest{value: val, isSet: true}
}

func (v NullableGetEligibleShipmentServicesRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetEligibleShipmentServicesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
