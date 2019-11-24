# golang
## 基本数据类型
### int
    int8 int16 int32 int64
    uint8 uint16 uint32 uint64

### float
    float32 float64(默认)开发推荐使用

### char
    golang中没有专门的字符类型，如果要存储单个字符（字母），一般使用byte来保存
    
boolean


## fmt
    fmt.Println("XXXXX")
    fmt.Printf("")          --fmt.Printf("n的数据类型%T 占用字节数是%a", n, unsafe.Sizeof(n))           n的数据类型int 占用字节数是8