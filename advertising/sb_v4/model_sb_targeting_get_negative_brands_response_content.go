package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBTargetingGetNegativeBrandsResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBTargetingGetNegativeBrandsResponseContent{}

// SBTargetingGetNegativeBrandsResponseContent Response object for /sb/negativeTargets/brands/recommendations containing list of brands for negative targeting.
type SBTargetingGetNegativeBrandsResponseContent struct {
	// List of Brands.
	Brands []SBTargetingBrand `json:"brands,omitempty"`
	// Operations that return paginated results include a pagination token in this field. To retrieve the next page of results, call the same operation and specify this token in the request. If the `NextToken` field is empty, there are no further results.
	NextToken *string `json:"nextToken,omitempty"`
}

// NewSBTargetingGetNegativeBrandsResponseContent instantiates a new SBTargetingGetNegativeBrandsResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBTargetingGetNegativeBrandsResponseContent() *SBTargetingGetNegativeBrandsResponseContent {
	this := SBTargetingGetNegativeBrandsResponseContent{}
	return &this
}

// NewSBTargetingGetNegativeBrandsResponseContentWithDefaults instantiates a new SBTargetingGetNegativeBrandsResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBTargetingGetNegativeBrandsResponseContentWithDefaults() *SBTargetingGetNegativeBrandsResponseContent {
	this := SBTargetingGetNegativeBrandsResponseContent{}
	return &this
}

// GetBrands returns the Brands field value if set, zero value otherwise.
func (o *SBTargetingGetNegativeBrandsResponseContent) GetBrands() []SBTargetingBrand {
	if o == nil || IsNil(o.Brands) {
		var ret []SBTargetingBrand
		return ret
	}
	return o.Brands
}

// GetBrandsOk returns a tuple with the Brands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBTargetingGetNegativeBrandsResponseContent) GetBrandsOk() ([]SBTargetingBrand, bool) {
	if o == nil || IsNil(o.Brands) {
		return nil, false
	}
	return o.Brands, true
}

// HasBrands returns a boolean if a field has been set.
func (o *SBTargetingGetNegativeBrandsResponseContent) HasBrands() bool {
	if o != nil && !IsNil(o.Brands) {
		return true
	}

	return false
}

// SetBrands gets a reference to the given []SBTargetingBrand and assigns it to the Brands field.
func (o *SBTargetingGetNegativeBrandsResponseContent) SetBrands(v []SBTargetingBrand) {
	o.Brands = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SBTargetingGetNegativeBrandsResponseContent) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBTargetingGetNegativeBrandsResponseContent) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SBTargetingGetNegativeBrandsResponseContent) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SBTargetingGetNegativeBrandsResponseContent) SetNextToken(v string) {
	o.NextToken = &v
}

func (o SBTargetingGetNegativeBrandsResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Brands) {
		toSerialize["brands"] = o.Brands
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	return toSerialize, nil
}

type NullableSBTargetingGetNegativeBrandsResponseContent struct {
	value *SBTargetingGetNegativeBrandsResponseContent
	isSet bool
}

func (v NullableSBTargetingGetNegativeBrandsResponseContent) Get() *SBTargetingGetNegativeBrandsResponseContent {
	return v.value
}

func (v *NullableSBTargetingGetNegativeBrandsResponseContent) Set(val *SBTargetingGetNegativeBrandsResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableSBTargetingGetNegativeBrandsResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableSBTargetingGetNegativeBrandsResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBTargetingGetNegativeBrandsResponseContent(val *SBTargetingGetNegativeBrandsResponseContent) *NullableSBTargetingGetNegativeBrandsResponseContent {
	return &NullableSBTargetingGetNegativeBrandsResponseContent{value: val, isSet: true}
}

func (v NullableSBTargetingGetNegativeBrandsResponseContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBTargetingGetNegativeBrandsResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
