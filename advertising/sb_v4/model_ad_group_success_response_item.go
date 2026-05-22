package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the AdGroupSuccessResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdGroupSuccessResponseItem{}

// AdGroupSuccessResponseItem struct for AdGroupSuccessResponseItem
type AdGroupSuccessResponseItem struct {
	AdGroup *AdGroup `json:"adGroup,omitempty"`
	// the index of the adGroup in the array from the request body.
	Index float32 `json:"index"`
	// the adGroup ID.
	AdGroupId *string `json:"adGroupId,omitempty"`
}

type _AdGroupSuccessResponseItem AdGroupSuccessResponseItem

// NewAdGroupSuccessResponseItem instantiates a new AdGroupSuccessResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdGroupSuccessResponseItem(index float32) *AdGroupSuccessResponseItem {
	this := AdGroupSuccessResponseItem{}
	this.Index = index
	return &this
}

// NewAdGroupSuccessResponseItemWithDefaults instantiates a new AdGroupSuccessResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdGroupSuccessResponseItemWithDefaults() *AdGroupSuccessResponseItem {
	this := AdGroupSuccessResponseItem{}
	return &this
}

// GetAdGroup returns the AdGroup field value if set, zero value otherwise.
func (o *AdGroupSuccessResponseItem) GetAdGroup() AdGroup {
	if o == nil || IsNil(o.AdGroup) {
		var ret AdGroup
		return ret
	}
	return *o.AdGroup
}

// GetAdGroupOk returns a tuple with the AdGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdGroupSuccessResponseItem) GetAdGroupOk() (*AdGroup, bool) {
	if o == nil || IsNil(o.AdGroup) {
		return nil, false
	}
	return o.AdGroup, true
}

// HasAdGroup returns a boolean if a field has been set.
func (o *AdGroupSuccessResponseItem) HasAdGroup() bool {
	if o != nil && !IsNil(o.AdGroup) {
		return true
	}

	return false
}

// SetAdGroup gets a reference to the given AdGroup and assigns it to the AdGroup field.
func (o *AdGroupSuccessResponseItem) SetAdGroup(v AdGroup) {
	o.AdGroup = &v
}

// GetIndex returns the Index field value
func (o *AdGroupSuccessResponseItem) GetIndex() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *AdGroupSuccessResponseItem) GetIndexOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *AdGroupSuccessResponseItem) SetIndex(v float32) {
	o.Index = v
}

// GetAdGroupId returns the AdGroupId field value if set, zero value otherwise.
func (o *AdGroupSuccessResponseItem) GetAdGroupId() string {
	if o == nil || IsNil(o.AdGroupId) {
		var ret string
		return ret
	}
	return *o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdGroupSuccessResponseItem) GetAdGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.AdGroupId) {
		return nil, false
	}
	return o.AdGroupId, true
}

// HasAdGroupId returns a boolean if a field has been set.
func (o *AdGroupSuccessResponseItem) HasAdGroupId() bool {
	if o != nil && !IsNil(o.AdGroupId) {
		return true
	}

	return false
}

// SetAdGroupId gets a reference to the given string and assigns it to the AdGroupId field.
func (o *AdGroupSuccessResponseItem) SetAdGroupId(v string) {
	o.AdGroupId = &v
}

func (o AdGroupSuccessResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdGroup) {
		toSerialize["adGroup"] = o.AdGroup
	}
	toSerialize["index"] = o.Index
	if !IsNil(o.AdGroupId) {
		toSerialize["adGroupId"] = o.AdGroupId
	}
	return toSerialize, nil
}

type NullableAdGroupSuccessResponseItem struct {
	value *AdGroupSuccessResponseItem
	isSet bool
}

func (v NullableAdGroupSuccessResponseItem) Get() *AdGroupSuccessResponseItem {
	return v.value
}

func (v *NullableAdGroupSuccessResponseItem) Set(val *AdGroupSuccessResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableAdGroupSuccessResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableAdGroupSuccessResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdGroupSuccessResponseItem(val *AdGroupSuccessResponseItem) *NullableAdGroupSuccessResponseItem {
	return &NullableAdGroupSuccessResponseItem{value: val, isSet: true}
}

func (v NullableAdGroupSuccessResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAdGroupSuccessResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
