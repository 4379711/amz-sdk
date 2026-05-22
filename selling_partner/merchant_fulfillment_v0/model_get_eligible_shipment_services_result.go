package merchant_fulfillment_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the GetEligibleShipmentServicesResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEligibleShipmentServicesResult{}

// GetEligibleShipmentServicesResult The payload for the `getEligibleShipmentServices` operation.
type GetEligibleShipmentServicesResult struct {
	// A list of shipping services offers.
	ShippingServiceList []ShippingService `json:"ShippingServiceList"`
	// List of services that are for some reason unavailable for this request
	RejectedShippingServiceList []RejectedShippingService `json:"RejectedShippingServiceList,omitempty"`
	// A list of temporarily unavailable carriers.
	TemporarilyUnavailableCarrierList []TemporarilyUnavailableCarrier `json:"TemporarilyUnavailableCarrierList,omitempty"`
	// List of carriers whose terms and conditions were not accepted by the seller.
	TermsAndConditionsNotAcceptedCarrierList []TermsAndConditionsNotAcceptedCarrier `json:"TermsAndConditionsNotAcceptedCarrierList,omitempty"`
}

type _GetEligibleShipmentServicesResult GetEligibleShipmentServicesResult

// NewGetEligibleShipmentServicesResult instantiates a new GetEligibleShipmentServicesResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEligibleShipmentServicesResult(shippingServiceList []ShippingService) *GetEligibleShipmentServicesResult {
	this := GetEligibleShipmentServicesResult{}
	this.ShippingServiceList = shippingServiceList
	return &this
}

// NewGetEligibleShipmentServicesResultWithDefaults instantiates a new GetEligibleShipmentServicesResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEligibleShipmentServicesResultWithDefaults() *GetEligibleShipmentServicesResult {
	this := GetEligibleShipmentServicesResult{}
	return &this
}

// GetShippingServiceList returns the ShippingServiceList field value
func (o *GetEligibleShipmentServicesResult) GetShippingServiceList() []ShippingService {
	if o == nil {
		var ret []ShippingService
		return ret
	}

	return o.ShippingServiceList
}

// GetShippingServiceListOk returns a tuple with the ShippingServiceList field value
// and a boolean to check if the value has been set.
func (o *GetEligibleShipmentServicesResult) GetShippingServiceListOk() ([]ShippingService, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShippingServiceList, true
}

// SetShippingServiceList sets field value
func (o *GetEligibleShipmentServicesResult) SetShippingServiceList(v []ShippingService) {
	o.ShippingServiceList = v
}

// GetRejectedShippingServiceList returns the RejectedShippingServiceList field value if set, zero value otherwise.
func (o *GetEligibleShipmentServicesResult) GetRejectedShippingServiceList() []RejectedShippingService {
	if o == nil || IsNil(o.RejectedShippingServiceList) {
		var ret []RejectedShippingService
		return ret
	}
	return o.RejectedShippingServiceList
}

// GetRejectedShippingServiceListOk returns a tuple with the RejectedShippingServiceList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEligibleShipmentServicesResult) GetRejectedShippingServiceListOk() ([]RejectedShippingService, bool) {
	if o == nil || IsNil(o.RejectedShippingServiceList) {
		return nil, false
	}
	return o.RejectedShippingServiceList, true
}

// HasRejectedShippingServiceList returns a boolean if a field has been set.
func (o *GetEligibleShipmentServicesResult) HasRejectedShippingServiceList() bool {
	if o != nil && !IsNil(o.RejectedShippingServiceList) {
		return true
	}

	return false
}

// SetRejectedShippingServiceList gets a reference to the given []RejectedShippingService and assigns it to the RejectedShippingServiceList field.
func (o *GetEligibleShipmentServicesResult) SetRejectedShippingServiceList(v []RejectedShippingService) {
	o.RejectedShippingServiceList = v
}

// GetTemporarilyUnavailableCarrierList returns the TemporarilyUnavailableCarrierList field value if set, zero value otherwise.
func (o *GetEligibleShipmentServicesResult) GetTemporarilyUnavailableCarrierList() []TemporarilyUnavailableCarrier {
	if o == nil || IsNil(o.TemporarilyUnavailableCarrierList) {
		var ret []TemporarilyUnavailableCarrier
		return ret
	}
	return o.TemporarilyUnavailableCarrierList
}

// GetTemporarilyUnavailableCarrierListOk returns a tuple with the TemporarilyUnavailableCarrierList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEligibleShipmentServicesResult) GetTemporarilyUnavailableCarrierListOk() ([]TemporarilyUnavailableCarrier, bool) {
	if o == nil || IsNil(o.TemporarilyUnavailableCarrierList) {
		return nil, false
	}
	return o.TemporarilyUnavailableCarrierList, true
}

// HasTemporarilyUnavailableCarrierList returns a boolean if a field has been set.
func (o *GetEligibleShipmentServicesResult) HasTemporarilyUnavailableCarrierList() bool {
	if o != nil && !IsNil(o.TemporarilyUnavailableCarrierList) {
		return true
	}

	return false
}

// SetTemporarilyUnavailableCarrierList gets a reference to the given []TemporarilyUnavailableCarrier and assigns it to the TemporarilyUnavailableCarrierList field.
func (o *GetEligibleShipmentServicesResult) SetTemporarilyUnavailableCarrierList(v []TemporarilyUnavailableCarrier) {
	o.TemporarilyUnavailableCarrierList = v
}

// GetTermsAndConditionsNotAcceptedCarrierList returns the TermsAndConditionsNotAcceptedCarrierList field value if set, zero value otherwise.
func (o *GetEligibleShipmentServicesResult) GetTermsAndConditionsNotAcceptedCarrierList() []TermsAndConditionsNotAcceptedCarrier {
	if o == nil || IsNil(o.TermsAndConditionsNotAcceptedCarrierList) {
		var ret []TermsAndConditionsNotAcceptedCarrier
		return ret
	}
	return o.TermsAndConditionsNotAcceptedCarrierList
}

// GetTermsAndConditionsNotAcceptedCarrierListOk returns a tuple with the TermsAndConditionsNotAcceptedCarrierList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEligibleShipmentServicesResult) GetTermsAndConditionsNotAcceptedCarrierListOk() ([]TermsAndConditionsNotAcceptedCarrier, bool) {
	if o == nil || IsNil(o.TermsAndConditionsNotAcceptedCarrierList) {
		return nil, false
	}
	return o.TermsAndConditionsNotAcceptedCarrierList, true
}

// HasTermsAndConditionsNotAcceptedCarrierList returns a boolean if a field has been set.
func (o *GetEligibleShipmentServicesResult) HasTermsAndConditionsNotAcceptedCarrierList() bool {
	if o != nil && !IsNil(o.TermsAndConditionsNotAcceptedCarrierList) {
		return true
	}

	return false
}

// SetTermsAndConditionsNotAcceptedCarrierList gets a reference to the given []TermsAndConditionsNotAcceptedCarrier and assigns it to the TermsAndConditionsNotAcceptedCarrierList field.
func (o *GetEligibleShipmentServicesResult) SetTermsAndConditionsNotAcceptedCarrierList(v []TermsAndConditionsNotAcceptedCarrier) {
	o.TermsAndConditionsNotAcceptedCarrierList = v
}

func (o GetEligibleShipmentServicesResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ShippingServiceList"] = o.ShippingServiceList
	if !IsNil(o.RejectedShippingServiceList) {
		toSerialize["RejectedShippingServiceList"] = o.RejectedShippingServiceList
	}
	if !IsNil(o.TemporarilyUnavailableCarrierList) {
		toSerialize["TemporarilyUnavailableCarrierList"] = o.TemporarilyUnavailableCarrierList
	}
	if !IsNil(o.TermsAndConditionsNotAcceptedCarrierList) {
		toSerialize["TermsAndConditionsNotAcceptedCarrierList"] = o.TermsAndConditionsNotAcceptedCarrierList
	}
	return toSerialize, nil
}

type NullableGetEligibleShipmentServicesResult struct {
	value *GetEligibleShipmentServicesResult
	isSet bool
}

func (v NullableGetEligibleShipmentServicesResult) Get() *GetEligibleShipmentServicesResult {
	return v.value
}

func (v *NullableGetEligibleShipmentServicesResult) Set(val *GetEligibleShipmentServicesResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEligibleShipmentServicesResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEligibleShipmentServicesResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEligibleShipmentServicesResult(val *GetEligibleShipmentServicesResult) *NullableGetEligibleShipmentServicesResult {
	return &NullableGetEligibleShipmentServicesResult{value: val, isSet: true}
}

func (v NullableGetEligibleShipmentServicesResult) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetEligibleShipmentServicesResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
