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

func GetTodayNotice(c *gin.Context) {
    // ListCard
	var buttons k.Kakao
	var quickReplies k.Kakao
	var items k.Kakao
	var label string
	var title string
    // ListCard

    // Parse
	var notices []models.Notice = Parse("", 30)
	now := time.Now()
	nowStr := fmt.Sprintf("%v.%02v.%v", now.Year()%100, int(now.Month()), now.Day())
	title = fmt.Sprintf("%v) 오늘 공지", nowStr)

	// Filtering out today's notice(s)
	for i, notice := range notices {
		if notice.Date != nowStr {
			notices = notices[:i]
			break
		}
	}

	if len(notices) > 5 {
		label = fmt.Sprintf("%v개 더보기", len(notices)-5)
		notices = notices[:5]
	} else {
		label = "아주대학교 공지"
	}

	if len(notices) == 0 {
		items.Add(k.ListItem{}.New("공지가 없습니다!", "", "http://k.kakaocdn.net/dn/APR96/btqqH7zLanY/kD5mIPX7TdD2NAxgP29cC0/1x1.jpg"))
	} else {
		for _, notice := range notices {
			if utf8.RuneCountInString(notice.Title) > 35 { // To count korean letters length correctly
				notice.Title = string([]rune(notice.Title)[0:32]) + "..."
			}
			description := fmt.Sprintf("%v %v", notice.Writer, notice.Date[len(notice.Date)-5:])
			noticeJSON := k.ListItemLink{}.New(notice.Title, description, "", notice.Link)
			items.Add(noticeJSON)
		}
	}

	buttons.Add(k.ShareButton{}.New("공유하기"))
	buttons.Add(k.LinkButton{}.New(label, AjouLink))

	quickReplies.Add(k.QuickReply{}.New("오늘", "오늘 공지 보여줘"))
	quickReplies.Add(k.QuickReply{}.New("어제", "어제 공지 보여줘"))

	c.JSON(200, k.BuildListCard(title, items, buttons, quickReplies))
}
```

```yaml
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
                            "label": "5개 더보기",
                            "webLinkUrl": "https://www.ajou.ac.kr/kr/ajou/notice.do"
                        }
                    ],
                    "header": {
                        "title": "21.02.19) 오늘 공지"
                    },
                    "items": [
                        {
                            ... 공지 ITEMS
                        }
                    ]
                }
            }
        ],
        "quickReplies": [
            {
                "action": "message",
                "label": "오늘",
                "messageText": "오늘 공지 보여줘"
            },
            {
                "action": "message",
                "label": "어제",
                "messageText": "어제 공지 보여줘"
            }
        ]
    },
    "version": "2.0"
}
```