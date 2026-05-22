package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the CreativePreviewConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreativePreviewConfiguration{}

// CreativePreviewConfiguration Optional configuration for creative preview.
type CreativePreviewConfiguration struct {
	Size *CreativePreviewConfigurationSize `json:"size,omitempty"`
	// The products to preview. Currently only the first product is previewable.
	Products []CreativePreviewConfigurationProductsInner `json:"products,omitempty"`
	// The URL where customers will land after clicking on its link. Must be provided if a LandingPageType is set. Please note that if a single product ad sets the landing page url, only one product ad can be added to the ad group. This field is not supported when using ASIN or SKU fields. ||Specifications| |------------------|------------------| |LandingPageType| Description| |STORE| The url should be in the format of https://www.amazon.com/stores/_* (using a correct Amazon url based on the marketplace)| |OFF_AMAZON_LINK| The url should be in the format of https://www.****.com. Note that this LandingPageType is not supported when using ASIN or SKU fields. A custom creative of headline, logo, image are require for this LandingPageType. | |MOMENT| Not yet supported. The url should be in the format of https://www.amazon.com/moments/promotion/{campaignId} (using a correct Amazon url based on the marketplace)|
	LandingPageURL  *string          `json:"landingPageURL,omitempty"`
	LandingPageType *LandingPageType `json:"landingPageType,omitempty"`
	// The name of the ad. Note that this field is not supported when using ASIN or SKU fields.
	AdName *string `json:"adName,omitempty"`
	// Preview the creative as if it is on a mobile environment.
	IsMobile *bool `json:"isMobile,omitempty"`
	// Preview the creative as if it is on an amazon site or third party site. The main difference is whether the preview will contain an AdChoices icon.
	IsOnAmazon *bool `json:"isOnAmazon,omitempty"`
}

// NewCreativePreviewConfiguration instantiates a new CreativePreviewConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreativePreviewConfiguration() *CreativePreviewConfiguration {
	this := CreativePreviewConfiguration{}
	return &this
}

// NewCreativePreviewConfigurationWithDefaults instantiates a new CreativePreviewConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreativePreviewConfigurationWithDefaults() *CreativePreviewConfiguration {
	this := CreativePreviewConfiguration{}
	return &this
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *CreativePreviewConfiguration) GetSize() CreativePreviewConfigurationSize {
	if o == nil || IsNil(o.Size) {
		var ret CreativePreviewConfigurationSize
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativePreviewConfiguration) GetSizeOk() (*CreativePreviewConfigurationSize, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *CreativePreviewConfiguration) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given CreativePreviewConfigurationSize and assigns it to the Size field.
func (o *CreativePreviewConfiguration) SetSize(v CreativePreviewConfigurationSize) {
	o.Size = &v
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *CreativePreviewConfiguration) GetProducts() []CreativePreviewConfigurationProductsInner {
	if o == nil || IsNil(o.Products) {
		var ret []CreativePreviewConfigurationProductsInner
		return ret
	}
	return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativePreviewConfiguration) GetProductsOk() ([]CreativePreviewConfigurationProductsInner, bool) {
	if o == nil || IsNil(o.Products) {
		return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *CreativePreviewConfiguration) HasProducts() bool {
	if o != nil && !IsNil(o.Products) {
		return true
	}

	return false
}

// SetProducts gets a reference to the given []CreativePreviewConfigurationProductsInner and assigns it to the Products field.
func (o *CreativePreviewConfiguration) SetProducts(v []CreativePreviewConfigurationProductsInner) {
	o.Products = v
}

// GetLandingPageURL returns the LandingPageURL field value if set, zero value otherwise.
func (o *CreativePreviewConfiguration) GetLandingPageURL() string {
	if o == nil || IsNil(o.LandingPageURL) {
		var ret string
		return ret
	}
	return *o.LandingPageURL
}

// GetLandingPageURLOk returns a tuple with the LandingPageURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativePreviewConfiguration) GetLandingPageURLOk() (*string, bool) {
	if o == nil || IsNil(o.LandingPageURL) {
		return nil, false
	}
	return o.LandingPageURL, true
}

// HasLandingPageURL returns a boolean if a field has been set.
func (o *CreativePreviewConfiguration) HasLandingPageURL() bool {
	if o != nil && !IsNil(o.LandingPageURL) {
		return true
	}

	return false
}

// SetLandingPageURL gets a reference to the given string and assigns it to the LandingPageURL field.
func (o *CreativePreviewConfiguration) SetLandingPageURL(v string) {
	o.LandingPageURL = &v
}

// GetLandingPageType returns the LandingPageType field value if set, zero value otherwise.
func (o *CreativePreviewConfiguration) GetLandingPageType() LandingPageType {
	if o == nil || IsNil(o.LandingPageType) {
		var ret LandingPageType
		return ret
	}
	return *o.LandingPageType
}

// GetLandingPageTypeOk returns a tuple with the LandingPageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativePreviewConfiguration) GetLandingPageTypeOk() (*LandingPageType, bool) {
	if o == nil || IsNil(o.LandingPageType) {
		return nil, false
	}
	return o.LandingPageType, true
}

// HasLandingPageType returns a boolean if a field has been set.
func (o *CreativePreviewConfiguration) HasLandingPageType() bool {
	if o != nil && !IsNil(o.LandingPageType) {
		return true
	}

	return false
}

// SetLandingPageType gets a reference to the given LandingPageType and assigns it to the LandingPageType field.
func (o *CreativePreviewConfiguration) SetLandingPageType(v LandingPageType) {
	o.LandingPageType = &v
}

// GetAdName returns the AdName field value if set, zero value otherwise.
func (o *CreativePreviewConfiguration) GetAdName() string {
	if o == nil || IsNil(o.AdName) {
		var ret string
		return ret
	}
	return *o.AdName
}

// GetAdNameOk returns a tuple with the AdName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativePreviewConfiguration) GetAdNameOk() (*string, bool) {
	if o == nil || IsNil(o.AdName) {
		return nil, false
	}
	return o.AdName, true
}

// HasAdName returns a boolean if a field has been set.
func (o *CreativePreviewConfiguration) HasAdName() bool {
	if o != nil && !IsNil(o.AdName) {
		return true
	}

	return false
}

// SetAdName gets a reference to the given string and assigns it to the AdName field.
func (o *CreativePreviewConfiguration) SetAdName(v string) {
	o.AdName = &v
}

// GetIsMobile returns the IsMobile field value if set, zero value otherwise.
func (o *CreativePreviewConfiguration) GetIsMobile() bool {
	if o == nil || IsNil(o.IsMobile) {
		var ret bool
		return ret
	}
	return *o.IsMobile
}

// GetIsMobileOk returns a tuple with the IsMobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativePreviewConfiguration) GetIsMobileOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMobile) {
		return nil, false
	}
	return o.IsMobile, true
}

// HasIsMobile returns a boolean if a field has been set.
func (o *CreativePreviewConfiguration) HasIsMobile() bool {
	if o != nil && !IsNil(o.IsMobile) {
		return true
	}

	return false
}

// SetIsMobile gets a reference to the given bool and assigns it to the IsMobile field.
func (o *CreativePreviewConfiguration) SetIsMobile(v bool) {
	o.IsMobile = &v
}

// GetIsOnAmazon returns the IsOnAmazon field value if set, zero value otherwise.
func (o *CreativePreviewConfiguration) GetIsOnAmazon() bool {
	if o == nil || IsNil(o.IsOnAmazon) {
		var ret bool
		return ret
	}
	return *o.IsOnAmazon
}

// GetIsOnAmazonOk returns a tuple with the IsOnAmazon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreativePreviewConfiguration) GetIsOnAmazonOk() (*bool, bool) {
	if o == nil || IsNil(o.IsOnAmazon) {
		return nil, false
	}
	return o.IsOnAmazon, true
}

// HasIsOnAmazon returns a boolean if a field has been set.
func (o *CreativePreviewConfiguration) HasIsOnAmazon() bool {
	if o != nil && !IsNil(o.IsOnAmazon) {
		return true
	}

	return false
}

// SetIsOnAmazon gets a reference to the given bool and assigns it to the IsOnAmazon field.
func (o *CreativePreviewConfiguration) SetIsOnAmazon(v bool) {
	o.IsOnAmazon = &v
}

func (o CreativePreviewConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Products) {
		toSerialize["products"] = o.Products
	}
	if !IsNil(o.LandingPageURL) {
		toSerialize["landingPageURL"] = o.LandingPageURL
	}
	if !IsNil(o.LandingPageType) {
		toSerialize["landingPageType"] = o.LandingPageType
	}
	if !IsNil(o.AdName) {
		toSerialize["adName"] = o.AdName
	}
	if !IsNil(o.IsMobile) {
		toSerialize["isMobile"] = o.IsMobile
	}
	if !IsNil(o.IsOnAmazon) {
		toSerialize["isOnAmazon"] = o.IsOnAmazon
	}
	return toSerialize, nil
}

type NullableCreativePreviewConfiguration struct {
	value *CreativePreviewConfiguration
	isSet bool
}

func (v NullableCreativePreviewConfiguration) Get() *CreativePreviewConfiguration {
	return v.value
}

func (v *NullableCreativePreviewConfiguration) Set(val *CreativePreviewConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableCreativePreviewConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableCreativePreviewConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreativePreviewConfiguration(val *CreativePreviewConfiguration) *NullableCreativePreviewConfiguration {
	return &NullableCreativePreviewConfiguration{value: val, isSet: true}
}

func (v NullableCreativePreviewConfiguration) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreativePreviewConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
