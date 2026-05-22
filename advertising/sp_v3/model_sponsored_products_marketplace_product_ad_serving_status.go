package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsMarketplaceProductAdServingStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsMarketplaceProductAdServingStatus{}

// SponsoredProductsMarketplaceProductAdServingStatus struct for SponsoredProductsMarketplaceProductAdServingStatus
type SponsoredProductsMarketplaceProductAdServingStatus struct {
	StatusReasons []SponsoredProductsGlobalProductAdMarketplaceServingReason `json:"statusReasons,omitempty"`
	Marketplace   *SponsoredProductsMarketplace                              `json:"marketplace,omitempty"`
	ServingStatus SponsoredProductsGlobalProductAdMarketplaceServingStatus   `json:"servingStatus"`
}

type _SponsoredProductsMarketplaceProductAdServingStatus SponsoredProductsMarketplaceProductAdServingStatus

// NewSponsoredProductsMarketplaceProductAdServingStatus instantiates a new SponsoredProductsMarketplaceProductAdServingStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsMarketplaceProductAdServingStatus(servingStatus SponsoredProductsGlobalProductAdMarketplaceServingStatus) *SponsoredProductsMarketplaceProductAdServingStatus {
	this := SponsoredProductsMarketplaceProductAdServingStatus{}
	this.ServingStatus = servingStatus
	return &this
}

// NewSponsoredProductsMarketplaceProductAdServingStatusWithDefaults instantiates a new SponsoredProductsMarketplaceProductAdServingStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsMarketplaceProductAdServingStatusWithDefaults() *SponsoredProductsMarketplaceProductAdServingStatus {
	this := SponsoredProductsMarketplaceProductAdServingStatus{}
	return &this
}

// GetStatusReasons returns the StatusReasons field value if set, zero value otherwise.
func (o *SponsoredProductsMarketplaceProductAdServingStatus) GetStatusReasons() []SponsoredProductsGlobalProductAdMarketplaceServingReason {
	if o == nil || IsNil(o.StatusReasons) {
		var ret []SponsoredProductsGlobalProductAdMarketplaceServingReason
		return ret
	}
	return o.StatusReasons
}

// GetStatusReasonsOk returns a tuple with the StatusReasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsMarketplaceProductAdServingStatus) GetStatusReasonsOk() ([]SponsoredProductsGlobalProductAdMarketplaceServingReason, bool) {
	if o == nil || IsNil(o.StatusReasons) {
		return nil, false
	}
	return o.StatusReasons, true
}

// HasStatusReasons returns a boolean if a field has been set.
func (o *SponsoredProductsMarketplaceProductAdServingStatus) HasStatusReasons() bool {
	if o != nil && !IsNil(o.StatusReasons) {
		return true
	}

	return false
}

// SetStatusReasons gets a reference to the given []SponsoredProductsGlobalProductAdMarketplaceServingReason and assigns it to the StatusReasons field.
func (o *SponsoredProductsMarketplaceProductAdServingStatus) SetStatusReasons(v []SponsoredProductsGlobalProductAdMarketplaceServingReason) {
	o.StatusReasons = v
}

// GetMarketplace returns the Marketplace field value if set, zero value otherwise.
func (o *SponsoredProductsMarketplaceProductAdServingStatus) GetMarketplace() SponsoredProductsMarketplace {
	if o == nil || IsNil(o.Marketplace) {
		var ret SponsoredProductsMarketplace
		return ret
	}
	return *o.Marketplace
}

// GetMarketplaceOk returns a tuple with the Marketplace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsMarketplaceProductAdServingStatus) GetMarketplaceOk() (*SponsoredProductsMarketplace, bool) {
	if o == nil || IsNil(o.Marketplace) {
		return nil, false
	}
	return o.Marketplace, true
}

// HasMarketplace returns a boolean if a field has been set.
func (o *SponsoredProductsMarketplaceProductAdServingStatus) HasMarketplace() bool {
	if o != nil && !IsNil(o.Marketplace) {
		return true
	}

	return false
}

// SetMarketplace gets a reference to the given SponsoredProductsMarketplace and assigns it to the Marketplace field.
func (o *SponsoredProductsMarketplaceProductAdServingStatus) SetMarketplace(v SponsoredProductsMarketplace) {
	o.Marketplace = &v
}

// GetServingStatus returns the ServingStatus field value
func (o *SponsoredProductsMarketplaceProductAdServingStatus) GetServingStatus() SponsoredProductsGlobalProductAdMarketplaceServingStatus {
	if o == nil {
		var ret SponsoredProductsGlobalProductAdMarketplaceServingStatus
		return ret
	}

	return o.ServingStatus
}

// GetServingStatusOk returns a tuple with the ServingStatus field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsMarketplaceProductAdServingStatus) GetServingStatusOk() (*SponsoredProductsGlobalProductAdMarketplaceServingStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServingStatus, true
}

// SetServingStatus sets field value
func (o *SponsoredProductsMarketplaceProductAdServingStatus) SetServingStatus(v SponsoredProductsGlobalProductAdMarketplaceServingStatus) {
	o.ServingStatus = v
}

func (o SponsoredProductsMarketplaceProductAdServingStatus) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsMarketplaceProductAdServingStatus struct {
	value *SponsoredProductsMarketplaceProductAdServingStatus
	isSet bool
}

func (v NullableSponsoredProductsMarketplaceProductAdServingStatus) Get() *SponsoredProductsMarketplaceProductAdServingStatus {
	return v.value
}

func (v *NullableSponsoredProductsMarketplaceProductAdServingStatus) Set(val *SponsoredProductsMarketplaceProductAdServingStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsMarketplaceProductAdServingStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsMarketplaceProductAdServingStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsMarketplaceProductAdServingStatus(val *SponsoredProductsMarketplaceProductAdServingStatus) *NullableSponsoredProductsMarketplaceProductAdServingStatus {
	return &NullableSponsoredProductsMarketplaceProductAdServingStatus{value: val, isSet: true}
}

func (v NullableSponsoredProductsMarketplaceProductAdServingStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsMarketplaceProductAdServingStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
