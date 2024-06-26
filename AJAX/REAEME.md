# AJAX

### 请求报文

- 行
  - GET /s?ie=utf-8  HTTP/1.1
- 头
  - Host:stguigu.com
  - Cookie: name=guigu
  - Content-type: application/x-www-form-urlencode
- 空行
- 体  GET的请求体是空  POST可以不为空
  - username=admin&password=passwd

### 相应报文

- 行
  - HTTP/1.1  200/OK
- 头
  - Content-type: text/html;charset=utf-8
  - COntent-length: 2048
  - Content-encoding: gzip
- 空行
- 体
  - `<html><head><body><h1>Hello</h1></body></head></html?`
