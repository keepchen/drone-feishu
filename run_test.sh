#!/bin/bash

export DRONE_SYSTEM_HOST=http://127.0.0.1
export DRONE_REPO_NAMESPACE=game
export DRONE_REPO_NAME=game-web-frontent
export DRONE_REPO=game/game-web-frontent
export DRONE_REPO_BRANCH=dev
export DRONE_COMMIT_SHA=abcdefghijklmn
export DRONE_COMMIT_MESSAGE="这是一条很长很长很长的测试消息，大概有25个字符吧"
export DRONE_COMMIT_LINK=http://127.0.0.1/game/game-web-frontent/-/commit/abcdefghijklmn
export DRONE_COMMIT_AUTHOR=keepchen
export DRONE_BUILD_STATUS=success # failure
export DRONE_BUILD_LINK=http://127.0.0.1/game/game-web-frontent/1
export PLUGIN_TOKEN= # 飞书的webhook token值
export PLUGIN_SECRET= # 飞书的签名校验secret
export PLUGIN_CARD_TITLE= # 卡片消息标题
export PLUGIN_SUCCESS_IMG_KEY= # 构建成功图片
export PLUGIN_FAILURE_IMG_KEY= # 构建失败图片
export PLUGIN_POWERED_BY_IMG_KEY= # 版权logo
export PLUGIN_POWERED_BY_IMG_ALT= # 版权logo的alt提示文字

# go build -o drone-feishu
# ./drone-feishu

go run .