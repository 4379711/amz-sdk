package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the AsinComponentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AsinComponentResponse{}

// AsinComponentResponse Pre-moderation result for an Asin component
type AsinComponentResponse struct {
	// The pre-moderation status of the component.
	PreModerationStatus *string `json:"preModerationStatus,omitempty"`
	// Type of Asin component.
	ComponentType *string `json:"componentType,omitempty"`
	// A list of policy violations for the component that were detected during pre moderation. Note that this field is present in the response only when preModerationStatus is set to REJECTED.
	PolicyViolations []AsinPolicyViolation `json:"policyViolations,omitempty"`
	// Pre-moderated Asin Id.
	Asin *string `json:"asin,omitempty"`
	// Id of the component. This is the same id sent as part of the request. This can be used to uniquely identify the component.
	Id *string `json:"id,omitempty"`
}

// NewAsinComponentResponse instantiates a new AsinComponentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAsinComponentResponse() *AsinComponentResponse {
	this := AsinComponentResponse{}
	return &this
}

// NewAsinComponentResponseWithDefaults instantiates a new AsinComponentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAsinComponentResponseWithDefaults() *AsinComponentResponse {
	this := AsinComponentResponse{}
	return &this
}

// GetPreModerationStatus returns the PreModerationStatus field value if set, zero value otherwise.
func (o *AsinComponentResponse) GetPreModerationStatus() string {
	if o == nil || IsNil(o.PreModerationStatus) {
		var ret string
		return ret
	}
	return *o.PreModerationStatus
}

// GetPreModerationStatusOk returns a tuple with the PreModerationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsinComponentResponse) GetPreModerationStatusOk() (*string, bool) {
	if o == nil || IsNil(o.PreModerationStatus) {
		return nil, false
	}
	return o.PreModerationStatus, true
}

// HasPreModerationStatus returns a boolean if a field has been set.
func (o *AsinComponentResponse) HasPreModerationStatus() bool {
	if o != nil && !IsNil(o.PreModerationStatus) {
		return true
	}

	return false
}

// SetPreModerationStatus gets a reference to the given string and assigns it to the PreModerationStatus field.
func (o *AsinComponentResponse) SetPreModerationStatus(v string) {
	o.PreModerationStatus = &v
}

// GetComponentType returns the ComponentType field value if set, zero value otherwise.
func (o *AsinComponentResponse) GetComponentType() string {
	if o == nil || IsNil(o.ComponentType) {
		var ret string
		return ret
	}
	return *o.ComponentType
}

// GetComponentTypeOk returns a tuple with the ComponentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsinComponentResponse) GetComponentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ComponentType) {
		return nil, false
	}
	return o.ComponentType, true
}

// HasComponentType returns a boolean if a field has been set.
func (o *AsinComponentResponse) HasComponentType() bool {
	if o != nil && !IsNil(o.ComponentType) {
		return true
	}

	return false
}

// SetComponentType gets a reference to the given string and assigns it to the ComponentType field.
func (o *AsinComponentResponse) SetComponentType(v string) {
	o.ComponentType = &v
}

// GetPolicyViolations returns the PolicyViolations field value if set, zero value otherwise.
func (o *AsinComponentResponse) GetPolicyViolations() []AsinPolicyViolation {
	if o == nil || IsNil(o.PolicyViolations) {
		var ret []AsinPolicyViolation
		return ret
	}
	return o.PolicyViolations
}

// GetPolicyViolationsOk returns a tuple with the PolicyViolations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsinComponentResponse) GetPolicyViolationsOk() ([]AsinPolicyViolation, bool) {
	if o == nil || IsNil(o.PolicyViolations) {
		return nil, false
	}
	return o.PolicyViolations, true
}

// HasPolicyViolations returns a boolean if a field has been set.
func (o *AsinComponentResponse) HasPolicyViolations() bool {
	if o != nil && !IsNil(o.PolicyViolations) {
		return true
	}

	return false
}

// SetPolicyViolations gets a reference to the given []AsinPolicyViolation and assigns it to the PolicyViolations field.
func (o *AsinComponentResponse) SetPolicyViolations(v []AsinPolicyViolation) {
	o.PolicyViolations = v
}

// GetAsin returns the Asin field value if set, zero value otherwise.
func (o *AsinComponentResponse) GetAsin() string {
	if o == nil || IsNil(o.Asin) {
		var ret string
		return ret
	}
	return *o.Asin
}

// GetAsinOk returns a tuple with the Asin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsinComponentResponse) GetAsinOk() (*string, bool) {
	if o == nil || IsNil(o.Asin) {
		return nil, false
	}
	return o.Asin, true
}

// HasAsin returns a boolean if a field has been set.
func (o *AsinComponentResponse) HasAsin() bool {
	if o != nil && !IsNil(o.Asin) {
		return true
	}

	return false
}

// SetAsin gets a reference to the given string and assigns it to the Asin field.
func (o *AsinComponentResponse) SetAsin(v string) {
	o.Asin = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AsinComponentResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsinComponentResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AsinComponentResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AsinComponentResponse) SetId(v string) {
	o.Id = &v
}

func (o AsinComponentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PreModerationStatus) {
		toSerialize["preModerationStatus"] = o.PreModerationStatus
	}
	if !IsNil(o.ComponentType) {
		toSerialize["componentType"] = o.ComponentType
	}
	if !IsNil(o.PolicyViolations) {
		toSerialize["policyViolations"] = o.PolicyViolations
	}
	if !IsNil(o.Asin) {
		toSerialize["asin"] = o.Asin
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableAsinComponentResponse struct {
	value *AsinComponentResponse
	isSet bool
}

func (v NullableAsinComponentResponse) Get() *AsinComponentResponse {
	return v.value
}

func (v *NullableAsinComponentResponse) Set(val *AsinComponentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAsinComponentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAsinComponentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsinComponentResponse(val *AsinComponentResponse) *NullableAsinComponentResponse {
	return &NullableAsinComponentResponse{value: val, isSet: true}
}

func (v NullableAsinComponentResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAsinComponentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
