package copier
import("reflect";"testing")
var _=reflect.TypeOf
func TestGoletaCopier003(t *testing.T){src:=struct{A int}{7};dst:=struct{A int}{};if e:=Copy(&dst,&src);e!=nil||dst.A!=7{t.Fatalf("dst=%+v err=%v",dst,e)}}

func TestGoletaCopier003AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	src:=struct{Name string}{"x"};dst:=struct{Name string}{};if e:=Copy(&dst,src);e!=nil||dst.Name!="x"{t.Fatalf("dst=%+v err=%v",dst,e)}
}
