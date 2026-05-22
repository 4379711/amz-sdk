package awd_20240509

import (
	"github.com/bytedance/sonic"
)

// checks if the InventoryListing type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InventoryListing{}

// InventoryListing AWD inventory payload.
type InventoryListing struct {
	// List of inventory summaries.
	Inventory []InventorySummary `json:"inventory"`
	// A token that is used to retrieve the next page of results. The response includes `nextToken` when the number of results exceeds the specified `maxResults` value. To get the next page of results, call the operation with this token and include the same arguments as the call that produced the token. To get a complete list, call this operation until `nextToken` is null. Note that this operation can return empty pages.
	NextToken *string `json:"nextToken,omitempty"`
}

type _InventoryListing InventoryListing

// NewInventoryListing instantiates a new InventoryListing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInventoryListing(inventory []InventorySummary) *InventoryListing {
	this := InventoryListing{}
	this.Inventory = inventory
	return &this
}

// NewInventoryListingWithDefaults instantiates a new InventoryListing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInventoryListingWithDefaults() *InventoryListing {
	this := InventoryListing{}
	return &this
}

// GetInventory returns the Inventory field value
func (o *InventoryListing) GetInventory() []InventorySummary {
	if o == nil {
		var ret []InventorySummary
		return ret
	}

	return o.Inventory
}

// GetInventoryOk returns a tuple with the Inventory field value
// and a boolean to check if the value has been set.
func (o *InventoryListing) GetInventoryOk() ([]InventorySummary, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inventory, true
}

// SetInventory sets field value
func (o *InventoryListing) SetInventory(v []InventorySummary) {
	o.Inventory = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *InventoryListing) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InventoryListing) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *InventoryListing) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *InventoryListing) SetNextToken(v string) {
	o.NextToken = &v
}

func (o InventoryListing) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["inventory"] = o.Inventory
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	return toSerialize, nil
}

type NullableInventoryListing struct {
	value *InventoryListing
	isSet bool
}

func (v NullableInventoryListing) Get() *InventoryListing {
	return v.value
}

func (v *NullableInventoryListing) Set(val *InventoryListing) {
	v.value = val
	v.isSet = true
}

func (v NullableInventoryListing) IsSet() bool {
	return v.isSet
}

func (v *NullableInventoryListing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInventoryListing(val *InventoryListing) *NullableInventoryListing {
	return &NullableInventoryListing{value: val, isSet: true}
}

func (v NullableInventoryListing) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableInventoryListing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
