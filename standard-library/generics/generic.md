# Generic

## 定义

```text
[P any]
[P ~any]   // 底层类型为P的任何类型
[S interface{ ~[]byte|string }]
[S ~[]E, E any]
[P Constraint[int]]
[_ any]
```