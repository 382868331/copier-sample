package copier
import("reflect";"testing")
var _=reflect.TypeOf
func TestGoletaCopier012(t *testing.T){to:=reflect.ValueOf(new(string)).Elem();from:=reflect.ValueOf(7);pair:=converterPair{SrcType:from.Type(),DstType:to.Type()};ok,e:=lookupAndCopyWithConverter(to,from,map[converterPair]TypeConverter{pair:{Fn:func(interface{})(interface{},error){return "seven",nil}}});if e!=nil||!ok||to.String()!="seven"{t.Fatalf("ok=%v to=%q err=%v",ok,to.String(),e)}}

func TestGoletaCopier012AdjacentBoundary(t *testing.T) {
	// Exercise a distinct adjacent boundary of the same public contract.
	to:=reflect.ValueOf(new(int)).Elem();from:=reflect.ValueOf("x");pair:=converterPair{SrcType:from.Type(),DstType:to.Type()};ok,_:=lookupAndCopyWithConverter(to,from,map[converterPair]TypeConverter{pair:{Fn:func(interface{})(interface{},error){return 3,nil}}});if !ok||to.Int()!=3{t.Fatal("converter not used")}
}
