package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateBrandVideoCreative type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateBrandVideoCreative{}

// CreateBrandVideoCreative struct for CreateBrandVideoCreative
type CreateBrandVideoCreative struct {
	Asins         []string       `json:"asins,omitempty"`
	BrandLogoCrop *BrandLogoCrop `json:"brandLogoCrop,omitempty"`
	BrandName     *string        `json:"brandName,omitempty"`
	// If set to true and the headline and/or video are not in the marketplace's default language, Amazon will attempt to translate them to the marketplace's default language. If Amazon is unable to translate them, the ad will be rejected by moderation. We only support translating headlines and videos from English to German, French, Italian, Spanish, Japanese, and Dutch. See developer notes for more information.
	ConsentToTranslate *bool    `json:"consentToTranslate,omitempty"`
	VideoAssetIds      []string `json:"videoAssetIds,omitempty"`
	BrandLogoAssetID   *string  `json:"brandLogoAssetID,omitempty"`
	// The headline text. Maximum length of the string is 50 characters for all marketplaces other than Japan, which has a maximum length of 35 characters.
	Headline *string `json:"headline,omitempty"`
}

// NewCreateBrandVideoCreative instantiates a new CreateBrandVideoCreative object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateBrandVideoCreative() *CreateBrandVideoCreative {
	this := CreateBrandVideoCreative{}
	return &this
}

// NewCreateBrandVideoCreativeWithDefaults instantiates a new CreateBrandVideoCreative object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateBrandVideoCreativeWithDefaults() *CreateBrandVideoCreative {
	this := CreateBrandVideoCreative{}
	return &this
}

// GetAsins returns the Asins field value if set, zero value otherwise.
func (o *CreateBrandVideoCreative) GetAsins() []string {
	if o == nil || IsNil(o.Asins) {
		var ret []string
		return ret
	}
	return o.Asins
}

// GetAsinsOk returns a tuple with the Asins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBrandVideoCreative) GetAsinsOk() ([]string, bool) {
	if o == nil || IsNil(o.Asins) {
		return nil, false
	}
	return o.Asins, true
}

// HasAsins returns a boolean if a field has been set.
func (o *CreateBrandVideoCreative) HasAsins() bool {
	if o != nil && !IsNil(o.Asins) {
		return true
	}

	return false
}

// SetAsins gets a reference to the given []string and assigns it to the Asins field.
func (o *CreateBrandVideoCreative) SetAsins(v []string) {
	o.Asins = v
}

// GetBrandLogoCrop returns the BrandLogoCrop field value if set, zero value otherwise.
func (o *CreateBrandVideoCreative) GetBrandLogoCrop() BrandLogoCrop {
	if o == nil || IsNil(o.BrandLogoCrop) {
		var ret BrandLogoCrop
		return ret
	}
	return *o.BrandLogoCrop
}

// GetBrandLogoCropOk returns a tuple with the BrandLogoCrop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBrandVideoCreative) GetBrandLogoCropOk() (*BrandLogoCrop, bool) {
	if o == nil || IsNil(o.BrandLogoCrop) {
		return nil, false
	}
	return o.BrandLogoCrop, true
}

// HasBrandLogoCrop returns a boolean if a field has been set.
func (o *CreateBrandVideoCreative) HasBrandLogoCrop() bool {
	if o != nil && !IsNil(o.BrandLogoCrop) {
		return true
	}

	return false
}

// SetBrandLogoCrop gets a reference to the given BrandLogoCrop and assigns it to the BrandLogoCrop field.
func (o *CreateBrandVideoCreative) SetBrandLogoCrop(v BrandLogoCrop) {
	o.BrandLogoCrop = &v
}

// GetBrandName returns the BrandName field value if set, zero value otherwise.
func (o *CreateBrandVideoCreative) GetBrandName() string {
	if o == nil || IsNil(o.BrandName) {
		var ret string
		return ret
	}
	return *o.BrandName
}

// GetBrandNameOk returns a tuple with the BrandName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBrandVideoCreative) GetBrandNameOk() (*string, bool) {
	if o == nil || IsNil(o.BrandName) {
		return nil, false
	}
	return o.BrandName, true
}

// HasBrandName returns a boolean if a field has been set.
func (o *CreateBrandVideoCreative) HasBrandName() bool {
	if o != nil && !IsNil(o.BrandName) {
		return true
	}

	return false
}

// SetBrandName gets a reference to the given string and assigns it to the BrandName field.
func (o *CreateBrandVideoCreative) SetBrandName(v string) {
	o.BrandName = &v
}

// GetConsentToTranslate returns the ConsentToTranslate field value if set, zero value otherwise.
func (o *CreateBrandVideoCreative) GetConsentToTranslate() bool {
	if o == nil || IsNil(o.ConsentToTranslate) {
		var ret bool
		return ret
	}
	return *o.ConsentToTranslate
}

// GetConsentToTranslateOk returns a tuple with the ConsentToTranslate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBrandVideoCreative) GetConsentToTranslateOk() (*bool, bool) {
	if o == nil || IsNil(o.ConsentToTranslate) {
		return nil, false
	}
	return o.ConsentToTranslate, true
}

// HasConsentToTranslate returns a boolean if a field has been set.
func (o *CreateBrandVideoCreative) HasConsentToTranslate() bool {
	if o != nil && !IsNil(o.ConsentToTranslate) {
		return true
	}

	return false
}

// SetConsentToTranslate gets a reference to the given bool and assigns it to the ConsentToTranslate field.
func (o *CreateBrandVideoCreative) SetConsentToTranslate(v bool) {
	o.ConsentToTranslate = &v
}

// GetVideoAssetIds returns the VideoAssetIds field value if set, zero value otherwise.
func (o *CreateBrandVideoCreative) GetVideoAssetIds() []string {
	if o == nil || IsNil(o.VideoAssetIds) {
		var ret []string
		return ret
	}
	return o.VideoAssetIds
}

// GetVideoAssetIdsOk returns a tuple with the VideoAssetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBrandVideoCreative) GetVideoAssetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.VideoAssetIds) {
		return nil, false
	}
	return o.VideoAssetIds, true
}

// HasVideoAssetIds returns a boolean if a field has been set.
func (o *CreateBrandVideoCreative) HasVideoAssetIds() bool {
	if o != nil && !IsNil(o.VideoAssetIds) {
		return true
	}

	return false
}

// SetVideoAssetIds gets a reference to the given []string and assigns it to the VideoAssetIds field.
func (o *CreateBrandVideoCreative) SetVideoAssetIds(v []string) {
	o.VideoAssetIds = v
}

// GetBrandLogoAssetID returns the BrandLogoAssetID field value if set, zero value otherwise.
func (o *CreateBrandVideoCreative) GetBrandLogoAssetID() string {
	if o == nil || IsNil(o.BrandLogoAssetID) {
		var ret string
		return ret
	}
	return *o.BrandLogoAssetID
}

// GetBrandLogoAssetIDOk returns a tuple with the BrandLogoAssetID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBrandVideoCreative) GetBrandLogoAssetIDOk() (*string, bool) {
	if o == nil || IsNil(o.BrandLogoAssetID) {
		return nil, false
	}
	return o.BrandLogoAssetID, true
}

// HasBrandLogoAssetID returns a boolean if a field has been set.
func (o *CreateBrandVideoCreative) HasBrandLogoAssetID() bool {
	if o != nil && !IsNil(o.BrandLogoAssetID) {
		return true
	}

	return false
}

// SetBrandLogoAssetID gets a reference to the given string and assigns it to the BrandLogoAssetID field.
func (o *CreateBrandVideoCreative) SetBrandLogoAssetID(v string) {
	o.BrandLogoAssetID = &v
}

// GetHeadline returns the Headline field value if set, zero value otherwise.
func (o *CreateBrandVideoCreative) GetHeadline() string {
	if o == nil || IsNil(o.Headline) {
		var ret string
		return ret
	}
	return *o.Headline
}

// GetHeadlineOk returns a tuple with the Headline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBrandVideoCreative) GetHeadlineOk() (*string, bool) {
	if o == nil || IsNil(o.Headline) {
		return nil, false
	}
	return o.Headline, true
}

// HasHeadline returns a boolean if a field has been set.
func (o *CreateBrandVideoCreative) HasHeadline() bool {
	if o != nil && !IsNil(o.Headline) {
		return true
	}

	return false
}

// SetHeadline gets a reference to the given string and assigns it to the Headline field.
func (o *CreateBrandVideoCreative) SetHeadline(v string) {
	o.Headline = &v
}

func (o CreateBrandVideoCreative) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Asins) {
		toSerialize["asins"] = o.Asins
	}
	if !IsNil(o.BrandLogoCrop) {
		toSerialize["brandLogoCrop"] = o.BrandLogoCrop
	}
	if !IsNil(o.BrandName) {
		toSerialize["brandName"] = o.BrandName
	}
	if !IsNil(o.ConsentToTranslate) {
		toSerialize["consentToTranslate"] = o.ConsentToTranslate
	}
	if !IsNil(o.VideoAssetIds) {
		toSerialize["videoAssetIds"] = o.VideoAssetIds
	}
	if !IsNil(o.BrandLogoAssetID) {
		toSerialize["brandLogoAssetID"] = o.BrandLogoAssetID
	}
	if !IsNil(o.Headline) {
		toSerialize["headline"] = o.Headline
	}
	return toSerialize, nil
}

type NullableCreateBrandVideoCreative struct {
	value *CreateBrandVideoCreative
	isSet bool
}

func (v NullableCreateBrandVideoCreative) Get() *CreateBrandVideoCreative {
	return v.value
}

func (v *NullableCreateBrandVideoCreative) Set(val *CreateBrandVideoCreative) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateBrandVideoCreative) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateBrandVideoCreative) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateBrandVideoCreative(val *CreateBrandVideoCreative) *NullableCreateBrandVideoCreative {
	return &NullableCreateBrandVideoCreative{value: val, isSet: true}
}

func (v NullableCreateBrandVideoCreative) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateBrandVideoCreative) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
