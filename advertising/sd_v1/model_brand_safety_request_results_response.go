package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the BrandSafetyRequestResultsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BrandSafetyRequestResultsResponse{}

// BrandSafetyRequestResultsResponse struct for BrandSafetyRequestResultsResponse
type BrandSafetyRequestResultsResponse struct {
	// A list of results for the given requestId
	Results []BrandSafetyRequestResult `json:"results,omitempty"`
}

// NewBrandSafetyRequestResultsResponse instantiates a new BrandSafetyRequestResultsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrandSafetyRequestResultsResponse() *BrandSafetyRequestResultsResponse {
	this := BrandSafetyRequestResultsResponse{}
	return &this
}

// NewBrandSafetyRequestResultsResponseWithDefaults instantiates a new BrandSafetyRequestResultsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandSafetyRequestResultsResponseWithDefaults() *BrandSafetyRequestResultsResponse {
	this := BrandSafetyRequestResultsResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *BrandSafetyRequestResultsResponse) GetResults() []BrandSafetyRequestResult {
	if o == nil || IsNil(o.Results) {
		var ret []BrandSafetyRequestResult
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandSafetyRequestResultsResponse) GetResultsOk() ([]BrandSafetyRequestResult, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *BrandSafetyRequestResultsResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []BrandSafetyRequestResult and assigns it to the Results field.
func (o *BrandSafetyRequestResultsResponse) SetResults(v []BrandSafetyRequestResult) {
	o.Results = v
}

func (o BrandSafetyRequestResultsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableBrandSafetyRequestResultsResponse struct {
	value *BrandSafetyRequestResultsResponse
	isSet bool
}

func (v NullableBrandSafetyRequestResultsResponse) Get() *BrandSafetyRequestResultsResponse {
	return v.value
}

func (v *NullableBrandSafetyRequestResultsResponse) Set(val *BrandSafetyRequestResultsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandSafetyRequestResultsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandSafetyRequestResultsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandSafetyRequestResultsResponse(val *BrandSafetyRequestResultsResponse) *NullableBrandSafetyRequestResultsResponse {
	return &NullableBrandSafetyRequestResultsResponse{value: val, isSet: true}
}

func (v NullableBrandSafetyRequestResultsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBrandSafetyRequestResultsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
