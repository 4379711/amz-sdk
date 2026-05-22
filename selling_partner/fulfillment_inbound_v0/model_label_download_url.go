package fulfillment_inbound_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the LabelDownloadURL type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LabelDownloadURL{}

// LabelDownloadURL Download URL for a label
type LabelDownloadURL struct {
	// URL to download the label for the package. Note: The URL will only be valid for 15 seconds
	DownloadURL *string `json:"DownloadURL,omitempty"`
}

// NewLabelDownloadURL instantiates a new LabelDownloadURL object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabelDownloadURL() *LabelDownloadURL {
	this := LabelDownloadURL{}
	return &this
}

// NewLabelDownloadURLWithDefaults instantiates a new LabelDownloadURL object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabelDownloadURLWithDefaults() *LabelDownloadURL {
	this := LabelDownloadURL{}
	return &this
}

// GetDownloadURL returns the DownloadURL field value if set, zero value otherwise.
func (o *LabelDownloadURL) GetDownloadURL() string {
	if o == nil || IsNil(o.DownloadURL) {
		var ret string
		return ret
	}
	return *o.DownloadURL
}

// GetDownloadURLOk returns a tuple with the DownloadURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelDownloadURL) GetDownloadURLOk() (*string, bool) {
	if o == nil || IsNil(o.DownloadURL) {
		return nil, false
	}
	return o.DownloadURL, true
}

// HasDownloadURL returns a boolean if a field has been set.
func (o *LabelDownloadURL) HasDownloadURL() bool {
	if o != nil && !IsNil(o.DownloadURL) {
		return true
	}

	return false
}

// SetDownloadURL gets a reference to the given string and assigns it to the DownloadURL field.
func (o *LabelDownloadURL) SetDownloadURL(v string) {
	o.DownloadURL = &v
}

func (o LabelDownloadURL) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DownloadURL) {
		toSerialize["DownloadURL"] = o.DownloadURL
	}
	return toSerialize, nil
}

type NullableLabelDownloadURL struct {
	value *LabelDownloadURL
	isSet bool
}

func (v NullableLabelDownloadURL) Get() *LabelDownloadURL {
	return v.value
}

func (v *NullableLabelDownloadURL) Set(val *LabelDownloadURL) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelDownloadURL) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelDownloadURL) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelDownloadURL(val *LabelDownloadURL) *NullableLabelDownloadURL {
	return &NullableLabelDownloadURL{value: val, isSet: true}
}

func (v NullableLabelDownloadURL) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableLabelDownloadURL) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
