## 文件目录
| 文件/目录名称 | 说明 |
|  ------  |  ------------  |
| app/common | 公共模块（请求、响应结构体等） |
| app/controllers | 业务调度器 |
| app/middleware | 中间件 |
| app/models | 数据库结构体 |
| app/services | 业务层 |
| bootstrap | 项目启动初始化 |
| config | 配置结构体 |
| global | 全局变量 |
| routes | 路由定义 |
| static | 静态资源（允许外部访问） |
| storage | 系统日志、文件等静态资源） |
| utils | 工具函数 |
| config.yaml | 配置文件 |
| main.go | 项目启动文件 |

## bootstrap目录中文件作用
| 文件名称 | 说明 |
| ------ | ------------ |
| config | 初始化项目配置(viper) |
| log | 初始化日志配置(zap)(lumberjack) |
| db | 初始化数据库配置(gorm) |

## 安装Gin
`go get -u github.com/gin-gonic/gin`

## 初始化 go.mod 文件
`go mod init <项目名称>`
