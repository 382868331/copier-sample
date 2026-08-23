package copier
import("reflect";"testing")
var _=reflect.TypeOf
func TestGoletaCopier005(t *testing.T){a:=reflect.TypeOf(struct{A int}{});b:=reflect.TypeOf(struct{B int}{});pair:=converterPair{SrcType:a,DstType:b};maps:=map[converterPair]FieldNameMapping{pair:{Mapping:map[string]string{"A":"B"}}};if getFieldNamesMapping(maps,a,b)["A"]!="B"{t.Fatal("mapping not found")}}

func TestGoletaCopier005AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	a:=reflect.TypeOf(struct{X int}{});b:=reflect.TypeOf(struct{Y int}{});pair:=converterPair{SrcType:a,DstType:b};if getFieldNamesMapping(map[converterPair]FieldNameMapping{pair:{Mapping:map[string]string{"X":"Y"}}},a,b)["X"]!="Y"{t.Fatal("mapping not found")}
}
