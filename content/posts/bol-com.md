+++ 
draft = false
date = 2026-06-07T23:52:03+02:00
title = "Bol.com Price Discount Alert"
description = "Get Price Discount Notification for Specific Product on bol.com"
slug = ""
authors = []
tags = ["git", "go", "shell", "bash", "python", "mail", "telegram"]
categories = ["code"]
externalLink = ""
series = []
+++


# Get Price Discount Notification for Specific Product on bol.com
**Source code [https://github.com/mozkaya1/bol.git](https://github.com/mozkaya1/bol.git)** 

> get bol.com link from product page .

![bol.com](/bol.com.png) 

### initial setting on shell code (./bol.sh)
> adjust settings on ./bol.sh
```bash
set "url" link for product on bol.com 
set "setPrice" for threshold price to be notified below  
set "cooldown"  check interval (seconds)
```
![initial setting](/initial-setting.png) 


### main GO code to get price on bol.com

> help context 
```bash
go run main.go -h
```
![go code help](/go-help.png) 

> run raw code  -- only gives output json of the product such as status,price,url ..
```bash
go run main.go

```
```bash
{"url":"https://www.bol.com/nl/nl/p/pd-tang-cv-waterpomptang-29-cm-isolatie-grip-12-profi-stalen-tang-met-softgrip-handvat/9300000160855987","price":19,"status":200,"discount":0}
```

### Notification -- (./bol.sh)
> I have used my own Telegram bot. You can use your own notification system as mail,notify so on ..

![ncode](/notification-code.png)


> Preview - Telegram

![telegram](/notification.png) 

### output log (./bol.log)

```bash 
[2026-06-07 15:24:19] -- Price dropped - % 37 --> https://www.bol.com/nl/nl/p/pd-tang-cv-waterpomptang-29-cm-isolatie-grip-12-profi-stalen-tang-met-softgrip-handvat/9300000160855987 -- 
[2026-06-07 15:24:20] -- No Discount -- {"url":"https://www.bol.com/nl/nl/p/pd-tang-cv-waterpomptang-29-cm-isolatie-grip-12-profi-stalen-tang-met-softgrip-handvat/9300000160855987","price":19,"status":200,"discount":0}
[2026-06-07 15:24:30] -- No Discount -- {"url":"https://www.bol.com/nl/nl/p/pd-tang-cv-waterpomptang-29-cm-isolatie-grip-12-profi-stalen-tang-met-softgrip-handvat/9300000160855987","price":19,"status":200,"discount":0}
[2026-06-07 15:24:41] -- No Discount -- {"url":"https://www.bol.com/nl/nl/p/pd-tang-cv-waterpomptang-29-cm-isolatie-grip-12-profi-stalen-tang-met-softgrip-handvat/9300000160855987","price":19,"status":200,"discount":0}
```

