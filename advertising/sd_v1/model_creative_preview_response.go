package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the CreativePreviewResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreativePreviewResponse{}

// CreativePreviewResponse struct for CreativePreviewResponse
type CreativePreviewResponse struct {
	PreviewHtml  string   `json:"previewHtml"`
	PreviewHtmls []string `json:"previewHtmls,omitempty"`
}

type _CreativePreviewResponse CreativePreviewResponse

// NewCreativePreviewResponse instantiates a new CreativePreviewResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreativePreviewResponse(previewHtml string) *CreativePreviewResponse {
	this := CreativePreviewResponse{}
	this.PreviewHtml = previewHtml
	return &this
}

// NewCreativePreviewResponseWithDefaults instantiates a new CreativePreviewResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreativePreviewResponseWithDefaults() *CreativePreviewResponse {
	this := CreativePreviewResponse{}
	return &this
}

// GetPreviewHtml returns the PreviewHtml field value
func (o *CreativePreviewResponse) GetPreviewHtml() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PreviewHtml
}

// GetPreviewHtmlOk returns a tuple with the PreviewHtml field value
// and a boolean to check if the value has been set.
func (o *CreativePreviewResponse) GetPreviewHtmlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreviewHtml, true
}

// SetPreviewHtml sets field value
func (o *CreativePreviewResponse) SetPreviewHtml(v string) {
	o.PreviewHtml = v
}

// GetPreviewHtmls returns the PreviewHtmls field value if set, zero value otherwise.
func (o *CreativePreviewResponse) GetPreviewHtmls() []string {
	if o == nil || IsNil(o.PreviewHtmls) {
		var ret []string
		return ret
	}
	return o.PreviewHtmls
}

// GetPreviewHtmlsOk returns a tuple with the PreviewHtmls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativePreviewResponse) GetPreviewHtmlsOk() ([]string, bool) {
	if o == nil || IsNil(o.PreviewHtmls) {
		return nil, false
	}
	return o.PreviewHtmls, true
}

// HasPreviewHtmls returns a boolean if a field has been set.
func (o *CreativePreviewResponse) HasPreviewHtmls() bool {
	if o != nil && !IsNil(o.PreviewHtmls) {
		return true
	}

	return false
}

// SetPreviewHtmls gets a reference to the given []string and assigns it to the PreviewHtmls field.
func (o *CreativePreviewResponse) SetPreviewHtmls(v []string) {
	o.PreviewHtmls = v
}

func (o CreativePreviewResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["previewHtml"] = o.PreviewHtml
	if !IsNil(o.PreviewHtmls) {
		toSerialize["previewHtmls"] = o.PreviewHtmls
	}
	return toSerialize, nil
}

type NullableCreativePreviewResponse struct {
	value *CreativePreviewResponse
	isSet bool
}

func (v NullableCreativePreviewResponse) Get() *CreativePreviewResponse {
	return v.value
}

func (v *NullableCreativePreviewResponse) Set(val *CreativePreviewResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreativePreviewResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreativePreviewResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreativePreviewResponse(val *CreativePreviewResponse) *NullableCreativePreviewResponse {
	return &NullableCreativePreviewResponse{value: val, isSet: true}
}

func (v NullableCreativePreviewResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreativePreviewResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
