# drone-feishu
FeiShu robot notification for drone ci plugin  

> Card Style inspired [guoxudongdocker/drone-dingtalk](https://hub.docker.com/r/guoxudongdocker/drone-dingtalk)  

#### QuickStart  

- docker
> [docker hub](https://hub.docker.com/r/keepchen/drone-feishu)
```shell
docker run keepchen/drone-feishu -e PLUGIN_TOKEN=your feishu webhook token \
                                 -e PLUGIN_SECRET=your feishu webhook sign secret
```  
- go run  
```shell  
# Tips: configure your own environment variables in `run_test.sh` first
sh run_test.sh 
```

#### Screenshot  

- success

![success](./card_success.png)  

- failure  

![failure](./card_failure.png)  

#### Test  

> According to your actual situation, set the environment variables in this run_test.sh file, and then run it.

#### Pipeline Settings (.drone.yml)  
```yaml
- name: feishu notification
  image: keepchen/drone-feishu 
  settings:
    token: # (required)
      from_secret: feishu_token
    secret: # (required)
      from_secret: feishu_secret
    card_title: ""
    success_img_key: ""
    failure_img_key: ""
    powered_by_img_key: ""
    powered_by_img_alt: ""
  commands:
    - env
  when:
    status: [failure, success]
```  
#### Environment
| variable | value | desc |
| ---- | ----| ---- |
| PLUGIN_TOKEN | xxxxxx-xxx-xxxx-xxxx-xxxxxxxxxxxx | 飞书的webhook token值 |
| PLUGIN_SECRET | xxxxxxxxxx | 飞书的签名校验secret |
| PLUGIN_CARD_TITLE | | 卡片消息标题 |
| PLUGIN_SUCCESS_IMG_KEY | | 构建成功图片 |
| PLUGIN_FAILURE_IMG_KEY | | 构建失败图片 |
| PLUGIN_POWERED_BY_IMG_KEY | | 版权logo |
| PLUGIN_POWERED_BY_IMG_ALT | | 版权logo的alt提示文字 |

- how to get my pictures `img_key`?  
see [doc](https://open.feishu.cn/document/ukTMukTMukTM/uUjNwUjL1YDM14SN2ATN)  
- what's the `token` & `secret`?  
![robot](./feishu-robot.png)

#### Finally  
Enjoy it. :)
