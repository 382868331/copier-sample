package copier
import("reflect";"testing")
var _=reflect.TypeOf
func TestGoletaCopier002(t *testing.T){m:=FieldNameMapping{SrcType:struct{A int}{},DstType:struct{B int}{},Mapping:map[string]string{"A":"B"}};if len((Option{FieldNameMapping:[]FieldNameMapping{m}}).fieldNameMapping())!=1{t.Fatal("mapping missing")}}
