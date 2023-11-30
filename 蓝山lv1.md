```go
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now().UnixMicro()
	var str string
	filePath := "d:/test_1.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer file.Close()
	str = "hello world "
	for i := 1; i <= 20000; i++ {
		file.Write([]byte(str))
	}
	end := time.Now().UnixMicro()
	fmt.Println("耗时", end-start)
}

```

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now().UnixMicro()
	filePath := "d:/test.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("创建文件失败，错误是=", err)
		return
	}
	defer file.Close()
	str := "hello world"
	write := bufio.NewWriter(file)
	for i := 1; i <= 20000; i++ {
		write.WriteString(str)
	}
	write.Flush()
	end := time.Now().UnixMicro()
	fmt.Println("耗时=", end-start)
}

```

