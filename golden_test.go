// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains simple golden tests for various examples.
// Besides validating the results when the implementation changes,
// it provides a way to look at the generated code without having
// to execute the print statements in one's head.

package main

import (
	"strings"
	"testing"
)

// Golden represents a test case.
type Golden struct {
	name    string
	input   string // input; the package clause is provided when running the test.
	output  string // exected output.
	flags   map[string]bool
	options map[string]string
}

var noFlags = map[string]bool{}
var noOptions = map[string]string{}

var golden = []Golden{
	{"day", dayIn, dayOut, noFlags, noOptions},
	{"offset", offsetIn, offsetOut, noFlags, noOptions},
	{"gap", gapIn, gapOut, noFlags, noOptions},
	{"num", numIn, numOut, noFlags, noOptions},
	{"unum", unumIn, unumOut, noFlags, noOptions},
	{"prime", primeIn, primeOut, noFlags, noOptions},
	{"prime", primeJsonIn, primeJsonOut, map[string]bool{IncludeJSON: true}, noOptions},
	{"prime", primeTextIn, primeTextOut, map[string]bool{IncludeText: true}, noOptions},
	{"prime", primeYamlIn, primeYamlOut, map[string]bool{IncludeYAML: true}, noOptions},
	{"prime", primeSqlIn, primeSqlOut, map[string]bool{IncludeSQL: true}, noOptions},
	{"prime", primeJsonAndSqlIn, primeJsonAndSqlOut, map[string]bool{IncludeJSON: true, IncludeSQL: true}, noOptions},
	{"prefix", prefixIn, dayOut, noFlags, map[string]string{TrimPrefix: "Day"}},
	{"camel", camelIn, camelOut, noFlags, noOptions},
	{"camel", camelIn, strings.Replace(camelOut, camelString, strings.ToUpper(camelString), 1), noFlags, map[string]string{TransformMethod: ToUpper}},
	{"camel", camelIn, strings.Replace(camelOut, camelString, strings.ToLower(camelString), 1), noFlags, map[string]string{TransformMethod: ToLower}},
	{"camel", camelIn, strings.Replace(camelOut, camelString, camelJSON, 1), noFlags, map[string]string{TransformMethod: ToJSON}},
	{"camel", camelIn, camelSnakeOut, noFlags, map[string]string{TransformMethod: ToSnake}},
	{"camel", camelIn, strings.Replace(camelSnakeOut, camelSnake, strings.ToUpper(camelSnake), 1), noFlags, map[string]string{TransformMethod: ToSnakeUpper}},
	{"camel", camelIn, strings.Replace(camelSnakeOut, camelSnake, camelKebab, 1), noFlags, map[string]string{TransformMethod: ToKebab}},
	{"camel", camelIn, strings.Replace(camelSnakeOut, camelSnake, strings.ToUpper(camelKebab), 1), noFlags, map[string]string{TransformMethod: ToKebabUpper}},
	{"camel", camelIn, camelIgnoreLowerOut, map[string]bool{IgnoreCase: true, IncludeJSON: true, AllowNumeric: true}, map[string]string{TransformMethod: ToLower}},
	{"camel", camelIn, camelIgnoreUpperOut, map[string]bool{IgnoreCase: true, IncludeJSON: true, AllowNumeric: true}, map[string]string{TransformMethod: ToUpper}},
	{"camel", camelIn, camelIgnoreJSONOut, map[string]bool{IgnoreCase: true, IncludeJSON: true, AllowNumeric: true}, map[string]string{TransformMethod: ToJSON}},
}

// Each example starts with "type XXX [u]int", with a single space separating them.

// Simple test: enumeration of type int starting at 0.
const dayIn = `type Day int
const (
	Monday Day = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)
`

const dayOut = `
const _DayName = "MondayTuesdayWednesdayThursdayFridaySaturdaySunday"

var _DayIndex = [...]uint8{0, 6, 13, 22, 30, 36, 44, 50}

func (i Day) String() string {
	if i < 0 || i >= Day(len(_DayIndex)-1) {
		return fmt.Sprintf("Day(%d)", i)
	}
	return _DayName[_DayIndex[i]:_DayIndex[i+1]]
}

var _DayValues = []Day{0, 1, 2, 3, 4, 5, 6}

var _DayNameToValueMap = map[string]Day{
	_DayName[0:6]:   0,
	_DayName[6:13]:  1,
	_DayName[13:22]: 2,
	_DayName[22:30]: 3,
	_DayName[30:36]: 4,
	_DayName[36:44]: 5,
	_DayName[44:50]: 6,
}

// DayString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func DayString(s string) (Day, error) {
	if val, ok := _DayNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Day values", s)
}

// DayValues returns all values of the enum
func DayValues() []Day {
	return _DayValues
}

// IsADay returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Day) IsADay() bool {
	for _, v := range _DayValues {
		if i == v {
			return true
		}
	}
	return false
}
`

// Camel case test: enumeration with camel case tags
const camelIn = `type Camel int
const (
	EnumFirst Camel = iota
	EnumSecond
	EnumThird
	EnumFourth
	EnumFifth
	EnumSixth
	EnumSeventh
)
`

const camelOut = `
const _CamelName = "EnumFirstEnumSecondEnumThirdEnumFourthEnumFifthEnumSixthEnumSeventh"

var _CamelIndex = [...]uint8{0, 9, 19, 28, 38, 47, 56, 67}

func (i Camel) String() string {
	if i < 0 || i >= Camel(len(_CamelIndex)-1) {
		return fmt.Sprintf("Camel(%d)", i)
	}
	return _CamelName[_CamelIndex[i]:_CamelIndex[i+1]]
}

var _CamelValues = []Camel{0, 1, 2, 3, 4, 5, 6}

var _CamelNameToValueMap = map[string]Camel{
	_CamelName[0:9]:   0,
	_CamelName[9:19]:  1,
	_CamelName[19:28]: 2,
	_CamelName[28:38]: 3,
	_CamelName[38:47]: 4,
	_CamelName[47:56]: 5,
	_CamelName[56:67]: 6,
}

// CamelString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func CamelString(s string) (Camel, error) {
	if val, ok := _CamelNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Camel values", s)
}

// CamelValues returns all values of the enum
func CamelValues() []Camel {
	return _CamelValues
}

// IsACamel returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Camel) IsACamel() bool {
	for _, v := range _CamelValues {
		if i == v {
			return true
		}
	}
	return false
}
`

const camelIgnoreLowerOut = `
const _CamelName = "enumfirstenumsecondenumthirdenumfourthenumfifthenumsixthenumseventh"

var _CamelIndex = [...]uint8{0, 9, 19, 28, 38, 47, 56, 67}

func (i Camel) String() string {
	if i < 0 || i >= Camel(len(_CamelIndex)-1) {
		return fmt.Sprintf("Camel(%d)", i)
	}
	return _CamelName[_CamelIndex[i]:_CamelIndex[i+1]]
}

var _CamelValues = []Camel{0, 1, 2, 3, 4, 5, 6}

var _CamelNameToValueMap = map[string]Camel{
	_CamelName[0:9]:   0,
	_CamelName[9:19]:  1,
	_CamelName[19:28]: 2,
	_CamelName[28:38]: 3,
	_CamelName[38:47]: 4,
	_CamelName[47:56]: 5,
	_CamelName[56:67]: 6,
}

// CamelString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func CamelString(s string) (Camel, error) {
	if val, ok := _CamelNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	i, err := strconv.Atoi(s)
	if err == nil {
		for _, v := range _CamelNameToValueMap {
			if int(v) == i {
				return v, nil
			}
		}
	}
	return 0, fmt.Errorf("%s does not belong to Camel values", s)
}

// CamelValues returns all values of the enum
func CamelValues() []Camel {
	return _CamelValues
}

// IsACamel returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Camel) IsACamel() bool {
	for _, v := range _CamelValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for Camel
func (i Camel) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Camel
func (i *Camel) UnmarshalJSON(data []byte) error {
	var s string
	var err error
	if err = json.Unmarshal(data, &s); err != nil {
		var val int
		if err = json.Unmarshal(data, &val); err != nil {
			return fmt.Errorf("Camel should be a string, got %s", data)
		}
		*i = Camel(val)
		if !i.IsACamel() {
			return fmt.Errorf("Invalid value for Camel (%d)", val)
		}
		return nil
	}

	*i, err = CamelString(s)
	return err
}
`

const camelIgnoreUpperOut = `
const _CamelName = "ENUMFIRSTENUMSECONDENUMTHIRDENUMFOURTHENUMFIFTHENUMSIXTHENUMSEVENTH"

var _CamelIndex = [...]uint8{0, 9, 19, 28, 38, 47, 56, 67}

func (i Camel) String() string {
	if i < 0 || i >= Camel(len(_CamelIndex)-1) {
		return fmt.Sprintf("Camel(%d)", i)
	}
	return _CamelName[_CamelIndex[i]:_CamelIndex[i+1]]
}

var _CamelValues = []Camel{0, 1, 2, 3, 4, 5, 6}

var _CamelNameToValueMap = map[string]Camel{
	_CamelName[0:9]:   0,
	_CamelName[9:19]:  1,
	_CamelName[19:28]: 2,
	_CamelName[28:38]: 3,
	_CamelName[38:47]: 4,
	_CamelName[47:56]: 5,
	_CamelName[56:67]: 6,
}

// CamelString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func CamelString(s string) (Camel, error) {
	if val, ok := _CamelNameToValueMap[strings.ToUpper(s)]; ok {
		return val, nil
	}
	i, err := strconv.Atoi(s)
	if err == nil {
		for _, v := range _CamelNameToValueMap {
			if int(v) == i {
				return v, nil
			}
		}
	}
	return 0, fmt.Errorf("%s does not belong to Camel values", s)
}

// CamelValues returns all values of the enum
func CamelValues() []Camel {
	return _CamelValues
}

// IsACamel returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Camel) IsACamel() bool {
	for _, v := range _CamelValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for Camel
func (i Camel) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Camel
func (i *Camel) UnmarshalJSON(data []byte) error {
	var s string
	var err error
	if err = json.Unmarshal(data, &s); err != nil {
		var val int
		if err = json.Unmarshal(data, &val); err != nil {
			return fmt.Errorf("Camel should be a string, got %s", data)
		}
		*i = Camel(val)
		if !i.IsACamel() {
			return fmt.Errorf("Invalid value for Camel (%d)", val)
		}
		return nil
	}

	*i, err = CamelString(s)
	return err
}
`

const camelIgnoreJSONOut = `
const _CamelName = "enumFirstenumSecondenumThirdenumFourthenumFifthenumSixthenumSeventh"

var _CamelIndex = [...]uint8{0, 9, 19, 28, 38, 47, 56, 67}

func (i Camel) String() string {
	if i < 0 || i >= Camel(len(_CamelIndex)-1) {
		return fmt.Sprintf("Camel(%d)", i)
	}
	return _CamelName[_CamelIndex[i]:_CamelIndex[i+1]]
}

var _CamelValues = []Camel{0, 1, 2, 3, 4, 5, 6}

var _CamelNameToValueMap = map[string]Camel{
	_CamelName[0:9]:   0,
	_CamelName[9:19]:  1,
	_CamelName[19:28]: 2,
	_CamelName[28:38]: 3,
	_CamelName[38:47]: 4,
	_CamelName[47:56]: 5,
	_CamelName[56:67]: 6,
}

// CamelString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func CamelString(s string) (Camel, error) {
	if val, ok := _CamelNameToValueMap[s]; ok {
		return val, nil
	}
	for k, v := range _CamelNameToValueMap {
		if strings.EqualFold(s, k) {
			return v, nil
		}
	}
	i, err := strconv.Atoi(s)
	if err == nil {
		for _, v := range _CamelNameToValueMap {
			if int(v) == i {
				return v, nil
			}
		}
	}
	return 0, fmt.Errorf("%s does not belong to Camel values", s)
}

// CamelValues returns all values of the enum
func CamelValues() []Camel {
	return _CamelValues
}

// IsACamel returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Camel) IsACamel() bool {
	for _, v := range _CamelValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for Camel
func (i Camel) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Camel
func (i *Camel) UnmarshalJSON(data []byte) error {
	var s string
	var err error
	if err = json.Unmarshal(data, &s); err != nil {
		var val int
		if err = json.Unmarshal(data, &val); err != nil {
			return fmt.Errorf("Camel should be a string, got %s", data)
		}
		*i = Camel(val)
		if !i.IsACamel() {
			return fmt.Errorf("Invalid value for Camel (%d)", val)
		}
		return nil
	}

	*i, err = CamelString(s)
	return err
}
`

const camelSnakeOut = `
const _CamelName = "enum_firstenum_secondenum_thirdenum_fourthenum_fifthenum_sixthenum_seventh"

var _CamelIndex = [...]uint8{0, 10, 21, 31, 42, 52, 62, 74}

func (i Camel) String() string {
	if i < 0 || i >= Camel(len(_CamelIndex)-1) {
		return fmt.Sprintf("Camel(%d)", i)
	}
	return _CamelName[_CamelIndex[i]:_CamelIndex[i+1]]
}

var _CamelValues = []Camel{0, 1, 2, 3, 4, 5, 6}

var _CamelNameToValueMap = map[string]Camel{
	_CamelName[0:10]:  0,
	_CamelName[10:21]: 1,
	_CamelName[21:31]: 2,
	_CamelName[31:42]: 3,
	_CamelName[42:52]: 4,
	_CamelName[52:62]: 5,
	_CamelName[62:74]: 6,
}

// CamelString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func CamelString(s string) (Camel, error) {
	if val, ok := _CamelNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Camel values", s)
}

// CamelValues returns all values of the enum
func CamelValues() []Camel {
	return _CamelValues
}

// IsACamel returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Camel) IsACamel() bool {
	for _, v := range _CamelValues {
		if i == v {
			return true
		}
	}
	return false
}
`

const camelString = "EnumFirstEnumSecondEnumThirdEnumFourthEnumFifthEnumSixthEnumSeventh"
const camelJSON = "enumFirstenumSecondenumThirdenumFourthenumFifthenumSixthenumSeventh"
const camelSnake = "enum_firstenum_secondenum_thirdenum_fourthenum_fifthenum_sixthenum_seventh"
const camelKebab = "enum-firstenum-secondenum-thirdenum-fourthenum-fifthenum-sixthenum-seventh"

// Enumeration with an offset.
// Also includes a duplicate.
const offsetIn = `type Number int
const (
	_ Number = iota
	One
	Two
	Three
	AnotherOne = One  // Duplicate; note that AnotherOne doesn't appear below.
)
`

const offsetOut = `
const _NumberName = "OneTwoThree"

var _NumberIndex = [...]uint8{0, 3, 6, 11}

func (i Number) String() string {
	i -= 1
	if i < 0 || i >= Number(len(_NumberIndex)-1) {
		return fmt.Sprintf("Number(%d)", i+1)
	}
	return _NumberName[_NumberIndex[i]:_NumberIndex[i+1]]
}

var _NumberValues = []Number{1, 2, 3}

var _NumberNameToValueMap = map[string]Number{
	_NumberName[0:3]:  1,
	_NumberName[3:6]:  2,
	_NumberName[6:11]: 3,
}

// NumberString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func NumberString(s string) (Number, error) {
	if val, ok := _NumberNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Number values", s)
}

// NumberValues returns all values of the enum
func NumberValues() []Number {
	return _NumberValues
}

// IsANumber returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Number) IsANumber() bool {
	for _, v := range _NumberValues {
		if i == v {
			return true
		}
	}
	return false
}
`

// Gaps and an offset.
const gapIn = `type Gap int
const (
	Two Gap = 2
	Three Gap = 3
	Five Gap = 5
	Six Gap = 6
	Seven Gap = 7
	Eight Gap = 8
	Nine Gap = 9
	Eleven Gap = 11
)
`

const gapOut = `
const (
	_GapName_0 = "TwoThree"
	_GapName_1 = "FiveSixSevenEightNine"
	_GapName_2 = "Eleven"
)

var (
	_GapIndex_0 = [...]uint8{0, 3, 8}
	_GapIndex_1 = [...]uint8{0, 4, 7, 12, 17, 21}
	_GapIndex_2 = [...]uint8{0, 6}
)

func (i Gap) String() string {
	switch {
	case 2 <= i && i <= 3:
		i -= 2
		return _GapName_0[_GapIndex_0[i]:_GapIndex_0[i+1]]
	case 5 <= i && i <= 9:
		i -= 5
		return _GapName_1[_GapIndex_1[i]:_GapIndex_1[i+1]]
	case i == 11:
		return _GapName_2
	default:
		return fmt.Sprintf("Gap(%d)", i)
	}
}

var _GapValues = []Gap{2, 3, 5, 6, 7, 8, 9, 11}

var _GapNameToValueMap = map[string]Gap{
	_GapName_0[0:3]:   2,
	_GapName_0[3:8]:   3,
	_GapName_1[0:4]:   5,
	_GapName_1[4:7]:   6,
	_GapName_1[7:12]:  7,
	_GapName_1[12:17]: 8,
	_GapName_1[17:21]: 9,
	_GapName_2[0:6]:   11,
}

// GapString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func GapString(s string) (Gap, error) {
	if val, ok := _GapNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Gap values", s)
}

// GapValues returns all values of the enum
func GapValues() []Gap {
	return _GapValues
}

// IsAGap returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Gap) IsAGap() bool {
	for _, v := range _GapValues {
		if i == v {
			return true
		}
	}
	return false
}
`

// Signed integers spanning zero.
const numIn = `type Num int
const (
	m_2 Num = -2 + iota
	m_1
	m0
	m1
	m2
)
`

const numOut = `
const _NumName = "m_2m_1m0m1m2"

var _NumIndex = [...]uint8{0, 3, 6, 8, 10, 12}

func (i Num) String() string {
	i -= -2
	if i < 0 || i >= Num(len(_NumIndex)-1) {
		return fmt.Sprintf("Num(%d)", i+-2)
	}
	return _NumName[_NumIndex[i]:_NumIndex[i+1]]
}

var _NumValues = []Num{-2, -1, 0, 1, 2}

var _NumNameToValueMap = map[string]Num{
	_NumName[0:3]:   -2,
	_NumName[3:6]:   -1,
	_NumName[6:8]:   0,
	_NumName[8:10]:  1,
	_NumName[10:12]: 2,
}

// NumString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func NumString(s string) (Num, error) {
	if val, ok := _NumNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Num values", s)
}

// NumValues returns all values of the enum
func NumValues() []Num {
	return _NumValues
}

// IsANum returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Num) IsANum() bool {
	for _, v := range _NumValues {
		if i == v {
			return true
		}
	}
	return false
}
`

// Unsigned integers spanning zero.
const unumIn = `type Unum uint
const (
	m_2 Unum = iota + 253
	m_1
)

const (
	m0 Unum = iota
	m1
	m2
)
`

const unumOut = `
const (
	_UnumName_0 = "m0m1m2"
	_UnumName_1 = "m_2m_1"
)

var (
	_UnumIndex_0 = [...]uint8{0, 2, 4, 6}
	_UnumIndex_1 = [...]uint8{0, 3, 6}
)

func (i Unum) String() string {
	switch {
	case 0 <= i && i <= 2:
		return _UnumName_0[_UnumIndex_0[i]:_UnumIndex_0[i+1]]
	case 253 <= i && i <= 254:
		i -= 253
		return _UnumName_1[_UnumIndex_1[i]:_UnumIndex_1[i+1]]
	default:
		return fmt.Sprintf("Unum(%d)", i)
	}
}

var _UnumValues = []Unum{0, 1, 2, 253, 254}

var _UnumNameToValueMap = map[string]Unum{
	_UnumName_0[0:2]: 0,
	_UnumName_0[2:4]: 1,
	_UnumName_0[4:6]: 2,
	_UnumName_1[0:3]: 253,
	_UnumName_1[3:6]: 254,
}

// UnumString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func UnumString(s string) (Unum, error) {
	if val, ok := _UnumNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Unum values", s)
}

// UnumValues returns all values of the enum
func UnumValues() []Unum {
	return _UnumValues
}

// IsAUnum returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Unum) IsAUnum() bool {
	for _, v := range _UnumValues {
		if i == v {
			return true
		}
	}
	return false
}
`

// Enough gaps to trigger a map implementation of the method.
// Also includes a duplicate to test that it doesn't cause problems
const primeIn = `type Prime int
const (
	p2 Prime = 2
	p3 Prime = 3
	p5 Prime = 5
	p7 Prime = 7
	p77 Prime = 7 // Duplicate; note that p77 doesn't appear below.
	p11 Prime = 11
	p13 Prime = 13
	p17 Prime = 17
	p19 Prime = 19
	p23 Prime = 23
	p29 Prime = 29
	p37 Prime = 31
	p41 Prime = 41
	p43 Prime = 43
)
`

const primeOut = `
const _PrimeName = "p2p3p5p7p11p13p17p19p23p29p37p41p43"

var _PrimeMap = map[Prime]string{
	2:  _PrimeName[0:2],
	3:  _PrimeName[2:4],
	5:  _PrimeName[4:6],
	7:  _PrimeName[6:8],
	11: _PrimeName[8:11],
	13: _PrimeName[11:14],
	17: _PrimeName[14:17],
	19: _PrimeName[17:20],
	23: _PrimeName[20:23],
	29: _PrimeName[23:26],
	31: _PrimeName[26:29],
	41: _PrimeName[29:32],
	43: _PrimeName[32:35],
}

func (i Prime) String() string {
	if str, ok := _PrimeMap[i]; ok {
		return str
	}
	return fmt.Sprintf("Prime(%d)", i)
}

var _PrimeValues = []Prime{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43}

var _PrimeNameToValueMap = map[string]Prime{
	_PrimeName[0:2]:   2,
	_PrimeName[2:4]:   3,
	_PrimeName[4:6]:   5,
	_PrimeName[6:8]:   7,
	_PrimeName[8:11]:  11,
	_PrimeName[11:14]: 13,
	_PrimeName[14:17]: 17,
	_PrimeName[17:20]: 19,
	_PrimeName[20:23]: 23,
	_PrimeName[23:26]: 29,
	_PrimeName[26:29]: 31,
	_PrimeName[29:32]: 41,
	_PrimeName[32:35]: 43,
}

// PrimeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func PrimeString(s string) (Prime, error) {
	if val, ok := _PrimeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Prime values", s)
}

// PrimeValues returns all values of the enum
func PrimeValues() []Prime {
	return _PrimeValues
}

// IsAPrime returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Prime) IsAPrime() bool {
	_, ok := _PrimeMap[i]
	return ok
}
`
const primeJsonIn = `type Prime int
const (
	p2 Prime = 2
	p3 Prime = 3
	p5 Prime = 5
	p7 Prime = 7
	p77 Prime = 7 // Duplicate; note that p77 doesn't appear below.
	p11 Prime = 11
	p13 Prime = 13
	p17 Prime = 17
	p19 Prime = 19
	p23 Prime = 23
	p29 Prime = 29
	p37 Prime = 31
	p41 Prime = 41
	p43 Prime = 43
)
`

const primeJsonOut = `
const _PrimeName = "p2p3p5p7p11p13p17p19p23p29p37p41p43"

var _PrimeMap = map[Prime]string{
	2:  _PrimeName[0:2],
	3:  _PrimeName[2:4],
	5:  _PrimeName[4:6],
	7:  _PrimeName[6:8],
	11: _PrimeName[8:11],
	13: _PrimeName[11:14],
	17: _PrimeName[14:17],
	19: _PrimeName[17:20],
	23: _PrimeName[20:23],
	29: _PrimeName[23:26],
	31: _PrimeName[26:29],
	41: _PrimeName[29:32],
	43: _PrimeName[32:35],
}

func (i Prime) String() string {
	if str, ok := _PrimeMap[i]; ok {
		return str
	}
	return fmt.Sprintf("Prime(%d)", i)
}

var _PrimeValues = []Prime{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43}

var _PrimeNameToValueMap = map[string]Prime{
	_PrimeName[0:2]:   2,
	_PrimeName[2:4]:   3,
	_PrimeName[4:6]:   5,
	_PrimeName[6:8]:   7,
	_PrimeName[8:11]:  11,
	_PrimeName[11:14]: 13,
	_PrimeName[14:17]: 17,
	_PrimeName[17:20]: 19,
	_PrimeName[20:23]: 23,
	_PrimeName[23:26]: 29,
	_PrimeName[26:29]: 31,
	_PrimeName[29:32]: 41,
	_PrimeName[32:35]: 43,
}

// PrimeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func PrimeString(s string) (Prime, error) {
	if val, ok := _PrimeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Prime values", s)
}

// PrimeValues returns all values of the enum
func PrimeValues() []Prime {
	return _PrimeValues
}

// IsAPrime returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Prime) IsAPrime() bool {
	_, ok := _PrimeMap[i]
	return ok
}

// MarshalJSON implements the json.Marshaler interface for Prime
func (i Prime) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Prime
func (i *Prime) UnmarshalJSON(data []byte) error {
	var s string
	var err error
	if err = json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Prime should be a string, got %s", data)
	}

	*i, err = PrimeString(s)
	return err
}
`

const primeTextIn = `type Prime int
const (
	p2 Prime = 2
	p3 Prime = 3
	p5 Prime = 5
	p7 Prime = 7
	p77 Prime = 7 // Duplicate; note that p77 doesn't appear below.
	p11 Prime = 11
	p13 Prime = 13
	p17 Prime = 17
	p19 Prime = 19
	p23 Prime = 23
	p29 Prime = 29
	p37 Prime = 31
	p41 Prime = 41
	p43 Prime = 43
)
`

const primeTextOut = `
const _PrimeName = "p2p3p5p7p11p13p17p19p23p29p37p41p43"

var _PrimeMap = map[Prime]string{
	2:  _PrimeName[0:2],
	3:  _PrimeName[2:4],
	5:  _PrimeName[4:6],
	7:  _PrimeName[6:8],
	11: _PrimeName[8:11],
	13: _PrimeName[11:14],
	17: _PrimeName[14:17],
	19: _PrimeName[17:20],
	23: _PrimeName[20:23],
	29: _PrimeName[23:26],
	31: _PrimeName[26:29],
	41: _PrimeName[29:32],
	43: _PrimeName[32:35],
}

func (i Prime) String() string {
	if str, ok := _PrimeMap[i]; ok {
		return str
	}
	return fmt.Sprintf("Prime(%d)", i)
}

var _PrimeValues = []Prime{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43}

var _PrimeNameToValueMap = map[string]Prime{
	_PrimeName[0:2]:   2,
	_PrimeName[2:4]:   3,
	_PrimeName[4:6]:   5,
	_PrimeName[6:8]:   7,
	_PrimeName[8:11]:  11,
	_PrimeName[11:14]: 13,
	_PrimeName[14:17]: 17,
	_PrimeName[17:20]: 19,
	_PrimeName[20:23]: 23,
	_PrimeName[23:26]: 29,
	_PrimeName[26:29]: 31,
	_PrimeName[29:32]: 41,
	_PrimeName[32:35]: 43,
}

// PrimeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func PrimeString(s string) (Prime, error) {
	if val, ok := _PrimeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Prime values", s)
}

// PrimeValues returns all values of the enum
func PrimeValues() []Prime {
	return _PrimeValues
}

// IsAPrime returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Prime) IsAPrime() bool {
	_, ok := _PrimeMap[i]
	return ok
}

// MarshalText implements the encoding.TextMarshaler interface for Prime
func (i Prime) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for Prime
func (i *Prime) UnmarshalText(text []byte) error {
	var err error
	*i, err = PrimeString(string(text))
	return err
}
`

const primeYamlIn = `type Prime int
const (
	p2 Prime = 2
	p3 Prime = 3
	p5 Prime = 5
	p7 Prime = 7
	p77 Prime = 7 // Duplicate; note that p77 doesn't appear below.
	p11 Prime = 11
	p13 Prime = 13
	p17 Prime = 17
	p19 Prime = 19
	p23 Prime = 23
	p29 Prime = 29
	p37 Prime = 31
	p41 Prime = 41
	p43 Prime = 43
)
`

const primeYamlOut = `
const _PrimeName = "p2p3p5p7p11p13p17p19p23p29p37p41p43"

var _PrimeMap = map[Prime]string{
	2:  _PrimeName[0:2],
	3:  _PrimeName[2:4],
	5:  _PrimeName[4:6],
	7:  _PrimeName[6:8],
	11: _PrimeName[8:11],
	13: _PrimeName[11:14],
	17: _PrimeName[14:17],
	19: _PrimeName[17:20],
	23: _PrimeName[20:23],
	29: _PrimeName[23:26],
	31: _PrimeName[26:29],
	41: _PrimeName[29:32],
	43: _PrimeName[32:35],
}

func (i Prime) String() string {
	if str, ok := _PrimeMap[i]; ok {
		return str
	}
	return fmt.Sprintf("Prime(%d)", i)
}

var _PrimeValues = []Prime{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43}

var _PrimeNameToValueMap = map[string]Prime{
	_PrimeName[0:2]:   2,
	_PrimeName[2:4]:   3,
	_PrimeName[4:6]:   5,
	_PrimeName[6:8]:   7,
	_PrimeName[8:11]:  11,
	_PrimeName[11:14]: 13,
	_PrimeName[14:17]: 17,
	_PrimeName[17:20]: 19,
	_PrimeName[20:23]: 23,
	_PrimeName[23:26]: 29,
	_PrimeName[26:29]: 31,
	_PrimeName[29:32]: 41,
	_PrimeName[32:35]: 43,
}

// PrimeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func PrimeString(s string) (Prime, error) {
	if val, ok := _PrimeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Prime values", s)
}

// PrimeValues returns all values of the enum
func PrimeValues() []Prime {
	return _PrimeValues
}

// IsAPrime returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Prime) IsAPrime() bool {
	_, ok := _PrimeMap[i]
	return ok
}

// MarshalYAML implements a YAML Marshaler for Prime
func (i Prime) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for Prime
func (i *Prime) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = PrimeString(s)
	return err
}
`

const primeSqlIn = `type Prime int
const (
	p2 Prime = 2
	p3 Prime = 3
	p5 Prime = 5
	p7 Prime = 7
	p77 Prime = 7 // Duplicate; note that p77 doesn't appear below.
	p11 Prime = 11
	p13 Prime = 13
	p17 Prime = 17
	p19 Prime = 19
	p23 Prime = 23
	p29 Prime = 29
	p37 Prime = 31
	p41 Prime = 41
	p43 Prime = 43
)
`

const primeSqlOut = `
const _PrimeName = "p2p3p5p7p11p13p17p19p23p29p37p41p43"

var _PrimeMap = map[Prime]string{
	2:  _PrimeName[0:2],
	3:  _PrimeName[2:4],
	5:  _PrimeName[4:6],
	7:  _PrimeName[6:8],
	11: _PrimeName[8:11],
	13: _PrimeName[11:14],
	17: _PrimeName[14:17],
	19: _PrimeName[17:20],
	23: _PrimeName[20:23],
	29: _PrimeName[23:26],
	31: _PrimeName[26:29],
	41: _PrimeName[29:32],
	43: _PrimeName[32:35],
}

func (i Prime) String() string {
	if str, ok := _PrimeMap[i]; ok {
		return str
	}
	return fmt.Sprintf("Prime(%d)", i)
}

var _PrimeValues = []Prime{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43}

var _PrimeNameToValueMap = map[string]Prime{
	_PrimeName[0:2]:   2,
	_PrimeName[2:4]:   3,
	_PrimeName[4:6]:   5,
	_PrimeName[6:8]:   7,
	_PrimeName[8:11]:  11,
	_PrimeName[11:14]: 13,
	_PrimeName[14:17]: 17,
	_PrimeName[17:20]: 19,
	_PrimeName[20:23]: 23,
	_PrimeName[23:26]: 29,
	_PrimeName[26:29]: 31,
	_PrimeName[29:32]: 41,
	_PrimeName[32:35]: 43,
}

// PrimeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func PrimeString(s string) (Prime, error) {
	if val, ok := _PrimeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Prime values", s)
}

// PrimeValues returns all values of the enum
func PrimeValues() []Prime {
	return _PrimeValues
}

// IsAPrime returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Prime) IsAPrime() bool {
	_, ok := _PrimeMap[i]
	return ok
}

func (i Prime) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *Prime) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	str, ok := value.(string)
	if !ok {
		bytes, ok := value.([]byte)
		if !ok {
			return fmt.Errorf("value is not a byte slice")
		}

		str = string(bytes[:])
	}

	val, err := PrimeString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}
`

const primeJsonAndSqlIn = `type Prime int
const (
	p2 Prime = 2
	p3 Prime = 3
	p5 Prime = 5
	p7 Prime = 7
	p77 Prime = 7 // Duplicate; note that p77 doesn't appear below.
	p11 Prime = 11
	p13 Prime = 13
	p17 Prime = 17
	p19 Prime = 19
	p23 Prime = 23
	p29 Prime = 29
	p37 Prime = 31
	p41 Prime = 41
	p43 Prime = 43
)
`

const primeJsonAndSqlOut = `
const _PrimeName = "p2p3p5p7p11p13p17p19p23p29p37p41p43"

var _PrimeMap = map[Prime]string{
	2:  _PrimeName[0:2],
	3:  _PrimeName[2:4],
	5:  _PrimeName[4:6],
	7:  _PrimeName[6:8],
	11: _PrimeName[8:11],
	13: _PrimeName[11:14],
	17: _PrimeName[14:17],
	19: _PrimeName[17:20],
	23: _PrimeName[20:23],
	29: _PrimeName[23:26],
	31: _PrimeName[26:29],
	41: _PrimeName[29:32],
	43: _PrimeName[32:35],
}

func (i Prime) String() string {
	if str, ok := _PrimeMap[i]; ok {
		return str
	}
	return fmt.Sprintf("Prime(%d)", i)
}

var _PrimeValues = []Prime{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43}

var _PrimeNameToValueMap = map[string]Prime{
	_PrimeName[0:2]:   2,
	_PrimeName[2:4]:   3,
	_PrimeName[4:6]:   5,
	_PrimeName[6:8]:   7,
	_PrimeName[8:11]:  11,
	_PrimeName[11:14]: 13,
	_PrimeName[14:17]: 17,
	_PrimeName[17:20]: 19,
	_PrimeName[20:23]: 23,
	_PrimeName[23:26]: 29,
	_PrimeName[26:29]: 31,
	_PrimeName[29:32]: 41,
	_PrimeName[32:35]: 43,
}

// PrimeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func PrimeString(s string) (Prime, error) {
	if val, ok := _PrimeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Prime values", s)
}

// PrimeValues returns all values of the enum
func PrimeValues() []Prime {
	return _PrimeValues
}

// IsAPrime returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Prime) IsAPrime() bool {
	_, ok := _PrimeMap[i]
	return ok
}

// MarshalJSON implements the json.Marshaler interface for Prime
func (i Prime) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Prime
func (i *Prime) UnmarshalJSON(data []byte) error {
	var s string
	var err error
	if err = json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Prime should be a string, got %s", data)
	}

	*i, err = PrimeString(s)
	return err
}

func (i Prime) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *Prime) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	str, ok := value.(string)
	if !ok {
		bytes, ok := value.([]byte)
		if !ok {
			return fmt.Errorf("value is not a byte slice")
		}

		str = string(bytes[:])
	}

	val, err := PrimeString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}
`

const prefixIn = `type Day int
const (
	DayMonday Day = iota
	DayTuesday
	DayWednesday
	DayThursday
	DayFriday
	DaySaturday
	DaySunday
)
`

func TestGolden(t *testing.T) {
	for _, test := range golden {
		runGoldenTest(t, test)
	}
}

func runGoldenTest(t *testing.T, test Golden) {
	var g Generator
	input := "package test\n" + test.input
	file := test.name + ".go"
	g.parsePackage(".", []string{file}, input)
	// Extract the name and type of the constant from the first line.
	tokens := strings.SplitN(test.input, " ", 3)
	if len(tokens) != 3 {
		t.Fatalf("%s: need type declaration on first line", test.name)
	}
	g.generate(tokens[1], test.flags, test.options)
	got := string(g.format())
	if got != test.output {
		t.Errorf("%s: got\n====\n%s====\nexpected\n====%s", test.name, got, test.output)
	}
}
