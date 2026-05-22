package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the VideoComponentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoComponentResponse{}

// VideoComponentResponse Pre moderation result for a video component
type VideoComponentResponse struct {
	// The pre moderation status of the component.
	PreModerationStatus *string `json:"preModerationStatus,omitempty"`
	// Type of the video component.
	ComponentType *string      `json:"componentType,omitempty"`
	LandingPage   *LandingPage `json:"landingPage,omitempty"`
	// A list of policy violations for the component that were detected during pre moderation. Note that this field is present in the response only when preModerationStatus is set to REJECTED.
	PolicyViolations []VideoPolicyViolation `json:"policyViolations,omitempty"`
	// Id of the component. This is the same id sent as part of the request. This can be used to uniquely identify the component.
	Id *string `json:"id,omitempty"`
	// Publicly accessible url of the video that got pre moderated.
	Url *string `json:"url,omitempty"`
}

// NewVideoComponentResponse instantiates a new VideoComponentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoComponentResponse() *VideoComponentResponse {
	this := VideoComponentResponse{}
	return &this
}

// NewVideoComponentResponseWithDefaults instantiates a new VideoComponentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoComponentResponseWithDefaults() *VideoComponentResponse {
	this := VideoComponentResponse{}
	return &this
}

// GetPreModerationStatus returns the PreModerationStatus field value if set, zero value otherwise.
func (o *VideoComponentResponse) GetPreModerationStatus() string {
	if o == nil || IsNil(o.PreModerationStatus) {
		var ret string
		return ret
	}
	return *o.PreModerationStatus
}

// GetPreModerationStatusOk returns a tuple with the PreModerationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoComponentResponse) GetPreModerationStatusOk() (*string, bool) {
	if o == nil || IsNil(o.PreModerationStatus) {
		return nil, false
	}
	return o.PreModerationStatus, true
}

// HasPreModerationStatus returns a boolean if a field has been set.
func (o *VideoComponentResponse) HasPreModerationStatus() bool {
	if o != nil && !IsNil(o.PreModerationStatus) {
		return true
	}

	return false
}

// SetPreModerationStatus gets a reference to the given string and assigns it to the PreModerationStatus field.
func (o *VideoComponentResponse) SetPreModerationStatus(v string) {
	o.PreModerationStatus = &v
}

// GetComponentType returns the ComponentType field value if set, zero value otherwise.
func (o *VideoComponentResponse) GetComponentType() string {
	if o == nil || IsNil(o.ComponentType) {
		var ret string
		return ret
	}
	return *o.ComponentType
}

// GetComponentTypeOk returns a tuple with the ComponentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoComponentResponse) GetComponentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ComponentType) {
		return nil, false
	}
	return o.ComponentType, true
}

// HasComponentType returns a boolean if a field has been set.
func (o *VideoComponentResponse) HasComponentType() bool {
	if o != nil && !IsNil(o.ComponentType) {
		return true
	}

	return false
}

// SetComponentType gets a reference to the given string and assigns it to the ComponentType field.
func (o *VideoComponentResponse) SetComponentType(v string) {
	o.ComponentType = &v
}

// GetLandingPage returns the LandingPage field value if set, zero value otherwise.
func (o *VideoComponentResponse) GetLandingPage() LandingPage {
	if o == nil || IsNil(o.LandingPage) {
		var ret LandingPage
		return ret
	}
	return *o.LandingPage
}

// GetLandingPageOk returns a tuple with the LandingPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoComponentResponse) GetLandingPageOk() (*LandingPage, bool) {
	if o == nil || IsNil(o.LandingPage) {
		return nil, false
	}
	return o.LandingPage, true
}

// HasLandingPage returns a boolean if a field has been set.
func (o *VideoComponentResponse) HasLandingPage() bool {
	if o != nil && !IsNil(o.LandingPage) {
		return true
	}

	return false
}

// SetLandingPage gets a reference to the given LandingPage and assigns it to the LandingPage field.
func (o *VideoComponentResponse) SetLandingPage(v LandingPage) {
	o.LandingPage = &v
}

// GetPolicyViolations returns the PolicyViolations field value if set, zero value otherwise.
func (o *VideoComponentResponse) GetPolicyViolations() []VideoPolicyViolation {
	if o == nil || IsNil(o.PolicyViolations) {
		var ret []VideoPolicyViolation
		return ret
	}
	return o.PolicyViolations
}

// GetPolicyViolationsOk returns a tuple with the PolicyViolations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoComponentResponse) GetPolicyViolationsOk() ([]VideoPolicyViolation, bool) {
	if o == nil || IsNil(o.PolicyViolations) {
		return nil, false
	}
	return o.PolicyViolations, true
}

// HasPolicyViolations returns a boolean if a field has been set.
func (o *VideoComponentResponse) HasPolicyViolations() bool {
	if o != nil && !IsNil(o.PolicyViolations) {
		return true
	}

	return false
}

// SetPolicyViolations gets a reference to the given []VideoPolicyViolation and assigns it to the PolicyViolations field.
func (o *VideoComponentResponse) SetPolicyViolations(v []VideoPolicyViolation) {
	o.PolicyViolations = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VideoComponentResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoComponentResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VideoComponentResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VideoComponentResponse) SetId(v string) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *VideoComponentResponse) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoComponentResponse) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *VideoComponentResponse) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *VideoComponentResponse) SetUrl(v string) {
	o.Url = &v
}

func (o VideoComponentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PreModerationStatus) {
		toSerialize["preModerationStatus"] = o.PreModerationStatus
	}
	if !IsNil(o.ComponentType) {
		toSerialize["componentType"] = o.ComponentType
	}
	if !IsNil(o.LandingPage) {
		toSerialize["landingPage"] = o.LandingPage
	}
	if !IsNil(o.PolicyViolations) {
		toSerialize["policyViolations"] = o.PolicyViolations
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableVideoComponentResponse struct {
	value *VideoComponentResponse
	isSet bool
}

func (v NullableVideoComponentResponse) Get() *VideoComponentResponse {
	return v.value
}

func (v *NullableVideoComponentResponse) Set(val *VideoComponentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoComponentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoComponentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoComponentResponse(val *VideoComponentResponse) *NullableVideoComponentResponse {
	return &NullableVideoComponentResponse{value: val, isSet: true}
}

func (v NullableVideoComponentResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableVideoComponentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
