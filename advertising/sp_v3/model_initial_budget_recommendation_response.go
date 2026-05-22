package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the InitialBudgetRecommendationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InitialBudgetRecommendationResponse{}

// InitialBudgetRecommendationResponse struct for InitialBudgetRecommendationResponse
type InitialBudgetRecommendationResponse struct {
	// A list of special events around the start and end date of the campaign.
	SpecialEvents []SpecialEvent `json:"specialEvents"`
	// Recommended daily budget for the new campaign. Note: value -1 means we don’t have enough information to provide a recommendation.
	DailyBudget float32 `json:"dailyBudget"`
	// Unique identifier for each recommendation.
	RecommendationId *string   `json:"recommendationId,omitempty"`
	Benchmark        Benchmark `json:"benchmark"`
}

type _InitialBudgetRecommendationResponse InitialBudgetRecommendationResponse

// NewInitialBudgetRecommendationResponse instantiates a new InitialBudgetRecommendationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInitialBudgetRecommendationResponse(specialEvents []SpecialEvent, dailyBudget float32, benchmark Benchmark) *InitialBudgetRecommendationResponse {
	this := InitialBudgetRecommendationResponse{}
	this.SpecialEvents = specialEvents
	this.DailyBudget = dailyBudget
	this.Benchmark = benchmark
	return &this
}

// NewInitialBudgetRecommendationResponseWithDefaults instantiates a new InitialBudgetRecommendationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInitialBudgetRecommendationResponseWithDefaults() *InitialBudgetRecommendationResponse {
	this := InitialBudgetRecommendationResponse{}
	return &this
}

// GetSpecialEvents returns the SpecialEvents field value
func (o *InitialBudgetRecommendationResponse) GetSpecialEvents() []SpecialEvent {
	if o == nil {
		var ret []SpecialEvent
		return ret
	}

	return o.SpecialEvents
}

// GetSpecialEventsOk returns a tuple with the SpecialEvents field value
// and a boolean to check if the value has been set.
func (o *InitialBudgetRecommendationResponse) GetSpecialEventsOk() ([]SpecialEvent, bool) {
	if o == nil {
		return nil, false
	}
	return o.SpecialEvents, true
}

// SetSpecialEvents sets field value
func (o *InitialBudgetRecommendationResponse) SetSpecialEvents(v []SpecialEvent) {
	o.SpecialEvents = v
}

// GetDailyBudget returns the DailyBudget field value
func (o *InitialBudgetRecommendationResponse) GetDailyBudget() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.DailyBudget
}

// GetDailyBudgetOk returns a tuple with the DailyBudget field value
// and a boolean to check if the value has been set.
func (o *InitialBudgetRecommendationResponse) GetDailyBudgetOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DailyBudget, true
}

// SetDailyBudget sets field value
func (o *InitialBudgetRecommendationResponse) SetDailyBudget(v float32) {
	o.DailyBudget = v
}

// GetRecommendationId returns the RecommendationId field value if set, zero value otherwise.
func (o *InitialBudgetRecommendationResponse) GetRecommendationId() string {
	if o == nil || IsNil(o.RecommendationId) {
		var ret string
		return ret
	}
	return *o.RecommendationId
}

// GetRecommendationIdOk returns a tuple with the RecommendationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InitialBudgetRecommendationResponse) GetRecommendationIdOk() (*string, bool) {
	if o == nil || IsNil(o.RecommendationId) {
		return nil, false
	}
	return o.RecommendationId, true
}

// HasRecommendationId returns a boolean if a field has been set.
func (o *InitialBudgetRecommendationResponse) HasRecommendationId() bool {
	if o != nil && !IsNil(o.RecommendationId) {
		return true
	}

	return false
}

// SetRecommendationId gets a reference to the given string and assigns it to the RecommendationId field.
func (o *InitialBudgetRecommendationResponse) SetRecommendationId(v string) {
	o.RecommendationId = &v
}

// GetBenchmark returns the Benchmark field value
func (o *InitialBudgetRecommendationResponse) GetBenchmark() Benchmark {
	if o == nil {
		var ret Benchmark
		return ret
	}

	return o.Benchmark
}

// GetBenchmarkOk returns a tuple with the Benchmark field value
// and a boolean to check if the value has been set.
func (o *InitialBudgetRecommendationResponse) GetBenchmarkOk() (*Benchmark, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Benchmark, true
}

// SetBenchmark sets field value
func (o *InitialBudgetRecommendationResponse) SetBenchmark(v Benchmark) {
	o.Benchmark = v
}

func (o InitialBudgetRecommendationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["specialEvents"] = o.SpecialEvents
	toSerialize["dailyBudget"] = o.DailyBudget
	if !IsNil(o.RecommendationId) {
		toSerialize["recommendationId"] = o.RecommendationId
	}
	toSerialize["benchmark"] = o.Benchmark
	return toSerialize, nil
}

type NullableInitialBudgetRecommendationResponse struct {
	value *InitialBudgetRecommendationResponse
	isSet bool
}

func (v NullableInitialBudgetRecommendationResponse) Get() *InitialBudgetRecommendationResponse {
	return v.value
}

func (v *NullableInitialBudgetRecommendationResponse) Set(val *InitialBudgetRecommendationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInitialBudgetRecommendationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInitialBudgetRecommendationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInitialBudgetRecommendationResponse(val *InitialBudgetRecommendationResponse) *NullableInitialBudgetRecommendationResponse {
	return &NullableInitialBudgetRecommendationResponse{value: val, isSet: true}
}

func (v NullableInitialBudgetRecommendationResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableInitialBudgetRecommendationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
