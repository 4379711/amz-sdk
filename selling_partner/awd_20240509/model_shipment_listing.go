package awd_20240509

import (
	"github.com/bytedance/sonic"
)

// checks if the ShipmentListing type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipmentListing{}

// ShipmentListing A list of inbound shipment summaries filtered by the attributes specified in the request.
type ShipmentListing struct {
	// A token that is used to retrieve the next page of results. The response includes `nextToken` when the number of results exceeds the specified `maxResults` value. To get the next page of results, call the operation with this token and include the same arguments as the call that produced the token. To get a complete list, call this operation until `nextToken` is null. Note that this operation can return empty pages.
	NextToken *string `json:"nextToken,omitempty"`
	// List of inbound shipment summaries.
	Shipments []InboundShipmentSummary `json:"shipments,omitempty"`
}

// NewShipmentListing instantiates a new ShipmentListing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentListing() *ShipmentListing {
	this := ShipmentListing{}
	return &this
}

// NewShipmentListingWithDefaults instantiates a new ShipmentListing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentListingWithDefaults() *ShipmentListing {
	this := ShipmentListing{}
	return &this
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ShipmentListing) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentListing) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ShipmentListing) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ShipmentListing) SetNextToken(v string) {
	o.NextToken = &v
}

// GetShipments returns the Shipments field value if set, zero value otherwise.
func (o *ShipmentListing) GetShipments() []InboundShipmentSummary {
	if o == nil || IsNil(o.Shipments) {
		var ret []InboundShipmentSummary
		return ret
	}
	return o.Shipments
}

// GetShipmentsOk returns a tuple with the Shipments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentListing) GetShipmentsOk() ([]InboundShipmentSummary, bool) {
	if o == nil || IsNil(o.Shipments) {
		return nil, false
	}
	return o.Shipments, true
}

// HasShipments returns a boolean if a field has been set.
func (o *ShipmentListing) HasShipments() bool {
	if o != nil && !IsNil(o.Shipments) {
		return true
	}

	return false
}

// SetShipments gets a reference to the given []InboundShipmentSummary and assigns it to the Shipments field.
func (o *ShipmentListing) SetShipments(v []InboundShipmentSummary) {
	o.Shipments = v
}

func (o ShipmentListing) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	if !IsNil(o.Shipments) {
		toSerialize["shipments"] = o.Shipments
	}
	return toSerialize, nil
}

type NullableShipmentListing struct {
	value *ShipmentListing
	isSet bool
}

func (v NullableShipmentListing) Get() *ShipmentListing {
	return v.value
}

func (v *NullableShipmentListing) Set(val *ShipmentListing) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentListing) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentListing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentListing(val *ShipmentListing) *NullableShipmentListing {
	return &NullableShipmentListing{value: val, isSet: true}
}

func (v NullableShipmentListing) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableShipmentListing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
