package promotions_20251201

import (
	"github.com/bytedance/sonic"
)

// checks if the SearchPromotionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchPromotionsResponse{}

// SearchPromotionsResponse The response schema for the `searchPromotions` operation. **Note:** The `selectionDetails` field is not present in the selection object within promotion summaries.
type SearchPromotionsResponse struct {
	// The total number of promotions matching the search criteria, across all pages. This count remains consistent across paginated requests. **Note:** In rare cases, individual records cannot be returned and are omitted from the response. When this happens, a page may contain fewer items than expected, and the combined number of items across all pages may be less than `totalResults`. The request itself still completes successfully.
	TotalResults int32       `json:"totalResults"`
	Pagination   *Pagination `json:"pagination,omitempty"`
	// A list of promotion summaries that matches the search criteria.
	Promotions []PromotionSummary `json:"promotions"`
}

// NewSearchPromotionsResponse instantiates a new SearchPromotionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchPromotionsResponse(totalResults int32, promotions []PromotionSummary) *SearchPromotionsResponse {
	this := SearchPromotionsResponse{}
	this.TotalResults = totalResults
	this.Promotions = promotions
	return &this
}

// NewSearchPromotionsResponseWithDefaults instantiates a new SearchPromotionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchPromotionsResponseWithDefaults() *SearchPromotionsResponse {
	this := SearchPromotionsResponse{}
	return &this
}

// GetTotalResults returns the TotalResults field value
func (o *SearchPromotionsResponse) GetTotalResults() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value
// and a boolean to check if the value has been set.
func (o *SearchPromotionsResponse) GetTotalResultsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalResults, true
}

// SetTotalResults sets field value
func (o *SearchPromotionsResponse) SetTotalResults(v int32) {
	o.TotalResults = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *SearchPromotionsResponse) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchPromotionsResponse) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *SearchPromotionsResponse) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *SearchPromotionsResponse) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetPromotions returns the Promotions field value
func (o *SearchPromotionsResponse) GetPromotions() []PromotionSummary {
	if o == nil {
		var ret []PromotionSummary
		return ret
	}

	return o.Promotions
}

// GetPromotionsOk returns a tuple with the Promotions field value
// and a boolean to check if the value has been set.
func (o *SearchPromotionsResponse) GetPromotionsOk() ([]PromotionSummary, bool) {
	if o == nil {
		return nil, false
	}
	return o.Promotions, true
}

// SetPromotions sets field value
func (o *SearchPromotionsResponse) SetPromotions(v []PromotionSummary) {
	o.Promotions = v
}

func (o SearchPromotionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["totalResults"] = o.TotalResults
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	toSerialize["promotions"] = o.Promotions
	return toSerialize, nil
}

type NullableSearchPromotionsResponse struct {
	value *SearchPromotionsResponse
	isSet bool
}

func (v NullableSearchPromotionsResponse) Get() *SearchPromotionsResponse {
	return v.value
}

func (v *NullableSearchPromotionsResponse) Set(val *SearchPromotionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchPromotionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchPromotionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchPromotionsResponse(val *SearchPromotionsResponse) *NullableSearchPromotionsResponse {
	return &NullableSearchPromotionsResponse{value: val, isSet: true}
}

func (v NullableSearchPromotionsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSearchPromotionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
