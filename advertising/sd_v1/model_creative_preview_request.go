package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the CreativePreviewRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreativePreviewRequest{}

// CreativePreviewRequest struct for CreativePreviewRequest
type CreativePreviewRequest struct {
	Creative              PreviewCreativeModel           `json:"creative"`
	PreviewConfiguration  CreativePreviewConfiguration   `json:"previewConfiguration"`
	PreviewConfigurations []CreativePreviewConfiguration `json:"previewConfigurations,omitempty"`
}

type _CreativePreviewRequest CreativePreviewRequest

// NewCreativePreviewRequest instantiates a new CreativePreviewRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreativePreviewRequest(creative PreviewCreativeModel, previewConfiguration CreativePreviewConfiguration) *CreativePreviewRequest {
	this := CreativePreviewRequest{}
	this.Creative = creative
	this.PreviewConfiguration = previewConfiguration
	return &this
}

// NewCreativePreviewRequestWithDefaults instantiates a new CreativePreviewRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreativePreviewRequestWithDefaults() *CreativePreviewRequest {
	this := CreativePreviewRequest{}
	return &this
}

// GetCreative returns the Creative field value
func (o *CreativePreviewRequest) GetCreative() PreviewCreativeModel {
	if o == nil {
		var ret PreviewCreativeModel
		return ret
	}

	return o.Creative
}

// GetCreativeOk returns a tuple with the Creative field value
// and a boolean to check if the value has been set.
func (o *CreativePreviewRequest) GetCreativeOk() (*PreviewCreativeModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Creative, true
}

// SetCreative sets field value
func (o *CreativePreviewRequest) SetCreative(v PreviewCreativeModel) {
	o.Creative = v
}

// GetPreviewConfiguration returns the PreviewConfiguration field value
func (o *CreativePreviewRequest) GetPreviewConfiguration() CreativePreviewConfiguration {
	if o == nil {
		var ret CreativePreviewConfiguration
		return ret
	}

	return o.PreviewConfiguration
}

// GetPreviewConfigurationOk returns a tuple with the PreviewConfiguration field value
// and a boolean to check if the value has been set.
func (o *CreativePreviewRequest) GetPreviewConfigurationOk() (*CreativePreviewConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreviewConfiguration, true
}

// SetPreviewConfiguration sets field value
func (o *CreativePreviewRequest) SetPreviewConfiguration(v CreativePreviewConfiguration) {
	o.PreviewConfiguration = v
}

// GetPreviewConfigurations returns the PreviewConfigurations field value if set, zero value otherwise.
func (o *CreativePreviewRequest) GetPreviewConfigurations() []CreativePreviewConfiguration {
	if o == nil || IsNil(o.PreviewConfigurations) {
		var ret []CreativePreviewConfiguration
		return ret
	}
	return o.PreviewConfigurations
}

// GetPreviewConfigurationsOk returns a tuple with the PreviewConfigurations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativePreviewRequest) GetPreviewConfigurationsOk() ([]CreativePreviewConfiguration, bool) {
	if o == nil || IsNil(o.PreviewConfigurations) {
		return nil, false
	}
	return o.PreviewConfigurations, true
}

// HasPreviewConfigurations returns a boolean if a field has been set.
func (o *CreativePreviewRequest) HasPreviewConfigurations() bool {
	if o != nil && !IsNil(o.PreviewConfigurations) {
		return true
	}

	return false
}

// SetPreviewConfigurations gets a reference to the given []CreativePreviewConfiguration and assigns it to the PreviewConfigurations field.
func (o *CreativePreviewRequest) SetPreviewConfigurations(v []CreativePreviewConfiguration) {
	o.PreviewConfigurations = v
}

func (o CreativePreviewRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["creative"] = o.Creative
	toSerialize["previewConfiguration"] = o.PreviewConfiguration
	if !IsNil(o.PreviewConfigurations) {
		toSerialize["previewConfigurations"] = o.PreviewConfigurations
	}
	return toSerialize, nil
}

type NullableCreativePreviewRequest struct {
	value *CreativePreviewRequest
	isSet bool
}

func (v NullableCreativePreviewRequest) Get() *CreativePreviewRequest {
	return v.value
}

func (v *NullableCreativePreviewRequest) Set(val *CreativePreviewRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreativePreviewRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreativePreviewRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreativePreviewRequest(val *CreativePreviewRequest) *NullableCreativePreviewRequest {
	return &NullableCreativePreviewRequest{value: val, isSet: true}
}

func (v NullableCreativePreviewRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreativePreviewRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
