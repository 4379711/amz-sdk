package listings_items_20210801

import (
	"github.com/bytedance/sonic"
)

// checks if the IssueEnforcementAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IssueEnforcementAction{}

// IssueEnforcementAction The enforcement action taken by Amazon that affect the publishing or status of a listing
type IssueEnforcementAction struct {
	// The enforcement action name.   Possible values:   * `LISTING_SUPPRESSED` - This enforcement takes down the current listing item's buyability.   * `ATTRIBUTE_SUPPRESSED` - An attribute's value on the listing item is invalid, which causes it to be rejected by Amazon.   * `CATALOG_ITEM_REMOVED` - This catalog item is inactive on Amazon, and all offers against it in the applicable marketplace are non-buyable.   * `SEARCH_SUPPRESSED` - This value indicates that the catalog item is hidden from search results.
	Action string `json:"action"`
}

type _IssueEnforcementAction IssueEnforcementAction

// NewIssueEnforcementAction instantiates a new IssueEnforcementAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssueEnforcementAction(action string) *IssueEnforcementAction {
	this := IssueEnforcementAction{}
	this.Action = action
	return &this
}

// NewIssueEnforcementActionWithDefaults instantiates a new IssueEnforcementAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssueEnforcementActionWithDefaults() *IssueEnforcementAction {
	this := IssueEnforcementAction{}
	return &this
}

// GetAction returns the Action field value
func (o *IssueEnforcementAction) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *IssueEnforcementAction) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *IssueEnforcementAction) SetAction(v string) {
	o.Action = v
}

func (o IssueEnforcementAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	return toSerialize, nil
}

type NullableIssueEnforcementAction struct {
	value *IssueEnforcementAction
	isSet bool
}

func (v NullableIssueEnforcementAction) Get() *IssueEnforcementAction {
	return v.value
}

func (v *NullableIssueEnforcementAction) Set(val *IssueEnforcementAction) {
	v.value = val
	v.isSet = true
}

func (v NullableIssueEnforcementAction) IsSet() bool {
	return v.isSet
}

func (v *NullableIssueEnforcementAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssueEnforcementAction(val *IssueEnforcementAction) *NullableIssueEnforcementAction {
	return &NullableIssueEnforcementAction{value: val, isSet: true}
}

func (v NullableIssueEnforcementAction) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableIssueEnforcementAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
