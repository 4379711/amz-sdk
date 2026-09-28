package promotions_20251201

import (
	"github.com/bytedance/sonic"
)

// checks if the SelectionDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SelectionDetails{}

// SelectionDetails Detailed selection information including items, rules, issues, and pagination.
type SelectionDetails struct {
	Rules *SelectionRules `json:"rules,omitempty"`
	// Item-level validation issues for items in this selection. Only present when `ISSUES` is included in the `includedData` query parameter.
	Issues []ItemIssue `json:"issues,omitempty"`
	// List of specific items to include. Only valid when `type` is `ITEMS` (maximum 100 items).
	Items      []Item      `json:"items,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

// NewSelectionDetails instantiates a new SelectionDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelectionDetails() *SelectionDetails {
	this := SelectionDetails{}
	return &this
}

// NewSelectionDetailsWithDefaults instantiates a new SelectionDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelectionDetailsWithDefaults() *SelectionDetails {
	this := SelectionDetails{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *SelectionDetails) GetRules() SelectionRules {
	if o == nil || IsNil(o.Rules) {
		var ret SelectionRules
		return ret
	}
	return *o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectionDetails) GetRulesOk() (*SelectionRules, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *SelectionDetails) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given SelectionRules and assigns it to the Rules field.
func (o *SelectionDetails) SetRules(v SelectionRules) {
	o.Rules = &v
}

// GetIssues returns the Issues field value if set, zero value otherwise.
func (o *SelectionDetails) GetIssues() []ItemIssue {
	if o == nil || IsNil(o.Issues) {
		var ret []ItemIssue
		return ret
	}
	return o.Issues
}

// GetIssuesOk returns a tuple with the Issues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectionDetails) GetIssuesOk() ([]ItemIssue, bool) {
	if o == nil || IsNil(o.Issues) {
		return nil, false
	}
	return o.Issues, true
}

// HasIssues returns a boolean if a field has been set.
func (o *SelectionDetails) HasIssues() bool {
	if o != nil && !IsNil(o.Issues) {
		return true
	}

	return false
}

// SetIssues gets a reference to the given []ItemIssue and assigns it to the Issues field.
func (o *SelectionDetails) SetIssues(v []ItemIssue) {
	o.Issues = v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SelectionDetails) GetItems() []Item {
	if o == nil || IsNil(o.Items) {
		var ret []Item
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectionDetails) GetItemsOk() ([]Item, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SelectionDetails) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Item and assigns it to the Items field.
func (o *SelectionDetails) SetItems(v []Item) {
	o.Items = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *SelectionDetails) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelectionDetails) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *SelectionDetails) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *SelectionDetails) SetPagination(v Pagination) {
	o.Pagination = &v
}

func (o SelectionDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	if !IsNil(o.Issues) {
		toSerialize["issues"] = o.Issues
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

type NullableSelectionDetails struct {
	value *SelectionDetails
	isSet bool
}

func (v NullableSelectionDetails) Get() *SelectionDetails {
	return v.value
}

func (v *NullableSelectionDetails) Set(val *SelectionDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableSelectionDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableSelectionDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelectionDetails(val *SelectionDetails) *NullableSelectionDetails {
	return &NullableSelectionDetails{value: val, isSet: true}
}

func (v NullableSelectionDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSelectionDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
