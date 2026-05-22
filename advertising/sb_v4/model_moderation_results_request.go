package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the ModerationResultsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModerationResultsRequest{}

// ModerationResultsRequest struct for ModerationResultsRequest
type ModerationResultsRequest struct {
	// Filter by specific version id of the ad. The API will return the ad's all versions moderation status if this field is empty.
	VersionIdFilter []string                       `json:"versionIdFilter,omitempty"`
	IdType          IdType                         `json:"idType"`
	AdProgramType   ModerationResultsAdProgramType `json:"adProgramType"`
	// Operations that return paginated results include a pagination token in this field. To retrieve the next page of results, call the same operation and specify this token in the request. If the `NextToken` field is empty, there are no further results.
	NextToken *string `json:"nextToken,omitempty"`
	// Sets a limit on the number of results returned by an operation.
	MaxResults int32 `json:"maxResults"`
	// Filter by specific moderation status.
	ModerationStatusFilter []ModerationStatus `json:"moderationStatusFilter,omitempty"`
	// The unique identifier of the ad which can be obtained after the ad is created using create APIs.
	Id string `json:"id"`
}

type _ModerationResultsRequest ModerationResultsRequest

// NewModerationResultsRequest instantiates a new ModerationResultsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModerationResultsRequest(idType IdType, adProgramType ModerationResultsAdProgramType, maxResults int32, id string) *ModerationResultsRequest {
	this := ModerationResultsRequest{}
	this.IdType = idType
	this.AdProgramType = adProgramType
	this.MaxResults = maxResults
	this.Id = id
	return &this
}

// NewModerationResultsRequestWithDefaults instantiates a new ModerationResultsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModerationResultsRequestWithDefaults() *ModerationResultsRequest {
	this := ModerationResultsRequest{}
	return &this
}

// GetVersionIdFilter returns the VersionIdFilter field value if set, zero value otherwise.
func (o *ModerationResultsRequest) GetVersionIdFilter() []string {
	if o == nil || IsNil(o.VersionIdFilter) {
		var ret []string
		return ret
	}
	return o.VersionIdFilter
}

// GetVersionIdFilterOk returns a tuple with the VersionIdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModerationResultsRequest) GetVersionIdFilterOk() ([]string, bool) {
	if o == nil || IsNil(o.VersionIdFilter) {
		return nil, false
	}
	return o.VersionIdFilter, true
}

// HasVersionIdFilter returns a boolean if a field has been set.
func (o *ModerationResultsRequest) HasVersionIdFilter() bool {
	if o != nil && !IsNil(o.VersionIdFilter) {
		return true
	}

	return false
}

// SetVersionIdFilter gets a reference to the given []string and assigns it to the VersionIdFilter field.
func (o *ModerationResultsRequest) SetVersionIdFilter(v []string) {
	o.VersionIdFilter = v
}

// GetIdType returns the IdType field value
func (o *ModerationResultsRequest) GetIdType() IdType {
	if o == nil {
		var ret IdType
		return ret
	}

	return o.IdType
}

// GetIdTypeOk returns a tuple with the IdType field value
// and a boolean to check if the value has been set.
func (o *ModerationResultsRequest) GetIdTypeOk() (*IdType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdType, true
}

// SetIdType sets field value
func (o *ModerationResultsRequest) SetIdType(v IdType) {
	o.IdType = v
}

// GetAdProgramType returns the AdProgramType field value
func (o *ModerationResultsRequest) GetAdProgramType() ModerationResultsAdProgramType {
	if o == nil {
		var ret ModerationResultsAdProgramType
		return ret
	}

	return o.AdProgramType
}

// GetAdProgramTypeOk returns a tuple with the AdProgramType field value
// and a boolean to check if the value has been set.
func (o *ModerationResultsRequest) GetAdProgramTypeOk() (*ModerationResultsAdProgramType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdProgramType, true
}

// SetAdProgramType sets field value
func (o *ModerationResultsRequest) SetAdProgramType(v ModerationResultsAdProgramType) {
	o.AdProgramType = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ModerationResultsRequest) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModerationResultsRequest) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ModerationResultsRequest) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ModerationResultsRequest) SetNextToken(v string) {
	o.NextToken = &v
}

// GetMaxResults returns the MaxResults field value
func (o *ModerationResultsRequest) GetMaxResults() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value
// and a boolean to check if the value has been set.
func (o *ModerationResultsRequest) GetMaxResultsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxResults, true
}

// SetMaxResults sets field value
func (o *ModerationResultsRequest) SetMaxResults(v int32) {
	o.MaxResults = v
}

// GetModerationStatusFilter returns the ModerationStatusFilter field value if set, zero value otherwise.
func (o *ModerationResultsRequest) GetModerationStatusFilter() []ModerationStatus {
	if o == nil || IsNil(o.ModerationStatusFilter) {
		var ret []ModerationStatus
		return ret
	}
	return o.ModerationStatusFilter
}

// GetModerationStatusFilterOk returns a tuple with the ModerationStatusFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModerationResultsRequest) GetModerationStatusFilterOk() ([]ModerationStatus, bool) {
	if o == nil || IsNil(o.ModerationStatusFilter) {
		return nil, false
	}
	return o.ModerationStatusFilter, true
}

// HasModerationStatusFilter returns a boolean if a field has been set.
func (o *ModerationResultsRequest) HasModerationStatusFilter() bool {
	if o != nil && !IsNil(o.ModerationStatusFilter) {
		return true
	}

	return false
}

// SetModerationStatusFilter gets a reference to the given []ModerationStatus and assigns it to the ModerationStatusFilter field.
func (o *ModerationResultsRequest) SetModerationStatusFilter(v []ModerationStatus) {
	o.ModerationStatusFilter = v
}

// GetId returns the Id field value
func (o *ModerationResultsRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ModerationResultsRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ModerationResultsRequest) SetId(v string) {
	o.Id = v
}

func (o ModerationResultsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VersionIdFilter) {
		toSerialize["versionIdFilter"] = o.VersionIdFilter
	}
	toSerialize["idType"] = o.IdType
	toSerialize["adProgramType"] = o.AdProgramType
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	toSerialize["maxResults"] = o.MaxResults
	if !IsNil(o.ModerationStatusFilter) {
		toSerialize["moderationStatusFilter"] = o.ModerationStatusFilter
	}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableModerationResultsRequest struct {
	value *ModerationResultsRequest
	isSet bool
}

func (v NullableModerationResultsRequest) Get() *ModerationResultsRequest {
	return v.value
}

func (v *NullableModerationResultsRequest) Set(val *ModerationResultsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModerationResultsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModerationResultsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModerationResultsRequest(val *ModerationResultsRequest) *NullableModerationResultsRequest {
	return &NullableModerationResultsRequest{value: val, isSet: true}
}

func (v NullableModerationResultsRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableModerationResultsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
