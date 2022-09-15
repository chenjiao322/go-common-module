## go-common-module

[![pipeline status](https://gitlab.geinc.cn/services/go-common-module/badges/master/pipeline.svg)](https://gitlab.geinc.cn/services/go-common-module/commits/master)

公司内使用公共库, 减少多个项目间相同的代码 项目中包括:

- **纯函数**
- 接入中间件的推荐方法
- 第三方库的插件

### 引入项目

##### 1. 配置私有仓库

命令行

`go env -w GOPRIVATE=gitlab.geinc.cn`

goland(windows)

<kbd>settings</kbd> > <kbd>Go</kbd> > <kbd>Go Modules</kbd> > <kbd>Environment</kbd> > <kbd>
GOPRIVATE=gitlab.geinc.cn</kbd>

goland(macOS)

<kbd>Preferences</kbd> > <kbd>Go</kbd> > <kbd>Go Modules</kbd> > <kbd>Environment</kbd> >
<kbd>GOPRIVATE=gitlab.geinc.cn</kbd>

2. 拉取依赖

`go get -u gitlab.geinc.cn/services/go-common-module`

可在goland的 <kbd>Run/Debug Configurations</kbd> > <kbd>Before Launch</kbd> > <kbd>go cmd</kbd> 中
添加`get -u gitlab.geinc.cn/services/go-common-module`,自动更新依赖

3. 出现错误时

可以先尝试直接git clone本仓库, 可以会出现需要输入密码的情况,先`git config --global credential.helper store`
再输入用户名密码,直到可以不输入任何信息,直接clone下来,再尝试 `go get`

##### 2.代码更新

- 建议始终保持最新版本
- 尽量保证较高的测试覆盖率
- 为减少管理成本,早期不使用tag,等到更新速度较慢后使用tag管理版本
- 每个文件夹中必须写readme.md,简述package的用途和用法 


#### 3. 单元测试

`go test ./... -cover`


