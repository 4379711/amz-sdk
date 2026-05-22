package product_metadata

import (
	"github.com/bytedance/sonic"
)

// checks if the ProductInfoModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductInfoModel{}

// ProductInfoModel struct for ProductInfoModel
type ProductInfoModel struct {
	Deal              *Deal              `json:"deal,omitempty"`
	CatalogType       *CatalogType       `json:"catalogType,omitempty"`
	TvPreviewMetadata *TvPreviewMetadata `json:"tvPreviewMetadata,omitempty"`
	// Stock availability:   * IN_STOCK - The item is in stock.   * IN_STOCK_SCARCE - The item is in stock, but stock levels are limited.   * OUT_OF_STOCK - The item is currently out of stock.   * PREORDER - The item is not yet available, but can be pre-ordered.   * LEADTIME - The item is only available after some amount of lead time.   * AVAILABLE_DATE - The item is not available, but will be available on a future date.
	Availability *string `json:"availability,omitempty"`
	// Product title of the item
	Title              *string                             `json:"title,omitempty"`
	GlobalStoreSetting *ProductInfoModelGlobalStoreSetting `json:"globalStoreSetting,omitempty"`
	AdEligibility      *LegacyEligibility                  `json:"adEligibility,omitempty"`
	// Parent ASIN of the book ASIN
	TitleAuthorityParentAsin *string `json:"titleAuthorityParentAsin,omitempty"`
	// List of ASIN variations of the current item
	Variations   []string   `json:"variations,omitempty"`
	Price        *ListPrice `json:"price,omitempty"`
	TvPreviewURL *string    `json:"tvPreviewURL,omitempty"`
	// Url to the product image
	ImageUrl         *string  `json:"imageUrl,omitempty"`
	EligibilityCodes []string `json:"eligibilityCodes,omitempty"`
	// Url to the product detail page in retail
	Href *string `json:"href,omitempty"`
	// List of ASIN variations of the book item
	TitleAuthority []string `json:"titleAuthority,omitempty"`
	// List of reasons that made this item ineligible to be advertised
	IneligibilityReasons []string  `json:"ineligibilityReasons,omitempty"`
	EligibilityReasons   []string  `json:"eligibilityReasons,omitempty"`
	TvIconURL            *string   `json:"tvIconURL,omitempty"`
	Merchant             *Merchant `json:"merchant,omitempty"`
	// List of ineligible status identifier
	IneligibilityCodes []string `json:"ineligibilityCodes,omitempty"`
	// Eligibility status for advertising:   * ELIGIBLE - Eligible for advertising   * INELIGIBLE - Ineligible for advertising
	EligibilityStatus *string            `json:"eligibilityStatus,omitempty"`
	AsinEligibility   *LegacyEligibility `json:"asinEligibility,omitempty"`
	CustomerReview    *CustomerReview    `json:"customerReview,omitempty"`
	// List of productSet
	ProductSets []ProductSet `json:"productSets,omitempty"`
	// ASIN of the item
	Asin      *string    `json:"asin,omitempty"`
	ListPrice *ListPrice `json:"listPrice,omitempty"`
}

// NewProductInfoModel instantiates a new ProductInfoModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductInfoModel() *ProductInfoModel {
	this := ProductInfoModel{}
	return &this
}

// NewProductInfoModelWithDefaults instantiates a new ProductInfoModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductInfoModelWithDefaults() *ProductInfoModel {
	this := ProductInfoModel{}
	return &this
}

// GetDeal returns the Deal field value if set, zero value otherwise.
func (o *ProductInfoModel) GetDeal() Deal {
	if o == nil || IsNil(o.Deal) {
		var ret Deal
		return ret
	}
	return *o.Deal
}

// GetDealOk returns a tuple with the Deal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductInfoModel) GetDealOk() (*Deal, bool) {
	if o == nil || IsNil(o.Deal) {
		return nil, false
	}
	return o.Deal, true
}

// HasDeal returns a boolean if a field has been set.
func (o *ProductInfoModel) HasDeal() bool {
	if o != nil && !IsNil(o.Deal) {
		return true
	}

	return false
}

// SetDeal gets a reference to the given Deal and assigns it to the Deal field.
func (o *ProductInfoModel) SetDeal(v Deal) {
	o.Deal = &v
}

// GetCatalogType returns the CatalogType field value if set, zero value otherwise.
func (o *ProductInfoModel) GetCatalogType() CatalogType {
	if o == nil || IsNil(o.CatalogType) {
		var ret CatalogType
		return ret
	}
	return *o.CatalogType
}

// GetCatalogTypeOk returns a tuple with the CatalogType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductInfoModel) GetCatalogTypeOk() (*CatalogType, bool) {
	if o == nil || IsNil(o.CatalogType) {
		return nil, false
	}
	return o.CatalogType, true
}

// HasCatalogType returns a boolean if a field has been set.
func (o *ProductInfoModel) HasCatalogType() bool {
	if o != nil && !IsNil(o.CatalogType) {
		return true
	}

	return false
}

// SetCatalogType gets a reference to the given CatalogType and assigns it to the CatalogType field.
func (o *ProductInfoModel) SetCatalogType(v CatalogType) {
	o.CatalogType = &v
}

// GetTvPreviewMetadata returns the TvPreviewMetadata field value if set, zero value otherwise.
func (o *ProductInfoModel) GetTvPreviewMetadata() TvPreviewMetadata {
	if o == nil || IsNil(o.TvPreviewMetadata) {
		var ret TvPreviewMetadata
		return ret
	}
	return *o.TvPreviewMetadata
}

// GetTvPreviewMetadataOk returns a tuple with the TvPreviewMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductInfoModel) GetTvPreviewMetadataOk() (*TvPreviewMetadata, bool) {
	if o == nil || IsNil(o.TvPreviewMetadata) {
		return nil, false
	}
	return o.TvPreviewMetadata, true
}

// HasTvPreviewMetadata returns a boolean if a field has been set.
func (o *ProductInfoModel) HasTvPreviewMetadata() bool {
	if o != nil && !IsNil(o.TvPreviewMetadata) {
		return true
	}

	return false
}

// SetTvPreviewMetadata gets a reference to the given TvPreviewMetadata and assigns it to the TvPreviewMetadata field.
func (o *ProductInfoModel) SetTvPreviewMetadata(v TvPreviewMetadata) {
	o.TvPreviewMetadata = &v
}

// GetAvailability returns the Availability field value if set, zero value otherwise.
func (o *ProductInfoModel) GetAvailability() string {
	if o == nil || IsNil(o.Availability) {
		var ret string
		return ret
	}
	return *o.Availability
}

// GetAvailabilityOk returns a tuple with the Availability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductInfoModel) GetAvailabilityOk() (*string, bool) {
	if o == nil || IsNil(o.Availability) {
		return nil, false
	}
	return o.Availability, true
}

// HasAvailability returns a boolean if a field has been set.
func (o *ProductInfoModel) HasAvailability() bool {
	if o != nil && !IsNil(o.Availability) {
		return true
	}

	return false
}

// SetAvailability gets a reference to the given string and assigns it to the Availability field.
func (o *ProductInfoModel) SetAvailability(v string) {
	o.Availability = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ProductInfoModel) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductInfoModel) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ProductInfoModel) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ProductInfoModel) SetTitle(v string) {
	o.Title = &v
}

// GetGlobalStoreSetting returns the GlobalStoreSetting field value if set, zero value otherwise.
func (o *ProductInfoModel) GetGlobalStoreSetting() ProductInfoModelGlobalStoreSetting {
	if o == nil || IsNil(o.GlobalStoreSetting) {
		var ret ProductInfoModelGlobalStoreSetting
		return ret
	}
	return *o.GlobalStoreSetting
}

// GetGlobalStoreSettingOk returns a tuple with the GlobalStoreSetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductInfoModel) GetGlobalStoreSettingOk() (*ProductInfoModelGlobalStoreSetting, bool) {
	if o == nil || IsNil(o.GlobalStoreSetting) {
		return nil, false
	}
	return o.GlobalStoreSetting, true
}

// HasGlobalStoreSetting returns a boolean if a field has been set.
func (o *ProductInfoModel) HasGlobalStoreSetting() bool {
	if o != nil && !IsNil(o.GlobalStoreSetting) {
		return true
	}

	return false
}

// SetGlobalStoreSetting gets a reference to the given ProductInfoModelGlobalStoreSetting and assigns it to the GlobalStoreSetting field.
func (o *ProductInfoModel) SetGlobalStoreSetting(v ProductInfoModelGlobalStoreSetting) {
	o.GlobalStoreSetting = &v
}

// GetAdEligibility returns the AdEligibility field value if set, zero value otherwise.
func (o *ProductInfoModel) GetAdEligibility() LegacyEligibility {
	if o == nil || IsNil(o.AdEligibility) {
		var ret LegacyEligibility
		return ret
	}
	return *o.AdEligibility
}

// GetAdEligibilityOk returns a tuple with the AdEligibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductInfoModel) GetAdEligibilityOk() (*LegacyEligibility, bool) {
	if o == nil || IsNil(o.AdEligibility) {
		return nil, false
	}
	return o.AdEligibility, true
}

// HasAdEligibility returns a boolean if a field has been set.
func (o *ProductInfoModel) HasAdEligibility() bool {
	if o != nil && !IsNil(o.AdEligibility) {
		return true
	}

	return false
}

// SetAdEligibility gets a reference to the given LegacyEligibility and assigns it to the AdEligibility field.
func (o *ProductInfoModel) SetAdEligibility(v LegacyEligibility) {
	o.AdEligibility = &v
}

// GetTitleAuthorityParentAsin returns the TitleAuthorityParentAsin field value if set, zero value otherwise.
func (o *ProductInfoModel) GetTitleAuthorityParentAsin() string {
	if o == nil || IsNil(o.TitleAuthorityParentAsin) {
		var ret string
		return ret
	}
	return *o.TitleAuthorityParentAsin
}

// GetTitleAuthorityParentAsinOk returns a tuple with the TitleAuthorityParentAsin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductInfoModel) GetTitleAuthorityParentAsinOk() (*string, bool) {
	if o == nil || IsNil(o.TitleAuthorityParentAsin) {
		return nil, false
	}
	return o.TitleAuthorityParentAsin, true
}

// HasTitleAuthorityParentAsin returns a boolean if a field has been set.
func (o *ProductInfoModel) HasTitleAuthorityParentAsin() bool {
	if o != nil && !IsNil(o.TitleAuthorityParentAsin) {
		return true
	}

	return false
}

// SetTitleAuthorityParentAsin gets a reference to the given string and assigns it to the TitleAuthorityParentAsin field.
func (o *ProductInfoModel) SetTitleAuthorityParentAsin(v string) {
	o.TitleAuthorityParentAsin = &v
}

// GetVariations returns the Variations field value if set, zero value otherwise.
func (o *ProductInfoModel) GetVariations() []string {
	if o == nil || IsNil(o.Variations) {
		var ret []string
		return ret
	}
	return o.Variations
}

// GetVariationsOk returns a tuple with the Variations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductInfoModel) GetVariationsOk() ([]string, bool) {
	if o == nil || IsNil(o.Variations) {
		return nil, false
	}
	return o.Variations, true
}

// HasVariations returns a boolean if a field has been set.
func (o *ProductInfoModel) HasVariations() bool {
	if o != nil && !IsNil(o.Variations) {
		return true
	}

	return false
}

// SetVariations gets a reference to the given []string and assigns it to the Variations field.
func (o *ProductInfoModel) SetVariations(v []string) {
	o.Variations = v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *ProductInfoModel) GetPrice() ListPrice {
	if o == nil || IsNil(o.Price) {
		var ret ListPrice
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductInfoModel) GetPriceOk() (*ListPrice, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *ProductInfoModel) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given ListPrice and assigns it to the Price field.
func (o *ProductInfoModel) SetPrice(v ListPrice) {
	o.Price = &v
}

// GetTvPreviewURL returns the TvPreviewURL field value if set, zero value otherwise.
func (o *ProductInfoModel) GetTvPreviewURL() string {
	if o == nil || IsNil(o.TvPreviewURL) {
		var ret string
		return ret
	}
	return *o.TvPreviewURL
}

// GetTvPreviewURLOk returns a tuple with the TvPreviewURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductInfoModel) GetTvPreviewURLOk() (*string, bool) {
	if o == nil || IsNil(o.TvPreviewURL) {
		return nil, false
	}
	return o.TvPreviewURL, true
}

// HasTvPreviewURL returns a boolean if a field has been set.
func (o *ProductInfoModel) HasTvPreviewURL() bool {
	if o != nil && !IsNil(o.TvPreviewURL) {
		return true
	}

	return false
}

// SetTvPreviewURL gets a reference to the given string and assigns it to the TvPreviewURL field.
func (o *ProductInfoModel) SetTvPreviewURL(v string) {
	o.TvPreviewURL = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *ProductInfoModel) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductInfoModel) GetImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *ProductInfoModel) HasImageUrl() bool {
	if o != nil && !IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *ProductInfoModel) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetEligibilityCodes returns the EligibilityCodes field value if set, zero value otherwise.
func (o *ProductInfoModel) GetEligibilityCodes() []string {
	if o == nil || IsNil(o.EligibilityCodes) {
		var ret []string
		return ret
	}
	return o.EligibilityCodes
}

// GetEligibilityCodesOk returns a tuple with the EligibilityCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductInfoModel) GetEligibilityCodesOk() ([]string, bool) {
	if o == nil || IsNil(o.EligibilityCodes) {
		return nil, false
	}
	return o.EligibilityCodes, true
}

// HasEligibilityCodes returns a boolean if a field has been set.
func (o *ProductInfoModel) HasEligibilityCodes() bool {
	if o != nil && !IsNil(o.EligibilityCodes) {
		return true
	}

	return false
}

// SetEligibilityCodes gets a reference to the given []string and assigns it to the EligibilityCodes field.
func (o *ProductInfoModel) SetEligibilityCodes(v []string) {
	o.EligibilityCodes = v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *ProductInfoModel) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductInfoModel) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *ProductInfoModel) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *ProductInfoModel) SetHref(v string) {
	o.Href = &v
}

// GetTitleAuthority returns the TitleAuthority field value if set, zero value otherwise.
func (o *ProductInfoModel) GetTitleAuthority() []string {
	if o == nil || IsNil(o.TitleAuthority) {
		var ret []string
		return ret
	}
	return o.TitleAuthority
}

// GetTitleAuthorityOk returns a tuple with the TitleAuthority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductInfoModel) GetTitleAuthorityOk() ([]string, bool) {
	if o == nil || IsNil(o.TitleAuthority) {
		return nil, false
	}
	return o.TitleAuthority, true
}

// HasTitleAuthority returns a boolean if a field has been set.
func (o *ProductInfoModel) HasTitleAuthority() bool {
	if o != nil && !IsNil(o.TitleAuthority) {
		return true
	}

	return false
}

// SetTitleAuthority gets a reference to the given []string and assigns it to the TitleAuthority field.
func (o *ProductInfoModel) SetTitleAuthority(v []string) {
	o.TitleAuthority = v
}

// GetIneligibilityReasons returns the IneligibilityReasons field value if set, zero value otherwise.
func (o *ProductInfoModel) GetIneligibilityReasons() []string {
	if o == nil || IsNil(o.IneligibilityReasons) {
		var ret []string
		return ret
	}
	return o.IneligibilityReasons
}

// GetIneligibilityReasonsOk returns a tuple with the IneligibilityReasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductInfoModel) GetIneligibilityReasonsOk() ([]string, bool) {
	if o == nil || IsNil(o.IneligibilityReasons) {
		return nil, false
	}
	return o.IneligibilityReasons, true
}

// HasIneligibilityReasons returns a boolean if a field has been set.
func (o *ProductInfoModel) HasIneligibilityReasons() bool {
	if o != nil && !IsNil(o.IneligibilityReasons) {
		return true
	}

	return false
}

// SetIneligibilityReasons gets a reference to the given []string and assigns it to the IneligibilityReasons field.
func (o *ProductInfoModel) SetIneligibilityReasons(v []string) {
	o.IneligibilityReasons = v
}

// GetEligibilityReasons returns the EligibilityReasons field value if set, zero value otherwise.
func (o *ProductInfoModel) GetEligibilityReasons() []string {
	if o == nil || IsNil(o.EligibilityReasons) {
		var ret []string
		return ret
	}
	return o.EligibilityReasons
}

// GetEligibilityReasonsOk returns a tuple with the EligibilityReasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductInfoModel) GetEligibilityReasonsOk() ([]string, bool) {
	if o == nil || IsNil(o.EligibilityReasons) {
		return nil, false
	}
	return o.EligibilityReasons, true
}

// HasEligibilityReasons returns a boolean if a field has been set.
func (o *ProductInfoModel) HasEligibilityReasons() bool {
	if o != nil && !IsNil(o.EligibilityReasons) {
		return true
	}

	return false
}

// SetEligibilityReasons gets a reference to the given []string and assigns it to the EligibilityReasons field.
func (o *ProductInfoModel) SetEligibilityReasons(v []string) {
	o.EligibilityReasons = v
}

// GetTvIconURL returns the TvIconURL field value if set, zero value otherwise.
func (o *ProductInfoModel) GetTvIconURL() string {
	if o == nil || IsNil(o.TvIconURL) {
		var ret string
		return ret
	}
	return *o.TvIconURL
}

// GetTvIconURLOk returns a tuple with the TvIconURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductInfoModel) GetTvIconURLOk() (*string, bool) {
	if o == nil || IsNil(o.TvIconURL) {
		return nil, false
	}
	return o.TvIconURL, true
}

// HasTvIconURL returns a boolean if a field has been set.
func (o *ProductInfoModel) HasTvIconURL() bool {
	if o != nil && !IsNil(o.TvIconURL) {
		return true
	}

	return false
}

// SetTvIconURL gets a reference to the given string and assigns it to the TvIconURL field.
func (o *ProductInfoModel) SetTvIconURL(v string) {
	o.TvIconURL = &v
}

// GetMerchant returns the Merchant field value if set, zero value otherwise.
func (o *ProductInfoModel) GetMerchant() Merchant {
	if o == nil || IsNil(o.Merchant) {
		var ret Merchant
		return ret
	}
	return *o.Merchant
}

// GetMerchantOk returns a tuple with the Merchant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductInfoModel) GetMerchantOk() (*Merchant, bool) {
	if o == nil || IsNil(o.Merchant) {
		return nil, false
	}
	return o.Merchant, true
}

// HasMerchant returns a boolean if a field has been set.
func (o *ProductInfoModel) HasMerchant() bool {
	if o != nil && !IsNil(o.Merchant) {
		return true
	}

	return false
}

// SetMerchant gets a reference to the given Merchant and assigns it to the Merchant field.
func (o *ProductInfoModel) SetMerchant(v Merchant) {
	o.Merchant = &v
}

// GetIneligibilityCodes returns the IneligibilityCodes field value if set, zero value otherwise.
func (o *ProductInfoModel) GetIneligibilityCodes() []string {
	if o == nil || IsNil(o.IneligibilityCodes) {
		var ret []string
		return ret
	}
	return o.IneligibilityCodes
}

// GetIneligibilityCodesOk returns a tuple with the IneligibilityCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductInfoModel) GetIneligibilityCodesOk() ([]string, bool) {
	if o == nil || IsNil(o.IneligibilityCodes) {
		return nil, false
	}
	return o.IneligibilityCodes, true
}

// HasIneligibilityCodes returns a boolean if a field has been set.
func (o *ProductInfoModel) HasIneligibilityCodes() bool {
	if o != nil && !IsNil(o.IneligibilityCodes) {
		return true
	}

	return false
}

// SetIneligibilityCodes gets a reference to the given []string and assigns it to the IneligibilityCodes field.
func (o *ProductInfoModel) SetIneligibilityCodes(v []string) {
	o.IneligibilityCodes = v
}

// GetEligibilityStatus returns the EligibilityStatus field value if set, zero value otherwise.
func (o *ProductInfoModel) GetEligibilityStatus() string {
	if o == nil || IsNil(o.EligibilityStatus) {
		var ret string
		return ret
	}
	return *o.EligibilityStatus
}

// GetEligibilityStatusOk returns a tuple with the EligibilityStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductInfoModel) GetEligibilityStatusOk() (*string, bool) {
	if o == nil || IsNil(o.EligibilityStatus) {
		return nil, false
	}
	return o.EligibilityStatus, true
}

// HasEligibilityStatus returns a boolean if a field has been set.
func (o *ProductInfoModel) HasEligibilityStatus() bool {
	if o != nil && !IsNil(o.EligibilityStatus) {
		return true
	}

	return false
}

// SetEligibilityStatus gets a reference to the given string and assigns it to the EligibilityStatus field.
func (o *ProductInfoModel) SetEligibilityStatus(v string) {
	o.EligibilityStatus = &v
}

// GetAsinEligibility returns the AsinEligibility field value if set, zero value otherwise.
func (o *ProductInfoModel) GetAsinEligibility() LegacyEligibility {
	if o == nil || IsNil(o.AsinEligibility) {
		var ret LegacyEligibility
		return ret
	}
	return *o.AsinEligibility
}

// GetAsinEligibilityOk returns a tuple with the AsinEligibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductInfoModel) GetAsinEligibilityOk() (*LegacyEligibility, bool) {
	if o == nil || IsNil(o.AsinEligibility) {
		return nil, false
	}
	return o.AsinEligibility, true
}

// HasAsinEligibility returns a boolean if a field has been set.
func (o *ProductInfoModel) HasAsinEligibility() bool {
	if o != nil && !IsNil(o.AsinEligibility) {
		return true
	}

	return false
}

// SetAsinEligibility gets a reference to the given LegacyEligibility and assigns it to the AsinEligibility field.
func (o *ProductInfoModel) SetAsinEligibility(v LegacyEligibility) {
	o.AsinEligibility = &v
}

// GetCustomerReview returns the CustomerReview field value if set, zero value otherwise.
func (o *ProductInfoModel) GetCustomerReview() CustomerReview {
	if o == nil || IsNil(o.CustomerReview) {
		var ret CustomerReview
		return ret
	}
	return *o.CustomerReview
}

// GetCustomerReviewOk returns a tuple with the CustomerReview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductInfoModel) GetCustomerReviewOk() (*CustomerReview, bool) {
	if o == nil || IsNil(o.CustomerReview) {
		return nil, false
	}
	return o.CustomerReview, true
}

// HasCustomerReview returns a boolean if a field has been set.
func (o *ProductInfoModel) HasCustomerReview() bool {
	if o != nil && !IsNil(o.CustomerReview) {
		return true
	}

	return false
}

// SetCustomerReview gets a reference to the given CustomerReview and assigns it to the CustomerReview field.
func (o *ProductInfoModel) SetCustomerReview(v CustomerReview) {
	o.CustomerReview = &v
}

// GetProductSets returns the ProductSets field value if set, zero value otherwise.
func (o *ProductInfoModel) GetProductSets() []ProductSet {
	if o == nil || IsNil(o.ProductSets) {
		var ret []ProductSet
		return ret
	}
	return o.ProductSets
}

// GetProductSetsOk returns a tuple with the ProductSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductInfoModel) GetProductSetsOk() ([]ProductSet, bool) {
	if o == nil || IsNil(o.ProductSets) {
		return nil, false
	}
	return o.ProductSets, true
}

// HasProductSets returns a boolean if a field has been set.
func (o *ProductInfoModel) HasProductSets() bool {
	if o != nil && !IsNil(o.ProductSets) {
		return true
	}

	return false
}

// SetProductSets gets a reference to the given []ProductSet and assigns it to the ProductSets field.
func (o *ProductInfoModel) SetProductSets(v []ProductSet) {
	o.ProductSets = v
}

// GetAsin returns the Asin field value if set, zero value otherwise.
func (o *ProductInfoModel) GetAsin() string {
	if o == nil || IsNil(o.Asin) {
		var ret string
		return ret
	}
	return *o.Asin
}

// GetAsinOk returns a tuple with the Asin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductInfoModel) GetAsinOk() (*string, bool) {
	if o == nil || IsNil(o.Asin) {
		return nil, false
	}
	return o.Asin, true
}

// HasAsin returns a boolean if a field has been set.
func (o *ProductInfoModel) HasAsin() bool {
	if o != nil && !IsNil(o.Asin) {
		return true
	}

	return false
}

// SetAsin gets a reference to the given string and assigns it to the Asin field.
func (o *ProductInfoModel) SetAsin(v string) {
	o.Asin = &v
}

// GetListPrice returns the ListPrice field value if set, zero value otherwise.
func (o *ProductInfoModel) GetListPrice() ListPrice {
	if o == nil || IsNil(o.ListPrice) {
		var ret ListPrice
		return ret
	}
	return *o.ListPrice
}

// GetListPriceOk returns a tuple with the ListPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductInfoModel) GetListPriceOk() (*ListPrice, bool) {
	if o == nil || IsNil(o.ListPrice) {
		return nil, false
	}
	return o.ListPrice, true
}

// HasListPrice returns a boolean if a field has been set.
func (o *ProductInfoModel) HasListPrice() bool {
	if o != nil && !IsNil(o.ListPrice) {
		return true
	}

	return false
}

// SetListPrice gets a reference to the given ListPrice and assigns it to the ListPrice field.
func (o *ProductInfoModel) SetListPrice(v ListPrice) {
	o.ListPrice = &v
}

func (o ProductInfoModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Deal) {
		toSerialize["deal"] = o.Deal
	}
	if !IsNil(o.CatalogType) {
		toSerialize["catalogType"] = o.CatalogType
	}
	if !IsNil(o.TvPreviewMetadata) {
		toSerialize["tvPreviewMetadata"] = o.TvPreviewMetadata
	}
	if !IsNil(o.Availability) {
		toSerialize["availability"] = o.Availability
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.GlobalStoreSetting) {
		toSerialize["globalStoreSetting"] = o.GlobalStoreSetting
	}
	if !IsNil(o.AdEligibility) {
		toSerialize["adEligibility"] = o.AdEligibility
	}
	if !IsNil(o.TitleAuthorityParentAsin) {
		toSerialize["titleAuthorityParentAsin"] = o.TitleAuthorityParentAsin
	}
	if !IsNil(o.Variations) {
		toSerialize["variations"] = o.Variations
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.TvPreviewURL) {
		toSerialize["tvPreviewURL"] = o.TvPreviewURL
	}
	if !IsNil(o.ImageUrl) {
		toSerialize["imageUrl"] = o.ImageUrl
	}
	if !IsNil(o.EligibilityCodes) {
		toSerialize["eligibilityCodes"] = o.EligibilityCodes
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.TitleAuthority) {
		toSerialize["titleAuthority"] = o.TitleAuthority
	}
	if !IsNil(o.IneligibilityReasons) {
		toSerialize["ineligibilityReasons"] = o.IneligibilityReasons
	}
	if !IsNil(o.EligibilityReasons) {
		toSerialize["eligibilityReasons"] = o.EligibilityReasons
	}
	if !IsNil(o.TvIconURL) {
		toSerialize["tvIconURL"] = o.TvIconURL
	}
	if !IsNil(o.Merchant) {
		toSerialize["merchant"] = o.Merchant
	}
	if !IsNil(o.IneligibilityCodes) {
		toSerialize["ineligibilityCodes"] = o.IneligibilityCodes
	}
	if !IsNil(o.EligibilityStatus) {
		toSerialize["eligibilityStatus"] = o.EligibilityStatus
	}
	if !IsNil(o.AsinEligibility) {
		toSerialize["asinEligibility"] = o.AsinEligibility
	}
	if !IsNil(o.CustomerReview) {
		toSerialize["customerReview"] = o.CustomerReview
	}
	if !IsNil(o.ProductSets) {
		toSerialize["productSets"] = o.ProductSets
	}
	if !IsNil(o.Asin) {
		toSerialize["asin"] = o.Asin
	}
	if !IsNil(o.ListPrice) {
		toSerialize["listPrice"] = o.ListPrice
	}
	return toSerialize, nil
}

type NullableProductInfoModel struct {
	value *ProductInfoModel
	isSet bool
}

func (v NullableProductInfoModel) Get() *ProductInfoModel {
	return v.value
}

func (v *NullableProductInfoModel) Set(val *ProductInfoModel) {
	v.value = val
	v.isSet = true
}

func (v NullableProductInfoModel) IsSet() bool {
	return v.isSet
}

func (v *NullableProductInfoModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductInfoModel(val *ProductInfoModel) *NullableProductInfoModel {
	return &NullableProductInfoModel{value: val, isSet: true}
}

func (v NullableProductInfoModel) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableProductInfoModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
