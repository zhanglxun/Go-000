### 第二课 error 学习笔记



#### Error vs Exception

1. error 的定义

   ```
   type error interface{
   	Error() string 
   }
   ```

2. 基础库中大量定义

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

