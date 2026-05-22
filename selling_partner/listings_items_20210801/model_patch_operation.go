package listings_items_20210801

import (
	"github.com/bytedance/sonic"
)

// checks if the PatchOperation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchOperation{}

// PatchOperation Individual JSON Patch operation for an HTTP PATCH request.
type PatchOperation struct {
	// Type of JSON Patch operation. Supported JSON Patch operations include add, replace, and delete. Refer to [JavaScript Object Notation (JSON) Patch](https://tools.ietf.org/html/rfc6902) for more information.
	Op string `json:"op"`
	// JSON Pointer path of the element to patch. Refer to [JavaScript Object Notation (JSON) Patch](https://tools.ietf.org/html/rfc6902) for more information.
	Path string `json:"path"`
	// JSON value to add, replace, or delete.
	Value []map[string]interface{} `json:"value,omitempty"`
}

type _PatchOperation PatchOperation

// NewPatchOperation instantiates a new PatchOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchOperation(op string, path string) *PatchOperation {
	this := PatchOperation{}
	this.Op = op
	this.Path = path
	return &this
}

// NewPatchOperationWithDefaults instantiates a new PatchOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchOperationWithDefaults() *PatchOperation {
	this := PatchOperation{}
	return &this
}

// GetOp returns the Op field value
func (o *PatchOperation) GetOp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Op
}

// GetOpOk returns a tuple with the Op field value
// and a boolean to check if the value has been set.
func (o *PatchOperation) GetOpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Op, true
}

// SetOp sets field value
func (o *PatchOperation) SetOp(v string) {
	o.Op = v
}

// GetPath returns the Path field value
func (o *PatchOperation) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *PatchOperation) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *PatchOperation) SetPath(v string) {
	o.Path = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PatchOperation) GetValue() []map[string]interface{} {
	if o == nil || IsNil(o.Value) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchOperation) GetValueOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PatchOperation) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given []map[string]interface{} and assigns it to the Value field.
func (o *PatchOperation) SetValue(v []map[string]interface{}) {
	o.Value = v
}

func (o PatchOperation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["op"] = o.Op
	toSerialize["path"] = o.Path
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullablePatchOperation struct {
	value *PatchOperation
	isSet bool
}

func (v NullablePatchOperation) Get() *PatchOperation {
	return v.value
}

func (v *NullablePatchOperation) Set(val *PatchOperation) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchOperation) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchOperation(val *PatchOperation) *NullablePatchOperation {
	return &NullablePatchOperation{value: val, isSet: true}
}

func (v NullablePatchOperation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePatchOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
