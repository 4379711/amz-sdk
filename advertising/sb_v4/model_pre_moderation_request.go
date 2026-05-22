package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the PreModerationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PreModerationRequest{}

// PreModerationRequest Components details that needs to be sent for pre moderation.
type PreModerationRequest struct {
	// Id of the brand/advertiser.
	RecordId *string `json:"recordId,omitempty"`
	// Asin components which needs to be pre moderated.
	AsinComponents []AsinComponent `json:"asinComponents,omitempty"`
	// Type of Ad program to which this pre moderation components belong to.
	AdProgram string `json:"adProgram"`
	// Specifying locale will translate the premoderation message into that locale's associated language.     | Locale | Language (ISO 639) | Country (ISO 3166) |   |-----|-----|-------|   | ar-AE | Arabic (ar) | United Arab Emirates (AE) |   | zh-CN | Chinese (zh) | China (CN) |   | nl-NL | Dutch (nl) | Netherlands (NL) |   | en-AU | English (en) | Australia (AU) |   | en-CA | English (en) | Canada (CA) |   | en-IN | English (en) | India (IN) |   | en-GB | English (en) | United Kingdom (GB) |   | en-US | English (en) | United States (US) |   | fr-CA | French (fr) | Canada (CA) |   | fr-FR | French (fr) | France (FR) |   | de-DE | German (de) | Germany (DE) |   | it-IT | Italian (it) | Italy (IT) |   | ja-JP | Japanese (ja) | Japan (JP) |   | ko-KR | Korean (ko) | South Korea (KR) |   | pt-BR | Portuguese (pt) | Brazil (BR) |   | es-ES | Spanish (es) | Spain (ES) |   | es-US | Spanish (es) | United States (US) |   | es-MX | Spanish (es) | Mexico (MX) |   | tr-TR | Turkish (tr) | Turkey (TR) |
	Locale string `json:"locale"`
	// Image components which needs to be pre moderated.
	ImageComponents []ImageComponent `json:"imageComponents,omitempty"`
	// Date components which needs to be pre moderated.
	DateComponents []DateComponent `json:"dateComponents,omitempty"`
	// Text components which needs to be pre moderated.
	TextComponents []TextComponent `json:"textComponents,omitempty"`
	// Video components which needs to be pre moderated.
	VideoComponents []VideoComponent `json:"videoComponents,omitempty"`
}

type _PreModerationRequest PreModerationRequest

// NewPreModerationRequest instantiates a new PreModerationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreModerationRequest(adProgram string, locale string) *PreModerationRequest {
	this := PreModerationRequest{}
	this.AdProgram = adProgram
	this.Locale = locale
	return &this
}

// NewPreModerationRequestWithDefaults instantiates a new PreModerationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreModerationRequestWithDefaults() *PreModerationRequest {
	this := PreModerationRequest{}
	return &this
}

// GetRecordId returns the RecordId field value if set, zero value otherwise.
func (o *PreModerationRequest) GetRecordId() string {
	if o == nil || IsNil(o.RecordId) {
		var ret string
		return ret
	}
	return *o.RecordId
}

// GetRecordIdOk returns a tuple with the RecordId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreModerationRequest) GetRecordIdOk() (*string, bool) {
	if o == nil || IsNil(o.RecordId) {
		return nil, false
	}
	return o.RecordId, true
}

// HasRecordId returns a boolean if a field has been set.
func (o *PreModerationRequest) HasRecordId() bool {
	if o != nil && !IsNil(o.RecordId) {
		return true
	}

	return false
}

// SetRecordId gets a reference to the given string and assigns it to the RecordId field.
func (o *PreModerationRequest) SetRecordId(v string) {
	o.RecordId = &v
}

// GetAsinComponents returns the AsinComponents field value if set, zero value otherwise.
func (o *PreModerationRequest) GetAsinComponents() []AsinComponent {
	if o == nil || IsNil(o.AsinComponents) {
		var ret []AsinComponent
		return ret
	}
	return o.AsinComponents
}

// GetAsinComponentsOk returns a tuple with the AsinComponents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreModerationRequest) GetAsinComponentsOk() ([]AsinComponent, bool) {
	if o == nil || IsNil(o.AsinComponents) {
		return nil, false
	}
	return o.AsinComponents, true
}

// HasAsinComponents returns a boolean if a field has been set.
func (o *PreModerationRequest) HasAsinComponents() bool {
	if o != nil && !IsNil(o.AsinComponents) {
		return true
	}

	return false
}

// SetAsinComponents gets a reference to the given []AsinComponent and assigns it to the AsinComponents field.
func (o *PreModerationRequest) SetAsinComponents(v []AsinComponent) {
	o.AsinComponents = v
}

// GetAdProgram returns the AdProgram field value
func (o *PreModerationRequest) GetAdProgram() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdProgram
}

// GetAdProgramOk returns a tuple with the AdProgram field value
// and a boolean to check if the value has been set.
func (o *PreModerationRequest) GetAdProgramOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdProgram, true
}

// SetAdProgram sets field value
func (o *PreModerationRequest) SetAdProgram(v string) {
	o.AdProgram = v
}

// GetLocale returns the Locale field value
func (o *PreModerationRequest) GetLocale() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value
// and a boolean to check if the value has been set.
func (o *PreModerationRequest) GetLocaleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locale, true
}

// SetLocale sets field value
func (o *PreModerationRequest) SetLocale(v string) {
	o.Locale = v
}

// GetImageComponents returns the ImageComponents field value if set, zero value otherwise.
func (o *PreModerationRequest) GetImageComponents() []ImageComponent {
	if o == nil || IsNil(o.ImageComponents) {
		var ret []ImageComponent
		return ret
	}
	return o.ImageComponents
}

// GetImageComponentsOk returns a tuple with the ImageComponents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreModerationRequest) GetImageComponentsOk() ([]ImageComponent, bool) {
	if o == nil || IsNil(o.ImageComponents) {
		return nil, false
	}
	return o.ImageComponents, true
}

// HasImageComponents returns a boolean if a field has been set.
func (o *PreModerationRequest) HasImageComponents() bool {
	if o != nil && !IsNil(o.ImageComponents) {
		return true
	}

	return false
}

// SetImageComponents gets a reference to the given []ImageComponent and assigns it to the ImageComponents field.
func (o *PreModerationRequest) SetImageComponents(v []ImageComponent) {
	o.ImageComponents = v
}

// GetDateComponents returns the DateComponents field value if set, zero value otherwise.
func (o *PreModerationRequest) GetDateComponents() []DateComponent {
	if o == nil || IsNil(o.DateComponents) {
		var ret []DateComponent
		return ret
	}
	return o.DateComponents
}

// GetDateComponentsOk returns a tuple with the DateComponents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreModerationRequest) GetDateComponentsOk() ([]DateComponent, bool) {
	if o == nil || IsNil(o.DateComponents) {
		return nil, false
	}
	return o.DateComponents, true
}

// HasDateComponents returns a boolean if a field has been set.
func (o *PreModerationRequest) HasDateComponents() bool {
	if o != nil && !IsNil(o.DateComponents) {
		return true
	}

	return false
}

// SetDateComponents gets a reference to the given []DateComponent and assigns it to the DateComponents field.
func (o *PreModerationRequest) SetDateComponents(v []DateComponent) {
	o.DateComponents = v
}

// GetTextComponents returns the TextComponents field value if set, zero value otherwise.
func (o *PreModerationRequest) GetTextComponents() []TextComponent {
	if o == nil || IsNil(o.TextComponents) {
		var ret []TextComponent
		return ret
	}
	return o.TextComponents
}

// GetTextComponentsOk returns a tuple with the TextComponents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreModerationRequest) GetTextComponentsOk() ([]TextComponent, bool) {
	if o == nil || IsNil(o.TextComponents) {
		return nil, false
	}
	return o.TextComponents, true
}

// HasTextComponents returns a boolean if a field has been set.
func (o *PreModerationRequest) HasTextComponents() bool {
	if o != nil && !IsNil(o.TextComponents) {
		return true
	}

	return false
}

// SetTextComponents gets a reference to the given []TextComponent and assigns it to the TextComponents field.
func (o *PreModerationRequest) SetTextComponents(v []TextComponent) {
	o.TextComponents = v
}

// GetVideoComponents returns the VideoComponents field value if set, zero value otherwise.
func (o *PreModerationRequest) GetVideoComponents() []VideoComponent {
	if o == nil || IsNil(o.VideoComponents) {
		var ret []VideoComponent
		return ret
	}
	return o.VideoComponents
}

// GetVideoComponentsOk returns a tuple with the VideoComponents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreModerationRequest) GetVideoComponentsOk() ([]VideoComponent, bool) {
	if o == nil || IsNil(o.VideoComponents) {
		return nil, false
	}
	return o.VideoComponents, true
}

// HasVideoComponents returns a boolean if a field has been set.
func (o *PreModerationRequest) HasVideoComponents() bool {
	if o != nil && !IsNil(o.VideoComponents) {
		return true
	}

	return false
}

// SetVideoComponents gets a reference to the given []VideoComponent and assigns it to the VideoComponents field.
func (o *PreModerationRequest) SetVideoComponents(v []VideoComponent) {
	o.VideoComponents = v
}

func (o PreModerationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RecordId) {
		toSerialize["recordId"] = o.RecordId
	}
	if !IsNil(o.AsinComponents) {
		toSerialize["asinComponents"] = o.AsinComponents
	}
	toSerialize["adProgram"] = o.AdProgram
	toSerialize["locale"] = o.Locale
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

type NullablePreModerationRequest struct {
	value *PreModerationRequest
	isSet bool
}

func (v NullablePreModerationRequest) Get() *PreModerationRequest {
	return v.value
}

func (v *NullablePreModerationRequest) Set(val *PreModerationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePreModerationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePreModerationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreModerationRequest(val *PreModerationRequest) *NullablePreModerationRequest {
	return &NullablePreModerationRequest{value: val, isSet: true}
}

func (v NullablePreModerationRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePreModerationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
