# CampusTrade
This is a simple managing system like project built by GoLang
### Project Structure
├───.github   // 存放Github工作流相关配置
│   ├───PULL_REQUEST_TEMPLATE    // PR模板
│   └───workflows    // 提交触发工作流
├───.idea     // IDEA界面配置文件夹
├───benchmark    // 性能测试文件
├───configs   // 配置文件
├───docs      // 项目相关文档
├───internal   // 项目源代码
│   ├───api   // 对外暴露的接口
│   ├───core  // 项目核心代码
│   ├───db    // 数据库相关代码
│   │   ├───mysql   // 存放MySQL的表映射以及相关操作
│   │   └───redis   // 存放Redis相关代码
│   ├───middleware  // 自定义中间件
│   ├───router      // 路由
│   ├───service     // 业务逻辑代码
│   │   ├───admin     // 管理员部分
│   │   ├───item      
│   │   ├───notification     // 站内通知
│   │   ├───transaction      // 交易
│   │   └───user             // 用户
│   └───utils     // 自定义工具结构体或方法函数
├───scripts    // 初始化脚本
└───test    // 单元测试以及集成测试