package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SPGetAssociatedCampaignsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SPGetAssociatedCampaignsResponse{}

// SPGetAssociatedCampaignsResponse struct for SPGetAssociatedCampaignsResponse
type SPGetAssociatedCampaignsResponse struct {
	// A list of campaigns that are associated to this budget rule.
	AssociatedCampaigns []AssociatedCampaign `json:"associatedCampaigns,omitempty"`
	// To retrieve the next page of results, call the same operation and specify this token in the request. If the `nextToken` field is empty, there are no further results.
	NextToken *string `json:"nextToken,omitempty"`
}

// NewSPGetAssociatedCampaignsResponse instantiates a new SPGetAssociatedCampaignsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSPGetAssociatedCampaignsResponse() *SPGetAssociatedCampaignsResponse {
	this := SPGetAssociatedCampaignsResponse{}
	return &this
}

// NewSPGetAssociatedCampaignsResponseWithDefaults instantiates a new SPGetAssociatedCampaignsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSPGetAssociatedCampaignsResponseWithDefaults() *SPGetAssociatedCampaignsResponse {
	this := SPGetAssociatedCampaignsResponse{}
	return &this
}

// GetAssociatedCampaigns returns the AssociatedCampaigns field value if set, zero value otherwise.
func (o *SPGetAssociatedCampaignsResponse) GetAssociatedCampaigns() []AssociatedCampaign {
	if o == nil || IsNil(o.AssociatedCampaigns) {
		var ret []AssociatedCampaign
		return ret
	}
	return o.AssociatedCampaigns
}

// GetAssociatedCampaignsOk returns a tuple with the AssociatedCampaigns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SPGetAssociatedCampaignsResponse) GetAssociatedCampaignsOk() ([]AssociatedCampaign, bool) {
	if o == nil || IsNil(o.AssociatedCampaigns) {
		return nil, false
	}
	return o.AssociatedCampaigns, true
}

// HasAssociatedCampaigns returns a boolean if a field has been set.
func (o *SPGetAssociatedCampaignsResponse) HasAssociatedCampaigns() bool {
	if o != nil && !IsNil(o.AssociatedCampaigns) {
		return true
	}

	return false
}

// SetAssociatedCampaigns gets a reference to the given []AssociatedCampaign and assigns it to the AssociatedCampaigns field.
func (o *SPGetAssociatedCampaignsResponse) SetAssociatedCampaigns(v []AssociatedCampaign) {
	o.AssociatedCampaigns = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SPGetAssociatedCampaignsResponse) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SPGetAssociatedCampaignsResponse) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SPGetAssociatedCampaignsResponse) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SPGetAssociatedCampaignsResponse) SetNextToken(v string) {
	o.NextToken = &v
}

func (o SPGetAssociatedCampaignsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssociatedCampaigns) {
		toSerialize["associatedCampaigns"] = o.AssociatedCampaigns
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	return toSerialize, nil
}

type NullableSPGetAssociatedCampaignsResponse struct {
	value *SPGetAssociatedCampaignsResponse
	isSet bool
}

func (v NullableSPGetAssociatedCampaignsResponse) Get() *SPGetAssociatedCampaignsResponse {
	return v.value
}

func (v *NullableSPGetAssociatedCampaignsResponse) Set(val *SPGetAssociatedCampaignsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSPGetAssociatedCampaignsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSPGetAssociatedCampaignsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSPGetAssociatedCampaignsResponse(val *SPGetAssociatedCampaignsResponse) *NullableSPGetAssociatedCampaignsResponse {
	return &NullableSPGetAssociatedCampaignsResponse{value: val, isSet: true}
}

func (v NullableSPGetAssociatedCampaignsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSPGetAssociatedCampaignsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
