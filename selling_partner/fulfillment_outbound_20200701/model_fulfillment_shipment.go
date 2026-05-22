package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
	"time"
)

// checks if the FulfillmentShipment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FulfillmentShipment{}

// FulfillmentShipment Delivery and item information for a shipment in a fulfillment order.
type FulfillmentShipment struct {
	// A shipment identifier assigned by Amazon.
	AmazonShipmentId string `json:"amazonShipmentId"`
	// An identifier for the fulfillment center that the shipment will be sent from.
	FulfillmentCenterId string `json:"fulfillmentCenterId"`
	// The current status of the shipment.
	FulfillmentShipmentStatus string `json:"fulfillmentShipmentStatus"`
	// Date timestamp
	ShippingDate *time.Time `json:"shippingDate,omitempty"`
	// Date timestamp
	EstimatedArrivalDate *time.Time `json:"estimatedArrivalDate,omitempty"`
	// Provides additional insight into shipment timeline. Primairly used to communicate that actual delivery dates aren't available.
	ShippingNotes []string `json:"shippingNotes,omitempty"`
	// An array of fulfillment shipment item information.
	FulfillmentShipmentItem []FulfillmentShipmentItem `json:"fulfillmentShipmentItem"`
	// An array of fulfillment shipment package information.
	FulfillmentShipmentPackage []FulfillmentShipmentPackage `json:"fulfillmentShipmentPackage,omitempty"`
}

type _FulfillmentShipment FulfillmentShipment

// NewFulfillmentShipment instantiates a new FulfillmentShipment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFulfillmentShipment(amazonShipmentId string, fulfillmentCenterId string, fulfillmentShipmentStatus string, fulfillmentShipmentItem []FulfillmentShipmentItem) *FulfillmentShipment {
	this := FulfillmentShipment{}
	this.AmazonShipmentId = amazonShipmentId
	this.FulfillmentCenterId = fulfillmentCenterId
	this.FulfillmentShipmentStatus = fulfillmentShipmentStatus
	this.FulfillmentShipmentItem = fulfillmentShipmentItem
	return &this
}

// NewFulfillmentShipmentWithDefaults instantiates a new FulfillmentShipment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFulfillmentShipmentWithDefaults() *FulfillmentShipment {
	this := FulfillmentShipment{}
	return &this
}

// GetAmazonShipmentId returns the AmazonShipmentId field value
func (o *FulfillmentShipment) GetAmazonShipmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AmazonShipmentId
}

// GetAmazonShipmentIdOk returns a tuple with the AmazonShipmentId field value
// and a boolean to check if the value has been set.
func (o *FulfillmentShipment) GetAmazonShipmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmazonShipmentId, true
}

// SetAmazonShipmentId sets field value
func (o *FulfillmentShipment) SetAmazonShipmentId(v string) {
	o.AmazonShipmentId = v
}

// GetFulfillmentCenterId returns the FulfillmentCenterId field value
func (o *FulfillmentShipment) GetFulfillmentCenterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FulfillmentCenterId
}

// GetFulfillmentCenterIdOk returns a tuple with the FulfillmentCenterId field value
// and a boolean to check if the value has been set.
func (o *FulfillmentShipment) GetFulfillmentCenterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FulfillmentCenterId, true
}

// SetFulfillmentCenterId sets field value
func (o *FulfillmentShipment) SetFulfillmentCenterId(v string) {
	o.FulfillmentCenterId = v
}

// GetFulfillmentShipmentStatus returns the FulfillmentShipmentStatus field value
func (o *FulfillmentShipment) GetFulfillmentShipmentStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FulfillmentShipmentStatus
}

// GetFulfillmentShipmentStatusOk returns a tuple with the FulfillmentShipmentStatus field value
// and a boolean to check if the value has been set.
func (o *FulfillmentShipment) GetFulfillmentShipmentStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FulfillmentShipmentStatus, true
}

// SetFulfillmentShipmentStatus sets field value
func (o *FulfillmentShipment) SetFulfillmentShipmentStatus(v string) {
	o.FulfillmentShipmentStatus = v
}

// GetShippingDate returns the ShippingDate field value if set, zero value otherwise.
func (o *FulfillmentShipment) GetShippingDate() time.Time {
	if o == nil || IsNil(o.ShippingDate) {
		var ret time.Time
		return ret
	}
	return *o.ShippingDate
}

// GetShippingDateOk returns a tuple with the ShippingDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentShipment) GetShippingDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ShippingDate) {
		return nil, false
	}
	return o.ShippingDate, true
}

// HasShippingDate returns a boolean if a field has been set.
func (o *FulfillmentShipment) HasShippingDate() bool {
	if o != nil && !IsNil(o.ShippingDate) {
		return true
	}

	return false
}

// SetShippingDate gets a reference to the given time.Time and assigns it to the ShippingDate field.
func (o *FulfillmentShipment) SetShippingDate(v time.Time) {
	o.ShippingDate = &v
}

// GetEstimatedArrivalDate returns the EstimatedArrivalDate field value if set, zero value otherwise.
func (o *FulfillmentShipment) GetEstimatedArrivalDate() time.Time {
	if o == nil || IsNil(o.EstimatedArrivalDate) {
		var ret time.Time
		return ret
	}
	return *o.EstimatedArrivalDate
}

// GetEstimatedArrivalDateOk returns a tuple with the EstimatedArrivalDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentShipment) GetEstimatedArrivalDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EstimatedArrivalDate) {
		return nil, false
	}
	return o.EstimatedArrivalDate, true
}

// HasEstimatedArrivalDate returns a boolean if a field has been set.
func (o *FulfillmentShipment) HasEstimatedArrivalDate() bool {
	if o != nil && !IsNil(o.EstimatedArrivalDate) {
		return true
	}

	return false
}

// SetEstimatedArrivalDate gets a reference to the given time.Time and assigns it to the EstimatedArrivalDate field.
func (o *FulfillmentShipment) SetEstimatedArrivalDate(v time.Time) {
	o.EstimatedArrivalDate = &v
}

// GetShippingNotes returns the ShippingNotes field value if set, zero value otherwise.
func (o *FulfillmentShipment) GetShippingNotes() []string {
	if o == nil || IsNil(o.ShippingNotes) {
		var ret []string
		return ret
	}
	return o.ShippingNotes
}

// GetShippingNotesOk returns a tuple with the ShippingNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentShipment) GetShippingNotesOk() ([]string, bool) {
	if o == nil || IsNil(o.ShippingNotes) {
		return nil, false
	}
	return o.ShippingNotes, true
}

// HasShippingNotes returns a boolean if a field has been set.
func (o *FulfillmentShipment) HasShippingNotes() bool {
	if o != nil && !IsNil(o.ShippingNotes) {
		return true
	}

	return false
}

// SetShippingNotes gets a reference to the given []string and assigns it to the ShippingNotes field.
func (o *FulfillmentShipment) SetShippingNotes(v []string) {
	o.ShippingNotes = v
}

// GetFulfillmentShipmentItem returns the FulfillmentShipmentItem field value
func (o *FulfillmentShipment) GetFulfillmentShipmentItem() []FulfillmentShipmentItem {
	if o == nil {
		var ret []FulfillmentShipmentItem
		return ret
	}

	return o.FulfillmentShipmentItem
}

// GetFulfillmentShipmentItemOk returns a tuple with the FulfillmentShipmentItem field value
// and a boolean to check if the value has been set.
func (o *FulfillmentShipment) GetFulfillmentShipmentItemOk() ([]FulfillmentShipmentItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.FulfillmentShipmentItem, true
}

// SetFulfillmentShipmentItem sets field value
func (o *FulfillmentShipment) SetFulfillmentShipmentItem(v []FulfillmentShipmentItem) {
	o.FulfillmentShipmentItem = v
}

// GetFulfillmentShipmentPackage returns the FulfillmentShipmentPackage field value if set, zero value otherwise.
func (o *FulfillmentShipment) GetFulfillmentShipmentPackage() []FulfillmentShipmentPackage {
	if o == nil || IsNil(o.FulfillmentShipmentPackage) {
		var ret []FulfillmentShipmentPackage
		return ret
	}
	return o.FulfillmentShipmentPackage
}

// GetFulfillmentShipmentPackageOk returns a tuple with the FulfillmentShipmentPackage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentShipment) GetFulfillmentShipmentPackageOk() ([]FulfillmentShipmentPackage, bool) {
	if o == nil || IsNil(o.FulfillmentShipmentPackage) {
		return nil, false
	}
	return o.FulfillmentShipmentPackage, true
}

// HasFulfillmentShipmentPackage returns a boolean if a field has been set.
func (o *FulfillmentShipment) HasFulfillmentShipmentPackage() bool {
	if o != nil && !IsNil(o.FulfillmentShipmentPackage) {
		return true
	}

	return false
}

// SetFulfillmentShipmentPackage gets a reference to the given []FulfillmentShipmentPackage and assigns it to the FulfillmentShipmentPackage field.
func (o *FulfillmentShipment) SetFulfillmentShipmentPackage(v []FulfillmentShipmentPackage) {
	o.FulfillmentShipmentPackage = v
}

func (o FulfillmentShipment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amazonShipmentId"] = o.AmazonShipmentId
	toSerialize["fulfillmentCenterId"] = o.FulfillmentCenterId
	toSerialize["fulfillmentShipmentStatus"] = o.FulfillmentShipmentStatus
	if !IsNil(o.ShippingDate) {
		toSerialize["shippingDate"] = o.ShippingDate
	}
	if !IsNil(o.EstimatedArrivalDate) {
		toSerialize["estimatedArrivalDate"] = o.EstimatedArrivalDate
	}
	if !IsNil(o.ShippingNotes) {
		toSerialize["shippingNotes"] = o.ShippingNotes
	}
	toSerialize["fulfillmentShipmentItem"] = o.FulfillmentShipmentItem
	if !IsNil(o.FulfillmentShipmentPackage) {
		toSerialize["fulfillmentShipmentPackage"] = o.FulfillmentShipmentPackage
	}
	return toSerialize, nil
}

type NullableFulfillmentShipment struct {
	value *FulfillmentShipment
	isSet bool
}

func (v NullableFulfillmentShipment) Get() *FulfillmentShipment {
	return v.value
}

func (v *NullableFulfillmentShipment) Set(val *FulfillmentShipment) {
	v.value = val
	v.isSet = true
}

func (v NullableFulfillmentShipment) IsSet() bool {
	return v.isSet
}

func (v *NullableFulfillmentShipment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFulfillmentShipment(val *FulfillmentShipment) *NullableFulfillmentShipment {
	return &NullableFulfillmentShipment{value: val, isSet: true}
}

func (v NullableFulfillmentShipment) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableFulfillmentShipment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
