package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the ChannelDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelDetails{}

// ChannelDetails Shipment source channel related information.
type ChannelDetails struct {
	ChannelType           ChannelType            `json:"channelType"`
	AmazonOrderDetails    *AmazonOrderDetails    `json:"amazonOrderDetails,omitempty"`
	AmazonShipmentDetails *AmazonShipmentDetails `json:"amazonShipmentDetails,omitempty"`
}

type _ChannelDetails ChannelDetails

// NewChannelDetails instantiates a new ChannelDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelDetails(channelType ChannelType) *ChannelDetails {
	this := ChannelDetails{}
	this.ChannelType = channelType
	return &this
}

// NewChannelDetailsWithDefaults instantiates a new ChannelDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelDetailsWithDefaults() *ChannelDetails {
	this := ChannelDetails{}
	return &this
}

// GetChannelType returns the ChannelType field value
func (o *ChannelDetails) GetChannelType() ChannelType {
	if o == nil {
		var ret ChannelType
		return ret
	}

	return o.ChannelType
}

// GetChannelTypeOk returns a tuple with the ChannelType field value
// and a boolean to check if the value has been set.
func (o *ChannelDetails) GetChannelTypeOk() (*ChannelType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelType, true
}

// SetChannelType sets field value
func (o *ChannelDetails) SetChannelType(v ChannelType) {
	o.ChannelType = v
}

// GetAmazonOrderDetails returns the AmazonOrderDetails field value if set, zero value otherwise.
func (o *ChannelDetails) GetAmazonOrderDetails() AmazonOrderDetails {
	if o == nil || IsNil(o.AmazonOrderDetails) {
		var ret AmazonOrderDetails
		return ret
	}
	return *o.AmazonOrderDetails
}

// GetAmazonOrderDetailsOk returns a tuple with the AmazonOrderDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelDetails) GetAmazonOrderDetailsOk() (*AmazonOrderDetails, bool) {
	if o == nil || IsNil(o.AmazonOrderDetails) {
		return nil, false
	}
	return o.AmazonOrderDetails, true
}

// HasAmazonOrderDetails returns a boolean if a field has been set.
func (o *ChannelDetails) HasAmazonOrderDetails() bool {
	if o != nil && !IsNil(o.AmazonOrderDetails) {
		return true
	}

	return false
}

// SetAmazonOrderDetails gets a reference to the given AmazonOrderDetails and assigns it to the AmazonOrderDetails field.
func (o *ChannelDetails) SetAmazonOrderDetails(v AmazonOrderDetails) {
	o.AmazonOrderDetails = &v
}

// GetAmazonShipmentDetails returns the AmazonShipmentDetails field value if set, zero value otherwise.
func (o *ChannelDetails) GetAmazonShipmentDetails() AmazonShipmentDetails {
	if o == nil || IsNil(o.AmazonShipmentDetails) {
		var ret AmazonShipmentDetails
		return ret
	}
	return *o.AmazonShipmentDetails
}

// GetAmazonShipmentDetailsOk returns a tuple with the AmazonShipmentDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelDetails) GetAmazonShipmentDetailsOk() (*AmazonShipmentDetails, bool) {
	if o == nil || IsNil(o.AmazonShipmentDetails) {
		return nil, false
	}
	return o.AmazonShipmentDetails, true
}

// HasAmazonShipmentDetails returns a boolean if a field has been set.
func (o *ChannelDetails) HasAmazonShipmentDetails() bool {
	if o != nil && !IsNil(o.AmazonShipmentDetails) {
		return true
	}

	return false
}

// SetAmazonShipmentDetails gets a reference to the given AmazonShipmentDetails and assigns it to the AmazonShipmentDetails field.
func (o *ChannelDetails) SetAmazonShipmentDetails(v AmazonShipmentDetails) {
	o.AmazonShipmentDetails = &v
}

func (o ChannelDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channelType"] = o.ChannelType
	if !IsNil(o.AmazonOrderDetails) {
		toSerialize["amazonOrderDetails"] = o.AmazonOrderDetails
	}
	if !IsNil(o.AmazonShipmentDetails) {
		toSerialize["amazonShipmentDetails"] = o.AmazonShipmentDetails
	}
	return toSerialize, nil
}

type NullableChannelDetails struct {
	value *ChannelDetails
	isSet bool
}

func (v NullableChannelDetails) Get() *ChannelDetails {
	return v.value
}

func (v *NullableChannelDetails) Set(val *ChannelDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelDetails(val *ChannelDetails) *NullableChannelDetails {
	return &NullableChannelDetails{value: val, isSet: true}
}

func (v NullableChannelDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableChannelDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
