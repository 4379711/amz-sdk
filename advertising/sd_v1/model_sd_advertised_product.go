package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDAdvertisedProduct type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDAdvertisedProduct{}

// SDAdvertisedProduct Product that advertisers want to advertise. Recommendations will be made for specified products. SDAdvertisedProduct can only contain either asins or landing pages. If landingPageUrl is used, there can only be one item for each request.
type SDAdvertisedProduct struct {
	// Amazon Standard Identification Number
	Asin            *string            `json:"asin,omitempty" validate:"regexp=[a-zA-Z0-9]{10}"`
	LandingPageType *SDLandingPageType `json:"landingPageType,omitempty"`
	// The URL where customers will land after clicking on its link. Must be provided if landingPageType field is set. This field is not supported when using asin field. ||Specifications| |------------------|------------------| |LandingPageType| Description| |OFF_AMAZON_LINK| The url should be in the format of https://www.****.com.|
	LandingPageURL *string `json:"landingPageURL,omitempty"`
}

// NewSDAdvertisedProduct instantiates a new SDAdvertisedProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDAdvertisedProduct() *SDAdvertisedProduct {
	this := SDAdvertisedProduct{}
	return &this
}

// NewSDAdvertisedProductWithDefaults instantiates a new SDAdvertisedProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDAdvertisedProductWithDefaults() *SDAdvertisedProduct {
	this := SDAdvertisedProduct{}
	return &this
}

// GetAsin returns the Asin field value if set, zero value otherwise.
func (o *SDAdvertisedProduct) GetAsin() string {
	if o == nil || IsNil(o.Asin) {
		var ret string
		return ret
	}
	return *o.Asin
}

// GetAsinOk returns a tuple with the Asin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDAdvertisedProduct) GetAsinOk() (*string, bool) {
	if o == nil || IsNil(o.Asin) {
		return nil, false
	}
	return o.Asin, true
}

// HasAsin returns a boolean if a field has been set.
func (o *SDAdvertisedProduct) HasAsin() bool {
	if o != nil && !IsNil(o.Asin) {
		return true
	}

	return false
}

// SetAsin gets a reference to the given string and assigns it to the Asin field.
func (o *SDAdvertisedProduct) SetAsin(v string) {
	o.Asin = &v
}

// GetLandingPageType returns the LandingPageType field value if set, zero value otherwise.
func (o *SDAdvertisedProduct) GetLandingPageType() SDLandingPageType {
	if o == nil || IsNil(o.LandingPageType) {
		var ret SDLandingPageType
		return ret
	}
	return *o.LandingPageType
}

// GetLandingPageTypeOk returns a tuple with the LandingPageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDAdvertisedProduct) GetLandingPageTypeOk() (*SDLandingPageType, bool) {
	if o == nil || IsNil(o.LandingPageType) {
		return nil, false
	}
	return o.LandingPageType, true
}

// HasLandingPageType returns a boolean if a field has been set.
func (o *SDAdvertisedProduct) HasLandingPageType() bool {
	if o != nil && !IsNil(o.LandingPageType) {
		return true
	}

	return false
}

// SetLandingPageType gets a reference to the given SDLandingPageType and assigns it to the LandingPageType field.
func (o *SDAdvertisedProduct) SetLandingPageType(v SDLandingPageType) {
	o.LandingPageType = &v
}

// GetLandingPageURL returns the LandingPageURL field value if set, zero value otherwise.
func (o *SDAdvertisedProduct) GetLandingPageURL() string {
	if o == nil || IsNil(o.LandingPageURL) {
		var ret string
		return ret
	}
	return *o.LandingPageURL
}

// GetLandingPageURLOk returns a tuple with the LandingPageURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDAdvertisedProduct) GetLandingPageURLOk() (*string, bool) {
	if o == nil || IsNil(o.LandingPageURL) {
		return nil, false
	}
	return o.LandingPageURL, true
}

// HasLandingPageURL returns a boolean if a field has been set.
func (o *SDAdvertisedProduct) HasLandingPageURL() bool {
	if o != nil && !IsNil(o.LandingPageURL) {
		return true
	}

	return false
}

// SetLandingPageURL gets a reference to the given string and assigns it to the LandingPageURL field.
func (o *SDAdvertisedProduct) SetLandingPageURL(v string) {
	o.LandingPageURL = &v
}

func (o SDAdvertisedProduct) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Asin) {
		toSerialize["asin"] = o.Asin
	}
	if !IsNil(o.LandingPageType) {
		toSerialize["landingPageType"] = o.LandingPageType
	}
	if !IsNil(o.LandingPageURL) {
		toSerialize["landingPageURL"] = o.LandingPageURL
	}
	return toSerialize, nil
}

type NullableSDAdvertisedProduct struct {
	value *SDAdvertisedProduct
	isSet bool
}

func (v NullableSDAdvertisedProduct) Get() *SDAdvertisedProduct {
	return v.value
}

func (v *NullableSDAdvertisedProduct) Set(val *SDAdvertisedProduct) {
	v.value = val
	v.isSet = true
}

func (v NullableSDAdvertisedProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableSDAdvertisedProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDAdvertisedProduct(val *SDAdvertisedProduct) *NullableSDAdvertisedProduct {
	return &NullableSDAdvertisedProduct{value: val, isSet: true}
}

func (v NullableSDAdvertisedProduct) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDAdvertisedProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
