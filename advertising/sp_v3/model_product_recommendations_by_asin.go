package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the ProductRecommendationsByASIN type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductRecommendationsByASIN{}

// ProductRecommendationsByASIN Product recommendations supplemented with relevant information.
type ProductRecommendationsByASIN struct {
	// An identifier to fetch next set of `ProductRecommendation` records in the result set if available. This will be null when at the end of result set.
	NextCursor *string `json:"nextCursor,omitempty"`
	// Optional parameter that links to the previous result set served. This parameter will be null on the first request.
	PreviousCursor *string `json:"previousCursor,omitempty"`
	// An array of `ProductRecommendation` objects.
	Recommendations []ProductRecommendation `json:"recommendations,omitempty"`
}

// NewProductRecommendationsByASIN instantiates a new ProductRecommendationsByASIN object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductRecommendationsByASIN() *ProductRecommendationsByASIN {
	this := ProductRecommendationsByASIN{}
	return &this
}

// NewProductRecommendationsByASINWithDefaults instantiates a new ProductRecommendationsByASIN object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductRecommendationsByASINWithDefaults() *ProductRecommendationsByASIN {
	this := ProductRecommendationsByASIN{}
	return &this
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise.
func (o *ProductRecommendationsByASIN) GetNextCursor() string {
	if o == nil || IsNil(o.NextCursor) {
		var ret string
		return ret
	}
	return *o.NextCursor
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductRecommendationsByASIN) GetNextCursorOk() (*string, bool) {
	if o == nil || IsNil(o.NextCursor) {
		return nil, false
	}
	return o.NextCursor, true
}

// HasNextCursor returns a boolean if a field has been set.
func (o *ProductRecommendationsByASIN) HasNextCursor() bool {
	if o != nil && !IsNil(o.NextCursor) {
		return true
	}

	return false
}

// SetNextCursor gets a reference to the given string and assigns it to the NextCursor field.
func (o *ProductRecommendationsByASIN) SetNextCursor(v string) {
	o.NextCursor = &v
}

// GetPreviousCursor returns the PreviousCursor field value if set, zero value otherwise.
func (o *ProductRecommendationsByASIN) GetPreviousCursor() string {
	if o == nil || IsNil(o.PreviousCursor) {
		var ret string
		return ret
	}
	return *o.PreviousCursor
}

// GetPreviousCursorOk returns a tuple with the PreviousCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductRecommendationsByASIN) GetPreviousCursorOk() (*string, bool) {
	if o == nil || IsNil(o.PreviousCursor) {
		return nil, false
	}
	return o.PreviousCursor, true
}

// HasPreviousCursor returns a boolean if a field has been set.
func (o *ProductRecommendationsByASIN) HasPreviousCursor() bool {
	if o != nil && !IsNil(o.PreviousCursor) {
		return true
	}

	return false
}

// SetPreviousCursor gets a reference to the given string and assigns it to the PreviousCursor field.
func (o *ProductRecommendationsByASIN) SetPreviousCursor(v string) {
	o.PreviousCursor = &v
}

// GetRecommendations returns the Recommendations field value if set, zero value otherwise.
func (o *ProductRecommendationsByASIN) GetRecommendations() []ProductRecommendation {
	if o == nil || IsNil(o.Recommendations) {
		var ret []ProductRecommendation
		return ret
	}
	return o.Recommendations
}

// GetRecommendationsOk returns a tuple with the Recommendations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductRecommendationsByASIN) GetRecommendationsOk() ([]ProductRecommendation, bool) {
	if o == nil || IsNil(o.Recommendations) {
		return nil, false
	}
	return o.Recommendations, true
}

// HasRecommendations returns a boolean if a field has been set.
func (o *ProductRecommendationsByASIN) HasRecommendations() bool {
	if o != nil && !IsNil(o.Recommendations) {
		return true
	}

	return false
}

// SetRecommendations gets a reference to the given []ProductRecommendation and assigns it to the Recommendations field.
func (o *ProductRecommendationsByASIN) SetRecommendations(v []ProductRecommendation) {
	o.Recommendations = v
}

func (o ProductRecommendationsByASIN) ToMap() (map[string]interface{}, error) {
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

type NullableProductRecommendationsByASIN struct {
	value *ProductRecommendationsByASIN
	isSet bool
}

func (v NullableProductRecommendationsByASIN) Get() *ProductRecommendationsByASIN {
	return v.value
}

func (v *NullableProductRecommendationsByASIN) Set(val *ProductRecommendationsByASIN) {
	v.value = val
	v.isSet = true
}

func (v NullableProductRecommendationsByASIN) IsSet() bool {
	return v.isSet
}

func (v *NullableProductRecommendationsByASIN) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductRecommendationsByASIN(val *ProductRecommendationsByASIN) *NullableProductRecommendationsByASIN {
	return &NullableProductRecommendationsByASIN{value: val, isSet: true}
}

func (v NullableProductRecommendationsByASIN) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableProductRecommendationsByASIN) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
