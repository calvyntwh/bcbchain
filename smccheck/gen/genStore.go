package gen

import (
	"github.com/bcbchain/bcbchain/smccheck/parsecode"
	"bytes"
	"go/ast"
	"path/filepath"
	"strings"
	"text/template"
)

var storeTemplate = `package {{.PackageName}}

{{if .Imports}}import ({{end}}
  {{range $v,$vv := .Imports}}
{{$v.Name}} {{$v.Path}}{{end}}
{{if .Imports}}){{end}}

// This file is auto generated by BCB-goland-plugin.
// Don't modified it

{{range $i, $s := .Stores}}
// {{$s|expNames}} {{$s|expType}}
//@:public:store{{$isMap := $s | isM}}{{$isLit := $s | isL}}{{$isBnN := $s | isN}}{{$isStar := $s | isS}}
{{if $isMap}}{{$isMLit := $s|isML}}{{$isMVStar := $s|isMS}}{{$isVMap := $s|isVM}}{{if (isNV ($s|expV))}}
func ({{$.ReceiverName}} *{{$.ContractName}}) _{{$s|expNames}}({{expK $s 0}}) bn.Number {
	temp := bn.N(0)
	return *{{$.ReceiverName}}.sdk.Helper().StateHelper().GetEx(fmt.Sprintf("/{{$s|expNames}}{{expK2K $s 0}}), &temp).(*bn.Number)
}{{else}}
func ({{$.ReceiverName}} *{{$.ContractName}}) _{{$s|expNames}}({{expK $s 0}}) {{$s|expV}} {
	return {{mp $s}}{{$.ReceiverName}}.sdk.Helper().StateHelper().GetEx(fmt.Sprintf("/{{$s|expNames}}{{expK2K $s 0}}), {{md $s}}).({{mr $s}})
}{{end}}
func ({{$.ReceiverName}} *{{$.ContractName}}) _chk{{$s|expNames|upperFirst}}({{expK $s 0}}) bool {
	return {{$.ReceiverName}}.sdk.Helper().StateHelper().Check(fmt.Sprintf("/{{$s|expNames}}{{expK2K $s 0}}))
}
func ({{$.ReceiverName}} *{{$.ContractName}}) _set{{$s|expNames|upperFirst}}({{expK $s 0}}, v {{$s|expV}}) {
	{{$.ReceiverName}}.sdk.Helper().StateHelper().Set(fmt.Sprintf("/{{$s|expNames}}{{expK2K $s 0}}), {{if $isMVStar}}v{{else}}&v{{end}})
}
func ({{$.ReceiverName}} *{{$.ContractName}}) _del{{$s|expNames|upperFirst}}({{expK $s 0}}) {
	{{$.ReceiverName}}.sdk.Helper().StateHelper().Delete(fmt.Sprintf("/{{$s|expNames}}{{expK2K $s 0}}))
}
{{else}}{{if $isBnN}}
func ({{$.ReceiverName}} *{{$.ContractName}}) _{{$s|expNames}}() bn.Number {
	temp := bn.N(0)
	return *{{$.ReceiverName}}.sdk.Helper().StateHelper().GetEx("/{{$s|expNames}}", &temp).(*bn.Number)
}{{else}}
func ({{$.ReceiverName}} *{{$.ContractName}}) _{{$s|expNames}}() {{$s|expType}} {
	return {{if not $isStar}}*{{end}}{{$.ReceiverName}}.sdk.Helper().StateHelper().GetEx("/{{$s|expNames}}", {{md $s}}).(*{{$s|expNoS}})
}{{end}}
func ({{$.ReceiverName}} *{{$.ContractName}}) _chk{{$s|expNames|upperFirst}}() bool {
	return {{$.ReceiverName}}.sdk.Helper().StateHelper().Check("/{{$s|expNames}}")
}
func ({{$.ReceiverName}} *{{$.ContractName}}) _set{{$s|expNames|upperFirst}}(v {{$s|expType}}) {
	{{$.ReceiverName}}.sdk.Helper().StateHelper().Set("/{{$s|expNames}}", {{if $isStar}}v{{else}}&v{{end}})
}
func ({{$.ReceiverName}} *{{$.ContractName}}) _del{{$s|expNames|upperFirst}}() {
	{{$.ReceiverName}}.sdk.Helper().StateHelper().Delete("/{{$s|expNames}}")
}
{{end}}
{{end}}

{{range $i, $c := .Caches}}
// {{$c|expNames}} {{$c|expType}}
//@:public:store:cache{{$isLit := $c|isL}}{{$isMap := $c|isM}}{{$isBnN := $c | isN}}{{$isStar := $c | isS}}
{{if $isMap}}{{$isMLit := $c|isML}}{{$isMVStar := $c|isMS}}{{$isVMap := $c|isVM}}
func ({{$.ReceiverName}} *{{$.ContractName}}) _set{{$c|expNames|upperFirst}}({{expK $c 0}}, v {{$c|expV}}) {
	{{$.ReceiverName}}.sdk.Helper().StateHelper().McSet(fmt.Sprintf("/{{$c|expNames}}{{expK2K $c 0}}), {{if $isMVStar}}v{{else}}&v{{end}})
}{{if (isNV ($c|expV))}}
func ({{$.ReceiverName}} *{{$.ContractName}}) _{{$c|expNames}}({{expK $c 0}}) bn.Number {
	temp := bn.N(0)
	return *{{$.ReceiverName}}.sdk.Helper().StateHelper().McGetEx(fmt.Sprintf("/{{$c|expNames}}{{expK2K $c 0}}), &temp).(*bn.Number)
}{{else}}
func ({{$.ReceiverName}} *{{$.ContractName}}) _{{$c|expNames}}({{expK $c 0}}) {{$c|expV}} {
	return {{mp $c}}{{$.ReceiverName}}.sdk.Helper().StateHelper().McGetEx(fmt.Sprintf("/{{$c|expNames}}{{expK2K $c 0}}), {{md $c}}).({{mr $c}})
}{{end}}
func ({{$.ReceiverName}} *{{$.ContractName}}) _clr{{$c|expNames|upperFirst}}({{expK $c 0}}) {
	{{$.ReceiverName}}.sdk.Helper().StateHelper().McClear(fmt.Sprintf("/{{$c|expNames}}{{expK2K $c 0}}))
}
func ({{$.ReceiverName}} *{{$.ContractName}}) _chk{{$c|expNames|upperFirst}}({{expK $c 0}}) bool {
	return {{$.ReceiverName}}.sdk.Helper().StateHelper().Check(fmt.Sprintf("/{{$c|expNames}}{{expK2K $c 0}}))
}
func ({{$.ReceiverName}} *{{$.ContractName}}) _del{{$c|expNames|upperFirst}}({{expK $c 0}}) {
	{{$.ReceiverName}}.sdk.Helper().StateHelper().McDelete(fmt.Sprintf("/{{$c|expNames}}{{expK2K $c 0}}))
}
{{else}}{{if $isBnN}}
func ({{$.ReceiverName}} *{{$.ContractName}}) _{{$c|expNames}}() bn.Number {
	temp := bn.N(0)
	return *{{$.ReceiverName}}.sdk.Helper().StateHelper().McGetEx("/{{$c|expNames}}", &temp).(*bn.Number)
} {{else}}
func ({{$.ReceiverName}} *{{$.ContractName}}) _{{$c|expNames}}() {{$c|expType}} {
	return {{if not $isStar}}*{{end}}{{$.ReceiverName}}.sdk.Helper().StateHelper().McGetEx("/{{$c|expNames}}", {{md $c}}).({{mr $c}})
}{{end}}
func ({{$.ReceiverName}} *{{$.ContractName}}) _chk{{$c|expNames|upperFirst}}() bool {
	return {{$.ReceiverName}}.sdk.Helper().StateHelper().McCheck("/{{$c|expNames}}")
}
func ({{$.ReceiverName}} *{{$.ContractName}}) _clr{{$c|expNames|upperFirst}}() {
	{{$.ReceiverName}}.sdk.Helper().StateHelper().McClear("/{{$c|expNames}}")
}
func ({{$.ReceiverName}} *{{$.ContractName}}) _set{{$c|expNames|upperFirst}}(v {{$c|expType}}) {
	{{$.ReceiverName}}.sdk.Helper().StateHelper().McSet("/{{$c|expNames}}", {{if $isStar}}v{{else}}&v{{end}})
}
func ({{$.ReceiverName}} *{{$.ContractName}}) _del{{$c|expNames|upperFirst}}() {
	{{$.ReceiverName}}.sdk.Helper().StateHelper().McDelete("/{{$c|expNames}}")
}
{{end}}
{{end}}
`

type storeExport struct {
	baseExport
	Imports map[parsecode.Import]struct{}
	Stores  []parsecode.Field
	Caches  []parsecode.Field
}

func res2Store(res *parsecode.Result) storeExport {
	store := storeExport{}
	store.PackageName = res.PackageName
	store.ReceiverName = strings.ToLower(string([]rune(res.ContractStructure)[0]))
	store.ContractName = res.ContractStructure
	imports := make(map[parsecode.Import]struct{})
	stores := make([]parsecode.Field, 0)
	caches := make([]parsecode.Field, 0)
	for _, s := range res.Stores {
		stores = append(stores, parsecode.FieldsExpand(s)...)
		for imp := range s.RelatedImport {
			imports[imp] = struct{}{}
		}
		if IsMap(s) {
			imports[parsecode.Import{Path: "\"fmt\""}] = struct{}{}
		}
	}
	for _, c := range res.StoreCaches {
		caches = append(caches, parsecode.FieldsExpand(c)...)
		for imp := range c.RelatedImport {
			imports[imp] = struct{}{}
		}
		if IsMap(c) {
			imports[parsecode.Import{Path: "\"fmt\""}] = struct{}{}
		}
	}
	store.Stores = stores
	store.Caches = caches
	store.Imports = imports
	return store
}

func GenStore(inPath string, res *parsecode.Result) {
	filename := filepath.Join(inPath, res.PackageName+"_autogen_store.go")

	funcMap := template.FuncMap{
		"upperFirst": parsecode.UpperFirst,
		"lowerFirst": parsecode.LowerFirst,
		"expNames":   parsecode.ExpandNames,
		"expType":    parsecode.ExpandType,
		"expNoS":     parsecode.ExpandTypeNoStar,
		"expK":       parsecode.ExpandMapFieldKey,
		"expV":       parsecode.ExpandMapFieldVal,
		"expVNoS":    parsecode.ExpandMapFieldValNoStar,
		"expK2K":     parsecode.ExpandMapFieldKeyToKey,
		"isM":        IsMap,
		"isS":        isStar,
		"isMS":       isMapFieldValStar,
		"isL":        IsLiteralType,
		"isML":       isMapValLiteral,
		"isN":        IsBnNumber,
		"isNV":       isBnNumberValue,
		"isVM":       isMapFieldValMap,
		"mp":         makePrefixStar,
		"mr":         makeReturnStr,
		"md":         makeDefaultValueStr,
	}
	tmpl, err := template.New("store").Funcs(funcMap).Parse(storeTemplate)
	if err != nil {
		panic(err)
	}

	store := res2Store(res)

	var buf bytes.Buffer

	if err = tmpl.Execute(&buf, store); err != nil {
		panic(err)
	}

	if err := parsecode.FmtAndWrite(filename, buf.String()); err != nil {
		panic(err)
	}

}

func IsMap(f parsecode.Field) bool {
	_, m := f.FieldType.(*ast.MapType)
	return m
}

func isStar(f parsecode.Field) bool {
	_, s := f.FieldType.(*ast.StarExpr)
	return s
}

func IsLiteralType(f parsecode.Field) bool {
	id, ok := f.FieldType.(*ast.Ident)
	if !ok {
		id1, ok := f.FieldType.(*ast.SelectorExpr)
		if ok {
			if _, ok = parsecode.LiteralTypes[id1.Sel.Name]; ok {
				return true
			}
		}

		id2, ok := f.FieldType.(*ast.StarExpr)
		if ok {
			if id, ok = id2.X.(*ast.Ident); ok {
				if _, ok = parsecode.LiteralTypes[id.Name]; ok {
					return true
				}
			} else if id1, ok = id2.X.(*ast.SelectorExpr); ok {
				if _, ok = parsecode.LiteralTypes[id1.Sel.Name]; ok {
					return true
				}
			}
		}
		return false
	}
	if _, ok2 := parsecode.LiteralTypes[id.Name]; ok2 {
		return true
	}
	return false
}

func IsLiteralTypeEx(f parsecode.Field) bool {

	_, ok := f.FieldType.(*ast.Ident)
	if !ok {
		id3, ok := f.FieldType.(*ast.ArrayType)
		if ok {
			return IsLiteralType(parsecode.Field{FieldType: id3.Elt})
		}
	}

	return false
}

func IsBnNumber(f parsecode.Field) bool {
	id, ok := f.FieldType.(*ast.SelectorExpr)
	if !ok {
		return false
	}
	if id.Sel.Name == parsecode.NumberType {
		return true
	}
	return false
}

func isBnNumberValue(name string) bool {
	splitName := strings.Split(name, ".")
	if len(splitName) == 2 && splitName[1] == parsecode.NumberType {
		return true
	}

	return false
}

func isMapValLiteral(f parsecode.Field) bool {
	mt, okM := f.FieldType.(*ast.MapType)
	if !okM {
		return false
	}
	id, ok := mt.Value.(*ast.Ident)
	if !ok {
		id1, ok := mt.Value.(*ast.SelectorExpr)
		if ok {
			if _, ok = parsecode.LiteralTypes[id1.Sel.Name]; ok {
				return true
			}
		}

		id2, ok := mt.Value.(*ast.StarExpr)
		if ok {
			if id, ok = id2.X.(*ast.Ident); ok {
				if _, ok = parsecode.LiteralTypes[id.Name]; ok {
					return true
				}
			} else if id1, ok = id2.X.(*ast.SelectorExpr); ok {
				if _, ok = parsecode.LiteralTypes[id1.Sel.Name]; ok {
					return true
				}
			}
		}
		return false
	}
	if _, ok2 := parsecode.LiteralTypes[id.Name]; ok2 {
		return true
	}
	return false
}

func isMapFieldValStar(f parsecode.Field) bool {
	m, ok := f.FieldType.(*ast.MapType)
	if !ok {
		return false
	}
	_, okStar := m.Value.(*ast.StarExpr)
	return okStar
}

func isMapFieldValMap(f parsecode.Field) bool {
	m, ok := f.FieldType.(*ast.MapType)
	if !ok {
		return false
	}
	_, okMap := m.Value.(*ast.MapType)
	return okMap
}

func makeDefaultValueStr(f parsecode.Field) string {
	if IsLiteralType(f) {
		return "new(" + parsecode.ExpandTypeNoStar(f) + ")"
	} else if IsMap(f) {
		m, ok := f.FieldType.(*ast.MapType)
		if !ok {
			return ""
		}

		if m1, ok := m.Value.(*ast.MapType); ok {
			m = m1
		}

		return makeDefaultValueStr(parsecode.Field{FieldType: m.Value})
	} else {
		return "&" + parsecode.ExpandTypeNoStar(f) + "{}"
	}
}

func makeReturnStr(f parsecode.Field) string {
	if IsMap(f) {
		m, ok := f.FieldType.(*ast.MapType)
		if !ok {
			return ""
		}

		if m1, ok := m.Value.(*ast.MapType); ok {
			m = m1
		}

		return "*" + parsecode.ExpandTypeNoStar(parsecode.Field{FieldType: m.Value})
	} else {
		return "*" + parsecode.ExpandTypeNoStar(f)
	}
}

func makePrefixStar(f parsecode.Field) string {
	m, ok := f.FieldType.(*ast.MapType)
	if !ok {
		return ""
	}

	if m1, ok := m.Value.(*ast.MapType); ok {
		m = m1
	}

	resultStr := ""
	if !isStar(parsecode.Field{FieldType: m.Value}) {
		resultStr += "*"
	}

	return resultStr
}