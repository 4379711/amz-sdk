package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the ReturnAuthorization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReturnAuthorization{}

// ReturnAuthorization Return authorization information for items accepted for return.
type ReturnAuthorization struct {
	// An identifier for the return authorization. This identifier associates return items with the return authorization used to return them.
	ReturnAuthorizationId string `json:"returnAuthorizationId"`
	// An identifier for the Amazon fulfillment center that the return items should be sent to.
	FulfillmentCenterId string  `json:"fulfillmentCenterId"`
	ReturnToAddress     Address `json:"returnToAddress"`
	// The return merchandise authorization (RMA) that Amazon needs to process the return.
	AmazonRmaId string `json:"amazonRmaId"`
	// A URL for a web page that contains the return authorization barcode and the mailing label. This does not include pre-paid shipping.
	RmaPageURL string `json:"rmaPageURL"`
}

type _ReturnAuthorization ReturnAuthorization

// NewReturnAuthorization instantiates a new ReturnAuthorization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnAuthorization(returnAuthorizationId string, fulfillmentCenterId string, returnToAddress Address, amazonRmaId string, rmaPageURL string) *ReturnAuthorization {
	this := ReturnAuthorization{}
	this.ReturnAuthorizationId = returnAuthorizationId
	this.FulfillmentCenterId = fulfillmentCenterId
	this.ReturnToAddress = returnToAddress
	this.AmazonRmaId = amazonRmaId
	this.RmaPageURL = rmaPageURL
	return &this
}

// NewReturnAuthorizationWithDefaults instantiates a new ReturnAuthorization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnAuthorizationWithDefaults() *ReturnAuthorization {
	this := ReturnAuthorization{}
	return &this
}

// GetReturnAuthorizationId returns the ReturnAuthorizationId field value
func (o *ReturnAuthorization) GetReturnAuthorizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReturnAuthorizationId
}

// GetReturnAuthorizationIdOk returns a tuple with the ReturnAuthorizationId field value
// and a boolean to check if the value has been set.
func (o *ReturnAuthorization) GetReturnAuthorizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReturnAuthorizationId, true
}

// SetReturnAuthorizationId sets field value
func (o *ReturnAuthorization) SetReturnAuthorizationId(v string) {
	o.ReturnAuthorizationId = v
}

// GetFulfillmentCenterId returns the FulfillmentCenterId field value
func (o *ReturnAuthorization) GetFulfillmentCenterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FulfillmentCenterId
}

// GetFulfillmentCenterIdOk returns a tuple with the FulfillmentCenterId field value
// and a boolean to check if the value has been set.
func (o *ReturnAuthorization) GetFulfillmentCenterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FulfillmentCenterId, true
}

// SetFulfillmentCenterId sets field value
func (o *ReturnAuthorization) SetFulfillmentCenterId(v string) {
	o.FulfillmentCenterId = v
}

// GetReturnToAddress returns the ReturnToAddress field value
func (o *ReturnAuthorization) GetReturnToAddress() Address {
	if o == nil {
		var ret Address
		return ret
	}

	return o.ReturnToAddress
}

// GetReturnToAddressOk returns a tuple with the ReturnToAddress field value
// and a boolean to check if the value has been set.
func (o *ReturnAuthorization) GetReturnToAddressOk() (*Address, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReturnToAddress, true
}

// SetReturnToAddress sets field value
func (o *ReturnAuthorization) SetReturnToAddress(v Address) {
	o.ReturnToAddress = v
}

// GetAmazonRmaId returns the AmazonRmaId field value
func (o *ReturnAuthorization) GetAmazonRmaId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AmazonRmaId
}

// GetAmazonRmaIdOk returns a tuple with the AmazonRmaId field value
// and a boolean to check if the value has been set.
func (o *ReturnAuthorization) GetAmazonRmaIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmazonRmaId, true
}

// SetAmazonRmaId sets field value
func (o *ReturnAuthorization) SetAmazonRmaId(v string) {
	o.AmazonRmaId = v
}

// GetRmaPageURL returns the RmaPageURL field value
func (o *ReturnAuthorization) GetRmaPageURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RmaPageURL
}

// GetRmaPageURLOk returns a tuple with the RmaPageURL field value
// and a boolean to check if the value has been set.
func (o *ReturnAuthorization) GetRmaPageURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RmaPageURL, true
}

// SetRmaPageURL sets field value
func (o *ReturnAuthorization) SetRmaPageURL(v string) {
	o.RmaPageURL = v
}

func (o ReturnAuthorization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["returnAuthorizationId"] = o.ReturnAuthorizationId
	toSerialize["fulfillmentCenterId"] = o.FulfillmentCenterId
	toSerialize["returnToAddress"] = o.ReturnToAddress
	toSerialize["amazonRmaId"] = o.AmazonRmaId
	toSerialize["rmaPageURL"] = o.RmaPageURL
	return toSerialize, nil
}

type NullableReturnAuthorization struct {
	value *ReturnAuthorization
	isSet bool
}

func (v NullableReturnAuthorization) Get() *ReturnAuthorization {
	return v.value
}

func (v *NullableReturnAuthorization) Set(val *ReturnAuthorization) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnAuthorization) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnAuthorization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnAuthorization(val *ReturnAuthorization) *NullableReturnAuthorization {
	return &NullableReturnAuthorization{value: val, isSet: true}
}

func (v NullableReturnAuthorization) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableReturnAuthorization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
