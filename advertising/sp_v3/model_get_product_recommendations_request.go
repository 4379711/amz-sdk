package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the GetProductRecommendationsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProductRecommendationsRequest{}

// GetProductRecommendationsRequest Request structure to get ASIN recommendations for a set of input ASINs.
type GetProductRecommendationsRequest struct {
	// A optional cursor value that can be used to fetch next or previous set of records.
	Cursor *string `json:"cursor,omitempty"`
	// List of input ASINs.
	AdAsins []string `json:"adAsins"`
	// Count of objects requested in the response. The count will be applied on the objects returned under `recommendations` array in response body.  <ul> <li>Requesting `application/vnd.spproductrecommendationresponse.themes.v3+json` mediatype applies the count on `ThemeRecommendation` objects.If no count value is passed a default of `5` is assumed. The API will return a maximum of `10` themes irrespective of how large the count is set. </li> <li>Requesting `application/vnd.spproductrecommendationresponse.asins.v3+json` mediatype applies count on the `ProductRecommendation` objects in response body.If no count value is passed a default of `100` is assumed. The API will return a maximum of `1000` recommendations irrespective of how large the count is set. </li> </ul> Please refer the response Schemas for more info.
	Count *int32 `json:"count,omitempty"`
	// Theme names and descriptions will be provided in the language for your supported locale. Available options are en_US (U.S. English), en_GB (UK English), zh_CN (Chinese), es_ES (Spanish), jp_JP (Japanese), de_DE (German), fr_FR (French), it_IT(Italian). If locale is not provided or unsupported, the theme names and descriptions will be returned in U.S. English (en_US).
	Locale *string `json:"locale,omitempty"`
}

type _GetProductRecommendationsRequest GetProductRecommendationsRequest

// NewGetProductRecommendationsRequest instantiates a new GetProductRecommendationsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProductRecommendationsRequest(adAsins []string) *GetProductRecommendationsRequest {
	this := GetProductRecommendationsRequest{}
	this.AdAsins = adAsins
	return &this
}

// NewGetProductRecommendationsRequestWithDefaults instantiates a new GetProductRecommendationsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProductRecommendationsRequestWithDefaults() *GetProductRecommendationsRequest {
	this := GetProductRecommendationsRequest{}
	return &this
}

// GetCursor returns the Cursor field value if set, zero value otherwise.
func (o *GetProductRecommendationsRequest) GetCursor() string {
	if o == nil || IsNil(o.Cursor) {
		var ret string
		return ret
	}
	return *o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductRecommendationsRequest) GetCursorOk() (*string, bool) {
	if o == nil || IsNil(o.Cursor) {
		return nil, false
	}
	return o.Cursor, true
}

// HasCursor returns a boolean if a field has been set.
func (o *GetProductRecommendationsRequest) HasCursor() bool {
	if o != nil && !IsNil(o.Cursor) {
		return true
	}

	return false
}

// SetCursor gets a reference to the given string and assigns it to the Cursor field.
func (o *GetProductRecommendationsRequest) SetCursor(v string) {
	o.Cursor = &v
}

// GetAdAsins returns the AdAsins field value
func (o *GetProductRecommendationsRequest) GetAdAsins() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AdAsins
}

// GetAdAsinsOk returns a tuple with the AdAsins field value
// and a boolean to check if the value has been set.
func (o *GetProductRecommendationsRequest) GetAdAsinsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdAsins, true
}

// SetAdAsins sets field value
func (o *GetProductRecommendationsRequest) SetAdAsins(v []string) {
	o.AdAsins = v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *GetProductRecommendationsRequest) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductRecommendationsRequest) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *GetProductRecommendationsRequest) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *GetProductRecommendationsRequest) SetCount(v int32) {
	o.Count = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *GetProductRecommendationsRequest) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductRecommendationsRequest) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *GetProductRecommendationsRequest) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *GetProductRecommendationsRequest) SetLocale(v string) {
	o.Locale = &v
}

func (o GetProductRecommendationsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cursor) {
		toSerialize["cursor"] = o.Cursor
	}
	toSerialize["adAsins"] = o.AdAsins
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	return toSerialize, nil
}

type NullableGetProductRecommendationsRequest struct {
	value *GetProductRecommendationsRequest
	isSet bool
}

func (v NullableGetProductRecommendationsRequest) Get() *GetProductRecommendationsRequest {
	return v.value
}

func (v *NullableGetProductRecommendationsRequest) Set(val *GetProductRecommendationsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProductRecommendationsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProductRecommendationsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProductRecommendationsRequest(val *GetProductRecommendationsRequest) *NullableGetProductRecommendationsRequest {
	return &NullableGetProductRecommendationsRequest{value: val, isSet: true}
}

func (v NullableGetProductRecommendationsRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetProductRecommendationsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
