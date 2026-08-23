package copier
import("reflect";"testing")
var _=reflect.TypeOf
func TestGoletaCopier012(t *testing.T){to:=reflect.ValueOf(new(string)).Elem();from:=reflect.ValueOf(7);pair:=converterPair{SrcType:from.Type(),DstType:to.Type()};ok,e:=lookupAndCopyWithConverter(to,from,map[converterPair]TypeConverter{pair:{Fn:func(interface{})(interface{},error){return "seven",nil}}});if e!=nil||!ok||to.String()!="seven"{t.Fatalf("ok=%v to=%q err=%v",ok,to.String(),e)}}
