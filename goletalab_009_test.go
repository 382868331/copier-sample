package copier
import("reflect";"testing")
var _=reflect.TypeOf
func TestGoletaCopier009(t *testing.T){n:=7;p:=&n;pp:=&p;if got:=indirect(reflect.ValueOf(pp));got.Kind()!=reflect.Int||got.Int()!=7{t.Fatalf("got=%v",got.Kind())}}

func TestGoletaCopier009AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	s:="x";p:=&s;pp:=&p;ppp:=&pp;if got:=indirect(reflect.ValueOf(ppp));got.Kind()!=reflect.String{t.Fatalf("got=%v",got.Kind())}
}
