package easy_ship_20220323

import (
	"github.com/bytedance/sonic"
)

// checks if the Dimensions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dimensions{}

// Dimensions The dimensions of the scheduled package.
type Dimensions struct {
	// The numerical value of the specified dimension.
	Length *float32 `json:"length,omitempty"`
	// The numerical value of the specified dimension.
	Width *float32 `json:"width,omitempty"`
	// The numerical value of the specified dimension.
	Height *float32      `json:"height,omitempty"`
	Unit   *UnitOfLength `json:"unit,omitempty"`
	// A string of up to 255 characters.
	Identifier *string `json:"identifier,omitempty"`
}

// NewDimensions instantiates a new Dimensions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDimensions() *Dimensions {
	this := Dimensions{}
	return &this
}

// NewDimensionsWithDefaults instantiates a new Dimensions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDimensionsWithDefaults() *Dimensions {
	this := Dimensions{}
	return &this
}

// GetLength returns the Length field value if set, zero value otherwise.
func (o *Dimensions) GetLength() float32 {
	if o == nil || IsNil(o.Length) {
		var ret float32
		return ret
	}
	return *o.Length
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dimensions) GetLengthOk() (*float32, bool) {
	if o == nil || IsNil(o.Length) {
		return nil, false
	}
	return o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *Dimensions) HasLength() bool {
	if o != nil && !IsNil(o.Length) {
		return true
	}

	return false
}

// SetLength gets a reference to the given float32 and assigns it to the Length field.
func (o *Dimensions) SetLength(v float32) {
	o.Length = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *Dimensions) GetWidth() float32 {
	if o == nil || IsNil(o.Width) {
		var ret float32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dimensions) GetWidthOk() (*float32, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *Dimensions) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given float32 and assigns it to the Width field.
func (o *Dimensions) SetWidth(v float32) {
	o.Width = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *Dimensions) GetHeight() float32 {
	if o == nil || IsNil(o.Height) {
		var ret float32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dimensions) GetHeightOk() (*float32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *Dimensions) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given float32 and assigns it to the Height field.
func (o *Dimensions) SetHeight(v float32) {
	o.Height = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *Dimensions) GetUnit() UnitOfLength {
	if o == nil || IsNil(o.Unit) {
		var ret UnitOfLength
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dimensions) GetUnitOk() (*UnitOfLength, bool) {
	if o == nil || IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *Dimensions) HasUnit() bool {
	if o != nil && !IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given UnitOfLength and assigns it to the Unit field.
func (o *Dimensions) SetUnit(v UnitOfLength) {
	o.Unit = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *Dimensions) GetIdentifier() string {
	if o == nil || IsNil(o.Identifier) {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dimensions) GetIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *Dimensions) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *Dimensions) SetIdentifier(v string) {
	o.Identifier = &v
}

func (o Dimensions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Length) {
		toSerialize["length"] = o.Length
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	if !IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}
	if !IsNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	return toSerialize, nil
}

type NullableDimensions struct {
	value *Dimensions
	isSet bool
}

func (v NullableDimensions) Get() *Dimensions {
	return v.value
}

func (v *NullableDimensions) Set(val *Dimensions) {
	v.value = val
	v.isSet = true
}

func (v NullableDimensions) IsSet() bool {
	return v.isSet
}

func (v *NullableDimensions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDimensions(val *Dimensions) *NullableDimensions {
	return &NullableDimensions{value: val, isSet: true}
}

func (v NullableDimensions) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDimensions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
