package invoices_api_model_20240619

import (
	"github.com/bytedance/sonic"
)

// checks if the ExportInvoicesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExportInvoicesResponse{}

// ExportInvoicesResponse Success.
type ExportInvoicesResponse struct {
	// The export identifier.
	ExportId *string `json:"exportId,omitempty"`
}

// NewExportInvoicesResponse instantiates a new ExportInvoicesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportInvoicesResponse() *ExportInvoicesResponse {
	this := ExportInvoicesResponse{}
	return &this
}

// NewExportInvoicesResponseWithDefaults instantiates a new ExportInvoicesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportInvoicesResponseWithDefaults() *ExportInvoicesResponse {
	this := ExportInvoicesResponse{}
	return &this
}

// GetExportId returns the ExportId field value if set, zero value otherwise.
func (o *ExportInvoicesResponse) GetExportId() string {
	if o == nil || IsNil(o.ExportId) {
		var ret string
		return ret
	}
	return *o.ExportId
}

// GetExportIdOk returns a tuple with the ExportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportInvoicesResponse) GetExportIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExportId) {
		return nil, false
	}
	return o.ExportId, true
}

// HasExportId returns a boolean if a field has been set.
func (o *ExportInvoicesResponse) HasExportId() bool {
	if o != nil && !IsNil(o.ExportId) {
		return true
	}

	return false
}

// SetExportId gets a reference to the given string and assigns it to the ExportId field.
func (o *ExportInvoicesResponse) SetExportId(v string) {
	o.ExportId = &v
}

func (o ExportInvoicesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExportId) {
		toSerialize["exportId"] = o.ExportId
	}
	return toSerialize, nil
}

type NullableExportInvoicesResponse struct {
	value *ExportInvoicesResponse
	isSet bool
}

func (v NullableExportInvoicesResponse) Get() *ExportInvoicesResponse {
	return v.value
}

func (v *NullableExportInvoicesResponse) Set(val *ExportInvoicesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableExportInvoicesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableExportInvoicesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportInvoicesResponse(val *ExportInvoicesResponse) *NullableExportInvoicesResponse {
	return &NullableExportInvoicesResponse{value: val, isSet: true}
}

func (v NullableExportInvoicesResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableExportInvoicesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
