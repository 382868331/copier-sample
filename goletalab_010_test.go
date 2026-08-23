package copier
import("reflect";"testing")
var _=reflect.TypeOf
func TestGoletaCopier010(t *testing.T){typ,ptr:=indirectType(reflect.TypeOf([]int{}));if typ.Kind()!=reflect.Int||!ptr{t.Fatalf("typ=%v ptr=%v",typ,ptr)}}
