package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the FulfillmentPreview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FulfillmentPreview{}

// FulfillmentPreview Information about a fulfillment order preview, including delivery and fee information based on shipping method.
type FulfillmentPreview struct {
	ShippingSpeedCategory ShippingSpeedCategory  `json:"shippingSpeedCategory"`
	ScheduledDeliveryInfo *ScheduledDeliveryInfo `json:"scheduledDeliveryInfo,omitempty"`
	// When true, this fulfillment order preview is fulfillable.
	IsFulfillable bool `json:"isFulfillable"`
	// When true, this fulfillment order preview is for COD (Cash On Delivery).
	IsCODCapable            bool    `json:"isCODCapable"`
	EstimatedShippingWeight *Weight `json:"estimatedShippingWeight,omitempty"`
	// An array of fee type and cost pairs.
	EstimatedFees []Fee `json:"estimatedFees,omitempty"`
	// An array of fulfillment preview shipment information.
	FulfillmentPreviewShipments []FulfillmentPreviewShipment `json:"fulfillmentPreviewShipments,omitempty"`
	// An array of unfulfillable preview item information.
	UnfulfillablePreviewItems []UnfulfillablePreviewItem `json:"unfulfillablePreviewItems,omitempty"`
	// String list
	OrderUnfulfillableReasons []string `json:"orderUnfulfillableReasons,omitempty"`
	// The marketplace the fulfillment order is placed against.
	MarketplaceId string `json:"marketplaceId"`
	// A list of features and their fulfillment policies to apply to the order.
	FeatureConstraints []FeatureSettings `json:"featureConstraints,omitempty"`
}

type _FulfillmentPreview FulfillmentPreview

// NewFulfillmentPreview instantiates a new FulfillmentPreview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFulfillmentPreview(shippingSpeedCategory ShippingSpeedCategory, isFulfillable bool, isCODCapable bool, marketplaceId string) *FulfillmentPreview {
	this := FulfillmentPreview{}
	this.ShippingSpeedCategory = shippingSpeedCategory
	this.IsFulfillable = isFulfillable
	this.IsCODCapable = isCODCapable
	this.MarketplaceId = marketplaceId
	return &this
}

// NewFulfillmentPreviewWithDefaults instantiates a new FulfillmentPreview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFulfillmentPreviewWithDefaults() *FulfillmentPreview {
	this := FulfillmentPreview{}
	return &this
}

// GetShippingSpeedCategory returns the ShippingSpeedCategory field value
func (o *FulfillmentPreview) GetShippingSpeedCategory() ShippingSpeedCategory {
	if o == nil {
		var ret ShippingSpeedCategory
		return ret
	}

	return o.ShippingSpeedCategory
}

// GetShippingSpeedCategoryOk returns a tuple with the ShippingSpeedCategory field value
// and a boolean to check if the value has been set.
func (o *FulfillmentPreview) GetShippingSpeedCategoryOk() (*ShippingSpeedCategory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShippingSpeedCategory, true
}

// SetShippingSpeedCategory sets field value
func (o *FulfillmentPreview) SetShippingSpeedCategory(v ShippingSpeedCategory) {
	o.ShippingSpeedCategory = v
}

// GetScheduledDeliveryInfo returns the ScheduledDeliveryInfo field value if set, zero value otherwise.
func (o *FulfillmentPreview) GetScheduledDeliveryInfo() ScheduledDeliveryInfo {
	if o == nil || IsNil(o.ScheduledDeliveryInfo) {
		var ret ScheduledDeliveryInfo
		return ret
	}
	return *o.ScheduledDeliveryInfo
}

// GetScheduledDeliveryInfoOk returns a tuple with the ScheduledDeliveryInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentPreview) GetScheduledDeliveryInfoOk() (*ScheduledDeliveryInfo, bool) {
	if o == nil || IsNil(o.ScheduledDeliveryInfo) {
		return nil, false
	}
	return o.ScheduledDeliveryInfo, true
}

// HasScheduledDeliveryInfo returns a boolean if a field has been set.
func (o *FulfillmentPreview) HasScheduledDeliveryInfo() bool {
	if o != nil && !IsNil(o.ScheduledDeliveryInfo) {
		return true
	}

	return false
}

// SetScheduledDeliveryInfo gets a reference to the given ScheduledDeliveryInfo and assigns it to the ScheduledDeliveryInfo field.
func (o *FulfillmentPreview) SetScheduledDeliveryInfo(v ScheduledDeliveryInfo) {
	o.ScheduledDeliveryInfo = &v
}

// GetIsFulfillable returns the IsFulfillable field value
func (o *FulfillmentPreview) GetIsFulfillable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsFulfillable
}

// GetIsFulfillableOk returns a tuple with the IsFulfillable field value
// and a boolean to check if the value has been set.
func (o *FulfillmentPreview) GetIsFulfillableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsFulfillable, true
}

// SetIsFulfillable sets field value
func (o *FulfillmentPreview) SetIsFulfillable(v bool) {
	o.IsFulfillable = v
}

// GetIsCODCapable returns the IsCODCapable field value
func (o *FulfillmentPreview) GetIsCODCapable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsCODCapable
}

// GetIsCODCapableOk returns a tuple with the IsCODCapable field value
// and a boolean to check if the value has been set.
func (o *FulfillmentPreview) GetIsCODCapableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsCODCapable, true
}

// SetIsCODCapable sets field value
func (o *FulfillmentPreview) SetIsCODCapable(v bool) {
	o.IsCODCapable = v
}

// GetEstimatedShippingWeight returns the EstimatedShippingWeight field value if set, zero value otherwise.
func (o *FulfillmentPreview) GetEstimatedShippingWeight() Weight {
	if o == nil || IsNil(o.EstimatedShippingWeight) {
		var ret Weight
		return ret
	}
	return *o.EstimatedShippingWeight
}

// GetEstimatedShippingWeightOk returns a tuple with the EstimatedShippingWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentPreview) GetEstimatedShippingWeightOk() (*Weight, bool) {
	if o == nil || IsNil(o.EstimatedShippingWeight) {
		return nil, false
	}
	return o.EstimatedShippingWeight, true
}

// HasEstimatedShippingWeight returns a boolean if a field has been set.
func (o *FulfillmentPreview) HasEstimatedShippingWeight() bool {
	if o != nil && !IsNil(o.EstimatedShippingWeight) {
		return true
	}

	return false
}

// SetEstimatedShippingWeight gets a reference to the given Weight and assigns it to the EstimatedShippingWeight field.
func (o *FulfillmentPreview) SetEstimatedShippingWeight(v Weight) {
	o.EstimatedShippingWeight = &v
}

// GetEstimatedFees returns the EstimatedFees field value if set, zero value otherwise.
func (o *FulfillmentPreview) GetEstimatedFees() []Fee {
	if o == nil || IsNil(o.EstimatedFees) {
		var ret []Fee
		return ret
	}
	return o.EstimatedFees
}

// GetEstimatedFeesOk returns a tuple with the EstimatedFees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentPreview) GetEstimatedFeesOk() ([]Fee, bool) {
	if o == nil || IsNil(o.EstimatedFees) {
		return nil, false
	}
	return o.EstimatedFees, true
}

// HasEstimatedFees returns a boolean if a field has been set.
func (o *FulfillmentPreview) HasEstimatedFees() bool {
	if o != nil && !IsNil(o.EstimatedFees) {
		return true
	}

	return false
}

// SetEstimatedFees gets a reference to the given []Fee and assigns it to the EstimatedFees field.
func (o *FulfillmentPreview) SetEstimatedFees(v []Fee) {
	o.EstimatedFees = v
}

// GetFulfillmentPreviewShipments returns the FulfillmentPreviewShipments field value if set, zero value otherwise.
func (o *FulfillmentPreview) GetFulfillmentPreviewShipments() []FulfillmentPreviewShipment {
	if o == nil || IsNil(o.FulfillmentPreviewShipments) {
		var ret []FulfillmentPreviewShipment
		return ret
	}
	return o.FulfillmentPreviewShipments
}

// GetFulfillmentPreviewShipmentsOk returns a tuple with the FulfillmentPreviewShipments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentPreview) GetFulfillmentPreviewShipmentsOk() ([]FulfillmentPreviewShipment, bool) {
	if o == nil || IsNil(o.FulfillmentPreviewShipments) {
		return nil, false
	}
	return o.FulfillmentPreviewShipments, true
}

// HasFulfillmentPreviewShipments returns a boolean if a field has been set.
func (o *FulfillmentPreview) HasFulfillmentPreviewShipments() bool {
	if o != nil && !IsNil(o.FulfillmentPreviewShipments) {
		return true
	}

	return false
}

// SetFulfillmentPreviewShipments gets a reference to the given []FulfillmentPreviewShipment and assigns it to the FulfillmentPreviewShipments field.
func (o *FulfillmentPreview) SetFulfillmentPreviewShipments(v []FulfillmentPreviewShipment) {
	o.FulfillmentPreviewShipments = v
}

// GetUnfulfillablePreviewItems returns the UnfulfillablePreviewItems field value if set, zero value otherwise.
func (o *FulfillmentPreview) GetUnfulfillablePreviewItems() []UnfulfillablePreviewItem {
	if o == nil || IsNil(o.UnfulfillablePreviewItems) {
		var ret []UnfulfillablePreviewItem
		return ret
	}
	return o.UnfulfillablePreviewItems
}

// GetUnfulfillablePreviewItemsOk returns a tuple with the UnfulfillablePreviewItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentPreview) GetUnfulfillablePreviewItemsOk() ([]UnfulfillablePreviewItem, bool) {
	if o == nil || IsNil(o.UnfulfillablePreviewItems) {
		return nil, false
	}
	return o.UnfulfillablePreviewItems, true
}

// HasUnfulfillablePreviewItems returns a boolean if a field has been set.
func (o *FulfillmentPreview) HasUnfulfillablePreviewItems() bool {
	if o != nil && !IsNil(o.UnfulfillablePreviewItems) {
		return true
	}

	return false
}

// SetUnfulfillablePreviewItems gets a reference to the given []UnfulfillablePreviewItem and assigns it to the UnfulfillablePreviewItems field.
func (o *FulfillmentPreview) SetUnfulfillablePreviewItems(v []UnfulfillablePreviewItem) {
	o.UnfulfillablePreviewItems = v
}

// GetOrderUnfulfillableReasons returns the OrderUnfulfillableReasons field value if set, zero value otherwise.
func (o *FulfillmentPreview) GetOrderUnfulfillableReasons() []string {
	if o == nil || IsNil(o.OrderUnfulfillableReasons) {
		var ret []string
		return ret
	}
	return o.OrderUnfulfillableReasons
}

// GetOrderUnfulfillableReasonsOk returns a tuple with the OrderUnfulfillableReasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentPreview) GetOrderUnfulfillableReasonsOk() ([]string, bool) {
	if o == nil || IsNil(o.OrderUnfulfillableReasons) {
		return nil, false
	}
	return o.OrderUnfulfillableReasons, true
}

// HasOrderUnfulfillableReasons returns a boolean if a field has been set.
func (o *FulfillmentPreview) HasOrderUnfulfillableReasons() bool {
	if o != nil && !IsNil(o.OrderUnfulfillableReasons) {
		return true
	}

	return false
}

// SetOrderUnfulfillableReasons gets a reference to the given []string and assigns it to the OrderUnfulfillableReasons field.
func (o *FulfillmentPreview) SetOrderUnfulfillableReasons(v []string) {
	o.OrderUnfulfillableReasons = v
}

// GetMarketplaceId returns the MarketplaceId field value
func (o *FulfillmentPreview) GetMarketplaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value
// and a boolean to check if the value has been set.
func (o *FulfillmentPreview) GetMarketplaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarketplaceId, true
}

// SetMarketplaceId sets field value
func (o *FulfillmentPreview) SetMarketplaceId(v string) {
	o.MarketplaceId = v
}

// GetFeatureConstraints returns the FeatureConstraints field value if set, zero value otherwise.
func (o *FulfillmentPreview) GetFeatureConstraints() []FeatureSettings {
	if o == nil || IsNil(o.FeatureConstraints) {
		var ret []FeatureSettings
		return ret
	}
	return o.FeatureConstraints
}

// GetFeatureConstraintsOk returns a tuple with the FeatureConstraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentPreview) GetFeatureConstraintsOk() ([]FeatureSettings, bool) {
	if o == nil || IsNil(o.FeatureConstraints) {
		return nil, false
	}
	return o.FeatureConstraints, true
}

// HasFeatureConstraints returns a boolean if a field has been set.
func (o *FulfillmentPreview) HasFeatureConstraints() bool {
	if o != nil && !IsNil(o.FeatureConstraints) {
		return true
	}

	return false
}

// SetFeatureConstraints gets a reference to the given []FeatureSettings and assigns it to the FeatureConstraints field.
func (o *FulfillmentPreview) SetFeatureConstraints(v []FeatureSettings) {
	o.FeatureConstraints = v
}

func (o FulfillmentPreview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["shippingSpeedCategory"] = o.ShippingSpeedCategory
	if !IsNil(o.ScheduledDeliveryInfo) {
		toSerialize["scheduledDeliveryInfo"] = o.ScheduledDeliveryInfo
	}
	toSerialize["isFulfillable"] = o.IsFulfillable
	toSerialize["isCODCapable"] = o.IsCODCapable
	if !IsNil(o.EstimatedShippingWeight) {
		toSerialize["estimatedShippingWeight"] = o.EstimatedShippingWeight
	}
	if !IsNil(o.EstimatedFees) {
		toSerialize["estimatedFees"] = o.EstimatedFees
	}
	if !IsNil(o.FulfillmentPreviewShipments) {
		toSerialize["fulfillmentPreviewShipments"] = o.FulfillmentPreviewShipments
	}
	if !IsNil(o.UnfulfillablePreviewItems) {
		toSerialize["unfulfillablePreviewItems"] = o.UnfulfillablePreviewItems
	}
	if !IsNil(o.OrderUnfulfillableReasons) {
		toSerialize["orderUnfulfillableReasons"] = o.OrderUnfulfillableReasons
	}
	toSerialize["marketplaceId"] = o.MarketplaceId
	if !IsNil(o.FeatureConstraints) {
		toSerialize["featureConstraints"] = o.FeatureConstraints
	}
	return toSerialize, nil
}

type NullableFulfillmentPreview struct {
	value *FulfillmentPreview
	isSet bool
}

func (v NullableFulfillmentPreview) Get() *FulfillmentPreview {
	return v.value
}

func (v *NullableFulfillmentPreview) Set(val *FulfillmentPreview) {
	v.value = val
	v.isSet = true
}

func (v NullableFulfillmentPreview) IsSet() bool {
	return v.isSet
}

func (v *NullableFulfillmentPreview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFulfillmentPreview(val *FulfillmentPreview) *NullableFulfillmentPreview {
	return &NullableFulfillmentPreview{value: val, isSet: true}
}

func (v NullableFulfillmentPreview) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableFulfillmentPreview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
