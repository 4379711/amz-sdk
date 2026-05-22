package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the AdSuccessResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdSuccessResponseItem{}

// AdSuccessResponseItem struct for AdSuccessResponseItem
type AdSuccessResponseItem struct {
	// the Ad ID.
	AdId *string         `json:"adId,omitempty"`
	Ad   *MultiAdGroupAd `json:"ad,omitempty"`
	// The index in the original list from the request.
	Index float32 `json:"index"`
}

type _AdSuccessResponseItem AdSuccessResponseItem

// NewAdSuccessResponseItem instantiates a new AdSuccessResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdSuccessResponseItem(index float32) *AdSuccessResponseItem {
	this := AdSuccessResponseItem{}
	this.Index = index
	return &this
}

// NewAdSuccessResponseItemWithDefaults instantiates a new AdSuccessResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdSuccessResponseItemWithDefaults() *AdSuccessResponseItem {
	this := AdSuccessResponseItem{}
	return &this
}

// GetAdId returns the AdId field value if set, zero value otherwise.
func (o *AdSuccessResponseItem) GetAdId() string {
	if o == nil || IsNil(o.AdId) {
		var ret string
		return ret
	}
	return *o.AdId
}

// GetAdIdOk returns a tuple with the AdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdSuccessResponseItem) GetAdIdOk() (*string, bool) {
	if o == nil || IsNil(o.AdId) {
		return nil, false
	}
	return o.AdId, true
}

// HasAdId returns a boolean if a field has been set.
func (o *AdSuccessResponseItem) HasAdId() bool {
	if o != nil && !IsNil(o.AdId) {
		return true
	}

	return false
}

// SetAdId gets a reference to the given string and assigns it to the AdId field.
func (o *AdSuccessResponseItem) SetAdId(v string) {
	o.AdId = &v
}

// GetAd returns the Ad field value if set, zero value otherwise.
func (o *AdSuccessResponseItem) GetAd() MultiAdGroupAd {
	if o == nil || IsNil(o.Ad) {
		var ret MultiAdGroupAd
		return ret
	}
	return *o.Ad
}

// GetAdOk returns a tuple with the Ad field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdSuccessResponseItem) GetAdOk() (*MultiAdGroupAd, bool) {
	if o == nil || IsNil(o.Ad) {
		return nil, false
	}
	return o.Ad, true
}

// HasAd returns a boolean if a field has been set.
func (o *AdSuccessResponseItem) HasAd() bool {
	if o != nil && !IsNil(o.Ad) {
		return true
	}

	return false
}

// SetAd gets a reference to the given MultiAdGroupAd and assigns it to the Ad field.
func (o *AdSuccessResponseItem) SetAd(v MultiAdGroupAd) {
	o.Ad = &v
}

// GetIndex returns the Index field value
func (o *AdSuccessResponseItem) GetIndex() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *AdSuccessResponseItem) GetIndexOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *AdSuccessResponseItem) SetIndex(v float32) {
	o.Index = v
}

func (o AdSuccessResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdId) {
		toSerialize["adId"] = o.AdId
	}
	if !IsNil(o.Ad) {
		toSerialize["ad"] = o.Ad
	}
	toSerialize["index"] = o.Index
	return toSerialize, nil
}

type NullableAdSuccessResponseItem struct {
	value *AdSuccessResponseItem
	isSet bool
}

func (v NullableAdSuccessResponseItem) Get() *AdSuccessResponseItem {
	return v.value
}

func (v *NullableAdSuccessResponseItem) Set(val *AdSuccessResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableAdSuccessResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableAdSuccessResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdSuccessResponseItem(val *AdSuccessResponseItem) *NullableAdSuccessResponseItem {
	return &NullableAdSuccessResponseItem{value: val, isSet: true}
}

func (v NullableAdSuccessResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAdSuccessResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
