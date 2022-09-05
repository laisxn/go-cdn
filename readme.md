# gin-cdn

### 
```
cp config_test.ini  config.ini
```

### 服务端监听接受文件
```
go run receive.go 
```

### 客户端传输单个文件
```
go run send.go 
```

### 客户端传输目录下所有文件
```
go run multi_send.go 
```