package copier
import("reflect";"testing")
var _=reflect.TypeOf
func TestGoletaCopier002(t *testing.T){m:=FieldNameMapping{SrcType:struct{A int}{},DstType:struct{B int}{},Mapping:map[string]string{"A":"B"}};if len((Option{FieldNameMapping:[]FieldNameMapping{m}}).fieldNameMapping())!=1{t.Fatal("mapping missing")}}

func TestGoletaCopier002AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	m:=FieldNameMapping{SrcType:struct{X int}{},DstType:struct{Y int}{},Mapping:map[string]string{"X":"Y"}};if len((Option{FieldNameMapping:[]FieldNameMapping{m}}).fieldNameMapping())!=1{t.Fatal("mapping missing")}
}
