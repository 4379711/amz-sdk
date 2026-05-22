package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the DateComponentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DateComponentResponse{}

// DateComponentResponse Pre-moderation result for a date component
type DateComponentResponse struct {
	// The pre-moderation status of the component.
	PreModerationStatus *string `json:"preModerationStatus,omitempty"`
	// Type of the date component.
	ComponentType *string `json:"componentType,omitempty"`
	// End date of the component.
	EndDate *string `json:"endDate,omitempty"`
	// A list of policy violations for the component that were detected during pre moderation. Note that this field is present in the response only when preModerationStatus is set to REJECTED.
	PolicyViolations []DatePolicyViolation `json:"policyViolations,omitempty"`
	// Id of the component. This is the same id sent as part of the request. This can be used to uniquely identify the component.
	Id *string `json:"id,omitempty"`
	// Start date of the component.
	StartDate *string `json:"startDate,omitempty"`
}

// NewDateComponentResponse instantiates a new DateComponentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDateComponentResponse() *DateComponentResponse {
	this := DateComponentResponse{}
	return &this
}

// NewDateComponentResponseWithDefaults instantiates a new DateComponentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDateComponentResponseWithDefaults() *DateComponentResponse {
	this := DateComponentResponse{}
	return &this
}

// GetPreModerationStatus returns the PreModerationStatus field value if set, zero value otherwise.
func (o *DateComponentResponse) GetPreModerationStatus() string {
	if o == nil || IsNil(o.PreModerationStatus) {
		var ret string
		return ret
	}
	return *o.PreModerationStatus
}

// GetPreModerationStatusOk returns a tuple with the PreModerationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateComponentResponse) GetPreModerationStatusOk() (*string, bool) {
	if o == nil || IsNil(o.PreModerationStatus) {
		return nil, false
	}
	return o.PreModerationStatus, true
}

// HasPreModerationStatus returns a boolean if a field has been set.
func (o *DateComponentResponse) HasPreModerationStatus() bool {
	if o != nil && !IsNil(o.PreModerationStatus) {
		return true
	}

	return false
}

// SetPreModerationStatus gets a reference to the given string and assigns it to the PreModerationStatus field.
func (o *DateComponentResponse) SetPreModerationStatus(v string) {
	o.PreModerationStatus = &v
}

// GetComponentType returns the ComponentType field value if set, zero value otherwise.
func (o *DateComponentResponse) GetComponentType() string {
	if o == nil || IsNil(o.ComponentType) {
		var ret string
		return ret
	}
	return *o.ComponentType
}

// GetComponentTypeOk returns a tuple with the ComponentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateComponentResponse) GetComponentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ComponentType) {
		return nil, false
	}
	return o.ComponentType, true
}

// HasComponentType returns a boolean if a field has been set.
func (o *DateComponentResponse) HasComponentType() bool {
	if o != nil && !IsNil(o.ComponentType) {
		return true
	}

	return false
}

// SetComponentType gets a reference to the given string and assigns it to the ComponentType field.
func (o *DateComponentResponse) SetComponentType(v string) {
	o.ComponentType = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *DateComponentResponse) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateComponentResponse) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *DateComponentResponse) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *DateComponentResponse) SetEndDate(v string) {
	o.EndDate = &v
}

// GetPolicyViolations returns the PolicyViolations field value if set, zero value otherwise.
func (o *DateComponentResponse) GetPolicyViolations() []DatePolicyViolation {
	if o == nil || IsNil(o.PolicyViolations) {
		var ret []DatePolicyViolation
		return ret
	}
	return o.PolicyViolations
}

// GetPolicyViolationsOk returns a tuple with the PolicyViolations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateComponentResponse) GetPolicyViolationsOk() ([]DatePolicyViolation, bool) {
	if o == nil || IsNil(o.PolicyViolations) {
		return nil, false
	}
	return o.PolicyViolations, true
}

// HasPolicyViolations returns a boolean if a field has been set.
func (o *DateComponentResponse) HasPolicyViolations() bool {
	if o != nil && !IsNil(o.PolicyViolations) {
		return true
	}

	return false
}

// SetPolicyViolations gets a reference to the given []DatePolicyViolation and assigns it to the PolicyViolations field.
func (o *DateComponentResponse) SetPolicyViolations(v []DatePolicyViolation) {
	o.PolicyViolations = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DateComponentResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateComponentResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DateComponentResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DateComponentResponse) SetId(v string) {
	o.Id = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *DateComponentResponse) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DateComponentResponse) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *DateComponentResponse) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *DateComponentResponse) SetStartDate(v string) {
	o.StartDate = &v
}

func (o DateComponentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PreModerationStatus) {
		toSerialize["preModerationStatus"] = o.PreModerationStatus
	}
	if !IsNil(o.ComponentType) {
		toSerialize["componentType"] = o.ComponentType
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.PolicyViolations) {
		toSerialize["policyViolations"] = o.PolicyViolations
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	return toSerialize, nil
}

type NullableDateComponentResponse struct {
	value *DateComponentResponse
	isSet bool
}

func (v NullableDateComponentResponse) Get() *DateComponentResponse {
	return v.value
}

func (v *NullableDateComponentResponse) Set(val *DateComponentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDateComponentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDateComponentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDateComponentResponse(val *DateComponentResponse) *NullableDateComponentResponse {
	return &NullableDateComponentResponse{value: val, isSet: true}
}

func (v NullableDateComponentResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDateComponentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
