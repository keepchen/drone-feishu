package main

import (
	"log"
	"os"
)

type DroneEnvironment struct {
	DroneSystemHost    string //drone系统地址
	DroneRepoNamespace string //git仓库所属分组
	DroneRepoName      string //git仓库名称
	DroneRepo          string //git仓库名称，等价于 DroneRepoNamespace + "/" + DroneRepoName
	DroneBranch        string //构建分支
	DroneCommitSha     string //git提交hash摘要值
	DroneCommitMessage string //git提交备注信息
	DroneCommitLink    string //git提交信息网页地址
	DroneCommitAuthor  string //git提交者名称
	DroneBuildStatus   string //构建状态
	DroneBuildLink     string //drone构建信息网页地址

	//以下是通过settings设置的环境变量
	//
	//settings:
	//	token:
	//		from_secret: xxx
	//	card_title: xxx
	//	success_img_key: xxx
	//	failure_img_key: xxx
	//	powered_by_img_key: xxx
	//	powered_by_img_alt: xxx
	PluginSecret          string //授权token
	PluginToken           string //授权token
	PluginCardTitle       string //通知卡片标题
	PluginSuccessImgKey   string //构建成功图片imgKey
	PluginFailureImgKey   string //构建失败图片imgKey
	PluginPoweredByImgKey string //powered by图片imgKey
	PluginPoweredByImgAlt string //powered by图片alt信息
}

// GetEnv 获取环境变量
func (droneEnv DroneEnvironment) GetEnv() DroneEnvironment {
	droneEnv.DroneSystemHost = os.Getenv("DRONE_SYSTEM_HOST")
	droneEnv.DroneRepoNamespace = os.Getenv("DRONE_REPO_NAMESPACE")
	droneEnv.DroneRepoName = os.Getenv("DRONE_REPO_NAME")
	droneEnv.DroneRepo = os.Getenv("DRONE_REPO")
	droneEnv.DroneBranch = os.Getenv("DRONE_BRANCH")
	droneEnv.DroneCommitSha = os.Getenv("DRONE_COMMIT_SHA")
	droneEnv.DroneCommitMessage = os.Getenv("DRONE_COMMIT_MESSAGE")
	droneEnv.DroneCommitLink = os.Getenv("DRONE_COMMIT_LINK")
	droneEnv.DroneCommitAuthor = os.Getenv("DRONE_COMMIT_AUTHOR")
	droneEnv.DroneBuildStatus = os.Getenv("DRONE_BUILD_STATUS")
	droneEnv.DroneBuildLink = os.Getenv("DRONE_BUILD_LINK")
	droneEnv.PluginToken = os.Getenv("PLUGIN_TOKEN")
	droneEnv.PluginSecret = os.Getenv("PLUGIN_SECRET")
	droneEnv.PluginCardTitle = os.Getenv("PLUGIN_CARD_TITLE")
	droneEnv.PluginSuccessImgKey = os.Getenv("PLUGIN_SUCCESS_IMG_KEY")
	droneEnv.PluginFailureImgKey = os.Getenv("PLUGIN_FAILURE_IMG_KEY")
	droneEnv.PluginPoweredByImgKey = os.Getenv("PLUGIN_POWERED_BY_IMG_KEY")
	droneEnv.PluginPoweredByImgAlt = os.Getenv("PLUGIN_POWERED_BY_IMG_ALT")

	if droneEnv.PluginToken == "" {
		log.Println("feishu webhook access token can not be empty")
		os.Exit(1)
	}

	if droneEnv.PluginSecret == "" {
		log.Println("feishu sign secret can not be empty")
		os.Exit(1)
	}

	return droneEnv
}
