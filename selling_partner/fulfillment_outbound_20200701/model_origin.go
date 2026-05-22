package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the Origin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Origin{}

// Origin The origin for the delivery offer.
type Origin struct {
	// The two digit country code the items should ship from. In ISO 3166-1 alpha-2 format.
	CountryCode string `json:"countryCode"`
}

type _Origin Origin

// NewOrigin instantiates a new Origin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrigin(countryCode string) *Origin {
	this := Origin{}
	this.CountryCode = countryCode
	return &this
}

// NewOriginWithDefaults instantiates a new Origin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOriginWithDefaults() *Origin {
	this := Origin{}
	return &this
}

// GetCountryCode returns the CountryCode field value
func (o *Origin) GetCountryCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value
// and a boolean to check if the value has been set.
func (o *Origin) GetCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryCode, true
}

// SetCountryCode sets field value
func (o *Origin) SetCountryCode(v string) {
	o.CountryCode = v
}

func (o Origin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["countryCode"] = o.CountryCode
	return toSerialize, nil
}

type NullableOrigin struct {
	value *Origin
	isSet bool
}

func (v NullableOrigin) Get() *Origin {
	return v.value
}

func (v *NullableOrigin) Set(val *Origin) {
	v.value = val
	v.isSet = true
}

func (v NullableOrigin) IsSet() bool {
	return v.isSet
}

func (v *NullableOrigin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrigin(val *Origin) *NullableOrigin {
	return &NullableOrigin{value: val, isSet: true}
}

func (v NullableOrigin) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOrigin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
