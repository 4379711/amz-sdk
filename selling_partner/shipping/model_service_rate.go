package shipping

import (
	"github.com/bytedance/sonic"
)

// checks if the ServiceRate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceRate{}

// ServiceRate The specific rate for a shipping service, or null if no service available.
type ServiceRate struct {
	TotalCharge    Currency           `json:"totalCharge"`
	BillableWeight Weight             `json:"billableWeight"`
	ServiceType    ServiceType        `json:"serviceType"`
	Promise        ShippingPromiseSet `json:"promise"`
}

type _ServiceRate ServiceRate

// NewServiceRate instantiates a new ServiceRate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceRate(totalCharge Currency, billableWeight Weight, serviceType ServiceType, promise ShippingPromiseSet) *ServiceRate {
	this := ServiceRate{}
	this.TotalCharge = totalCharge
	this.BillableWeight = billableWeight
	this.ServiceType = serviceType
	this.Promise = promise
	return &this
}

// NewServiceRateWithDefaults instantiates a new ServiceRate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceRateWithDefaults() *ServiceRate {
	this := ServiceRate{}
	return &this
}

// GetTotalCharge returns the TotalCharge field value
func (o *ServiceRate) GetTotalCharge() Currency {
	if o == nil {
		var ret Currency
		return ret
	}

	return o.TotalCharge
}

// GetTotalChargeOk returns a tuple with the TotalCharge field value
// and a boolean to check if the value has been set.
func (o *ServiceRate) GetTotalChargeOk() (*Currency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCharge, true
}

// SetTotalCharge sets field value
func (o *ServiceRate) SetTotalCharge(v Currency) {
	o.TotalCharge = v
}

// GetBillableWeight returns the BillableWeight field value
func (o *ServiceRate) GetBillableWeight() Weight {
	if o == nil {
		var ret Weight
		return ret
	}

	return o.BillableWeight
}

// GetBillableWeightOk returns a tuple with the BillableWeight field value
// and a boolean to check if the value has been set.
func (o *ServiceRate) GetBillableWeightOk() (*Weight, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BillableWeight, true
}

// SetBillableWeight sets field value
func (o *ServiceRate) SetBillableWeight(v Weight) {
	o.BillableWeight = v
}

// GetServiceType returns the ServiceType field value
func (o *ServiceRate) GetServiceType() ServiceType {
	if o == nil {
		var ret ServiceType
		return ret
	}

	return o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value
// and a boolean to check if the value has been set.
func (o *ServiceRate) GetServiceTypeOk() (*ServiceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceType, true
}

// SetServiceType sets field value
func (o *ServiceRate) SetServiceType(v ServiceType) {
	o.ServiceType = v
}

// GetPromise returns the Promise field value
func (o *ServiceRate) GetPromise() ShippingPromiseSet {
	if o == nil {
		var ret ShippingPromiseSet
		return ret
	}

	return o.Promise
}

// GetPromiseOk returns a tuple with the Promise field value
// and a boolean to check if the value has been set.
func (o *ServiceRate) GetPromiseOk() (*ShippingPromiseSet, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Promise, true
}

// SetPromise sets field value
func (o *ServiceRate) SetPromise(v ShippingPromiseSet) {
	o.Promise = v
}

func (o ServiceRate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["totalCharge"] = o.TotalCharge
	toSerialize["billableWeight"] = o.BillableWeight
	toSerialize["serviceType"] = o.ServiceType
	toSerialize["promise"] = o.Promise
	return toSerialize, nil
}

type NullableServiceRate struct {
	value *ServiceRate
	isSet bool
}

func (v NullableServiceRate) Get() *ServiceRate {
	return v.value
}

func (v *NullableServiceRate) Set(val *ServiceRate) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceRate) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceRate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceRate(val *ServiceRate) *NullableServiceRate {
	return &NullableServiceRate{value: val, isSet: true}
}

func (v NullableServiceRate) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableServiceRate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
