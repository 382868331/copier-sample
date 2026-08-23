package copier
import("reflect";"testing")
var _=reflect.TypeOf
func TestGoletaCopier007(t *testing.T){v:=reflect.ValueOf(0);if !shouldIgnore(v,0,true)||shouldIgnore(v,tagOverride,true){t.Fatal("override semantics inverted")}}
