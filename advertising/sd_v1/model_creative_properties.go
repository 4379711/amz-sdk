package sd_v1

import (
	"fmt"
	"github.com/bytedance/sonic"
)

// CreativeProperties Select customizations on your creative from any combination of headline, logo, custom image and backgrounds.
type CreativeProperties struct {
	BackgroundCreativeProperties  *BackgroundCreativeProperties
	CustomImageCreativeProperties *CustomImageCreativeProperties
	HeadlineCreativeProperties    *HeadlineCreativeProperties
	LeadGenCreativeProperties     *LeadGenCreativeProperties
	LogoCreativeProperties        *LogoCreativeProperties
	VideoCreativeProperties       *VideoCreativeProperties
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CreativeProperties) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into BackgroundCreativeProperties
	err = sonic.Unmarshal(data, &dst.BackgroundCreativeProperties)
	if err == nil {
		jsonBackgroundCreativeProperties, _ := sonic.Marshal(dst.BackgroundCreativeProperties)
		if string(jsonBackgroundCreativeProperties) == "{}" { // empty struct
			dst.BackgroundCreativeProperties = nil
		} else {
			return nil // data stored in dst.BackgroundCreativeProperties, return on the first match
		}
	} else {
		dst.BackgroundCreativeProperties = nil
	}

	// try to unmarshal JSON data into CustomImageCreativeProperties
	err = sonic.Unmarshal(data, &dst.CustomImageCreativeProperties)
	if err == nil {
		jsonCustomImageCreativeProperties, _ := sonic.Marshal(dst.CustomImageCreativeProperties)
		if string(jsonCustomImageCreativeProperties) == "{}" { // empty struct
			dst.CustomImageCreativeProperties = nil
		} else {
			return nil // data stored in dst.CustomImageCreativeProperties, return on the first match
		}
	} else {
		dst.CustomImageCreativeProperties = nil
	}

	// try to unmarshal JSON data into HeadlineCreativeProperties
	err = sonic.Unmarshal(data, &dst.HeadlineCreativeProperties)
	if err == nil {
		jsonHeadlineCreativeProperties, _ := sonic.Marshal(dst.HeadlineCreativeProperties)
		if string(jsonHeadlineCreativeProperties) == "{}" { // empty struct
			dst.HeadlineCreativeProperties = nil
		} else {
			return nil // data stored in dst.HeadlineCreativeProperties, return on the first match
		}
	} else {
		dst.HeadlineCreativeProperties = nil
	}

	// try to unmarshal JSON data into LeadGenCreativeProperties
	err = sonic.Unmarshal(data, &dst.LeadGenCreativeProperties)
	if err == nil {
		jsonLeadGenCreativeProperties, _ := sonic.Marshal(dst.LeadGenCreativeProperties)
		if string(jsonLeadGenCreativeProperties) == "{}" { // empty struct
			dst.LeadGenCreativeProperties = nil
		} else {
			return nil // data stored in dst.LeadGenCreativeProperties, return on the first match
		}
	} else {
		dst.LeadGenCreativeProperties = nil
	}

	// try to unmarshal JSON data into LogoCreativeProperties
	err = sonic.Unmarshal(data, &dst.LogoCreativeProperties)
	if err == nil {
		jsonLogoCreativeProperties, _ := sonic.Marshal(dst.LogoCreativeProperties)
		if string(jsonLogoCreativeProperties) == "{}" { // empty struct
			dst.LogoCreativeProperties = nil
		} else {
			return nil // data stored in dst.LogoCreativeProperties, return on the first match
		}
	} else {
		dst.LogoCreativeProperties = nil
	}

	// try to unmarshal JSON data into VideoCreativeProperties
	err = sonic.Unmarshal(data, &dst.VideoCreativeProperties)
	if err == nil {
		jsonVideoCreativeProperties, _ := sonic.Marshal(dst.VideoCreativeProperties)
		if string(jsonVideoCreativeProperties) == "{}" { // empty struct
			dst.VideoCreativeProperties = nil
		} else {
			return nil // data stored in dst.VideoCreativeProperties, return on the first match
		}
	} else {
		dst.VideoCreativeProperties = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(CreativeProperties)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *CreativeProperties) MarshalJSON() ([]byte, error) {
	if src.BackgroundCreativeProperties != nil {
		return sonic.Marshal(&src.BackgroundCreativeProperties)
	}

	if src.CustomImageCreativeProperties != nil {
		return sonic.Marshal(&src.CustomImageCreativeProperties)
	}

	if src.HeadlineCreativeProperties != nil {
		return sonic.Marshal(&src.HeadlineCreativeProperties)
	}

	if src.LeadGenCreativeProperties != nil {
		return sonic.Marshal(&src.LeadGenCreativeProperties)
	}

	if src.LogoCreativeProperties != nil {
		return sonic.Marshal(&src.LogoCreativeProperties)
	}

	if src.VideoCreativeProperties != nil {
		return sonic.Marshal(&src.VideoCreativeProperties)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableCreativeProperties struct {
	value *CreativeProperties
	isSet bool
}

func (v NullableCreativeProperties) Get() *CreativeProperties {
	return v.value
}

func (v *NullableCreativeProperties) Set(val *CreativeProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableCreativeProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableCreativeProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreativeProperties(val *CreativeProperties) *NullableCreativeProperties {
	return &NullableCreativeProperties{value: val, isSet: true}
}

func (v NullableCreativeProperties) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreativeProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
