package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateStoreSpotlightCreative type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateStoreSpotlightCreative{}

// CreateStoreSpotlightCreative struct for CreateStoreSpotlightCreative
type CreateStoreSpotlightCreative struct {
	BrandLogoCrop *BrandLogoCrop `json:"brandLogoCrop,omitempty"`
	BrandName     *string        `json:"brandName,omitempty"`
	Subpages      []Subpage      `json:"subpages,omitempty"`
	// If set to true and the headline and/or video are not in the marketplace's default language, Amazon will attempt to translate them to the marketplace's default language. If Amazon is unable to translate them, the ad will be rejected by moderation. We only support translating headlines and videos from English to German, French, Italian, Spanish, Japanese, and Dutch. See developer notes for more information.
	ConsentToTranslate *bool   `json:"consentToTranslate,omitempty"`
	BrandLogoAssetID   *string `json:"brandLogoAssetID,omitempty"`
	// The headline text. Maximum length of the string is 50 characters for all marketplaces other than Japan, which has a maximum length of 35 characters.
	Headline *string `json:"headline,omitempty"`
}

// NewCreateStoreSpotlightCreative instantiates a new CreateStoreSpotlightCreative object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateStoreSpotlightCreative() *CreateStoreSpotlightCreative {
	this := CreateStoreSpotlightCreative{}
	return &this
}

// NewCreateStoreSpotlightCreativeWithDefaults instantiates a new CreateStoreSpotlightCreative object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateStoreSpotlightCreativeWithDefaults() *CreateStoreSpotlightCreative {
	this := CreateStoreSpotlightCreative{}
	return &this
}

// GetBrandLogoCrop returns the BrandLogoCrop field value if set, zero value otherwise.
func (o *CreateStoreSpotlightCreative) GetBrandLogoCrop() BrandLogoCrop {
	if o == nil || IsNil(o.BrandLogoCrop) {
		var ret BrandLogoCrop
		return ret
	}
	return *o.BrandLogoCrop
}

// GetBrandLogoCropOk returns a tuple with the BrandLogoCrop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStoreSpotlightCreative) GetBrandLogoCropOk() (*BrandLogoCrop, bool) {
	if o == nil || IsNil(o.BrandLogoCrop) {
		return nil, false
	}
	return o.BrandLogoCrop, true
}

// HasBrandLogoCrop returns a boolean if a field has been set.
func (o *CreateStoreSpotlightCreative) HasBrandLogoCrop() bool {
	if o != nil && !IsNil(o.BrandLogoCrop) {
		return true
	}

	return false
}

// SetBrandLogoCrop gets a reference to the given BrandLogoCrop and assigns it to the BrandLogoCrop field.
func (o *CreateStoreSpotlightCreative) SetBrandLogoCrop(v BrandLogoCrop) {
	o.BrandLogoCrop = &v
}

// GetBrandName returns the BrandName field value if set, zero value otherwise.
func (o *CreateStoreSpotlightCreative) GetBrandName() string {
	if o == nil || IsNil(o.BrandName) {
		var ret string
		return ret
	}
	return *o.BrandName
}

// GetBrandNameOk returns a tuple with the BrandName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStoreSpotlightCreative) GetBrandNameOk() (*string, bool) {
	if o == nil || IsNil(o.BrandName) {
		return nil, false
	}
	return o.BrandName, true
}

// HasBrandName returns a boolean if a field has been set.
func (o *CreateStoreSpotlightCreative) HasBrandName() bool {
	if o != nil && !IsNil(o.BrandName) {
		return true
	}

	return false
}

// SetBrandName gets a reference to the given string and assigns it to the BrandName field.
func (o *CreateStoreSpotlightCreative) SetBrandName(v string) {
	o.BrandName = &v
}

// GetSubpages returns the Subpages field value if set, zero value otherwise.
func (o *CreateStoreSpotlightCreative) GetSubpages() []Subpage {
	if o == nil || IsNil(o.Subpages) {
		var ret []Subpage
		return ret
	}
	return o.Subpages
}

// GetSubpagesOk returns a tuple with the Subpages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStoreSpotlightCreative) GetSubpagesOk() ([]Subpage, bool) {
	if o == nil || IsNil(o.Subpages) {
		return nil, false
	}
	return o.Subpages, true
}

// HasSubpages returns a boolean if a field has been set.
func (o *CreateStoreSpotlightCreative) HasSubpages() bool {
	if o != nil && !IsNil(o.Subpages) {
		return true
	}

	return false
}

// SetSubpages gets a reference to the given []Subpage and assigns it to the Subpages field.
func (o *CreateStoreSpotlightCreative) SetSubpages(v []Subpage) {
	o.Subpages = v
}

// GetConsentToTranslate returns the ConsentToTranslate field value if set, zero value otherwise.
func (o *CreateStoreSpotlightCreative) GetConsentToTranslate() bool {
	if o == nil || IsNil(o.ConsentToTranslate) {
		var ret bool
		return ret
	}
	return *o.ConsentToTranslate
}

// GetConsentToTranslateOk returns a tuple with the ConsentToTranslate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStoreSpotlightCreative) GetConsentToTranslateOk() (*bool, bool) {
	if o == nil || IsNil(o.ConsentToTranslate) {
		return nil, false
	}
	return o.ConsentToTranslate, true
}

// HasConsentToTranslate returns a boolean if a field has been set.
func (o *CreateStoreSpotlightCreative) HasConsentToTranslate() bool {
	if o != nil && !IsNil(o.ConsentToTranslate) {
		return true
	}

	return false
}

// SetConsentToTranslate gets a reference to the given bool and assigns it to the ConsentToTranslate field.
func (o *CreateStoreSpotlightCreative) SetConsentToTranslate(v bool) {
	o.ConsentToTranslate = &v
}

// GetBrandLogoAssetID returns the BrandLogoAssetID field value if set, zero value otherwise.
func (o *CreateStoreSpotlightCreative) GetBrandLogoAssetID() string {
	if o == nil || IsNil(o.BrandLogoAssetID) {
		var ret string
		return ret
	}
	return *o.BrandLogoAssetID
}

// GetBrandLogoAssetIDOk returns a tuple with the BrandLogoAssetID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStoreSpotlightCreative) GetBrandLogoAssetIDOk() (*string, bool) {
	if o == nil || IsNil(o.BrandLogoAssetID) {
		return nil, false
	}
	return o.BrandLogoAssetID, true
}

// HasBrandLogoAssetID returns a boolean if a field has been set.
func (o *CreateStoreSpotlightCreative) HasBrandLogoAssetID() bool {
	if o != nil && !IsNil(o.BrandLogoAssetID) {
		return true
	}

	return false
}

// SetBrandLogoAssetID gets a reference to the given string and assigns it to the BrandLogoAssetID field.
func (o *CreateStoreSpotlightCreative) SetBrandLogoAssetID(v string) {
	o.BrandLogoAssetID = &v
}

// GetHeadline returns the Headline field value if set, zero value otherwise.
func (o *CreateStoreSpotlightCreative) GetHeadline() string {
	if o == nil || IsNil(o.Headline) {
		var ret string
		return ret
	}
	return *o.Headline
}

// GetHeadlineOk returns a tuple with the Headline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStoreSpotlightCreative) GetHeadlineOk() (*string, bool) {
	if o == nil || IsNil(o.Headline) {
		return nil, false
	}
	return o.Headline, true
}

// HasHeadline returns a boolean if a field has been set.
func (o *CreateStoreSpotlightCreative) HasHeadline() bool {
	if o != nil && !IsNil(o.Headline) {
		return true
	}

	return false
}

// SetHeadline gets a reference to the given string and assigns it to the Headline field.
func (o *CreateStoreSpotlightCreative) SetHeadline(v string) {
	o.Headline = &v
}

func (o CreateStoreSpotlightCreative) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BrandLogoCrop) {
		toSerialize["brandLogoCrop"] = o.BrandLogoCrop
	}
	if !IsNil(o.BrandName) {
		toSerialize["brandName"] = o.BrandName
	}
	if !IsNil(o.Subpages) {
		toSerialize["subpages"] = o.Subpages
	}
	if !IsNil(o.ConsentToTranslate) {
		toSerialize["consentToTranslate"] = o.ConsentToTranslate
	}
	if !IsNil(o.BrandLogoAssetID) {
		toSerialize["brandLogoAssetID"] = o.BrandLogoAssetID
	}
	if !IsNil(o.Headline) {
		toSerialize["headline"] = o.Headline
	}
	return toSerialize, nil
}

type NullableCreateStoreSpotlightCreative struct {
	value *CreateStoreSpotlightCreative
	isSet bool
}

func (v NullableCreateStoreSpotlightCreative) Get() *CreateStoreSpotlightCreative {
	return v.value
}

func (v *NullableCreateStoreSpotlightCreative) Set(val *CreateStoreSpotlightCreative) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateStoreSpotlightCreative) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateStoreSpotlightCreative) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateStoreSpotlightCreative(val *CreateStoreSpotlightCreative) *NullableCreateStoreSpotlightCreative {
	return &NullableCreateStoreSpotlightCreative{value: val, isSet: true}
}

func (v NullableCreateStoreSpotlightCreative) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateStoreSpotlightCreative) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
