package main

import (
	"fmt"
	"strings"
)

var (
	defaultTitle         = "Drone CI Notification"
	defaultSuccessImgKey = "img_v2_8205e2c9-ed56-48f4-a38d-efafaa69e35g"
	defaultFailureImgKey = "img_v2_e38b7ff9-2d11-467f-80c1-4c40fca7afag"
	logoImgKey           = "img_v2_1276726d-d435-4f52-b042-a76093814d1g"
	logoAltStr           = "keepchen"
)

type CardConfig struct {
	WideScreenMode bool `json:"wide_screen_mode"`
}

type CardElement struct {
	Tag      *string        `json:"tag,omitempty"`
	ImgKey   *string        `json:"img_key,omitempty"`
	Content  *string        `json:"content,omitempty"`
	Title    *CardTitle     `json:"title,omitempty"`
	Alt      *CardAlt       `json:"alt,omitempty"`
	Elements []*CardElement `json:"elements,omitempty"`
}

type CardAlt struct {
	Tag     string `json:"tag"`
	Content string `json:"content"`
}

type CardTitle struct {
	Tag     string `json:"tag"`
	Content string `json:"content"`
}

type CardHeader struct {
	Template string    `json:"template"`
	Title    CardTitle `json:"title"`
}

type Card struct {
	Config   CardConfig     `json:"config"`
	Header   CardHeader     `json:"header"`
	Elements []*CardElement `json:"elements,omitempty"`
}

// Build 组装卡片信息
func (Card) Build(repo, branch, author, status, commitMsg, repoUrl, droneUrl string,
	cardTitle, successImgKey, failureImgKey, poweredByImgKey, poweredByImgAlt string) Card {

	//override default card setting
	if cardTitle != "" {
		defaultTitle = cardTitle
	}
	if successImgKey != "" {
		defaultSuccessImgKey = successImgKey
	}
	if failureImgKey != "" {
		defaultFailureImgKey = failureImgKey
	}
	if poweredByImgKey != "" {
		logoImgKey = poweredByImgKey
	}
	if poweredByImgAlt != "" {
		logoAltStr = poweredByImgAlt
	}
	var titleColor = "blue"
	if status != "success" {
		titleColor = "red"
	}

	var card Card
	card.Config = CardConfig{
		WideScreenMode: true,
	}

	card.Header = CardHeader{
		Template: titleColor,
		Title: CardTitle{
			Content: defaultTitle,
			Tag:     "plain_text",
		},
	}

	{
		var (
			tag     = "markdown"
			content = fmt.Sprintf("**%s**", repo)
		)
		card.Elements = append(card.Elements, &CardElement{
			Tag:     &tag,
			Content: &content,
		})
	}

	{
		var (
			tag = "hr"
		)
		card.Elements = append(card.Elements, &CardElement{
			Tag: &tag,
		})
	}

	{
		var (
			tag     = "markdown"
			content = fmt.Sprintf("*%s分支*\n**提交者：** *%s*", branch, author)
		)
		card.Elements = append(card.Elements, &CardElement{
			Tag:     &tag,
			Content: &content,
		})
	}

	{
		var (
			tag   = "img"
			title = CardTitle{
				Tag:     "lark_md",
				Content: "",
			}
			imgKey = ""
			alt    = CardAlt{
				Tag:     "plain_text",
				Content: "",
			}
		)
		if status == "success" {
			imgKey = defaultSuccessImgKey
			alt.Content = "success"
		} else {
			imgKey = defaultFailureImgKey
			alt.Content = "failure"
		}
		card.Elements = append(card.Elements, &CardElement{
			Tag:    &tag,
			Title:  &title,
			ImgKey: &imgKey,
			Alt:    &alt,
		})
	}

	//debug multiple lines commit message
//	commitMsg = `
//Merge branch 'branch1' into 'main'
//
//feat:1.do something，2.another...
//
//See merge request game/game-web-frontend!478
//`

	{
		var (
			tag     = "markdown"
			content = fmt.Sprintf("**Commit信息：**  \n%s", handleMultipleLinesCommitMsg(commitMsg))
		)
		card.Elements = append(card.Elements, &CardElement{
			Tag:     &tag,
			Content: &content,
		})
	}

	{
		var (
			tag = "hr"
		)
		card.Elements = append(card.Elements, &CardElement{
			Tag: &tag,
		})
	}

	{
		var (
			tag     = "markdown"
			content = fmt.Sprintf("[查看Commit信息](%s) | [查看构建信息](%s)", repoUrl, droneUrl)
		)
		card.Elements = append(card.Elements, &CardElement{
			Tag:     &tag,
			Content: &content,
		})
	}

	{
		var (
			elements []*CardElement
			tag      = "plain_text"
			content  = "Powered By"
			element1 = &CardElement{
				Tag:     &tag,
				Content: &content,
			}
			tag2   = "img"
			imgKey = logoImgKey
			alt    = CardAlt{
				Tag:     "plain_text",
				Content: logoAltStr,
			}
			element2 = &CardElement{
				Tag:    &tag2,
				ImgKey: &imgKey,
				Alt:    &alt,
			}
		)
		elements = append(elements, element1, element2)

		var tag3 = "note"
		card.Elements = append(card.Elements, &CardElement{
			Tag:      &tag3,
			Elements: elements,
		})
	}

	return card
}

// handle multiple lines commit message and return markdown string
func handleMultipleLinesCommitMsg(commitMsg string) string {
	commitMsgLines := strings.Split(commitMsg, "\n")
	var markdownHolders = make([]string, 1)
	for _, line := range commitMsgLines {
		if len(strings.Trim(line, "")) == 0 {
			continue
		}
		markdownHolders = append(markdownHolders, fmt.Sprintf("*%s*", line))
	}

	return strings.Join(markdownHolders, "\n")
}