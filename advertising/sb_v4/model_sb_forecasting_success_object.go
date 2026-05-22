package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBForecastingSuccessObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBForecastingSuccessObject{}

// SBForecastingSuccessObject struct for SBForecastingSuccessObject
type SBForecastingSuccessObject struct {
	// Correlates the campaign to the campaign list index specified in the request. Zero-based.
	Index    *int32                        `json:"index,omitempty"`
	Campaign *SBForecastingSuccessCampaign `json:"campaign,omitempty"`
}

// NewSBForecastingSuccessObject instantiates a new SBForecastingSuccessObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBForecastingSuccessObject() *SBForecastingSuccessObject {
	this := SBForecastingSuccessObject{}
	return &this
}

// NewSBForecastingSuccessObjectWithDefaults instantiates a new SBForecastingSuccessObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBForecastingSuccessObjectWithDefaults() *SBForecastingSuccessObject {
	this := SBForecastingSuccessObject{}
	return &this
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *SBForecastingSuccessObject) GetIndex() int32 {
	if o == nil || IsNil(o.Index) {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBForecastingSuccessObject) GetIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *SBForecastingSuccessObject) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *SBForecastingSuccessObject) SetIndex(v int32) {
	o.Index = &v
}

// GetCampaign returns the Campaign field value if set, zero value otherwise.
func (o *SBForecastingSuccessObject) GetCampaign() SBForecastingSuccessCampaign {
	if o == nil || IsNil(o.Campaign) {
		var ret SBForecastingSuccessCampaign
		return ret
	}
	return *o.Campaign
}

// GetCampaignOk returns a tuple with the Campaign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBForecastingSuccessObject) GetCampaignOk() (*SBForecastingSuccessCampaign, bool) {
	if o == nil || IsNil(o.Campaign) {
		return nil, false
	}
	return o.Campaign, true
}

// HasCampaign returns a boolean if a field has been set.
func (o *SBForecastingSuccessObject) HasCampaign() bool {
	if o != nil && !IsNil(o.Campaign) {
		return true
	}

	return false
}

// SetCampaign gets a reference to the given SBForecastingSuccessCampaign and assigns it to the Campaign field.
func (o *SBForecastingSuccessObject) SetCampaign(v SBForecastingSuccessCampaign) {
	o.Campaign = &v
}

func (o SBForecastingSuccessObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Index) {
		toSerialize["index"] = o.Index
	}
	if !IsNil(o.Campaign) {
		toSerialize["campaign"] = o.Campaign
	}
	return toSerialize, nil
}

type NullableSBForecastingSuccessObject struct {
	value *SBForecastingSuccessObject
	isSet bool
}

func (v NullableSBForecastingSuccessObject) Get() *SBForecastingSuccessObject {
	return v.value
}

func (v *NullableSBForecastingSuccessObject) Set(val *SBForecastingSuccessObject) {
	v.value = val
	v.isSet = true
}

func (v NullableSBForecastingSuccessObject) IsSet() bool {
	return v.isSet
}

func (v *NullableSBForecastingSuccessObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBForecastingSuccessObject(val *SBForecastingSuccessObject) *NullableSBForecastingSuccessObject {
	return &NullableSBForecastingSuccessObject{value: val, isSet: true}
}

func (v NullableSBForecastingSuccessObject) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBForecastingSuccessObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
