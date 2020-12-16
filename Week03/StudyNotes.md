### 第三课 concurrency笔记



#### Gorountine

1. 操作系统并行和并发的概念，线程，进程，内核态

2. main也是作为 goroutine运行的。

3. GPM 的逻辑要弄明白；一个P 和M 可能对于多个G 

4. 并发不是并行；

5. 代码超时控制

   ```
   unc main() {
     ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
     defer cancel()
     
     ch := make(chan result)
     go func() {
       record, err := search(ctx)
       ch <- result{record, err}
     }()
     
     select {
     case <- ctx.Done():
       // timeout
     case res := <-ch:
       // get result
     }
   }
   ```

6. 不鼓励select{}的方式，阻塞main 线程，来提供不退出的服务。

7. log.Fatal(err)的方式阻塞,不要在野生地方使用，在main 里面；

8. 鼓励采用chan 来监听； 

9. go 出去的 goroutine 要管住他的生命周期；不鼓励野生的 go一个协程出去；


#### Memory model

1. 为了保证共享内存的正确性（可见性、有序性、原子性），内存模型定义了共享内存系统中多线程程序读写操作行为的规范。

   通过这些规则来规范对内存的读写操作，从而保证指令执行的正确性。它与处理器有关、与缓存有关、与并发有关、与编译器也有关。

   它解决了 CPU 多级缓存、指令重排等导致的内存访问问题，保证了并发场景下的一致性、原子性和有序性。

2. 内存模型，基本就是原子的，可见性。另外有：happensBefor。atomic.Add()

3. Happens before：如果A happens before B，那么A的执行结果对B可见（**并不一定表示A比B先执行，如果A与B执行的顺序对结果没有影响是可以重排序的**）

4. 在单线程环境下，所有的表达式，按照代码中的先后顺序，具有Happens Before关系

5. 按时

#### Package sync

1. 包的异步，原子CAS等

2. 传统的线程模型(通常在编写 Java、C++ 和Python 程序时使用)程序员在线程之间通信需要使用共享内存。通常，共享数据结构由锁保护，线程将争用这些锁来访问数据。在某些情况下，通过使用线程安全的数据结构(如Python的Queue)，这会变得更容易。

3. Do not communicate by sharing memory; instead, share memory by communicating.

4. 一个 in 的chan  一个 out 的chan , 非常清真。

5. race detector 竞赛检测器。

   ```
   go build -race 8.go #查问题
   go test -race 
   ```

   Go tools compain -s

6. 秀了半天data race 在jerry  和 ben 的 两个并行互相调用的例子。

#### chan

1. 按时

#### Package context

1. 可以做到级联的取消，级联的超时，显示的传递；

#### Reference(学习资料)

Cpu的缓存对go 程序的影响

```
https://mp.weixin.qq.com/s/vnm9yztpfYA4w-IM6XqyIA
```

Effective Go 这本书。

```
https://learnku.com/docs/effective-go/2020/function/6242
```

