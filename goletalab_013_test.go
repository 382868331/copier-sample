package copier
import("reflect";"testing")
var _=reflect.TypeOf
func TestGoletaCopier013(t *testing.T){to:=reflect.ValueOf(new(string)).Elem();to.SetString("old");from:=reflect.ValueOf(1);pair:=converterPair{SrcType:from.Type(),DstType:to.Type()};ok,e:=lookupAndCopyWithConverter(to,from,map[converterPair]TypeConverter{pair:{Fn:func(interface{})(interface{},error){return nil,nil}}});if e!=nil||!ok||to.String()!=""{t.Fatalf("to=%q",to.String())}}

func TestGoletaCopier013AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	to:=reflect.ValueOf(new(*int)).Elem();from:=reflect.ValueOf("x");pair:=converterPair{SrcType:from.Type(),DstType:to.Type()};ok,_:=lookupAndCopyWithConverter(to,from,map[converterPair]TypeConverter{pair:{Fn:func(interface{})(interface{},error){return nil,nil}}});if !ok||!to.IsNil(){t.Fatal("nil not assigned")}
}
