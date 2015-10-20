// This file was generated by nomdl/codegen.

package test

import (
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

var __testPackageInFile_leafDep_CachedRef = __testPackageInFile_leafDep_Ref()

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func __testPackageInFile_leafDep_Ref() ref.Ref {
	p := types.NewPackage([]types.TypeRef{
		types.MakeStructTypeRef("S",
			[]types.Field{
				types.Field{"s", types.MakePrimitiveTypeRef(types.StringKind), false},
				types.Field{"b", types.MakePrimitiveTypeRef(types.BoolKind), false},
			},
			types.Choices{},
		),
		types.MakeEnumTypeRef("E", "e1", "e2", "e3"),
	}, []ref.Ref{})
	return types.RegisterPackage(&p)
}

// S

type S struct {
	m types.Map
}

func NewS() S {
	return S{types.NewMap(
		types.NewString("$type"), types.MakeTypeRef(__testPackageInFile_leafDep_CachedRef, 0),
		types.NewString("s"), types.NewString(""),
		types.NewString("b"), types.Bool(false),
	)}
}

type SDef struct {
	S string
	B bool
}

func (def SDef) New() S {
	return S{
		types.NewMap(
			types.NewString("$type"), types.MakeTypeRef(__testPackageInFile_leafDep_CachedRef, 0),
			types.NewString("s"), types.NewString(def.S),
			types.NewString("b"), types.Bool(def.B),
		)}
}

func (s S) Def() (d SDef) {
	d.S = s.m.Get(types.NewString("s")).(types.String).String()
	d.B = bool(s.m.Get(types.NewString("b")).(types.Bool))
	return
}

var __typeRefForS = types.MakeTypeRef(__testPackageInFile_leafDep_CachedRef, 0)

func (m S) TypeRef() types.TypeRef {
	return __typeRefForS
}

func init() {
	types.RegisterFromValFunction(__typeRefForS, func(v types.Value) types.NomsValue {
		return SFromVal(v)
	})
}

func SFromVal(val types.Value) S {
	// TODO: Validate here
	return S{val.(types.Map)}
}

func (s S) NomsValue() types.Value {
	return s.m
}

func (s S) Equals(other types.Value) bool {
	if other, ok := other.(S); ok {
		return s.m.Equals(other.m)
	}
	return false
}

func (s S) Ref() ref.Ref {
	return s.m.Ref()
}

func (s S) Chunks() (futures []types.Future) {
	futures = append(futures, s.TypeRef().Chunks()...)
	futures = append(futures, s.m.Chunks()...)
	return
}

func (s S) S() string {
	return s.m.Get(types.NewString("s")).(types.String).String()
}

func (s S) SetS(val string) S {
	return S{s.m.Set(types.NewString("s"), types.NewString(val))}
}

func (s S) B() bool {
	return bool(s.m.Get(types.NewString("b")).(types.Bool))
}

func (s S) SetB(val bool) S {
	return S{s.m.Set(types.NewString("b"), types.Bool(val))}
}

// E

type E uint32

const (
	E1 E = iota
	E2
	E3
)
