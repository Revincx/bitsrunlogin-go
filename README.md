# BitSrunLoginGo

[![Lisense](https://img.shields.io/github/license/Mmx233/BitSrunLoginGo)](https://github.com/Mmx233/BitSrunLoginGo/blob/main/LICENSE)
[![Release](https://img.shields.io/github/v/release/Mmx233/BitSrunLoginGo?color=blueviolet&include_prereleases)](https://github.com/Mmx233/BitSrunLoginGo/releases)
[![GoReport](https://goreportcard.com/badge/github.com/Mmx233/BitSrunLoginGo)](https://goreportcard.com/report/github.com/Mmx233/BitSrunLoginGo)

深澜校园网登录脚本Go语言版。GO语言可以直接交叉编译出mips架构可执行程序（路由器）（主流平台更不用说了），从而免除安装环境。

> 登录逻辑来自： https://github.com/coffeehat/BIT-srun-login-script

> 对Openwrt更加友好的ipk编译版： [Mmx233/BitSrunLoginGo_Openwrt](https://github.com/Mmx233/BitSrunLoginGo_Openwrt) ，该版本压缩了binary文件，节省闪存空间

## :hammer_and_wrench:构建

**\*运行程序不需要golang环境**

建议安装使用最新版golang

直接编译本系统可执行程序：

```shell
go build
```

交叉编译（Linux）：

```shell
export GOGGC=0
export GOOS=windows #系统
export GOARCH=amd64 #架构
go build
```

交叉编译（Powershell）：

```shell
$env:GOGGC=0
$env:GOOS='linux' #系统
$env:GOARCH='amd64' #架构
go build
```

golang支持的系统与架构请自行查询

## :gear:运行

编译结果为可执行文件，直接启动即可

可以通过添加启动参数`--config`指定配置文件路径，默认为当前目录的`Config.yaml`

支持`json`、`yaml`、`yml`、`toml`、`hcl`、`tfvars`等，仅对`json`和`yaml`进行了优化与测试

```shell
./autoLogin --config=/demo/i.json
```

首次运行将自动生成配置文件

Config.json说明：

```json5
{
  "form": {
    "domain": "www.msftconnecttest.com", //登录地址ip或域名
    "username": "", //账号
    "user_type": "cmcc", //运营商类型，详情看下方
    "password": "" //密码
  },
  "meta": { //登录参数
    "n": "200",
    "type": "1",
    "acid": "5",
    "enc": "srun_bx1"
  },
  "settings": {
    "basic": { //基础设置
      "https": false, //访问校园网API时直接使用http URL
      "skip_cert_verify": false, //是否忽略证书验证
      "interfaces": "", //网卡名称正则（注意JSON转义），如：eth0\\.[2-3]，不为空时为多网卡模式
      "skip_net_check": false, //是否跳过网络检查（仅非守护模式）
      "net_check_url": "https://www.baidu.com/", //网络检查使用的URL
      "timeout": 5 //网络请求超时时间（秒）
    },
    "guardian": { //守护模式
      "enable": false,
      "duration": 300, //网络检查周期（秒）
    }, 
    "daemon": { //后台模式（不建议windows使用）
      "enable": false,
      "path": ".BitSrun", //守护监听文件路径，确保只有单守护运行
    },
    "debug": {
      "enable": false, //开启debug模式，报错将更加详细
      "write_log": false, //写日志文件
      "log_path": "./" //日志文件存放路径
    }
  }
}
```

登录参数从原网页登陆时对`/srun_portal`的请求抓取，抓取时请把浏览器控制台的`preserve log`（保留日志）启用。

运营商类型在原网页会被自动附加在账号后，请把`@`后面的部分填入`user_type`，没有则留空（删掉默认的）

## :jigsaw: 作为module使用

**\*本项目使用了AGPL V3 License，请酌情引用**

示例：

```go
package main

import (
	"github.com/Mmx233/BitSrunLoginGo/v1"
	"github.com/Mmx233/BitSrunLoginGo/v1/transfer"
)

func main() {
	//具体用法请查看struct注释
	if e:=BitSrun.Login(&srunTransfer.Login{
		Https:      false,
		OutPut:    false,
		Debug:  false,
		WriteLog: false,
		CheckNet:  false,
		Transport: nil,
		LoginInfo: srunTransfer.LoginInfo{
			Form: &srunTransfer.LoginForm{
				Domain:   "",
				UserName: "",
				UserType: "",
				PassWord: "",
			},
			Meta: &srunTransfer.LoginMeta{
				N:    "",
				Type: "",
				Acid: "",
				Enc:  "",
			},
		},
	});e!=nil {
		panic(e)
    }
}
```