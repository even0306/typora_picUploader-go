# nextcloudUploader_typora
typora plugin，nextcloudPicUploader  

在学习go语言，刚好没有找到喜欢的typora图片上传插件，便自己边查边学边练，做出来了。练手项目，写的不好

nextcloud使用__File sharing__插件实现直链下载，使用开放api接口上传

#### 配置文件
```
{
    "uploadUrl": "http[s]://[host]/remote.php/dav/files/[user]/[path]/",    //nextcloud的上传地址
    "downloadUrl": "http[s]://[host]/apps/sharingpath/[user]/[path]/",  //nextcloud通过File sharing插件产生的下载地址
    "user": "",         //nextcloud的账号
    "passwd": "",       //nextcloud的密码
    "proxy": ""         //代理
}
```
将config.json文件放在执行文件同目录即可