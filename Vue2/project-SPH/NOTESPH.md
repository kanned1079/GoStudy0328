## 初始化项目
1. 设置运行时自动打开浏览器
   - 修改 `package.json` 文件中的 `"serve": "vue-cli-service serve --open",`
2. 关闭eslint语法校验
    - 修改 `vue.config.js` 文件中加入 `lintOnSave: false,`
3. src文件夹简写方法 配置别名为@ 在 `jsconfig.json` 文件中

## 前端路由分析
#### 路由组件
- Home Search Login Register
#### 非路由组件
- Header(首页 搜索页) Footer(在登录页面没有)