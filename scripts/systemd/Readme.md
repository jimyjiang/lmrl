添加文件：/etc/systemd/system/lmrl.servic

启用并启动服务：

```bash
   sudo systemctl daemon-reload
   sudo systemctl enable lmrl
   sudo systemctl start lmrl
```

查看日志：

```bash
   journalctl -f -u lmrl
```
