package easy_ship_20220323

import (
	"github.com/bytedance/sonic"
)

// checks if the ScheduledPackageId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScheduledPackageId{}

// ScheduledPackageId Identifies the scheduled package to be updated.
type ScheduledPackageId struct {
	// An Amazon-defined order identifier. Identifies the order that the seller wants to deliver using Amazon Easy Ship.
	AmazonOrderId string `json:"amazonOrderId"`
	// An Amazon-defined identifier for the scheduled package.
	PackageId *string `json:"packageId,omitempty"`
}

type _ScheduledPackageId ScheduledPackageId

// NewScheduledPackageId instantiates a new ScheduledPackageId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduledPackageId(amazonOrderId string) *ScheduledPackageId {
	this := ScheduledPackageId{}
	this.AmazonOrderId = amazonOrderId
	return &this
}

// NewScheduledPackageIdWithDefaults instantiates a new ScheduledPackageId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduledPackageIdWithDefaults() *ScheduledPackageId {
	this := ScheduledPackageId{}
	return &this
}

// GetAmazonOrderId returns the AmazonOrderId field value
func (o *ScheduledPackageId) GetAmazonOrderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AmazonOrderId
}

// GetAmazonOrderIdOk returns a tuple with the AmazonOrderId field value
// and a boolean to check if the value has been set.
func (o *ScheduledPackageId) GetAmazonOrderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmazonOrderId, true
}

// SetAmazonOrderId sets field value
func (o *ScheduledPackageId) SetAmazonOrderId(v string) {
	o.AmazonOrderId = v
}

// GetPackageId returns the PackageId field value if set, zero value otherwise.
func (o *ScheduledPackageId) GetPackageId() string {
	if o == nil || IsNil(o.PackageId) {
		var ret string
		return ret
	}
	return *o.PackageId
}

// GetPackageIdOk returns a tuple with the PackageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledPackageId) GetPackageIdOk() (*string, bool) {
	if o == nil || IsNil(o.PackageId) {
		return nil, false
	}
	return o.PackageId, true
}

// HasPackageId returns a boolean if a field has been set.
func (o *ScheduledPackageId) HasPackageId() bool {
	if o != nil && !IsNil(o.PackageId) {
		return true
	}

	return false
}

// SetPackageId gets a reference to the given string and assigns it to the PackageId field.
func (o *ScheduledPackageId) SetPackageId(v string) {
	o.PackageId = &v
}

func (o ScheduledPackageId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amazonOrderId"] = o.AmazonOrderId
	if !IsNil(o.PackageId) {
		toSerialize["packageId"] = o.PackageId
	}
	return toSerialize, nil
}

type NullableScheduledPackageId struct {
	value *ScheduledPackageId
	isSet bool
}

func (v NullableScheduledPackageId) Get() *ScheduledPackageId {
	return v.value
}

func (v *NullableScheduledPackageId) Set(val *ScheduledPackageId) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduledPackageId) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduledPackageId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduledPackageId(val *ScheduledPackageId) *NullableScheduledPackageId {
	return &NullableScheduledPackageId{value: val, isSet: true}
}

func (v NullableScheduledPackageId) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableScheduledPackageId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
