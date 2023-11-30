```go
package main

import (
	"fmt"
	"os"
	"time"
)

// timestampWriter 是实现了 io.Writer 接口的结构体
type timestampWriter struct {
	timestamp time.Time // 记录时间戳
	file      *os.File  // 文件对象
}

// Write 方法实现了 io.Writer 接口，将带有时间戳的日志信息写入文件
func (t *timestampWriter) Write(p []byte) (int, error) {
	// 将时间戳格式化为字符串
	timestampStr := t.timestamp.Format("2006-01-02 15:04:05")
	// 构建带时间戳的日志信息
	newP := []byte(fmt.Sprintf("[%s] %s", timestampStr, p))
	// 写入日志信息到文件
	_, err := t.file.Write(newP)
	if err != nil {
		return 0, err
	}
	// 返回写入的字节数和 nil 错误
	n := len(p)
	return n, nil
}

func main() {
	filePath := "d:/test_1.txt"
	// 创建或打开一个文件用于写入
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close() // 延迟关闭文件

	// 创建一个 timestampWriter 实例，用于写入日志信息到文件中
	logWriter := &timestampWriter{timestamp: time.Now(), file: file}

	str := "输入密码"
	// 写入带有格式的日志信息到文件中
	fmt.Fprintf(logWriter, "用户登录:%s\n", str)

	// 模拟延迟
	time.Sleep(2 * time.Second)
	// 再次写入日志信息到文件中，不使用 fmt.Fprintf，而是使用 fmt.Fprintln
	fmt.Fprintln(logWriter, "用户执行操作A\n")

	// 模拟延迟
	time.Sleep(1 * time.Second)
	// 继续写入日志信息到文件中
	fmt.Fprintln(logWriter, "用户执行操作B")
}
```

