package promotions_20251201

import (
	"github.com/bytedance/sonic"
)

// checks if the Selection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Selection{}

// Selection Defines which products qualify for a benefit, allowing either specific items to be selected by ASIN or the entire catalog with optional exclusions.
type Selection struct {
	// The unique identifier for this selection configuration. **Note:** This field is absent when `type` is `CATALOG`.
	SelectionId *string `json:"selectionId,omitempty"`
	// The revision identifier for the selection. Pass this value to the `getSelection` operation's `revisionId` query parameter to retrieve the correct version of the selection. A promotion may have multiple selection revisions when an update is in progress. **Note:** This field is absent when `type` is `CATALOG`.
	RevisionId *int32 `json:"revisionId,omitempty"`
	// Selection type.
	Type             string            `json:"type"`
	SelectionDetails *SelectionDetails `json:"selectionDetails,omitempty"`
}

// NewSelection instantiates a new Selection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelection(type_ string) *Selection {
	this := Selection{}
	this.Type = type_
	return &this
}

// NewSelectionWithDefaults instantiates a new Selection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelectionWithDefaults() *Selection {
	this := Selection{}
	return &this
}

// GetSelectionId returns the SelectionId field value if set, zero value otherwise.
func (o *Selection) GetSelectionId() string {
	if o == nil || IsNil(o.SelectionId) {
		var ret string
		return ret
	}
	return *o.SelectionId
}

// GetSelectionIdOk returns a tuple with the SelectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Selection) GetSelectionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SelectionId) {
		return nil, false
	}
	return o.SelectionId, true
}

// HasSelectionId returns a boolean if a field has been set.
func (o *Selection) HasSelectionId() bool {
	if o != nil && !IsNil(o.SelectionId) {
		return true
	}

	return false
}

// SetSelectionId gets a reference to the given string and assigns it to the SelectionId field.
func (o *Selection) SetSelectionId(v string) {
	o.SelectionId = &v
}

// GetRevisionId returns the RevisionId field value if set, zero value otherwise.
func (o *Selection) GetRevisionId() int32 {
	if o == nil || IsNil(o.RevisionId) {
		var ret int32
		return ret
	}
	return *o.RevisionId
}

// GetRevisionIdOk returns a tuple with the RevisionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Selection) GetRevisionIdOk() (*int32, bool) {
	if o == nil || IsNil(o.RevisionId) {
		return nil, false
	}
	return o.RevisionId, true
}

// HasRevisionId returns a boolean if a field has been set.
func (o *Selection) HasRevisionId() bool {
	if o != nil && !IsNil(o.RevisionId) {
		return true
	}

	return false
}

// SetRevisionId gets a reference to the given int32 and assigns it to the RevisionId field.
func (o *Selection) SetRevisionId(v int32) {
	o.RevisionId = &v
}

// GetType returns the Type field value
func (o *Selection) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Selection) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Selection) SetType(v string) {
	o.Type = v
}

// GetSelectionDetails returns the SelectionDetails field value if set, zero value otherwise.
func (o *Selection) GetSelectionDetails() SelectionDetails {
	if o == nil || IsNil(o.SelectionDetails) {
		var ret SelectionDetails
		return ret
	}
	return *o.SelectionDetails
}

// GetSelectionDetailsOk returns a tuple with the SelectionDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Selection) GetSelectionDetailsOk() (*SelectionDetails, bool) {
	if o == nil || IsNil(o.SelectionDetails) {
		return nil, false
	}
	return o.SelectionDetails, true
}

// HasSelectionDetails returns a boolean if a field has been set.
func (o *Selection) HasSelectionDetails() bool {
	if o != nil && !IsNil(o.SelectionDetails) {
		return true
	}

	return false
}

// SetSelectionDetails gets a reference to the given SelectionDetails and assigns it to the SelectionDetails field.
func (o *Selection) SetSelectionDetails(v SelectionDetails) {
	o.SelectionDetails = &v
}

func (o Selection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SelectionId) {
		toSerialize["selectionId"] = o.SelectionId
	}
	if !IsNil(o.RevisionId) {
		toSerialize["revisionId"] = o.RevisionId
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.SelectionDetails) {
		toSerialize["selectionDetails"] = o.SelectionDetails
	}
	return toSerialize, nil
}

type NullableSelection struct {
	value *Selection
	isSet bool
}

func (v NullableSelection) Get() *Selection {
	return v.value
}

func (v *NullableSelection) Set(val *Selection) {
	v.value = val
	v.isSet = true
}

func (v NullableSelection) IsSet() bool {
	return v.isSet
}

func (v *NullableSelection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelection(val *Selection) *NullableSelection {
	return &NullableSelection{value: val, isSet: true}
}

func (v NullableSelection) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSelection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
