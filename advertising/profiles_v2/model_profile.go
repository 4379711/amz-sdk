package profiles_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the Profile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Profile{}

// Profile struct for Profile
type Profile struct {
	ProfileId   *int64       `json:"profileId,omitempty"`
	CountryCode *CountryCode `json:"countryCode,omitempty"`
	// The currency used for all monetary values for entities under this profile. |Region|`countryCode`|Country Name|`currencyCode`| |-----|------|------|------| |NA|BR|Brazil|BRL| |NA|CA|Canada|CAD| |NA|MX|Mexico|MXN| |NA|US|United States|USD| |EU|AE|United Arab Emirates|AED| |EU|BE|Belgium|EUR| |EU|DE|Germany|EUR| |EU|EG|Egypt|EGP| |EU|ES|Spain|EUR| |EU|FR|France|EUR| |EU|IN|India|INR| |EU|IT|Italy|EUR| |EU|NL|The Netherlands|EUR| |EU|PL|Poland|PLN| |EU|SA|Saudi Arabia|SAR| |EU|SE|Sweden|SEK| |EU|TR|Turkey|TRY| |EU|UK|United Kingdom|GBP| |EU|ZA| South Africa | ZAR| |FE|AU|Australia|AUD| |FE|JP|Japan|JPY| |FE|SG|Singapore|SGD|
	CurrencyCode *string `json:"currencyCode,omitempty"`
	// Note that this field applies to Sponsored Product campaigns for seller type accounts only. Not supported for vendor type accounts.
	DailyBudget *float32 `json:"dailyBudget,omitempty"`
	// The time zone used for all date-based campaign management and reporting. |Region|`countryCode`|Country Name|`timezone`| |------|-----|-----|------| |NA|BR|Brazil|America/Sao_Paulo| |NA|CA|Canada|America/Los_Angeles| |NA|MX|Mexico|America/Los_Angeles| |NA|US|United States|America/Los_Angeles| |EU|AE|United Arab Emirates|Asia/Dubai| |EU|BE|Belgium|Europe/Brussels| |EU|DE|Germany|Europe/Paris| |EU|EG|Egypt|Africa/Cairo| |EU|ES|Spain|Europe/Paris| |EU|FR|France|Europe/Paris| |EU|IN|India|Asia/Kolkata| |EU|IT|Italy|Europe/Paris| |EU|NL|The Netherlands|Europe/Amsterdam| |EU|PL|Poland|Europe/Warsaw| |EU|SA|Saudi Arabia|Asia/Riyadh| |EU|SE|Sweden|Europe/Stockholm| |EU|TR|Turkey|Europe/Istanbul| |EU|UK|United Kingdom|Europe/London| |EU|ZA| South Africa | Africa/Johannesburg | |FE|AU|Australia|Australia/Sydney| |FE|JP|Japan|Asia/Tokyo| |FE|SG|Singapore|Asia/Singapore|
	Timezone    *string      `json:"timezone,omitempty"`
	AccountInfo *AccountInfo `json:"accountInfo,omitempty"`
}

// NewProfile instantiates a new Profile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfile() *Profile {
	this := Profile{}
	return &this
}

// NewProfileWithDefaults instantiates a new Profile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileWithDefaults() *Profile {
	this := Profile{}
	return &this
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *Profile) GetProfileId() int64 {
	if o == nil || IsNil(o.ProfileId) {
		var ret int64
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Profile) GetProfileIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ProfileId) {
		return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *Profile) HasProfileId() bool {
	if o != nil && !IsNil(o.ProfileId) {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given int64 and assigns it to the ProfileId field.
func (o *Profile) SetProfileId(v int64) {
	o.ProfileId = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *Profile) GetCountryCode() CountryCode {
	if o == nil || IsNil(o.CountryCode) {
		var ret CountryCode
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Profile) GetCountryCodeOk() (*CountryCode, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *Profile) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given CountryCode and assigns it to the CountryCode field.
func (o *Profile) SetCountryCode(v CountryCode) {
	o.CountryCode = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *Profile) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Profile) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *Profile) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *Profile) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetDailyBudget returns the DailyBudget field value if set, zero value otherwise.
func (o *Profile) GetDailyBudget() float32 {
	if o == nil || IsNil(o.DailyBudget) {
		var ret float32
		return ret
	}
	return *o.DailyBudget
}

// GetDailyBudgetOk returns a tuple with the DailyBudget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Profile) GetDailyBudgetOk() (*float32, bool) {
	if o == nil || IsNil(o.DailyBudget) {
		return nil, false
	}
	return o.DailyBudget, true
}

// HasDailyBudget returns a boolean if a field has been set.
func (o *Profile) HasDailyBudget() bool {
	if o != nil && !IsNil(o.DailyBudget) {
		return true
	}

	return false
}

// SetDailyBudget gets a reference to the given float32 and assigns it to the DailyBudget field.
func (o *Profile) SetDailyBudget(v float32) {
	o.DailyBudget = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *Profile) GetTimezone() string {
	if o == nil || IsNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Profile) GetTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *Profile) HasTimezone() bool {
	if o != nil && !IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *Profile) SetTimezone(v string) {
	o.Timezone = &v
}

// GetAccountInfo returns the AccountInfo field value if set, zero value otherwise.
func (o *Profile) GetAccountInfo() AccountInfo {
	if o == nil || IsNil(o.AccountInfo) {
		var ret AccountInfo
		return ret
	}
	return *o.AccountInfo
}

// GetAccountInfoOk returns a tuple with the AccountInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Profile) GetAccountInfoOk() (*AccountInfo, bool) {
	if o == nil || IsNil(o.AccountInfo) {
		return nil, false
	}
	return o.AccountInfo, true
}

// HasAccountInfo returns a boolean if a field has been set.
func (o *Profile) HasAccountInfo() bool {
	if o != nil && !IsNil(o.AccountInfo) {
		return true
	}

	return false
}

// SetAccountInfo gets a reference to the given AccountInfo and assigns it to the AccountInfo field.
func (o *Profile) SetAccountInfo(v AccountInfo) {
	o.AccountInfo = &v
}

func (o Profile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProfileId) {
		toSerialize["profileId"] = o.ProfileId
	}
	if !IsNil(o.CountryCode) {
		toSerialize["countryCode"] = o.CountryCode
	}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currencyCode"] = o.CurrencyCode
	}
	if !IsNil(o.DailyBudget) {
		toSerialize["dailyBudget"] = o.DailyBudget
	}
	if !IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}
	if !IsNil(o.AccountInfo) {
		toSerialize["accountInfo"] = o.AccountInfo
	}
	return toSerialize, nil
}

type NullableProfile struct {
	value *Profile
	isSet bool
}

func (v NullableProfile) Get() *Profile {
	return v.value
}

func (v *NullableProfile) Set(val *Profile) {
	v.value = val
	v.isSet = true
}

func (v NullableProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfile(val *Profile) *NullableProfile {
	return &NullableProfile{value: val, isSet: true}
}

func (v NullableProfile) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
