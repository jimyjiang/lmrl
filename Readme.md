启动一个 3001 端口的服务器，通过 nginx 反向代理，可以通过访问http://47.47.101.147.164:8080/lmrl来访问该站点
本站点首页展示灵命日粮的列表，点击后收听音频。
从服务器指定目录读取音频 mp3 文件的属性，动态展示列表，列表展示最近的 60 个音频文件。
配置，见 nginx 配置
编译： make build
部署： make deploy

mp3 文件的下载与上传见： /Users/jimmy.jiang/doc/bairex/aliyun/nginx.md, 灵命日粮章节。
