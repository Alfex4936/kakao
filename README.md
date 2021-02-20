<div align="center">
<p>
    <img width="680" src="https://github.com/Alfex4936/kakaoChatbot-Ajou/blob/main/imgs/chatbot.png">
</p>
<h1>카카오톡 챗봇 빌더 도우미</h1>
<h2>Go언어용 WIP</h2>
</div>


## Gin 웹프레임워크 예제
```go
import k "github.com/Alfex4936/kakao"

// POST /simpletext
func returnSimpleText(c *gin.Context) {
	c.JSON(200, k.SimpleText{}.Build("메시지 입니다."))
}

// POST /listcard
func returnListCard(c *gin.Context) {
	listCard := k.ListCard{}.New(true) // whether to use quickReplies or not

	listCard.Title = "Hello ListCard"

	// ListItem: Title, Description, imageUrl
	listCard.Items.Add(k.ListItem{}.New("안녕하세요!", "", "http://image"))
	// LinkListItem: Title, Description, imageUrl, address
	listCard.Items.Add(k.ListItemLink{}.New("안녕하세요!", "", "", "https://www.naver.com/"))

	listCard.Buttons.Add(k.ShareButton{}.New("공유하기"))
	listCard.Buttons.Add(k.LinkButton{}.New("네이버 링크", "https://www.naver.com/"))

	// QuickReplies: label, message (메시지 없으면 라벨로 발화)
	listCard.QuickReplies.Add(k.QuickReply{}.New("오늘", "오늘 날씨 알려줘"))
	listCard.QuickReplies.Add(k.QuickReply{}.New("어제", "어제 날씨 알려줘"))

	c.JSON(200, listCard.Build())
}
```

```yaml
# SimpleText /simpletext
{
    "template": {
        "outputs": [
            {
                "simpleText": {
                    "text": "메시지 입니다."
                }
            }
        ]
    },
    "version": "2.0"
}

# ListCard /listcard
{
    "template": {
        "outputs": [
            {
                "listCard": {
                    "buttons": [
                        {
                            "action": "share",
                            "label": "공유하기"
                        },
                        {
                            "action": "webLink",
                            "label": "네이버 링크",
                            "webLinkUrl": "https://www.naver.com/"
                        }
                    ],
                    "header": {
                        "title": "Hello ListCard"
                    },
                    "items": [
                        {
                            "imageUrl": "http://image",
                            "title": "안녕하세요!"
                        },
                        {
                            "title": "안녕하세요!",
                            "link": {
                                "web": "https://www.naver.com/"
                            }
                        }
                    ]
                }
            }
        ],
        "quickReplies": [
            {
                "action": "message",
                "label": "오늘",
                "messageText": "오늘 날씨 알려줘"
            },
            {
                "action": "message",
                "label": "어제",
                "messageText": "어제 날씨 알려줘"
            }
        ]
    },
    "version": "2.0"
}
```