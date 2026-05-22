package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the DeleteSponsoredBrandsAdGroupsBetaRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteSponsoredBrandsAdGroupsBetaRequestContent{}

// DeleteSponsoredBrandsAdGroupsBetaRequestContent struct for DeleteSponsoredBrandsAdGroupsBetaRequestContent
type DeleteSponsoredBrandsAdGroupsBetaRequestContent struct {
	AdGroupIdFilter *ObjectIdFilter `json:"adGroupIdFilter,omitempty"`
}

// NewDeleteSponsoredBrandsAdGroupsBetaRequestContent instantiates a new DeleteSponsoredBrandsAdGroupsBetaRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteSponsoredBrandsAdGroupsBetaRequestContent() *DeleteSponsoredBrandsAdGroupsBetaRequestContent {
	this := DeleteSponsoredBrandsAdGroupsBetaRequestContent{}
	return &this
}

// NewDeleteSponsoredBrandsAdGroupsBetaRequestContentWithDefaults instantiates a new DeleteSponsoredBrandsAdGroupsBetaRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteSponsoredBrandsAdGroupsBetaRequestContentWithDefaults() *DeleteSponsoredBrandsAdGroupsBetaRequestContent {
	this := DeleteSponsoredBrandsAdGroupsBetaRequestContent{}
	return &this
}

// GetAdGroupIdFilter returns the AdGroupIdFilter field value if set, zero value otherwise.
func (o *DeleteSponsoredBrandsAdGroupsBetaRequestContent) GetAdGroupIdFilter() ObjectIdFilter {
	if o == nil || IsNil(o.AdGroupIdFilter) {
		var ret ObjectIdFilter
		return ret
	}
	return *o.AdGroupIdFilter
}

// GetAdGroupIdFilterOk returns a tuple with the AdGroupIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteSponsoredBrandsAdGroupsBetaRequestContent) GetAdGroupIdFilterOk() (*ObjectIdFilter, bool) {
	if o == nil || IsNil(o.AdGroupIdFilter) {
		return nil, false
	}
	return o.AdGroupIdFilter, true
}

// HasAdGroupIdFilter returns a boolean if a field has been set.
func (o *DeleteSponsoredBrandsAdGroupsBetaRequestContent) HasAdGroupIdFilter() bool {
	if o != nil && !IsNil(o.AdGroupIdFilter) {
		return true
	}

	return false
}

// SetAdGroupIdFilter gets a reference to the given ObjectIdFilter and assigns it to the AdGroupIdFilter field.
func (o *DeleteSponsoredBrandsAdGroupsBetaRequestContent) SetAdGroupIdFilter(v ObjectIdFilter) {
	o.AdGroupIdFilter = &v
}

func (o DeleteSponsoredBrandsAdGroupsBetaRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdGroupIdFilter) {
		toSerialize["adGroupIdFilter"] = o.AdGroupIdFilter
	}
	return toSerialize, nil
}

type NullableDeleteSponsoredBrandsAdGroupsBetaRequestContent struct {
	value *DeleteSponsoredBrandsAdGroupsBetaRequestContent
	isSet bool
}

func (v NullableDeleteSponsoredBrandsAdGroupsBetaRequestContent) Get() *DeleteSponsoredBrandsAdGroupsBetaRequestContent {
	return v.value
}

func (v *NullableDeleteSponsoredBrandsAdGroupsBetaRequestContent) Set(val *DeleteSponsoredBrandsAdGroupsBetaRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteSponsoredBrandsAdGroupsBetaRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteSponsoredBrandsAdGroupsBetaRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteSponsoredBrandsAdGroupsBetaRequestContent(val *DeleteSponsoredBrandsAdGroupsBetaRequestContent) *NullableDeleteSponsoredBrandsAdGroupsBetaRequestContent {
	return &NullableDeleteSponsoredBrandsAdGroupsBetaRequestContent{value: val, isSet: true}
}

func (v NullableDeleteSponsoredBrandsAdGroupsBetaRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDeleteSponsoredBrandsAdGroupsBetaRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
