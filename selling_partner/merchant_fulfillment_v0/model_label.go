package merchant_fulfillment_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the Label type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Label{}

// Label Data for creating a shipping label and dimensions for printing the label.
type Label struct {
	// Custom text to print on the label. Note: Custom text is only included on labels that are in ZPL format (ZPL203). FedEx does not support `CustomTextForLabel`.
	CustomTextForLabel *string             `json:"CustomTextForLabel,omitempty"`
	Dimensions         LabelDimensions     `json:"Dimensions"`
	FileContents       FileContents        `json:"FileContents"`
	LabelFormat        *LabelFormat        `json:"LabelFormat,omitempty"`
	StandardIdForLabel *StandardIdForLabel `json:"StandardIdForLabel,omitempty"`
}

type _Label Label

// NewLabel instantiates a new Label object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabel(dimensions LabelDimensions, fileContents FileContents) *Label {
	this := Label{}
	this.Dimensions = dimensions
	this.FileContents = fileContents
	return &this
}

// NewLabelWithDefaults instantiates a new Label object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabelWithDefaults() *Label {
	this := Label{}
	return &this
}

// GetCustomTextForLabel returns the CustomTextForLabel field value if set, zero value otherwise.
func (o *Label) GetCustomTextForLabel() string {
	if o == nil || IsNil(o.CustomTextForLabel) {
		var ret string
		return ret
	}
	return *o.CustomTextForLabel
}

// GetCustomTextForLabelOk returns a tuple with the CustomTextForLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Label) GetCustomTextForLabelOk() (*string, bool) {
	if o == nil || IsNil(o.CustomTextForLabel) {
		return nil, false
	}
	return o.CustomTextForLabel, true
}

// HasCustomTextForLabel returns a boolean if a field has been set.
func (o *Label) HasCustomTextForLabel() bool {
	if o != nil && !IsNil(o.CustomTextForLabel) {
		return true
	}

	return false
}

// SetCustomTextForLabel gets a reference to the given string and assigns it to the CustomTextForLabel field.
func (o *Label) SetCustomTextForLabel(v string) {
	o.CustomTextForLabel = &v
}

// GetDimensions returns the Dimensions field value
func (o *Label) GetDimensions() LabelDimensions {
	if o == nil {
		var ret LabelDimensions
		return ret
	}

	return o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value
// and a boolean to check if the value has been set.
func (o *Label) GetDimensionsOk() (*LabelDimensions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dimensions, true
}

// SetDimensions sets field value
func (o *Label) SetDimensions(v LabelDimensions) {
	o.Dimensions = v
}

// GetFileContents returns the FileContents field value
func (o *Label) GetFileContents() FileContents {
	if o == nil {
		var ret FileContents
		return ret
	}

	return o.FileContents
}

// GetFileContentsOk returns a tuple with the FileContents field value
// and a boolean to check if the value has been set.
func (o *Label) GetFileContentsOk() (*FileContents, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileContents, true
}

// SetFileContents sets field value
func (o *Label) SetFileContents(v FileContents) {
	o.FileContents = v
}

// GetLabelFormat returns the LabelFormat field value if set, zero value otherwise.
func (o *Label) GetLabelFormat() LabelFormat {
	if o == nil || IsNil(o.LabelFormat) {
		var ret LabelFormat
		return ret
	}
	return *o.LabelFormat
}

// GetLabelFormatOk returns a tuple with the LabelFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Label) GetLabelFormatOk() (*LabelFormat, bool) {
	if o == nil || IsNil(o.LabelFormat) {
		return nil, false
	}
	return o.LabelFormat, true
}

// HasLabelFormat returns a boolean if a field has been set.
func (o *Label) HasLabelFormat() bool {
	if o != nil && !IsNil(o.LabelFormat) {
		return true
	}

	return false
}

// SetLabelFormat gets a reference to the given LabelFormat and assigns it to the LabelFormat field.
func (o *Label) SetLabelFormat(v LabelFormat) {
	o.LabelFormat = &v
}

// GetStandardIdForLabel returns the StandardIdForLabel field value if set, zero value otherwise.
func (o *Label) GetStandardIdForLabel() StandardIdForLabel {
	if o == nil || IsNil(o.StandardIdForLabel) {
		var ret StandardIdForLabel
		return ret
	}
	return *o.StandardIdForLabel
}

// GetStandardIdForLabelOk returns a tuple with the StandardIdForLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Label) GetStandardIdForLabelOk() (*StandardIdForLabel, bool) {
	if o == nil || IsNil(o.StandardIdForLabel) {
		return nil, false
	}
	return o.StandardIdForLabel, true
}

// HasStandardIdForLabel returns a boolean if a field has been set.
func (o *Label) HasStandardIdForLabel() bool {
	if o != nil && !IsNil(o.StandardIdForLabel) {
		return true
	}

	return false
}

// SetStandardIdForLabel gets a reference to the given StandardIdForLabel and assigns it to the StandardIdForLabel field.
func (o *Label) SetStandardIdForLabel(v StandardIdForLabel) {
	o.StandardIdForLabel = &v
}

func (o Label) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomTextForLabel) {
		toSerialize["CustomTextForLabel"] = o.CustomTextForLabel
	}
	toSerialize["Dimensions"] = o.Dimensions
	toSerialize["FileContents"] = o.FileContents
	if !IsNil(o.LabelFormat) {
		toSerialize["LabelFormat"] = o.LabelFormat
	}
	if !IsNil(o.StandardIdForLabel) {
		toSerialize["StandardIdForLabel"] = o.StandardIdForLabel
	}
	return toSerialize, nil
}

type NullableLabel struct {
	value *Label
	isSet bool
}

func (v NullableLabel) Get() *Label {
	return v.value
}

func (v *NullableLabel) Set(val *Label) {
	v.value = val
	v.isSet = true
}

func (v NullableLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabel(val *Label) *NullableLabel {
	return &NullableLabel{value: val, isSet: true}
}

func (v NullableLabel) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
