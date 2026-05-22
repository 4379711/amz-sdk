package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the GetCategoryRecommendationsForAsinsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCategoryRecommendationsForAsinsRequest{}

// GetCategoryRecommendationsForAsinsRequest Request object to retrieve Category Recommendations based on the input ASINs.
type GetCategoryRecommendationsForAsinsRequest struct {
	// List of input ASINs. This API does not check if the ASINs are valid ASINs.
	Asins []string `json:"asins,omitempty"`
	// Enable this if you would like to retrieve categories which are ancestor nodes of the original recommended categories. This may increase the number of categories returned, but decrease the relevancy of those categories.
	IncludeAncestor *bool `json:"includeAncestor,omitempty"`
}

// NewGetCategoryRecommendationsForAsinsRequest instantiates a new GetCategoryRecommendationsForAsinsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCategoryRecommendationsForAsinsRequest() *GetCategoryRecommendationsForAsinsRequest {
	this := GetCategoryRecommendationsForAsinsRequest{}
	return &this
}

// NewGetCategoryRecommendationsForAsinsRequestWithDefaults instantiates a new GetCategoryRecommendationsForAsinsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCategoryRecommendationsForAsinsRequestWithDefaults() *GetCategoryRecommendationsForAsinsRequest {
	this := GetCategoryRecommendationsForAsinsRequest{}
	return &this
}

// GetAsins returns the Asins field value if set, zero value otherwise.
func (o *GetCategoryRecommendationsForAsinsRequest) GetAsins() []string {
	if o == nil || IsNil(o.Asins) {
		var ret []string
		return ret
	}
	return o.Asins
}

// GetAsinsOk returns a tuple with the Asins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCategoryRecommendationsForAsinsRequest) GetAsinsOk() ([]string, bool) {
	if o == nil || IsNil(o.Asins) {
		return nil, false
	}
	return o.Asins, true
}

// HasAsins returns a boolean if a field has been set.
func (o *GetCategoryRecommendationsForAsinsRequest) HasAsins() bool {
	if o != nil && !IsNil(o.Asins) {
		return true
	}

	return false
}

// SetAsins gets a reference to the given []string and assigns it to the Asins field.
func (o *GetCategoryRecommendationsForAsinsRequest) SetAsins(v []string) {
	o.Asins = v
}

// GetIncludeAncestor returns the IncludeAncestor field value if set, zero value otherwise.
func (o *GetCategoryRecommendationsForAsinsRequest) GetIncludeAncestor() bool {
	if o == nil || IsNil(o.IncludeAncestor) {
		var ret bool
		return ret
	}
	return *o.IncludeAncestor
}

// GetIncludeAncestorOk returns a tuple with the IncludeAncestor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCategoryRecommendationsForAsinsRequest) GetIncludeAncestorOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeAncestor) {
		return nil, false
	}
	return o.IncludeAncestor, true
}

// HasIncludeAncestor returns a boolean if a field has been set.
func (o *GetCategoryRecommendationsForAsinsRequest) HasIncludeAncestor() bool {
	if o != nil && !IsNil(o.IncludeAncestor) {
		return true
	}

	return false
}

// SetIncludeAncestor gets a reference to the given bool and assigns it to the IncludeAncestor field.
func (o *GetCategoryRecommendationsForAsinsRequest) SetIncludeAncestor(v bool) {
	o.IncludeAncestor = &v
}

func (o GetCategoryRecommendationsForAsinsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Asins) {
		toSerialize["asins"] = o.Asins
	}
	if !IsNil(o.IncludeAncestor) {
		toSerialize["includeAncestor"] = o.IncludeAncestor
	}
	return toSerialize, nil
}

type NullableGetCategoryRecommendationsForAsinsRequest struct {
	value *GetCategoryRecommendationsForAsinsRequest
	isSet bool
}

func (v NullableGetCategoryRecommendationsForAsinsRequest) Get() *GetCategoryRecommendationsForAsinsRequest {
	return v.value
}

func (v *NullableGetCategoryRecommendationsForAsinsRequest) Set(val *GetCategoryRecommendationsForAsinsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCategoryRecommendationsForAsinsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCategoryRecommendationsForAsinsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCategoryRecommendationsForAsinsRequest(val *GetCategoryRecommendationsForAsinsRequest) *NullableGetCategoryRecommendationsForAsinsRequest {
	return &NullableGetCategoryRecommendationsForAsinsRequest{value: val, isSet: true}
}

func (v NullableGetCategoryRecommendationsForAsinsRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetCategoryRecommendationsForAsinsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
