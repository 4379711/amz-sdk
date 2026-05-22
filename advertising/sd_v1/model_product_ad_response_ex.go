package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the ProductAdResponseEx type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductAdResponseEx{}

// ProductAdResponseEx struct for ProductAdResponseEx
type ProductAdResponseEx struct {
	// The identifier of the ad.
	AdId *float32 `json:"adId,omitempty"`
	// The identifier of the ad group associated with the ad.
	AdGroupId *float32 `json:"adGroupId,omitempty"`
	// The identifier of the campaign associated with the ad.
	CampaignId *float32 `json:"campaignId,omitempty"`
	// The URL where customers will land after clicking on its link. Must be provided if a LandingPageType is set. Please note that if a single product ad sets the landing page url, only one product ad can be added to the ad group. This field is not supported when using ASIN or SKU fields. ||Specifications| |------------------|------------------| |LandingPageType| Description| |STORE| The url should be in the format of https://www.amazon.com/stores/_* (using a correct Amazon url based on the marketplace)| |OFF_AMAZON_LINK| The url should be in the format of https://www.****.com. Note that this LandingPageType is not supported when using ASIN or SKU fields. A custom creative of headline, logo, image are require for this LandingPageType. | |MOMENT| Not yet supported. The url should be in the format of https://www.amazon.com/moments/promotion/{campaignId} (using a correct Amazon url based on the marketplace)|
	LandingPageURL  *string          `json:"landingPageURL,omitempty"`
	LandingPageType *LandingPageType `json:"landingPageType,omitempty"`
	// The name of the ad. Note that this field is not supported when using ASIN or SKU fields.
	AdName *string `json:"adName,omitempty"`
	// The ASIN of the product being advertised. This parameter is included in the response for sellers and vendors.
	Asin *string `json:"asin,omitempty"`
	// The SKU of the product being advertised. This parameter is included in the response for sellers.
	Sku *string `json:"sku,omitempty"`
	// The state of the product ad.
	State *string `json:"state,omitempty"`
	// The status of the product ad.
	ServingStatus *string `json:"servingStatus,omitempty"`
	// Epoch date the product ad was created.
	CreationDate *int64 `json:"creationDate,omitempty"`
	// Epoch date of the last update to any property associated with the product ad.
	LastUpdatedDate *int64 `json:"lastUpdatedDate,omitempty"`
}

// NewProductAdResponseEx instantiates a new ProductAdResponseEx object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductAdResponseEx() *ProductAdResponseEx {
	this := ProductAdResponseEx{}
	return &this
}

// NewProductAdResponseExWithDefaults instantiates a new ProductAdResponseEx object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductAdResponseExWithDefaults() *ProductAdResponseEx {
	this := ProductAdResponseEx{}
	return &this
}

// GetAdId returns the AdId field value if set, zero value otherwise.
func (o *ProductAdResponseEx) GetAdId() float32 {
	if o == nil || IsNil(o.AdId) {
		var ret float32
		return ret
	}
	return *o.AdId
}

// GetAdIdOk returns a tuple with the AdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAdResponseEx) GetAdIdOk() (*float32, bool) {
	if o == nil || IsNil(o.AdId) {
		return nil, false
	}
	return o.AdId, true
}

// HasAdId returns a boolean if a field has been set.
func (o *ProductAdResponseEx) HasAdId() bool {
	if o != nil && !IsNil(o.AdId) {
		return true
	}

	return false
}

// SetAdId gets a reference to the given float32 and assigns it to the AdId field.
func (o *ProductAdResponseEx) SetAdId(v float32) {
	o.AdId = &v
}

// GetAdGroupId returns the AdGroupId field value if set, zero value otherwise.
func (o *ProductAdResponseEx) GetAdGroupId() float32 {
	if o == nil || IsNil(o.AdGroupId) {
		var ret float32
		return ret
	}
	return *o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAdResponseEx) GetAdGroupIdOk() (*float32, bool) {
	if o == nil || IsNil(o.AdGroupId) {
		return nil, false
	}
	return o.AdGroupId, true
}

// HasAdGroupId returns a boolean if a field has been set.
func (o *ProductAdResponseEx) HasAdGroupId() bool {
	if o != nil && !IsNil(o.AdGroupId) {
		return true
	}

	return false
}

// SetAdGroupId gets a reference to the given float32 and assigns it to the AdGroupId field.
func (o *ProductAdResponseEx) SetAdGroupId(v float32) {
	o.AdGroupId = &v
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *ProductAdResponseEx) GetCampaignId() float32 {
	if o == nil || IsNil(o.CampaignId) {
		var ret float32
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAdResponseEx) GetCampaignIdOk() (*float32, bool) {
	if o == nil || IsNil(o.CampaignId) {
		return nil, false
	}
	return o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *ProductAdResponseEx) HasCampaignId() bool {
	if o != nil && !IsNil(o.CampaignId) {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given float32 and assigns it to the CampaignId field.
func (o *ProductAdResponseEx) SetCampaignId(v float32) {
	o.CampaignId = &v
}

// GetLandingPageURL returns the LandingPageURL field value if set, zero value otherwise.
func (o *ProductAdResponseEx) GetLandingPageURL() string {
	if o == nil || IsNil(o.LandingPageURL) {
		var ret string
		return ret
	}
	return *o.LandingPageURL
}

// GetLandingPageURLOk returns a tuple with the LandingPageURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAdResponseEx) GetLandingPageURLOk() (*string, bool) {
	if o == nil || IsNil(o.LandingPageURL) {
		return nil, false
	}
	return o.LandingPageURL, true
}

// HasLandingPageURL returns a boolean if a field has been set.
func (o *ProductAdResponseEx) HasLandingPageURL() bool {
	if o != nil && !IsNil(o.LandingPageURL) {
		return true
	}

	return false
}

// SetLandingPageURL gets a reference to the given string and assigns it to the LandingPageURL field.
func (o *ProductAdResponseEx) SetLandingPageURL(v string) {
	o.LandingPageURL = &v
}

// GetLandingPageType returns the LandingPageType field value if set, zero value otherwise.
func (o *ProductAdResponseEx) GetLandingPageType() LandingPageType {
	if o == nil || IsNil(o.LandingPageType) {
		var ret LandingPageType
		return ret
	}
	return *o.LandingPageType
}

// GetLandingPageTypeOk returns a tuple with the LandingPageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAdResponseEx) GetLandingPageTypeOk() (*LandingPageType, bool) {
	if o == nil || IsNil(o.LandingPageType) {
		return nil, false
	}
	return o.LandingPageType, true
}

// HasLandingPageType returns a boolean if a field has been set.
func (o *ProductAdResponseEx) HasLandingPageType() bool {
	if o != nil && !IsNil(o.LandingPageType) {
		return true
	}

	return false
}

// SetLandingPageType gets a reference to the given LandingPageType and assigns it to the LandingPageType field.
func (o *ProductAdResponseEx) SetLandingPageType(v LandingPageType) {
	o.LandingPageType = &v
}

// GetAdName returns the AdName field value if set, zero value otherwise.
func (o *ProductAdResponseEx) GetAdName() string {
	if o == nil || IsNil(o.AdName) {
		var ret string
		return ret
	}
	return *o.AdName
}

// GetAdNameOk returns a tuple with the AdName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAdResponseEx) GetAdNameOk() (*string, bool) {
	if o == nil || IsNil(o.AdName) {
		return nil, false
	}
	return o.AdName, true
}

// HasAdName returns a boolean if a field has been set.
func (o *ProductAdResponseEx) HasAdName() bool {
	if o != nil && !IsNil(o.AdName) {
		return true
	}

	return false
}

// SetAdName gets a reference to the given string and assigns it to the AdName field.
func (o *ProductAdResponseEx) SetAdName(v string) {
	o.AdName = &v
}

// GetAsin returns the Asin field value if set, zero value otherwise.
func (o *ProductAdResponseEx) GetAsin() string {
	if o == nil || IsNil(o.Asin) {
		var ret string
		return ret
	}
	return *o.Asin
}

// GetAsinOk returns a tuple with the Asin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAdResponseEx) GetAsinOk() (*string, bool) {
	if o == nil || IsNil(o.Asin) {
		return nil, false
	}
	return o.Asin, true
}

// HasAsin returns a boolean if a field has been set.
func (o *ProductAdResponseEx) HasAsin() bool {
	if o != nil && !IsNil(o.Asin) {
		return true
	}

	return false
}

// SetAsin gets a reference to the given string and assigns it to the Asin field.
func (o *ProductAdResponseEx) SetAsin(v string) {
	o.Asin = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *ProductAdResponseEx) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAdResponseEx) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *ProductAdResponseEx) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *ProductAdResponseEx) SetSku(v string) {
	o.Sku = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ProductAdResponseEx) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAdResponseEx) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ProductAdResponseEx) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ProductAdResponseEx) SetState(v string) {
	o.State = &v
}

// GetServingStatus returns the ServingStatus field value if set, zero value otherwise.
func (o *ProductAdResponseEx) GetServingStatus() string {
	if o == nil || IsNil(o.ServingStatus) {
		var ret string
		return ret
	}
	return *o.ServingStatus
}

// GetServingStatusOk returns a tuple with the ServingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAdResponseEx) GetServingStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ServingStatus) {
		return nil, false
	}
	return o.ServingStatus, true
}

// HasServingStatus returns a boolean if a field has been set.
func (o *ProductAdResponseEx) HasServingStatus() bool {
	if o != nil && !IsNil(o.ServingStatus) {
		return true
	}

	return false
}

// SetServingStatus gets a reference to the given string and assigns it to the ServingStatus field.
func (o *ProductAdResponseEx) SetServingStatus(v string) {
	o.ServingStatus = &v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *ProductAdResponseEx) GetCreationDate() int64 {
	if o == nil || IsNil(o.CreationDate) {
		var ret int64
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAdResponseEx) GetCreationDateOk() (*int64, bool) {
	if o == nil || IsNil(o.CreationDate) {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *ProductAdResponseEx) HasCreationDate() bool {
	if o != nil && !IsNil(o.CreationDate) {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given int64 and assigns it to the CreationDate field.
func (o *ProductAdResponseEx) SetCreationDate(v int64) {
	o.CreationDate = &v
}

// GetLastUpdatedDate returns the LastUpdatedDate field value if set, zero value otherwise.
func (o *ProductAdResponseEx) GetLastUpdatedDate() int64 {
	if o == nil || IsNil(o.LastUpdatedDate) {
		var ret int64
		return ret
	}
	return *o.LastUpdatedDate
}

// GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAdResponseEx) GetLastUpdatedDateOk() (*int64, bool) {
	if o == nil || IsNil(o.LastUpdatedDate) {
		return nil, false
	}
	return o.LastUpdatedDate, true
}

// HasLastUpdatedDate returns a boolean if a field has been set.
func (o *ProductAdResponseEx) HasLastUpdatedDate() bool {
	if o != nil && !IsNil(o.LastUpdatedDate) {
		return true
	}

	return false
}

// SetLastUpdatedDate gets a reference to the given int64 and assigns it to the LastUpdatedDate field.
func (o *ProductAdResponseEx) SetLastUpdatedDate(v int64) {
	o.LastUpdatedDate = &v
}

func (o ProductAdResponseEx) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdId) {
		toSerialize["adId"] = o.AdId
	}
	if !IsNil(o.AdGroupId) {
		toSerialize["adGroupId"] = o.AdGroupId
	}
	if !IsNil(o.CampaignId) {
		toSerialize["campaignId"] = o.CampaignId
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
	if !IsNil(o.Asin) {
		toSerialize["asin"] = o.Asin
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.ServingStatus) {
		toSerialize["servingStatus"] = o.ServingStatus
	}
	if !IsNil(o.CreationDate) {
		toSerialize["creationDate"] = o.CreationDate
	}
	if !IsNil(o.LastUpdatedDate) {
		toSerialize["lastUpdatedDate"] = o.LastUpdatedDate
	}
	return toSerialize, nil
}

type NullableProductAdResponseEx struct {
	value *ProductAdResponseEx
	isSet bool
}

func (v NullableProductAdResponseEx) Get() *ProductAdResponseEx {
	return v.value
}

func (v *NullableProductAdResponseEx) Set(val *ProductAdResponseEx) {
	v.value = val
	v.isSet = true
}

func (v NullableProductAdResponseEx) IsSet() bool {
	return v.isSet
}

func (v *NullableProductAdResponseEx) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductAdResponseEx(val *ProductAdResponseEx) *NullableProductAdResponseEx {
	return &NullableProductAdResponseEx{value: val, isSet: true}
}

func (v NullableProductAdResponseEx) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableProductAdResponseEx) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
