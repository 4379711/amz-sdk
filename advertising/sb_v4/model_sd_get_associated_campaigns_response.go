package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SDGetAssociatedCampaignsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDGetAssociatedCampaignsResponse{}

// SDGetAssociatedCampaignsResponse struct for SDGetAssociatedCampaignsResponse
type SDGetAssociatedCampaignsResponse struct {
	// A list of campaigns that are associated to this budget rule.
	AssociatedCampaigns []AssociatedCampaign `json:"associatedCampaigns,omitempty"`
	// To retrieve the next page of results, call the same operation and specify this token in the request. If the `nextToken` field is empty, there are no further results.
	NextToken *string `json:"nextToken,omitempty"`
}

// NewSDGetAssociatedCampaignsResponse instantiates a new SDGetAssociatedCampaignsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDGetAssociatedCampaignsResponse() *SDGetAssociatedCampaignsResponse {
	this := SDGetAssociatedCampaignsResponse{}
	return &this
}

// NewSDGetAssociatedCampaignsResponseWithDefaults instantiates a new SDGetAssociatedCampaignsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDGetAssociatedCampaignsResponseWithDefaults() *SDGetAssociatedCampaignsResponse {
	this := SDGetAssociatedCampaignsResponse{}
	return &this
}

// GetAssociatedCampaigns returns the AssociatedCampaigns field value if set, zero value otherwise.
func (o *SDGetAssociatedCampaignsResponse) GetAssociatedCampaigns() []AssociatedCampaign {
	if o == nil || IsNil(o.AssociatedCampaigns) {
		var ret []AssociatedCampaign
		return ret
	}
	return o.AssociatedCampaigns
}

// GetAssociatedCampaignsOk returns a tuple with the AssociatedCampaigns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDGetAssociatedCampaignsResponse) GetAssociatedCampaignsOk() ([]AssociatedCampaign, bool) {
	if o == nil || IsNil(o.AssociatedCampaigns) {
		return nil, false
	}
	return o.AssociatedCampaigns, true
}

// HasAssociatedCampaigns returns a boolean if a field has been set.
func (o *SDGetAssociatedCampaignsResponse) HasAssociatedCampaigns() bool {
	if o != nil && !IsNil(o.AssociatedCampaigns) {
		return true
	}

	return false
}

// SetAssociatedCampaigns gets a reference to the given []AssociatedCampaign and assigns it to the AssociatedCampaigns field.
func (o *SDGetAssociatedCampaignsResponse) SetAssociatedCampaigns(v []AssociatedCampaign) {
	o.AssociatedCampaigns = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *SDGetAssociatedCampaignsResponse) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDGetAssociatedCampaignsResponse) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *SDGetAssociatedCampaignsResponse) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *SDGetAssociatedCampaignsResponse) SetNextToken(v string) {
	o.NextToken = &v
}

func (o SDGetAssociatedCampaignsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssociatedCampaigns) {
		toSerialize["associatedCampaigns"] = o.AssociatedCampaigns
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	return toSerialize, nil
}

type NullableSDGetAssociatedCampaignsResponse struct {
	value *SDGetAssociatedCampaignsResponse
	isSet bool
}

func (v NullableSDGetAssociatedCampaignsResponse) Get() *SDGetAssociatedCampaignsResponse {
	return v.value
}

func (v *NullableSDGetAssociatedCampaignsResponse) Set(val *SDGetAssociatedCampaignsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSDGetAssociatedCampaignsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSDGetAssociatedCampaignsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDGetAssociatedCampaignsResponse(val *SDGetAssociatedCampaignsResponse) *NullableSDGetAssociatedCampaignsResponse {
	return &NullableSDGetAssociatedCampaignsResponse{value: val, isSet: true}
}

func (v NullableSDGetAssociatedCampaignsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDGetAssociatedCampaignsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
