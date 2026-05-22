package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the ProductRecommendationsByTheme type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductRecommendationsByTheme{}

// ProductRecommendationsByTheme Product recommendations grouped by theme attribute.
type ProductRecommendationsByTheme struct {
	// An identifier to fetch next set of `ThemeRecommendation` records in the result set if available. This will be null when at the end of result set.
	NextCursor *string `json:"nextCursor,omitempty"`
	// Optional parameter that links to the previous result set served to the requester.
	PreviousCursor *string `json:"previousCursor,omitempty"`
	// An array of `ThemeRecommendation` objects
	Recommendations []ThemeRecommendation `json:"recommendations,omitempty"`
}

// NewProductRecommendationsByTheme instantiates a new ProductRecommendationsByTheme object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductRecommendationsByTheme() *ProductRecommendationsByTheme {
	this := ProductRecommendationsByTheme{}
	return &this
}

// NewProductRecommendationsByThemeWithDefaults instantiates a new ProductRecommendationsByTheme object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductRecommendationsByThemeWithDefaults() *ProductRecommendationsByTheme {
	this := ProductRecommendationsByTheme{}
	return &this
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise.
func (o *ProductRecommendationsByTheme) GetNextCursor() string {
	if o == nil || IsNil(o.NextCursor) {
		var ret string
		return ret
	}
	return *o.NextCursor
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductRecommendationsByTheme) GetNextCursorOk() (*string, bool) {
	if o == nil || IsNil(o.NextCursor) {
		return nil, false
	}
	return o.NextCursor, true
}

// HasNextCursor returns a boolean if a field has been set.
func (o *ProductRecommendationsByTheme) HasNextCursor() bool {
	if o != nil && !IsNil(o.NextCursor) {
		return true
	}

	return false
}

// SetNextCursor gets a reference to the given string and assigns it to the NextCursor field.
func (o *ProductRecommendationsByTheme) SetNextCursor(v string) {
	o.NextCursor = &v
}

// GetPreviousCursor returns the PreviousCursor field value if set, zero value otherwise.
func (o *ProductRecommendationsByTheme) GetPreviousCursor() string {
	if o == nil || IsNil(o.PreviousCursor) {
		var ret string
		return ret
	}
	return *o.PreviousCursor
}

// GetPreviousCursorOk returns a tuple with the PreviousCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductRecommendationsByTheme) GetPreviousCursorOk() (*string, bool) {
	if o == nil || IsNil(o.PreviousCursor) {
		return nil, false
	}
	return o.PreviousCursor, true
}

// HasPreviousCursor returns a boolean if a field has been set.
func (o *ProductRecommendationsByTheme) HasPreviousCursor() bool {
	if o != nil && !IsNil(o.PreviousCursor) {
		return true
	}

	return false
}

// SetPreviousCursor gets a reference to the given string and assigns it to the PreviousCursor field.
func (o *ProductRecommendationsByTheme) SetPreviousCursor(v string) {
	o.PreviousCursor = &v
}

// GetRecommendations returns the Recommendations field value if set, zero value otherwise.
func (o *ProductRecommendationsByTheme) GetRecommendations() []ThemeRecommendation {
	if o == nil || IsNil(o.Recommendations) {
		var ret []ThemeRecommendation
		return ret
	}
	return o.Recommendations
}

// GetRecommendationsOk returns a tuple with the Recommendations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductRecommendationsByTheme) GetRecommendationsOk() ([]ThemeRecommendation, bool) {
	if o == nil || IsNil(o.Recommendations) {
		return nil, false
	}
	return o.Recommendations, true
}

// HasRecommendations returns a boolean if a field has been set.
func (o *ProductRecommendationsByTheme) HasRecommendations() bool {
	if o != nil && !IsNil(o.Recommendations) {
		return true
	}

	return false
}

// SetRecommendations gets a reference to the given []ThemeRecommendation and assigns it to the Recommendations field.
func (o *ProductRecommendationsByTheme) SetRecommendations(v []ThemeRecommendation) {
	o.Recommendations = v
}

func (o ProductRecommendationsByTheme) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextCursor) {
		toSerialize["nextCursor"] = o.NextCursor
	}
	if !IsNil(o.PreviousCursor) {
		toSerialize["previousCursor"] = o.PreviousCursor
	}
	if !IsNil(o.Recommendations) {
		toSerialize["recommendations"] = o.Recommendations
	}
	return toSerialize, nil
}

type NullableProductRecommendationsByTheme struct {
	value *ProductRecommendationsByTheme
	isSet bool
}

func (v NullableProductRecommendationsByTheme) Get() *ProductRecommendationsByTheme {
	return v.value
}

func (v *NullableProductRecommendationsByTheme) Set(val *ProductRecommendationsByTheme) {
	v.value = val
	v.isSet = true
}

func (v NullableProductRecommendationsByTheme) IsSet() bool {
	return v.isSet
}

func (v *NullableProductRecommendationsByTheme) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductRecommendationsByTheme(val *ProductRecommendationsByTheme) *NullableProductRecommendationsByTheme {
	return &NullableProductRecommendationsByTheme{value: val, isSet: true}
}

func (v NullableProductRecommendationsByTheme) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableProductRecommendationsByTheme) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
