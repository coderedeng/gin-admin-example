## 运行步骤

1. 克隆代码库：

```plaintext
git clone git clone https://github.com/coderedeng/gin-admin-example.git
```

2. 进入项目目录：

```plaintext
cd gin-admin-example-main
```

3. 安装依赖：

```plaintext
go mod tidy
```

4. 配置数据库：

- 在项目根目录下找到 `sql/gpa.sql` 文件，初始化数据库结构。

5. 运行项目：

```plaintext
go run main.go
```

6. 访问接口管理界面：

- 生成swagger文档
-     swag init
  
- 在浏览器中输入 

-     swagger后端文档地址:http://127.0.0.1:8888/swagger/index.html
      knife4g后端文档地址:http://127.0.0.1:8888/doc/index
