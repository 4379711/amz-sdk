package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the DeleteSponsoredBrandsAdsBetaRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteSponsoredBrandsAdsBetaRequestContent{}

// DeleteSponsoredBrandsAdsBetaRequestContent struct for DeleteSponsoredBrandsAdsBetaRequestContent
type DeleteSponsoredBrandsAdsBetaRequestContent struct {
	AdIdFilter *ObjectIdFilter `json:"adIdFilter,omitempty"`
}

// NewDeleteSponsoredBrandsAdsBetaRequestContent instantiates a new DeleteSponsoredBrandsAdsBetaRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteSponsoredBrandsAdsBetaRequestContent() *DeleteSponsoredBrandsAdsBetaRequestContent {
	this := DeleteSponsoredBrandsAdsBetaRequestContent{}
	return &this
}

// NewDeleteSponsoredBrandsAdsBetaRequestContentWithDefaults instantiates a new DeleteSponsoredBrandsAdsBetaRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteSponsoredBrandsAdsBetaRequestContentWithDefaults() *DeleteSponsoredBrandsAdsBetaRequestContent {
	this := DeleteSponsoredBrandsAdsBetaRequestContent{}
	return &this
}

// GetAdIdFilter returns the AdIdFilter field value if set, zero value otherwise.
func (o *DeleteSponsoredBrandsAdsBetaRequestContent) GetAdIdFilter() ObjectIdFilter {
	if o == nil || IsNil(o.AdIdFilter) {
		var ret ObjectIdFilter
		return ret
	}
	return *o.AdIdFilter
}

// GetAdIdFilterOk returns a tuple with the AdIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteSponsoredBrandsAdsBetaRequestContent) GetAdIdFilterOk() (*ObjectIdFilter, bool) {
	if o == nil || IsNil(o.AdIdFilter) {
		return nil, false
	}
	return o.AdIdFilter, true
}

// HasAdIdFilter returns a boolean if a field has been set.
func (o *DeleteSponsoredBrandsAdsBetaRequestContent) HasAdIdFilter() bool {
	if o != nil && !IsNil(o.AdIdFilter) {
		return true
	}

	return false
}

// SetAdIdFilter gets a reference to the given ObjectIdFilter and assigns it to the AdIdFilter field.
func (o *DeleteSponsoredBrandsAdsBetaRequestContent) SetAdIdFilter(v ObjectIdFilter) {
	o.AdIdFilter = &v
}

func (o DeleteSponsoredBrandsAdsBetaRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdIdFilter) {
		toSerialize["adIdFilter"] = o.AdIdFilter
	}
	return toSerialize, nil
}

type NullableDeleteSponsoredBrandsAdsBetaRequestContent struct {
	value *DeleteSponsoredBrandsAdsBetaRequestContent
	isSet bool
}

func (v NullableDeleteSponsoredBrandsAdsBetaRequestContent) Get() *DeleteSponsoredBrandsAdsBetaRequestContent {
	return v.value
}

func (v *NullableDeleteSponsoredBrandsAdsBetaRequestContent) Set(val *DeleteSponsoredBrandsAdsBetaRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteSponsoredBrandsAdsBetaRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteSponsoredBrandsAdsBetaRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteSponsoredBrandsAdsBetaRequestContent(val *DeleteSponsoredBrandsAdsBetaRequestContent) *NullableDeleteSponsoredBrandsAdsBetaRequestContent {
	return &NullableDeleteSponsoredBrandsAdsBetaRequestContent{value: val, isSet: true}
}

func (v NullableDeleteSponsoredBrandsAdsBetaRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDeleteSponsoredBrandsAdsBetaRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
