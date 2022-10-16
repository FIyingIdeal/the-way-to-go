# the-way-to-go

## Go 环境安装

- [Go代理设置](https://goproxy.io/zh/docs/getting-started.html)

## Go语言学习教程

- [《The Way to Go》中文译本](https://learnku.com/docs/the-way-to-go)


## Go 项目的包结构

- [Go中的目录结构](https://github.com/zhouhaibing089/Blog/issues/5)


## Go 语言中的类型

- 基础类型

- 复合类型

从方法传递角度来说，会分为值类型和引用类型：
- 值类型
    - 结构体

- 引用类型
    - 切片
    -  

#### 结构体类型

Go语言中的结构体类型类似于Java语言中的 `POJO`（Plain Ordinary Java Object，简单Java对象，其中有一些属性及其 `getter/setter` 方法的类，没有业务逻辑），包含一些不重名的字段（Field）。定义方式如下：

```go
type identifier struct {
    filed1 type1
    field2 type2
}
```