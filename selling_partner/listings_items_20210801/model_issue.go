package listings_items_20210801

import (
	"github.com/bytedance/sonic"
)

// checks if the Issue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Issue{}

// Issue An issue with a listings item.
type Issue struct {
	// An issue code that identifies the type of issue.
	Code string `json:"code"`
	// A message that describes the issue.
	Message string `json:"message"`
	// The severity of the issue.
	Severity string `json:"severity"`
	// The names of the attributes associated with the issue, if applicable.
	AttributeNames []string `json:"attributeNames,omitempty"`
	// List of issue categories.   Possible vales:   * `INVALID_ATTRIBUTE` - Indicating an invalid attribute in the listing.   * `MISSING_ATTRIBUTE` - Highlighting a missing attribute in the listing.   * `INVALID_IMAGE` - Signifying an invalid image in the listing.   * `MISSING_IMAGE` - Noting the absence of an image in the listing.   * `INVALID_PRICE` - Pertaining to issues with the listing's price-related attributes.   * `MISSING_PRICE` - Pointing out the absence of a price attribute in the listing.   * `DUPLICATE` - Identifying listings with potential duplicate problems, such as this ASIN potentially being a duplicate of another ASIN.   * `QUALIFICATION_REQUIRED` - Indicating that the listing requires qualification-related approval.
	Categories   []string           `json:"categories"`
	Enforcements *IssueEnforcements `json:"enforcements,omitempty"`
}

type _Issue Issue

// NewIssue instantiates a new Issue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssue(code string, message string, severity string, categories []string) *Issue {
	this := Issue{}
	this.Code = code
	this.Message = message
	this.Severity = severity
	this.Categories = categories
	return &this
}

// NewIssueWithDefaults instantiates a new Issue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssueWithDefaults() *Issue {
	this := Issue{}
	return &this
}

// GetCode returns the Code field value
func (o *Issue) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *Issue) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *Issue) SetCode(v string) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *Issue) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *Issue) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *Issue) SetMessage(v string) {
	o.Message = v
}

// GetSeverity returns the Severity field value
func (o *Issue) GetSeverity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *Issue) GetSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value
func (o *Issue) SetSeverity(v string) {
	o.Severity = v
}

// GetAttributeNames returns the AttributeNames field value if set, zero value otherwise.
func (o *Issue) GetAttributeNames() []string {
	if o == nil || IsNil(o.AttributeNames) {
		var ret []string
		return ret
	}
	return o.AttributeNames
}

// GetAttributeNamesOk returns a tuple with the AttributeNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Issue) GetAttributeNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.AttributeNames) {
		return nil, false
	}
	return o.AttributeNames, true
}

// HasAttributeNames returns a boolean if a field has been set.
func (o *Issue) HasAttributeNames() bool {
	if o != nil && !IsNil(o.AttributeNames) {
		return true
	}

	return false
}

// SetAttributeNames gets a reference to the given []string and assigns it to the AttributeNames field.
func (o *Issue) SetAttributeNames(v []string) {
	o.AttributeNames = v
}

// GetCategories returns the Categories field value
func (o *Issue) GetCategories() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value
// and a boolean to check if the value has been set.
func (o *Issue) GetCategoriesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Categories, true
}

// SetCategories sets field value
func (o *Issue) SetCategories(v []string) {
	o.Categories = v
}

// GetEnforcements returns the Enforcements field value if set, zero value otherwise.
func (o *Issue) GetEnforcements() IssueEnforcements {
	if o == nil || IsNil(o.Enforcements) {
		var ret IssueEnforcements
		return ret
	}
	return *o.Enforcements
}

// GetEnforcementsOk returns a tuple with the Enforcements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Issue) GetEnforcementsOk() (*IssueEnforcements, bool) {
	if o == nil || IsNil(o.Enforcements) {
		return nil, false
	}
	return o.Enforcements, true
}

// HasEnforcements returns a boolean if a field has been set.
func (o *Issue) HasEnforcements() bool {
	if o != nil && !IsNil(o.Enforcements) {
		return true
	}

	return false
}

// SetEnforcements gets a reference to the given IssueEnforcements and assigns it to the Enforcements field.
func (o *Issue) SetEnforcements(v IssueEnforcements) {
	o.Enforcements = &v
}

func (o Issue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	toSerialize["severity"] = o.Severity
	if !IsNil(o.AttributeNames) {
		toSerialize["attributeNames"] = o.AttributeNames
	}
	toSerialize["categories"] = o.Categories
	if !IsNil(o.Enforcements) {
		toSerialize["enforcements"] = o.Enforcements
	}
	return toSerialize, nil
}

type NullableIssue struct {
	value *Issue
	isSet bool
}

func (v NullableIssue) Get() *Issue {
	return v.value
}

func (v *NullableIssue) Set(val *Issue) {
	v.value = val
	v.isSet = true
}

func (v NullableIssue) IsSet() bool {
	return v.isSet
}

func (v *NullableIssue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssue(val *Issue) *NullableIssue {
	return &NullableIssue{value: val, isSet: true}
}

func (v NullableIssue) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableIssue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
