package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the GetCampaignRecommendationsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCampaignRecommendationsResponse{}

// GetCampaignRecommendationsResponse struct for GetCampaignRecommendationsResponse
type GetCampaignRecommendationsResponse struct {
	// An identifier to fetch next set of campaign recommendations records in the result set if available. This will be null when at the end of result set.
	NextToken *string `json:"nextToken,omitempty"`
	// List of campaign recommendations.
	Recommendations []CampaignRecommendation `json:"recommendations"`
}

type _GetCampaignRecommendationsResponse GetCampaignRecommendationsResponse

// NewGetCampaignRecommendationsResponse instantiates a new GetCampaignRecommendationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCampaignRecommendationsResponse(recommendations []CampaignRecommendation) *GetCampaignRecommendationsResponse {
	this := GetCampaignRecommendationsResponse{}
	this.Recommendations = recommendations
	return &this
}

// NewGetCampaignRecommendationsResponseWithDefaults instantiates a new GetCampaignRecommendationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCampaignRecommendationsResponseWithDefaults() *GetCampaignRecommendationsResponse {
	this := GetCampaignRecommendationsResponse{}
	return &this
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *GetCampaignRecommendationsResponse) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCampaignRecommendationsResponse) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *GetCampaignRecommendationsResponse) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *GetCampaignRecommendationsResponse) SetNextToken(v string) {
	o.NextToken = &v
}

// GetRecommendations returns the Recommendations field value
func (o *GetCampaignRecommendationsResponse) GetRecommendations() []CampaignRecommendation {
	if o == nil {
		var ret []CampaignRecommendation
		return ret
	}

	return o.Recommendations
}

// GetRecommendationsOk returns a tuple with the Recommendations field value
// and a boolean to check if the value has been set.
func (o *GetCampaignRecommendationsResponse) GetRecommendationsOk() ([]CampaignRecommendation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Recommendations, true
}

// SetRecommendations sets field value
func (o *GetCampaignRecommendationsResponse) SetRecommendations(v []CampaignRecommendation) {
	o.Recommendations = v
}

func (o GetCampaignRecommendationsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	toSerialize["recommendations"] = o.Recommendations
	return toSerialize, nil
}

type NullableGetCampaignRecommendationsResponse struct {
	value *GetCampaignRecommendationsResponse
	isSet bool
}

func (v NullableGetCampaignRecommendationsResponse) Get() *GetCampaignRecommendationsResponse {
	return v.value
}

func (v *NullableGetCampaignRecommendationsResponse) Set(val *GetCampaignRecommendationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCampaignRecommendationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCampaignRecommendationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCampaignRecommendationsResponse(val *GetCampaignRecommendationsResponse) *NullableGetCampaignRecommendationsResponse {
	return &NullableGetCampaignRecommendationsResponse{value: val, isSet: true}
}

func (v NullableGetCampaignRecommendationsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetCampaignRecommendationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
