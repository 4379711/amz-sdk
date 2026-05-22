package catalog_items_20220401

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemSearchResults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemSearchResults{}

// ItemSearchResults Items in the Amazon catalog and search-related metadata.
type ItemSearchResults struct {
	// For searches that are based on `identifiers`, `numberOfResults` is the total number of Amazon catalog items found. For searches that are based on `keywords`, `numberOfResults` is the estimated total number of Amazon catalog items that are matched by the search query. Only results up to the page count limit are returned per request regardless of the number found.  **Note:** The maximum number of items (ASINs) that can be returned and paged through is 1,000.
	NumberOfResults int32       `json:"numberOfResults"`
	Pagination      Pagination  `json:"pagination"`
	Refinements     Refinements `json:"refinements"`
	// A list of items from the Amazon catalog.
	Items []Item `json:"items"`
}

type _ItemSearchResults ItemSearchResults

// NewItemSearchResults instantiates a new ItemSearchResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemSearchResults(numberOfResults int32, pagination Pagination, refinements Refinements, items []Item) *ItemSearchResults {
	this := ItemSearchResults{}
	this.NumberOfResults = numberOfResults
	this.Pagination = pagination
	this.Refinements = refinements
	this.Items = items
	return &this
}

// NewItemSearchResultsWithDefaults instantiates a new ItemSearchResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemSearchResultsWithDefaults() *ItemSearchResults {
	this := ItemSearchResults{}
	return &this
}

// GetNumberOfResults returns the NumberOfResults field value
func (o *ItemSearchResults) GetNumberOfResults() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumberOfResults
}

// GetNumberOfResultsOk returns a tuple with the NumberOfResults field value
// and a boolean to check if the value has been set.
func (o *ItemSearchResults) GetNumberOfResultsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumberOfResults, true
}

// SetNumberOfResults sets field value
func (o *ItemSearchResults) SetNumberOfResults(v int32) {
	o.NumberOfResults = v
}

// GetPagination returns the Pagination field value
func (o *ItemSearchResults) GetPagination() Pagination {
	if o == nil {
		var ret Pagination
		return ret
	}

	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value
// and a boolean to check if the value has been set.
func (o *ItemSearchResults) GetPaginationOk() (*Pagination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pagination, true
}

// SetPagination sets field value
func (o *ItemSearchResults) SetPagination(v Pagination) {
	o.Pagination = v
}

// GetRefinements returns the Refinements field value
func (o *ItemSearchResults) GetRefinements() Refinements {
	if o == nil {
		var ret Refinements
		return ret
	}

	return o.Refinements
}

// GetRefinementsOk returns a tuple with the Refinements field value
// and a boolean to check if the value has been set.
func (o *ItemSearchResults) GetRefinementsOk() (*Refinements, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Refinements, true
}

// SetRefinements sets field value
func (o *ItemSearchResults) SetRefinements(v Refinements) {
	o.Refinements = v
}

// GetItems returns the Items field value
func (o *ItemSearchResults) GetItems() []Item {
	if o == nil {
		var ret []Item
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ItemSearchResults) GetItemsOk() ([]Item, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *ItemSearchResults) SetItems(v []Item) {
	o.Items = v
}

func (o ItemSearchResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["numberOfResults"] = o.NumberOfResults
	toSerialize["pagination"] = o.Pagination
	toSerialize["refinements"] = o.Refinements
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

type NullableItemSearchResults struct {
	value *ItemSearchResults
	isSet bool
}

func (v NullableItemSearchResults) Get() *ItemSearchResults {
	return v.value
}

func (v *NullableItemSearchResults) Set(val *ItemSearchResults) {
	v.value = val
	v.isSet = true
}

func (v NullableItemSearchResults) IsSet() bool {
	return v.isSet
}

func (v *NullableItemSearchResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemSearchResults(val *ItemSearchResults) *NullableItemSearchResults {
	return &NullableItemSearchResults{value: val, isSet: true}
}

func (v NullableItemSearchResults) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemSearchResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
