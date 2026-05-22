package finances_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the ServiceFeeEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceFeeEvent{}

// ServiceFeeEvent A service fee on the seller's account.
type ServiceFeeEvent struct {
	// An Amazon-defined identifier for an order.
	AmazonOrderId *string `json:"AmazonOrderId,omitempty"`
	// A short description of the service fee reason.
	FeeReason *string `json:"FeeReason,omitempty"`
	// A list of fee component information.
	FeeList []FeeComponent `json:"FeeList,omitempty"`
	// The seller SKU of the item. The seller SKU is qualified by the seller's seller ID, which is included with every call to the Selling Partner API.
	SellerSKU *string `json:"SellerSKU,omitempty"`
	// A unique identifier assigned by Amazon to products stored in and fulfilled from an Amazon fulfillment center.
	FnSKU *string `json:"FnSKU,omitempty"`
	// A short description of the service fee event.
	FeeDescription *string `json:"FeeDescription,omitempty"`
	// The Amazon Standard Identification Number (ASIN) of the item.
	ASIN *string `json:"ASIN,omitempty"`
	// The name of the store where the event occurred.
	StoreName *string `json:"StoreName,omitempty"`
}

// NewServiceFeeEvent instantiates a new ServiceFeeEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceFeeEvent() *ServiceFeeEvent {
	this := ServiceFeeEvent{}
	return &this
}

// NewServiceFeeEventWithDefaults instantiates a new ServiceFeeEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceFeeEventWithDefaults() *ServiceFeeEvent {
	this := ServiceFeeEvent{}
	return &this
}

// GetAmazonOrderId returns the AmazonOrderId field value if set, zero value otherwise.
func (o *ServiceFeeEvent) GetAmazonOrderId() string {
	if o == nil || IsNil(o.AmazonOrderId) {
		var ret string
		return ret
	}
	return *o.AmazonOrderId
}

// GetAmazonOrderIdOk returns a tuple with the AmazonOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceFeeEvent) GetAmazonOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.AmazonOrderId) {
		return nil, false
	}
	return o.AmazonOrderId, true
}

// HasAmazonOrderId returns a boolean if a field has been set.
func (o *ServiceFeeEvent) HasAmazonOrderId() bool {
	if o != nil && !IsNil(o.AmazonOrderId) {
		return true
	}

	return false
}

// SetAmazonOrderId gets a reference to the given string and assigns it to the AmazonOrderId field.
func (o *ServiceFeeEvent) SetAmazonOrderId(v string) {
	o.AmazonOrderId = &v
}

// GetFeeReason returns the FeeReason field value if set, zero value otherwise.
func (o *ServiceFeeEvent) GetFeeReason() string {
	if o == nil || IsNil(o.FeeReason) {
		var ret string
		return ret
	}
	return *o.FeeReason
}

// GetFeeReasonOk returns a tuple with the FeeReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceFeeEvent) GetFeeReasonOk() (*string, bool) {
	if o == nil || IsNil(o.FeeReason) {
		return nil, false
	}
	return o.FeeReason, true
}

// HasFeeReason returns a boolean if a field has been set.
func (o *ServiceFeeEvent) HasFeeReason() bool {
	if o != nil && !IsNil(o.FeeReason) {
		return true
	}

	return false
}

// SetFeeReason gets a reference to the given string and assigns it to the FeeReason field.
func (o *ServiceFeeEvent) SetFeeReason(v string) {
	o.FeeReason = &v
}

// GetFeeList returns the FeeList field value if set, zero value otherwise.
func (o *ServiceFeeEvent) GetFeeList() []FeeComponent {
	if o == nil || IsNil(o.FeeList) {
		var ret []FeeComponent
		return ret
	}
	return o.FeeList
}

// GetFeeListOk returns a tuple with the FeeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceFeeEvent) GetFeeListOk() ([]FeeComponent, bool) {
	if o == nil || IsNil(o.FeeList) {
		return nil, false
	}
	return o.FeeList, true
}

// HasFeeList returns a boolean if a field has been set.
func (o *ServiceFeeEvent) HasFeeList() bool {
	if o != nil && !IsNil(o.FeeList) {
		return true
	}

	return false
}

// SetFeeList gets a reference to the given []FeeComponent and assigns it to the FeeList field.
func (o *ServiceFeeEvent) SetFeeList(v []FeeComponent) {
	o.FeeList = v
}

// GetSellerSKU returns the SellerSKU field value if set, zero value otherwise.
func (o *ServiceFeeEvent) GetSellerSKU() string {
	if o == nil || IsNil(o.SellerSKU) {
		var ret string
		return ret
	}
	return *o.SellerSKU
}

// GetSellerSKUOk returns a tuple with the SellerSKU field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceFeeEvent) GetSellerSKUOk() (*string, bool) {
	if o == nil || IsNil(o.SellerSKU) {
		return nil, false
	}
	return o.SellerSKU, true
}

// HasSellerSKU returns a boolean if a field has been set.
func (o *ServiceFeeEvent) HasSellerSKU() bool {
	if o != nil && !IsNil(o.SellerSKU) {
		return true
	}

	return false
}

// SetSellerSKU gets a reference to the given string and assigns it to the SellerSKU field.
func (o *ServiceFeeEvent) SetSellerSKU(v string) {
	o.SellerSKU = &v
}

// GetFnSKU returns the FnSKU field value if set, zero value otherwise.
func (o *ServiceFeeEvent) GetFnSKU() string {
	if o == nil || IsNil(o.FnSKU) {
		var ret string
		return ret
	}
	return *o.FnSKU
}

// GetFnSKUOk returns a tuple with the FnSKU field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceFeeEvent) GetFnSKUOk() (*string, bool) {
	if o == nil || IsNil(o.FnSKU) {
		return nil, false
	}
	return o.FnSKU, true
}

// HasFnSKU returns a boolean if a field has been set.
func (o *ServiceFeeEvent) HasFnSKU() bool {
	if o != nil && !IsNil(o.FnSKU) {
		return true
	}

	return false
}

// SetFnSKU gets a reference to the given string and assigns it to the FnSKU field.
func (o *ServiceFeeEvent) SetFnSKU(v string) {
	o.FnSKU = &v
}

// GetFeeDescription returns the FeeDescription field value if set, zero value otherwise.
func (o *ServiceFeeEvent) GetFeeDescription() string {
	if o == nil || IsNil(o.FeeDescription) {
		var ret string
		return ret
	}
	return *o.FeeDescription
}

// GetFeeDescriptionOk returns a tuple with the FeeDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceFeeEvent) GetFeeDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.FeeDescription) {
		return nil, false
	}
	return o.FeeDescription, true
}

// HasFeeDescription returns a boolean if a field has been set.
func (o *ServiceFeeEvent) HasFeeDescription() bool {
	if o != nil && !IsNil(o.FeeDescription) {
		return true
	}

	return false
}

// SetFeeDescription gets a reference to the given string and assigns it to the FeeDescription field.
func (o *ServiceFeeEvent) SetFeeDescription(v string) {
	o.FeeDescription = &v
}

// GetASIN returns the ASIN field value if set, zero value otherwise.
func (o *ServiceFeeEvent) GetASIN() string {
	if o == nil || IsNil(o.ASIN) {
		var ret string
		return ret
	}
	return *o.ASIN
}

// GetASINOk returns a tuple with the ASIN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceFeeEvent) GetASINOk() (*string, bool) {
	if o == nil || IsNil(o.ASIN) {
		return nil, false
	}
	return o.ASIN, true
}

// HasASIN returns a boolean if a field has been set.
func (o *ServiceFeeEvent) HasASIN() bool {
	if o != nil && !IsNil(o.ASIN) {
		return true
	}

	return false
}

// SetASIN gets a reference to the given string and assigns it to the ASIN field.
func (o *ServiceFeeEvent) SetASIN(v string) {
	o.ASIN = &v
}

// GetStoreName returns the StoreName field value if set, zero value otherwise.
func (o *ServiceFeeEvent) GetStoreName() string {
	if o == nil || IsNil(o.StoreName) {
		var ret string
		return ret
	}
	return *o.StoreName
}

// GetStoreNameOk returns a tuple with the StoreName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceFeeEvent) GetStoreNameOk() (*string, bool) {
	if o == nil || IsNil(o.StoreName) {
		return nil, false
	}
	return o.StoreName, true
}

// HasStoreName returns a boolean if a field has been set.
func (o *ServiceFeeEvent) HasStoreName() bool {
	if o != nil && !IsNil(o.StoreName) {
		return true
	}

	return false
}

// SetStoreName gets a reference to the given string and assigns it to the StoreName field.
func (o *ServiceFeeEvent) SetStoreName(v string) {
	o.StoreName = &v
}

func (o ServiceFeeEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AmazonOrderId) {
		toSerialize["AmazonOrderId"] = o.AmazonOrderId
	}
	if !IsNil(o.FeeReason) {
		toSerialize["FeeReason"] = o.FeeReason
	}
	if !IsNil(o.FeeList) {
		toSerialize["FeeList"] = o.FeeList
	}
	if !IsNil(o.SellerSKU) {
		toSerialize["SellerSKU"] = o.SellerSKU
	}
	if !IsNil(o.FnSKU) {
		toSerialize["FnSKU"] = o.FnSKU
	}
	if !IsNil(o.FeeDescription) {
		toSerialize["FeeDescription"] = o.FeeDescription
	}
	if !IsNil(o.ASIN) {
		toSerialize["ASIN"] = o.ASIN
	}
	if !IsNil(o.StoreName) {
		toSerialize["StoreName"] = o.StoreName
	}
	return toSerialize, nil
}

type NullableServiceFeeEvent struct {
	value *ServiceFeeEvent
	isSet bool
}

func (v NullableServiceFeeEvent) Get() *ServiceFeeEvent {
	return v.value
}

func (v *NullableServiceFeeEvent) Set(val *ServiceFeeEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceFeeEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceFeeEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceFeeEvent(val *ServiceFeeEvent) *NullableServiceFeeEvent {
	return &NullableServiceFeeEvent{value: val, isSet: true}
}

func (v NullableServiceFeeEvent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableServiceFeeEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
