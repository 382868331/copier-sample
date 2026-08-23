package copier
import("reflect";"testing")
var _=reflect.TypeOf
func TestGoletaCopier010(t *testing.T){typ,ptr:=indirectType(reflect.TypeOf([]int{}));if typ.Kind()!=reflect.Int||!ptr{t.Fatalf("typ=%v ptr=%v",typ,ptr)}}

func TestGoletaCopier010AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	typ,ptr:=indirectType(reflect.TypeOf([]*string{}));if typ.Kind()!=reflect.String||!ptr{t.Fatalf("typ=%v ptr=%v",typ,ptr)}
}
