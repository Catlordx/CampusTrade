## 项目要求

### 1.编码规范

项目编码规范严格使用[Google](https://go.dev/doc/effective_go)编码规范
1. 代码必须通过Gofmt格式化

2. 字段、结构体和函数等在包外的可见性由首字母大小写确定

3. 包名必须全部为小写字母，并且仅包含一个单词，不得使用下划线和驼峰式命名。根据公约，包名应为其所在目录的目录名

   > Another convention is that the package name is the base name of its source directory; the package in `src/encoding/base64` is imported as `"encoding/base64"` but has name `base64`, not `encoding_base64` and not `encodingBase64`.

4. 不要定义Getter和Setter函数，如果你希望由类似于Getter函数的功能，假设有一个字段名为owner,那么可以定义一个方法Owner()返回owner的值。如果希望有类似于Setter函数的功能可以定义为SetOwner()

5. 按照惯例，单方法接口以方法名加上-er 后缀或类似的修饰来命名，以构造一个代理名词：如阅读器、书写器、格式化器、关闭提示器等。

6. 读取、写入、关闭、刷新、字符串等都有规范的签名和含义。为避免混淆，除非你的方法具有相同的签名和含义，否则不要使用这些名称。反之，如果您的类型实现了一个与知名类型上的方法具有相同含义的方法，则应赋予其相同的名称和签名；将字符串转换方法称为 String，而不是 ToString。

7. 在Go中，采用驼峰式命名法而非下划线分割多词标识符

8. 

### 2.测试工具
暂定使用Go标准库中的testing模块进行测试、benchmark
assert采用社区提供的包 "github.com/stretchr/testify/assert"
使用官方提供的`pprof`包进行性能分析



### 3.CI/CD
CI/CD借助GithubWorkFlow实现
