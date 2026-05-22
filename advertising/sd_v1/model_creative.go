package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the Creative type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Creative{}

// Creative Creative model.
type Creative struct {
	// Unique identifier of the creative.
	CreativeId float32 `json:"creativeId"`
	// The identifier of the ad group.
	AdGroupId    int64                          `json:"adGroupId"`
	CreativeType CreativeTypeInCreativeResponse `json:"creativeType"`
	Properties   CreativeProperties             `json:"properties"`
	// The moderation status of the creative
	ModerationStatus string `json:"moderationStatus"`
}

type _Creative Creative

// NewCreative instantiates a new Creative object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreative(creativeId float32, adGroupId int64, creativeType CreativeTypeInCreativeResponse, properties CreativeProperties, moderationStatus string) *Creative {
	this := Creative{}
	this.CreativeId = creativeId
	this.AdGroupId = adGroupId
	this.CreativeType = creativeType
	this.Properties = properties
	this.ModerationStatus = moderationStatus
	return &this
}

// NewCreativeWithDefaults instantiates a new Creative object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreativeWithDefaults() *Creative {
	this := Creative{}
	return &this
}

// GetCreativeId returns the CreativeId field value
func (o *Creative) GetCreativeId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CreativeId
}

// GetCreativeIdOk returns a tuple with the CreativeId field value
// and a boolean to check if the value has been set.
func (o *Creative) GetCreativeIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreativeId, true
}

// SetCreativeId sets field value
func (o *Creative) SetCreativeId(v float32) {
	o.CreativeId = v
}

// GetAdGroupId returns the AdGroupId field value
func (o *Creative) GetAdGroupId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *Creative) GetAdGroupIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *Creative) SetAdGroupId(v int64) {
	o.AdGroupId = v
}

// GetCreativeType returns the CreativeType field value
func (o *Creative) GetCreativeType() CreativeTypeInCreativeResponse {
	if o == nil {
		var ret CreativeTypeInCreativeResponse
		return ret
	}

	return o.CreativeType
}

// GetCreativeTypeOk returns a tuple with the CreativeType field value
// and a boolean to check if the value has been set.
func (o *Creative) GetCreativeTypeOk() (*CreativeTypeInCreativeResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreativeType, true
}

// SetCreativeType sets field value
func (o *Creative) SetCreativeType(v CreativeTypeInCreativeResponse) {
	o.CreativeType = v
}

// GetProperties returns the Properties field value
func (o *Creative) GetProperties() CreativeProperties {
	if o == nil {
		var ret CreativeProperties
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *Creative) GetPropertiesOk() (*CreativeProperties, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *Creative) SetProperties(v CreativeProperties) {
	o.Properties = v
}

// GetModerationStatus returns the ModerationStatus field value
func (o *Creative) GetModerationStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModerationStatus
}

// GetModerationStatusOk returns a tuple with the ModerationStatus field value
// and a boolean to check if the value has been set.
func (o *Creative) GetModerationStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModerationStatus, true
}

// SetModerationStatus sets field value
func (o *Creative) SetModerationStatus(v string) {
	o.ModerationStatus = v
}

func (o Creative) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["creativeId"] = o.CreativeId
	toSerialize["adGroupId"] = o.AdGroupId
	toSerialize["creativeType"] = o.CreativeType
	toSerialize["properties"] = o.Properties
	toSerialize["moderationStatus"] = o.ModerationStatus
	return toSerialize, nil
}

type NullableCreative struct {
	value *Creative
	isSet bool
}

func (v NullableCreative) Get() *Creative {
	return v.value
}

func (v *NullableCreative) Set(val *Creative) {
	v.value = val
	v.isSet = true
}

func (v NullableCreative) IsSet() bool {
	return v.isSet
}

func (v *NullableCreative) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreative(val *Creative) *NullableCreative {
	return &NullableCreative{value: val, isSet: true}
}

func (v NullableCreative) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreative) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
