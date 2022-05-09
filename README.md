# golang-web-application

## Go 内置的标准库




### 如何创建Web Server

Web服务器就是用来处理Web请求 (**Handle** Web Request)


```go
func ListenAndServe(addr string, handler Handler) error {
	server := &Server{Addr: addr, Handler: handler}
	return server.ListenAndServe()
}
```


```go
func (srv *Server) ListenAndServe() error {
    ...
	ln, err := net.Listen("tcp", addr)
    ...
	return srv.Serve(ln)
}
```

```go
func (srv *Server) Serve(l net.Listener) error {
	...
	for {
		rw, err := l.Accept()
		...
		connCtx := ctx
		...
		c := srv.newConn(rw)
		c.setState(c.rwc, StateNew, runHooks) // before Serve can return
		go c.serve(connCtx)
	}
}
```

```go

// Serve a new connection.
func (c *conn) serve(ctx context.Context) {
    ...
    for{
        ...
		serverHandler{c.server}.ServeHTTP(w, w.req)
        ...
	}
}
```

*http.ListenAndServer()*

通过源码分析，我们得出如下调用关系图，最后的ServeHttp()就是我们常用的HTTP Handler被调用的位置

```
Server 
    --> ListenAndServe()
        -->Serve()
            --> newConn() --> conn {Server :...}
            conn
                -->serve()
                    -->ServeHttp()

```


http.handler是一个接口，其中定义了ServeHttp()方法，其中ServeHttp有两个参数

- HTTPResponesWriter
- 指向Request的Struct的指针

- http.Handle
- http.HandleFunc



## [julienschmidt/httprouter](https://github.com/julienschmidt/httprouter)








## 数据库


