package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the ListGlobalTargetableCategoriesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListGlobalTargetableCategoriesRequest{}

// ListGlobalTargetableCategoriesRequest struct for ListGlobalTargetableCategoriesRequest
type ListGlobalTargetableCategoriesRequest struct {
	// A list of country codes to fetch category tree <br> <table><thead><tr><th>Country Code</th><th>Country</th></tr></thead><tbody><tr><td>US</td><td>United States</td></tr><tr><td>CA</td><td>Canada</td></tr><tr><td>MX</td><td>Mexico</td></tr><tr><td>BR</td><td>Brazil</td></tr><tr><td>UK</td><td>United Kingdom</td></tr><tr><td>DE</td><td>Germany</td></tr><tr><td>FR</td><td>France</td></tr><tr><td>ES</td><td>Spain</td></tr><tr><td>IT</td><td>Italy</td></tr><tr><td>IN</td><td>India</td></tr><tr><td>AE</td><td>United Arab Emirates</td></tr><tr><td>SA</td><td>Saudi Arabia</td></tr><tr><td>NL</td><td>Netherlands</td></tr><tr><td>PL</td><td>Poland</td></tr><tr><td>BE</td><td>Belgium</td></tr><tr><td>SE</td><td>Sweden</td></tr><tr><td>TR</td><td>Turkey</td></tr><tr><td>EG</td><td>Egypt</td></tr><tr><td>JP</td><td>Japan</td></tr><tr><td>AU</td><td>Australia</td></tr><tr><td>SG</td><td>Singapore</td></tr></tbody></table>
	CountryCodes []string `json:"countryCodes"`
	// The locale to which the caller wishes to translate the targetable categories to. For example, if the caller wishes to receive the targetable categories in Simplified Chinese, the locale parameter should be set to zh_CN. If no locale is provided, the returned targetable categories will be in the default language of the marketplace. <br><br> <b>Supported locale's</b> : \"ar_AE\",\"de_DE\",\"en_AE\",\"en_AU\",\"en_CA\",\"en_GB\",\"en_IN\",\"en_SG\",\"en_US\",\"es_ES\",\"es_MX\",\"fr_CA\",\"fr_FR\",\"hi_IN\",\"it_IT\",\"ja_JP\",\"ko_KR\",\"nl_NL\",\"pl_PL\",\"pt_BR\",\"sv_SE\",\"ta_IN\",\"th_TH\",\"tr_TR\",\"vi_VN\",\"zh_CN\"
	Locale *string `json:"locale,omitempty"`
}

type _ListGlobalTargetableCategoriesRequest ListGlobalTargetableCategoriesRequest

// NewListGlobalTargetableCategoriesRequest instantiates a new ListGlobalTargetableCategoriesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListGlobalTargetableCategoriesRequest(countryCodes []string) *ListGlobalTargetableCategoriesRequest {
	this := ListGlobalTargetableCategoriesRequest{}
	this.CountryCodes = countryCodes
	return &this
}

// NewListGlobalTargetableCategoriesRequestWithDefaults instantiates a new ListGlobalTargetableCategoriesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListGlobalTargetableCategoriesRequestWithDefaults() *ListGlobalTargetableCategoriesRequest {
	this := ListGlobalTargetableCategoriesRequest{}
	return &this
}

// GetCountryCodes returns the CountryCodes field value
func (o *ListGlobalTargetableCategoriesRequest) GetCountryCodes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.CountryCodes
}

// GetCountryCodesOk returns a tuple with the CountryCodes field value
// and a boolean to check if the value has been set.
func (o *ListGlobalTargetableCategoriesRequest) GetCountryCodesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CountryCodes, true
}

// SetCountryCodes sets field value
func (o *ListGlobalTargetableCategoriesRequest) SetCountryCodes(v []string) {
	o.CountryCodes = v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *ListGlobalTargetableCategoriesRequest) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListGlobalTargetableCategoriesRequest) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *ListGlobalTargetableCategoriesRequest) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *ListGlobalTargetableCategoriesRequest) SetLocale(v string) {
	o.Locale = &v
}

func (o ListGlobalTargetableCategoriesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["countryCodes"] = o.CountryCodes
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	return toSerialize, nil
}

type NullableListGlobalTargetableCategoriesRequest struct {
	value *ListGlobalTargetableCategoriesRequest
	isSet bool
}

func (v NullableListGlobalTargetableCategoriesRequest) Get() *ListGlobalTargetableCategoriesRequest {
	return v.value
}

func (v *NullableListGlobalTargetableCategoriesRequest) Set(val *ListGlobalTargetableCategoriesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableListGlobalTargetableCategoriesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableListGlobalTargetableCategoriesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListGlobalTargetableCategoriesRequest(val *ListGlobalTargetableCategoriesRequest) *NullableListGlobalTargetableCategoriesRequest {
	return &NullableListGlobalTargetableCategoriesRequest{value: val, isSet: true}
}

func (v NullableListGlobalTargetableCategoriesRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListGlobalTargetableCategoriesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
