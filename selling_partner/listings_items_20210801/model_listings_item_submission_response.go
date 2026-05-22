package listings_items_20210801

import (
	"github.com/bytedance/sonic"
)

// checks if the ListingsItemSubmissionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListingsItemSubmissionResponse{}

// ListingsItemSubmissionResponse Response containing the results of a submission to the Selling Partner API for Listings Items.
type ListingsItemSubmissionResponse struct {
	// A selling partner provided identifier for an Amazon listing.
	Sku string `json:"sku"`
	// The status of the listings item submission.
	Status string `json:"status"`
	// The unique identifier of the listings item submission.
	SubmissionId string `json:"submissionId"`
	// Listings item issues related to the listings item submission.
	Issues []Issue `json:"issues,omitempty"`
	// Identity attributes associated with the item in the Amazon catalog, such as the ASIN.
	Identifiers []ItemIdentifiersByMarketplace `json:"identifiers,omitempty"`
}

type _ListingsItemSubmissionResponse ListingsItemSubmissionResponse

// NewListingsItemSubmissionResponse instantiates a new ListingsItemSubmissionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListingsItemSubmissionResponse(sku string, status string, submissionId string) *ListingsItemSubmissionResponse {
	this := ListingsItemSubmissionResponse{}
	this.Sku = sku
	this.Status = status
	this.SubmissionId = submissionId
	return &this
}

// NewListingsItemSubmissionResponseWithDefaults instantiates a new ListingsItemSubmissionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListingsItemSubmissionResponseWithDefaults() *ListingsItemSubmissionResponse {
	this := ListingsItemSubmissionResponse{}
	return &this
}

// GetSku returns the Sku field value
func (o *ListingsItemSubmissionResponse) GetSku() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sku
}

// GetSkuOk returns a tuple with the Sku field value
// and a boolean to check if the value has been set.
func (o *ListingsItemSubmissionResponse) GetSkuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sku, true
}

// SetSku sets field value
func (o *ListingsItemSubmissionResponse) SetSku(v string) {
	o.Sku = v
}

// GetStatus returns the Status field value
func (o *ListingsItemSubmissionResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ListingsItemSubmissionResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ListingsItemSubmissionResponse) SetStatus(v string) {
	o.Status = v
}

// GetSubmissionId returns the SubmissionId field value
func (o *ListingsItemSubmissionResponse) GetSubmissionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubmissionId
}

// GetSubmissionIdOk returns a tuple with the SubmissionId field value
// and a boolean to check if the value has been set.
func (o *ListingsItemSubmissionResponse) GetSubmissionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubmissionId, true
}

// SetSubmissionId sets field value
func (o *ListingsItemSubmissionResponse) SetSubmissionId(v string) {
	o.SubmissionId = v
}

// GetIssues returns the Issues field value if set, zero value otherwise.
func (o *ListingsItemSubmissionResponse) GetIssues() []Issue {
	if o == nil || IsNil(o.Issues) {
		var ret []Issue
		return ret
	}
	return o.Issues
}

// GetIssuesOk returns a tuple with the Issues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingsItemSubmissionResponse) GetIssuesOk() ([]Issue, bool) {
	if o == nil || IsNil(o.Issues) {
		return nil, false
	}
	return o.Issues, true
}

// HasIssues returns a boolean if a field has been set.
func (o *ListingsItemSubmissionResponse) HasIssues() bool {
	if o != nil && !IsNil(o.Issues) {
		return true
	}

	return false
}

// SetIssues gets a reference to the given []Issue and assigns it to the Issues field.
func (o *ListingsItemSubmissionResponse) SetIssues(v []Issue) {
	o.Issues = v
}

// GetIdentifiers returns the Identifiers field value if set, zero value otherwise.
func (o *ListingsItemSubmissionResponse) GetIdentifiers() []ItemIdentifiersByMarketplace {
	if o == nil || IsNil(o.Identifiers) {
		var ret []ItemIdentifiersByMarketplace
		return ret
	}
	return o.Identifiers
}

// GetIdentifiersOk returns a tuple with the Identifiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingsItemSubmissionResponse) GetIdentifiersOk() ([]ItemIdentifiersByMarketplace, bool) {
	if o == nil || IsNil(o.Identifiers) {
		return nil, false
	}
	return o.Identifiers, true
}

// HasIdentifiers returns a boolean if a field has been set.
func (o *ListingsItemSubmissionResponse) HasIdentifiers() bool {
	if o != nil && !IsNil(o.Identifiers) {
		return true
	}

	return false
}

// SetIdentifiers gets a reference to the given []ItemIdentifiersByMarketplace and assigns it to the Identifiers field.
func (o *ListingsItemSubmissionResponse) SetIdentifiers(v []ItemIdentifiersByMarketplace) {
	o.Identifiers = v
}

func (o ListingsItemSubmissionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sku"] = o.Sku
	toSerialize["status"] = o.Status
	toSerialize["submissionId"] = o.SubmissionId
	if !IsNil(o.Issues) {
		toSerialize["issues"] = o.Issues
	}
	if !IsNil(o.Identifiers) {
		toSerialize["identifiers"] = o.Identifiers
	}
	return toSerialize, nil
}

type NullableListingsItemSubmissionResponse struct {
	value *ListingsItemSubmissionResponse
	isSet bool
}

func (v NullableListingsItemSubmissionResponse) Get() *ListingsItemSubmissionResponse {
	return v.value
}

func (v *NullableListingsItemSubmissionResponse) Set(val *ListingsItemSubmissionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListingsItemSubmissionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListingsItemSubmissionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListingsItemSubmissionResponse(val *ListingsItemSubmissionResponse) *NullableListingsItemSubmissionResponse {
	return &NullableListingsItemSubmissionResponse{value: val, isSet: true}
}

func (v NullableListingsItemSubmissionResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListingsItemSubmissionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
