# keep-running

Keeping your commands running!

使你的命令保持运行状态！

## Installation

你可以使用以下命令安装 `keep-running`，不要忘记安装 Go 环境：

```bash
go install github.com/Lofanmi/keep-running@latest
```

## Usage

如果你用过 ollama 部署模型，你会经常看到这个错误信息：

```
➜  ~ ollama run qwen:7b             
pulling manifest 
pulling 87f26aae09c7...  42% ▕███████████████                     ▏ 1.9 GB/4.5 GB                  
Error: max retries exceeded: Get "https://registry.ollama.ai/v2/library/qwen/blobs/sha256:87f26aae09c7f052de93ff98a2282f05822cc6de4af1a2a159c5bd1acbd10ec4": read tcp 192.168.10.101:38410->172.67.182.229:443: read: connection reset by peer
➜  ~ ollama run qwen:7b
pulling manifest 
pulling 87f26aae09c7...  70% ▕█████████████████████████           ▏ 3.2 GB/4.5 GB  2.5 MB/s    9m5s
Error: max retries exceeded: Get "https://dd20bb891979d25aebc8bec07b2b3bbc.r2.cloudflarestorage.com/ollama/docker/registry/v2/blobs/sha256/87/87f26aae09c7f052de93ff98a2282f05822cc6de4af1a2a159c5bd1acbd10ec4/data?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=66040c77ac1b787c3af820529859349a%!F(MISSING)20240419%!F(MISSING)auto%!F(MISSING)s3%!F(MISSING)aws4_request&X-Amz-Date=20240419T071250Z&X-Amz-Expires=1200&X-Amz-SignedHeaders=host&X-Amz-Signature=129289c17cae6380401d07941d75d089d7c13a176d65ffc89be8d3938f4ca3d2": read tcp 192.168.10.101:44436->104.18.8.90:443: read: connection reset by peer
```

出现这个错误信息，是因为某些不可描述的原因引起的网络不稳定，导致下载镜像失败。

这时候你可以使用 `keep-running` 命令来保持你的命令运行：

```bash
keep-running ollama run qwen:7b
```

这样，当命令失败退出时，`keep-running` 会自动重新运行你的命令（程序会自动重试 10000 次），直到你的命令成功为止。

## TODO

1. 支持自定义重试次数
2. 支持自定义重试间隔

## License

MIT
