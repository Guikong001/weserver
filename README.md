# 微信公众号服务器
> 用以微信公众号的后端，提供登录验证功能
> fork自Justsong

## 功能
+ [x] Access Token 自动刷新 & 提供外部访问接口
+ [x] 自定义菜单（需要你的公众号有这个权限）
+ [x] 登录验证
+ [x] AI 回复
+ [ ] 自定义回复

## 展示
![demo1](https://user-images.githubusercontent.com/39998050/200124147-3338a2eb-8193-4068-ae6f-276cfe16a708.png)
![demo2](https://user-images.githubusercontent.com/39998050/200124177-78636b4c-0aac-4860-a138-68f3d92477b9.png)

## 部署
### 手动部署
1. 从 [GitHub Releases](https://github.com/songquanpeng/wechat-server/releases/latest) 下载可执行文件或者从源码编译：
   ```shell
   git clone https://github.com/songquanpeng/wechat-server.git
   go mod download
   go build -ldflags "-s -w" -o wechat-server
   ````
2. 运行：
   ```shell
   chmod u+x wechat-server
   ./wechat-server --port 3000 --log-dir ./logs
   ```
3. 访问 [http://localhost:3000/](http://localhost:3000/) 并登录。初始账号用户名为 `root`，密码为 `123456`。

更加详细的部署教程[参见此处](https://iamazing.cn/page/how-to-deploy-a-website)。

### 基于 Docker 进行部署
执行：`docker run -d --restart always -p 3000:3000 -v /home/ubuntu/data/wechat-server:/data justsong/wechat-server`

数据将会保存在宿主机的 `/home/ubuntu/data/wechat-server` 目录。

## 配置
1. 从 [GitHub Releases](https://github.com/songquanpeng/wechat-server/releases/latest) 下载可执行文件。
2. 系统本身开箱即用，有一些环境变量可供配置：
   1. `REDIS_CONN_STRING`: 设置之后，将启用 Redis。
      + 例如：`REDIS_CONN_STRING=redis://default:redispw@localhost:49153`
   2. `SESSION_SECRET`:设置之后，将使用给定会话密钥。
      + 例如：`SESSION_SECRET=random_string`
   3. `SQL_DSN`: 设置之后，将使用目标数据库而非 SQLite。
      + 例如：`SQL_DSN=root:123456@tcp(localhost:3306)/gofile`
3. 运行: 
   1. `chmod u+x wechat-server`
   2. `./wechat-server --port 3000`
4. 初始账户用户名为 `root`，密码为 `123456`，记得登录后立刻修改密码。
5. 前往[微信公众号配置页面 -> 设置与开发 -> 基本配置](https://mp.weixin.qq.com/)获取 AppID 和 AppSecret，并在我们的配置页面填入上述信息，另外还需要配置 IP 白名单，按照页面上的提示完成即可。
6. 前往[微信公众号配置页面 -> 设置与开发 -> 基本配置](https://mp.weixin.qq.com/)填写以下配置：
   1. `URL` 填：`https://<your.domain>/api/wechat`
   2. `Token` 首先在我们的配置页面随便填写一个 Token，然后在微信公众号的配置页面填入同一个 Token 即可。
   3. `EncodingAESKey` 点随机生成，然后在我们的配置页面填入该值。
   4. 消息加解密方式选择明文模式。
7. 之后保存设置并启用设置。
8. 当前版本需要重启服务才能应用配置信息，因此请重启服务。

## API
### 获取 Access Token
1. 请求方法：`GET`
2. URL：`/api/wechat/access_token`
3. 无参数，但是需要设置 HTTP 头部：`Authorization: <token>`

### 通过验证码查询用户 ID
1. 请求方法：`GET`
2. URL：`/api/wechat/user?code=<code>`
3. 需要设置 HTTP 头部：`Authorization: <token>`

### 注意
需要将 `<token>` 和 `<code>` 替换为实际的内容。
