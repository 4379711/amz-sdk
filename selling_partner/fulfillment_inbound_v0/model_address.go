package fulfillment_inbound_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the Address type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Address{}

// Address Specific details to identify a place.
type Address struct {
	// Name of the individual or business.
	Name string `json:"Name"`
	// The street address information.
	AddressLine1 string `json:"AddressLine1"`
	// Additional street address information, if required.
	AddressLine2 *string `json:"AddressLine2,omitempty"`
	// The district or county.
	DistrictOrCounty *string `json:"DistrictOrCounty,omitempty"`
	// The city.
	City string `json:"City"`
	// The state or province code.  If state or province codes are used in your marketplace, it is recommended that you include one with your request. This helps Amazon to select the most appropriate Amazon fulfillment center for your inbound shipment plan.
	StateOrProvinceCode string `json:"StateOrProvinceCode"`
	// The country code in two-character ISO 3166-1 alpha-2 format.
	CountryCode string `json:"CountryCode"`
	// The postal code.  If postal codes are used in your marketplace, we recommended that you include one with your request. This helps Amazon select the most appropriate Amazon fulfillment center for the inbound shipment plan.
	PostalCode string `json:"PostalCode"`
}

type _Address Address

// NewAddress instantiates a new Address object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddress(name string, addressLine1 string, city string, stateOrProvinceCode string, countryCode string, postalCode string) *Address {
	this := Address{}
	this.Name = name
	this.AddressLine1 = addressLine1
	this.City = city
	this.StateOrProvinceCode = stateOrProvinceCode
	this.CountryCode = countryCode
	this.PostalCode = postalCode
	return &this
}

// NewAddressWithDefaults instantiates a new Address object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressWithDefaults() *Address {
	this := Address{}
	return &this
}

// GetName returns the Name field value
func (o *Address) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Address) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Address) SetName(v string) {
	o.Name = v
}

// GetAddressLine1 returns the AddressLine1 field value
func (o *Address) GetAddressLine1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddressLine1
}

// GetAddressLine1Ok returns a tuple with the AddressLine1 field value
// and a boolean to check if the value has been set.
func (o *Address) GetAddressLine1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressLine1, true
}

// SetAddressLine1 sets field value
func (o *Address) SetAddressLine1(v string) {
	o.AddressLine1 = v
}

// GetAddressLine2 returns the AddressLine2 field value if set, zero value otherwise.
func (o *Address) GetAddressLine2() string {
	if o == nil || IsNil(o.AddressLine2) {
		var ret string
		return ret
	}
	return *o.AddressLine2
}

// GetAddressLine2Ok returns a tuple with the AddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetAddressLine2Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine2) {
		return nil, false
	}
	return o.AddressLine2, true
}

// HasAddressLine2 returns a boolean if a field has been set.
func (o *Address) HasAddressLine2() bool {
	if o != nil && !IsNil(o.AddressLine2) {
		return true
	}

	return false
}

// SetAddressLine2 gets a reference to the given string and assigns it to the AddressLine2 field.
func (o *Address) SetAddressLine2(v string) {
	o.AddressLine2 = &v
}

// GetDistrictOrCounty returns the DistrictOrCounty field value if set, zero value otherwise.
func (o *Address) GetDistrictOrCounty() string {
	if o == nil || IsNil(o.DistrictOrCounty) {
		var ret string
		return ret
	}
	return *o.DistrictOrCounty
}

// GetDistrictOrCountyOk returns a tuple with the DistrictOrCounty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetDistrictOrCountyOk() (*string, bool) {
	if o == nil || IsNil(o.DistrictOrCounty) {
		return nil, false
	}
	return o.DistrictOrCounty, true
}

// HasDistrictOrCounty returns a boolean if a field has been set.
func (o *Address) HasDistrictOrCounty() bool {
	if o != nil && !IsNil(o.DistrictOrCounty) {
		return true
	}

	return false
}

// SetDistrictOrCounty gets a reference to the given string and assigns it to the DistrictOrCounty field.
func (o *Address) SetDistrictOrCounty(v string) {
	o.DistrictOrCounty = &v
}

// GetCity returns the City field value
func (o *Address) GetCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *Address) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *Address) SetCity(v string) {
	o.City = v
}

// GetStateOrProvinceCode returns the StateOrProvinceCode field value
func (o *Address) GetStateOrProvinceCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StateOrProvinceCode
}

// GetStateOrProvinceCodeOk returns a tuple with the StateOrProvinceCode field value
// and a boolean to check if the value has been set.
func (o *Address) GetStateOrProvinceCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StateOrProvinceCode, true
}

// SetStateOrProvinceCode sets field value
func (o *Address) SetStateOrProvinceCode(v string) {
	o.StateOrProvinceCode = v
}

// GetCountryCode returns the CountryCode field value
func (o *Address) GetCountryCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value
// and a boolean to check if the value has been set.
func (o *Address) GetCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryCode, true
}

// SetCountryCode sets field value
func (o *Address) SetCountryCode(v string) {
	o.CountryCode = v
}

// GetPostalCode returns the PostalCode field value
func (o *Address) GetPostalCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value
// and a boolean to check if the value has been set.
func (o *Address) GetPostalCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostalCode, true
}

// SetPostalCode sets field value
func (o *Address) SetPostalCode(v string) {
	o.PostalCode = v
}

func (o Address) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Name"] = o.Name
	toSerialize["AddressLine1"] = o.AddressLine1
	if !IsNil(o.AddressLine2) {
		toSerialize["AddressLine2"] = o.AddressLine2
	}
	if !IsNil(o.DistrictOrCounty) {
		toSerialize["DistrictOrCounty"] = o.DistrictOrCounty
	}
	toSerialize["City"] = o.City
	toSerialize["StateOrProvinceCode"] = o.StateOrProvinceCode
	toSerialize["CountryCode"] = o.CountryCode
	toSerialize["PostalCode"] = o.PostalCode
	return toSerialize, nil
}

type NullableAddress struct {
	value *Address
	isSet bool
}

func (v NullableAddress) Get() *Address {
	return v.value
}

func (v *NullableAddress) Set(val *Address) {
	v.value = val
	v.isSet = true
}

func (v NullableAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddress(val *Address) *NullableAddress {
	return &NullableAddress{value: val, isSet: true}
}

func (v NullableAddress) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
