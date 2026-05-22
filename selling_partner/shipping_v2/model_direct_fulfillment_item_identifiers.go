package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the DirectFulfillmentItemIdentifiers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DirectFulfillmentItemIdentifiers{}

// DirectFulfillmentItemIdentifiers Item identifiers for an item in a direct fulfillment shipment.
type DirectFulfillmentItemIdentifiers struct {
	// A unique identifier for an item provided by the client for a direct fulfillment shipment. This is only populated for direct fulfillment multi-piece shipments. It is required if a vendor wants to change the configuration of the packages in which the purchase order is shipped.
	LineItemID string `json:"lineItemID"`
	// A unique identifier for an item provided by the client for a direct fulfillment shipment. This is only populated if a single line item has multiple pieces. Defaults to 1.
	PieceNumber *string `json:"pieceNumber,omitempty"`
}

type _DirectFulfillmentItemIdentifiers DirectFulfillmentItemIdentifiers

// NewDirectFulfillmentItemIdentifiers instantiates a new DirectFulfillmentItemIdentifiers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDirectFulfillmentItemIdentifiers(lineItemID string) *DirectFulfillmentItemIdentifiers {
	this := DirectFulfillmentItemIdentifiers{}
	this.LineItemID = lineItemID
	return &this
}

// NewDirectFulfillmentItemIdentifiersWithDefaults instantiates a new DirectFulfillmentItemIdentifiers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDirectFulfillmentItemIdentifiersWithDefaults() *DirectFulfillmentItemIdentifiers {
	this := DirectFulfillmentItemIdentifiers{}
	return &this
}

// GetLineItemID returns the LineItemID field value
func (o *DirectFulfillmentItemIdentifiers) GetLineItemID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LineItemID
}

// GetLineItemIDOk returns a tuple with the LineItemID field value
// and a boolean to check if the value has been set.
func (o *DirectFulfillmentItemIdentifiers) GetLineItemIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LineItemID, true
}

// SetLineItemID sets field value
func (o *DirectFulfillmentItemIdentifiers) SetLineItemID(v string) {
	o.LineItemID = v
}

// GetPieceNumber returns the PieceNumber field value if set, zero value otherwise.
func (o *DirectFulfillmentItemIdentifiers) GetPieceNumber() string {
	if o == nil || IsNil(o.PieceNumber) {
		var ret string
		return ret
	}
	return *o.PieceNumber
}

// GetPieceNumberOk returns a tuple with the PieceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectFulfillmentItemIdentifiers) GetPieceNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PieceNumber) {
		return nil, false
	}
	return o.PieceNumber, true
}

// HasPieceNumber returns a boolean if a field has been set.
func (o *DirectFulfillmentItemIdentifiers) HasPieceNumber() bool {
	if o != nil && !IsNil(o.PieceNumber) {
		return true
	}

	return false
}

// SetPieceNumber gets a reference to the given string and assigns it to the PieceNumber field.
func (o *DirectFulfillmentItemIdentifiers) SetPieceNumber(v string) {
	o.PieceNumber = &v
}

func (o DirectFulfillmentItemIdentifiers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["lineItemID"] = o.LineItemID
	if !IsNil(o.PieceNumber) {
		toSerialize["pieceNumber"] = o.PieceNumber
	}
	return toSerialize, nil
}

type NullableDirectFulfillmentItemIdentifiers struct {
	value *DirectFulfillmentItemIdentifiers
	isSet bool
}

func (v NullableDirectFulfillmentItemIdentifiers) Get() *DirectFulfillmentItemIdentifiers {
	return v.value
}

func (v *NullableDirectFulfillmentItemIdentifiers) Set(val *DirectFulfillmentItemIdentifiers) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectFulfillmentItemIdentifiers) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectFulfillmentItemIdentifiers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectFulfillmentItemIdentifiers(val *DirectFulfillmentItemIdentifiers) *NullableDirectFulfillmentItemIdentifiers {
	return &NullableDirectFulfillmentItemIdentifiers{value: val, isSet: true}
}

func (v NullableDirectFulfillmentItemIdentifiers) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDirectFulfillmentItemIdentifiers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
