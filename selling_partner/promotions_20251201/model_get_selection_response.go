package promotions_20251201

import (
	"github.com/bytedance/sonic"
)

// checks if the GetSelectionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSelectionResponse{}

// GetSelectionResponse The response schema for `getSelection`. The `selectionDetails` field is always present when `type` is `ITEMS`.
type GetSelectionResponse struct {
	Selection Selection `json:"selection"`
}

// NewGetSelectionResponse instantiates a new GetSelectionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSelectionResponse(selection Selection) *GetSelectionResponse {
	this := GetSelectionResponse{}
	this.Selection = selection
	return &this
}

// NewGetSelectionResponseWithDefaults instantiates a new GetSelectionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSelectionResponseWithDefaults() *GetSelectionResponse {
	this := GetSelectionResponse{}
	return &this
}

// GetSelection returns the Selection field value
func (o *GetSelectionResponse) GetSelection() Selection {
	if o == nil {
		var ret Selection
		return ret
	}

	return o.Selection
}

// GetSelectionOk returns a tuple with the Selection field value
// and a boolean to check if the value has been set.
func (o *GetSelectionResponse) GetSelectionOk() (*Selection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Selection, true
}

// SetSelection sets field value
func (o *GetSelectionResponse) SetSelection(v Selection) {
	o.Selection = v
}

func (o GetSelectionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["selection"] = o.Selection
	return toSerialize, nil
}

type NullableGetSelectionResponse struct {
	value *GetSelectionResponse
	isSet bool
}

func (v NullableGetSelectionResponse) Get() *GetSelectionResponse {
	return v.value
}

func (v *NullableGetSelectionResponse) Set(val *GetSelectionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSelectionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSelectionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSelectionResponse(val *GetSelectionResponse) *NullableGetSelectionResponse {
	return &NullableGetSelectionResponse{value: val, isSet: true}
}

func (v NullableGetSelectionResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetSelectionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
