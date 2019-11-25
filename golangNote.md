# golang

## 环境变量配置

%GOROOT% = GOSDK安装目录根目录

%PATH% = %PATH%;%GOROOT%\bin

%GOPATH% = 源代码所在目录

## 编译 执行go程序

go build [file].go      --编译

go build -o [newFileName].exe [fileName].go   --编译并指定新文件名

go run [file].exe    --编译执行程序

## 变量定义

var n2 int = 100
var n1 = 300
n3 := 100

## 基本数据类型

### int

    int8 int16 int32 int64
    uint8 uint16 uint32 uint64

### float

    float32 float64(默认)开发推荐使用

### 字符类型char

    golang中没有专门的字符类型，如果要存储单个字符（字母），一般使用byte来保存
    如果保存的字符在ASCII表 可以直接保存到byte
    如果大于255，考虑使用int保存
    如何需要按照字符方式输出，需要格式化输出，即 fmt.Printf("%c",c1)

### bool布尔类型

    bool类型占1个字节

### 字符串类型 string

    字符串一旦赋值了，字符串就不能更改。在go中字符串是不可变的
    两个表现形式：
    "" 双引号
    `` 反引号
    字符串可以进行拼接 "+"    "+="      可以换行拼接   多行拼接+要保留在上行

## 基本数据类型默认值

    int     0
    float   0
    string  ""
    bool    false

## 基本数据类型转换

    golang中数据类型不能自动转换
    var i int32 = 100
    var b int8 = int8(i)    --类型转换方式

## 基本数据类型和string的转换

    fmt.Sprintf     --Sprintf根据format参数生成格式化的字符串并返回
    // 基本数据类型转string
    var num1 int = 99
    var num2 float64 = 23.456
    var b bool = true
    var myChar byte = 'h'
    var str string // 空的

    // 方式1 fmt.Sprintf方法
    str = fmt.Sprintf("%d", num1)
    fmt.Printf("str type %T str=%q\n", str, str)

    str = fmt.Sprintf("%f", num2)
    fmt.Printf("str type %T str=%q\n", str, str)

    str = fmt.Sprintf("%t", b)
    fmt.Printf("str type %T str=%q\n", str, str)

    str = fmt.Sprintf("%c", myChar)
    fmt.Printf("str type %T str=%q\n", str, str)

    // 方式2 使用strconv包的函数
    str = strconv.FormatInt(int64(num1), 10)
    fmt.Printf("str type %T str=%q\n", str, str)

    // strconv.FormatFloat(num2, 'f', 10 , 64)
    // 说明： 'f'格式  10：表示小数位保留10位   64表示这个小数是64位的
    str = strconv.FormatFloat(num2, 'f', 10 , 64)
    fmt.Printf("str type %T str=%q\n", str, str)

    str = strconv.FormatBool(b)
    fmt.Printf("str type %T str=%q\n", str, str)

    // 使用strconv包中有一个函数Itoa
    var num5 int64 = 4567
    str = strconv.Itoa(int(num5))
    fmt.Printf("str type %T str=%q\n", str, str)

## string转基本数据类型

    string类型转成基本数据类型时，如何不能转换成目标类型，golang直接将其转成0，false 类型默认值

    // string转基本数据类型
    var str1 string = "true"
    var b3 bool
    // 1.strconv.ParseBool(str) 函数会返回两个值 （vaule bool, err error）
    // 2.只想获取到value bool，不需要获取err error，所以使用 _来忽略
    b3 , _ = strconv.ParseBool(str1)
    fmt.Printf("b3 type %T b3=%v\n", b3, b3)

    var str2 string = "1234590"
    var i3 int64
    i3 , _ = strconv.ParseInt(str2, 10, 64)
    fmt.Printf("i3 type %T i3=%v\n", i3, i3)

    var str3 string = "123.456"
    var f4 float64
    f4 , _ = strconv.ParseFloat(str3, 64)
    fmt.Printf("f4 type %T f4=%v\n", f4, f4)

## fmt

    fmt.Println("XXXXX")
    fmt.Printf("")     --fmt.Printf("n的数据类型%T 占用字节数是%a", n, unsafe.Sizeof(n))  n的数据类型int 占用字节数是8      %v 表示按原值输出

    引入包使用
    import "fmt"
    import (
        "fmt"
    )

    import (
        _ "fmt" // 如果没有使用到一个包，但是又不想删除，前面加一个_, 表示忽略
    )

## 指针

    1.基本数据类型，变量存的就是值，也叫值类型
    2.获取变量的地址，用&，比如：var num int，获取num的地址就是 &num
    3.指针类型，指针变量存的是一个地址，这个地址指向的空间存的才是值，比如：var ptr *int = &num
    4.获取指针类型所指向的值，使用：*，比如 var ptr *int，使用*ptr获取ptr指向的值
