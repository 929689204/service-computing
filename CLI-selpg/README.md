# CLI-selpg

## 项目描述

> 使用 Go 语言开发的 Linux 命令行实用程序Selpg。
> 
> Selpg从标准输入或从作为命令行参数给出的文件名读取文本输入。它允许用户指定来自该输入并随后将被输出的页面范围,然后输出到标准输出或是文件中。
>
>程序开发参考于[开发 Linux 命令行实用程序](https://www.ibm.com/developerworks/cn/linux/shell/clutil/index.html) 。

## 代码设计
### 引入需要的包
```go
import(
	"bufio"
	"fmt"
	"os"
	"strconv"
)
```

### 数据结构

```go
type selpgArgs struct{
    startPage int
    endPage int
    inFilename string
    pageLen int
    pageType int
    printDest string
}
```

其中，pageLen 表示每页的行数，可以被 “-l” 指令修改；pageType 表示页的类型，l 为确定页行数的文本，f 为页数由ASCII码确定的换页字符定界的文本；

### 主函数

```go
func main(){
	var sa selpgArgs
	ac := len(os.Args)
	av := os.Args
	progname = os.Args[0]
	
	sa.startPage = -1
	sa.endPage = -1
	sa.inFilename = ""
	sa.pageLen = 72
	sa.pageType = 'l'
	sa.printDest = ""

	processArgs(ac, av, &sa)
	processInput(sa)
}
```

其中，processArgs函数用来判断指令的具体要求和传递相应参数，processInput函数用来按指令操作输入输出流

## 测试结果



1. `$ selpg`，错误用法时提示USAGE

```
   $ ./selpg :
./selpg: not enough arguments

USAGE: ./selpg -sstart_page -eend_page [ -f | -llines_per_page ][ -ddest ] [ in_filename ]

```

2. `$ selpg -s1 -e2 -l10 input.txt`，将input.txt第一页和第二页输出，每一页10行，终端显示为：

```
   $ ./selpg -s1 -e2 -l10 input.txt
input file line 1
input file line 2
input file line 3
input file line 4
input file line 5
input file line 6
input file line 7
input file line 8
input file line 9
input file line 10
input file line 11
input file line 12
input file line 13
input file line 14
input file line 15
input file line 16
input file line 17
input file line 18
input file line 19
input file line 20

```

3. `$ selpg -s1 -e2 < input.txt`，selpg 读取标准输入，标准输入被 shell / 内核重定向为来自 “input.txt” 而不是显式命名的文件名参数；
```
$ ./selpg -s1 -e2 < input.txt
input file line 1
input file line 2
input file line 3
input file line 4
input file line 5
input file line 6
input file line 7
input file line 8
input file line 9
input file line 10
···
input file line 50
```



4. `selpg -s3 -e5 -l6 input.txt >output.txt`，selpg 将第 3 页到第 5 页写至标准输出，每页6行；标准输出被 shell／内核重定向至 “output.txt”，终端没有输出，使用`cat output.txt`命令在终端显示 output.txt 内容：

```
$ ./selpg -s3 -e5 -l6 input.txt >output.txt
$ cat output.txt 
input file line 13
input file line 14
input file line 15
input file line 16
input file line 17
input file line 18
input file line 19
input file line 20
input file line 21
input file line 22
input file line 23
input file line 24
input file line 25
input file line 26
input file line 27
input file line 28
input file line 29
input file line 30

```

5. `selpg -s10 -e20 input.txt 2>error.txt`，不符合标准的信息将被输出至error.txt，终端不显示错误信息，使用`cat error.txt`命令在终端显示error内容：

```
   $ selpg -s10 -e20 -l3 input.txt 2>error.txt
$ cat error.txt 
bash: selpg: 未找到命令...
$ ./selpg -s10 -e20 -l3 input.txt 2>error.txt
input file line 28
input file line 29
input file line 30
input file line 31
input file line 32
input file line 33
input file line 34
input file line 35
input file line 36
input file line 37
input file line 38
input file line 39
input file line 40
input file line 41
input file line 42
input file line 43
input file line 44
input file line 45
input file line 46
input file line 47
input file line 48
input file line 49
input file line 50
$ cat error.txt
./selpg: end_page (20) greater than total pages (17), less output than expected


```

6. `selpg -s1 -e2 -f input.txt`，假定页由换页符定界，第 1 页到第 2 页被写至 selpg 的标准输出；

7. `selpg -s1 -e2 -dxxx input.txt`，输出错误信息：

```
]$ ./selpg -s1 -e2 -dxxx input.txt
./selpg: could not open pipe to ""
```
