# goroutine

1.依据文档图6-1，用中文描述 Reactive 动机

2.使用 go HTTPClient 实现图 6-2 的 Naive Approach

3.为每个 HTTP 请求设计一个 goroutine ，利用 Channel搭建基于消息的异步机制，实现图 6-3

4.对比两种实现，用数据说明 go 异步 REST 服务协作的优势

---------------
### Reactive 动机
如图6-2所示，因为同步方法是顺序处理请求的，所以会有资源的浪费。我们不必要去阻塞线程，它们可能会被用于一些实际的工作。
![](https://github.com/JXONEV/goroutine/raw/master/image/1.png)

### Naive Approach
![](https://github.com/JXONEV/goroutine/raw/master/image/2.png)

### Goroutine
![](https://github.com/JXONEV/goroutine/raw/master/image/3.png)

### 对比
![](https://github.com/JXONEV/goroutine/raw/master/image/4.png)

从执行时间来看，异步方法的花费要小于同步方法，说明它能够有效地利用资源。
