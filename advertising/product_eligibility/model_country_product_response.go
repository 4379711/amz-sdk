package product_eligibility

import (
	"github.com/bytedance/sonic"
)

// checks if the CountryProductResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CountryProductResponse{}

// CountryProductResponse A country and the county's corresponding product identifiers list with advertising eligibility response.
type CountryProductResponse struct {
	// A list of product advertising eligibility responses.
	ProductResponseList []ProductResponse `json:"productResponseList"`
	// A country code in the ISO 3166-1 alpha-2 format. Eligibility in this country will be checked for the corresponding list of product identifiers. |Country|Country ISO 3166 Code| |-------|---------------------| |Australia|AU| |Belgium|BE| |Brazil|BR| |Canada|CA| |Chile|CL| |Colombia|CO| |Egypt|EG| |France|FR| |Germany|DE| |India|IN| |Ireland|IE| |Italy|IT| |Japan|JP| |Mexico|MX| |Netherlands|NL| |Nigeria|NG| |Poland|PL| |Russia|RU| |Saudi Arabia|SA| |Singapore|SG| |South Africa|ZA| |South Korea|KR| |Spain|ES| |Sweden|SE| |Turkey|TR| |United Arab Emirates|AE| |United Kingdom|GB| |United States|US|
	CountryCode    string          `json:"countryCode"`
	CountryFailure *CountryFailure `json:"countryFailure,omitempty"`
}

type _CountryProductResponse CountryProductResponse

// NewCountryProductResponse instantiates a new CountryProductResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountryProductResponse(productResponseList []ProductResponse, countryCode string) *CountryProductResponse {
	this := CountryProductResponse{}
	this.ProductResponseList = productResponseList
	this.CountryCode = countryCode
	return &this
}

// NewCountryProductResponseWithDefaults instantiates a new CountryProductResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountryProductResponseWithDefaults() *CountryProductResponse {
	this := CountryProductResponse{}
	return &this
}

// GetProductResponseList returns the ProductResponseList field value
func (o *CountryProductResponse) GetProductResponseList() []ProductResponse {
	if o == nil {
		var ret []ProductResponse
		return ret
	}

	return o.ProductResponseList
}

// GetProductResponseListOk returns a tuple with the ProductResponseList field value
// and a boolean to check if the value has been set.
func (o *CountryProductResponse) GetProductResponseListOk() ([]ProductResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductResponseList, true
}

// SetProductResponseList sets field value
func (o *CountryProductResponse) SetProductResponseList(v []ProductResponse) {
	o.ProductResponseList = v
}

// GetCountryCode returns the CountryCode field value
func (o *CountryProductResponse) GetCountryCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value
// and a boolean to check if the value has been set.
func (o *CountryProductResponse) GetCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryCode, true
}

// SetCountryCode sets field value
func (o *CountryProductResponse) SetCountryCode(v string) {
	o.CountryCode = v
}

// GetCountryFailure returns the CountryFailure field value if set, zero value otherwise.
func (o *CountryProductResponse) GetCountryFailure() CountryFailure {
	if o == nil || IsNil(o.CountryFailure) {
		var ret CountryFailure
		return ret
	}
	return *o.CountryFailure
}

// GetCountryFailureOk returns a tuple with the CountryFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryProductResponse) GetCountryFailureOk() (*CountryFailure, bool) {
	if o == nil || IsNil(o.CountryFailure) {
		return nil, false
	}
	return o.CountryFailure, true
}

// HasCountryFailure returns a boolean if a field has been set.
func (o *CountryProductResponse) HasCountryFailure() bool {
	if o != nil && !IsNil(o.CountryFailure) {
		return true
	}

	return false
}

// SetCountryFailure gets a reference to the given CountryFailure and assigns it to the CountryFailure field.
func (o *CountryProductResponse) SetCountryFailure(v CountryFailure) {
	o.CountryFailure = &v
}

func (o CountryProductResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["productResponseList"] = o.ProductResponseList
	toSerialize["countryCode"] = o.CountryCode
	if !IsNil(o.CountryFailure) {
		toSerialize["countryFailure"] = o.CountryFailure
	}
	return toSerialize, nil
}

type NullableCountryProductResponse struct {
	value *CountryProductResponse
	isSet bool
}

func (v NullableCountryProductResponse) Get() *CountryProductResponse {
	return v.value
}

func (v *NullableCountryProductResponse) Set(val *CountryProductResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCountryProductResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCountryProductResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountryProductResponse(val *CountryProductResponse) *NullableCountryProductResponse {
	return &NullableCountryProductResponse{value: val, isSet: true}
}

func (v NullableCountryProductResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCountryProductResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
