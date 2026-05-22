package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the PreModerationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PreModerationResponse{}

// PreModerationResponse Information regarding the policy violations if present for the components, sent for pre moderation.
type PreModerationResponse struct {
	// Id of the brand/advertiser.
	RecordId *string `json:"recordId,omitempty"`
	// Pre moderation result of the asin components. It will have information regarding the policy violations present if any.
	AsinComponents []AsinComponentResponse `json:"asinComponents,omitempty"`
	// Unique Id for the moderation Request.
	PreModerationId *string `json:"preModerationId,omitempty"`
	// Type of Ad program to which the pre moderation components belong to.
	AdProgram *string `json:"adProgram,omitempty"`
	// Locale value that was passed in request.
	Locale *string `json:"locale,omitempty"`
	// Pre moderation result of the image components. It will have information regarding the policy violations present if any.
	ImageComponents []ImageComponentResponse `json:"imageComponents,omitempty"`
	// Pre moderation result of the date components. It will have information regarding the policy violations present if any.
	DateComponents []DateComponentResponse `json:"dateComponents,omitempty"`
	// Pre moderation result of the text components. It will have information regarding the policy violations present if any.
	TextComponents []TextComponentResponse `json:"textComponents,omitempty"`
	// Pre moderation result of the video components. It will have information regarding the policy violations present if any.
	VideoComponents []VideoComponentResponse `json:"videoComponents,omitempty"`
}

// NewPreModerationResponse instantiates a new PreModerationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreModerationResponse() *PreModerationResponse {
	this := PreModerationResponse{}
	return &this
}

// NewPreModerationResponseWithDefaults instantiates a new PreModerationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreModerationResponseWithDefaults() *PreModerationResponse {
	this := PreModerationResponse{}
	return &this
}

// GetRecordId returns the RecordId field value if set, zero value otherwise.
func (o *PreModerationResponse) GetRecordId() string {
	if o == nil || IsNil(o.RecordId) {
		var ret string
		return ret
	}
	return *o.RecordId
}

// GetRecordIdOk returns a tuple with the RecordId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreModerationResponse) GetRecordIdOk() (*string, bool) {
	if o == nil || IsNil(o.RecordId) {
		return nil, false
	}
	return o.RecordId, true
}

// HasRecordId returns a boolean if a field has been set.
func (o *PreModerationResponse) HasRecordId() bool {
	if o != nil && !IsNil(o.RecordId) {
		return true
	}

	return false
}

// SetRecordId gets a reference to the given string and assigns it to the RecordId field.
func (o *PreModerationResponse) SetRecordId(v string) {
	o.RecordId = &v
}

// GetAsinComponents returns the AsinComponents field value if set, zero value otherwise.
func (o *PreModerationResponse) GetAsinComponents() []AsinComponentResponse {
	if o == nil || IsNil(o.AsinComponents) {
		var ret []AsinComponentResponse
		return ret
	}
	return o.AsinComponents
}

// GetAsinComponentsOk returns a tuple with the AsinComponents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreModerationResponse) GetAsinComponentsOk() ([]AsinComponentResponse, bool) {
	if o == nil || IsNil(o.AsinComponents) {
		return nil, false
	}
	return o.AsinComponents, true
}

// HasAsinComponents returns a boolean if a field has been set.
func (o *PreModerationResponse) HasAsinComponents() bool {
	if o != nil && !IsNil(o.AsinComponents) {
		return true
	}

	return false
}

// SetAsinComponents gets a reference to the given []AsinComponentResponse and assigns it to the AsinComponents field.
func (o *PreModerationResponse) SetAsinComponents(v []AsinComponentResponse) {
	o.AsinComponents = v
}

// GetPreModerationId returns the PreModerationId field value if set, zero value otherwise.
func (o *PreModerationResponse) GetPreModerationId() string {
	if o == nil || IsNil(o.PreModerationId) {
		var ret string
		return ret
	}
	return *o.PreModerationId
}

// GetPreModerationIdOk returns a tuple with the PreModerationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreModerationResponse) GetPreModerationIdOk() (*string, bool) {
	if o == nil || IsNil(o.PreModerationId) {
		return nil, false
	}
	return o.PreModerationId, true
}

// HasPreModerationId returns a boolean if a field has been set.
func (o *PreModerationResponse) HasPreModerationId() bool {
	if o != nil && !IsNil(o.PreModerationId) {
		return true
	}

	return false
}

// SetPreModerationId gets a reference to the given string and assigns it to the PreModerationId field.
func (o *PreModerationResponse) SetPreModerationId(v string) {
	o.PreModerationId = &v
}

// GetAdProgram returns the AdProgram field value if set, zero value otherwise.
func (o *PreModerationResponse) GetAdProgram() string {
	if o == nil || IsNil(o.AdProgram) {
		var ret string
		return ret
	}
	return *o.AdProgram
}

// GetAdProgramOk returns a tuple with the AdProgram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreModerationResponse) GetAdProgramOk() (*string, bool) {
	if o == nil || IsNil(o.AdProgram) {
		return nil, false
	}
	return o.AdProgram, true
}

// HasAdProgram returns a boolean if a field has been set.
func (o *PreModerationResponse) HasAdProgram() bool {
	if o != nil && !IsNil(o.AdProgram) {
		return true
	}

	return false
}

// SetAdProgram gets a reference to the given string and assigns it to the AdProgram field.
func (o *PreModerationResponse) SetAdProgram(v string) {
	o.AdProgram = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *PreModerationResponse) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreModerationResponse) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *PreModerationResponse) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *PreModerationResponse) SetLocale(v string) {
	o.Locale = &v
}

// GetImageComponents returns the ImageComponents field value if set, zero value otherwise.
func (o *PreModerationResponse) GetImageComponents() []ImageComponentResponse {
	if o == nil || IsNil(o.ImageComponents) {
		var ret []ImageComponentResponse
		return ret
	}
	return o.ImageComponents
}

// GetImageComponentsOk returns a tuple with the ImageComponents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreModerationResponse) GetImageComponentsOk() ([]ImageComponentResponse, bool) {
	if o == nil || IsNil(o.ImageComponents) {
		return nil, false
	}
	return o.ImageComponents, true
}

// HasImageComponents returns a boolean if a field has been set.
func (o *PreModerationResponse) HasImageComponents() bool {
	if o != nil && !IsNil(o.ImageComponents) {
		return true
	}

	return false
}

// SetImageComponents gets a reference to the given []ImageComponentResponse and assigns it to the ImageComponents field.
func (o *PreModerationResponse) SetImageComponents(v []ImageComponentResponse) {
	o.ImageComponents = v
}

// GetDateComponents returns the DateComponents field value if set, zero value otherwise.
func (o *PreModerationResponse) GetDateComponents() []DateComponentResponse {
	if o == nil || IsNil(o.DateComponents) {
		var ret []DateComponentResponse
		return ret
	}
	return o.DateComponents
}

// GetDateComponentsOk returns a tuple with the DateComponents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreModerationResponse) GetDateComponentsOk() ([]DateComponentResponse, bool) {
	if o == nil || IsNil(o.DateComponents) {
		return nil, false
	}
	return o.DateComponents, true
}

// HasDateComponents returns a boolean if a field has been set.
func (o *PreModerationResponse) HasDateComponents() bool {
	if o != nil && !IsNil(o.DateComponents) {
		return true
	}

	return false
}

// SetDateComponents gets a reference to the given []DateComponentResponse and assigns it to the DateComponents field.
func (o *PreModerationResponse) SetDateComponents(v []DateComponentResponse) {
	o.DateComponents = v
}

// GetTextComponents returns the TextComponents field value if set, zero value otherwise.
func (o *PreModerationResponse) GetTextComponents() []TextComponentResponse {
	if o == nil || IsNil(o.TextComponents) {
		var ret []TextComponentResponse
		return ret
	}
	return o.TextComponents
}

// GetTextComponentsOk returns a tuple with the TextComponents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreModerationResponse) GetTextComponentsOk() ([]TextComponentResponse, bool) {
	if o == nil || IsNil(o.TextComponents) {
		return nil, false
	}
	return o.TextComponents, true
}

// HasTextComponents returns a boolean if a field has been set.
func (o *PreModerationResponse) HasTextComponents() bool {
	if o != nil && !IsNil(o.TextComponents) {
		return true
	}

	return false
}

// SetTextComponents gets a reference to the given []TextComponentResponse and assigns it to the TextComponents field.
func (o *PreModerationResponse) SetTextComponents(v []TextComponentResponse) {
	o.TextComponents = v
}

// GetVideoComponents returns the VideoComponents field value if set, zero value otherwise.
func (o *PreModerationResponse) GetVideoComponents() []VideoComponentResponse {
	if o == nil || IsNil(o.VideoComponents) {
		var ret []VideoComponentResponse
		return ret
	}
	return o.VideoComponents
}

// GetVideoComponentsOk returns a tuple with the VideoComponents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreModerationResponse) GetVideoComponentsOk() ([]VideoComponentResponse, bool) {
	if o == nil || IsNil(o.VideoComponents) {
		return nil, false
	}
	return o.VideoComponents, true
}

// HasVideoComponents returns a boolean if a field has been set.
func (o *PreModerationResponse) HasVideoComponents() bool {
	if o != nil && !IsNil(o.VideoComponents) {
		return true
	}

	return false
}

// SetVideoComponents gets a reference to the given []VideoComponentResponse and assigns it to the VideoComponents field.
func (o *PreModerationResponse) SetVideoComponents(v []VideoComponentResponse) {
	o.VideoComponents = v
}

func (o PreModerationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RecordId) {
		toSerialize["recordId"] = o.RecordId
	}
	if !IsNil(o.AsinComponents) {
		toSerialize["asinComponents"] = o.AsinComponents
	}
	if !IsNil(o.PreModerationId) {
		toSerialize["preModerationId"] = o.PreModerationId
	}
	if !IsNil(o.AdProgram) {
		toSerialize["adProgram"] = o.AdProgram
	}
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	if !IsNil(o.ImageComponents) {
		toSerialize["imageComponents"] = o.ImageComponents
	}
	if !IsNil(o.DateComponents) {
		toSerialize["dateComponents"] = o.DateComponents
	}
	if !IsNil(o.TextComponents) {
		toSerialize["textComponents"] = o.TextComponents
	}
	if !IsNil(o.VideoComponents) {
		toSerialize["videoComponents"] = o.VideoComponents
	}
	return toSerialize, nil
}

type NullablePreModerationResponse struct {
	value *PreModerationResponse
	isSet bool
}

func (v NullablePreModerationResponse) Get() *PreModerationResponse {
	return v.value
}

func (v *NullablePreModerationResponse) Set(val *PreModerationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePreModerationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePreModerationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreModerationResponse(val *PreModerationResponse) *NullablePreModerationResponse {
	return &NullablePreModerationResponse{value: val, isSet: true}
}

func (v NullablePreModerationResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePreModerationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
