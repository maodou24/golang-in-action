# router

集中注册 HTTP 路由、路由组和对应处理函数，并组装需要应用的 `middleware`。

依赖 `controller` 和 `middleware`。只负责请求分发与中间件组合，不解析业务参数、不实现业务规则，也不直接访问数据层。
