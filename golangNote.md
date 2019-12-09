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

### 基本数据类型默认值

    int     0
    float   0
    string  ""
    bool    false

### 基本数据类型转换

    golang中数据类型不能自动转换
    var i int32 = 100
    var b int8 = int8(i)    --类型转换方式

### 基本数据类型和string的转换

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

### string转基本数据类型

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

## 值类型和引用类型

值类型：int float bool string,数组和结构体struct
值类型，变量直接存储值，内存通常在栈中分配

引用类型：指针 slice切片 map 管道chan interface
引用类型：变量存储的是一个地址，这个地址对应的空间才真正存储数据（值），内存通常在堆上分配，
当没有任何变量引用这个地址时，该地址对应的数据空间就变成一个垃圾，由GC来回收

## 标识符命名规范

如果变量名、函数名、常量名首字母大写，则可以被其他的包访问；如果首字母小写，则只能在本包中使用
首字母大写（public） 首字母小写（private）
    // 定义公有
    var FileName string = "代办"
    -- 外部引用
    import (
        "go_code/project01/referVar/refer"
    )
    fmt.Println(refer.FileName)

## 算数运算符

运算符  /           整数之间做除法，只保留整数部分而舍弃小数部分。例如: x:=19/5    --3
运算符  %           等价 a % b = a - a / b * b
运算符 ++ or --     只能独立使用  只有后++，后--

两数交换案例
a:= 7
b:= 2
temp:= a
a:=b
b:=temp

a:=10
b:=20
a = a + b
b = a - b   // b = a + b - b  => b = a
a = a - b   // a = a + b - a  => a = b

## 键盘输入语句

导入fmt包
调用fmt包的 fmt.Scanln() 或者 fmt.Scanf()
var name string
fmt.Scanln(&name)
fmt.Scanf("%s", &name)

## 进制

var i int = 5
fmt.Printf("%b", i)

## switch

匹配项后面不用带break,case后的表达式可以有多个，使用，间隔
方式1：
var score int = 30
switch {
    case score > 90:
        ...
    case score >= 70 && score <= 90:
        ...
    default :
        ...
}

方式2：
switch score:= 90; {
    case score > 90:
        ...
    case score >= 70 && score <= 90:
        ...
    default :
        ...
}

switch穿透 fallthrough

## for

方式1：
for i:=1; i<=10; i++ {
    fmt.Println(i)
}
方式2：
j := 1
for j <= 10 {
    fmt.Println(j)
    j++
}
方式3：
for {
}
==
for ;; {

}
方式4：for-range
var str string = "abc~ok"
for index, val := range str {
    fmt.Printf("index=%d, val=%c \n", index, val)
}

字符串遍历
var str string = "abc~ok"
for i:= 0; i < len(str); i++ {
    fmt.Printf("%c \n", str[i])
}

如何字符串含有中文，传统遍历会出现乱码，传统遍历是按照字节遍历，而一个汉字在utf8编码是对应3个字节
var str1 string = "hello,world!上海"
str2 := []rune(str1)
for i:= 0; i < len(str2); i++ {
        fmt.Printf("%c \n", str2[i])
}

## 获取随机数

为了生成一个随机数，还需要给rand设置一个种子
time.Now().UnixNano() : 返回一个从1970.01.01 的0时0分0秒的豪秒数
rand.Seed(time.Now().Unix())
rand.Intn(100)      获取[0, 100)的随机数

## 函数

函数            分为自定义函数和系统函数
func 函数名（形参列表）（返回值列表）{
    执行语句...
    return 返回值列表
}
1)如何返回多个值时，在接收时，希望忽略某个返回值，则使用 _ 符号表示占位忽略
2)如果返回值只有一个，（返回值类型列表）可以不写()
func getSumAndSub(n1 int, n2 int) (int, int) {
    sum := n1 + n2
    sub := n1 - n2
    return sum, sub
}
func main() {
    res1,res2 := getSumAndSub(2, 1)
    fmt.Println("sum=",res1,", sub=", res2)
    _, res3 := getSumAndSub(5, 1)
    fmt.Println(" sub=", res3)
}

## 包 package

打包基本语法
package 包名

引入包的基本语法
import "包的路径"
在import包时，路径从$GOPATH$的src下开始，不用带src，编译器会自动从src下开始引入
如果包名过长，可以取个别名
import (
    "fmt"
    util "go_code/project01/func/utils"
)
相同包下 不能有相同的函数名
编译项目
go build -o bin/my.exe go_code/project01/func/main

## 函数递归调用

一个函数在函数体内又调用了本身，称为递归调用
但是递归一定要有退出条件，否则无限循环调用
斐波那契数  递归算法

## 函数调用注意点

1.基本数据类型和数组默认都是值传递
2.如果希望函数内的变量能修改函数外的变量，可以传入地址&
3.go语言不支持函数重载
4.在go中，函数也是一种数据类型，可以赋值给一个变量，则该变量就是一个函数类型变量了。通过该变量可以对函数调用
5.函数既然是一种数据类型，因此在go中，函数可以作为形参，并且调用
    func getSum(n1 int,n2 int) int {
        return n1 + n2
    }

    func myFum(funvar func(int,int) int, n1 int ,n2 int) int {
        return funvar(n1, n2)
    }

    func main() {
        res := myFum(getSum, 5, 10)
        fmt.Println("main res=" , res)
    }
6.为了简化数据类型定义，go支持自定义数据类型
    type mySum func(int,int) int ,mySum等价于一个函数类型func(int,int) int
7.支持对函数返回值命名
    func getSumAndSub(n1 int, n2 int) (sum int, sub int) {
        sum = n1 + n2
        sub = n1 - n2
        return
    }
8.使用 _ 标识符，忽略返回值
9.go支持可变参数

## init函数

每一个源文件都可以包含一个init函数，该函数会在main函数执行前，被go运行框架调用，也就是说init会在main函数前被调用
如果一个文件同时包含全局变量定义,init函数和main函数，那么执行的流程 全局变量定义->init函数->main函数
main函数 引用了其他包文件 函数执行顺序
引入包变量定义-> 引入包init-> main变量定义-> main的init函数-> main函数

## 匿名函数

匿名函数使用方式1
在定义匿名函数时就直接调用，这种方式只能调用一次
    res := func (n1 int, n2 int) int {
        return n1 + n2
    }(10, 20)
    fmt.Println("res=",res)

匿名函数使用方式2
将匿名函数赋给一个变量（函数变量），在通过该变量来调用匿名函数
    a := func (n1 int, n2 int) int {
        return n1 - n2
    }
    res2 := a(20, 10)
    fmt.Println("res2=",res2)

全局匿名函数
    var (
        // fun1就是一个全局匿名函数
        Fun1 = func(n1 int, n2 int) int {
            return n1 * n2
        }
    )

## 闭包

闭包就是一个函数和与其相关的引用环境组合的一个整体（实体）
    func AddUpper() func(int) int {
        var n int = 10
        return func (x int) int {
            n = n + x
            return n
        }
    }
    func main() {
        f := AddUpper()
        fmt.Println(f(1)) // 11
        fmt.Println(f(2)) // 13
        fmt.Println(f(3)) // 16
    }
闭包说明
返回的是一个匿名函数，但是这个匿名函数引用到函数外的n，因此这个匿名函数就和这个n形成一个整体，构成闭包。
闭包是类，函数时操作，n是字段。函数和它使用的n构成闭包

func makeSuffix(suffix string) func(string) string{
    return func(fileName string) string {
        if strings.HasSuffix(fileName, suffix) {
            return fileName
        } else {
            return fileName + suffix
        }
    }
}
代码说明
1.返回的函数和makeSuffix(suffix string)的suffix变量和返回的函数组合成一个闭包，因为引用的函数引用到suffix这个变量
2.闭包的好处，如果使用传统的方法，也可以实现这个功能，但是传统的方法需要每次都传入后缀名，比如.jpg，而闭包因为可以保留上次引用的某个值，所以我们传入一次就可以反复使用

## strings.HasSuffix

该函数可以判断某个字符串是否有指定的后缀
strings.HasSuffix(s string, suffix string) bool

## 函数的defer

在函数中，经常需要创建资源（比如：数据库连接、文件句柄、锁等），为了在函数完毕后，及时的释放资源，提供defer（延时机制）
    // 当执行到defer时，暂时不执行，会将defer后面的语句压入到独立的栈（defer栈中）
    // 当函数执行完毕后，再从defer栈，按照先入后出的方式出栈，执行
    defer fmt.Println("n1",n1)
    defer fmt.Println("n2",n2)
所以先 打印n2 , 再打印n1
在defer压入栈时，同时也会将相关的值拷贝同时入栈

defer最佳实践：
file = openFile（文件名）
defer file.close()

connect = openDataBase()
defer connect.close()

## 作用域

全局变量错误
name := "tom"  // × 报错 在全局定义中
原因是 name := "tom"  等价于 var name string   name = "tom"
赋值语句不能在函数体外

## 字符串常用的系统函数

1.len       统计字符串长度,按字节len(str)
len("hello")         5
len("hello北")       8
golang的编码统一为utf-8（ascii的字符（字母和数字）占一个字节，汉字占用3个字节）

2.字符串遍历，同时处理有中文的问题 r := []rune(str)

3.字符串转整数 n,err := strconv.Atoi("12")
    n, err := strconv.Atoi("123")
    if err != nil {
        fmt.Println("错误:", err)
    } else {
        fmt.Println("n=", n)
    }

4.整数转字符串 str := strconv.Itoa("12")
str = strconv.Itoa(123)
fmt.Printf("str=%v,str=%T", str, str)

5.字符串转 []byte:  var bytes = []byte("hello go")

6.[]byte转字符串： str = string([]byte{97,98,99})

7.10进制转2,8,16进制: str = strconv.FormatInt(123,2) // 2->8,16

8.查找子串是否在指定的字符串中： strings.Contains("seafood", "foo") // true

9.统计一个字符串有几个指定的子串: strings.Count("seafood sea", sea) // 2

10.不区分大小写的字符串比较(==是区分字母大小写的): strings.EqualFold("abc","ABC") // True

11.返回子串在字符串第一次出现的index值，如果没有，返回-1：strings.Index("NLT_abc","abc")// 4

12.返回子串在字符串最后一次出现的index值如果没有，返回-1：strings.LastIndex("go golang","go")// 3

13.将指定的子串替换成另外一个子串： strings.Replace("go go hello", "go", "go语言", n)
n可以指定希望替换几个，如果n=-1表示全部替换

14.按照指定的某个字符，为分割标识，将一个字符串拆分成字符串数组
strings.Split("hello,world,ok",",")
strArr := strings.Split("hello,world,ok",",")
for i:= 0; i < len(strArr); i++ {
    fmt.Printf("str[%v]=%v\n", i, strArr[i])
}

15.将字符串的字母进行大小写转换
strings.ToLower("GO")   //go
strings.ToUpper("go")   // GO

16.将字符串左右两边的空格去掉
strings.TrimSpace("   short boy peter   ")  // "short boy peter"

17.将字符串左右两边指定的字符串去掉
strings.Trim("! hello! "," !")          // 去掉左右两边的"!"和" "

18.将字符串左边指定的字符串去掉
strings.TrimLeft("! hello! "," !")

19.将字符串右边指定的字符串去掉
strings.TrimRighte("! hello! "," !")  

20.判断字符串是否以指定字符串开头：strings.HasPerfix("ftp://192.168.10.1","ftp") // true

21.判断字符串是否以指定字符串结束: strings.HasSuffix("ftp://192.168.10.1","ftp") // false

## 时间和日期相关函数

需要导入time包
time.Time 类型，用于表示时间
1.获取当前时间
    now := time.Now()
2.通过now可以获取 年月日 时分秒
now.Year()
now.Month()  // May => 5    int(now.Month())
now.Day()
now.Hour()
now.Minute()
now.Second()
3.格式化日期时间
(1)格式化的第一种方式
fmt.Printf("当前年月日 时分秒 %d-%d-%d %d:%d:%d\n",now.Year(), now.Month(),
now.Day(),now.Hour(),now.Minute(),now.Second())

dateStr := fmt.Sprintf("当前年月日 时分秒 %d-%d-%d %d:%d:%d\n",now.Year(), now.Month(),
now.Day(),now.Hour(),now.Minute(),now.Second())

(2)格式化的第二种方式
fmt.Printf(now.Format("2006-01-02 15:04:05"))
fmt.Println("")
fmt.Printf(now.Format("2006-01-02"))
fmt.Println("")
fmt.Printf(now.Format("15:04:05"))