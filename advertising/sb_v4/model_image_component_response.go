package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ImageComponentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageComponentResponse{}

// ImageComponentResponse Pre moderation result for a image component
type ImageComponentResponse struct {
	// The pre moderation status of the component.
	PreModerationStatus *string `json:"preModerationStatus,omitempty"`
	// Type of the image component.
	ComponentType *string      `json:"componentType,omitempty"`
	LandingPage   *LandingPage `json:"landingPage,omitempty"`
	// A list of policy violations for the component that were detected during pre moderation. Note that this field is present in the response only when preModerationStatus is set to REJECTED.
	PolicyViolations []ImagePolicyViolation `json:"policyViolations,omitempty"`
	// Id of the component. This is the same id sent as part of the request. This can be used to uniquely identify the component.
	Id *string `json:"id,omitempty"`
	// Publicly accessible url of the image that got pre moderated.
	Url *string `json:"url,omitempty"`
}

// NewImageComponentResponse instantiates a new ImageComponentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageComponentResponse() *ImageComponentResponse {
	this := ImageComponentResponse{}
	return &this
}

// NewImageComponentResponseWithDefaults instantiates a new ImageComponentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageComponentResponseWithDefaults() *ImageComponentResponse {
	this := ImageComponentResponse{}
	return &this
}

// GetPreModerationStatus returns the PreModerationStatus field value if set, zero value otherwise.
func (o *ImageComponentResponse) GetPreModerationStatus() string {
	if o == nil || IsNil(o.PreModerationStatus) {
		var ret string
		return ret
	}
	return *o.PreModerationStatus
}

// GetPreModerationStatusOk returns a tuple with the PreModerationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageComponentResponse) GetPreModerationStatusOk() (*string, bool) {
	if o == nil || IsNil(o.PreModerationStatus) {
		return nil, false
	}
	return o.PreModerationStatus, true
}

// HasPreModerationStatus returns a boolean if a field has been set.
func (o *ImageComponentResponse) HasPreModerationStatus() bool {
	if o != nil && !IsNil(o.PreModerationStatus) {
		return true
	}

	return false
}

// SetPreModerationStatus gets a reference to the given string and assigns it to the PreModerationStatus field.
func (o *ImageComponentResponse) SetPreModerationStatus(v string) {
	o.PreModerationStatus = &v
}

// GetComponentType returns the ComponentType field value if set, zero value otherwise.
func (o *ImageComponentResponse) GetComponentType() string {
	if o == nil || IsNil(o.ComponentType) {
		var ret string
		return ret
	}
	return *o.ComponentType
}

// GetComponentTypeOk returns a tuple with the ComponentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageComponentResponse) GetComponentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ComponentType) {
		return nil, false
	}
	return o.ComponentType, true
}

// HasComponentType returns a boolean if a field has been set.
func (o *ImageComponentResponse) HasComponentType() bool {
	if o != nil && !IsNil(o.ComponentType) {
		return true
	}

	return false
}

// SetComponentType gets a reference to the given string and assigns it to the ComponentType field.
func (o *ImageComponentResponse) SetComponentType(v string) {
	o.ComponentType = &v
}

// GetLandingPage returns the LandingPage field value if set, zero value otherwise.
func (o *ImageComponentResponse) GetLandingPage() LandingPage {
	if o == nil || IsNil(o.LandingPage) {
		var ret LandingPage
		return ret
	}
	return *o.LandingPage
}

// GetLandingPageOk returns a tuple with the LandingPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageComponentResponse) GetLandingPageOk() (*LandingPage, bool) {
	if o == nil || IsNil(o.LandingPage) {
		return nil, false
	}
	return o.LandingPage, true
}

// HasLandingPage returns a boolean if a field has been set.
func (o *ImageComponentResponse) HasLandingPage() bool {
	if o != nil && !IsNil(o.LandingPage) {
		return true
	}

	return false
}

// SetLandingPage gets a reference to the given LandingPage and assigns it to the LandingPage field.
func (o *ImageComponentResponse) SetLandingPage(v LandingPage) {
	o.LandingPage = &v
}

// GetPolicyViolations returns the PolicyViolations field value if set, zero value otherwise.
func (o *ImageComponentResponse) GetPolicyViolations() []ImagePolicyViolation {
	if o == nil || IsNil(o.PolicyViolations) {
		var ret []ImagePolicyViolation
		return ret
	}
	return o.PolicyViolations
}

// GetPolicyViolationsOk returns a tuple with the PolicyViolations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageComponentResponse) GetPolicyViolationsOk() ([]ImagePolicyViolation, bool) {
	if o == nil || IsNil(o.PolicyViolations) {
		return nil, false
	}
	return o.PolicyViolations, true
}

// HasPolicyViolations returns a boolean if a field has been set.
func (o *ImageComponentResponse) HasPolicyViolations() bool {
	if o != nil && !IsNil(o.PolicyViolations) {
		return true
	}

	return false
}

// SetPolicyViolations gets a reference to the given []ImagePolicyViolation and assigns it to the PolicyViolations field.
func (o *ImageComponentResponse) SetPolicyViolations(v []ImagePolicyViolation) {
	o.PolicyViolations = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ImageComponentResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageComponentResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ImageComponentResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ImageComponentResponse) SetId(v string) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ImageComponentResponse) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageComponentResponse) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ImageComponentResponse) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ImageComponentResponse) SetUrl(v string) {
	o.Url = &v
}

func (o ImageComponentResponse) ToMap() (map[string]interface{}, error) {
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

type NullableImageComponentResponse struct {
	value *ImageComponentResponse
	isSet bool
}

func (v NullableImageComponentResponse) Get() *ImageComponentResponse {
	return v.value
}

func (v *NullableImageComponentResponse) Set(val *ImageComponentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableImageComponentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableImageComponentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageComponentResponse(val *ImageComponentResponse) *NullableImageComponentResponse {
	return &NullableImageComponentResponse{value: val, isSet: true}
}

func (v NullableImageComponentResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableImageComponentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
