namespace go base

struct ThriftAny {
    1: string TypeUrl = ""
    2: binary Value = ""
}

struct User4 {
    1: string Name = "",
    2: ThriftAny Value,
}
