package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SDHeadlineRecommendationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDHeadlineRecommendationRequest{}

// SDHeadlineRecommendationRequest Request structure of SD headline recommendation API.
type SDHeadlineRecommendationRequest struct {
	// An array of ASINs associated with the creative.
	Asins []string `json:"asins,omitempty"`
	// Maximum number of recommendations that API should return. Response will [0, maxNumRecommendations] recommendations (recommendations are not guaranteed as there can be instances where the ML model can not generate policy compliant headlines for the given set of asins).
	MaxNumRecommendations *float32 `json:"maxNumRecommendations,omitempty"`
	AdFormat              *string  `json:"adFormat,omitempty"`
}

// NewSDHeadlineRecommendationRequest instantiates a new SDHeadlineRecommendationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDHeadlineRecommendationRequest() *SDHeadlineRecommendationRequest {
	this := SDHeadlineRecommendationRequest{}
	return &this
}

// NewSDHeadlineRecommendationRequestWithDefaults instantiates a new SDHeadlineRecommendationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDHeadlineRecommendationRequestWithDefaults() *SDHeadlineRecommendationRequest {
	this := SDHeadlineRecommendationRequest{}
	return &this
}

// GetAsins returns the Asins field value if set, zero value otherwise.
func (o *SDHeadlineRecommendationRequest) GetAsins() []string {
	if o == nil || IsNil(o.Asins) {
		var ret []string
		return ret
	}
	return o.Asins
}

// GetAsinsOk returns a tuple with the Asins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDHeadlineRecommendationRequest) GetAsinsOk() ([]string, bool) {
	if o == nil || IsNil(o.Asins) {
		return nil, false
	}
	return o.Asins, true
}

// HasAsins returns a boolean if a field has been set.
func (o *SDHeadlineRecommendationRequest) HasAsins() bool {
	if o != nil && !IsNil(o.Asins) {
		return true
	}

	return false
}

// SetAsins gets a reference to the given []string and assigns it to the Asins field.
func (o *SDHeadlineRecommendationRequest) SetAsins(v []string) {
	o.Asins = v
}

// GetMaxNumRecommendations returns the MaxNumRecommendations field value if set, zero value otherwise.
func (o *SDHeadlineRecommendationRequest) GetMaxNumRecommendations() float32 {
	if o == nil || IsNil(o.MaxNumRecommendations) {
		var ret float32
		return ret
	}
	return *o.MaxNumRecommendations
}

// GetMaxNumRecommendationsOk returns a tuple with the MaxNumRecommendations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDHeadlineRecommendationRequest) GetMaxNumRecommendationsOk() (*float32, bool) {
	if o == nil || IsNil(o.MaxNumRecommendations) {
		return nil, false
	}
	return o.MaxNumRecommendations, true
}

// HasMaxNumRecommendations returns a boolean if a field has been set.
func (o *SDHeadlineRecommendationRequest) HasMaxNumRecommendations() bool {
	if o != nil && !IsNil(o.MaxNumRecommendations) {
		return true
	}

	return false
}

// SetMaxNumRecommendations gets a reference to the given float32 and assigns it to the MaxNumRecommendations field.
func (o *SDHeadlineRecommendationRequest) SetMaxNumRecommendations(v float32) {
	o.MaxNumRecommendations = &v
}

// GetAdFormat returns the AdFormat field value if set, zero value otherwise.
func (o *SDHeadlineRecommendationRequest) GetAdFormat() string {
	if o == nil || IsNil(o.AdFormat) {
		var ret string
		return ret
	}
	return *o.AdFormat
}

// GetAdFormatOk returns a tuple with the AdFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDHeadlineRecommendationRequest) GetAdFormatOk() (*string, bool) {
	if o == nil || IsNil(o.AdFormat) {
		return nil, false
	}
	return o.AdFormat, true
}

// HasAdFormat returns a boolean if a field has been set.
func (o *SDHeadlineRecommendationRequest) HasAdFormat() bool {
	if o != nil && !IsNil(o.AdFormat) {
		return true
	}

	return false
}

// SetAdFormat gets a reference to the given string and assigns it to the AdFormat field.
func (o *SDHeadlineRecommendationRequest) SetAdFormat(v string) {
	o.AdFormat = &v
}

func (o SDHeadlineRecommendationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Asins) {
		toSerialize["asins"] = o.Asins
	}
	if !IsNil(o.MaxNumRecommendations) {
		toSerialize["maxNumRecommendations"] = o.MaxNumRecommendations
	}
	if !IsNil(o.AdFormat) {
		toSerialize["adFormat"] = o.AdFormat
	}
	return toSerialize, nil
}

type NullableSDHeadlineRecommendationRequest struct {
	value *SDHeadlineRecommendationRequest
	isSet bool
}

func (v NullableSDHeadlineRecommendationRequest) Get() *SDHeadlineRecommendationRequest {
	return v.value
}

func (v *NullableSDHeadlineRecommendationRequest) Set(val *SDHeadlineRecommendationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSDHeadlineRecommendationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSDHeadlineRecommendationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDHeadlineRecommendationRequest(val *SDHeadlineRecommendationRequest) *NullableSDHeadlineRecommendationRequest {
	return &NullableSDHeadlineRecommendationRequest{value: val, isSet: true}
}

func (v NullableSDHeadlineRecommendationRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDHeadlineRecommendationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
