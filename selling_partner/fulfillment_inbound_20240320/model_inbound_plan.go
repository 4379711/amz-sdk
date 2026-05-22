package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
	"time"
)

// checks if the InboundPlan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InboundPlan{}

// InboundPlan Inbound plan containing details of the inbound workflow.
type InboundPlan struct {
	// The time at which the inbound plan was created. In [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) datetime with pattern `yyyy-MM-ddTHH:mm:ssZ`.
	CreatedAt time.Time `json:"createdAt"`
	// Identifier of an inbound plan.
	InboundPlanId string `json:"inboundPlanId" validate:"regexp=^[a-zA-Z0-9-]*$"`
	// The time at which the inbound plan was last updated. In [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) datetime format with pattern `yyyy-MM-ddTHH:mm:ssZ`.
	LastUpdatedAt time.Time `json:"lastUpdatedAt"`
	// A list of marketplace IDs.
	MarketplaceIds []string `json:"marketplaceIds"`
	// Human-readable name of the inbound plan.
	Name string `json:"name"`
	// Packing options for the inbound plan. This property will be populated when it has been generated via the corresponding operation. If there is a chosen placement option, only packing options for that placement option will be returned. If there are confirmed shipments, only packing options for those shipments will be returned. Query the packing option for more details.
	PackingOptions []PackingOptionSummary `json:"packingOptions,omitempty"`
	// Placement options for the inbound plan. This property will be populated when it has been generated via the corresponding operation. If there is a chosen placement option, that will be the only returned option. Query the placement option for more details.
	PlacementOptions []PlacementOptionSummary `json:"placementOptions,omitempty"`
	// A list of shipment IDs for the inbound plan. This property is populated when it has been generated with the `confirmPlacementOptions` operation. Only shipments from the chosen placement option are returned. Query the shipment for more details.
	Shipments     []ShipmentSummary `json:"shipments,omitempty"`
	SourceAddress Address           `json:"sourceAddress"`
	// Current status of the inbound plan. Possible values: `ACTIVE`, `VOIDED`, `SHIPPED`, `ERRORED`.
	Status string `json:"status"`
}

type _InboundPlan InboundPlan

// NewInboundPlan instantiates a new InboundPlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInboundPlan(createdAt time.Time, inboundPlanId string, lastUpdatedAt time.Time, marketplaceIds []string, name string, sourceAddress Address, status string) *InboundPlan {
	this := InboundPlan{}
	this.CreatedAt = createdAt
	this.InboundPlanId = inboundPlanId
	this.LastUpdatedAt = lastUpdatedAt
	this.MarketplaceIds = marketplaceIds
	this.Name = name
	this.SourceAddress = sourceAddress
	this.Status = status
	return &this
}

// NewInboundPlanWithDefaults instantiates a new InboundPlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInboundPlanWithDefaults() *InboundPlan {
	this := InboundPlan{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *InboundPlan) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *InboundPlan) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *InboundPlan) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetInboundPlanId returns the InboundPlanId field value
func (o *InboundPlan) GetInboundPlanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InboundPlanId
}

// GetInboundPlanIdOk returns a tuple with the InboundPlanId field value
// and a boolean to check if the value has been set.
func (o *InboundPlan) GetInboundPlanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InboundPlanId, true
}

// SetInboundPlanId sets field value
func (o *InboundPlan) SetInboundPlanId(v string) {
	o.InboundPlanId = v
}

// GetLastUpdatedAt returns the LastUpdatedAt field value
func (o *InboundPlan) GetLastUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdatedAt
}

// GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field value
// and a boolean to check if the value has been set.
func (o *InboundPlan) GetLastUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedAt, true
}

// SetLastUpdatedAt sets field value
func (o *InboundPlan) SetLastUpdatedAt(v time.Time) {
	o.LastUpdatedAt = v
}

// GetMarketplaceIds returns the MarketplaceIds field value
func (o *InboundPlan) GetMarketplaceIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.MarketplaceIds
}

// GetMarketplaceIdsOk returns a tuple with the MarketplaceIds field value
// and a boolean to check if the value has been set.
func (o *InboundPlan) GetMarketplaceIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MarketplaceIds, true
}

// SetMarketplaceIds sets field value
func (o *InboundPlan) SetMarketplaceIds(v []string) {
	o.MarketplaceIds = v
}

// GetName returns the Name field value
func (o *InboundPlan) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InboundPlan) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InboundPlan) SetName(v string) {
	o.Name = v
}

// GetPackingOptions returns the PackingOptions field value if set, zero value otherwise.
func (o *InboundPlan) GetPackingOptions() []PackingOptionSummary {
	if o == nil || IsNil(o.PackingOptions) {
		var ret []PackingOptionSummary
		return ret
	}
	return o.PackingOptions
}

// GetPackingOptionsOk returns a tuple with the PackingOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundPlan) GetPackingOptionsOk() ([]PackingOptionSummary, bool) {
	if o == nil || IsNil(o.PackingOptions) {
		return nil, false
	}
	return o.PackingOptions, true
}

// HasPackingOptions returns a boolean if a field has been set.
func (o *InboundPlan) HasPackingOptions() bool {
	if o != nil && !IsNil(o.PackingOptions) {
		return true
	}

	return false
}

// SetPackingOptions gets a reference to the given []PackingOptionSummary and assigns it to the PackingOptions field.
func (o *InboundPlan) SetPackingOptions(v []PackingOptionSummary) {
	o.PackingOptions = v
}

// GetPlacementOptions returns the PlacementOptions field value if set, zero value otherwise.
func (o *InboundPlan) GetPlacementOptions() []PlacementOptionSummary {
	if o == nil || IsNil(o.PlacementOptions) {
		var ret []PlacementOptionSummary
		return ret
	}
	return o.PlacementOptions
}

// GetPlacementOptionsOk returns a tuple with the PlacementOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundPlan) GetPlacementOptionsOk() ([]PlacementOptionSummary, bool) {
	if o == nil || IsNil(o.PlacementOptions) {
		return nil, false
	}
	return o.PlacementOptions, true
}

// HasPlacementOptions returns a boolean if a field has been set.
func (o *InboundPlan) HasPlacementOptions() bool {
	if o != nil && !IsNil(o.PlacementOptions) {
		return true
	}

	return false
}

// SetPlacementOptions gets a reference to the given []PlacementOptionSummary and assigns it to the PlacementOptions field.
func (o *InboundPlan) SetPlacementOptions(v []PlacementOptionSummary) {
	o.PlacementOptions = v
}

// GetShipments returns the Shipments field value if set, zero value otherwise.
func (o *InboundPlan) GetShipments() []ShipmentSummary {
	if o == nil || IsNil(o.Shipments) {
		var ret []ShipmentSummary
		return ret
	}
	return o.Shipments
}

// GetShipmentsOk returns a tuple with the Shipments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundPlan) GetShipmentsOk() ([]ShipmentSummary, bool) {
	if o == nil || IsNil(o.Shipments) {
		return nil, false
	}
	return o.Shipments, true
}

// HasShipments returns a boolean if a field has been set.
func (o *InboundPlan) HasShipments() bool {
	if o != nil && !IsNil(o.Shipments) {
		return true
	}

	return false
}

// SetShipments gets a reference to the given []ShipmentSummary and assigns it to the Shipments field.
func (o *InboundPlan) SetShipments(v []ShipmentSummary) {
	o.Shipments = v
}

// GetSourceAddress returns the SourceAddress field value
func (o *InboundPlan) GetSourceAddress() Address {
	if o == nil {
		var ret Address
		return ret
	}

	return o.SourceAddress
}

// GetSourceAddressOk returns a tuple with the SourceAddress field value
// and a boolean to check if the value has been set.
func (o *InboundPlan) GetSourceAddressOk() (*Address, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceAddress, true
}

// SetSourceAddress sets field value
func (o *InboundPlan) SetSourceAddress(v Address) {
	o.SourceAddress = v
}

// GetStatus returns the Status field value
func (o *InboundPlan) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InboundPlan) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *InboundPlan) SetStatus(v string) {
	o.Status = v
}

func (o InboundPlan) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["inboundPlanId"] = o.InboundPlanId
	toSerialize["lastUpdatedAt"] = o.LastUpdatedAt
	toSerialize["marketplaceIds"] = o.MarketplaceIds
	toSerialize["name"] = o.Name
	if !IsNil(o.PackingOptions) {
		toSerialize["packingOptions"] = o.PackingOptions
	}
	if !IsNil(o.PlacementOptions) {
		toSerialize["placementOptions"] = o.PlacementOptions
	}
	if !IsNil(o.Shipments) {
		toSerialize["shipments"] = o.Shipments
	}
	toSerialize["sourceAddress"] = o.SourceAddress
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullableInboundPlan struct {
	value *InboundPlan
	isSet bool
}

func (v NullableInboundPlan) Get() *InboundPlan {
	return v.value
}

func (v *NullableInboundPlan) Set(val *InboundPlan) {
	v.value = val
	v.isSet = true
}

func (v NullableInboundPlan) IsSet() bool {
	return v.isSet
}

func (v *NullableInboundPlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInboundPlan(val *InboundPlan) *NullableInboundPlan {
	return &NullableInboundPlan{value: val, isSet: true}
}

func (v NullableInboundPlan) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableInboundPlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
