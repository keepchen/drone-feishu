package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	var (
		droneEnv = (DroneEnvironment{}).GetEnv()
		payload  Payload
	)
	card := (Card{}).Build(droneEnv.DroneRepo, droneEnv.DroneBranch, droneEnv.DroneCommitAuthor,
		droneEnv.DroneBuildStatus, droneEnv.DroneCommitMessage, droneEnv.DroneCommitLink, droneEnv.DroneBuildLink,
		droneEnv.PluginCardTitle, droneEnv.PluginSuccessImgKey, droneEnv.PluginFailureImgKey,
		droneEnv.PluginPoweredByImgKey, droneEnv.PluginPoweredByImgAlt)

	payload.Content = card
	payload.MsgType = "interactive"

	nowTimestamp := time.Now().Unix()
	payload.Timestamp = nowTimestamp
	sign, err := GenSign(droneEnv.PluginSecret, nowTimestamp)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	payload.Sign = sign

	js, err := json.Marshal(&payload)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = sendMessage(js, droneEnv.PluginToken)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}

//发送消息
func sendMessage(payload []byte, token string) error {
	var feiShuUrl = fmt.Sprintf("https://open.feishu.cn/open-apis/bot/v2/hook/%s", token)
	req, err := http.NewRequest("POST", feiShuUrl, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	statusCode := resp.StatusCode
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(statusCode, string(body))

	return nil
}
