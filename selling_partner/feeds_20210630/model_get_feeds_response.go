package feeds_20210630

import (
	"github.com/bytedance/sonic"
)

// checks if the GetFeedsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFeedsResponse{}

// GetFeedsResponse Response schema.
type GetFeedsResponse struct {
	// A list of feeds.
	Feeds []Feed `json:"feeds"`
	// Returned when the number of results exceeds pageSize. To get the next page of results, call the getFeeds operation with this token as the only parameter.
	NextToken *string `json:"nextToken,omitempty"`
}

type _GetFeedsResponse GetFeedsResponse

// NewGetFeedsResponse instantiates a new GetFeedsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFeedsResponse(feeds []Feed) *GetFeedsResponse {
	this := GetFeedsResponse{}
	this.Feeds = feeds
	return &this
}

// NewGetFeedsResponseWithDefaults instantiates a new GetFeedsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFeedsResponseWithDefaults() *GetFeedsResponse {
	this := GetFeedsResponse{}
	return &this
}

// GetFeeds returns the Feeds field value
func (o *GetFeedsResponse) GetFeeds() []Feed {
	if o == nil {
		var ret []Feed
		return ret
	}

	return o.Feeds
}

// GetFeedsOk returns a tuple with the Feeds field value
// and a boolean to check if the value has been set.
func (o *GetFeedsResponse) GetFeedsOk() ([]Feed, bool) {
	if o == nil {
		return nil, false
	}
	return o.Feeds, true
}

// SetFeeds sets field value
func (o *GetFeedsResponse) SetFeeds(v []Feed) {
	o.Feeds = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *GetFeedsResponse) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFeedsResponse) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *GetFeedsResponse) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *GetFeedsResponse) SetNextToken(v string) {
	o.NextToken = &v
}

func (o GetFeedsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["feeds"] = o.Feeds
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	return toSerialize, nil
}

type NullableGetFeedsResponse struct {
	value *GetFeedsResponse
	isSet bool
}

func (v NullableGetFeedsResponse) Get() *GetFeedsResponse {
	return v.value
}

func (v *NullableGetFeedsResponse) Set(val *GetFeedsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFeedsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFeedsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFeedsResponse(val *GetFeedsResponse) *NullableGetFeedsResponse {
	return &NullableGetFeedsResponse{value: val, isSet: true}
}

func (v NullableGetFeedsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetFeedsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
