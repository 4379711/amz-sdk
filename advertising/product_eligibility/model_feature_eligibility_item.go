package product_eligibility

import (
	"github.com/bytedance/sonic"
)

// checks if the FeatureEligibilityItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeatureEligibilityItem{}

// FeatureEligibilityItem struct for FeatureEligibilityItem
type FeatureEligibilityItem struct {
	// The marketplace to check the feature access in, will be obfuscated or not depending on the input
	MarketplaceId *string `json:"marketplaceId,omitempty"`
	// String identifier for the status.
	Reasons  []FeatureReasonItem `json:"reasons,omitempty"`
	Resource *Resource           `json:"resource,omitempty"`
	Action   *Action             `json:"action,omitempty"`
	// Boolean value where if true, advertiser is eligible to access the given feature.
	IsEligible *bool `json:"isEligible,omitempty"`
	// The id of the item that is Marketplace + Resource + Action
	ItemRequestId *string `json:"itemRequestId,omitempty"`
	// the http status code of the item
	HttpStatusCode *float32 `json:"httpStatusCode,omitempty"`
}

// NewFeatureEligibilityItem instantiates a new FeatureEligibilityItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureEligibilityItem() *FeatureEligibilityItem {
	this := FeatureEligibilityItem{}
	return &this
}

// NewFeatureEligibilityItemWithDefaults instantiates a new FeatureEligibilityItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureEligibilityItemWithDefaults() *FeatureEligibilityItem {
	this := FeatureEligibilityItem{}
	return &this
}

// GetMarketplaceId returns the MarketplaceId field value if set, zero value otherwise.
func (o *FeatureEligibilityItem) GetMarketplaceId() string {
	if o == nil || IsNil(o.MarketplaceId) {
		var ret string
		return ret
	}
	return *o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureEligibilityItem) GetMarketplaceIdOk() (*string, bool) {
	if o == nil || IsNil(o.MarketplaceId) {
		return nil, false
	}
	return o.MarketplaceId, true
}

// HasMarketplaceId returns a boolean if a field has been set.
func (o *FeatureEligibilityItem) HasMarketplaceId() bool {
	if o != nil && !IsNil(o.MarketplaceId) {
		return true
	}

	return false
}

// SetMarketplaceId gets a reference to the given string and assigns it to the MarketplaceId field.
func (o *FeatureEligibilityItem) SetMarketplaceId(v string) {
	o.MarketplaceId = &v
}

// GetReasons returns the Reasons field value if set, zero value otherwise.
func (o *FeatureEligibilityItem) GetReasons() []FeatureReasonItem {
	if o == nil || IsNil(o.Reasons) {
		var ret []FeatureReasonItem
		return ret
	}
	return o.Reasons
}

// GetReasonsOk returns a tuple with the Reasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureEligibilityItem) GetReasonsOk() ([]FeatureReasonItem, bool) {
	if o == nil || IsNil(o.Reasons) {
		return nil, false
	}
	return o.Reasons, true
}

// HasReasons returns a boolean if a field has been set.
func (o *FeatureEligibilityItem) HasReasons() bool {
	if o != nil && !IsNil(o.Reasons) {
		return true
	}

	return false
}

// SetReasons gets a reference to the given []FeatureReasonItem and assigns it to the Reasons field.
func (o *FeatureEligibilityItem) SetReasons(v []FeatureReasonItem) {
	o.Reasons = v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *FeatureEligibilityItem) GetResource() Resource {
	if o == nil || IsNil(o.Resource) {
		var ret Resource
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureEligibilityItem) GetResourceOk() (*Resource, bool) {
	if o == nil || IsNil(o.Resource) {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *FeatureEligibilityItem) HasResource() bool {
	if o != nil && !IsNil(o.Resource) {
		return true
	}

	return false
}

// SetResource gets a reference to the given Resource and assigns it to the Resource field.
func (o *FeatureEligibilityItem) SetResource(v Resource) {
	o.Resource = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *FeatureEligibilityItem) GetAction() Action {
	if o == nil || IsNil(o.Action) {
		var ret Action
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureEligibilityItem) GetActionOk() (*Action, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *FeatureEligibilityItem) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given Action and assigns it to the Action field.
func (o *FeatureEligibilityItem) SetAction(v Action) {
	o.Action = &v
}

// GetIsEligible returns the IsEligible field value if set, zero value otherwise.
func (o *FeatureEligibilityItem) GetIsEligible() bool {
	if o == nil || IsNil(o.IsEligible) {
		var ret bool
		return ret
	}
	return *o.IsEligible
}

// GetIsEligibleOk returns a tuple with the IsEligible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureEligibilityItem) GetIsEligibleOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEligible) {
		return nil, false
	}
	return o.IsEligible, true
}

// HasIsEligible returns a boolean if a field has been set.
func (o *FeatureEligibilityItem) HasIsEligible() bool {
	if o != nil && !IsNil(o.IsEligible) {
		return true
	}

	return false
}

// SetIsEligible gets a reference to the given bool and assigns it to the IsEligible field.
func (o *FeatureEligibilityItem) SetIsEligible(v bool) {
	o.IsEligible = &v
}

// GetItemRequestId returns the ItemRequestId field value if set, zero value otherwise.
func (o *FeatureEligibilityItem) GetItemRequestId() string {
	if o == nil || IsNil(o.ItemRequestId) {
		var ret string
		return ret
	}
	return *o.ItemRequestId
}

// GetItemRequestIdOk returns a tuple with the ItemRequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureEligibilityItem) GetItemRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.ItemRequestId) {
		return nil, false
	}
	return o.ItemRequestId, true
}

// HasItemRequestId returns a boolean if a field has been set.
func (o *FeatureEligibilityItem) HasItemRequestId() bool {
	if o != nil && !IsNil(o.ItemRequestId) {
		return true
	}

	return false
}

// SetItemRequestId gets a reference to the given string and assigns it to the ItemRequestId field.
func (o *FeatureEligibilityItem) SetItemRequestId(v string) {
	o.ItemRequestId = &v
}

// GetHttpStatusCode returns the HttpStatusCode field value if set, zero value otherwise.
func (o *FeatureEligibilityItem) GetHttpStatusCode() float32 {
	if o == nil || IsNil(o.HttpStatusCode) {
		var ret float32
		return ret
	}
	return *o.HttpStatusCode
}

// GetHttpStatusCodeOk returns a tuple with the HttpStatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureEligibilityItem) GetHttpStatusCodeOk() (*float32, bool) {
	if o == nil || IsNil(o.HttpStatusCode) {
		return nil, false
	}
	return o.HttpStatusCode, true
}

// HasHttpStatusCode returns a boolean if a field has been set.
func (o *FeatureEligibilityItem) HasHttpStatusCode() bool {
	if o != nil && !IsNil(o.HttpStatusCode) {
		return true
	}

	return false
}

// SetHttpStatusCode gets a reference to the given float32 and assigns it to the HttpStatusCode field.
func (o *FeatureEligibilityItem) SetHttpStatusCode(v float32) {
	o.HttpStatusCode = &v
}

func (o FeatureEligibilityItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MarketplaceId) {
		toSerialize["marketplaceId"] = o.MarketplaceId
	}
	if !IsNil(o.Reasons) {
		toSerialize["reasons"] = o.Reasons
	}
	if !IsNil(o.Resource) {
		toSerialize["resource"] = o.Resource
	}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.IsEligible) {
		toSerialize["isEligible"] = o.IsEligible
	}
	if !IsNil(o.ItemRequestId) {
		toSerialize["itemRequestId"] = o.ItemRequestId
	}
	if !IsNil(o.HttpStatusCode) {
		toSerialize["httpStatusCode"] = o.HttpStatusCode
	}
	return toSerialize, nil
}

type NullableFeatureEligibilityItem struct {
	value *FeatureEligibilityItem
	isSet bool
}

func (v NullableFeatureEligibilityItem) Get() *FeatureEligibilityItem {
	return v.value
}

func (v *NullableFeatureEligibilityItem) Set(val *FeatureEligibilityItem) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureEligibilityItem) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureEligibilityItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureEligibilityItem(val *FeatureEligibilityItem) *NullableFeatureEligibilityItem {
	return &NullableFeatureEligibilityItem{value: val, isSet: true}
}

func (v NullableFeatureEligibilityItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableFeatureEligibilityItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
