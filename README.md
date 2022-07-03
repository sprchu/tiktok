# tiktok

此项目为字节跳动青训营后端大项目，客户端由字节跳动提供，此仓库仅为后端代码。

api 文档地址：https://www.apifox.cn/apidoc/shared-56b867b2-e909-4c0d-a033-10328c97046e/api-21003227


## 项目组织

```
.
├── api                     # 总路由
│  ├── config
│  ├── Dockerfile
│  ├── main.go
│  ├── social               # 社交相关。下面是 gozero 标准 layout
│  │  ├── api.go
│  │  └── internal
│  │     ├── handler
│  │     ├── logic          # 主要业务逻辑
│  │     ├── middleware
│  │     ├── svc
│  │     └── types          # 根据 .api 文件生成的类型
│  ├── user                 # 用户相关。同上
│  └── videomgr             # 视频相关。同上
│     ├── api.go
│     ├── internal
│     └── storage           # 上传文件接口相关代码
│        ├── local.go
│        └── storage.go
│
├── docker-compose.yml
├── go.mod                  # 统一各服务依赖关系
│
├── schema                  # 项目 schema
│  ├── api                  # api 文件
│  ├── sql                  # sql 文件
│  └── proto                # proto 文件
│
├── servebase               # 服务基础代码
│
├── social                                  # 社交服务。下面是 gozero 标准 layout
│  ├── Dockerfile
│  ├── model                                # db 层
│  └── rpc                                  # rpc 层
│     ├── internal
│     │  ├── config
│     │  ├── logic                          # 主要业务逻辑
│     │  ├── server
│     │  └── svc
│     ├── service
│     ├── social.go
│     └── types                             # 根据 .proto 文件生成的类型
│        └── social
│           ├── social.pb.go
│           └── social_grpc.pb.go
│
├── template                                # gozero 生成模板，仅作少量修改
│  ├── api
│  └── rpc
│
├── tests                                   # 测试相关
│
├── user                                    # 用户服务，同上
│  ├── Dockerfile
│  ├── model
│  └── rpc
└── videomgr                                # 视频服务，同上
   ├── Dockerfile
   ├── model
   └── rpc

```
