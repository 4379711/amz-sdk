package product_metadata

import (
	"github.com/bytedance/sonic"
)

// checks if the ProductMetadataModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductMetadataModel{}

// ProductMetadataModel struct for ProductMetadataModel
type ProductMetadataModel struct {
	PriceToPay *PriceToPay `json:"priceToPay,omitempty"`
	// List of ineligible status identifier
	IneligibilityCodes []string `json:"ineligibilityCodes,omitempty"`
	// Stock availability:   * IN_STOCK - The item is in stock.   * IN_STOCK_SCARCE - The item is in stock, but stock levels are limited.   * OUT_OF_STOCK - The item is currently out of stock.   * PREORDER - The item is not yet available, but can be pre-ordered.   * LEADTIME - The item is only available after some amount of lead time.   * AVAILABLE_DATE - The item is not available, but will be available on a future date.
	Availability *string `json:"availability,omitempty"`
	// Product title of the item
	Title *string `json:"title,omitempty"`
	// Eligibility status for advertising:   * ELIGIBLE - Eligible for advertising   * INELIGIBLE - Ineligible for advertising
	EligibilityStatus  *string                             `json:"eligibilityStatus,omitempty"`
	BasisPrice         *BasisPrice                         `json:"basisPrice,omitempty"`
	GlobalStoreSetting *ProductInfoModelGlobalStoreSetting `json:"globalStoreSetting,omitempty"`
	// Best seller rank position in the category
	BestSellerRank *string `json:"bestSellerRank,omitempty"`
	// Date the item was first available on Amazon
	CreatedDate *string `json:"createdDate,omitempty"`
	// Url to the product image
	ImageUrl *string `json:"imageUrl,omitempty"`
	// ASIN of the item
	Asin *string `json:"asin,omitempty"`
	// sku of the item
	Sku *string `json:"sku,omitempty"`
	// Category (browse node) name of the ASIN
	Category *string `json:"category,omitempty"`
	// Brand name of the item
	Brand *string `json:"brand,omitempty"`
	// List of ASIN variations of the current item
	VariationList []string `json:"variationList,omitempty"`
	// List of reasons that made this item ineligible to be advertised
	IneligibilityReasons []string `json:"ineligibilityReasons,omitempty"`
}

// NewProductMetadataModel instantiates a new ProductMetadataModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductMetadataModel() *ProductMetadataModel {
	this := ProductMetadataModel{}
	return &this
}

// NewProductMetadataModelWithDefaults instantiates a new ProductMetadataModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductMetadataModelWithDefaults() *ProductMetadataModel {
	this := ProductMetadataModel{}
	return &this
}

// GetPriceToPay returns the PriceToPay field value if set, zero value otherwise.
func (o *ProductMetadataModel) GetPriceToPay() PriceToPay {
	if o == nil || IsNil(o.PriceToPay) {
		var ret PriceToPay
		return ret
	}
	return *o.PriceToPay
}

// GetPriceToPayOk returns a tuple with the PriceToPay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductMetadataModel) GetPriceToPayOk() (*PriceToPay, bool) {
	if o == nil || IsNil(o.PriceToPay) {
		return nil, false
	}
	return o.PriceToPay, true
}

// HasPriceToPay returns a boolean if a field has been set.
func (o *ProductMetadataModel) HasPriceToPay() bool {
	if o != nil && !IsNil(o.PriceToPay) {
		return true
	}

	return false
}

// SetPriceToPay gets a reference to the given PriceToPay and assigns it to the PriceToPay field.
func (o *ProductMetadataModel) SetPriceToPay(v PriceToPay) {
	o.PriceToPay = &v
}

// GetIneligibilityCodes returns the IneligibilityCodes field value if set, zero value otherwise.
func (o *ProductMetadataModel) GetIneligibilityCodes() []string {
	if o == nil || IsNil(o.IneligibilityCodes) {
		var ret []string
		return ret
	}
	return o.IneligibilityCodes
}

// GetIneligibilityCodesOk returns a tuple with the IneligibilityCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductMetadataModel) GetIneligibilityCodesOk() ([]string, bool) {
	if o == nil || IsNil(o.IneligibilityCodes) {
		return nil, false
	}
	return o.IneligibilityCodes, true
}

// HasIneligibilityCodes returns a boolean if a field has been set.
func (o *ProductMetadataModel) HasIneligibilityCodes() bool {
	if o != nil && !IsNil(o.IneligibilityCodes) {
		return true
	}

	return false
}

// SetIneligibilityCodes gets a reference to the given []string and assigns it to the IneligibilityCodes field.
func (o *ProductMetadataModel) SetIneligibilityCodes(v []string) {
	o.IneligibilityCodes = v
}

// GetAvailability returns the Availability field value if set, zero value otherwise.
func (o *ProductMetadataModel) GetAvailability() string {
	if o == nil || IsNil(o.Availability) {
		var ret string
		return ret
	}
	return *o.Availability
}

// GetAvailabilityOk returns a tuple with the Availability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductMetadataModel) GetAvailabilityOk() (*string, bool) {
	if o == nil || IsNil(o.Availability) {
		return nil, false
	}
	return o.Availability, true
}

// HasAvailability returns a boolean if a field has been set.
func (o *ProductMetadataModel) HasAvailability() bool {
	if o != nil && !IsNil(o.Availability) {
		return true
	}

	return false
}

// SetAvailability gets a reference to the given string and assigns it to the Availability field.
func (o *ProductMetadataModel) SetAvailability(v string) {
	o.Availability = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ProductMetadataModel) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductMetadataModel) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ProductMetadataModel) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ProductMetadataModel) SetTitle(v string) {
	o.Title = &v
}

// GetEligibilityStatus returns the EligibilityStatus field value if set, zero value otherwise.
func (o *ProductMetadataModel) GetEligibilityStatus() string {
	if o == nil || IsNil(o.EligibilityStatus) {
		var ret string
		return ret
	}
	return *o.EligibilityStatus
}

// GetEligibilityStatusOk returns a tuple with the EligibilityStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductMetadataModel) GetEligibilityStatusOk() (*string, bool) {
	if o == nil || IsNil(o.EligibilityStatus) {
		return nil, false
	}
	return o.EligibilityStatus, true
}

// HasEligibilityStatus returns a boolean if a field has been set.
func (o *ProductMetadataModel) HasEligibilityStatus() bool {
	if o != nil && !IsNil(o.EligibilityStatus) {
		return true
	}

	return false
}

// SetEligibilityStatus gets a reference to the given string and assigns it to the EligibilityStatus field.
func (o *ProductMetadataModel) SetEligibilityStatus(v string) {
	o.EligibilityStatus = &v
}

// GetBasisPrice returns the BasisPrice field value if set, zero value otherwise.
func (o *ProductMetadataModel) GetBasisPrice() BasisPrice {
	if o == nil || IsNil(o.BasisPrice) {
		var ret BasisPrice
		return ret
	}
	return *o.BasisPrice
}

// GetBasisPriceOk returns a tuple with the BasisPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductMetadataModel) GetBasisPriceOk() (*BasisPrice, bool) {
	if o == nil || IsNil(o.BasisPrice) {
		return nil, false
	}
	return o.BasisPrice, true
}

// HasBasisPrice returns a boolean if a field has been set.
func (o *ProductMetadataModel) HasBasisPrice() bool {
	if o != nil && !IsNil(o.BasisPrice) {
		return true
	}

	return false
}

// SetBasisPrice gets a reference to the given BasisPrice and assigns it to the BasisPrice field.
func (o *ProductMetadataModel) SetBasisPrice(v BasisPrice) {
	o.BasisPrice = &v
}

// GetGlobalStoreSetting returns the GlobalStoreSetting field value if set, zero value otherwise.
func (o *ProductMetadataModel) GetGlobalStoreSetting() ProductInfoModelGlobalStoreSetting {
	if o == nil || IsNil(o.GlobalStoreSetting) {
		var ret ProductInfoModelGlobalStoreSetting
		return ret
	}
	return *o.GlobalStoreSetting
}

// GetGlobalStoreSettingOk returns a tuple with the GlobalStoreSetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductMetadataModel) GetGlobalStoreSettingOk() (*ProductInfoModelGlobalStoreSetting, bool) {
	if o == nil || IsNil(o.GlobalStoreSetting) {
		return nil, false
	}
	return o.GlobalStoreSetting, true
}

// HasGlobalStoreSetting returns a boolean if a field has been set.
func (o *ProductMetadataModel) HasGlobalStoreSetting() bool {
	if o != nil && !IsNil(o.GlobalStoreSetting) {
		return true
	}

	return false
}

// SetGlobalStoreSetting gets a reference to the given ProductInfoModelGlobalStoreSetting and assigns it to the GlobalStoreSetting field.
func (o *ProductMetadataModel) SetGlobalStoreSetting(v ProductInfoModelGlobalStoreSetting) {
	o.GlobalStoreSetting = &v
}

// GetBestSellerRank returns the BestSellerRank field value if set, zero value otherwise.
func (o *ProductMetadataModel) GetBestSellerRank() string {
	if o == nil || IsNil(o.BestSellerRank) {
		var ret string
		return ret
	}
	return *o.BestSellerRank
}

// GetBestSellerRankOk returns a tuple with the BestSellerRank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductMetadataModel) GetBestSellerRankOk() (*string, bool) {
	if o == nil || IsNil(o.BestSellerRank) {
		return nil, false
	}
	return o.BestSellerRank, true
}

// HasBestSellerRank returns a boolean if a field has been set.
func (o *ProductMetadataModel) HasBestSellerRank() bool {
	if o != nil && !IsNil(o.BestSellerRank) {
		return true
	}

	return false
}

// SetBestSellerRank gets a reference to the given string and assigns it to the BestSellerRank field.
func (o *ProductMetadataModel) SetBestSellerRank(v string) {
	o.BestSellerRank = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *ProductMetadataModel) GetCreatedDate() string {
	if o == nil || IsNil(o.CreatedDate) {
		var ret string
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductMetadataModel) GetCreatedDateOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *ProductMetadataModel) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given string and assigns it to the CreatedDate field.
func (o *ProductMetadataModel) SetCreatedDate(v string) {
	o.CreatedDate = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *ProductMetadataModel) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductMetadataModel) GetImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *ProductMetadataModel) HasImageUrl() bool {
	if o != nil && !IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *ProductMetadataModel) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetAsin returns the Asin field value if set, zero value otherwise.
func (o *ProductMetadataModel) GetAsin() string {
	if o == nil || IsNil(o.Asin) {
		var ret string
		return ret
	}
	return *o.Asin
}

// GetAsinOk returns a tuple with the Asin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductMetadataModel) GetAsinOk() (*string, bool) {
	if o == nil || IsNil(o.Asin) {
		return nil, false
	}
	return o.Asin, true
}

// HasAsin returns a boolean if a field has been set.
func (o *ProductMetadataModel) HasAsin() bool {
	if o != nil && !IsNil(o.Asin) {
		return true
	}

	return false
}

// SetAsin gets a reference to the given string and assigns it to the Asin field.
func (o *ProductMetadataModel) SetAsin(v string) {
	o.Asin = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *ProductMetadataModel) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductMetadataModel) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *ProductMetadataModel) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *ProductMetadataModel) SetSku(v string) {
	o.Sku = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *ProductMetadataModel) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductMetadataModel) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *ProductMetadataModel) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *ProductMetadataModel) SetCategory(v string) {
	o.Category = &v
}

// GetBrand returns the Brand field value if set, zero value otherwise.
func (o *ProductMetadataModel) GetBrand() string {
	if o == nil || IsNil(o.Brand) {
		var ret string
		return ret
	}
	return *o.Brand
}

// GetBrandOk returns a tuple with the Brand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductMetadataModel) GetBrandOk() (*string, bool) {
	if o == nil || IsNil(o.Brand) {
		return nil, false
	}
	return o.Brand, true
}

// HasBrand returns a boolean if a field has been set.
func (o *ProductMetadataModel) HasBrand() bool {
	if o != nil && !IsNil(o.Brand) {
		return true
	}

	return false
}

// SetBrand gets a reference to the given string and assigns it to the Brand field.
func (o *ProductMetadataModel) SetBrand(v string) {
	o.Brand = &v
}

// GetVariationList returns the VariationList field value if set, zero value otherwise.
func (o *ProductMetadataModel) GetVariationList() []string {
	if o == nil || IsNil(o.VariationList) {
		var ret []string
		return ret
	}
	return o.VariationList
}

// GetVariationListOk returns a tuple with the VariationList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductMetadataModel) GetVariationListOk() ([]string, bool) {
	if o == nil || IsNil(o.VariationList) {
		return nil, false
	}
	return o.VariationList, true
}

// HasVariationList returns a boolean if a field has been set.
func (o *ProductMetadataModel) HasVariationList() bool {
	if o != nil && !IsNil(o.VariationList) {
		return true
	}

	return false
}

// SetVariationList gets a reference to the given []string and assigns it to the VariationList field.
func (o *ProductMetadataModel) SetVariationList(v []string) {
	o.VariationList = v
}

// GetIneligibilityReasons returns the IneligibilityReasons field value if set, zero value otherwise.
func (o *ProductMetadataModel) GetIneligibilityReasons() []string {
	if o == nil || IsNil(o.IneligibilityReasons) {
		var ret []string
		return ret
	}
	return o.IneligibilityReasons
}

// GetIneligibilityReasonsOk returns a tuple with the IneligibilityReasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductMetadataModel) GetIneligibilityReasonsOk() ([]string, bool) {
	if o == nil || IsNil(o.IneligibilityReasons) {
		return nil, false
	}
	return o.IneligibilityReasons, true
}

// HasIneligibilityReasons returns a boolean if a field has been set.
func (o *ProductMetadataModel) HasIneligibilityReasons() bool {
	if o != nil && !IsNil(o.IneligibilityReasons) {
		return true
	}

	return false
}

// SetIneligibilityReasons gets a reference to the given []string and assigns it to the IneligibilityReasons field.
func (o *ProductMetadataModel) SetIneligibilityReasons(v []string) {
	o.IneligibilityReasons = v
}

func (o ProductMetadataModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PriceToPay) {
		toSerialize["priceToPay"] = o.PriceToPay
	}
	if !IsNil(o.IneligibilityCodes) {
		toSerialize["ineligibilityCodes"] = o.IneligibilityCodes
	}
	if !IsNil(o.Availability) {
		toSerialize["availability"] = o.Availability
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.EligibilityStatus) {
		toSerialize["eligibilityStatus"] = o.EligibilityStatus
	}
	if !IsNil(o.BasisPrice) {
		toSerialize["basisPrice"] = o.BasisPrice
	}
	if !IsNil(o.GlobalStoreSetting) {
		toSerialize["globalStoreSetting"] = o.GlobalStoreSetting
	}
	if !IsNil(o.BestSellerRank) {
		toSerialize["bestSellerRank"] = o.BestSellerRank
	}
	if !IsNil(o.CreatedDate) {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if !IsNil(o.ImageUrl) {
		toSerialize["imageUrl"] = o.ImageUrl
	}
	if !IsNil(o.Asin) {
		toSerialize["asin"] = o.Asin
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Brand) {
		toSerialize["brand"] = o.Brand
	}
	if !IsNil(o.VariationList) {
		toSerialize["variationList"] = o.VariationList
	}
	if !IsNil(o.IneligibilityReasons) {
		toSerialize["ineligibilityReasons"] = o.IneligibilityReasons
	}
	return toSerialize, nil
}

type NullableProductMetadataModel struct {
	value *ProductMetadataModel
	isSet bool
}

func (v NullableProductMetadataModel) Get() *ProductMetadataModel {
	return v.value
}

func (v *NullableProductMetadataModel) Set(val *ProductMetadataModel) {
	v.value = val
	v.isSet = true
}

func (v NullableProductMetadataModel) IsSet() bool {
	return v.isSet
}

func (v *NullableProductMetadataModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductMetadataModel(val *ProductMetadataModel) *NullableProductMetadataModel {
	return &NullableProductMetadataModel{value: val, isSet: true}
}

func (v NullableProductMetadataModel) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableProductMetadataModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
