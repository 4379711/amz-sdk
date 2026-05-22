package product_metadata

import (
	"github.com/bytedance/sonic"
)

// checks if the TvPreviewMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TvPreviewMetadata{}

// TvPreviewMetadata struct for TvPreviewMetadata
type TvPreviewMetadata struct {
	ShortSynopsis   *string `json:"shortSynopsis,omitempty"`
	TvBackgroundURL *string `json:"tvBackgroundURL,omitempty"`
	TvIconURL       *string `json:"tvIconURL,omitempty"`
	LinkIn          *string `json:"linkIn,omitempty"`
	Title           *string `json:"title,omitempty"`
}

// NewTvPreviewMetadata instantiates a new TvPreviewMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTvPreviewMetadata() *TvPreviewMetadata {
	this := TvPreviewMetadata{}
	return &this
}

// NewTvPreviewMetadataWithDefaults instantiates a new TvPreviewMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTvPreviewMetadataWithDefaults() *TvPreviewMetadata {
	this := TvPreviewMetadata{}
	return &this
}

// GetShortSynopsis returns the ShortSynopsis field value if set, zero value otherwise.
func (o *TvPreviewMetadata) GetShortSynopsis() string {
	if o == nil || IsNil(o.ShortSynopsis) {
		var ret string
		return ret
	}
	return *o.ShortSynopsis
}

// GetShortSynopsisOk returns a tuple with the ShortSynopsis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvPreviewMetadata) GetShortSynopsisOk() (*string, bool) {
	if o == nil || IsNil(o.ShortSynopsis) {
		return nil, false
	}
	return o.ShortSynopsis, true
}

// HasShortSynopsis returns a boolean if a field has been set.
func (o *TvPreviewMetadata) HasShortSynopsis() bool {
	if o != nil && !IsNil(o.ShortSynopsis) {
		return true
	}

	return false
}

// SetShortSynopsis gets a reference to the given string and assigns it to the ShortSynopsis field.
func (o *TvPreviewMetadata) SetShortSynopsis(v string) {
	o.ShortSynopsis = &v
}

// GetTvBackgroundURL returns the TvBackgroundURL field value if set, zero value otherwise.
func (o *TvPreviewMetadata) GetTvBackgroundURL() string {
	if o == nil || IsNil(o.TvBackgroundURL) {
		var ret string
		return ret
	}
	return *o.TvBackgroundURL
}

// GetTvBackgroundURLOk returns a tuple with the TvBackgroundURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvPreviewMetadata) GetTvBackgroundURLOk() (*string, bool) {
	if o == nil || IsNil(o.TvBackgroundURL) {
		return nil, false
	}
	return o.TvBackgroundURL, true
}

// HasTvBackgroundURL returns a boolean if a field has been set.
func (o *TvPreviewMetadata) HasTvBackgroundURL() bool {
	if o != nil && !IsNil(o.TvBackgroundURL) {
		return true
	}

	return false
}

// SetTvBackgroundURL gets a reference to the given string and assigns it to the TvBackgroundURL field.
func (o *TvPreviewMetadata) SetTvBackgroundURL(v string) {
	o.TvBackgroundURL = &v
}

// GetTvIconURL returns the TvIconURL field value if set, zero value otherwise.
func (o *TvPreviewMetadata) GetTvIconURL() string {
	if o == nil || IsNil(o.TvIconURL) {
		var ret string
		return ret
	}
	return *o.TvIconURL
}

// GetTvIconURLOk returns a tuple with the TvIconURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvPreviewMetadata) GetTvIconURLOk() (*string, bool) {
	if o == nil || IsNil(o.TvIconURL) {
		return nil, false
	}
	return o.TvIconURL, true
}

// HasTvIconURL returns a boolean if a field has been set.
func (o *TvPreviewMetadata) HasTvIconURL() bool {
	if o != nil && !IsNil(o.TvIconURL) {
		return true
	}

	return false
}

// SetTvIconURL gets a reference to the given string and assigns it to the TvIconURL field.
func (o *TvPreviewMetadata) SetTvIconURL(v string) {
	o.TvIconURL = &v
}

// GetLinkIn returns the LinkIn field value if set, zero value otherwise.
func (o *TvPreviewMetadata) GetLinkIn() string {
	if o == nil || IsNil(o.LinkIn) {
		var ret string
		return ret
	}
	return *o.LinkIn
}

// GetLinkInOk returns a tuple with the LinkIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvPreviewMetadata) GetLinkInOk() (*string, bool) {
	if o == nil || IsNil(o.LinkIn) {
		return nil, false
	}
	return o.LinkIn, true
}

// HasLinkIn returns a boolean if a field has been set.
func (o *TvPreviewMetadata) HasLinkIn() bool {
	if o != nil && !IsNil(o.LinkIn) {
		return true
	}

	return false
}

// SetLinkIn gets a reference to the given string and assigns it to the LinkIn field.
func (o *TvPreviewMetadata) SetLinkIn(v string) {
	o.LinkIn = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *TvPreviewMetadata) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TvPreviewMetadata) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *TvPreviewMetadata) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *TvPreviewMetadata) SetTitle(v string) {
	o.Title = &v
}

func (o TvPreviewMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ShortSynopsis) {
		toSerialize["shortSynopsis"] = o.ShortSynopsis
	}
	if !IsNil(o.TvBackgroundURL) {
		toSerialize["tvBackgroundURL"] = o.TvBackgroundURL
	}
	if !IsNil(o.TvIconURL) {
		toSerialize["tvIconURL"] = o.TvIconURL
	}
	if !IsNil(o.LinkIn) {
		toSerialize["linkIn"] = o.LinkIn
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableTvPreviewMetadata struct {
	value *TvPreviewMetadata
	isSet bool
}

func (v NullableTvPreviewMetadata) Get() *TvPreviewMetadata {
	return v.value
}

func (v *NullableTvPreviewMetadata) Set(val *TvPreviewMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableTvPreviewMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableTvPreviewMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTvPreviewMetadata(val *TvPreviewMetadata) *NullableTvPreviewMetadata {
	return &NullableTvPreviewMetadata{value: val, isSet: true}
}

func (v NullableTvPreviewMetadata) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTvPreviewMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
