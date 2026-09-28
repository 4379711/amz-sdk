package promotions_20251201

import (
	"github.com/bytedance/sonic"
)

// checks if the Merchandising type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Merchandising{}

// Merchandising The merchandising configuration for displaying promotions on the retail website.
type Merchandising struct {
	// Whether the promotion badge and details are displayed on the product detail pages and search results on the retail website.
	DisplayOnWebsite *string `json:"displayOnWebsite,omitempty"`
}

// NewMerchandising instantiates a new Merchandising object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchandising() *Merchandising {
	this := Merchandising{}
	return &this
}

// NewMerchandisingWithDefaults instantiates a new Merchandising object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchandisingWithDefaults() *Merchandising {
	this := Merchandising{}
	return &this
}

// GetDisplayOnWebsite returns the DisplayOnWebsite field value if set, zero value otherwise.
func (o *Merchandising) GetDisplayOnWebsite() string {
	if o == nil || IsNil(o.DisplayOnWebsite) {
		var ret string
		return ret
	}
	return *o.DisplayOnWebsite
}

// GetDisplayOnWebsiteOk returns a tuple with the DisplayOnWebsite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Merchandising) GetDisplayOnWebsiteOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayOnWebsite) {
		return nil, false
	}
	return o.DisplayOnWebsite, true
}

// HasDisplayOnWebsite returns a boolean if a field has been set.
func (o *Merchandising) HasDisplayOnWebsite() bool {
	if o != nil && !IsNil(o.DisplayOnWebsite) {
		return true
	}

	return false
}

// SetDisplayOnWebsite gets a reference to the given string and assigns it to the DisplayOnWebsite field.
func (o *Merchandising) SetDisplayOnWebsite(v string) {
	o.DisplayOnWebsite = &v
}

func (o Merchandising) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisplayOnWebsite) {
		toSerialize["displayOnWebsite"] = o.DisplayOnWebsite
	}
	return toSerialize, nil
}

type NullableMerchandising struct {
	value *Merchandising
	isSet bool
}

func (v NullableMerchandising) Get() *Merchandising {
	return v.value
}

func (v *NullableMerchandising) Set(val *Merchandising) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchandising) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchandising) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchandising(val *Merchandising) *NullableMerchandising {
	return &NullableMerchandising{value: val, isSet: true}
}

func (v NullableMerchandising) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableMerchandising) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
