## install guide

```shell

# install extractwords (从文件夹中提取英文单词)
go install github.com/ooooo-youwillsee/ooooo-tools/cmd/extractwords@latest

# usage: extractwords -f words_file -d words_dir -o output.txt
extractwords -f /path/to/words.file -o output.txt

# install pullimagetodocker (拉取镜像，推送到docker中)
go install github.com/ooooo-youwillsee/ooooo-tools/cmd/pullimagetodocker@latest

# usage: pullimagetodocker -f k8s_yaml_file -d k8s_yaml_dir -u docker.username
pullimagetodocker -f /path/to/k8s.yaml -u youwillsee -o output
```