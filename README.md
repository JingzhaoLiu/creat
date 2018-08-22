## 项目生成脚手架

- 使用环境mac

生成一个空的go项目组，带go插件的各种依赖，
在vscode中使用只需要安装一个go插件就可以愉快的编码了

- dirname(自定义文件名)
--bin 
--- go插件依赖
--pkg 
--- go插件依赖
--src
---main.go  



## 使用



```  shell
//rule

./copy <dir name> <path>

// example

cd bin
./copy calcpro /Users/martin/Desktop/

```



## 字符串截取
``` go
split := strings.Split(str, "Desktop/")
i:= strings.Index(str, "liu/")
fmt.Println(split[1])
fmt.Println(str[i:])

```