package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateProductCollectionCreative type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateProductCollectionCreative{}

// CreateProductCollectionCreative struct for CreateProductCollectionCreative
type CreateProductCollectionCreative struct {
	BrandLogoCrop      *BrandLogoCrop   `json:"brandLogoCrop,omitempty"`
	Asins              []string         `json:"asins,omitempty"`
	BrandName          *string          `json:"brandName,omitempty"`
	CustomImageAssetId *string          `json:"customImageAssetId,omitempty"`
	CustomImageCrop    *CustomImageCrop `json:"customImageCrop,omitempty"`
	BrandLogoAssetID   *string          `json:"brandLogoAssetID,omitempty"`
	// The headline text. Maximum length of the string is 50 characters for all marketplaces other than Japan, which has a maximum length of 35 characters.
	Headline *string `json:"headline,omitempty"`
}

// NewCreateProductCollectionCreative instantiates a new CreateProductCollectionCreative object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProductCollectionCreative() *CreateProductCollectionCreative {
	this := CreateProductCollectionCreative{}
	return &this
}

// NewCreateProductCollectionCreativeWithDefaults instantiates a new CreateProductCollectionCreative object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProductCollectionCreativeWithDefaults() *CreateProductCollectionCreative {
	this := CreateProductCollectionCreative{}
	return &this
}

// GetBrandLogoCrop returns the BrandLogoCrop field value if set, zero value otherwise.
func (o *CreateProductCollectionCreative) GetBrandLogoCrop() BrandLogoCrop {
	if o == nil || IsNil(o.BrandLogoCrop) {
		var ret BrandLogoCrop
		return ret
	}
	return *o.BrandLogoCrop
}

// GetBrandLogoCropOk returns a tuple with the BrandLogoCrop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductCollectionCreative) GetBrandLogoCropOk() (*BrandLogoCrop, bool) {
	if o == nil || IsNil(o.BrandLogoCrop) {
		return nil, false
	}
	return o.BrandLogoCrop, true
}

// HasBrandLogoCrop returns a boolean if a field has been set.
func (o *CreateProductCollectionCreative) HasBrandLogoCrop() bool {
	if o != nil && !IsNil(o.BrandLogoCrop) {
		return true
	}

	return false
}

// SetBrandLogoCrop gets a reference to the given BrandLogoCrop and assigns it to the BrandLogoCrop field.
func (o *CreateProductCollectionCreative) SetBrandLogoCrop(v BrandLogoCrop) {
	o.BrandLogoCrop = &v
}

// GetAsins returns the Asins field value if set, zero value otherwise.
func (o *CreateProductCollectionCreative) GetAsins() []string {
	if o == nil || IsNil(o.Asins) {
		var ret []string
		return ret
	}
	return o.Asins
}

// GetAsinsOk returns a tuple with the Asins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductCollectionCreative) GetAsinsOk() ([]string, bool) {
	if o == nil || IsNil(o.Asins) {
		return nil, false
	}
	return o.Asins, true
}

// HasAsins returns a boolean if a field has been set.
func (o *CreateProductCollectionCreative) HasAsins() bool {
	if o != nil && !IsNil(o.Asins) {
		return true
	}

	return false
}

// SetAsins gets a reference to the given []string and assigns it to the Asins field.
func (o *CreateProductCollectionCreative) SetAsins(v []string) {
	o.Asins = v
}

// GetBrandName returns the BrandName field value if set, zero value otherwise.
func (o *CreateProductCollectionCreative) GetBrandName() string {
	if o == nil || IsNil(o.BrandName) {
		var ret string
		return ret
	}
	return *o.BrandName
}

// GetBrandNameOk returns a tuple with the BrandName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductCollectionCreative) GetBrandNameOk() (*string, bool) {
	if o == nil || IsNil(o.BrandName) {
		return nil, false
	}
	return o.BrandName, true
}

// HasBrandName returns a boolean if a field has been set.
func (o *CreateProductCollectionCreative) HasBrandName() bool {
	if o != nil && !IsNil(o.BrandName) {
		return true
	}

	return false
}

// SetBrandName gets a reference to the given string and assigns it to the BrandName field.
func (o *CreateProductCollectionCreative) SetBrandName(v string) {
	o.BrandName = &v
}

// GetCustomImageAssetId returns the CustomImageAssetId field value if set, zero value otherwise.
func (o *CreateProductCollectionCreative) GetCustomImageAssetId() string {
	if o == nil || IsNil(o.CustomImageAssetId) {
		var ret string
		return ret
	}
	return *o.CustomImageAssetId
}

// GetCustomImageAssetIdOk returns a tuple with the CustomImageAssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductCollectionCreative) GetCustomImageAssetIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomImageAssetId) {
		return nil, false
	}
	return o.CustomImageAssetId, true
}

// HasCustomImageAssetId returns a boolean if a field has been set.
func (o *CreateProductCollectionCreative) HasCustomImageAssetId() bool {
	if o != nil && !IsNil(o.CustomImageAssetId) {
		return true
	}

	return false
}

// SetCustomImageAssetId gets a reference to the given string and assigns it to the CustomImageAssetId field.
func (o *CreateProductCollectionCreative) SetCustomImageAssetId(v string) {
	o.CustomImageAssetId = &v
}

// GetCustomImageCrop returns the CustomImageCrop field value if set, zero value otherwise.
func (o *CreateProductCollectionCreative) GetCustomImageCrop() CustomImageCrop {
	if o == nil || IsNil(o.CustomImageCrop) {
		var ret CustomImageCrop
		return ret
	}
	return *o.CustomImageCrop
}

// GetCustomImageCropOk returns a tuple with the CustomImageCrop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductCollectionCreative) GetCustomImageCropOk() (*CustomImageCrop, bool) {
	if o == nil || IsNil(o.CustomImageCrop) {
		return nil, false
	}
	return o.CustomImageCrop, true
}

// HasCustomImageCrop returns a boolean if a field has been set.
func (o *CreateProductCollectionCreative) HasCustomImageCrop() bool {
	if o != nil && !IsNil(o.CustomImageCrop) {
		return true
	}

	return false
}

// SetCustomImageCrop gets a reference to the given CustomImageCrop and assigns it to the CustomImageCrop field.
func (o *CreateProductCollectionCreative) SetCustomImageCrop(v CustomImageCrop) {
	o.CustomImageCrop = &v
}

// GetBrandLogoAssetID returns the BrandLogoAssetID field value if set, zero value otherwise.
func (o *CreateProductCollectionCreative) GetBrandLogoAssetID() string {
	if o == nil || IsNil(o.BrandLogoAssetID) {
		var ret string
		return ret
	}
	return *o.BrandLogoAssetID
}

// GetBrandLogoAssetIDOk returns a tuple with the BrandLogoAssetID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductCollectionCreative) GetBrandLogoAssetIDOk() (*string, bool) {
	if o == nil || IsNil(o.BrandLogoAssetID) {
		return nil, false
	}
	return o.BrandLogoAssetID, true
}

// HasBrandLogoAssetID returns a boolean if a field has been set.
func (o *CreateProductCollectionCreative) HasBrandLogoAssetID() bool {
	if o != nil && !IsNil(o.BrandLogoAssetID) {
		return true
	}

	return false
}

// SetBrandLogoAssetID gets a reference to the given string and assigns it to the BrandLogoAssetID field.
func (o *CreateProductCollectionCreative) SetBrandLogoAssetID(v string) {
	o.BrandLogoAssetID = &v
}

// GetHeadline returns the Headline field value if set, zero value otherwise.
func (o *CreateProductCollectionCreative) GetHeadline() string {
	if o == nil || IsNil(o.Headline) {
		var ret string
		return ret
	}
	return *o.Headline
}

// GetHeadlineOk returns a tuple with the Headline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductCollectionCreative) GetHeadlineOk() (*string, bool) {
	if o == nil || IsNil(o.Headline) {
		return nil, false
	}
	return o.Headline, true
}

// HasHeadline returns a boolean if a field has been set.
func (o *CreateProductCollectionCreative) HasHeadline() bool {
	if o != nil && !IsNil(o.Headline) {
		return true
	}

	return false
}

// SetHeadline gets a reference to the given string and assigns it to the Headline field.
func (o *CreateProductCollectionCreative) SetHeadline(v string) {
	o.Headline = &v
}

func (o CreateProductCollectionCreative) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BrandLogoCrop) {
		toSerialize["brandLogoCrop"] = o.BrandLogoCrop
	}
	if !IsNil(o.Asins) {
		toSerialize["asins"] = o.Asins
	}
	if !IsNil(o.BrandName) {
		toSerialize["brandName"] = o.BrandName
	}
	if !IsNil(o.CustomImageAssetId) {
		toSerialize["customImageAssetId"] = o.CustomImageAssetId
	}
	if !IsNil(o.CustomImageCrop) {
		toSerialize["customImageCrop"] = o.CustomImageCrop
	}
	if !IsNil(o.BrandLogoAssetID) {
		toSerialize["brandLogoAssetID"] = o.BrandLogoAssetID
	}
	if !IsNil(o.Headline) {
		toSerialize["headline"] = o.Headline
	}
	return toSerialize, nil
}

type NullableCreateProductCollectionCreative struct {
	value *CreateProductCollectionCreative
	isSet bool
}

func (v NullableCreateProductCollectionCreative) Get() *CreateProductCollectionCreative {
	return v.value
}

func (v *NullableCreateProductCollectionCreative) Set(val *CreateProductCollectionCreative) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProductCollectionCreative) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProductCollectionCreative) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProductCollectionCreative(val *CreateProductCollectionCreative) *NullableCreateProductCollectionCreative {
	return &NullableCreateProductCollectionCreative{value: val, isSet: true}
}

func (v NullableCreateProductCollectionCreative) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateProductCollectionCreative) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
