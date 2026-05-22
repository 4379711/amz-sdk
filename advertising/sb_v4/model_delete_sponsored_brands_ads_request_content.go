package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the DeleteSponsoredBrandsAdsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteSponsoredBrandsAdsRequestContent{}

// DeleteSponsoredBrandsAdsRequestContent struct for DeleteSponsoredBrandsAdsRequestContent
type DeleteSponsoredBrandsAdsRequestContent struct {
	AdIdFilter *ObjectIdFilter `json:"adIdFilter,omitempty"`
}

// NewDeleteSponsoredBrandsAdsRequestContent instantiates a new DeleteSponsoredBrandsAdsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteSponsoredBrandsAdsRequestContent() *DeleteSponsoredBrandsAdsRequestContent {
	this := DeleteSponsoredBrandsAdsRequestContent{}
	return &this
}

// NewDeleteSponsoredBrandsAdsRequestContentWithDefaults instantiates a new DeleteSponsoredBrandsAdsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteSponsoredBrandsAdsRequestContentWithDefaults() *DeleteSponsoredBrandsAdsRequestContent {
	this := DeleteSponsoredBrandsAdsRequestContent{}
	return &this
}

// GetAdIdFilter returns the AdIdFilter field value if set, zero value otherwise.
func (o *DeleteSponsoredBrandsAdsRequestContent) GetAdIdFilter() ObjectIdFilter {
	if o == nil || IsNil(o.AdIdFilter) {
		var ret ObjectIdFilter
		return ret
	}
	return *o.AdIdFilter
}

// GetAdIdFilterOk returns a tuple with the AdIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteSponsoredBrandsAdsRequestContent) GetAdIdFilterOk() (*ObjectIdFilter, bool) {
	if o == nil || IsNil(o.AdIdFilter) {
		return nil, false
	}
	return o.AdIdFilter, true
}

// HasAdIdFilter returns a boolean if a field has been set.
func (o *DeleteSponsoredBrandsAdsRequestContent) HasAdIdFilter() bool {
	if o != nil && !IsNil(o.AdIdFilter) {
		return true
	}

	return false
}

// SetAdIdFilter gets a reference to the given ObjectIdFilter and assigns it to the AdIdFilter field.
func (o *DeleteSponsoredBrandsAdsRequestContent) SetAdIdFilter(v ObjectIdFilter) {
	o.AdIdFilter = &v
}

func (o DeleteSponsoredBrandsAdsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdIdFilter) {
		toSerialize["adIdFilter"] = o.AdIdFilter
	}
	return toSerialize, nil
}

type NullableDeleteSponsoredBrandsAdsRequestContent struct {
	value *DeleteSponsoredBrandsAdsRequestContent
	isSet bool
}

func (v NullableDeleteSponsoredBrandsAdsRequestContent) Get() *DeleteSponsoredBrandsAdsRequestContent {
	return v.value
}

func (v *NullableDeleteSponsoredBrandsAdsRequestContent) Set(val *DeleteSponsoredBrandsAdsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteSponsoredBrandsAdsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteSponsoredBrandsAdsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteSponsoredBrandsAdsRequestContent(val *DeleteSponsoredBrandsAdsRequestContent) *NullableDeleteSponsoredBrandsAdsRequestContent {
	return &NullableDeleteSponsoredBrandsAdsRequestContent{value: val, isSet: true}
}

func (v NullableDeleteSponsoredBrandsAdsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDeleteSponsoredBrandsAdsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
