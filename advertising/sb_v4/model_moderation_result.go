package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ModerationResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModerationResult{}

// ModerationResult struct for ModerationResult
type ModerationResult struct {
	// The version identifier that helps to keep track of multiple versions of a submitted ad. In case of Sponsored Brands this is the creative version id.
	VersionId        *string           `json:"versionId,omitempty"`
	IdType           *IdType           `json:"idType,omitempty"`
	ModerationStatus *ModerationStatus `json:"moderationStatus,omitempty"`
	// A list of policy violations for a campaign that has failed moderation. Note that this field is present in the response only when moderationStatus is set to REJECTED.
	PolicyViolations []PolicyViolation `json:"policyViolations,omitempty"`
	// Expected date and time by which moderation will be complete. The format is ISO 8601 in UTC time zone. Note that this field is present in the response only when moderationStatus is set to IN_PROGRESS.
	EtaForModeration *string `json:"etaForModeration,omitempty"`
	// The unique identifier of the ad which can be obtained after the ad is created using create APIs.
	Id *string `json:"id,omitempty"`
}

// NewModerationResult instantiates a new ModerationResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModerationResult() *ModerationResult {
	this := ModerationResult{}
	return &this
}

// NewModerationResultWithDefaults instantiates a new ModerationResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModerationResultWithDefaults() *ModerationResult {
	this := ModerationResult{}
	return &this
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *ModerationResult) GetVersionId() string {
	if o == nil || IsNil(o.VersionId) {
		var ret string
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModerationResult) GetVersionIdOk() (*string, bool) {
	if o == nil || IsNil(o.VersionId) {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *ModerationResult) HasVersionId() bool {
	if o != nil && !IsNil(o.VersionId) {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given string and assigns it to the VersionId field.
func (o *ModerationResult) SetVersionId(v string) {
	o.VersionId = &v
}

// GetIdType returns the IdType field value if set, zero value otherwise.
func (o *ModerationResult) GetIdType() IdType {
	if o == nil || IsNil(o.IdType) {
		var ret IdType
		return ret
	}
	return *o.IdType
}

// GetIdTypeOk returns a tuple with the IdType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModerationResult) GetIdTypeOk() (*IdType, bool) {
	if o == nil || IsNil(o.IdType) {
		return nil, false
	}
	return o.IdType, true
}

// HasIdType returns a boolean if a field has been set.
func (o *ModerationResult) HasIdType() bool {
	if o != nil && !IsNil(o.IdType) {
		return true
	}

	return false
}

// SetIdType gets a reference to the given IdType and assigns it to the IdType field.
func (o *ModerationResult) SetIdType(v IdType) {
	o.IdType = &v
}

// GetModerationStatus returns the ModerationStatus field value if set, zero value otherwise.
func (o *ModerationResult) GetModerationStatus() ModerationStatus {
	if o == nil || IsNil(o.ModerationStatus) {
		var ret ModerationStatus
		return ret
	}
	return *o.ModerationStatus
}

// GetModerationStatusOk returns a tuple with the ModerationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModerationResult) GetModerationStatusOk() (*ModerationStatus, bool) {
	if o == nil || IsNil(o.ModerationStatus) {
		return nil, false
	}
	return o.ModerationStatus, true
}

// HasModerationStatus returns a boolean if a field has been set.
func (o *ModerationResult) HasModerationStatus() bool {
	if o != nil && !IsNil(o.ModerationStatus) {
		return true
	}

	return false
}

// SetModerationStatus gets a reference to the given ModerationStatus and assigns it to the ModerationStatus field.
func (o *ModerationResult) SetModerationStatus(v ModerationStatus) {
	o.ModerationStatus = &v
}

// GetPolicyViolations returns the PolicyViolations field value if set, zero value otherwise.
func (o *ModerationResult) GetPolicyViolations() []PolicyViolation {
	if o == nil || IsNil(o.PolicyViolations) {
		var ret []PolicyViolation
		return ret
	}
	return o.PolicyViolations
}

// GetPolicyViolationsOk returns a tuple with the PolicyViolations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModerationResult) GetPolicyViolationsOk() ([]PolicyViolation, bool) {
	if o == nil || IsNil(o.PolicyViolations) {
		return nil, false
	}
	return o.PolicyViolations, true
}

// HasPolicyViolations returns a boolean if a field has been set.
func (o *ModerationResult) HasPolicyViolations() bool {
	if o != nil && !IsNil(o.PolicyViolations) {
		return true
	}

	return false
}

// SetPolicyViolations gets a reference to the given []PolicyViolation and assigns it to the PolicyViolations field.
func (o *ModerationResult) SetPolicyViolations(v []PolicyViolation) {
	o.PolicyViolations = v
}

// GetEtaForModeration returns the EtaForModeration field value if set, zero value otherwise.
func (o *ModerationResult) GetEtaForModeration() string {
	if o == nil || IsNil(o.EtaForModeration) {
		var ret string
		return ret
	}
	return *o.EtaForModeration
}

// GetEtaForModerationOk returns a tuple with the EtaForModeration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModerationResult) GetEtaForModerationOk() (*string, bool) {
	if o == nil || IsNil(o.EtaForModeration) {
		return nil, false
	}
	return o.EtaForModeration, true
}

// HasEtaForModeration returns a boolean if a field has been set.
func (o *ModerationResult) HasEtaForModeration() bool {
	if o != nil && !IsNil(o.EtaForModeration) {
		return true
	}

	return false
}

// SetEtaForModeration gets a reference to the given string and assigns it to the EtaForModeration field.
func (o *ModerationResult) SetEtaForModeration(v string) {
	o.EtaForModeration = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModerationResult) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModerationResult) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModerationResult) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ModerationResult) SetId(v string) {
	o.Id = &v
}

func (o ModerationResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VersionId) {
		toSerialize["versionId"] = o.VersionId
	}
	if !IsNil(o.IdType) {
		toSerialize["idType"] = o.IdType
	}
	if !IsNil(o.ModerationStatus) {
		toSerialize["moderationStatus"] = o.ModerationStatus
	}
	if !IsNil(o.PolicyViolations) {
		toSerialize["policyViolations"] = o.PolicyViolations
	}
	if !IsNil(o.EtaForModeration) {
		toSerialize["etaForModeration"] = o.EtaForModeration
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableModerationResult struct {
	value *ModerationResult
	isSet bool
}

func (v NullableModerationResult) Get() *ModerationResult {
	return v.value
}

func (v *NullableModerationResult) Set(val *ModerationResult) {
	v.value = val
	v.isSet = true
}

func (v NullableModerationResult) IsSet() bool {
	return v.isSet
}

func (v *NullableModerationResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModerationResult(val *ModerationResult) *NullableModerationResult {
	return &NullableModerationResult{value: val, isSet: true}
}

func (v NullableModerationResult) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableModerationResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
