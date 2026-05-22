package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the HeadlineSuggestionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HeadlineSuggestionRequest{}

// HeadlineSuggestionRequest Request structure of headline suggestion API.
type HeadlineSuggestionRequest struct {
	// An array of ASINs associated with the creative. Note do not pass an empty array, this results in an error.
	Asins []string `json:"asins,omitempty"`
	// An array of Store Pages associated with SB Spotlight Creative.
	StorePages []StorePage `json:"storePages,omitempty"`
	// Maximum number of suggestions that API should return. Response will [0, maxNumSuggestions] suggestions (suggestions are not guaranteed).
	MaxNumSuggestions *float32 `json:"maxNumSuggestions,omitempty"`
	AdFormat          *string  `json:"adFormat,omitempty"`
}

// NewHeadlineSuggestionRequest instantiates a new HeadlineSuggestionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeadlineSuggestionRequest() *HeadlineSuggestionRequest {
	this := HeadlineSuggestionRequest{}
	return &this
}

// NewHeadlineSuggestionRequestWithDefaults instantiates a new HeadlineSuggestionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeadlineSuggestionRequestWithDefaults() *HeadlineSuggestionRequest {
	this := HeadlineSuggestionRequest{}
	return &this
}

// GetAsins returns the Asins field value if set, zero value otherwise.
func (o *HeadlineSuggestionRequest) GetAsins() []string {
	if o == nil || IsNil(o.Asins) {
		var ret []string
		return ret
	}
	return o.Asins
}

// GetAsinsOk returns a tuple with the Asins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeadlineSuggestionRequest) GetAsinsOk() ([]string, bool) {
	if o == nil || IsNil(o.Asins) {
		return nil, false
	}
	return o.Asins, true
}

// HasAsins returns a boolean if a field has been set.
func (o *HeadlineSuggestionRequest) HasAsins() bool {
	if o != nil && !IsNil(o.Asins) {
		return true
	}

	return false
}

// SetAsins gets a reference to the given []string and assigns it to the Asins field.
func (o *HeadlineSuggestionRequest) SetAsins(v []string) {
	o.Asins = v
}

// GetStorePages returns the StorePages field value if set, zero value otherwise.
func (o *HeadlineSuggestionRequest) GetStorePages() []StorePage {
	if o == nil || IsNil(o.StorePages) {
		var ret []StorePage
		return ret
	}
	return o.StorePages
}

// GetStorePagesOk returns a tuple with the StorePages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeadlineSuggestionRequest) GetStorePagesOk() ([]StorePage, bool) {
	if o == nil || IsNil(o.StorePages) {
		return nil, false
	}
	return o.StorePages, true
}

// HasStorePages returns a boolean if a field has been set.
func (o *HeadlineSuggestionRequest) HasStorePages() bool {
	if o != nil && !IsNil(o.StorePages) {
		return true
	}

	return false
}

// SetStorePages gets a reference to the given []StorePage and assigns it to the StorePages field.
func (o *HeadlineSuggestionRequest) SetStorePages(v []StorePage) {
	o.StorePages = v
}

// GetMaxNumSuggestions returns the MaxNumSuggestions field value if set, zero value otherwise.
func (o *HeadlineSuggestionRequest) GetMaxNumSuggestions() float32 {
	if o == nil || IsNil(o.MaxNumSuggestions) {
		var ret float32
		return ret
	}
	return *o.MaxNumSuggestions
}

// GetMaxNumSuggestionsOk returns a tuple with the MaxNumSuggestions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeadlineSuggestionRequest) GetMaxNumSuggestionsOk() (*float32, bool) {
	if o == nil || IsNil(o.MaxNumSuggestions) {
		return nil, false
	}
	return o.MaxNumSuggestions, true
}

// HasMaxNumSuggestions returns a boolean if a field has been set.
func (o *HeadlineSuggestionRequest) HasMaxNumSuggestions() bool {
	if o != nil && !IsNil(o.MaxNumSuggestions) {
		return true
	}

	return false
}

// SetMaxNumSuggestions gets a reference to the given float32 and assigns it to the MaxNumSuggestions field.
func (o *HeadlineSuggestionRequest) SetMaxNumSuggestions(v float32) {
	o.MaxNumSuggestions = &v
}

// GetAdFormat returns the AdFormat field value if set, zero value otherwise.
func (o *HeadlineSuggestionRequest) GetAdFormat() string {
	if o == nil || IsNil(o.AdFormat) {
		var ret string
		return ret
	}
	return *o.AdFormat
}

// GetAdFormatOk returns a tuple with the AdFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HeadlineSuggestionRequest) GetAdFormatOk() (*string, bool) {
	if o == nil || IsNil(o.AdFormat) {
		return nil, false
	}
	return o.AdFormat, true
}

// HasAdFormat returns a boolean if a field has been set.
func (o *HeadlineSuggestionRequest) HasAdFormat() bool {
	if o != nil && !IsNil(o.AdFormat) {
		return true
	}

	return false
}

// SetAdFormat gets a reference to the given string and assigns it to the AdFormat field.
func (o *HeadlineSuggestionRequest) SetAdFormat(v string) {
	o.AdFormat = &v
}

func (o HeadlineSuggestionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Asins) {
		toSerialize["asins"] = o.Asins
	}
	if !IsNil(o.StorePages) {
		toSerialize["storePages"] = o.StorePages
	}
	if !IsNil(o.MaxNumSuggestions) {
		toSerialize["maxNumSuggestions"] = o.MaxNumSuggestions
	}
	if !IsNil(o.AdFormat) {
		toSerialize["adFormat"] = o.AdFormat
	}
	return toSerialize, nil
}

type NullableHeadlineSuggestionRequest struct {
	value *HeadlineSuggestionRequest
	isSet bool
}

func (v NullableHeadlineSuggestionRequest) Get() *HeadlineSuggestionRequest {
	return v.value
}

func (v *NullableHeadlineSuggestionRequest) Set(val *HeadlineSuggestionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableHeadlineSuggestionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableHeadlineSuggestionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeadlineSuggestionRequest(val *HeadlineSuggestionRequest) *NullableHeadlineSuggestionRequest {
	return &NullableHeadlineSuggestionRequest{value: val, isSet: true}
}

func (v NullableHeadlineSuggestionRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableHeadlineSuggestionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
