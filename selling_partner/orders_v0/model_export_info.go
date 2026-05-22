package orders_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the ExportInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExportInfo{}

// ExportInfo Contains information that is related to the export of an order item.
type ExportInfo struct {
	ExportCharge *Money `json:"ExportCharge,omitempty"`
	// Holds the `ExportCharge` collection model that is associated with the specified order item.\\n\\n**Possible values**: `AMAZON_FACILITATED`: Import/export charge is withheld by Amazon and remitted to the customs authority by the carrier on behalf of the buyer/seller.
	ExportChargeModel *string `json:"ExportChargeModel,omitempty"`
}

// NewExportInfo instantiates a new ExportInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportInfo() *ExportInfo {
	this := ExportInfo{}
	return &this
}

// NewExportInfoWithDefaults instantiates a new ExportInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportInfoWithDefaults() *ExportInfo {
	this := ExportInfo{}
	return &this
}

// GetExportCharge returns the ExportCharge field value if set, zero value otherwise.
func (o *ExportInfo) GetExportCharge() Money {
	if o == nil || IsNil(o.ExportCharge) {
		var ret Money
		return ret
	}
	return *o.ExportCharge
}

// GetExportChargeOk returns a tuple with the ExportCharge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportInfo) GetExportChargeOk() (*Money, bool) {
	if o == nil || IsNil(o.ExportCharge) {
		return nil, false
	}
	return o.ExportCharge, true
}

// HasExportCharge returns a boolean if a field has been set.
func (o *ExportInfo) HasExportCharge() bool {
	if o != nil && !IsNil(o.ExportCharge) {
		return true
	}

	return false
}

// SetExportCharge gets a reference to the given Money and assigns it to the ExportCharge field.
func (o *ExportInfo) SetExportCharge(v Money) {
	o.ExportCharge = &v
}

// GetExportChargeModel returns the ExportChargeModel field value if set, zero value otherwise.
func (o *ExportInfo) GetExportChargeModel() string {
	if o == nil || IsNil(o.ExportChargeModel) {
		var ret string
		return ret
	}
	return *o.ExportChargeModel
}

// GetExportChargeModelOk returns a tuple with the ExportChargeModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportInfo) GetExportChargeModelOk() (*string, bool) {
	if o == nil || IsNil(o.ExportChargeModel) {
		return nil, false
	}
	return o.ExportChargeModel, true
}

// HasExportChargeModel returns a boolean if a field has been set.
func (o *ExportInfo) HasExportChargeModel() bool {
	if o != nil && !IsNil(o.ExportChargeModel) {
		return true
	}

	return false
}

// SetExportChargeModel gets a reference to the given string and assigns it to the ExportChargeModel field.
func (o *ExportInfo) SetExportChargeModel(v string) {
	o.ExportChargeModel = &v
}

func (o ExportInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExportCharge) {
		toSerialize["ExportCharge"] = o.ExportCharge
	}
	if !IsNil(o.ExportChargeModel) {
		toSerialize["ExportChargeModel"] = o.ExportChargeModel
	}
	return toSerialize, nil
}

type NullableExportInfo struct {
	value *ExportInfo
	isSet bool
}

func (v NullableExportInfo) Get() *ExportInfo {
	return v.value
}

func (v *NullableExportInfo) Set(val *ExportInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableExportInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableExportInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportInfo(val *ExportInfo) *NullableExportInfo {
	return &NullableExportInfo{value: val, isSet: true}
}

func (v NullableExportInfo) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableExportInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
