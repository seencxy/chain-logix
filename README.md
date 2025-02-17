# 区块链物流管理平台

- 为什么要开发这个平台?
- 平台使用了什么技术?
- 适用于哪些人群和组织？

## 为什么要开发这个平台

全球物流行业正面临多重挑战，包括货物跟踪的不透明、低效、安全性问题，以及复杂的供应链管理。随着全球化的深入，提供有效、透明和可靠的物流解决方案变得至关重要。为了迎合这一需求，我们推出了区块链物流管理平台，充分利用区块链技术的不可篡改性、去中心化和透明度，为现代物流提供创新解决方案。

## 平台使用的技术

平台采用最新的区块链技术，特别是联盟链模式，以确保数据的安全性和高效的交易处理。我们整合了智能合约技术，实现自动处理复杂的物流流程，如自动支付和合同执行。采用前后端分离架构，后端使用 **Go** 语言开发，采用 Gin 框架；前端使用 **Vue3**，结合 **Element-UI** 和 **Daisy-UI** 等多种 UI 库；数据存储采用 MySQL、Redis，并重要数据上链部署。

## 适用群众和组织

这个平台适用于各种规模的物流公司、供应链管理者、制造商以及零售商。它尤其适合那些寻求优化供应链管理、降低运营成本、提高货物追踪能力和增强客户信任的企业。



## 身份分配

### 1. 供应链管理员（Supply Chain Manager）

**职责**：

- 监控和管理整个供应链流程。
- 使用区块链平台来追踪货物从生产到交付的每个阶段。
- 确保所有交易和记录在区块链上准确无误地执行和更新。
- 分析供应链数据，以优化流程和减少浪费。

**特权**：

- 访问全面的供应链数据和分析报告。
- 能够发起和管理智能合约，如自动付款和交货确认。
- 实施和更新供应链安全和合规性措施。

### 2. 用户（user）

**特权**：

* 能够方便地进行物流追踪。
* 具备发送物流请求的能力。

### 3. 总管理员（admin）

**职责: **

* 有效管理平台用户。
* 负责审核供应链管理员。



## 平台开发技术栈

#### 前端

- **技术栈：** 使用Vue.js结合Vite构建高效的前端应用。Vite作为现代化的前端构建工具，可以提供快速的热重载和丰富的插件生态。
- **UI框架：** 采用Tailwind CSS组件库，这是一个功能强大的实用工具优先的CSS框架，可以快速构建定制化的响应式界面。
- **用户体验：** 重点关注用户交互和界面设计，确保界面直观、易用，特别是对于老年用户来说，界面应简洁、易懂。
- **图表组件(chart): ** 添加chart组件，通过引入这个图表组件，我们希望增强用户对物流运输的大概数量，从而帮助他们做出更明智的决策，能够更加直观的分析。
- **高德地图: ** 因为需要实现住总

#### 后端

- **编程语言：** 使用Go语言，它以高效的性能和简洁的语法著称，非常适合构建高性能的服务端应用。
- **Web框架：** 选择Gin框架，这是一个高性能的HTTP Web框架，简单易用且效率高，适合快速开发。
- **与区块链的交互：** 需要开发与FISCO BCOS区块链平台交互的接口，包括智能合约的调用、交易处理等。
- **Alipay: ** 选择支付宝作为支付平台，方便用户付款以及退款，安全信用度高，以及后续养老金发放也采用支付宝账号转账，避免出现资金问题。
- **数据库: ** 本平台将重要数据存储在链上，但并不是将所有数据存储在链上，将部分私密数据存储在mysql，将数据持久化，因为有验证码，将验证码等有过期时间的数据需要及时更新。

#### 区块链平台

* **webase: ** 搭建webase作为合约部署
* **Fisco Bcos Browser: **搭建fisco bcos作为区块浏览器，提供给用户进行区块以及交易搜索



