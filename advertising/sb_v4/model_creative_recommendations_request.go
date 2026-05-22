package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreativeRecommendationsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreativeRecommendationsRequest{}

// CreativeRecommendationsRequest Request structure of creative recommendations API.
type CreativeRecommendationsRequest struct {
	// An array of ASINs associated with the creative. Note, do not pass an empty array, this results in an error.
	Asins []string `json:"asins"`
	// Ad format of the creative.
	AdFormat string `json:"adFormat"`
	// Required recommendations details.
	RequiredRecommendations []RequiredRecommendations `json:"requiredRecommendations"`
}

type _CreativeRecommendationsRequest CreativeRecommendationsRequest

// NewCreativeRecommendationsRequest instantiates a new CreativeRecommendationsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreativeRecommendationsRequest(asins []string, adFormat string, requiredRecommendations []RequiredRecommendations) *CreativeRecommendationsRequest {
	this := CreativeRecommendationsRequest{}
	this.Asins = asins
	this.AdFormat = adFormat
	this.RequiredRecommendations = requiredRecommendations
	return &this
}

// NewCreativeRecommendationsRequestWithDefaults instantiates a new CreativeRecommendationsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreativeRecommendationsRequestWithDefaults() *CreativeRecommendationsRequest {
	this := CreativeRecommendationsRequest{}
	return &this
}

// GetAsins returns the Asins field value
func (o *CreativeRecommendationsRequest) GetAsins() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Asins
}

// GetAsinsOk returns a tuple with the Asins field value
// and a boolean to check if the value has been set.
func (o *CreativeRecommendationsRequest) GetAsinsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Asins, true
}

// SetAsins sets field value
func (o *CreativeRecommendationsRequest) SetAsins(v []string) {
	o.Asins = v
}

// GetAdFormat returns the AdFormat field value
func (o *CreativeRecommendationsRequest) GetAdFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdFormat
}

// GetAdFormatOk returns a tuple with the AdFormat field value
// and a boolean to check if the value has been set.
func (o *CreativeRecommendationsRequest) GetAdFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdFormat, true
}

// SetAdFormat sets field value
func (o *CreativeRecommendationsRequest) SetAdFormat(v string) {
	o.AdFormat = v
}

// GetRequiredRecommendations returns the RequiredRecommendations field value
func (o *CreativeRecommendationsRequest) GetRequiredRecommendations() []RequiredRecommendations {
	if o == nil {
		var ret []RequiredRecommendations
		return ret
	}

	return o.RequiredRecommendations
}

// GetRequiredRecommendationsOk returns a tuple with the RequiredRecommendations field value
// and a boolean to check if the value has been set.
func (o *CreativeRecommendationsRequest) GetRequiredRecommendationsOk() ([]RequiredRecommendations, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequiredRecommendations, true
}

// SetRequiredRecommendations sets field value
func (o *CreativeRecommendationsRequest) SetRequiredRecommendations(v []RequiredRecommendations) {
	o.RequiredRecommendations = v
}

func (o CreativeRecommendationsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["asins"] = o.Asins
	toSerialize["adFormat"] = o.AdFormat
	toSerialize["requiredRecommendations"] = o.RequiredRecommendations
	return toSerialize, nil
}

type NullableCreativeRecommendationsRequest struct {
	value *CreativeRecommendationsRequest
	isSet bool
}

func (v NullableCreativeRecommendationsRequest) Get() *CreativeRecommendationsRequest {
	return v.value
}

func (v *NullableCreativeRecommendationsRequest) Set(val *CreativeRecommendationsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreativeRecommendationsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreativeRecommendationsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreativeRecommendationsRequest(val *CreativeRecommendationsRequest) *NullableCreativeRecommendationsRequest {
	return &NullableCreativeRecommendationsRequest{value: val, isSet: true}
}

func (v NullableCreativeRecommendationsRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreativeRecommendationsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
