# xerrors

- 错误栈
- 可扩展
    - 可以附加错误码
    - 可以附加用户文案
    - 可以附加具体错误详情
- WithMessage "aa" not found
- WithStack
- Wrap
- Is
- Cause Unwrap
- As
- 序列化和反序列化
- NewWithContext 记录错误ID
- FormatError 调试模式，
- **错误链**
- Named 命名空间
- 错误码
- 用户文案，最好是在逻辑层（service logic）增加。（流程）
- 三类错误
    - 上游系统错误
    - 本系统错误（包含系统以来的数据库等本服务资源）
    - 下游系统错误（下游业务系统）
- 尽量避免 sentinel 错误，应当立即处理。
    - 例如 gorm.ErrNotFound 可以通过指针判空或者数组长度判零，或者通过显示的返回

## TODO

https://go.googlesource.com/proposal/+/master/design/go2draft.md
https://github.com/golang/exp/tree/master/errors
import "golang.org/x/xerrors"

```
有个问题
处理错误的时候，不清楚错误是带栈的还是不带栈的
一种方法是，所有返回错误的地方都 Wrap 一下
另一种方法是，代码分层比较好。这样也可以区分

还有个问题
就是错误码的问题。
这个包，没有考虑错误码。
这个也需要自己额外关注。但是错误码，只是协议层(go-kit Transport层，序列化和反序列化错误时需要的)需要的东西。

还有个问题
就是区分内部错误和外部错误。
内部错误，可能直接占用一个码，不需要对外告诉具体细节（生产环境）
外部错误，需要把错误信息，返回给调用方。

产生错误的时候，不一定知道是内部错误，还是外部错误
// 产生错误的时候，只表明是什么错误（是什么类型的错误），而不表明是内部错误还是外部错误
```
