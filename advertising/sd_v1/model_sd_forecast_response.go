package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDForecastResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDForecastResponse{}

// SDForecastResponse Response to a request for SD forecasting.
type SDForecastResponse struct {
	BidOptimization *string `json:"bidOptimization,omitempty"`
	// Forecasts for campaign start date and end date. Default end date is start date plus 7 days.
	LifetimeForecasts []Forecast `json:"lifetimeForecasts,omitempty"`
	// Weekly average forecasts.
	WeeklyForecasts []Forecast `json:"weeklyForecasts,omitempty"`
	// Daily average forecasts.
	DailyForecasts []Forecast `json:"dailyForecasts,omitempty"`
	// Forecasting curves.
	Curves         []Curve         `json:"curves,omitempty"`
	ForecastStatus *ForecastStatus `json:"forecastStatus,omitempty"`
}

// NewSDForecastResponse instantiates a new SDForecastResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDForecastResponse() *SDForecastResponse {
	this := SDForecastResponse{}
	return &this
}

// NewSDForecastResponseWithDefaults instantiates a new SDForecastResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDForecastResponseWithDefaults() *SDForecastResponse {
	this := SDForecastResponse{}
	return &this
}

// GetBidOptimization returns the BidOptimization field value if set, zero value otherwise.
func (o *SDForecastResponse) GetBidOptimization() string {
	if o == nil || IsNil(o.BidOptimization) {
		var ret string
		return ret
	}
	return *o.BidOptimization
}

// GetBidOptimizationOk returns a tuple with the BidOptimization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDForecastResponse) GetBidOptimizationOk() (*string, bool) {
	if o == nil || IsNil(o.BidOptimization) {
		return nil, false
	}
	return o.BidOptimization, true
}

// HasBidOptimization returns a boolean if a field has been set.
func (o *SDForecastResponse) HasBidOptimization() bool {
	if o != nil && !IsNil(o.BidOptimization) {
		return true
	}

	return false
}

// SetBidOptimization gets a reference to the given string and assigns it to the BidOptimization field.
func (o *SDForecastResponse) SetBidOptimization(v string) {
	o.BidOptimization = &v
}

// GetLifetimeForecasts returns the LifetimeForecasts field value if set, zero value otherwise.
func (o *SDForecastResponse) GetLifetimeForecasts() []Forecast {
	if o == nil || IsNil(o.LifetimeForecasts) {
		var ret []Forecast
		return ret
	}
	return o.LifetimeForecasts
}

// GetLifetimeForecastsOk returns a tuple with the LifetimeForecasts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDForecastResponse) GetLifetimeForecastsOk() ([]Forecast, bool) {
	if o == nil || IsNil(o.LifetimeForecasts) {
		return nil, false
	}
	return o.LifetimeForecasts, true
}

// HasLifetimeForecasts returns a boolean if a field has been set.
func (o *SDForecastResponse) HasLifetimeForecasts() bool {
	if o != nil && !IsNil(o.LifetimeForecasts) {
		return true
	}

	return false
}

// SetLifetimeForecasts gets a reference to the given []Forecast and assigns it to the LifetimeForecasts field.
func (o *SDForecastResponse) SetLifetimeForecasts(v []Forecast) {
	o.LifetimeForecasts = v
}

// GetWeeklyForecasts returns the WeeklyForecasts field value if set, zero value otherwise.
func (o *SDForecastResponse) GetWeeklyForecasts() []Forecast {
	if o == nil || IsNil(o.WeeklyForecasts) {
		var ret []Forecast
		return ret
	}
	return o.WeeklyForecasts
}

// GetWeeklyForecastsOk returns a tuple with the WeeklyForecasts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDForecastResponse) GetWeeklyForecastsOk() ([]Forecast, bool) {
	if o == nil || IsNil(o.WeeklyForecasts) {
		return nil, false
	}
	return o.WeeklyForecasts, true
}

// HasWeeklyForecasts returns a boolean if a field has been set.
func (o *SDForecastResponse) HasWeeklyForecasts() bool {
	if o != nil && !IsNil(o.WeeklyForecasts) {
		return true
	}

	return false
}

// SetWeeklyForecasts gets a reference to the given []Forecast and assigns it to the WeeklyForecasts field.
func (o *SDForecastResponse) SetWeeklyForecasts(v []Forecast) {
	o.WeeklyForecasts = v
}

// GetDailyForecasts returns the DailyForecasts field value if set, zero value otherwise.
func (o *SDForecastResponse) GetDailyForecasts() []Forecast {
	if o == nil || IsNil(o.DailyForecasts) {
		var ret []Forecast
		return ret
	}
	return o.DailyForecasts
}

// GetDailyForecastsOk returns a tuple with the DailyForecasts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDForecastResponse) GetDailyForecastsOk() ([]Forecast, bool) {
	if o == nil || IsNil(o.DailyForecasts) {
		return nil, false
	}
	return o.DailyForecasts, true
}

// HasDailyForecasts returns a boolean if a field has been set.
func (o *SDForecastResponse) HasDailyForecasts() bool {
	if o != nil && !IsNil(o.DailyForecasts) {
		return true
	}

	return false
}

// SetDailyForecasts gets a reference to the given []Forecast and assigns it to the DailyForecasts field.
func (o *SDForecastResponse) SetDailyForecasts(v []Forecast) {
	o.DailyForecasts = v
}

// GetCurves returns the Curves field value if set, zero value otherwise.
func (o *SDForecastResponse) GetCurves() []Curve {
	if o == nil || IsNil(o.Curves) {
		var ret []Curve
		return ret
	}
	return o.Curves
}

// GetCurvesOk returns a tuple with the Curves field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDForecastResponse) GetCurvesOk() ([]Curve, bool) {
	if o == nil || IsNil(o.Curves) {
		return nil, false
	}
	return o.Curves, true
}

// HasCurves returns a boolean if a field has been set.
func (o *SDForecastResponse) HasCurves() bool {
	if o != nil && !IsNil(o.Curves) {
		return true
	}

	return false
}

// SetCurves gets a reference to the given []Curve and assigns it to the Curves field.
func (o *SDForecastResponse) SetCurves(v []Curve) {
	o.Curves = v
}

// GetForecastStatus returns the ForecastStatus field value if set, zero value otherwise.
func (o *SDForecastResponse) GetForecastStatus() ForecastStatus {
	if o == nil || IsNil(o.ForecastStatus) {
		var ret ForecastStatus
		return ret
	}
	return *o.ForecastStatus
}

// GetForecastStatusOk returns a tuple with the ForecastStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDForecastResponse) GetForecastStatusOk() (*ForecastStatus, bool) {
	if o == nil || IsNil(o.ForecastStatus) {
		return nil, false
	}
	return o.ForecastStatus, true
}

// HasForecastStatus returns a boolean if a field has been set.
func (o *SDForecastResponse) HasForecastStatus() bool {
	if o != nil && !IsNil(o.ForecastStatus) {
		return true
	}

	return false
}

// SetForecastStatus gets a reference to the given ForecastStatus and assigns it to the ForecastStatus field.
func (o *SDForecastResponse) SetForecastStatus(v ForecastStatus) {
	o.ForecastStatus = &v
}

func (o SDForecastResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BidOptimization) {
		toSerialize["bidOptimization"] = o.BidOptimization
	}
	if !IsNil(o.LifetimeForecasts) {
		toSerialize["lifetimeForecasts"] = o.LifetimeForecasts
	}
	if !IsNil(o.WeeklyForecasts) {
		toSerialize["weeklyForecasts"] = o.WeeklyForecasts
	}
	if !IsNil(o.DailyForecasts) {
		toSerialize["dailyForecasts"] = o.DailyForecasts
	}
	if !IsNil(o.Curves) {
		toSerialize["curves"] = o.Curves
	}
	if !IsNil(o.ForecastStatus) {
		toSerialize["forecastStatus"] = o.ForecastStatus
	}
	return toSerialize, nil
}

type NullableSDForecastResponse struct {
	value *SDForecastResponse
	isSet bool
}

func (v NullableSDForecastResponse) Get() *SDForecastResponse {
	return v.value
}

func (v *NullableSDForecastResponse) Set(val *SDForecastResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSDForecastResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSDForecastResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDForecastResponse(val *SDForecastResponse) *NullableSDForecastResponse {
	return &NullableSDForecastResponse{value: val, isSet: true}
}

func (v NullableSDForecastResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDForecastResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
