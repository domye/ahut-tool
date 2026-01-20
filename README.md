# AHUT-Tool

AHUT-Tool 是一个基于 Wails 框架开发的桌面应用程序，旨在为安徽工业大学（AHUT）学生提供便捷的工具。

## 功能特性

- **教务系统登录**：安全地登录学校教务系统
- **成绩查询**：快速查询各学期成绩及统计信息
- **现代化界面**：使用 Vue3 和 Ant Design Vue 构建的用户友好界面
- **桌面应用**：原生桌面应用体验，无需浏览器即可使用

## 技术栈

- **前端**: Vue 3、TypeScript、Ant Design Vue、Pinia、Vue Router
- **后端**: Go 语言
- **框架**: Wails（用于构建桌面应用）
- **HTTP 客户端**: Resty
- **UI 组件库**: Ant Design Vue

## 项目结构

```
ahut-tool/
├── backend/              # Go 后端代码
│   ├── jwxt/             # 教务系统相关功能
│   ├── models/           # 数据模型定义
│   ├── pay/              # 支付相关功能
│   ├── utils/            # 工具函数
│   └── app.go            # 应用主入口
├── frontend/             # 前端代码
│   ├── src/              # 源代码
│   │   ├── components/   # Vue 组件
│   │   ├── views/        # 页面视图
│   │   ├── router/       # 路由配置
│   │   └── store/        # Pinia 状态管理
│   ├── public/           # 静态资源
│   └── package.json      # 前端依赖配置
├── main.go               # 应用程序主入口
└── wails.json            # Wails 配置文件
```

## 安装与运行

### 前提条件

- Go 1.18+
- Node.js 16+
- npm 或 yarn

### 开发环境设置

1. 克隆仓库到本地：

```bash
git clone <repository-url>
cd ahut-tool
```

2. 安装前端依赖：

```bash
cd frontend
npm install
```

3. 返回项目根目录并安装 Go 模块：

```bash
cd ..
go mod tidy
```

4. 运行开发模式：

```bash
wails dev
```

### 生产构建

要构建最终的应用程序：

```bash
wails build
```

这将生成一个可在目标平台上运行的独立可执行文件。

## 使用方法

1. 启动应用程序
2. 在登录页面输入学号和密码
3. 登录成功后，可以查询成绩等教务信息
4. 应用会安全地处理你的登录凭据

## 安全性

- 用户密码通过 Base64 编码和加密传输
- 所有敏感数据在本地处理，不会上传至第三方服务器
- Cookie 管理确保会话安全

## 贡献

欢迎提交 Issue 和 Pull Request 来帮助改进这个项目！

## 许可证

该项目采用 MIT 许可证 - 查看 [LICENSE](./LICENSE) 文件了解详情。

## 作者

- **Domye** - [a1523610551@163.com](mailto:a1523610551@163.com)

## 致谢

- 感谢 Wails 团队提供的优秀框架
- 感谢 Vue、Go 社区提供的强大工具