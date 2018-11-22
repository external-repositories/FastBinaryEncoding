// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// Version: 1.1.0.0

package test
import _ "../proto"

import "encoding/json"

type EnumTyped uint8

const (
    EnumTyped_ENUM_VALUE_0 = EnumTyped(0) + 0
    EnumTyped_ENUM_VALUE_1 = EnumTyped(int('1')) + 0
    EnumTyped_ENUM_VALUE_2 = EnumTyped(int('1')) + 1
    EnumTyped_ENUM_VALUE_3 = EnumTyped(int('3')) + 0
    EnumTyped_ENUM_VALUE_4 = EnumTyped(int('3')) + 1
    EnumTyped_ENUM_VALUE_5 = EnumTyped_ENUM_VALUE_3
)

func (e EnumTyped) String() string {
    switch e {
    case EnumTyped_ENUM_VALUE_0:
        return "ENUM_VALUE_0"
    case EnumTyped_ENUM_VALUE_1:
        return "ENUM_VALUE_1"
    case EnumTyped_ENUM_VALUE_2:
        return "ENUM_VALUE_2"
    case EnumTyped_ENUM_VALUE_3:
        return "ENUM_VALUE_3"
    case EnumTyped_ENUM_VALUE_4:
        return "ENUM_VALUE_4"
    case EnumTyped_ENUM_VALUE_5:
        return "ENUM_VALUE_5"
    }
    return "<unknown>"
}

func (e EnumTyped) MarshalJSON() ([]byte, error) {
    return json.Marshal(uint8(e))
}

func (e *EnumTyped) UnmarshalJSON(b []byte) error {
    var value uint8
    err := json.Unmarshal(b, &value)
    if err != nil {
        return err
    }
    *e = EnumTyped(value)
    return nil
}