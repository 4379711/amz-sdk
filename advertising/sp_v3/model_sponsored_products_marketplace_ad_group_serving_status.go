package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsMarketplaceAdGroupServingStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsMarketplaceAdGroupServingStatus{}

// SponsoredProductsMarketplaceAdGroupServingStatus struct for SponsoredProductsMarketplaceAdGroupServingStatus
type SponsoredProductsMarketplaceAdGroupServingStatus struct {
	StatusReasons []SponsoredProductsGlobalAdGroupMarketplaceServingReason `json:"statusReasons,omitempty"`
	Marketplace   *SponsoredProductsMarketplace                            `json:"marketplace,omitempty"`
	ServingStatus SponsoredProductsGlobalAdGroupMarketplaceServingStatus   `json:"servingStatus"`
}

type _SponsoredProductsMarketplaceAdGroupServingStatus SponsoredProductsMarketplaceAdGroupServingStatus

// NewSponsoredProductsMarketplaceAdGroupServingStatus instantiates a new SponsoredProductsMarketplaceAdGroupServingStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsMarketplaceAdGroupServingStatus(servingStatus SponsoredProductsGlobalAdGroupMarketplaceServingStatus) *SponsoredProductsMarketplaceAdGroupServingStatus {
	this := SponsoredProductsMarketplaceAdGroupServingStatus{}
	this.ServingStatus = servingStatus
	return &this
}

// NewSponsoredProductsMarketplaceAdGroupServingStatusWithDefaults instantiates a new SponsoredProductsMarketplaceAdGroupServingStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsMarketplaceAdGroupServingStatusWithDefaults() *SponsoredProductsMarketplaceAdGroupServingStatus {
	this := SponsoredProductsMarketplaceAdGroupServingStatus{}
	return &this
}

// GetStatusReasons returns the StatusReasons field value if set, zero value otherwise.
func (o *SponsoredProductsMarketplaceAdGroupServingStatus) GetStatusReasons() []SponsoredProductsGlobalAdGroupMarketplaceServingReason {
	if o == nil || IsNil(o.StatusReasons) {
		var ret []SponsoredProductsGlobalAdGroupMarketplaceServingReason
		return ret
	}
	return o.StatusReasons
}

// GetStatusReasonsOk returns a tuple with the StatusReasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsMarketplaceAdGroupServingStatus) GetStatusReasonsOk() ([]SponsoredProductsGlobalAdGroupMarketplaceServingReason, bool) {
	if o == nil || IsNil(o.StatusReasons) {
		return nil, false
	}
	return o.StatusReasons, true
}

// HasStatusReasons returns a boolean if a field has been set.
func (o *SponsoredProductsMarketplaceAdGroupServingStatus) HasStatusReasons() bool {
	if o != nil && !IsNil(o.StatusReasons) {
		return true
	}

	return false
}

// SetStatusReasons gets a reference to the given []SponsoredProductsGlobalAdGroupMarketplaceServingReason and assigns it to the StatusReasons field.
func (o *SponsoredProductsMarketplaceAdGroupServingStatus) SetStatusReasons(v []SponsoredProductsGlobalAdGroupMarketplaceServingReason) {
	o.StatusReasons = v
}

// GetMarketplace returns the Marketplace field value if set, zero value otherwise.
func (o *SponsoredProductsMarketplaceAdGroupServingStatus) GetMarketplace() SponsoredProductsMarketplace {
	if o == nil || IsNil(o.Marketplace) {
		var ret SponsoredProductsMarketplace
		return ret
	}
	return *o.Marketplace
}

// GetMarketplaceOk returns a tuple with the Marketplace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsMarketplaceAdGroupServingStatus) GetMarketplaceOk() (*SponsoredProductsMarketplace, bool) {
	if o == nil || IsNil(o.Marketplace) {
		return nil, false
	}
	return o.Marketplace, true
}

// HasMarketplace returns a boolean if a field has been set.
func (o *SponsoredProductsMarketplaceAdGroupServingStatus) HasMarketplace() bool {
	if o != nil && !IsNil(o.Marketplace) {
		return true
	}

	return false
}

// SetMarketplace gets a reference to the given SponsoredProductsMarketplace and assigns it to the Marketplace field.
func (o *SponsoredProductsMarketplaceAdGroupServingStatus) SetMarketplace(v SponsoredProductsMarketplace) {
	o.Marketplace = &v
}

// GetServingStatus returns the ServingStatus field value
func (o *SponsoredProductsMarketplaceAdGroupServingStatus) GetServingStatus() SponsoredProductsGlobalAdGroupMarketplaceServingStatus {
	if o == nil {
		var ret SponsoredProductsGlobalAdGroupMarketplaceServingStatus
		return ret
	}

	return o.ServingStatus
}

// GetServingStatusOk returns a tuple with the ServingStatus field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsMarketplaceAdGroupServingStatus) GetServingStatusOk() (*SponsoredProductsGlobalAdGroupMarketplaceServingStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServingStatus, true
}

// SetServingStatus sets field value
func (o *SponsoredProductsMarketplaceAdGroupServingStatus) SetServingStatus(v SponsoredProductsGlobalAdGroupMarketplaceServingStatus) {
	o.ServingStatus = v
}

func (o SponsoredProductsMarketplaceAdGroupServingStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StatusReasons) {
		toSerialize["statusReasons"] = o.StatusReasons
	}
	if !IsNil(o.Marketplace) {
		toSerialize["marketplace"] = o.Marketplace
	}
	toSerialize["servingStatus"] = o.ServingStatus
	return toSerialize, nil
}

type NullableSponsoredProductsMarketplaceAdGroupServingStatus struct {
	value *SponsoredProductsMarketplaceAdGroupServingStatus
	isSet bool
}

func (v NullableSponsoredProductsMarketplaceAdGroupServingStatus) Get() *SponsoredProductsMarketplaceAdGroupServingStatus {
	return v.value
}

func (v *NullableSponsoredProductsMarketplaceAdGroupServingStatus) Set(val *SponsoredProductsMarketplaceAdGroupServingStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsMarketplaceAdGroupServingStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsMarketplaceAdGroupServingStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsMarketplaceAdGroupServingStatus(val *SponsoredProductsMarketplaceAdGroupServingStatus) *NullableSponsoredProductsMarketplaceAdGroupServingStatus {
	return &NullableSponsoredProductsMarketplaceAdGroupServingStatus{value: val, isSet: true}
}

func (v NullableSponsoredProductsMarketplaceAdGroupServingStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsMarketplaceAdGroupServingStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
