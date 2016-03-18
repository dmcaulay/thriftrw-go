// Code generated by thriftrw

package testdata

import "github.com/thriftrw/thriftrw-go/wire"

type ContactInfo struct{ EmailAddress string }

func (v *ContactInfo) ToWire() wire.Value {
	var fs [1]wire.Field
	i := 0
	fs[i] = wire.Field{ID: 1, Value: wire.NewValueString(v.EmailAddress)}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fs[:i]})
}
func (v *ContactInfo) FromWire(w wire.Value) error {
	var err error
	for _, f := range w.GetStruct().Fields {
		switch f.ID {
		case 1:
			if f.Value.Type() == wire.TBinary {
				v.EmailAddress, err = f.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
func _ContactInfo_Read(w wire.Value) (*ContactInfo, error) {
	var v ContactInfo
	err := v.FromWire(w)
	return &v, err
}

type Edge struct {
	End   *Point
	Start *Point
}

func (v *Edge) ToWire() wire.Value {
	var fs [2]wire.Field
	i := 0
	fs[i] = wire.Field{ID: 2, Value: v.End.ToWire()}
	i++
	fs[i] = wire.Field{ID: 1, Value: v.Start.ToWire()}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fs[:i]})
}
func (v *Edge) FromWire(w wire.Value) error {
	var err error
	for _, f := range w.GetStruct().Fields {
		switch f.ID {
		case 2:
			if f.Value.Type() == wire.TStruct {
				v.End, err = _Point_Read(f.Value)
				if err != nil {
					return err
				}
			}
		case 1:
			if f.Value.Type() == wire.TStruct {
				v.Start, err = _Point_Read(f.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
func _Edge_Read(w wire.Value) (*Edge, error) {
	var v Edge
	err := v.FromWire(w)
	return &v, err
}

type _List_EnumDefault_ValueList []EnumDefault

func (v _List_EnumDefault_ValueList) ForEach(f func(wire.Value) error) error {
	for _, x := range v {
		err := f(x.ToWire())
		if err != nil {
			return err
		}
	}
	return nil
}
func (v _List_EnumDefault_ValueList) Close() {
}

type _Map_EnumWithDuplicateValues_I32_MapItemList map[EnumWithDuplicateValues]int32

func (m _Map_EnumWithDuplicateValues_I32_MapItemList) ForEach(f func(wire.MapItem) error) error {
	for k, v := range m {
		err := f(wire.MapItem{Key: k.ToWire(), Value: wire.NewValueI32(v)})
		if err != nil {
			return err
		}
	}
	return nil
}
func (m _Map_EnumWithDuplicateValues_I32_MapItemList) Close() {
}

type _Set_EnumWithValues_ValueList map[EnumWithValues]struct{}

func (v _Set_EnumWithValues_ValueList) ForEach(f func(wire.Value) error) error {
	for x := range v {
		err := f(x.ToWire())
		if err != nil {
			return err
		}
	}
	return nil
}
func (v _Set_EnumWithValues_ValueList) Close() {
}
func _List_EnumDefault_Read(l wire.List) ([]EnumDefault, error) {
	if l.ValueType != wire.TI32 {
		return nil, nil
	}
	o := make([]EnumDefault, 0, l.Size)
	err := l.Items.ForEach(func(x wire.Value) error {
		i, err := _EnumDefault_Read(x)
		if err != nil {
			return err
		}
		o = append(o, i)
		return nil
	})
	l.Items.Close()
	return o, err
}
func _Map_EnumWithDuplicateValues_I32_Read(m wire.Map) (map[EnumWithDuplicateValues]int32, error) {
	if m.KeyType != wire.TI32 {
		return nil, nil
	}
	if m.ValueType != wire.TI32 {
		return nil, nil
	}
	o := make(map[EnumWithDuplicateValues]int32, m.Size)
	err := m.Items.ForEach(func(x wire.MapItem) error {
		k, err := _EnumWithDuplicateValues_Read(x.Key)
		if err != nil {
			return err
		}
		v, err := x.Value.GetI32(), error(nil)
		if err != nil {
			return err
		}
		o[k] = v
		return nil
	})
	m.Items.Close()
	return o, err
}
func _Set_EnumWithValues_Read(s wire.Set) (map[EnumWithValues]struct{}, error) {
	if s.ValueType != wire.TI32 {
		return nil, nil
	}
	o := make(map[EnumWithValues]struct{}, s.Size)
	err := s.Items.ForEach(func(x wire.Value) error {
		i, err := _EnumWithValues_Read(x)
		if err != nil {
			return err
		}
		o[i] = struct{}{}
		return nil
	})
	s.Items.Close()
	return o, err
}

type EnumContainers struct {
	ListOfEnums []EnumDefault
	MapOfEnums  map[EnumWithDuplicateValues]int32
	SetOfEnums  map[EnumWithValues]struct{}
}

func (v *EnumContainers) ToWire() wire.Value {
	var fs [3]wire.Field
	i := 0
	if v.ListOfEnums != nil {
		fs[i] = wire.Field{ID: 1, Value: wire.NewValueList(wire.List{ValueType: wire.TI32, Size: len(v.ListOfEnums), Items: _List_EnumDefault_ValueList(v.ListOfEnums)})}
		i++
	}
	if v.MapOfEnums != nil {
		fs[i] = wire.Field{ID: 3, Value: wire.NewValueMap(wire.Map{KeyType: wire.TI32, ValueType: wire.TI32, Size: len(v.MapOfEnums), Items: _Map_EnumWithDuplicateValues_I32_MapItemList(v.MapOfEnums)})}
		i++
	}
	if v.SetOfEnums != nil {
		fs[i] = wire.Field{ID: 2, Value: wire.NewValueSet(wire.Set{ValueType: wire.TI32, Size: len(v.SetOfEnums), Items: _Set_EnumWithValues_ValueList(v.SetOfEnums)})}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fs[:i]})
}
func (v *EnumContainers) FromWire(w wire.Value) error {
	var err error
	for _, f := range w.GetStruct().Fields {
		switch f.ID {
		case 1:
			if f.Value.Type() == wire.TList {
				v.ListOfEnums, err = _List_EnumDefault_Read(f.Value.GetList())
				if err != nil {
					return err
				}
			}
		case 3:
			if f.Value.Type() == wire.TMap {
				v.MapOfEnums, err = _Map_EnumWithDuplicateValues_I32_Read(f.Value.GetMap())
				if err != nil {
					return err
				}
			}
		case 2:
			if f.Value.Type() == wire.TSet {
				v.SetOfEnums, err = _Set_EnumWithValues_Read(f.Value.GetSet())
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
func _EnumContainers_Read(w wire.Value) (*EnumContainers, error) {
	var v EnumContainers
	err := v.FromWire(w)
	return &v, err
}

type EnumDefault int32

const (
	EnumDefaultFoo EnumDefault = 0
	EnumDefaultBar EnumDefault = 1
	EnumDefaultBaz EnumDefault = 2
)

func (v EnumDefault) ToWire() wire.Value {
	return wire.NewValueI32(int32(v))
}
func (v *EnumDefault) FromWire(w wire.Value) error {
	*v = (EnumDefault)(w.GetI32())
	return nil
}
func _EnumDefault_Read(w wire.Value) (EnumDefault, error) {
	var v EnumDefault
	err := v.FromWire(w)
	return v, err
}

type EnumWithDuplicateName int32

const (
	EnumWithDuplicateNameA EnumWithDuplicateName = 0
	EnumWithDuplicateNameB EnumWithDuplicateName = 1
	EnumWithDuplicateNameC EnumWithDuplicateName = 2
	EnumWithDuplicateNameP EnumWithDuplicateName = 3
	EnumWithDuplicateNameQ EnumWithDuplicateName = 4
	EnumWithDuplicateNameR EnumWithDuplicateName = 5
	EnumWithDuplicateNameX EnumWithDuplicateName = 6
	EnumWithDuplicateNameY EnumWithDuplicateName = 7
	EnumWithDuplicateNameZ EnumWithDuplicateName = 8
)

func (v EnumWithDuplicateName) ToWire() wire.Value {
	return wire.NewValueI32(int32(v))
}
func (v *EnumWithDuplicateName) FromWire(w wire.Value) error {
	*v = (EnumWithDuplicateName)(w.GetI32())
	return nil
}
func _EnumWithDuplicateName_Read(w wire.Value) (EnumWithDuplicateName, error) {
	var v EnumWithDuplicateName
	err := v.FromWire(w)
	return v, err
}

type EnumWithDuplicateValues int32

const (
	EnumWithDuplicateValuesP EnumWithDuplicateValues = 0
	EnumWithDuplicateValuesQ EnumWithDuplicateValues = -1
	EnumWithDuplicateValuesR EnumWithDuplicateValues = 0
)

func (v EnumWithDuplicateValues) ToWire() wire.Value {
	return wire.NewValueI32(int32(v))
}
func (v *EnumWithDuplicateValues) FromWire(w wire.Value) error {
	*v = (EnumWithDuplicateValues)(w.GetI32())
	return nil
}
func _EnumWithDuplicateValues_Read(w wire.Value) (EnumWithDuplicateValues, error) {
	var v EnumWithDuplicateValues
	err := v.FromWire(w)
	return v, err
}

type EnumWithValues int32

const (
	EnumWithValuesX EnumWithValues = 123
	EnumWithValuesY EnumWithValues = 456
	EnumWithValuesZ EnumWithValues = 789
)

func (v EnumWithValues) ToWire() wire.Value {
	return wire.NewValueI32(int32(v))
}
func (v *EnumWithValues) FromWire(w wire.Value) error {
	*v = (EnumWithValues)(w.GetI32())
	return nil
}
func _EnumWithValues_Read(w wire.Value) (EnumWithValues, error) {
	var v EnumWithValues
	err := v.FromWire(w)
	return v, err
}

type Frame struct {
	Size    *Size
	TopLeft *Point
}

func (v *Frame) ToWire() wire.Value {
	var fs [2]wire.Field
	i := 0
	fs[i] = wire.Field{ID: 2, Value: v.Size.ToWire()}
	i++
	fs[i] = wire.Field{ID: 1, Value: v.TopLeft.ToWire()}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fs[:i]})
}
func (v *Frame) FromWire(w wire.Value) error {
	var err error
	for _, f := range w.GetStruct().Fields {
		switch f.ID {
		case 2:
			if f.Value.Type() == wire.TStruct {
				v.Size, err = _Size_Read(f.Value)
				if err != nil {
					return err
				}
			}
		case 1:
			if f.Value.Type() == wire.TStruct {
				v.TopLeft, err = _Point_Read(f.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
func _Frame_Read(w wire.Value) (*Frame, error) {
	var v Frame
	err := v.FromWire(w)
	return &v, err
}

type _List_Edge_ValueList []*Edge

func (v _List_Edge_ValueList) ForEach(f func(wire.Value) error) error {
	for _, x := range v {
		err := f(x.ToWire())
		if err != nil {
			return err
		}
	}
	return nil
}
func (v _List_Edge_ValueList) Close() {
}
func _List_Edge_Read(l wire.List) ([]*Edge, error) {
	if l.ValueType != wire.TStruct {
		return nil, nil
	}
	o := make([]*Edge, 0, l.Size)
	err := l.Items.ForEach(func(x wire.Value) error {
		i, err := _Edge_Read(x)
		if err != nil {
			return err
		}
		o = append(o, i)
		return nil
	})
	l.Items.Close()
	return o, err
}

type Graph struct{ Edges []*Edge }

func (v *Graph) ToWire() wire.Value {
	var fs [1]wire.Field
	i := 0
	fs[i] = wire.Field{ID: 1, Value: wire.NewValueList(wire.List{ValueType: wire.TStruct, Size: len(v.Edges), Items: _List_Edge_ValueList(v.Edges)})}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fs[:i]})
}
func (v *Graph) FromWire(w wire.Value) error {
	var err error
	for _, f := range w.GetStruct().Fields {
		switch f.ID {
		case 1:
			if f.Value.Type() == wire.TList {
				v.Edges, err = _List_Edge_Read(f.Value.GetList())
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
func _Graph_Read(w wire.Value) (*Graph, error) {
	var v Graph
	err := v.FromWire(w)
	return &v, err
}

type Point struct {
	X float64
	Y float64
}

func (v *Point) ToWire() wire.Value {
	var fs [2]wire.Field
	i := 0
	fs[i] = wire.Field{ID: 1, Value: wire.NewValueDouble(v.X)}
	i++
	fs[i] = wire.Field{ID: 2, Value: wire.NewValueDouble(v.Y)}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fs[:i]})
}
func (v *Point) FromWire(w wire.Value) error {
	var err error
	for _, f := range w.GetStruct().Fields {
		switch f.ID {
		case 1:
			if f.Value.Type() == wire.TDouble {
				v.X, err = f.Value.GetDouble(), error(nil)
				if err != nil {
					return err
				}
			}
		case 2:
			if f.Value.Type() == wire.TDouble {
				v.Y, err = f.Value.GetDouble(), error(nil)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
func _Point_Read(w wire.Value) (*Point, error) {
	var v Point
	err := v.FromWire(w)
	return &v, err
}

type _List_Binary_ValueList [][]byte

func (v _List_Binary_ValueList) ForEach(f func(wire.Value) error) error {
	for _, x := range v {
		err := f(wire.NewValueBinary(x))
		if err != nil {
			return err
		}
	}
	return nil
}
func (v _List_Binary_ValueList) Close() {
}

type _List_I64_ValueList []int64

func (v _List_I64_ValueList) ForEach(f func(wire.Value) error) error {
	for _, x := range v {
		err := f(wire.NewValueI64(x))
		if err != nil {
			return err
		}
	}
	return nil
}
func (v _List_I64_ValueList) Close() {
}

type _Map_I32_String_MapItemList map[int32]string

func (m _Map_I32_String_MapItemList) ForEach(f func(wire.MapItem) error) error {
	for k, v := range m {
		err := f(wire.MapItem{Key: wire.NewValueI32(k), Value: wire.NewValueString(v)})
		if err != nil {
			return err
		}
	}
	return nil
}
func (m _Map_I32_String_MapItemList) Close() {
}

type _Map_String_Bool_MapItemList map[string]bool

func (m _Map_String_Bool_MapItemList) ForEach(f func(wire.MapItem) error) error {
	for k, v := range m {
		err := f(wire.MapItem{Key: wire.NewValueString(k), Value: wire.NewValueBool(v)})
		if err != nil {
			return err
		}
	}
	return nil
}
func (m _Map_String_Bool_MapItemList) Close() {
}

type _Set_Byte_ValueList map[int8]struct{}

func (v _Set_Byte_ValueList) ForEach(f func(wire.Value) error) error {
	for x := range v {
		err := f(wire.NewValueI8(x))
		if err != nil {
			return err
		}
	}
	return nil
}
func (v _Set_Byte_ValueList) Close() {
}

type _Set_String_ValueList map[string]struct{}

func (v _Set_String_ValueList) ForEach(f func(wire.Value) error) error {
	for x := range v {
		err := f(wire.NewValueString(x))
		if err != nil {
			return err
		}
	}
	return nil
}
func (v _Set_String_ValueList) Close() {
}
func _List_Binary_Read(l wire.List) ([][]byte, error) {
	if l.ValueType != wire.TBinary {
		return nil, nil
	}
	o := make([][]byte, 0, l.Size)
	err := l.Items.ForEach(func(x wire.Value) error {
		i, err := x.GetBinary(), error(nil)
		if err != nil {
			return err
		}
		o = append(o, i)
		return nil
	})
	l.Items.Close()
	return o, err
}
func _List_I64_Read(l wire.List) ([]int64, error) {
	if l.ValueType != wire.TI64 {
		return nil, nil
	}
	o := make([]int64, 0, l.Size)
	err := l.Items.ForEach(func(x wire.Value) error {
		i, err := x.GetI64(), error(nil)
		if err != nil {
			return err
		}
		o = append(o, i)
		return nil
	})
	l.Items.Close()
	return o, err
}
func _Map_I32_String_Read(m wire.Map) (map[int32]string, error) {
	if m.KeyType != wire.TI32 {
		return nil, nil
	}
	if m.ValueType != wire.TBinary {
		return nil, nil
	}
	o := make(map[int32]string, m.Size)
	err := m.Items.ForEach(func(x wire.MapItem) error {
		k, err := x.Key.GetI32(), error(nil)
		if err != nil {
			return err
		}
		v, err := x.Value.GetString(), error(nil)
		if err != nil {
			return err
		}
		o[k] = v
		return nil
	})
	m.Items.Close()
	return o, err
}
func _Map_String_Bool_Read(m wire.Map) (map[string]bool, error) {
	if m.KeyType != wire.TBinary {
		return nil, nil
	}
	if m.ValueType != wire.TBool {
		return nil, nil
	}
	o := make(map[string]bool, m.Size)
	err := m.Items.ForEach(func(x wire.MapItem) error {
		k, err := x.Key.GetString(), error(nil)
		if err != nil {
			return err
		}
		v, err := x.Value.GetBool(), error(nil)
		if err != nil {
			return err
		}
		o[k] = v
		return nil
	})
	m.Items.Close()
	return o, err
}
func _Set_Byte_Read(s wire.Set) (map[int8]struct{}, error) {
	if s.ValueType != wire.TI8 {
		return nil, nil
	}
	o := make(map[int8]struct{}, s.Size)
	err := s.Items.ForEach(func(x wire.Value) error {
		i, err := x.GetI8(), error(nil)
		if err != nil {
			return err
		}
		o[i] = struct{}{}
		return nil
	})
	s.Items.Close()
	return o, err
}
func _Set_String_Read(s wire.Set) (map[string]struct{}, error) {
	if s.ValueType != wire.TBinary {
		return nil, nil
	}
	o := make(map[string]struct{}, s.Size)
	err := s.Items.ForEach(func(x wire.Value) error {
		i, err := x.GetString(), error(nil)
		if err != nil {
			return err
		}
		o[i] = struct{}{}
		return nil
	})
	s.Items.Close()
	return o, err
}

type PrimitiveContainers struct {
	ListOfBinary      [][]byte
	ListOfInts        []int64
	MapOfIntToString  map[int32]string
	MapOfStringToBool map[string]bool
	SetOfBytes        map[int8]struct{}
	SetOfStrings      map[string]struct{}
}

func (v *PrimitiveContainers) ToWire() wire.Value {
	var fs [6]wire.Field
	i := 0
	if v.ListOfBinary != nil {
		fs[i] = wire.Field{ID: 1, Value: wire.NewValueList(wire.List{ValueType: wire.TBinary, Size: len(v.ListOfBinary), Items: _List_Binary_ValueList(v.ListOfBinary)})}
		i++
	}
	if v.ListOfInts != nil {
		fs[i] = wire.Field{ID: 2, Value: wire.NewValueList(wire.List{ValueType: wire.TI64, Size: len(v.ListOfInts), Items: _List_I64_ValueList(v.ListOfInts)})}
		i++
	}
	if v.MapOfIntToString != nil {
		fs[i] = wire.Field{ID: 5, Value: wire.NewValueMap(wire.Map{KeyType: wire.TI32, ValueType: wire.TBinary, Size: len(v.MapOfIntToString), Items: _Map_I32_String_MapItemList(v.MapOfIntToString)})}
		i++
	}
	if v.MapOfStringToBool != nil {
		fs[i] = wire.Field{ID: 6, Value: wire.NewValueMap(wire.Map{KeyType: wire.TBinary, ValueType: wire.TBool, Size: len(v.MapOfStringToBool), Items: _Map_String_Bool_MapItemList(v.MapOfStringToBool)})}
		i++
	}
	if v.SetOfBytes != nil {
		fs[i] = wire.Field{ID: 4, Value: wire.NewValueSet(wire.Set{ValueType: wire.TI8, Size: len(v.SetOfBytes), Items: _Set_Byte_ValueList(v.SetOfBytes)})}
		i++
	}
	if v.SetOfStrings != nil {
		fs[i] = wire.Field{ID: 3, Value: wire.NewValueSet(wire.Set{ValueType: wire.TBinary, Size: len(v.SetOfStrings), Items: _Set_String_ValueList(v.SetOfStrings)})}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fs[:i]})
}
func (v *PrimitiveContainers) FromWire(w wire.Value) error {
	var err error
	for _, f := range w.GetStruct().Fields {
		switch f.ID {
		case 1:
			if f.Value.Type() == wire.TList {
				v.ListOfBinary, err = _List_Binary_Read(f.Value.GetList())
				if err != nil {
					return err
				}
			}
		case 2:
			if f.Value.Type() == wire.TList {
				v.ListOfInts, err = _List_I64_Read(f.Value.GetList())
				if err != nil {
					return err
				}
			}
		case 5:
			if f.Value.Type() == wire.TMap {
				v.MapOfIntToString, err = _Map_I32_String_Read(f.Value.GetMap())
				if err != nil {
					return err
				}
			}
		case 6:
			if f.Value.Type() == wire.TMap {
				v.MapOfStringToBool, err = _Map_String_Bool_Read(f.Value.GetMap())
				if err != nil {
					return err
				}
			}
		case 4:
			if f.Value.Type() == wire.TSet {
				v.SetOfBytes, err = _Set_Byte_Read(f.Value.GetSet())
				if err != nil {
					return err
				}
			}
		case 3:
			if f.Value.Type() == wire.TSet {
				v.SetOfStrings, err = _Set_String_Read(f.Value.GetSet())
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
func _PrimitiveContainers_Read(w wire.Value) (*PrimitiveContainers, error) {
	var v PrimitiveContainers
	err := v.FromWire(w)
	return &v, err
}

type _List_String_ValueList []string

func (v _List_String_ValueList) ForEach(f func(wire.Value) error) error {
	for _, x := range v {
		err := f(wire.NewValueString(x))
		if err != nil {
			return err
		}
	}
	return nil
}
func (v _List_String_ValueList) Close() {
}

type _Map_I64_Double_MapItemList map[int64]float64

func (m _Map_I64_Double_MapItemList) ForEach(f func(wire.MapItem) error) error {
	for k, v := range m {
		err := f(wire.MapItem{Key: wire.NewValueI64(k), Value: wire.NewValueDouble(v)})
		if err != nil {
			return err
		}
	}
	return nil
}
func (m _Map_I64_Double_MapItemList) Close() {
}

type _Set_I32_ValueList map[int32]struct{}

func (v _Set_I32_ValueList) ForEach(f func(wire.Value) error) error {
	for x := range v {
		err := f(wire.NewValueI32(x))
		if err != nil {
			return err
		}
	}
	return nil
}
func (v _Set_I32_ValueList) Close() {
}
func _List_String_Read(l wire.List) ([]string, error) {
	if l.ValueType != wire.TBinary {
		return nil, nil
	}
	o := make([]string, 0, l.Size)
	err := l.Items.ForEach(func(x wire.Value) error {
		i, err := x.GetString(), error(nil)
		if err != nil {
			return err
		}
		o = append(o, i)
		return nil
	})
	l.Items.Close()
	return o, err
}
func _Map_I64_Double_Read(m wire.Map) (map[int64]float64, error) {
	if m.KeyType != wire.TI64 {
		return nil, nil
	}
	if m.ValueType != wire.TDouble {
		return nil, nil
	}
	o := make(map[int64]float64, m.Size)
	err := m.Items.ForEach(func(x wire.MapItem) error {
		k, err := x.Key.GetI64(), error(nil)
		if err != nil {
			return err
		}
		v, err := x.Value.GetDouble(), error(nil)
		if err != nil {
			return err
		}
		o[k] = v
		return nil
	})
	m.Items.Close()
	return o, err
}
func _Set_I32_Read(s wire.Set) (map[int32]struct{}, error) {
	if s.ValueType != wire.TI32 {
		return nil, nil
	}
	o := make(map[int32]struct{}, s.Size)
	err := s.Items.ForEach(func(x wire.Value) error {
		i, err := x.GetI32(), error(nil)
		if err != nil {
			return err
		}
		o[i] = struct{}{}
		return nil
	})
	s.Items.Close()
	return o, err
}

type PrimitiveContainersRequired struct {
	ListOfStrings      []string
	MapOfIntsToDoubles map[int64]float64
	SetOfInts          map[int32]struct{}
}

func (v *PrimitiveContainersRequired) ToWire() wire.Value {
	var fs [3]wire.Field
	i := 0
	fs[i] = wire.Field{ID: 1, Value: wire.NewValueList(wire.List{ValueType: wire.TBinary, Size: len(v.ListOfStrings), Items: _List_String_ValueList(v.ListOfStrings)})}
	i++
	fs[i] = wire.Field{ID: 3, Value: wire.NewValueMap(wire.Map{KeyType: wire.TI64, ValueType: wire.TDouble, Size: len(v.MapOfIntsToDoubles), Items: _Map_I64_Double_MapItemList(v.MapOfIntsToDoubles)})}
	i++
	fs[i] = wire.Field{ID: 2, Value: wire.NewValueSet(wire.Set{ValueType: wire.TI32, Size: len(v.SetOfInts), Items: _Set_I32_ValueList(v.SetOfInts)})}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fs[:i]})
}
func (v *PrimitiveContainersRequired) FromWire(w wire.Value) error {
	var err error
	for _, f := range w.GetStruct().Fields {
		switch f.ID {
		case 1:
			if f.Value.Type() == wire.TList {
				v.ListOfStrings, err = _List_String_Read(f.Value.GetList())
				if err != nil {
					return err
				}
			}
		case 3:
			if f.Value.Type() == wire.TMap {
				v.MapOfIntsToDoubles, err = _Map_I64_Double_Read(f.Value.GetMap())
				if err != nil {
					return err
				}
			}
		case 2:
			if f.Value.Type() == wire.TSet {
				v.SetOfInts, err = _Set_I32_Read(f.Value.GetSet())
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
func _PrimitiveContainersRequired_Read(w wire.Value) (*PrimitiveContainersRequired, error) {
	var v PrimitiveContainersRequired
	err := v.FromWire(w)
	return &v, err
}

type PrimitiveOptionalStruct struct {
	BinaryField []byte
	BoolField   *bool
	ByteField   *int8
	DoubleField *float64
	Int16Field  *int16
	Int32Field  *int32
	Int64Field  *int64
	StringField *string
}

func (v *PrimitiveOptionalStruct) ToWire() wire.Value {
	var fs [8]wire.Field
	i := 0
	if v.BinaryField != nil {
		fs[i] = wire.Field{ID: 8, Value: wire.NewValueBinary(v.BinaryField)}
		i++
	}
	if v.BoolField != nil {
		fs[i] = wire.Field{ID: 1, Value: wire.NewValueBool(*v.BoolField)}
		i++
	}
	if v.ByteField != nil {
		fs[i] = wire.Field{ID: 2, Value: wire.NewValueI8(*v.ByteField)}
		i++
	}
	if v.DoubleField != nil {
		fs[i] = wire.Field{ID: 6, Value: wire.NewValueDouble(*v.DoubleField)}
		i++
	}
	if v.Int16Field != nil {
		fs[i] = wire.Field{ID: 3, Value: wire.NewValueI16(*v.Int16Field)}
		i++
	}
	if v.Int32Field != nil {
		fs[i] = wire.Field{ID: 4, Value: wire.NewValueI32(*v.Int32Field)}
		i++
	}
	if v.Int64Field != nil {
		fs[i] = wire.Field{ID: 5, Value: wire.NewValueI64(*v.Int64Field)}
		i++
	}
	if v.StringField != nil {
		fs[i] = wire.Field{ID: 7, Value: wire.NewValueString(*v.StringField)}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fs[:i]})
}
func (v *PrimitiveOptionalStruct) FromWire(w wire.Value) error {
	var err error
	for _, f := range w.GetStruct().Fields {
		switch f.ID {
		case 8:
			if f.Value.Type() == wire.TBinary {
				v.BinaryField, err = f.Value.GetBinary(), error(nil)
				if err != nil {
					return err
				}
			}
		case 1:
			if f.Value.Type() == wire.TBool {
				x, err := f.Value.GetBool(), error(nil)
				v.BoolField = &x
				if err != nil {
					return err
				}
			}
		case 2:
			if f.Value.Type() == wire.TI8 {
				x2, err := f.Value.GetI8(), error(nil)
				v.ByteField = &x2
				if err != nil {
					return err
				}
			}
		case 6:
			if f.Value.Type() == wire.TDouble {
				x3, err := f.Value.GetDouble(), error(nil)
				v.DoubleField = &x3
				if err != nil {
					return err
				}
			}
		case 3:
			if f.Value.Type() == wire.TI16 {
				x4, err := f.Value.GetI16(), error(nil)
				v.Int16Field = &x4
				if err != nil {
					return err
				}
			}
		case 4:
			if f.Value.Type() == wire.TI32 {
				x5, err := f.Value.GetI32(), error(nil)
				v.Int32Field = &x5
				if err != nil {
					return err
				}
			}
		case 5:
			if f.Value.Type() == wire.TI64 {
				x6, err := f.Value.GetI64(), error(nil)
				v.Int64Field = &x6
				if err != nil {
					return err
				}
			}
		case 7:
			if f.Value.Type() == wire.TBinary {
				x7, err := f.Value.GetString(), error(nil)
				v.StringField = &x7
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
func _PrimitiveOptionalStruct_Read(w wire.Value) (*PrimitiveOptionalStruct, error) {
	var v PrimitiveOptionalStruct
	err := v.FromWire(w)
	return &v, err
}

type PrimitiveRequiredStruct struct {
	BinaryField []byte
	BoolField   bool
	ByteField   int8
	DoubleField float64
	Int16Field  int16
	Int32Field  int32
	Int64Field  int64
	StringField string
}

func (v *PrimitiveRequiredStruct) ToWire() wire.Value {
	var fs [8]wire.Field
	i := 0
	fs[i] = wire.Field{ID: 8, Value: wire.NewValueBinary(v.BinaryField)}
	i++
	fs[i] = wire.Field{ID: 1, Value: wire.NewValueBool(v.BoolField)}
	i++
	fs[i] = wire.Field{ID: 2, Value: wire.NewValueI8(v.ByteField)}
	i++
	fs[i] = wire.Field{ID: 6, Value: wire.NewValueDouble(v.DoubleField)}
	i++
	fs[i] = wire.Field{ID: 3, Value: wire.NewValueI16(v.Int16Field)}
	i++
	fs[i] = wire.Field{ID: 4, Value: wire.NewValueI32(v.Int32Field)}
	i++
	fs[i] = wire.Field{ID: 5, Value: wire.NewValueI64(v.Int64Field)}
	i++
	fs[i] = wire.Field{ID: 7, Value: wire.NewValueString(v.StringField)}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fs[:i]})
}
func (v *PrimitiveRequiredStruct) FromWire(w wire.Value) error {
	var err error
	for _, f := range w.GetStruct().Fields {
		switch f.ID {
		case 8:
			if f.Value.Type() == wire.TBinary {
				v.BinaryField, err = f.Value.GetBinary(), error(nil)
				if err != nil {
					return err
				}
			}
		case 1:
			if f.Value.Type() == wire.TBool {
				v.BoolField, err = f.Value.GetBool(), error(nil)
				if err != nil {
					return err
				}
			}
		case 2:
			if f.Value.Type() == wire.TI8 {
				v.ByteField, err = f.Value.GetI8(), error(nil)
				if err != nil {
					return err
				}
			}
		case 6:
			if f.Value.Type() == wire.TDouble {
				v.DoubleField, err = f.Value.GetDouble(), error(nil)
				if err != nil {
					return err
				}
			}
		case 3:
			if f.Value.Type() == wire.TI16 {
				v.Int16Field, err = f.Value.GetI16(), error(nil)
				if err != nil {
					return err
				}
			}
		case 4:
			if f.Value.Type() == wire.TI32 {
				v.Int32Field, err = f.Value.GetI32(), error(nil)
				if err != nil {
					return err
				}
			}
		case 5:
			if f.Value.Type() == wire.TI64 {
				v.Int64Field, err = f.Value.GetI64(), error(nil)
				if err != nil {
					return err
				}
			}
		case 7:
			if f.Value.Type() == wire.TBinary {
				v.StringField, err = f.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
func _PrimitiveRequiredStruct_Read(w wire.Value) (*PrimitiveRequiredStruct, error) {
	var v PrimitiveRequiredStruct
	err := v.FromWire(w)
	return &v, err
}

type Size struct {
	Height float64
	Width  float64
}

func (v *Size) ToWire() wire.Value {
	var fs [2]wire.Field
	i := 0
	fs[i] = wire.Field{ID: 2, Value: wire.NewValueDouble(v.Height)}
	i++
	fs[i] = wire.Field{ID: 1, Value: wire.NewValueDouble(v.Width)}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fs[:i]})
}
func (v *Size) FromWire(w wire.Value) error {
	var err error
	for _, f := range w.GetStruct().Fields {
		switch f.ID {
		case 2:
			if f.Value.Type() == wire.TDouble {
				v.Height, err = f.Value.GetDouble(), error(nil)
				if err != nil {
					return err
				}
			}
		case 1:
			if f.Value.Type() == wire.TDouble {
				v.Width, err = f.Value.GetDouble(), error(nil)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
func _Size_Read(w wire.Value) (*Size, error) {
	var v Size
	err := v.FromWire(w)
	return &v, err
}

type User struct {
	Contact *ContactInfo
	Name    string
}

func (v *User) ToWire() wire.Value {
	var fs [2]wire.Field
	i := 0
	if v.Contact != nil {
		fs[i] = wire.Field{ID: 2, Value: v.Contact.ToWire()}
		i++
	}
	fs[i] = wire.Field{ID: 1, Value: wire.NewValueString(v.Name)}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fs[:i]})
}
func (v *User) FromWire(w wire.Value) error {
	var err error
	for _, f := range w.GetStruct().Fields {
		switch f.ID {
		case 2:
			if f.Value.Type() == wire.TStruct {
				v.Contact, err = _ContactInfo_Read(f.Value)
				if err != nil {
					return err
				}
			}
		case 1:
			if f.Value.Type() == wire.TBinary {
				v.Name, err = f.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
func _User_Read(w wire.Value) (*User, error) {
	var v User
	err := v.FromWire(w)
	return &v, err
}