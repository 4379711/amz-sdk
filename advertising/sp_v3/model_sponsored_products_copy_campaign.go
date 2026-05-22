package sp_v3

import "github.com/bytedance/sonic"

// checks if the SponsoredProductsCopyCampaign type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCopyCampaign{}

// SponsoredProductsCopyCampaign struct for SponsoredProductsCopyCampaign
type SponsoredProductsCopyCampaign struct {
	TargetCampaignAttributes SponsoredProductsTargetCampaignAttributes `json:"targetCampaignAttributes"`
	// entity object identifier
	SourceCampaignId string `json:"sourceCampaignId"`
}

type _SponsoredProductsCopyCampaign SponsoredProductsCopyCampaign

// NewSponsoredProductsCopyCampaign instantiates a new SponsoredProductsCopyCampaign object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCopyCampaign(targetCampaignAttributes SponsoredProductsTargetCampaignAttributes, sourceCampaignId string) *SponsoredProductsCopyCampaign {
	this := SponsoredProductsCopyCampaign{}
	this.TargetCampaignAttributes = targetCampaignAttributes
	this.SourceCampaignId = sourceCampaignId
	return &this
}

// NewSponsoredProductsCopyCampaignWithDefaults instantiates a new SponsoredProductsCopyCampaign object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCopyCampaignWithDefaults() *SponsoredProductsCopyCampaign {
	this := SponsoredProductsCopyCampaign{}
	return &this
}

// GetTargetCampaignAttributes returns the TargetCampaignAttributes field value
func (o *SponsoredProductsCopyCampaign) GetTargetCampaignAttributes() SponsoredProductsTargetCampaignAttributes {
	if o == nil {
		var ret SponsoredProductsTargetCampaignAttributes
		return ret
	}

	return o.TargetCampaignAttributes
}

// GetTargetCampaignAttributesOk returns a tuple with the TargetCampaignAttributes field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCopyCampaign) GetTargetCampaignAttributesOk() (*SponsoredProductsTargetCampaignAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetCampaignAttributes, true
}

// SetTargetCampaignAttributes sets field value
func (o *SponsoredProductsCopyCampaign) SetTargetCampaignAttributes(v SponsoredProductsTargetCampaignAttributes) {
	o.TargetCampaignAttributes = v
}

// GetSourceCampaignId returns the SourceCampaignId field value
func (o *SponsoredProductsCopyCampaign) GetSourceCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceCampaignId
}

// GetSourceCampaignIdOk returns a tuple with the SourceCampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCopyCampaign) GetSourceCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceCampaignId, true
}

// SetSourceCampaignId sets field value
func (o *SponsoredProductsCopyCampaign) SetSourceCampaignId(v string) {
	o.SourceCampaignId = v
}

func (o SponsoredProductsCopyCampaign) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["targetCampaignAttributes"] = o.TargetCampaignAttributes
	toSerialize["sourceCampaignId"] = o.SourceCampaignId
	return toSerialize, nil
}

type NullableSponsoredProductsCopyCampaign struct {
	value *SponsoredProductsCopyCampaign
	isSet bool
}

func (v NullableSponsoredProductsCopyCampaign) Get() *SponsoredProductsCopyCampaign {
	return v.value
}

func (v *NullableSponsoredProductsCopyCampaign) Set(val *SponsoredProductsCopyCampaign) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCopyCampaign) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCopyCampaign) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCopyCampaign(val *SponsoredProductsCopyCampaign) *NullableSponsoredProductsCopyCampaign {
	return &NullableSponsoredProductsCopyCampaign{value: val, isSet: true}
}

func (v NullableSponsoredProductsCopyCampaign) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCopyCampaign) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
