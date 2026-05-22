package invoices_api_model_20240619

import (
	"github.com/bytedance/sonic"
)

// checks if the GetInvoicesExportsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetInvoicesExportsResponse{}

// GetInvoicesExportsResponse Success.
type GetInvoicesExportsResponse struct {
	// A list of exports.
	Exports []Export `json:"exports,omitempty"`
	// This token is returned when the number of results exceeds the specified `pageSize` value. To get the next page of results, call the `getInvoices` operation and include this token with the previous call parameters.
	NextToken *string `json:"nextToken,omitempty"`
}

// NewGetInvoicesExportsResponse instantiates a new GetInvoicesExportsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetInvoicesExportsResponse() *GetInvoicesExportsResponse {
	this := GetInvoicesExportsResponse{}
	return &this
}

// NewGetInvoicesExportsResponseWithDefaults instantiates a new GetInvoicesExportsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetInvoicesExportsResponseWithDefaults() *GetInvoicesExportsResponse {
	this := GetInvoicesExportsResponse{}
	return &this
}

// GetExports returns the Exports field value if set, zero value otherwise.
func (o *GetInvoicesExportsResponse) GetExports() []Export {
	if o == nil || IsNil(o.Exports) {
		var ret []Export
		return ret
	}
	return o.Exports
}

// GetExportsOk returns a tuple with the Exports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInvoicesExportsResponse) GetExportsOk() ([]Export, bool) {
	if o == nil || IsNil(o.Exports) {
		return nil, false
	}
	return o.Exports, true
}

// HasExports returns a boolean if a field has been set.
func (o *GetInvoicesExportsResponse) HasExports() bool {
	if o != nil && !IsNil(o.Exports) {
		return true
	}

	return false
}

// SetExports gets a reference to the given []Export and assigns it to the Exports field.
func (o *GetInvoicesExportsResponse) SetExports(v []Export) {
	o.Exports = v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *GetInvoicesExportsResponse) GetNextToken() string {
	if o == nil || IsNil(o.NextToken) {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInvoicesExportsResponse) GetNextTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextToken) {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *GetInvoicesExportsResponse) HasNextToken() bool {
	if o != nil && !IsNil(o.NextToken) {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *GetInvoicesExportsResponse) SetNextToken(v string) {
	o.NextToken = &v
}

func (o GetInvoicesExportsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Exports) {
		toSerialize["exports"] = o.Exports
	}
	if !IsNil(o.NextToken) {
		toSerialize["nextToken"] = o.NextToken
	}
	return toSerialize, nil
}

type NullableGetInvoicesExportsResponse struct {
	value *GetInvoicesExportsResponse
	isSet bool
}

func (v NullableGetInvoicesExportsResponse) Get() *GetInvoicesExportsResponse {
	return v.value
}

func (v *NullableGetInvoicesExportsResponse) Set(val *GetInvoicesExportsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInvoicesExportsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInvoicesExportsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInvoicesExportsResponse(val *GetInvoicesExportsResponse) *NullableGetInvoicesExportsResponse {
	return &NullableGetInvoicesExportsResponse{value: val, isSet: true}
}

func (v NullableGetInvoicesExportsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetInvoicesExportsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
