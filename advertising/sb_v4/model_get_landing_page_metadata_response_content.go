package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the GetLandingPageMetadataResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetLandingPageMetadataResponseContent{}

// GetLandingPageMetadataResponseContent struct for GetLandingPageMetadataResponseContent
type GetLandingPageMetadataResponseContent struct {
	// The type of landing page, such as store page, product list (simple landing page), custom url. | Page Type    | |--------------| | PRODUCT_LIST | | STORE        | | CUSTOM_URL   | | DETAIL_PAGE  |
	PageType string `json:"pageType"`
	// A canonical URL is the URL that represents the best version of landing page URL from a group of duplicate landing page URLs. For example, if you have two landing page URLs for the same page (such as amazon.it/HSA/pages/default?pageId=B59A592C-8A12-4684-A37E-2416FD594D87 and amazon.it/stores/page/B59A592C-8A12-4684-A37E-2416FD594D87), we chooses one as canonical. In this case, canonical url is amazon.it/stores/page/B59A592C-8A12-4684-A37E-2416FD594D87
	CanonicalUrl string `json:"canonicalUrl"`
	// A human-readable description of the unSupportedReasonCode field.
	UnSupportedReason *string `json:"unSupportedReason,omitempty"`
	// This field determines whether the landing page is supported for the ad product.
	IsSupported *bool `json:"isSupported,omitempty"`
	// Enumerated code for why landing page is unsupported. | Reason Code                 | | SB_DETAIL_PAGE_UNSUPPORTED  | | SB_GATEWAY_PAGE_UNSUPPORTED | | SB_SEARCH_PAGE_UNSUPPORTED  | | SB_BROWSE_PAGE_UNSUPPORTED  | | SB_OTHER_PAGE_UNSUPPORTED   |
	UnSupportedReasonCode *string `json:"unSupportedReasonCode,omitempty"`
}

type _GetLandingPageMetadataResponseContent GetLandingPageMetadataResponseContent

// NewGetLandingPageMetadataResponseContent instantiates a new GetLandingPageMetadataResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLandingPageMetadataResponseContent(pageType string, canonicalUrl string) *GetLandingPageMetadataResponseContent {
	this := GetLandingPageMetadataResponseContent{}
	this.PageType = pageType
	this.CanonicalUrl = canonicalUrl
	return &this
}

// NewGetLandingPageMetadataResponseContentWithDefaults instantiates a new GetLandingPageMetadataResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLandingPageMetadataResponseContentWithDefaults() *GetLandingPageMetadataResponseContent {
	this := GetLandingPageMetadataResponseContent{}
	return &this
}

// GetPageType returns the PageType field value
func (o *GetLandingPageMetadataResponseContent) GetPageType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PageType
}

// GetPageTypeOk returns a tuple with the PageType field value
// and a boolean to check if the value has been set.
func (o *GetLandingPageMetadataResponseContent) GetPageTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageType, true
}

// SetPageType sets field value
func (o *GetLandingPageMetadataResponseContent) SetPageType(v string) {
	o.PageType = v
}

// GetCanonicalUrl returns the CanonicalUrl field value
func (o *GetLandingPageMetadataResponseContent) GetCanonicalUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CanonicalUrl
}

// GetCanonicalUrlOk returns a tuple with the CanonicalUrl field value
// and a boolean to check if the value has been set.
func (o *GetLandingPageMetadataResponseContent) GetCanonicalUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanonicalUrl, true
}

// SetCanonicalUrl sets field value
func (o *GetLandingPageMetadataResponseContent) SetCanonicalUrl(v string) {
	o.CanonicalUrl = v
}

// GetUnSupportedReason returns the UnSupportedReason field value if set, zero value otherwise.
func (o *GetLandingPageMetadataResponseContent) GetUnSupportedReason() string {
	if o == nil || IsNil(o.UnSupportedReason) {
		var ret string
		return ret
	}
	return *o.UnSupportedReason
}

// GetUnSupportedReasonOk returns a tuple with the UnSupportedReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLandingPageMetadataResponseContent) GetUnSupportedReasonOk() (*string, bool) {
	if o == nil || IsNil(o.UnSupportedReason) {
		return nil, false
	}
	return o.UnSupportedReason, true
}

// HasUnSupportedReason returns a boolean if a field has been set.
func (o *GetLandingPageMetadataResponseContent) HasUnSupportedReason() bool {
	if o != nil && !IsNil(o.UnSupportedReason) {
		return true
	}

	return false
}

// SetUnSupportedReason gets a reference to the given string and assigns it to the UnSupportedReason field.
func (o *GetLandingPageMetadataResponseContent) SetUnSupportedReason(v string) {
	o.UnSupportedReason = &v
}

// GetIsSupported returns the IsSupported field value if set, zero value otherwise.
func (o *GetLandingPageMetadataResponseContent) GetIsSupported() bool {
	if o == nil || IsNil(o.IsSupported) {
		var ret bool
		return ret
	}
	return *o.IsSupported
}

// GetIsSupportedOk returns a tuple with the IsSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLandingPageMetadataResponseContent) GetIsSupportedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSupported) {
		return nil, false
	}
	return o.IsSupported, true
}

// HasIsSupported returns a boolean if a field has been set.
func (o *GetLandingPageMetadataResponseContent) HasIsSupported() bool {
	if o != nil && !IsNil(o.IsSupported) {
		return true
	}

	return false
}

// SetIsSupported gets a reference to the given bool and assigns it to the IsSupported field.
func (o *GetLandingPageMetadataResponseContent) SetIsSupported(v bool) {
	o.IsSupported = &v
}

// GetUnSupportedReasonCode returns the UnSupportedReasonCode field value if set, zero value otherwise.
func (o *GetLandingPageMetadataResponseContent) GetUnSupportedReasonCode() string {
	if o == nil || IsNil(o.UnSupportedReasonCode) {
		var ret string
		return ret
	}
	return *o.UnSupportedReasonCode
}

// GetUnSupportedReasonCodeOk returns a tuple with the UnSupportedReasonCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLandingPageMetadataResponseContent) GetUnSupportedReasonCodeOk() (*string, bool) {
	if o == nil || IsNil(o.UnSupportedReasonCode) {
		return nil, false
	}
	return o.UnSupportedReasonCode, true
}

// HasUnSupportedReasonCode returns a boolean if a field has been set.
func (o *GetLandingPageMetadataResponseContent) HasUnSupportedReasonCode() bool {
	if o != nil && !IsNil(o.UnSupportedReasonCode) {
		return true
	}

	return false
}

// SetUnSupportedReasonCode gets a reference to the given string and assigns it to the UnSupportedReasonCode field.
func (o *GetLandingPageMetadataResponseContent) SetUnSupportedReasonCode(v string) {
	o.UnSupportedReasonCode = &v
}

func (o GetLandingPageMetadataResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pageType"] = o.PageType
	toSerialize["canonicalUrl"] = o.CanonicalUrl
	if !IsNil(o.UnSupportedReason) {
		toSerialize["unSupportedReason"] = o.UnSupportedReason
	}
	if !IsNil(o.IsSupported) {
		toSerialize["isSupported"] = o.IsSupported
	}
	if !IsNil(o.UnSupportedReasonCode) {
		toSerialize["unSupportedReasonCode"] = o.UnSupportedReasonCode
	}
	return toSerialize, nil
}

type NullableGetLandingPageMetadataResponseContent struct {
	value *GetLandingPageMetadataResponseContent
	isSet bool
}

func (v NullableGetLandingPageMetadataResponseContent) Get() *GetLandingPageMetadataResponseContent {
	return v.value
}

func (v *NullableGetLandingPageMetadataResponseContent) Set(val *GetLandingPageMetadataResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLandingPageMetadataResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLandingPageMetadataResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLandingPageMetadataResponseContent(val *GetLandingPageMetadataResponseContent) *NullableGetLandingPageMetadataResponseContent {
	return &NullableGetLandingPageMetadataResponseContent{value: val, isSet: true}
}

func (v NullableGetLandingPageMetadataResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetLandingPageMetadataResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
