package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the ContentTargetingPredicate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentTargetingPredicate{}

// ContentTargetingPredicate A predicate to match against in the content targeting expression.
type ContentTargetingPredicate struct {
	Type *string `json:"type,omitempty"`
	// The value to be targeted.  The following table shows all possible values of the `contentCategorySameAs` predicate. | Category              | Subcategory                             | Value                    | |-----------------------|-----------------------------------------|--------------------------| | Movies and Television | All Movies and Television               | amzn1.iab-content.SPSHQ5 | | Movies and Television | Action or Adventure                     | amzn1.iab-content.325    | | Movies and Television | Animation or Anime                      | amzn1.iab-content.641    | | Movies and Television | Biographies                             | amzn1.iab-content.44     | | Movies and Television | Comedy                                  | amzn1.iab-content.646    | | Movies and Television | Documentary                             | amzn1.iab-content.332    | | Movies and Television | Drama                                   | amzn1.iab-content.647    | | Movies and Television | Factual                                 | amzn1.iab-content.648    | | Movies and Television | Family                                  | amzn1.iab-content.645    | | Movies and Television | Fantasy                                 | amzn1.iab-content.335    | | Movies and Television | History                                 | amzn1.iab-content.EZWB7V | | Movies and Television | Holiday                                 | amzn1.iab-content.649    | | Movies and Television | Horror                                  | amzn1.iab-content.336    | | Movies and Television | Lifestyle                               | amzn1.iab-content.TIFQA5 | | Movies and Television | Music Video                             | amzn1.iab-content.650    | | Movies and Television | Musicals                                | amzn1.iab-content.156    | | Movies and Television | Mystery                                 | amzn1.iab-content.331    | | Movies and Television | Reality TV                              | amzn1.iab-content.651    | | Movies and Television | Romance                                 | amzn1.iab-content.326    | | Movies and Television | Science Fiction                         | amzn1.iab-content.652    | | Movies and Television | Soap Opera                              | amzn1.iab-content.642    | | Movies and Television | Special Interest (Indie or Art House)   | amzn1.iab-content.643    | | Movies and Television | Sports Radio                            | amzn1.iab-content.370    | | Movies and Television | Talk Show                               | amzn1.iab-content.A0AH3G | | Movies and Television | True Crime                              | amzn1.iab-content.KHPC5A | | Movies and Television | Western                                 | amzn1.iab-content.KHPC6A | | Music and Radio       | All Music and Radio                     | amzn1.iab-content.338    | | Music and Radio       | Blues                                   | amzn1.iab-content.360    | | Music and Radio       | Classical Music                         | amzn1.iab-content.346    | | Music and Radio       | Comedy (Music and Audio)                | amzn1.iab-content.348    | | Music and Radio       | Pop, Contemporary Hits, or Top 40 Music | amzn1.iab-content.349    | | Music and Radio       | Country Music                           | amzn1.iab-content.350    | | Music and Radio       | Dance and Electronic Music              | amzn1.iab-content.351    | | Music and Radio       | Hip Hop Music                           | amzn1.iab-content.355    | | Music and Radio       | Inspirational or New Age Music          | amzn1.iab-content.356    | | Music and Radio       | Jazz                                    | amzn1.iab-content.357    | | Music and Radio       | Oldies or Adult Standards               | amzn1.iab-content.358    | | Music and Radio       | R&B, Soul or Funk Music                 | amzn1.iab-content.362    | | Music and Radio       | Reggae                                  | amzn1.iab-content.359    | | Music and Radio       | Rock Music                              | amzn1.iab-content.363    | | Music and Radio       | Songwriters or Folk                     | amzn1.iab-content.353    | | Music and Radio       | World or International Music            | amzn1.iab-content.352    | | Video Games           | All Video Games                         | amzn1.iab-content.680    | | Video Games           | Action-Adventure Games                  | amzn1.iab-content.691    | | Video Games           | Casual Games                            | amzn1.iab-content.693    | | Video Games           | Puzzle Video Games                      | amzn1.iab-content.698    | | Video Games           | Racing Video Games                      | amzn1.iab-content.VK7KD0 | | Video Games           | Role-Playing Video Games                | amzn1.iab-content.687    | | Video Games           | Simulation Video Games                  | amzn1.iab-content.688    | | Video Games           | Sports Video Games                      | amzn1.iab-content.689    | | Video Games           | Strategy Video Games                    | amzn1.iab-content.690    | | Video Games           | PC Games                                | amzn1.iab-content.684    | | Video Games           | Mobile Games                            | amzn1.iab-content.683    | | Video Games           | Console Games                           | amzn1.iab-content.681    | | Video Games           | eSports                                 | amzn1.iab-content.682    |
	Value *string `json:"value,omitempty"`
}

// NewContentTargetingPredicate instantiates a new ContentTargetingPredicate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentTargetingPredicate() *ContentTargetingPredicate {
	this := ContentTargetingPredicate{}
	return &this
}

// NewContentTargetingPredicateWithDefaults instantiates a new ContentTargetingPredicate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentTargetingPredicateWithDefaults() *ContentTargetingPredicate {
	this := ContentTargetingPredicate{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ContentTargetingPredicate) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentTargetingPredicate) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ContentTargetingPredicate) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ContentTargetingPredicate) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ContentTargetingPredicate) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentTargetingPredicate) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ContentTargetingPredicate) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ContentTargetingPredicate) SetValue(v string) {
	o.Value = &v
}

func (o ContentTargetingPredicate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableContentTargetingPredicate struct {
	value *ContentTargetingPredicate
	isSet bool
}

func (v NullableContentTargetingPredicate) Get() *ContentTargetingPredicate {
	return v.value
}

func (v *NullableContentTargetingPredicate) Set(val *ContentTargetingPredicate) {
	v.value = val
	v.isSet = true
}

func (v NullableContentTargetingPredicate) IsSet() bool {
	return v.isSet
}

func (v *NullableContentTargetingPredicate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentTargetingPredicate(val *ContentTargetingPredicate) *NullableContentTargetingPredicate {
	return &NullableContentTargetingPredicate{value: val, isSet: true}
}

func (v NullableContentTargetingPredicate) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableContentTargetingPredicate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
