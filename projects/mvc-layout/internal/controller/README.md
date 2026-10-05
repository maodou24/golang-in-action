# controller

HTTP 请求处理入口，负责解析请求、参数校验和必要的边界转换；调用 `service` 执行业务，并将结果组装为 API 响应。

依赖 `api/v1/request`、`api/v1/response` 和 `service`。不直接实现业务规则、注册路由或访问数据库。
