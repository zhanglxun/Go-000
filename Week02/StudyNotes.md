### 第二课 error 学习笔记



#### Error vs Exception

1. error 的定义

   ```
   type error interface{
   	Error() string 
   }
   ```

2. Panic ：

   1. 在程序启动的时候，如果有强依赖的服务出现故障时 `panic`退出。
   2. 在程序启动的时候，如果发现有配置明显不符合要求， 可以 `panic` 退出（防御编程）。
   3. 在程序入口处，例如 `gin` 中间件需要使用 `recovery` 预防 `panic` 程序退出。
   4. 在程序中我们应该避免使用野生的 `goroutine`。
      1. 如果是在请求中需要执行异步任务，应该使用异步 `worker` ，消息通知的方式进行处理，避免请求量大时大量 `goroutine` 创建。
      2. 如果需要使用 `goroutine` 时，应该使用同一的 `Go` 函数进行创建，这个函数中会进行 `recovery` ，避免因为野生 `goroutine` panic 导致主进程退出。

3. 基础库中大量定义

4. 

5. 

#### Error Type



#### Go 1.13 Erros



#### Go 2 Error Inspection



#### 课堂作业

1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？



#### 学习资料

Vim-go 的插件，可以在vim 中编写和工作

```
https://github.com/fatih/vim-go

内存模型
https://mp.weixin.qq.com/s/eDd212DhjIRGpytBkgfzAg
```

