package main

import (
	"fmt"
	"reflect"
)

// 使用反射来遍历结构体的字段 调用结构体的方法
// 并获取结构体标签的值
// 怎么拿到json的标签的

// func (v Value) Call(in []Value) []Value
// Call方法使用输入的参数in调用v持有的函数。
// 例如，如果len(in) == 3，v.Call(in)代表调用v(in[0], in[1], in[2])（其中Value值表示其持有值）。
// 如果v的Kind不是Func会panic。它返回函数所有输出结果的Value封装的切片。
// 和go代码一样，每一个输入实参的持有值都必须可以直接赋值给函数对应输入参数的类型。
// 如果v持有值是可变参数函数，Call方法会自行创建一个代表可变参数的切片，将对应可变参数的值都拷贝到里面。

// func (v Value) Method(i int) Value
// 默认按方法名排序对应i的值 i从0开始
// 传入参数和返回参数都是 []reflect.Value

//type Type interface {
//	// Kind返回该接口的具体分类
//	Kind() Kind
//	// Name返回该类型在自身包内的类型名，如果是未命名类型会返回""
//	Name() string
//	// PkgPath返回类型的包路径，即明确指定包的import路径，如"encoding/base64"
//	// 如果类型为内建类型(string, error)或未命名类型(*T, struct{}, []int)，会返回""
//	PkgPath() string
//	// 返回类型的字符串表示。该字符串可能会使用短包名（如用base64代替"encoding/base64"）
//	// 也不保证每个类型的字符串表示不同。如果要比较两个类型是否相等，请直接用Type类型比较。
//	String() string
//	// 返回要保存一个该类型的值需要多少字节；类似unsafe.Sizeof
//	Size() uintptr
//	// 返回当从内存中申请一个该类型值时，会对齐的字节数
//	Align() int
//	// 返回当该类型作为结构体的字段时，会对齐的字节数
//	FieldAlign() int
//	// 如果该类型实现了u代表的接口，会返回真
//	Implements(u Type) bool
//	// 如果该类型的值可以直接赋值给u代表的类型，返回真
//	AssignableTo(u Type) bool
//	// 如该类型的值可以转换为u代表的类型，返回真
//	ConvertibleTo(u Type) bool
//	// 返回该类型的字位数。如果该类型的Kind不是Int、Uint、Float或Complex，会panic
//	Bits() int
//	// 返回array类型的长度，如非数组类型将panic
//	Len() int
//	// 返回该类型的元素类型，如果该类型的Kind不是Array、Chan、Map、Ptr或Slice，会panic
//	Elem() Type
//	// 返回map类型的键的类型。如非映射类型将panic
//	Key() Type
//	// 返回一个channel类型的方向，如非通道类型将会panic
//	ChanDir() ChanDir
//	// 返回struct类型的字段数（匿名字段算作一个字段），如非结构体类型将panic
//	NumField() int
//	// 返回struct类型的第i个字段的类型，如非结构体或者i不在[0, NumField())内将会panic
//	Field(i int) StructField
//	// 返回索引序列指定的嵌套字段的类型，
//	// 等价于用索引中每个值链式调用本方法，如非结构体将会panic
//	FieldByIndex(index []int) StructField
//	// 返回该类型名为name的字段（会查找匿名字段及其子字段），
//	// 布尔值说明是否找到，如非结构体将panic
//	FieldByName(name string) (StructField, bool)
//	// 返回该类型第一个字段名满足函数match的字段，布尔值说明是否找到，如非结构体将会panic
//	FieldByNameFunc(match func(string) bool) (StructField, bool)
//	// 如果函数类型的最后一个输入参数是"..."形式的参数，IsVariadic返回真
//	// 如果这样，t.In(t.NumIn() - 1)返回参数的隐式的实际类型（声明类型的切片）
//	// 如非函数类型将panic
//	IsVariadic() bool
//	// 返回func类型的参数个数，如果不是函数，将会panic
//	NumIn() int
//	// 返回func类型的第i个参数的类型，如非函数或者i不在[0, NumIn())内将会panic
//	In(i int) Type
//	// 返回func类型的返回值个数，如果不是函数，将会panic
//	NumOut() int
//	// 返回func类型的第i个返回值的类型，如非函数或者i不在[0, NumOut())内将会panic
//	Out(i int) Type
//	// 返回该类型的方法集中方法的数目
//	// 匿名字段的方法会被计算；主体类型的方法会屏蔽匿名字段的同名方法；
//	// 匿名字段导致的歧义方法会滤除
//	NumMethod() int
//	// 返回该类型方法集中的第i个方法，i不在[0, NumMethod())范围内时，将导致panic
//	// 对非接口类型T或*T，返回值的Type字段和Func字段描述方法的未绑定函数状态
//	// 对接口类型，返回值的Type字段描述方法的签名，Func字段为nil
//	Method(int) Method
//	// 根据方法名返回该类型方法集中的方法，使用一个布尔值说明是否发现该方法
//	// 对非接口类型T或*T，返回值的Type字段和Func字段描述方法的未绑定函数状态
//	// 对接口类型，返回值的Type字段描述方法的签名，Func字段为nil
//	MethodByName(string) (Method, bool)
//	// 内含隐藏或非导出方法
//}

func main() {
	most := Monster75{
		Name:  "kanna",
		Age:   20,
		Score: 91,
		Sex:   "male",
	}
	testReflect(most) // 相当于找人执行了它的方法
}

type Monster75 struct {
	// func (tag StructTag) Get(key string) string
	Name  string `json:"name"` // name就是这个key
	Age   int    `json:"age"`
	Score float32
	Sex   string
}

func (s Monster75) Print() {
	fmt.Println("----------")
	fmt.Println(s)
	fmt.Println("----------")
}

func (s Monster75) GetSum(a, b int) int {
	return a + b
}

func (s Monster75) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func testReflect(in interface{}) {
	// 获取到reflect.Type 类型
	rTpy := reflect.TypeOf(in)
	fmt.Println("rTyp =", rTpy)
	// 获取到reflect.Value 类型
	rVal := reflect.ValueOf(in)
	fmt.Println("rVal =", rVal)
	// 获取到类别
	kd := rVal.Kind()
	fmt.Println("类别 =", kd)
	// 如果传入的不是结构体
	if kd != reflect.Struct {
		// 可以在这里对传入类型进行判断 这里是比较的常量数字
		fmt.Println("这不是一个结构体")
		return
	}
	// 获取该结构体有多少个字段

	// Field(i int) StructField
	// 返回索引序列指定的嵌套字段的类型，
	// 等价于用索引中每个值链式调用本方法，如非结构体将会panic

	//type StructField struct {
	//	// Name是字段的名字。PkgPath是非导出字段的包路径，对导出字段该字段为""。
	//	// 参见http://golang.org/ref/spec#Uniqueness_of_identifiers
	//	Name    string
	//	PkgPath string
	//	Type      Type      // 字段的类型
	//	Tag       StructTag // 字段的标签
	//	Offset    uintptr   // 字段在结构体中的字节偏移量
	//	Index     []int     // 用于Type.FieldByIndex时的索引切片
	//	Anonymous bool      // 是否匿名字段
	//}

	// func (tag StructTag) Get(key string) string
	// Get方法返回标签字符串中键key对应的值。如果标签中没有该键，会返回""。
	// 如果标签不符合标准格式，Get的返回值是不确定的。

	// 获取这个结构体有多少个字段
	numField := rVal.NumField()
	fmt.Println("字段数 = ", numField)
	for i := 0; i < numField; i++ {
		fmt.Printf("字段%d的值是%v\n", i, rVal.Field(i)) // 这是值
		tagVal := rTpy.Field(i).Tag.Get("json")
		if tagVal != "" { // 如果该字段有"json"标签才显示
			fmt.Printf("字段%d的标签是%v\n", i, tagVal) // 这才是标签 俩不一样
		} else {
			fmt.Printf("字段%d没有标签\n", i)
		}
	}

	// 获取这个结构体有多少个方法
	numFunc := rTpy.NumMethod() // 如果方法开头为小写则不算
	fmt.Println("方法数 = ", numFunc)

	// 获取到第二个方法 并调用
	// 注意这里的方法顺序是按照方法第一个字母的ASCII排序
	rVal.Method(1).Call(nil)

	// func (v Value) Call(in []Value) []Value
	var params []reflect.Value                   // 先声明这样的一个切片 调用Call需要的
	params = append(params, reflect.ValueOf(10)) // 这里是把一个常量转换为了reflec.Value类型
	params = append(params, reflect.ValueOf(40))
	res := rVal.Method(0).Call(params) // 这还是一个切片 因为Call接受和返回的都是切片
	fmt.Println("res =", res[0].Int()) // 这里是Value类型 不能Int(resp[0])强转
	// 上面要是类型不确定 那么还需要使用断言
	fmt.Println("res.len =", len(res)) // 那么这里返回的切片的长度也是1
}
