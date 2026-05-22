package product_metadata

import (
	"github.com/bytedance/sonic"
)

// checks if the ProductMetadataRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductMetadataRequest{}

// ProductMetadataRequest struct for ProductMetadataRequest
type ProductMetadataRequest struct {
	// Whether item details such as name, image, and price is required.
	CheckItemDetails *bool `json:"checkItemDetails,omitempty"`
	// Specific SKUs to search for in the advertiser's inventory. Currently only support SP program type for sellers. Cannot use together with asins or searchStr input types.
	Skus []string `json:"skus,omitempty"`
	// Whether advertising eligibility info is required
	CheckEligibility *bool `json:"checkEligibility,omitempty"`
	//  This will return only GlobalStore listings related to [GlobalStore Program](https://sellercentral.amazon.com/help/hub/reference/external/202139180) and not local listings
	IsGlobalStoreSelection *bool `json:"isGlobalStoreSelection,omitempty"`
	// Number of items to be returned on this page index.
	PageSize int32 `json:"pageSize"`
	// Optional locale for detail and eligibility response strings. Default to the marketplace locale.
	Locale *string `json:"locale,omitempty"`
	// Specific asins to search for in the advertiser's inventory. Cannot use together with skus or searchStr input types.
	Asins []string `json:"asins,omitempty"`
	// Pagination token used for the suggested sort type or for author merchant
	CursorToken *string `json:"cursorToken,omitempty"`
	// Program type. Required if checks advertising eligibility:   * SP - Sponsored Product   * SB - Sponsored Brand   * SD - Sponsored Display
	AdType *string `json:"adType,omitempty"`
	// Specific string in the item title to search for in the advertiser's inventory. Case insensitive. Cannot use together with asins or skus input types.
	SearchStr *string `json:"searchStr,omitempty"`
	// Index of the page to be returned; For author, this value will be ignored, should use cursorToken instead. For seller and vendor, results are capped at 10k ((pageIndex + 1) * pageSize).
	PageIndex int32 `json:"pageIndex"`
	// Sort order (has to be DESC for the suggested sort type):   * ASC - Ascending, from A to Z   * DESC - Descending, from Z to A
	SortOrder *string `json:"sortOrder,omitempty"`
	// Sort option for the result. Currently only support SP program type for sellers:   * SUGGESTED - Suggested products are those most likely to engage customers, and have a higher chance of generating clicks if advertised.   * CREATED_DATE - Date the item listing was created
	SortBy *string `json:"sortBy,omitempty"`
}

type _ProductMetadataRequest ProductMetadataRequest

// NewProductMetadataRequest instantiates a new ProductMetadataRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductMetadataRequest(pageSize int32, pageIndex int32) *ProductMetadataRequest {
	this := ProductMetadataRequest{}
	var checkItemDetails bool = false
	this.CheckItemDetails = &checkItemDetails
	var checkEligibility bool = false
	this.CheckEligibility = &checkEligibility
	var isGlobalStoreSelection bool = false
	this.IsGlobalStoreSelection = &isGlobalStoreSelection
	this.PageSize = pageSize
	this.PageIndex = pageIndex
	var sortOrder string = "DESC"
	this.SortOrder = &sortOrder
	return &this
}

// NewProductMetadataRequestWithDefaults instantiates a new ProductMetadataRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductMetadataRequestWithDefaults() *ProductMetadataRequest {
	this := ProductMetadataRequest{}
	var checkItemDetails bool = false
	this.CheckItemDetails = &checkItemDetails
	var checkEligibility bool = false
	this.CheckEligibility = &checkEligibility
	var isGlobalStoreSelection bool = false
	this.IsGlobalStoreSelection = &isGlobalStoreSelection
	var sortOrder string = "DESC"
	this.SortOrder = &sortOrder
	return &this
}

// GetCheckItemDetails returns the CheckItemDetails field value if set, zero value otherwise.
func (o *ProductMetadataRequest) GetCheckItemDetails() bool {
	if o == nil || IsNil(o.CheckItemDetails) {
		var ret bool
		return ret
	}
	return *o.CheckItemDetails
}

// GetCheckItemDetailsOk returns a tuple with the CheckItemDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductMetadataRequest) GetCheckItemDetailsOk() (*bool, bool) {
	if o == nil || IsNil(o.CheckItemDetails) {
		return nil, false
	}
	return o.CheckItemDetails, true
}

// HasCheckItemDetails returns a boolean if a field has been set.
func (o *ProductMetadataRequest) HasCheckItemDetails() bool {
	if o != nil && !IsNil(o.CheckItemDetails) {
		return true
	}

	return false
}

// SetCheckItemDetails gets a reference to the given bool and assigns it to the CheckItemDetails field.
func (o *ProductMetadataRequest) SetCheckItemDetails(v bool) {
	o.CheckItemDetails = &v
}

// GetSkus returns the Skus field value if set, zero value otherwise.
func (o *ProductMetadataRequest) GetSkus() []string {
	if o == nil || IsNil(o.Skus) {
		var ret []string
		return ret
	}
	return o.Skus
}

// GetSkusOk returns a tuple with the Skus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductMetadataRequest) GetSkusOk() ([]string, bool) {
	if o == nil || IsNil(o.Skus) {
		return nil, false
	}
	return o.Skus, true
}

// HasSkus returns a boolean if a field has been set.
func (o *ProductMetadataRequest) HasSkus() bool {
	if o != nil && !IsNil(o.Skus) {
		return true
	}

	return false
}

// SetSkus gets a reference to the given []string and assigns it to the Skus field.
func (o *ProductMetadataRequest) SetSkus(v []string) {
	o.Skus = v
}

// GetCheckEligibility returns the CheckEligibility field value if set, zero value otherwise.
func (o *ProductMetadataRequest) GetCheckEligibility() bool {
	if o == nil || IsNil(o.CheckEligibility) {
		var ret bool
		return ret
	}
	return *o.CheckEligibility
}

// GetCheckEligibilityOk returns a tuple with the CheckEligibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductMetadataRequest) GetCheckEligibilityOk() (*bool, bool) {
	if o == nil || IsNil(o.CheckEligibility) {
		return nil, false
	}
	return o.CheckEligibility, true
}

// HasCheckEligibility returns a boolean if a field has been set.
func (o *ProductMetadataRequest) HasCheckEligibility() bool {
	if o != nil && !IsNil(o.CheckEligibility) {
		return true
	}

	return false
}

// SetCheckEligibility gets a reference to the given bool and assigns it to the CheckEligibility field.
func (o *ProductMetadataRequest) SetCheckEligibility(v bool) {
	o.CheckEligibility = &v
}

// GetIsGlobalStoreSelection returns the IsGlobalStoreSelection field value if set, zero value otherwise.
func (o *ProductMetadataRequest) GetIsGlobalStoreSelection() bool {
	if o == nil || IsNil(o.IsGlobalStoreSelection) {
		var ret bool
		return ret
	}
	return *o.IsGlobalStoreSelection
}

// GetIsGlobalStoreSelectionOk returns a tuple with the IsGlobalStoreSelection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductMetadataRequest) GetIsGlobalStoreSelectionOk() (*bool, bool) {
	if o == nil || IsNil(o.IsGlobalStoreSelection) {
		return nil, false
	}
	return o.IsGlobalStoreSelection, true
}

// HasIsGlobalStoreSelection returns a boolean if a field has been set.
func (o *ProductMetadataRequest) HasIsGlobalStoreSelection() bool {
	if o != nil && !IsNil(o.IsGlobalStoreSelection) {
		return true
	}

	return false
}

// SetIsGlobalStoreSelection gets a reference to the given bool and assigns it to the IsGlobalStoreSelection field.
func (o *ProductMetadataRequest) SetIsGlobalStoreSelection(v bool) {
	o.IsGlobalStoreSelection = &v
}

// GetPageSize returns the PageSize field value
func (o *ProductMetadataRequest) GetPageSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *ProductMetadataRequest) GetPageSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *ProductMetadataRequest) SetPageSize(v int32) {
	o.PageSize = v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *ProductMetadataRequest) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductMetadataRequest) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *ProductMetadataRequest) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *ProductMetadataRequest) SetLocale(v string) {
	o.Locale = &v
}

// GetAsins returns the Asins field value if set, zero value otherwise.
func (o *ProductMetadataRequest) GetAsins() []string {
	if o == nil || IsNil(o.Asins) {
		var ret []string
		return ret
	}
	return o.Asins
}

// GetAsinsOk returns a tuple with the Asins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductMetadataRequest) GetAsinsOk() ([]string, bool) {
	if o == nil || IsNil(o.Asins) {
		return nil, false
	}
	return o.Asins, true
}

// HasAsins returns a boolean if a field has been set.
func (o *ProductMetadataRequest) HasAsins() bool {
	if o != nil && !IsNil(o.Asins) {
		return true
	}

	return false
}

// SetAsins gets a reference to the given []string and assigns it to the Asins field.
func (o *ProductMetadataRequest) SetAsins(v []string) {
	o.Asins = v
}

// GetCursorToken returns the CursorToken field value if set, zero value otherwise.
func (o *ProductMetadataRequest) GetCursorToken() string {
	if o == nil || IsNil(o.CursorToken) {
		var ret string
		return ret
	}
	return *o.CursorToken
}

// GetCursorTokenOk returns a tuple with the CursorToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductMetadataRequest) GetCursorTokenOk() (*string, bool) {
	if o == nil || IsNil(o.CursorToken) {
		return nil, false
	}
	return o.CursorToken, true
}

// HasCursorToken returns a boolean if a field has been set.
func (o *ProductMetadataRequest) HasCursorToken() bool {
	if o != nil && !IsNil(o.CursorToken) {
		return true
	}

	return false
}

// SetCursorToken gets a reference to the given string and assigns it to the CursorToken field.
func (o *ProductMetadataRequest) SetCursorToken(v string) {
	o.CursorToken = &v
}

// GetAdType returns the AdType field value if set, zero value otherwise.
func (o *ProductMetadataRequest) GetAdType() string {
	if o == nil || IsNil(o.AdType) {
		var ret string
		return ret
	}
	return *o.AdType
}

// GetAdTypeOk returns a tuple with the AdType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductMetadataRequest) GetAdTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AdType) {
		return nil, false
	}
	return o.AdType, true
}

// HasAdType returns a boolean if a field has been set.
func (o *ProductMetadataRequest) HasAdType() bool {
	if o != nil && !IsNil(o.AdType) {
		return true
	}

	return false
}

// SetAdType gets a reference to the given string and assigns it to the AdType field.
func (o *ProductMetadataRequest) SetAdType(v string) {
	o.AdType = &v
}

// GetSearchStr returns the SearchStr field value if set, zero value otherwise.
func (o *ProductMetadataRequest) GetSearchStr() string {
	if o == nil || IsNil(o.SearchStr) {
		var ret string
		return ret
	}
	return *o.SearchStr
}

// GetSearchStrOk returns a tuple with the SearchStr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductMetadataRequest) GetSearchStrOk() (*string, bool) {
	if o == nil || IsNil(o.SearchStr) {
		return nil, false
	}
	return o.SearchStr, true
}

// HasSearchStr returns a boolean if a field has been set.
func (o *ProductMetadataRequest) HasSearchStr() bool {
	if o != nil && !IsNil(o.SearchStr) {
		return true
	}

	return false
}

// SetSearchStr gets a reference to the given string and assigns it to the SearchStr field.
func (o *ProductMetadataRequest) SetSearchStr(v string) {
	o.SearchStr = &v
}

// GetPageIndex returns the PageIndex field value
func (o *ProductMetadataRequest) GetPageIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PageIndex
}

// GetPageIndexOk returns a tuple with the PageIndex field value
// and a boolean to check if the value has been set.
func (o *ProductMetadataRequest) GetPageIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageIndex, true
}

// SetPageIndex sets field value
func (o *ProductMetadataRequest) SetPageIndex(v int32) {
	o.PageIndex = v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *ProductMetadataRequest) GetSortOrder() string {
	if o == nil || IsNil(o.SortOrder) {
		var ret string
		return ret
	}
	return *o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductMetadataRequest) GetSortOrderOk() (*string, bool) {
	if o == nil || IsNil(o.SortOrder) {
		return nil, false
	}
	return o.SortOrder, true
}

// HasSortOrder returns a boolean if a field has been set.
func (o *ProductMetadataRequest) HasSortOrder() bool {
	if o != nil && !IsNil(o.SortOrder) {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given string and assigns it to the SortOrder field.
func (o *ProductMetadataRequest) SetSortOrder(v string) {
	o.SortOrder = &v
}

// GetSortBy returns the SortBy field value if set, zero value otherwise.
func (o *ProductMetadataRequest) GetSortBy() string {
	if o == nil || IsNil(o.SortBy) {
		var ret string
		return ret
	}
	return *o.SortBy
}

// GetSortByOk returns a tuple with the SortBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductMetadataRequest) GetSortByOk() (*string, bool) {
	if o == nil || IsNil(o.SortBy) {
		return nil, false
	}
	return o.SortBy, true
}

// HasSortBy returns a boolean if a field has been set.
func (o *ProductMetadataRequest) HasSortBy() bool {
	if o != nil && !IsNil(o.SortBy) {
		return true
	}

	return false
}

// SetSortBy gets a reference to the given string and assigns it to the SortBy field.
func (o *ProductMetadataRequest) SetSortBy(v string) {
	o.SortBy = &v
}

func (o ProductMetadataRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CheckItemDetails) {
		toSerialize["checkItemDetails"] = o.CheckItemDetails
	}
	if !IsNil(o.Skus) {
		toSerialize["skus"] = o.Skus
	}
	if !IsNil(o.CheckEligibility) {
		toSerialize["checkEligibility"] = o.CheckEligibility
	}
	if !IsNil(o.IsGlobalStoreSelection) {
		toSerialize["isGlobalStoreSelection"] = o.IsGlobalStoreSelection
	}
	toSerialize["pageSize"] = o.PageSize
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	if !IsNil(o.Asins) {
		toSerialize["asins"] = o.Asins
	}
	if !IsNil(o.CursorToken) {
		toSerialize["cursorToken"] = o.CursorToken
	}
	if !IsNil(o.AdType) {
		toSerialize["adType"] = o.AdType
	}
	if !IsNil(o.SearchStr) {
		toSerialize["searchStr"] = o.SearchStr
	}
	toSerialize["pageIndex"] = o.PageIndex
	if !IsNil(o.SortOrder) {
		toSerialize["sortOrder"] = o.SortOrder
	}
	if !IsNil(o.SortBy) {
		toSerialize["sortBy"] = o.SortBy
	}
	return toSerialize, nil
}

type NullableProductMetadataRequest struct {
	value *ProductMetadataRequest
	isSet bool
}

func (v NullableProductMetadataRequest) Get() *ProductMetadataRequest {
	return v.value
}

func (v *NullableProductMetadataRequest) Set(val *ProductMetadataRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductMetadataRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductMetadataRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductMetadataRequest(val *ProductMetadataRequest) *NullableProductMetadataRequest {
	return &NullableProductMetadataRequest{value: val, isSet: true}
}

func (v NullableProductMetadataRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableProductMetadataRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
