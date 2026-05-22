package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the PatchDocument type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchDocument{}

// PatchDocument JSONPatch request document.
type PatchDocument struct {
	// The JSONPatch operation type.
	Op string `json:"op"`
	// A path constructed from the JSON object to be updated.
	Path string `json:"path"`
	// The value used by the operation specified in the `op` field.
	Value interface{} `json:"value,omitempty"`
}

type _PatchDocument PatchDocument

// NewPatchDocument instantiates a new PatchDocument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchDocument(op string, path string) *PatchDocument {
	this := PatchDocument{}
	this.Op = op
	this.Path = path
	return &this
}

// NewPatchDocumentWithDefaults instantiates a new PatchDocument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchDocumentWithDefaults() *PatchDocument {
	this := PatchDocument{}
	return &this
}

// GetOp returns the Op field value
func (o *PatchDocument) GetOp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Op
}

// GetOpOk returns a tuple with the Op field value
// and a boolean to check if the value has been set.
func (o *PatchDocument) GetOpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Op, true
}

// SetOp sets field value
func (o *PatchDocument) SetOp(v string) {
	o.Op = v
}

// GetPath returns the Path field value
func (o *PatchDocument) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *PatchDocument) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *PatchDocument) SetPath(v string) {
	o.Path = v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchDocument) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchDocument) GetValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PatchDocument) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *PatchDocument) SetValue(v interface{}) {
	o.Value = v
}

func (o PatchDocument) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["op"] = o.Op
	toSerialize["path"] = o.Path
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullablePatchDocument struct {
	value *PatchDocument
	isSet bool
}

func (v NullablePatchDocument) Get() *PatchDocument {
	return v.value
}

func (v *NullablePatchDocument) Set(val *PatchDocument) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchDocument) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchDocument(val *PatchDocument) *NullablePatchDocument {
	return &NullablePatchDocument{value: val, isSet: true}
}

func (v NullablePatchDocument) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePatchDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
