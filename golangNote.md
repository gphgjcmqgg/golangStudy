# golang

## 环境变量配置

%GOROOT% = GOSDK安装目录根目录

%PATH% = %PATH%;%GOROOT%\bin

%GOPATH% = 源代码所在目录

## 编译 执行go程序

go build [file].go      --编译

go build -o [newFileName].exe [fileName].go   --编译并指定新文件名

go run [file].go    --编译执行程序

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

## 时间常量

时间单位换算
const {
    Nanosecond Duration = 1 //纳秒
    Microsecond = 1000 * Nanosecond //微秒
    Millisecond = 1000 * Microsecond //毫秒
    Second = 1000 * Millisecond    //秒
    Minute = 60 * Second    //分钟
    Hour = 60 * Minute    //小时
}

时间常量案例
i:=0
for {
    fmt.Println(i)
    i++
    time.Sleep(time.Millisecond * 100) // 休眠0.1秒 100毫秒
    if i == 100 {
        break
    }
}

time.Sleep(d Duration)          // 休眠

time的Unix和UnixNano的方法      获取随机数
time.Now().Unix()    time.Now().UnixNano()

## go内置函数(buildin)

1.len   返回长度 如string、array、slice、map、channel
2.new   用来分配内存，主要用来分配值类型，比如int、float32、struct 返回的是指针
    num2 := new(int)
    // num2的类型%T == *int
    // num2的值是== 地址 0xc04200e0e0	（这个地址系统分配）
    // num2的地址 == 地址 0xc042004030	（这个地址系统分配）
    // num2指向的值 == 0
    fmt.Printf("num2的类型%T, num2的值是%v, num2的地址%v, num2指向的值%v\n", num2 , num2 , &num2, *num2)
    *num2 = 100
    fmt.Printf("num2的类型%T, num2的值是%v, num2的地址%v, num2指向的值%v\n", num2 , num2 , &num2, *num2)

## go错误处理机制

go引入的错误处理方式：defer、panic、recover，
go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理

defer + recover来处理异常
defer func () {
    err := recover() // recover()内置函数、可以捕获到异常
    if err != nil { // 说明捕获到错误
        fmt.Println("err=",err)
    }
}()

## go自定义错误处理

errors.New("错误说明"), 会返回一个error类型的值，表示一个错误
panic内置函数，接收一个interface{}类型的值作为参数。可以接收recover类型变量,输出错误信息，并退出程序。

## 数组

数组可以存放多个同一类型数据。数组也是一种数据类型，在go中，数组是值类型

数组定义
var 数组名 [数组大小]数据类型
var a [6]float64
赋初值
a[0] = 1.0

数组的地址可以通过数组名来获取&a
数组的元素的地址、就是数组元素第一个元素的地址

    // 数组几种初始化的方式
    var numArr01 [3]int = [3]int{1,2,3}
    fmt.Println("numArr01=", numArr01)

    var numArr02 = [3]int{5,6,7}
    fmt.Println("numArr02=", numArr02)

    var numArr03 = [...]int{7,8,9}
    fmt.Println("numArr03=", numArr03)

    var numArr04 = [...]int{1:700,0:800,2:900}
    fmt.Println("numArr04=", numArr04)

    numArr05 := [3]int{15,16,17}
    fmt.Println("numArr05=", numArr05)

数组遍历

1. for循环遍历
2. for-range遍历
for i,v := range numArr03 {
    fmt.Printf("index=%d, value=%v \n", i, v)
}

数组的注意细节
1.数组是多个相同类型数据的组合，一个数组一旦声明，其长度固定，不能动态变化
2.数组创建后，如果没有赋值，有默认值
3.数组下标从0开始
4.go的数组属值类型，在默认情况下是值传递，因此会进行值拷贝，数组间不会相互影响
5.如想在其他函数中，去修改原来的数组，可以使用引用传递（指针方式）
func test(arr *[3]int) {
    (*arr)[0] = 55
}
6.长度是数组类型的一部分，在传递函数参数时需要考虑数组的长度。

## 切片 slice

为什么需要切片
1.需要保存个数不确定的数据
2.切片是数组的一个引用，切片是引用类型，在进行传递时，遵守引用传递的机制
3.切片的使用和数组类似，遍历切片、访问切片的元素和求切片长度len都一样
4.切片的长度是可以变化的
5.切片定义的基本语法
var 变量名 []数据类型
var a []int

slice案例
var intArr [5]int = [5]int{1,22,33,66,99}
// slice为切片
// intArr[1:3]表示slice引用到intArr这个数组的第二个元素到第三个元素
// 引用intArr数组的起始下标为1，最后的下标为3（但是不包含3）
slice := intArr[1:3]
fmt.Println(slice)

slice 分析   保存了 一个引用首地址，一个切片长度，一个切片容量
type slice struct {
    ptr *[2]int
    len int
    cap int
}

切片使用的三种方式
1.定义一个切片，然后让切片去引用一个已经创建的数组
var intArr [5]int = [5]int{1,22,33,66,99}
slice := intArr[1:3]

2.通过make来创建切片
基本语法：
var 切片名 []type = make([]type,len,[cap])
参数说明: type：就是数据类型  len：大小 cap：指定切片容量，可选
make方式创建切片可以指定切片的大小和容量
如何没有给切片的各个元素赋值，那么就会使用默认值
通过make方式创建的切片对应的数组是由make底层维护，对外不可见，即只能通过slice去访问

3.定义一个切片，直接就指定具体数组，使用原理类似make的方式
var slice []int = []int{1, 3, 5}

切片注意细节
1.切片初始化时 var slice = arr[startIndex:endIndex]
说明:从arr数组下标为startIndex，取到下标为endIndex的元素（不含arr[endIndex]）

2.切片初始化可以简写
var slice = arr[0:end]            ==>var slice = arr[:end]
var slice = arr[start:len(arr)]   ==>var slice = arr[start:]
var slice = arr[0:len(arr)]       ==>var slice = arr[:]

3.切片可以继续切片

4.切片可以用append内置函数，进行动态追加扩容
var slice []int = []int{100,200,300}
slice = append(slice, 400,500,600)
// 切片追加切片
slice1 = append(slice1, slice1...)

5.切片的拷贝操作
切片使用copy内置函数完成拷贝
copy(slice1, slice2)
(1)copy函数的参数数据类型是切片
(2)copy的两个切片互相独立

## string和slice的关系

1.string底层是一个byte数组，所以可以进行切片处理
str := "hello@World"
slice = str[6:] // World

2.string是不可变的，不能通过str[0] = 'z'方式来修改字符串

3.如果需要修改字符串，可以先将string->[]byte /或者 []rune ->修改 ->重新转成string
转成[]byte后，可以处理英文数字，但是不能处理中文
原因是[]byte字节来处理，而一个汉字，是三个字节，因此会出现乱码
解决方法 是将 string转成 []rune即可，因为 []rune是按字符处理的，兼容汉字
str2 := "hello@ma"
sliceStr :=  []byte(str2)
sliceStr[6] = 'b'
sliceStr[7] = 'a'
str2 = string(sliceStr)

str2 = "hello@北京"
sliceChineseStr :=  []rune(str2)
sliceStr[6] = '东'

## 排序和查找

1.排序的分类
1)内部排序
    将需要处理的所有数据都加载到内部存储器中进行排序
    包括（交换式排序法、选择式排序法和插入式排序法）
2)外部排序
    数据量过大，无法全部加载到内存中，需要借助外部存储进行排序。
    包括（合并排序法和直接合并排序法）

交换式排序法：是属于内部排序法，是运用数据值比较后，依判断规则对数据位置进行交换，以达到排序的目的
1.交换式排序法又分为两种：
    1)冒泡排序法(Bubble sort)
    2)快速排序法(Quick sort)

冒泡排序的规则
1.一共会经过arr.length -1次的轮数比较，每一轮将会确定一个数的位置
2.每一轮的比较次数在逐渐的减少
3.当发现前面的一个数比后面的一个数大的时候，就可以进行交换

常用的查找有两种
1)顺序查找
2)二分查找（该数组有序）会使用递归

二分查找思路
1.arr是一个有序数组，并且是从小到大排序
2.先找到中间的下标 middle = (leftIndex + rightIndex) / 2, 然后让中间下标的值和findVal进行比较
    2.1如果arr[middle] > findVal,就应该向 leftIndex --- (middle -1)范围查找
    2.2如果arr[middle] < findVal,就应该向 (middle +1) --- rightIndex 范围查找
    2.3如果arr[middle] == findVal, 就找到
    2.4上面的2.1、2.2、2.3的逻辑会递归执行
3.分析递归的退出条件(leftIndex > rightIndex)

## 二维数组

定义/声明二维数组
var arr [2][3]int
arr     保存了两个指针，每个指针对应一个一维数组

二维数组直接初始化
var 数组名 [2][3]int = [2][3]int{{1,2,3},{4,5,6}}
var 数组名 [2][3]int = [...][3]int{{1,2,3},{4,5,6}}
var 数组名 = [2][3]int{{1,2,3},{4,5,6}}
var 数组名 = [...][3]int{{1,2,3},{4,5,6}}

## map

map是key-vale数据结构，又称为字段或者关联数组。
基本语法
var map变量名 map[keytype]valuetype
var a map[string]string
var a map[string]int
var a map[int]string
var a map[string]map[string]string
注意：声明是不会分配内存的，初始化需要make，分配内存后才能赋值和使用

var a map[string]string
// 在使用map前，需要先make，make的作用就是给map分配数据空间
// map的key-value是无序的
a = make(map[string]string,10)
a["no1"] = "123"
a["no2"] = "456"

map使用的方式：
1.先声明，在make
    var a map[string]string
    a = make(map[string]string,10)
2.声明直接make
    var a = make(map[string]string)
3.声明直接赋值
var b map[string]string = map[string]string {
    "no1":"成都",
}

map的增删改查操作
1.map增加和更新
map["key"] = value  // 如何key还没有，就是增加，如果key存在就是修改

2.map删除
delete是一个内建函数，如果key存在，就删除该key-value,如果key不存在,不操作，但是也不会报错
    delete(map, key)
map一次删除所有key细节：
1.遍历所有key，逐一删除
2.直接make一个新的空间

3.map查找
val, findRes = heroes["no1"]
if findRes {
    // 找到val
} else {
    // 没有no1这个key
}

4.map遍历
map遍历使用for-range的结构遍历

5.map长度
使用len(map)

6.map切片
切片的数据类型如果是map，称为slice of map, map切片，这样使用则 map个数就可以动态变化
var slice []map[string]string
// map切片
var slice []map[string]string
slice = make([]map[string]string,2)
if slice[0] == nil {
    slice[0]= make(map[string]string,2)
    slice[0]["name"] = "jack"
    slice[0]["age"] = "34"
}

fmt.Println(slice)
// 动态增加map，使用append
monster := map[string]string {
    "name" : "zhangsan",
    "age" : "200",
}
slice = append(slice,monster)
fmt.Println(slice)

7.map排序

如何按照map的key的顺序进行排序输出
1.先将map的key放入到切片中
2.对切片进行排序
3.遍历切片，然后按照key来输出map的值
map1 := map[int]int {
    10:99,
    1:5,
    4:23,
    7:8,
}
// 排序
var keys []int
for k,_ := range map1 {
        keys = append(keys,k)
}
sort.Ints(keys)

for _ ,k := range keys {
        fmt.Printf("map[%v]=%v\t", k, map1[k])
}

map使用细节
1.map是引用类型，遵守引用类型传递的机制，在一个函数接受map，修改后，会直接修改原map
2.map的容量达到后，再想map增加元素，会自动扩容，不会发生panic，也就是说map能动态的增长键值对
3.map的value经常使用struct类型，更适合管理复杂的数据

## 面向对象

## 结构体

golang是基于struct来实现OOP特性，golang面向接口编程是非常重要的特性
继承：golang是通过匿名字段来实现

声明结构体
type 结构体名称 struct {
    field1 type
    field2 type
}
结构体是值类型，默认值拷贝

注意：
指针、slice和map的零值时nil，即还没有分配空间
使用切片、map前一定要先make，分配空间

创建结构体变量和访问结构体字段
1.直接声明
    var person Person
2.{}
    var person Person = Person{}
    var cat2 Cat = Cat{"Tom", 20, "Red", "旅游"}
3.&
    var person *Person = new(Person)
4.{}
    var person *Person = &Person{}
    var cat2 *Cat = &Cat{"Tom", 20, "Red", "旅游"}
说明：
1.方式3和方式4返回的是结构体指针
2.结构体指针访问字段的标准方式应该是：(*结构体指针).变量名
比如(*person).Name = "Tom"
3.但go做了一个简化，也支持结构体指针.字段名,
比如person.Name = "Tom"
go编译器底层对person.Name做了转化(*person).Name

结构体使用注意事项和细节
1)结构体的所有字段在内存中是连续的
2)结构体是用户单独定义的类型，和其它类型进行转换时需要有完全相同的字段(名称、个数和类型)
3)结构体进行type重新定义（相当于取别名）,Golang认为是新的数据类型，但是相互之间可以强转
4)struct的每个字段上，可以写上一个tag，该tag可以通过反射机制获取，常见的使用场景就是序列化和反序列化

客户端                                 服务器端
1.接受到一个json字符串                  1.假设返回一个Person变量值
2.进行处理                             2.Person序列化[json] ->字符串

转json用"encoding/json"
jsonStr,err := json.Marshal(person) ->  转成json字符串
例子：
type Person struct {
    Name string `json:"name"`           \\ `json:"name"`就是tag
    Age int `json:"age"`
}

var person = Person{"boy",50}
fmt.Println("person=", person)
jsonStr,err := json.Marshal(person)         \\ json.Marshal函数中使用反射
if err != nil {
    fmt.Println("err=", err)
} else {
    fmt.Println("jsonStr=", string(jsonStr))
}

## 方法

golang中的方法是作用在指定的数据类型上的（即：和指定的数据类型绑定），因此自定义类型，都可以有方法，而不仅仅是struct

方法的声明和调用
type A struct {
    Num int
}

func(a A) test(){
    fmt.Println(a.Name)
}

对上面语法的说明：
1.func(a A) test(){} 表示A结构体有一方法,方法名为test
2.(a A)体现test方法是和A类型绑定的
例子：
type Person struct{
    Name string
}

// 给A类型绑定一个方法
func (p Person) test() {
    fmt.Println("test()", p.Name)
}

func main() {
    var p Person
    p.Name = "tom"
    p.test()
}
上面代码总结：
1.test方法和Person类型绑定
2.test方法只能通过Person类型的变量来调用，而不能直接调用，也不能使用其他类型变量来调用
3.func(p Person) test(){}       p表示哪个Person变量调用，这个p就是它的副本，这点和函数传参非常相似
4.p这个形参名，由程序员指定，不是固定

方法的声明（定义）
func(receiver type) methodName(参数列表) (返回值列表) {
    方法体
    return 返回值
}

func(stu *Student) String() string {
    str := fmt.Sprintf("Name=[%v] Age=[%v]", stu.Name, stu.Age)
    return str
}

main() {
    stu := Student{
        Name: "Tom",
        Tom: 20,
    }
    fmt.Printf(stu)
}
1.如果一个变量实现了String()这个方法，那么fmt.Println默认会调用这个变量的String()进行输出

## 方法和函数区别

1)调用方式不一样
    函数的调用方式：    函数名(实参列表)
    方法的调用方式：    变量.方法名(实参列表)

2)对于普通函数，接受者为值类型时，不能将指针类型的数据直接传递，反之亦然
3)对于方法(如struct的方法)，接收者为值类型时，可以直接用指针类型的变量调用方法，反过来同样可以

总结：
1.不管调用形式如何，真正决定是值拷贝还是地址拷贝，看这个方法是和哪个类型绑定
2.如果和是值类型，比如（p Person），则是值拷贝，如果和指针类型，比如（p *Person），则是地址拷贝

## 面向对象编程三大特征

1.封装
1)将结构体、字段（属性）的首字母小写（类似private）
2)给结构体所在包提供一个工厂模式的函数，首字母大写，类似一个构造函数
3)提供一个首字母大写的Set方法，用于对属性判断并赋值
4)提供一个首字母大写的Get方法，用于获取属性的值

2.继承
在golang中，如果一个struct嵌套了另一个匿名结构体，那么这个结构体可以直接访问匿名结构体的字段和方法，从而实现了继承特性

嵌套匿名结构体的基本语法
type Goods struct {
    Name string
    Price int
}

type Book struct {
    Goods     // 这里就是嵌套匿名结构体Goods
    Writer string
}

继承的深入讨论
1.结构体可以使用嵌套匿名结构体所有的字段和方法，即：首字母大写或者小写的字段、方法，都可以使用
2.匿名结构体字段访问可以简化
如
    var book Book
    book.Goods.Name  // book.Name 简化
    总结：
    1)直接通过b访问字段或者方法时，其执行流程如下比如b.Name
    2)编译器会先看对应的类型有没有Name，如果有，直接调用B类型的Name字段
    3)如果没有就去看B中嵌入的匿名结构体A有没有声明Name字段，如果有就调用，如果没有继续查找，如果都找不到就报错
3.当结构体和匿名结构体有相同的字段或者方法时，编译器采用就近访问原则访问，如果希望访问匿名结构体的字段和方法，可以通过匿名结构体名来区分
4.结构体嵌入两个（或多个）匿名结构体，如两个匿名结构体有相同的字段和方法（同时结构体本身没有同名的字段和方法），在访问时，就必须明确指定匿名结构体名字，否则编译报错
5.如果一个struct嵌套了一个有名结构体，这种模式就是组合，如果是组合关系，那么在访问组合的结构体的字段或方法时，必须带上结构体的名字
6.嵌套匿名结构体后，也可以在创建结构体变量（实例）时，直接指定各个匿名结构体字段的值
tv := TV { Goods {"电视", 5000.00}, Brand{"海尔", "山东"},}

## 接口

基本介绍
interface类型可以定义一组方法，但是这些不需要实现。并且interface不能包含任何变量。到某个自定义类型要使用的时候，在根据具体情况把这些方法写出来。

基本语法
type 接口名 interface {
    method1(参数列表)返回值列表
    method2(参数列表)返回值列表
}
实现接口所有方法
func(t 自定义类型) method1(参数列表)返回值列表{
    // 方法实现
}
func(t 自定义类型) method2(参数列表)返回值列表{
    // 方法实现
}
小结说明
1)接口里的所有方法都没有方法体，即接口的方法都是没有实现的方法。接口体现了程序设计的多态和高内聚低耦合的思想
2)golang中的接口不需要显式的实现。只要一个变量，含有接口类型中的所有方法，那么这个变量就实现了这个接口。因此，golang中没有implement这样的关键字

接口的注意事项和细节
1)接口本身不能创建实例，但是可以指向一个实现了该接口的自定义类型的变量（实例）
2)接口中所有的方法都没有方法体，即都没有实现的方法
3)在golang中，一个自定义类型需要将某个接口的所有方法都实现，我们说这个自定义类型实现了该接口
4)一个自定义类型只有实现了某个接口，才能将该自定义类型的实例(变量)赋给接口类型
5)只要是自定义数据类型，就可以实现接口，不单单是结构体类型
6)一个自定义类型可以实现多个接口
7)golang接口中不能有任何变量
8)一个接口A可以继承多个别的接口b、c，这时如果要实现A接口，也必须将B,C接口的方法也全部实现
9)interface类型默认是一个指针（引用类型），如果没有对interface初始化就使用那么会输出nil
10)空接口interface{}没有任何方法，所以所有类型都实现了空接口

实现接口 vs 继承
实现接口是对继承机制的补充

接口和继承解决的问题不同
继承的价值主要在于：解决代码的复用性和可维护性
接口的价值主要在于：设计好各种规范（方法），让其他自定义类型去实现这些方法

继承是满足is-a的关系，接口是满足like-a的关系
接口在一定程度上实现代码解耦

面向对象编程多态
变量（实例）具有多种形态。面向对象的第三大特征，在go语言，多态特征是通过接口实现的。可以按照统一的接口来调用不同的实现。这时接口变量就呈现不同形态。

接口体现多态的两种方式
1.多态参数
2.多态数组
var usbArr = [3]Usb
usbArr[0] = Phone{}
usbArr[1] = Phone{}
usbArr[2] = Camera{}

类型断言
var a interface{}
var point Point = Point{1,2}
a = point
var b Point
// b = a // 报错
b = a.(Point)  // 类型断言
fmt.Println(b)
基本介绍
类型断言，由于接口是一般类型，不知道具体类型，如果要转成具体类型，就需要使用类型断言

在进行类型断言时，如果类型不匹配，就会报panic，因此进行类型断言时，要确保原来的空接口指向的就是断言的类型。
如何在进行断言时，带上检测机制，如果成功就ok，否则也不要报panic
y,ok := x.(float64)
if ok {
    fmt.Println("convert success")
} else {
    fmt.Println("convert fail")
}

