# desktop-components
golang 桌面组件

# 环境准备
## 安装go,qt
## 按照要求配置环境变量
* QT_DIR
* QT_VERSION
## 项目跟目录执行如下命令
* `go mod vendor`
* `go mod tidy`
* `qtsetup`

# 打包
项目根目录下执行 `sh scripts\build.sh pkg\helloworld\helloworld.go `