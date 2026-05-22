package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
	"time"
)

// checks if the PackingOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PackingOption{}

// PackingOption A packing option contains a set of pack groups plus additional information about the packing option, such as any discounts or fees if it's selected.
type PackingOption struct {
	// Discount for the offered option.
	Discounts []Incentive `json:"discounts"`
	// The time at which this packing option is no longer valid. In [ISO 8601](https://developer-docs.amazon.com/sp-api/docs/iso-8601) datetime format with pattern `yyyy-MM-ddTHH:mm:ss.sssZ`.
	Expiration *time.Time `json:"expiration,omitempty"`
	// Fee for the offered option.
	Fees []Incentive `json:"fees"`
	// Packing group IDs.
	PackingGroups []string `json:"packingGroups"`
	// Identifier of a packing option.
	PackingOptionId string `json:"packingOptionId" validate:"regexp=^[a-zA-Z0-9-]*$"`
	// The status of the packing option. Possible values: `OFFERED`, `ACCEPTED`, `EXPIRED`.
	Status string `json:"status"`
	// A list of possible configurations for this option.
	SupportedConfigurations []PackingConfiguration `json:"supportedConfigurations"`
	// **This field is deprecated**. Use the `shippingRequirements` property under `supportedConfigurations` instead. List of supported shipping modes.
	SupportedShippingConfigurations []ShippingConfiguration `json:"supportedShippingConfigurations"`
}

type _PackingOption PackingOption

// NewPackingOption instantiates a new PackingOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackingOption(discounts []Incentive, fees []Incentive, packingGroups []string, packingOptionId string, status string, supportedConfigurations []PackingConfiguration, supportedShippingConfigurations []ShippingConfiguration) *PackingOption {
	this := PackingOption{}
	this.Discounts = discounts
	this.Fees = fees
	this.PackingGroups = packingGroups
	this.PackingOptionId = packingOptionId
	this.Status = status
	this.SupportedConfigurations = supportedConfigurations
	this.SupportedShippingConfigurations = supportedShippingConfigurations
	return &this
}

// NewPackingOptionWithDefaults instantiates a new PackingOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackingOptionWithDefaults() *PackingOption {
	this := PackingOption{}
	return &this
}

// GetDiscounts returns the Discounts field value
func (o *PackingOption) GetDiscounts() []Incentive {
	if o == nil {
		var ret []Incentive
		return ret
	}

	return o.Discounts
}

// GetDiscountsOk returns a tuple with the Discounts field value
// and a boolean to check if the value has been set.
func (o *PackingOption) GetDiscountsOk() ([]Incentive, bool) {
	if o == nil {
		return nil, false
	}
	return o.Discounts, true
}

// SetDiscounts sets field value
func (o *PackingOption) SetDiscounts(v []Incentive) {
	o.Discounts = v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *PackingOption) GetExpiration() time.Time {
	if o == nil || IsNil(o.Expiration) {
		var ret time.Time
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackingOption) GetExpirationOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Expiration) {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *PackingOption) HasExpiration() bool {
	if o != nil && !IsNil(o.Expiration) {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given time.Time and assigns it to the Expiration field.
func (o *PackingOption) SetExpiration(v time.Time) {
	o.Expiration = &v
}

// GetFees returns the Fees field value
func (o *PackingOption) GetFees() []Incentive {
	if o == nil {
		var ret []Incentive
		return ret
	}

	return o.Fees
}

// GetFeesOk returns a tuple with the Fees field value
// and a boolean to check if the value has been set.
func (o *PackingOption) GetFeesOk() ([]Incentive, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fees, true
}

// SetFees sets field value
func (o *PackingOption) SetFees(v []Incentive) {
	o.Fees = v
}

// GetPackingGroups returns the PackingGroups field value
func (o *PackingOption) GetPackingGroups() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PackingGroups
}

// GetPackingGroupsOk returns a tuple with the PackingGroups field value
// and a boolean to check if the value has been set.
func (o *PackingOption) GetPackingGroupsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PackingGroups, true
}

// SetPackingGroups sets field value
func (o *PackingOption) SetPackingGroups(v []string) {
	o.PackingGroups = v
}

// GetPackingOptionId returns the PackingOptionId field value
func (o *PackingOption) GetPackingOptionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PackingOptionId
}

// GetPackingOptionIdOk returns a tuple with the PackingOptionId field value
// and a boolean to check if the value has been set.
func (o *PackingOption) GetPackingOptionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackingOptionId, true
}

// SetPackingOptionId sets field value
func (o *PackingOption) SetPackingOptionId(v string) {
	o.PackingOptionId = v
}

// GetStatus returns the Status field value
func (o *PackingOption) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PackingOption) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PackingOption) SetStatus(v string) {
	o.Status = v
}

// GetSupportedConfigurations returns the SupportedConfigurations field value
func (o *PackingOption) GetSupportedConfigurations() []PackingConfiguration {
	if o == nil {
		var ret []PackingConfiguration
		return ret
	}

	return o.SupportedConfigurations
}

// GetSupportedConfigurationsOk returns a tuple with the SupportedConfigurations field value
// and a boolean to check if the value has been set.
func (o *PackingOption) GetSupportedConfigurationsOk() ([]PackingConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return o.SupportedConfigurations, true
}

// SetSupportedConfigurations sets field value
func (o *PackingOption) SetSupportedConfigurations(v []PackingConfiguration) {
	o.SupportedConfigurations = v
}

// GetSupportedShippingConfigurations returns the SupportedShippingConfigurations field value
func (o *PackingOption) GetSupportedShippingConfigurations() []ShippingConfiguration {
	if o == nil {
		var ret []ShippingConfiguration
		return ret
	}

	return o.SupportedShippingConfigurations
}

// GetSupportedShippingConfigurationsOk returns a tuple with the SupportedShippingConfigurations field value
// and a boolean to check if the value has been set.
func (o *PackingOption) GetSupportedShippingConfigurationsOk() ([]ShippingConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return o.SupportedShippingConfigurations, true
}

// SetSupportedShippingConfigurations sets field value
func (o *PackingOption) SetSupportedShippingConfigurations(v []ShippingConfiguration) {
	o.SupportedShippingConfigurations = v
}

func (o PackingOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["discounts"] = o.Discounts
	if !IsNil(o.Expiration) {
		toSerialize["expiration"] = o.Expiration
	}
	toSerialize["fees"] = o.Fees
	toSerialize["packingGroups"] = o.PackingGroups
	toSerialize["packingOptionId"] = o.PackingOptionId
	toSerialize["status"] = o.Status
	toSerialize["supportedConfigurations"] = o.SupportedConfigurations
	toSerialize["supportedShippingConfigurations"] = o.SupportedShippingConfigurations
	return toSerialize, nil
}

type NullablePackingOption struct {
	value *PackingOption
	isSet bool
}

func (v NullablePackingOption) Get() *PackingOption {
	return v.value
}

func (v *NullablePackingOption) Set(val *PackingOption) {
	v.value = val
	v.isSet = true
}

func (v NullablePackingOption) IsSet() bool {
	return v.isSet
}

func (v *NullablePackingOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackingOption(val *PackingOption) *NullablePackingOption {
	return &NullablePackingOption{value: val, isSet: true}
}

func (v NullablePackingOption) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePackingOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
