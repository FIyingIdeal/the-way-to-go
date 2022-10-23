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
    - map
    - channel

#### 数组类型

#### 切片类型 - slice

- **简介**
    slice 是一个引用类型，其底层是基于数组实现的，可以理解为可变长度的数组，与java中的 `ArrrayList` 类似。

- **声明方式**
    `var identifier []type`
    slice 的声明方式与数组特别相似，只是省略了初始长度。以声明一个int类型的切片和声明一个长度为 5 的数组对照来看：

    ```GO
    // int 类型的 slice 声明
    var s []int
    // int 类型长度为5的数组声明
    var a [5]int
    ```

- **初始化**
    1. `var slice1 []type = arr1[start:end]` : 由数组 arr1 从 start 索引到 end-1 索引之间的元素构成的子集（切分数组，start:end 被称为切片表达式）；
    2. `slice2 := &arr2` : 与 `var slice2 []type = arr2[:]` 和 `var slice2 []type = arr2[0:len(arr2)]` 一样，等于完整的 arr2 数组；
    3. `slice3 := make([]type, len, cap)` : 当相关数组还没有定义时，可以使用 `make()` 函数来创建一个切片，同时会创建好数组。其中 cap 是可选参数，表示底层数组的长度，当 cap 被省略，只保留 len 时，即 `slice3 := make([]type, len)` 时，len 既是数组的长度也是切片的初始长度；
        ```GO
        // 以下两种方法生成相同的切片
        make([]int, 50, 100)
        new([100]int)[0:50]
        ```


#### 结构体类型

- **简介** 
    Go语言中的结构体类型类似于Java语言中的 `POJO`（Plain Ordinary Java Object，简单Java对象，其中有一些属性及其 `getter/setter` 方法的类，没有业务逻辑），包含一些不重名的字段（Field）
    
- **声明方式**    
    一般的结构体声明方式如下所示：

    ```go
    type identifier struct {
        // 字段名 类型
        filed1 type1
        field2 type2
    }
    ```
    此外，结构体中可以包含一个或多个匿名（或内嵌）字段，即这些字段没有显示的名字，只有类型，此时类型就是字段的名字，也因此一个结构体中对于每一种数据类型只能有一个匿名字段。

    ```go
    package main

    import "fmt"

    type Student struct {
        name string
        // int 是一个匿名字段，类型就是字段名
        int
    }

    func main() {
        s := Student{"zhangsan", 22}
        // 类型就是字段名，因此可以通过 s.int 访问
        age := s.int
        fmt.Println("age is :", age)
    }
    ```

- **内嵌结构体**
    结构体也可以作为一个匿名字段放在另一个结构体里，此时外层结构体可以直接访问内层结构体的字段。如下例所示：
    ```go
    package main

    import "fmt"

    type Inner struct {
        a int
        b int
    }

    type Outter struct {
        c int
        Inner
    }

    func main() {
        outter := Outter{3, Inner{1, 2}}
        // 通过 outter 可以直接访问匿名结构体 Inner 的 a 和 b 字段
        a := outter.a
        b := outter.b
        c := outter.c
        fmt.Printf("a is %d, b is %d, c is %d", a, b, c)
    }
    ```
    这类似面向对象语言中的“继承”。但需要考虑一个问题：继承关系里的字段名重复时该怎么处理？尤其是当一个结构体内部包含的两个匿名结构体中出现重复字段（即同名字段在同一个层级）时会发生什么？
    > 1. 当不同层级的字段重名时，外层的名字会覆盖内层的名字，但两者内存空间都保留，这提供了一种重载字段或方法的方式；
    > 2. 如果相同的名字在同一级别出现了两次，如果这个名字被程序使用了，将会引发一个错误（不使用没关系）。没有办法来解决这种问题引起的二义性，必须由程序员自己修正。
    ```go
    package main

    import "fmt"

    type A struct{ a int }
    type B struct{ a, b int }

    type C struct {
        A
        B
        b float32
    }

    func main() {
        c := C{A{11}, B{21, 22}, 3.1}
        // 报错：ambiguous selector c.acompiler
        // 不知道是 c.A.a 还是 c.B.a
        //aa := c.a

        // 可以通过加匿名字段名的方式明确指定要用哪个结构体内的字段
        aa := c.A.a
        // 11
        fmt.Println("aa is ", aa)

        ba := c.B.a
        // 21
        fmt.Println("ba is ", ba)

        // 外层结构体C中的字段b覆盖了内层匿名结构体B中的字段b
        cb := c.b
        // 3.1
        fmt.Println("bb is ", cb)

        bb := c.B.b
        // 22
        fmt.Println("bb is ", bb)
    }
    ```

- **方法**


- **String()**
与 Java 中的 toString() 方法一样，格式化输出内容。
    | 格式化标识符 | 说明 |
    |---|---|
    | %T | 输出类型的完全规格
    | %#v | 输出实例的完整输出，包括它的字段
    | %p | 输出指针变量中的值，即被指向变量的地址 |




#### 指针类型

- **简介**
    理解指针的相关文章：
    - [800 字彻底理解 Go 指针](https://learnku.com/go/t/35168)
    - [C - 指针](https://www.jianshu.com/p/63f3bfb58687)

- **声明方式**
    `var intP *int`
    `intP1 := &verb`
    当一个指针被定义后没有分配到任何变量是，它的默认值是 nil。

- **反引用**
    在指针类型前面加上 * 号（前缀）来获取指针所指向的内容，这里的 * 号是一个类型更改器。使用一个指针引用一个值被称为间接引用。 
    ```GO
    // 指针变量 intP
    var intP *int
    // 取指针变量 intP 中的值
    var a int = *intP
    ```

- **空间占用**
在 32 位机器上占用 4 个字节，在 64 位机器上占用 8 个字节，并且与它所指向的值的大小无关

- **限制**
    - Go语言中限制不能得到一个文字或常亮的地址；
        ```go
        const i = 5
        ptr := &i //error: cannot take the address of i
        ptr2 := &10 //error: cannot take the address of 10
        ```
    - 指针运算 `pointer+2` 在Go语言中是不被允许的；
    - 对一个空指针的反向引用是不合法的：
        ```go
        // 空指针
        var p *int = nil
        // 直接给空指针赋值（*p就称为反向引用）会使程序崩溃
        *p = 0
        ```



| 类型名称 | 值 OR 引用类型 | 初始值 | 长度 |
|---|---|---|---|
| 切片 - slice | 引用类型 | nil | 0 |

## new() 与 make() 的区别

二者看起来没什么区别，都是在对上分配内存，但它们的行为不同，适用于不同的类型。

- new(T) 为每一个新的类型 T 分配一片内存，初始化为0，并且返回类型为 *T 的内存地址：这种方法返回一个指向类型为 T，值为0的地址的指针。它适用于**值类型**，如 **数组** 和 **结构体**，相当于 &T{}；

- make(T) 返回一个类型为 T 的初始值，它只适用于 3 种内建的引用类型：切片、map 和 channel；