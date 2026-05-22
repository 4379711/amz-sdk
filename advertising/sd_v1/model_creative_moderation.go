package sd_v1

import (
	"time"

	"github.com/bytedance/sonic"
)

// checks if the CreativeModeration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreativeModeration{}

// CreativeModeration System generated Creative moderation.
type CreativeModeration struct {
	// Unique identifier of the creative.
	CreativeId   float32                        `json:"creativeId"`
	CreativeType CreativeTypeInCreativeResponse `json:"creativeType"`
	// The moderation status of the creative. |Status|Description| |------|-----------| |APPROVED|Moderation for the creative is complete.| |IN_PROGRESS|Moderation for the creative is in progress. The expected date and time for completion are specfied in the `etaForModeration` field.| |REJECTED|The creative has failed moderation. Specific information about the content that violated policy is available in `policyViolations`.|
	ModerationStatus string `json:"moderationStatus"`
	// Expected date and time by which moderation will be complete.
	EtaForModeration time.Time `json:"etaForModeration"`
	// A list of policy violations for a creative that has failed moderation.
	PolicyViolations []CreativeModerationPolicyViolationsInner `json:"policyViolations"`
}

type _CreativeModeration CreativeModeration

// NewCreativeModeration instantiates a new CreativeModeration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreativeModeration(creativeId float32, creativeType CreativeTypeInCreativeResponse, moderationStatus string, etaForModeration time.Time, policyViolations []CreativeModerationPolicyViolationsInner) *CreativeModeration {
	this := CreativeModeration{}
	this.CreativeId = creativeId
	this.CreativeType = creativeType
	this.ModerationStatus = moderationStatus
	this.EtaForModeration = etaForModeration
	this.PolicyViolations = policyViolations
	return &this
}

// NewCreativeModerationWithDefaults instantiates a new CreativeModeration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreativeModerationWithDefaults() *CreativeModeration {
	this := CreativeModeration{}
	return &this
}

// GetCreativeId returns the CreativeId field value
func (o *CreativeModeration) GetCreativeId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CreativeId
}

// GetCreativeIdOk returns a tuple with the CreativeId field value
// and a boolean to check if the value has been set.
func (o *CreativeModeration) GetCreativeIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreativeId, true
}

// SetCreativeId sets field value
func (o *CreativeModeration) SetCreativeId(v float32) {
	o.CreativeId = v
}

// GetCreativeType returns the CreativeType field value
func (o *CreativeModeration) GetCreativeType() CreativeTypeInCreativeResponse {
	if o == nil {
		var ret CreativeTypeInCreativeResponse
		return ret
	}

	return o.CreativeType
}

// GetCreativeTypeOk returns a tuple with the CreativeType field value
// and a boolean to check if the value has been set.
func (o *CreativeModeration) GetCreativeTypeOk() (*CreativeTypeInCreativeResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreativeType, true
}

// SetCreativeType sets field value
func (o *CreativeModeration) SetCreativeType(v CreativeTypeInCreativeResponse) {
	o.CreativeType = v
}

// GetModerationStatus returns the ModerationStatus field value
func (o *CreativeModeration) GetModerationStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModerationStatus
}

// GetModerationStatusOk returns a tuple with the ModerationStatus field value
// and a boolean to check if the value has been set.
func (o *CreativeModeration) GetModerationStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModerationStatus, true
}

// SetModerationStatus sets field value
func (o *CreativeModeration) SetModerationStatus(v string) {
	o.ModerationStatus = v
}

// GetEtaForModeration returns the EtaForModeration field value
func (o *CreativeModeration) GetEtaForModeration() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EtaForModeration
}

// GetEtaForModerationOk returns a tuple with the EtaForModeration field value
// and a boolean to check if the value has been set.
func (o *CreativeModeration) GetEtaForModerationOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EtaForModeration, true
}

// SetEtaForModeration sets field value
func (o *CreativeModeration) SetEtaForModeration(v time.Time) {
	o.EtaForModeration = v
}

// GetPolicyViolations returns the PolicyViolations field value
func (o *CreativeModeration) GetPolicyViolations() []CreativeModerationPolicyViolationsInner {
	if o == nil {
		var ret []CreativeModerationPolicyViolationsInner
		return ret
	}

	return o.PolicyViolations
}

// GetPolicyViolationsOk returns a tuple with the PolicyViolations field value
// and a boolean to check if the value has been set.
func (o *CreativeModeration) GetPolicyViolationsOk() ([]CreativeModerationPolicyViolationsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.PolicyViolations, true
}

// SetPolicyViolations sets field value
func (o *CreativeModeration) SetPolicyViolations(v []CreativeModerationPolicyViolationsInner) {
	o.PolicyViolations = v
}

func (o CreativeModeration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["creativeId"] = o.CreativeId
	toSerialize["creativeType"] = o.CreativeType
	toSerialize["moderationStatus"] = o.ModerationStatus
	toSerialize["etaForModeration"] = o.EtaForModeration
	toSerialize["policyViolations"] = o.PolicyViolations
	return toSerialize, nil
}

type NullableCreativeModeration struct {
	value *CreativeModeration
	isSet bool
}

func (v NullableCreativeModeration) Get() *CreativeModeration {
	return v.value
}

func (v *NullableCreativeModeration) Set(val *CreativeModeration) {
	v.value = val
	v.isSet = true
}

func (v NullableCreativeModeration) IsSet() bool {
	return v.isSet
}

func (v *NullableCreativeModeration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreativeModeration(val *CreativeModeration) *NullableCreativeModeration {
	return &NullableCreativeModeration{value: val, isSet: true}
}

func (v NullableCreativeModeration) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreativeModeration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
