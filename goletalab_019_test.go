package copier

import ("fmt";"reflect";"strings";"testing")

var _=fmt.Sprint
var _=strings.Contains
var _=reflect.DeepEqual
type goletaConvSource019 struct{V string}
type goletaConvDest019 struct{V int}
func goletaConvOption019()Option{return Option{Converters:[]TypeConverter{{SrcType:string(""),DstType:int(0),Fn:func(interface{})(interface{},error){return nil,fmt.Errorf("rejected input")}}}}}

func TestGoletaCopier019(t *testing.T){src:=goletaConvSource019{V:"bad"};var dst goletaConvDest019;err:=CopyWithOption(&dst,src,goletaConvOption019());if err==nil||!strings.Contains(err.Error(),"rejected"){t.Fatalf("err=%v",err)}}

func TestGoletaCopier019Boundary(t *testing.T) {
 src:=goletaConvSource019{V:"other"}
 var dst goletaConvDest019
 err:=CopyWithOption(&dst,src,goletaConvOption019())
 if err==nil{t.Fatal("converter error lost")}
}
