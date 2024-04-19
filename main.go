package main

import (
	"log"
	"os"
	"os/exec"
)

const (
	RetryTimes = 10000
)

func main() {
	cmdArgs := os.Args[1:]
	if len(cmdArgs) <= 0 {
		log.Fatalf("请输入命令，例：keep-running ollama run qwen:7b")
	}
	cmd, args := cmdArgs[0], cmdArgs[1:]
	for i := 0; i < RetryTimes; i++ {
		c := exec.Command(cmd, args...)
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		log.Printf("运行命令：%s", c)
		if err := c.Start(); err != nil {
			log.Printf("启动失败：%s %s", c, err)
			continue
		}
		if err := c.Wait(); err != nil {
			log.Printf("运行失败：%s %s", c, err)
			continue
		}
		if c.ProcessState.Success() {
			log.Printf("命令成功结束：%s", c)
			break
		}
	}
}
