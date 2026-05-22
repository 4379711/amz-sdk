package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the DeleteSponsoredBrandsAdGroupsRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteSponsoredBrandsAdGroupsRequestContent{}

// DeleteSponsoredBrandsAdGroupsRequestContent struct for DeleteSponsoredBrandsAdGroupsRequestContent
type DeleteSponsoredBrandsAdGroupsRequestContent struct {
	AdGroupIdFilter *ObjectIdFilter `json:"adGroupIdFilter,omitempty"`
}

// NewDeleteSponsoredBrandsAdGroupsRequestContent instantiates a new DeleteSponsoredBrandsAdGroupsRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteSponsoredBrandsAdGroupsRequestContent() *DeleteSponsoredBrandsAdGroupsRequestContent {
	this := DeleteSponsoredBrandsAdGroupsRequestContent{}
	return &this
}

// NewDeleteSponsoredBrandsAdGroupsRequestContentWithDefaults instantiates a new DeleteSponsoredBrandsAdGroupsRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteSponsoredBrandsAdGroupsRequestContentWithDefaults() *DeleteSponsoredBrandsAdGroupsRequestContent {
	this := DeleteSponsoredBrandsAdGroupsRequestContent{}
	return &this
}

// GetAdGroupIdFilter returns the AdGroupIdFilter field value if set, zero value otherwise.
func (o *DeleteSponsoredBrandsAdGroupsRequestContent) GetAdGroupIdFilter() ObjectIdFilter {
	if o == nil || IsNil(o.AdGroupIdFilter) {
		var ret ObjectIdFilter
		return ret
	}
	return *o.AdGroupIdFilter
}

// GetAdGroupIdFilterOk returns a tuple with the AdGroupIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteSponsoredBrandsAdGroupsRequestContent) GetAdGroupIdFilterOk() (*ObjectIdFilter, bool) {
	if o == nil || IsNil(o.AdGroupIdFilter) {
		return nil, false
	}
	return o.AdGroupIdFilter, true
}

// HasAdGroupIdFilter returns a boolean if a field has been set.
func (o *DeleteSponsoredBrandsAdGroupsRequestContent) HasAdGroupIdFilter() bool {
	if o != nil && !IsNil(o.AdGroupIdFilter) {
		return true
	}

	return false
}

// SetAdGroupIdFilter gets a reference to the given ObjectIdFilter and assigns it to the AdGroupIdFilter field.
func (o *DeleteSponsoredBrandsAdGroupsRequestContent) SetAdGroupIdFilter(v ObjectIdFilter) {
	o.AdGroupIdFilter = &v
}

func (o DeleteSponsoredBrandsAdGroupsRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdGroupIdFilter) {
		toSerialize["adGroupIdFilter"] = o.AdGroupIdFilter
	}
	return toSerialize, nil
}

type NullableDeleteSponsoredBrandsAdGroupsRequestContent struct {
	value *DeleteSponsoredBrandsAdGroupsRequestContent
	isSet bool
}

func (v NullableDeleteSponsoredBrandsAdGroupsRequestContent) Get() *DeleteSponsoredBrandsAdGroupsRequestContent {
	return v.value
}

func (v *NullableDeleteSponsoredBrandsAdGroupsRequestContent) Set(val *DeleteSponsoredBrandsAdGroupsRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteSponsoredBrandsAdGroupsRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteSponsoredBrandsAdGroupsRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteSponsoredBrandsAdGroupsRequestContent(val *DeleteSponsoredBrandsAdGroupsRequestContent) *NullableDeleteSponsoredBrandsAdGroupsRequestContent {
	return &NullableDeleteSponsoredBrandsAdGroupsRequestContent{value: val, isSet: true}
}

func (v NullableDeleteSponsoredBrandsAdGroupsRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDeleteSponsoredBrandsAdGroupsRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
