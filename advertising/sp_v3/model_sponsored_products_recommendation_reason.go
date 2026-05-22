package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsRecommendationReason type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsRecommendationReason{}

// SponsoredProductsRecommendationReason Provides a reason for why this target is being recommended for harvesting
type SponsoredProductsRecommendationReason struct {
	// The reason for the recommendation
	Reason *string `json:"reason,omitempty"`
	// The data supporting the recommendation reason
	Data *string `json:"data,omitempty"`
}

// NewSponsoredProductsRecommendationReason instantiates a new SponsoredProductsRecommendationReason object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsRecommendationReason() *SponsoredProductsRecommendationReason {
	this := SponsoredProductsRecommendationReason{}
	return &this
}

// NewSponsoredProductsRecommendationReasonWithDefaults instantiates a new SponsoredProductsRecommendationReason object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsRecommendationReasonWithDefaults() *SponsoredProductsRecommendationReason {
	this := SponsoredProductsRecommendationReason{}
	return &this
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *SponsoredProductsRecommendationReason) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsRecommendationReason) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *SponsoredProductsRecommendationReason) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *SponsoredProductsRecommendationReason) SetReason(v string) {
	o.Reason = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SponsoredProductsRecommendationReason) GetData() string {
	if o == nil || IsNil(o.Data) {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsRecommendationReason) GetDataOk() (*string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SponsoredProductsRecommendationReason) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *SponsoredProductsRecommendationReason) SetData(v string) {
	o.Data = &v
}

func (o SponsoredProductsRecommendationReason) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableSponsoredProductsRecommendationReason struct {
	value *SponsoredProductsRecommendationReason
	isSet bool
}

func (v NullableSponsoredProductsRecommendationReason) Get() *SponsoredProductsRecommendationReason {
	return v.value
}

func (v *NullableSponsoredProductsRecommendationReason) Set(val *SponsoredProductsRecommendationReason) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsRecommendationReason) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsRecommendationReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsRecommendationReason(val *SponsoredProductsRecommendationReason) *NullableSponsoredProductsRecommendationReason {
	return &NullableSponsoredProductsRecommendationReason{value: val, isSet: true}
}

func (v NullableSponsoredProductsRecommendationReason) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsRecommendationReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
