package listings_restrictions_20210801

import (
	"github.com/bytedance/sonic"
)

// checks if the Restriction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Restriction{}

// Restriction A listing restriction, optionally qualified by a condition, with a list of reasons for the restriction.
type Restriction struct {
	// A marketplace identifier. Identifies the Amazon marketplace where the restriction is enforced.
	MarketplaceId string `json:"marketplaceId"`
	// The condition that applies to the restriction.
	ConditionType *string `json:"conditionType,omitempty"`
	// A list of reasons for the restriction.
	Reasons []Reason `json:"reasons,omitempty"`
}

type _Restriction Restriction

// NewRestriction instantiates a new Restriction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestriction(marketplaceId string) *Restriction {
	this := Restriction{}
	this.MarketplaceId = marketplaceId
	return &this
}

// NewRestrictionWithDefaults instantiates a new Restriction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestrictionWithDefaults() *Restriction {
	this := Restriction{}
	return &this
}

// GetMarketplaceId returns the MarketplaceId field value
func (o *Restriction) GetMarketplaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value
// and a boolean to check if the value has been set.
func (o *Restriction) GetMarketplaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarketplaceId, true
}

// SetMarketplaceId sets field value
func (o *Restriction) SetMarketplaceId(v string) {
	o.MarketplaceId = v
}

// GetConditionType returns the ConditionType field value if set, zero value otherwise.
func (o *Restriction) GetConditionType() string {
	if o == nil || IsNil(o.ConditionType) {
		var ret string
		return ret
	}
	return *o.ConditionType
}

// GetConditionTypeOk returns a tuple with the ConditionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Restriction) GetConditionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ConditionType) {
		return nil, false
	}
	return o.ConditionType, true
}

// HasConditionType returns a boolean if a field has been set.
func (o *Restriction) HasConditionType() bool {
	if o != nil && !IsNil(o.ConditionType) {
		return true
	}

	return false
}

// SetConditionType gets a reference to the given string and assigns it to the ConditionType field.
func (o *Restriction) SetConditionType(v string) {
	o.ConditionType = &v
}

// GetReasons returns the Reasons field value if set, zero value otherwise.
func (o *Restriction) GetReasons() []Reason {
	if o == nil || IsNil(o.Reasons) {
		var ret []Reason
		return ret
	}
	return o.Reasons
}

// GetReasonsOk returns a tuple with the Reasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Restriction) GetReasonsOk() ([]Reason, bool) {
	if o == nil || IsNil(o.Reasons) {
		return nil, false
	}
	return o.Reasons, true
}

// HasReasons returns a boolean if a field has been set.
func (o *Restriction) HasReasons() bool {
	if o != nil && !IsNil(o.Reasons) {
		return true
	}

	return false
}

// SetReasons gets a reference to the given []Reason and assigns it to the Reasons field.
func (o *Restriction) SetReasons(v []Reason) {
	o.Reasons = v
}

func (o Restriction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["marketplaceId"] = o.MarketplaceId
	if !IsNil(o.ConditionType) {
		toSerialize["conditionType"] = o.ConditionType
	}
	if !IsNil(o.Reasons) {
		toSerialize["reasons"] = o.Reasons
	}
	return toSerialize, nil
}

type NullableRestriction struct {
	value *Restriction
	isSet bool
}

func (v NullableRestriction) Get() *Restriction {
	return v.value
}

func (v *NullableRestriction) Set(val *Restriction) {
	v.value = val
	v.isSet = true
}

func (v NullableRestriction) IsSet() bool {
	return v.isSet
}

func (v *NullableRestriction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestriction(val *Restriction) *NullableRestriction {
	return &NullableRestriction{value: val, isSet: true}
}

func (v NullableRestriction) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRestriction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
