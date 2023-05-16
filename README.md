## install guide

```shell

# install extractwords (从文件夹中提取英文单词)
go install github.com/ooooo-youwillsee/ooooo-tools/cmd/extractwords@latest

# usage: extractwords -f words_file -d words_dir -o output.txt
extractwords -f /path/to/words.file -o output.txt

# install syncimage (拉取镜像，推送到docker中)
go install github.com/ooooo-youwillsee/ooooo-tools/cmd/syncimage@latest

# usage: 
# syncimage remote -f /Users/ooooo/Downloads/tekton.yaml -r docker.io/youwillsee -i gcr.io -e https_proxy=http://127.0.0.1:1080 -e http_proxy=http://127.0.0.1:1080 -v
# syncimage local -f /Users/ooooo/Downloads/tekton.yaml -i gcr.io -e https_proxy=http://127.0.0.1:1080 -e http_proxy=http://127.0.0.1:1080 -v

syncimage -f /path/to/k8s.yaml -u youwillsee -v
```