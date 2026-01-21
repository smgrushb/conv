# Conv

## 关于版权与协议

本项目是对 [coven](https://github.com/petersunbag/coven) 项目Fork后的深度修改版本，原项目使用 [MIT 协议](https://opensource.org/licenses/MIT) 发布，原作者为 [petersunbag](https://github.com/petersunbag)。

当前项目中，原始代码已被大幅重构或替换，新增部分以 [Apache License 2.0](https://www.apache.org/licenses/LICENSE-2.0) 许可发布。

具体详见 [LICENSE 文件](./LICENSE)。

向原作者的贡献表示诚挚的感谢

## 特别声明

1. 本文档由`claude-3.7-sonnet-thinking`阅读源码后生成，已经检查后调整部分内容。
2. 本工具预期用途是对**只读**数据在不同数据结构之间进行映射，**非常不建议**在映射后对源数据或映射结果做写操作。
3. 本工具大量使用reflect/unsafe等操作，对不同go版本的兼容性测试没有进行全面测试，在**生产环境**使用前**建议**您进行全面的效果测试以保证您的项目正常运行（注：经迭代，本工具目前的Hack代码仅剩string和[]byte的零拷贝）
4. 由于我的测试文件功能覆盖不全且写的比较随意，故没有提交到仓库，仅在本地自测。如有测试需求，请自行编写测试文件
5. 当源类型与目标类型**完全一致**（Same Type）时，本工具会执行**按位直接读写**的**浅拷贝(Direct Memory Copy)**，而不是完全深拷贝。这意味着在同类型转换场景下，即使数据实例存在循环引用（环状结构），也能成功完成映射（复制了指针，保持了环状结构），而不会触发**栈溢出（Stack Overflow）**。
6. 出于性能和功能定位考虑：本工具**不支持**转换**数据实例形成闭环**（内存地址循环引用，如 A->B->A/A->A）的数据结构（除非满足上述同类型浅拷贝条件），强行转换会导致无限递归并引发**栈溢出（Stack Overflow）**。请确保源数据符合预期。（注：本工具**支持**结构体类型定义层面的递归嵌套，只要运行时数据不成环即可）

## 主要用途

- **数据层转换**：在API、数据库模型、服务间通信等场景中进行数据结构转换
- **协议转换**：原生Go类型与Protocol Buffers类型之间的互相转换
- **类型适配**：不同类型系统之间的转换与适配（如字符串与数值、数组与字符串等）
- **数据序列化/反序列化**：在不同表示形式之间转换数据

## 特性

- 高效的内存操作（使用unsafe直接操作内存，减少开销）
- 类型安全的API（基于Go泛型）
- 丰富的转换选项（时间格式化、字段映射、自定义转换逻辑等）
- 全面的错误处理与结果包装
- 优雅处理复杂嵌套结构
- 支持Protocol Buffers类型

## 目录

### 基础类型转换
- [基本类型转换](#基本类型转换)
- [基础类型与any互转](#基础类型与any互转)
- [指针类型转换](#指针类型转换)
- [时间类型处理](#时间类型处理)

### 复合类型转换
- [结构体转换](#结构体转换)
- [切片和数组转换](#切片和数组转换)
- [Map转换](#map转换)

### 高级功能
- [Protocol Buffers转换](#protocol-buffers转换)
- [结果处理和错误检查](#结果处理和错误检查)
- [两阶段转换](#两阶段转换)
- [自定义转换器](#自定义转换器)

### 配置与优化
- [全局选项设置](#全局选项设置)
- [注意](#注意)

## 详细用法

### 基本类型转换

```go
import (
    "github.com/smgrushb/conv"
)

// 字符串到数值类型
intVal := conv.OstrichConvert[int]("123")         // 123
int64Val := conv.OstrichConvert[int64]("123")     // 123
floatVal := conv.OstrichConvert[float64]("123.45") // 123.45

// 数值类型互转
int32Val := conv.OstrichConvert[int32](int64(123456))  // 123456
uintVal := conv.OstrichConvert[uint](123)             // 123

// 布尔值转换
boolVal := conv.OstrichConvert[bool](1)         // true
boolVal2 := conv.OstrichConvert[bool]("true")   // true
intFromBool := conv.OstrichConvert[int](true)   // 1

// 带错误处理的转换（推荐用于生产环境）
if result, err := conv.Convert[int]("123abc"); err != nil {
    // 转换成功
    intVal := result
} else {
    // 处理错误
}
```

### 基础类型与any互转

Conv库支持Go的基础类型和`any`类型(interface{})之间的无缝转换，便于处理动态类型数据和类型未知的场景。

#### 基础类型转any

将任何类型转换为`any`类型非常简单，Conv能保留原始值的类型信息：

```go
// 基础类型转any
intVal := 42
anyVal := conv.OstrichConvert[any](intVal)  // any类型，内部保存int值42

// 复杂类型转any
type User struct {
    Name string
    Age  int
}
user := User{Name: "John", Age: 30}
anyUser := conv.OstrichConvert[any](user)  // any类型，内部保存User结构体
```

#### any转基础类型

从`any`类型转换到具体类型时，Conv会尝试根据目标类型进行适当的转换：

```go
// any转基础类型
anyVal := any(42)                           // int类型存储在any中
intFromAny := conv.OstrichConvert[int](anyVal)      // 42
floatFromAny := conv.OstrichConvert[float64](anyVal) // 42.0
strFromAny := conv.OstrichConvert[string](anyVal)    // "42"

// 复杂类型
anyString := any("hello")  
bytesFromAny := conv.OstrichConvert[[]byte](anyString)  // []byte("hello")

// 智能类型匹配 - 即使存储类型不完全相同，也能尝试转换
anyFloat := any(42.5)                           // float64存储在any中
intFromFloat := conv.OstrichConvert[int](anyFloat)      // 42
```

#### 处理不确定类型的数据

在处理JSON、配置文件或其他动态数据源时，`any`转换特别有用：

```go
// 假设这是从JSON解析出来的数据
data := map[string]any{
    "id": 1,
    "name": "Product",
    "price": 99.9,
    "in_stock": true,
    "tags": []any{"electronics", "sale"},
}

// 提取并转换特定字段
id := conv.OstrichConvert[int64](data["id"])            // 1
name := conv.OstrichConvert[string](data["name"])       // "Product"
price := conv.OstrichConvert[float64](data["price"])    // 99.9
inStock := conv.OstrichConvert[bool](data["in_stock"])  // true

// 处理切片
tags := conv.OstrichConvert[[]string](data["tags"])     // []string{"electronics", "sale"}
```

#### 处理未知字段类型

处理字段类型未知的情况，可以使用类型断言加Conv转换：

```go
// 遍历map中的any值并根据实际类型处理
processAnyMap := func(data map[string]any) {
    for k, v := range data {
        switch v.(type) {
        case int, int64, float64:
            // 数值类型，转为int64处理
            numVal := conv.OstrichConvert[int64](v)
            fmt.Printf("数值字段 %s: %d\n", k, numVal)
        case string:
            // 字符串，保持原样
            strVal := conv.OstrichConvert[string](v)
            fmt.Printf("字符串字段 %s: %s\n", k, strVal)
        case bool:
            // 布尔值
            boolVal := conv.OstrichConvert[bool](v)
            fmt.Printf("布尔字段 %s: %t\n", k, boolVal)
        case []any:
            // 切片，转为string切片
            sliceVal := conv.OstrichConvert[[]string](v)
            fmt.Printf("切片字段 %s: %v\n", k, sliceVal)
        default:
            // 其他类型，转为string
            fmt.Printf("其他字段 %s: %s\n", k, conv.OstrichConvert[string](v))
        }
    }
}
```

#### 注意事项

使用`any`类型转换时需要注意：

1. 虽然Conv尝试智能转换，但并非所有类型都能互相转换，转换失败时会返回目标类型的零值
2. 对于严格的类型安全需求，应使用`Convert`方法并检查错误，而非`OstrichConvert`
3. 从`any`转换到结构体时，确保结构体字段与源数据的键匹配（通过标签或字段名）
4. 性能考虑：使用`any`类型会引入反射开销，频繁操作时请考虑缓存转换结果或使用其他方案

### 指针类型转换

Conv库提供了强大的指针类型转换能力，支持各种指针类型之间的互相转换，以及指针与非指针类型之间的转换。

#### 基础指针类型转换

```go
// 不同指针类型之间的转换
intPtr := new(int)
*intPtr = 42
int64Ptr := conv.OstrichConvert[*int64](intPtr)  // *int64(42)
stringPtr := conv.OstrichConvert[*string](intPtr) // *string("42")

// nil指针处理
var nilIntPtr *int = nil
nilStringPtr := conv.OstrichConvert[*string](nilIntPtr) // nil

// 多级指针转换
intPtrPtr := &intPtr
int64PtrPtr := conv.OstrichConvert[**int64](intPtrPtr) // **int64(42)
```

#### 指针与值类型互转

Conv支持在指针和值类型之间进行智能转换，根据需要自动解引用或创建指针：

```go
// 值到指针
val := 42
valPtr := conv.OstrichConvert[*int](val)  // *int(42)

// 指针到值
intPtr := new(int)
*intPtr = 42
backToVal := conv.OstrichConvert[int](intPtr)  // 42

// 智能处理nil值
var nilPtr *int = nil
defaultVal := conv.OstrichConvert[int](nilPtr)  // 0 (零值)
```

#### 结构体指针转换

```go
// 源结构体
type Person struct {
    Name string
    Age  int
}

// 目标结构体
type User struct {
    Name string
    Age  string
}

// 结构体指针转换
person := &Person{Name: "John", Age: 30}
userPtr := conv.OstrichConvert[*User](person)
// userPtr.Name == "John"
// userPtr.Age == "30"

// 结构体指针到值
user := conv.OstrichConvert[User](person)
// user.Name == "John"
// user.Age == "30"
```

### 时间类型处理

```go
timeObj := time.Unix(1640995200, 0)

// time.Time转字符串（使用默认格式）
timeStr1 := conv.OstrichConvert[string](timeObj) // "2022-01-01 08:00:00"

// 自定义时间格式
timeStr2 := conv.OstrichConvert[string](timeObj, 
    option.TimeFormat("2006-01-02"))  // "2022-01-01"

// 字符串转time.Time
timeVal := conv.OstrichConvert[time.Time]("2022-01-01 08:00:00")

// 设置最小时间（小于此时间的会转为零值）
timeStr3 := conv.OstrichConvert[string](timeObj.Add(-time.Second),
    option.MinUnixBy(timeObj))
// 如果timeObj早于2022-01-01 08:00:00，则返回空字符串
```

#### 注意事项

字符串转time.Time时使用ParseInLocation附加本地时区，但存在两个特殊场景使用UTC：
- 时间字面量是`1970-01-01 00:00:00`，即时间戳零值
- 时间字面量是`0001-01-01 00:00:00`，即time.Time零值

## 复合类型转换

### 结构体转换

#### 基于标签的字段映射

```go
// 数据库实体
type UserEntity struct {
    ID          int64   `json:"user_id"` // 注意这里是user_id
    Name        string  `json:"user_name"` // 注意这里是user_name
    Age         int     `json:"age"`
    IsActive    bool    `json:"is_active"`
    CreatedTime int64   `json:"created_time"` // 注意这里是created_time
}

// API响应DTO
type UserDTO struct {
    ID        string  `json:"id"`
    Name      string  `json:"name"`
    Age       string  `json:"age"`
    Active    bool    `json:"is_active"`
    CreatedAt string  `json:"created_at" conv:"created_time"` // conv标签指定映射源字段
}

// 基本结构体转换
entity := UserEntity{
    ID: 123, 
    Name: "John", 
    Age: 30, 
    IsActive: true, 
    CreatedTime: 1641789600,
}

var dto UserDTO
err := conv.ConvertTo(entity, &dto)
// dto会包含所有匹配的字段，CreatedAt会匹配到entity.CreatedTime
// 因为在dto中使用了conv:"created_time"标签

// 注意: conv标签会覆盖json标签，如果源结构体中没有对应conv标签的字段，
// 该字段将不会被映射，会保持零值
```

#### 标签映射机制说明

Conv库在结构体转换中的映射原则：

1. **默认使用json标签**：默认情况下，Conv使用结构体字段的json标签进行匹配
2. **优先级标签override**：如果设置了conv标签，它会完全覆盖json标签
3. **标签不回退**：当使用conv标签指定映射时，如果找不到对应标签的源字段，不会回退到json标签匹配
4. **全局自定义**：可以通过`SetStructTageName`和`SetStructPriorityTagName`更改默认和优先标签名

```go
// 更改默认使用的标签
conv.SetStructTageName("bson") // 使用bson标签代替json

// 更改优先级标签
conv.SetStructPriorityTagName("mapping") // 使用mapping标签代替conv作为优先标签
```

#### 嵌套结构体转换

```go
type Address struct {
    Street  string `json:"street"`
    City    string `json:"city"`
    ZipCode string `json:"zip_code"`
}

type UserWithAddress struct {
    ID      int     `json:"user_id"`
    Name    string  `json:"user_name"`
    Address Address `json:"address"`
}

type AddressDTO struct {
    Street  string `json:"street"`
    City    string `json:"city"`
    ZipCode int    `json:"zip_code"` // 注意类型不同
}

type UserWithAddressDTO struct {
    ID      string     `json:"id"`
    Name    string     `json:"name"`
    Address AddressDTO `json:"address"`
}

// 嵌套结构体会自动递归转换
user := UserWithAddress{
    ID: 1,
    Name: "John",
    Address: Address{
        Street: "123 Main St",
        City: "New York",
        ZipCode: "10001",
    },
}

userDTO := conv.OstrichConvert[UserWithAddressDTO](user)
// userDTO.Address.ZipCode 会自动从字符串转为整数
```

#### 匿名结构体转换

Conv库支持Go的匿名结构体（嵌入式字段）转换，包括带指针和非指针类型的嵌入字段。匿名结构体的字段会被"提升"到外层结构体，Conv可以正确处理这种情况。

```go
// 基础信息结构体
type BaseInfo struct {
    ID        int64  `json:"id"`
    CreatedAt int64  `json:"created_at"`
    UpdatedAt int64  `json:"updated_at"`
}

// 地址信息结构体
type AddressInfo struct {
    Country string `json:"country"`
    City    string `json:"city"`
    Street  string `json:"street"`
}

// 带非指针匿名字段的结构体
type CustomerEntity struct {
    BaseInfo               // 嵌入BaseInfo，字段会被提升
    Name      string       `json:"name"`
    Age       int          `json:"age"`
    AddressInfo            // 嵌入AddressInfo，字段会被提升
}

// 带指针匿名字段的结构体
type VipCustomerEntity struct {
    *BaseInfo              // 嵌入BaseInfo指针
    Name       string      `json:"name"`
    Level      int         `json:"level"`
    *AddressInfo           // 嵌入AddressInfo指针
}

// 目标DTO结构体
type CustomerDTO struct {
    ID        string `json:"id"`
    Name      string `json:"name"`
    Age       string `json:"age"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
    Country   string `json:"country"`
    City      string `json:"city"`
    Street    string `json:"street"`
}

// 非指针匿名字段转换示例
customer := CustomerEntity{
    BaseInfo: BaseInfo{
        ID:        123,
        CreatedAt: 1641789600,
        UpdatedAt: 1641889600,
    },
    Name: "John",
    Age:  30,
    AddressInfo: AddressInfo{
        Country: "USA",
        City:    "New York",
        Street:  "Fifth Avenue",
    },
}

// 转换时会自动处理提升的字段
customerDTO := conv.OstrichConvert[CustomerDTO](customer)
// customerDTO.ID, customerDTO.CreatedAt, customerDTO.UpdatedAt会从BaseInfo获取
// customerDTO.Country, customerDTO.City, customerDTO.Street会从AddressInfo获取

// 指针匿名字段转换示例
vipCustomer := VipCustomerEntity{
    BaseInfo: &BaseInfo{
        ID:        456,
        CreatedAt: 1651789600,
        UpdatedAt: 1651889600,
    },
    Name:  "Alice",
    Level: 2,
    AddressInfo: &AddressInfo{
        Country: "Canada",
        City:    "Toronto",
        Street:  "King Street",
    },
}

// 同样可以正确处理指针类型的嵌入字段
vipCustomerDTO := conv.OstrichConvert[CustomerDTO](vipCustomer)

// 处理nil指针的情况
nilPointerCustomer := VipCustomerEntity{
    BaseInfo: &BaseInfo{
        ID:        789,
        CreatedAt: 1661789600,
        UpdatedAt: 1661889600,
    },
    Name:        "Bob",
    Level:       3,
    AddressInfo: nil, // 空指针
}

// Conv将跳过nil指针的匿名字段，相关目标字段会保持零值
nilPointerDTO := conv.OstrichConvert[CustomerDTO](nilPointerCustomer)
// nilPointerDTO.Country, nilPointerDTO.City, nilPointerDTO.Street会是零值
```

在上面的例子中，Conv库处理匿名结构体字段的方式：

1. **非指针匿名字段**：嵌入的字段会直接提升并参与转换过程
2. **指针类型匿名字段**：如果指针非nil，会解引用并提升字段进行转换
3. **nil指针匿名字段**：会跳过nil指针引用的字段，相关目标字段会保持零值
4. **多层嵌套**：支持多级嵌套的匿名字段转换

这种设计支持Go的组合模式，方便从数据库实体层到API层的数据转换，避免了冗余的手动映射代码。

#### 方法调用转换

Conv库支持一种强大的结构体转换功能：在转换过程中调用源结构体的方法，并将方法的返回值赋值给目标结构体的对应字段。这使得转换过程不仅可以映射现有字段，还能执行计算和业务逻辑。

##### 基本方法调用

当目标结构体的字段名在源结构体中不存在对应字段，但存在同名方法时，Conv会自动调用该方法并使用其返回值：

```go
// 源结构体
type Product struct {
    Price    float64 `json:"price"`
    Quantity int     `json:"quantity"`
}

// 计算总价的方法
func (p Product) Total() float64 {
    return p.Price * float64(p.Quantity)
}

// 目标结构体
type ProductDTO struct {
    Price    string  `json:"price"`
    Quantity string  `json:"quantity"`
    Total    string  // 没有标签时，字段名直接匹配源结构体的Total()方法
}

// 转换示例
product := Product{Price: 15.5, Quantity: 3}
dto1 := conv.OstrichConvert[ProductDTO](product)

// dto1.Total 将包含 "46.5" (15.5 * 3 的结果转为字符串)
```

重要：当目标字段使用了json标签时，需要额外处理才能触发方法调用。有两种解决方案：

1. 使用conv标签明确指向源结构体的方法名：

```go
// 目标结构体使用conv标签
type ProductDTOWithTags struct {
    Price    string  `json:"price"`
    Quantity string  `json:"quantity"`
    Total    string  `json:"total" conv:"Total"`  // 使用conv标签明确指向Total()方法
}

dto2 := conv.OstrichConvert[ProductDTOWithTags](product)
// dto2.Total 将包含 "46.5"
```

2. 使用option.AliasMap设置方法名到字段名的映射：

```go
// 目标结构体只有json标签
type ProductDTOWithJsonTags struct {
    Price    string  `json:"price"`
    Quantity string  `json:"quantity"`
    Total    string  `json:"total"`  // 只有json标签
}

// 使用别名映射
dto3 := conv.OstrichConvert[ProductDTOWithJsonTags](product, 
    option.AliasMap(map[string]string{
        "Total": "total",  // 将Total()方法映射到total字段
    }))
// dto3.Total 将包含 "46.5"
```

这种方式可以避免在转换前后编写大量手动计算代码，使数据转换更加简洁。

##### 带错误/状态检查的方法调用

Conv库对于返回多个值的方法有特殊处理：如果方法返回两个值，且第二个返回值类型为`error`或`bool`，Conv会根据第二个返回值决定是否进行赋值：

```go
// 源结构体
type Account struct {
    encryptedPassword string
}

// 返回错误的方法
func (u Account) Password() (string, error) {
    if u.encryptedPassword == "" {
        return "", errors.New("password not set")
    }
    // 解密密码的逻辑
    return "decrypted-password", nil
}

// 返回布尔值的方法
func (u Account) FullName() (string, bool) {
    // 假设根据某些条件决定是否有完整姓名
    return "John Doe", true
}

// 目标结构体
type AccountDTO struct {
    Password string `json:"password" conv:"Password"` // 使用conv标签指向方法
    FullName string `json:"full_name" conv:"FullName"` // 使用conv标签指向方法
}

// 转换示例
account := Account{encryptedPassword: "encrypted-data"}
dto := conv.OstrichConvert[AccountDTO](account)

// dto.Password 将被设置为 "decrypted-password"，因为方法返回nil错误
// dto.FullName 将被设置为 "John Doe"，因为方法返回true
```

转换逻辑如下：
- 对于`(value, error)`返回类型，只有当error为nil时才会赋值
- 对于`(value, bool)`返回类型，只有当bool为true时才会赋值
- 如果根据条件不赋值，目标字段将保持零值

##### 方法类型字段

除了调用结构体方法外，Conv还支持调用结构体中的方法类型字段：

```go
// 源结构体
type DataProcessor struct {
    CalculateTotal func() float64           // 无参数方法字段
    FormatData     func(data string) string // 带参数方法字段
}

// 设置方法字段
processor := DataProcessor{
    CalculateTotal: func() float64 { return 100.5 },
    FormatData: func(data string) string { 
        return strings.ToUpper(data) 
    },
}

// 目标结构体
type ProcessorDTO struct {
    CalculateTotal string `json:"calculate_total" conv:"CalculateTotal"` // 使用conv标签
    FormatData     string `json:"format_data"`     // 与无参数函数字段同名
}

// 转换示例
dto := conv.OstrichConvert[ProcessorDTO](processor)

// dto.CalculateTotal 将包含 "100.5"
// 注意：对于带参数的方法字段(如FormatData)，Conv不会自动调用，因为无法确定参数值
```

注意事项：
1. 只有无参数的方法和函数字段才会被自动调用
2. 方法名称区分大小写，必须与目标字段名完全匹配
3. 方法必须是导出的(首字母大写)才能被调用

#### 结构体转换map

Conv支持将结构体转换成`map[string]string`和`map[string]any`

```go
type User struct {
    id    int
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Score int    `json:"score"`
}

// 方法或方法类型字段的返回值不会被映射到map中
func (u *User) Id() int{
	return u.id
}

type Community struct {
    Member []User `json:"member"`
}
type CommunityDTO struct {
    Member []*map[string]string `json:"member"`
}
user := User{id: 1, Name: "John", Age: 30}
userMap1 := conv.OstrichConvert[map[string]any](user) // {"name": "John", "age": 30, "score": 0}

community := Community{Member: []User{user}}
dto := conv.OstrichConvert[CommunityDTO](community) // {"member":[{"name": "John", "age": "30", "score": 0}]}

// 默认不携带私有字段，默认不跳过零值字段
userMap2 := conv.OstrichConvert[map[string]string](user, option.IncludePrivateFields(), option.IgnoreEmptyFields()) // {"id": "1", "name": "John", "age": 30}
```

### 切片和数组转换

```go
// 基本切片转换
intSlice := []int{1, 2, 3}
int64Slice := conv.OstrichConvert[[]int64](intSlice)  // [1, 2, 3]
stringSlice := conv.OstrichConvert[[]string](intSlice) // ["1", "2", "3"]

// 字符串与切片互转
csv := "a,b,c"
strSlice1 := conv.OstrichConvert[[]string](csv,
option.CustomConverter(convextend.String2Strings())) // ["a", "b", "c"]

// 自定义分隔符
custom := "a|b|c"
strSlice2 := conv.OstrichConvert[[]string](custom, 
    option.CustomConverter(convextend.String2Strings().Sep("|"))) // ["a", "b", "c"]

// 数字字符串转数字切片
numStr := "1,2,3"
numSlice := conv.OstrichConvert[[]int](numStr, 
    option.CustomConverter(convextend.String2Ints())) // [1, 2, 3]

// 自定义空字符串处理策略
emptyStr := ""
emptySlice := conv.OstrichConvert[[]string](emptyStr, 
    option.CustomConverter(convextend.String2Strings().SplitStrategy(convextend.EmptySplit)))
// 使用EmptySplit会返回空切片[]，而不是[""]

// 切片转字符串
ints := []int{1, 2, 3}
str := conv.OstrichConvert[string](ints, 
    option.CustomConverter(convextend.Ints2String().Sep("-")))  // "1-2-3"
```

### Map转换

```go
// 基本map转换（键和值类型都可转换）
mapIntStr := map[int]string{1: "a", 2: "b"}
mapStrStr := conv.OstrichConvert[map[string]string](mapIntStr)  // {"1": "a", "2": "b"}

// Map转KVP切片
mapData := map[string]int{"a": 1, "b": 2}
kvpSlice := conv.OstrichConvert[[]model.KeyValuePair[string, int]](mapData,
    option.CustomConverter(convextend.Map2KVP[string, int](model.NewKeyValuePair[string, int]())))
// 转换为包含键值对结构体的切片

// 自定义KVP类型
type MyKVP struct {
    K string `json:"key"`
    V int    `json:"value"`
}

func (m *MyKVP) GetKey() string { return m.K }
func (m *MyKVP) SetKey(k string) { m.K = k }
func (m *MyKVP) GetValue() int { return m.V }
func (m *MyKVP) SetValue(v int) { m.V = v }

customKVPSlice := conv.OstrichConvert[[]MyKVP](mapData,
    option.CustomConverter(convextend.Map2KVP[string, int](&MyKVP{})))

// KVP切片转回Map
kvpToMap := conv.OstrichConvert[map[string]int](kvpSlice,
    option.CustomConverter(convextend.KVP2Map[string, int](model.NewKeyValuePair[string, int]())))
```

## 高级功能

### Protocol Buffers转换

```go
import (
    "github.com/smgrushb/conv"
    "github.com/smgrushb/conv/extend"
    "google.golang.org/protobuf/types/known/timestamppb"
    "google.golang.org/protobuf/types/known/wrapperspb"
    "google.golang.org/protobuf/types/known/durationpb"
    "google.golang.org/protobuf/types/known/structpb"
)

func init() {
    // 初始化所有Protocol Buffers转换器
	convextend.ConvProto()
}

// 时间类型转换
ti := time.Now()
pbTimestamp := conv.OstrichConvert[*timestamppb.Timestamp](ti)
backToTime := conv.OstrichConvert[time.Time](pbTimestamp)

// Duration转换
d := time.Hour * 24
pbDuration := conv.OstrichConvert[*durationpb.Duration](d)
backToDuration := conv.OstrichConvert[time.Duration](pbDuration)

// Wrapper类型转换
b := true
boolValue := conv.OstrichConvert[*wrapperspb.BoolValue](b)
i32 := int32(42)
intValue := conv.OstrichConvert[*wrapperspb.Int32Value](i32)
f64 := 3.14
doubleValue := conv.OstrichConvert[*wrapperspb.DoubleValue](f64)
str := "hello"
stringValue := conv.OstrichConvert[*wrapperspb.StringValue](str)
bytes := []byte("world")
bytesValue := conv.OstrichConvert[*wrapperspb.BytesValue](bytes)

// 从Wrapper类型转回原始类型
boolBack := conv.OstrichConvert[bool](boolValue)
intBack := conv.OstrichConvert[int32](intValue)
floatBack := conv.OstrichConvert[float64](doubleValue)
stringBack := conv.OstrichConvert[string](stringValue)
bytesBack := conv.OstrichConvert[[]byte](bytesValue)

// Struct类型转换
mapData := map[string]any{
    "name":  "John",
    "age":   30.0,
    "items": []any{"book", "pen"},
}
pbStruct := conv.OstrichConvert[*structpb.Struct](mapData)
backToMap := conv.OstrichConvert[map[string]any](pbStruct)

// ListValue转换
sliceData := []any{"a", 1.1, true}
pbList := conv.OstrichConvert[*structpb.ListValue](sliceData)
backToSlice := conv.OstrichConvert[[]any](pbList)
```

### 结果处理和错误检查

```go
// 标准错误处理模式（推荐用于生产环境）
result, err := conv.Convert[int]("123abc")
if err == nil {
    intVal := result
    // 使用转换后的值
} else {
    // 处理错误
}

// 两阶段转换带错误处理
result, err := conv.TwoPhaseConv[FinalType, IntermediateType](sourceValue)
if err == nil {
    finalVal := result.Value()
} else {
    // 处理错误
}

```

### 两阶段转换

```go
// 通过中间类型进行复杂转换
// 例：UserModel -> Map -> UserDTO

type UserModel struct {
    ID   int    `json:"user_id"`
    Name string `json:"user_name"`
}

type UserDTO struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}

// 两阶段转换：先转为map，再转为目标结构体
model := UserModel{ID: 123, Name: "John"}
dto := conv.OstrichTwoPhaseConvert[UserDTO, map[string]interface{}](model,
    option.AliasMap(map[string]string{
        "user_id": "id",
        "user_name": "name",
    }))

// 复杂场景：XML -> Map -> 结构体
xmlStr := `<user><id>123</id><name>John</name></user>`
// 假设有中间转换步骤将XML转为map
mapData := parseXML(xmlStr) // 自定义函数
dto = conv.OstrichTwoPhaseConvert[UserDTO, map[string]interface{}](mapData)
```

### 自定义转换器

```go
// 完整自定义转换器示例
import (
    "github.com/smgrushb/conv/internal"
    "reflect"
    "strings"
    "unsafe"
)

// 例：转换全大写字符串到标题格式
type UpperToTitleConverter struct{}

func (c *UpperToTitleConverter) Is(dstTyp, srcTyp reflect.Type) bool {
    return srcTyp.Kind() == reflect.String && dstTyp.Kind() == reflect.String
}

func (c *UpperToTitleConverter) Converter() func(dPtr, sPtr unsafe.Pointer) {
    return func(dPtr, sPtr unsafe.Pointer) {
        src := *(*string)(sPtr)
        if strings.ToUpper(src) == src {
            *(*string)(dPtr) = strings.Title(strings.ToLower(src))
        } else {
            *(*string)(dPtr) = src
        }
    }
}

func (c *UpperToTitleConverter) Key() string {
    return "[UpperToTitleConverter]"
}

// 使用自定义转换器
upperStr := "HELLO WORLD"
titleStr := conv.OstrichConvert[string](upperStr, 
    option.Custom(&UpperToTitleConverter{}))
// 返回 "Hello World"
```

## 配置与优化

### 全局选项设置

```go
// 更改结构体标签名称
conv.SetStructTageName("bson")  // 使用bson标签进行字段匹配

// 设置优先级标签
conv.SetStructPriorityTagName("priority")  // priority标签优先于默认标签
```

### 注意

- 自定义转换器应考虑线程安全
- 默认的"鸵鸟模式"（Ostrich）API会在转换失败时使用零值，适合不关心错误的场景
- 对于生产环境，推荐使用带错误处理的Convert API而非OstrichConvert
- 复杂嵌套结构转换可能会影响性能，考虑分步转换
- 零拷贝操作虽然高效但会创建对原始数据的依赖，确保原始数据在转换结果使用期间不被修改

## 依赖项

- `google.golang.org/protobuf`：对Protocol Buffers的支持