package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the TargetingPredicateLegacy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TargetingPredicateLegacy{}

// TargetingPredicateLegacy struct for TargetingPredicateLegacy
type TargetingPredicateLegacy struct {
	Type *string `json:"type,omitempty"`
	// The value to be targeted.
	Value *string `json:"value,omitempty"`
	// The type of event that the value applies to. Only available for similarProduct and exactProduct currently. * views event type corresponds to a customer who viewed the detail page of the product(s).
	EventType *string `json:"eventType,omitempty"`
}

// NewTargetingPredicateLegacy instantiates a new TargetingPredicateLegacy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetingPredicateLegacy() *TargetingPredicateLegacy {
	this := TargetingPredicateLegacy{}
	return &this
}

// NewTargetingPredicateLegacyWithDefaults instantiates a new TargetingPredicateLegacy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetingPredicateLegacyWithDefaults() *TargetingPredicateLegacy {
	this := TargetingPredicateLegacy{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TargetingPredicateLegacy) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetingPredicateLegacy) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TargetingPredicateLegacy) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TargetingPredicateLegacy) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TargetingPredicateLegacy) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetingPredicateLegacy) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TargetingPredicateLegacy) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *TargetingPredicateLegacy) SetValue(v string) {
	o.Value = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *TargetingPredicateLegacy) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetingPredicateLegacy) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *TargetingPredicateLegacy) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *TargetingPredicateLegacy) SetEventType(v string) {
	o.EventType = &v
}

func (o TargetingPredicateLegacy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.EventType) {
		toSerialize["eventType"] = o.EventType
	}
	return toSerialize, nil
}

type NullableTargetingPredicateLegacy struct {
	value *TargetingPredicateLegacy
	isSet bool
}

func (v NullableTargetingPredicateLegacy) Get() *TargetingPredicateLegacy {
	return v.value
}

func (v *NullableTargetingPredicateLegacy) Set(val *TargetingPredicateLegacy) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetingPredicateLegacy) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetingPredicateLegacy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetingPredicateLegacy(val *TargetingPredicateLegacy) *NullableTargetingPredicateLegacy {
	return &NullableTargetingPredicateLegacy{value: val, isSet: true}
}

func (v NullableTargetingPredicateLegacy) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTargetingPredicateLegacy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
