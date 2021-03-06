// Code generated by immutableGen. DO NOT EDIT.

package imm

//immutableVet:skipFile

import (
	"myitcv.io/immutable"
)

//
// strEntrySelect is an immutable type and has the following template:
//
// 	map[string]Label
//
type strEntrySelect struct {
	theMap  map[string]Label
	mutable bool
	__tmpl  _Imm_strEntrySelect
}

var _ immutable.Immutable = new(strEntrySelect)
var _ = new(strEntrySelect).__tmpl

func newStrEntrySelect(inits ...func(m *strEntrySelect)) *strEntrySelect {
	res := newStrEntrySelectCap(0)
	if len(inits) == 0 {
		return res
	}

	return res.WithMutable(func(m *strEntrySelect) {
		for _, i := range inits {
			i(m)
		}
	})
}

func newStrEntrySelectCap(l int) *strEntrySelect {
	return &strEntrySelect{
		theMap: make(map[string]Label, l),
	}
}

func (m *strEntrySelect) Mutable() bool {
	return m.mutable
}

func (m *strEntrySelect) Len() int {
	if m == nil {
		return 0
	}

	return len(m.theMap)
}

func (m *strEntrySelect) Get(k string) (Label, bool) {
	v, ok := m.theMap[k]
	return v, ok
}

func (m *strEntrySelect) AsMutable() *strEntrySelect {
	if m == nil {
		return nil
	}

	if m.Mutable() {
		return m
	}

	res := m.dup()
	res.mutable = true

	return res
}

func (m *strEntrySelect) dup() *strEntrySelect {
	resMap := make(map[string]Label, len(m.theMap))

	for k := range m.theMap {
		resMap[k] = m.theMap[k]
	}

	res := &strEntrySelect{
		theMap: resMap,
	}

	return res
}

func (m *strEntrySelect) AsImmutable(v *strEntrySelect) *strEntrySelect {
	if m == nil {
		return nil
	}

	if v == m {
		return m
	}

	m.mutable = false
	return m
}

func (m *strEntrySelect) Range() map[string]Label {
	if m == nil {
		return nil
	}

	return m.theMap
}

func (m *strEntrySelect) WithMutable(f func(mi *strEntrySelect)) *strEntrySelect {
	res := m.AsMutable()
	f(res)
	res = res.AsImmutable(m)

	return res
}

func (m *strEntrySelect) WithImmutable(f func(mi *strEntrySelect)) *strEntrySelect {
	prev := m.mutable
	m.mutable = false
	f(m)
	m.mutable = prev

	return m
}

func (m *strEntrySelect) Set(k string, v Label) *strEntrySelect {
	if m.mutable {
		m.theMap[k] = v
		return m
	}

	res := m.dup()
	res.theMap[k] = v

	return res
}

func (m *strEntrySelect) Del(k string) *strEntrySelect {
	if _, ok := m.theMap[k]; !ok {
		return m
	}

	if m.mutable {
		delete(m.theMap, k)
		return m
	}

	res := m.dup()
	delete(res.theMap, k)

	return res
}
func (s *strEntrySelect) IsDeeplyNonMutable(seen map[interface{}]bool) bool {
	if s == nil {
		return true
	}

	if s.Mutable() {
		return false
	}
	return true
}

//
// LabelEntries is an immutable type and has the following template:
//
// 	[]Label
//
type LabelEntries struct {
	theSlice []Label
	mutable  bool
	__tmpl   _Imm_LabelEntries
}

var _ immutable.Immutable = new(LabelEntries)
var _ = new(LabelEntries).__tmpl

func NewLabelEntries(s ...Label) *LabelEntries {
	c := make([]Label, len(s))
	copy(c, s)

	return &LabelEntries{
		theSlice: c,
	}
}

func NewLabelEntriesLen(l int) *LabelEntries {
	c := make([]Label, l)

	return &LabelEntries{
		theSlice: c,
	}
}

func (m *LabelEntries) Mutable() bool {
	return m.mutable
}

func (m *LabelEntries) Len() int {
	if m == nil {
		return 0
	}

	return len(m.theSlice)
}

func (m *LabelEntries) Get(i int) Label {
	return m.theSlice[i]
}

func (m *LabelEntries) AsMutable() *LabelEntries {
	if m == nil {
		return nil
	}

	if m.Mutable() {
		return m
	}

	res := m.dup()
	res.mutable = true

	return res
}

func (m *LabelEntries) dup() *LabelEntries {
	resSlice := make([]Label, len(m.theSlice))

	for i := range m.theSlice {
		resSlice[i] = m.theSlice[i]
	}

	res := &LabelEntries{
		theSlice: resSlice,
	}

	return res
}

func (m *LabelEntries) AsImmutable(v *LabelEntries) *LabelEntries {
	if m == nil {
		return nil
	}

	if v == m {
		return m
	}

	m.mutable = false
	return m
}

func (m *LabelEntries) Range() []Label {
	if m == nil {
		return nil
	}

	return m.theSlice
}

func (m *LabelEntries) WithMutable(f func(mi *LabelEntries)) *LabelEntries {
	res := m.AsMutable()
	f(res)
	res = res.AsImmutable(m)

	return res
}

func (m *LabelEntries) WithImmutable(f func(mi *LabelEntries)) *LabelEntries {
	prev := m.mutable
	m.mutable = false
	f(m)
	m.mutable = prev

	return m
}

func (m *LabelEntries) Set(i int, v Label) *LabelEntries {
	if m.mutable {
		m.theSlice[i] = v
		return m
	}

	res := m.dup()
	res.theSlice[i] = v

	return res
}

func (m *LabelEntries) Append(v ...Label) *LabelEntries {
	if m.mutable {
		m.theSlice = append(m.theSlice, v...)
		return m
	}

	res := m.dup()
	res.theSlice = append(res.theSlice, v...)

	return res
}
func (s *LabelEntries) IsDeeplyNonMutable(seen map[interface{}]bool) bool {
	if s == nil {
		return true
	}

	if s.Mutable() {
		return false
	}
	return true
}

//
// entriesKeysSelect is an immutable type and has the following template:
//
// 	[]entryKey
//
type entriesKeysSelect struct {
	theSlice []entryKey
	mutable  bool
	__tmpl   _Imm_entriesKeysSelect
}

var _ immutable.Immutable = new(entriesKeysSelect)
var _ = new(entriesKeysSelect).__tmpl

func newEntriesKeysSelect(s ...entryKey) *entriesKeysSelect {
	c := make([]entryKey, len(s))
	copy(c, s)

	return &entriesKeysSelect{
		theSlice: c,
	}
}

func newEntriesKeysSelectLen(l int) *entriesKeysSelect {
	c := make([]entryKey, l)

	return &entriesKeysSelect{
		theSlice: c,
	}
}

func (m *entriesKeysSelect) Mutable() bool {
	return m.mutable
}

func (m *entriesKeysSelect) Len() int {
	if m == nil {
		return 0
	}

	return len(m.theSlice)
}

func (m *entriesKeysSelect) Get(i int) entryKey {
	return m.theSlice[i]
}

func (m *entriesKeysSelect) AsMutable() *entriesKeysSelect {
	if m == nil {
		return nil
	}

	if m.Mutable() {
		return m
	}

	res := m.dup()
	res.mutable = true

	return res
}

func (m *entriesKeysSelect) dup() *entriesKeysSelect {
	resSlice := make([]entryKey, len(m.theSlice))

	for i := range m.theSlice {
		resSlice[i] = m.theSlice[i]
	}

	res := &entriesKeysSelect{
		theSlice: resSlice,
	}

	return res
}

func (m *entriesKeysSelect) AsImmutable(v *entriesKeysSelect) *entriesKeysSelect {
	if m == nil {
		return nil
	}

	if v == m {
		return m
	}

	m.mutable = false
	return m
}

func (m *entriesKeysSelect) Range() []entryKey {
	if m == nil {
		return nil
	}

	return m.theSlice
}

func (m *entriesKeysSelect) WithMutable(f func(mi *entriesKeysSelect)) *entriesKeysSelect {
	res := m.AsMutable()
	f(res)
	res = res.AsImmutable(m)

	return res
}

func (m *entriesKeysSelect) WithImmutable(f func(mi *entriesKeysSelect)) *entriesKeysSelect {
	prev := m.mutable
	m.mutable = false
	f(m)
	m.mutable = prev

	return m
}

func (m *entriesKeysSelect) Set(i int, v entryKey) *entriesKeysSelect {
	if m.mutable {
		m.theSlice[i] = v
		return m
	}

	res := m.dup()
	res.theSlice[i] = v

	return res
}

func (m *entriesKeysSelect) Append(v ...entryKey) *entriesKeysSelect {
	if m.mutable {
		m.theSlice = append(m.theSlice, v...)
		return m
	}

	res := m.dup()
	res.theSlice = append(res.theSlice, v...)

	return res
}
func (s *entriesKeysSelect) IsDeeplyNonMutable(seen map[interface{}]bool) bool {
	if s == nil {
		return true
	}

	if s.Mutable() {
		return false
	}
	return true
}
