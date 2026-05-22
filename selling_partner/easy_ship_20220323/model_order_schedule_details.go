package easy_ship_20220323

import (
	"github.com/bytedance/sonic"
)

// checks if the OrderScheduleDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderScheduleDetails{}

// OrderScheduleDetails This object allows users to specify an order to be scheduled. Only the amazonOrderId is required.
type OrderScheduleDetails struct {
	// An Amazon-defined order identifier. Identifies the order that the seller wants to deliver using Amazon Easy Ship.
	AmazonOrderId  string          `json:"amazonOrderId"`
	PackageDetails *PackageDetails `json:"packageDetails,omitempty"`
}

type _OrderScheduleDetails OrderScheduleDetails

// NewOrderScheduleDetails instantiates a new OrderScheduleDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderScheduleDetails(amazonOrderId string) *OrderScheduleDetails {
	this := OrderScheduleDetails{}
	this.AmazonOrderId = amazonOrderId
	return &this
}

// NewOrderScheduleDetailsWithDefaults instantiates a new OrderScheduleDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderScheduleDetailsWithDefaults() *OrderScheduleDetails {
	this := OrderScheduleDetails{}
	return &this
}

// GetAmazonOrderId returns the AmazonOrderId field value
func (o *OrderScheduleDetails) GetAmazonOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AmazonOrderId
}

// GetAmazonOrderIdOk returns a tuple with the AmazonOrderId field value
// and a boolean to check if the value has been set.
func (o *OrderScheduleDetails) GetAmazonOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmazonOrderId, true
}

// SetAmazonOrderId sets field value
func (o *OrderScheduleDetails) SetAmazonOrderId(v string) {
	o.AmazonOrderId = v
}

// GetPackageDetails returns the PackageDetails field value if set, zero value otherwise.
func (o *OrderScheduleDetails) GetPackageDetails() PackageDetails {
	if o == nil || IsNil(o.PackageDetails) {
		var ret PackageDetails
		return ret
	}
	return *o.PackageDetails
}

// GetPackageDetailsOk returns a tuple with the PackageDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderScheduleDetails) GetPackageDetailsOk() (*PackageDetails, bool) {
	if o == nil || IsNil(o.PackageDetails) {
		return nil, false
	}
	return o.PackageDetails, true
}

// HasPackageDetails returns a boolean if a field has been set.
func (o *OrderScheduleDetails) HasPackageDetails() bool {
	if o != nil && !IsNil(o.PackageDetails) {
		return true
	}

	return false
}

// SetPackageDetails gets a reference to the given PackageDetails and assigns it to the PackageDetails field.
func (o *OrderScheduleDetails) SetPackageDetails(v PackageDetails) {
	o.PackageDetails = &v
}

func (o OrderScheduleDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amazonOrderId"] = o.AmazonOrderId
	if !IsNil(o.PackageDetails) {
		toSerialize["packageDetails"] = o.PackageDetails
	}
	return toSerialize, nil
}

type NullableOrderScheduleDetails struct {
	value *OrderScheduleDetails
	isSet bool
}

func (v NullableOrderScheduleDetails) Get() *OrderScheduleDetails {
	return v.value
}

func (v *NullableOrderScheduleDetails) Set(val *OrderScheduleDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderScheduleDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderScheduleDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderScheduleDetails(val *OrderScheduleDetails) *NullableOrderScheduleDetails {
	return &NullableOrderScheduleDetails{value: val, isSet: true}
}

func (v NullableOrderScheduleDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOrderScheduleDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
