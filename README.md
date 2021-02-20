<div align="center">
<p>
    <img width="680" src="https://raw.githubusercontent.com/Alfex4936/kakaoChatbot-Ajou/main/imgs/chatbot.png">
</p>
<a href="https://pkg.go.dev/github.com/Alfex4936/kakao"><img src="https://pkg.go.dev/badge/github.com/Alfex4936/kakao.svg" alt="Go Reference"></a>

<h2>카카오톡 챗봇 빌더 도우미</h2>
<h3>Go언어 전용</h3>
</div>

# 소개

파이썬처럼 dictionary 타입을 쉽게 이용할 수 있지 않기 때문에,

SimpleText, SimpleImage, ListCard, Carousel, BasicCard, ContextControl JSON 데이터를 쉽게 만들 수 있도록 도와줍니다.

# 설치
```console
ubuntu:~$ go get -u github.com/Alfex4936/kakao
```


# 사용법 (gin 프레임워크 기준)

## 카카오 JSON 데이터 Bind

예제) 유저 발화문 얻기: kjson.UserRequest.Utterance

```go
// JSON 요청 처리
var kjson k.Request
if err := c.BindJSON(&kjson); err != nil {
	c.AbortWithStatusJSON(200, k.SimpleText{}.Build("에러"))
	return
}

fmt.Println(kjson.UserRequest.Utterance)  // 유저 발화문
```

## SimpleText, SimpleImage ListCard

SimpleText는 메시지만 넘기면 됩니다.

SimpleImage는 이미지 주소와 로딩 실패 메시지를 넘기면 됩니다.

ListCard는 New() 초기화 후, 아이템들 입력 후 Build()

```go
import k "github.com/Alfex4936/kakao"

// 카카오 POST json 데이터
var kjson k.Request
if err := c.BindJSON(&kjson); err != nil {
    c.AbortWithStatusJSON(200, k.SimpleText{}.Build("잠시 후 다시"))
    return
}

// POST /simpletext
func returnSimpleText(c *gin.Context) {
	c.PureJSON(200, k.SimpleText{}.Build("메시지 입니다."))
}

// POST /simpleimage
func returnSimpleText(c *gin.Context) {
	c.PureJSON(200, k.SimpleImage{}.Build("http://", "ALT"))
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

	c.PureJSON(200, listCard.Build())
}
```

## BasicCard, Carousel

BasicCard는 New(썸네일, 버튼 bool) 초기화 후, 아이템들 입력 후 Build()

Carousel은 New(케로셀헤더 bool) 초기화 후, 아이템들 입력 후 Build()

```go
import k "github.com/Alfex4936/kakao"

// BasicCard 만들기
func returnBasicCard(c *gin.Context) {
	basicCard := k.BasicCard{}.New(true, true)  // 썸네일, 버튼 사용 여부
	basicCard.Title = "제목입니다."
	basicCard.Desc = "설명입니다."
	basicCard.Thumbnail = k.Thumbnail{}.New("http://썸네일링크")

	basicCard.Buttons.Add(k.LinkButton{}.New("날씨 홈피", "http://날씨 사이트"))
	c.PureJSON(200, basicCard.Build())
}

// Carousel 만들기
func returnCarousel(c *gin.Context) {
	carousel := k.Carousel{}.New(false)  // CarouselHeader 사용 여부

    for _, person := range people.PhoneNumber {
        // basicCard 케로셀에 담기
		card1 := k.BasicCard{}.New(false, true)
		card1.Title = fmt.Sprintf("%v (%v)", person.Name, person.Email)
		card1.Desc = fmt.Sprintf("전화번호: %v\n부서명: %v", intel+person.TelNo, person.DeptNm)

        // 전화 버튼, 웹 링크 버튼 케로셀에 담기
		card1.Buttons.Add(k.CallButton{}.New("전화", intel+person.TelNo))
		card1.Buttons.Add(k.LinkButton{}.New("이메일", fmt.Sprintf("mailto:%s?subject=안녕하세요.", person.Email)))

		carousel.Cards.Add(card1)
	}

	c.PureJSON(200, carousel.Build())
}
```

## ContextControl

name, lifeSpan, params을 받습니다.

(params는 필수 여부가 아니므로 nil 처리로 표시 X)

```go
import k "github.com/Alfex4936/kakao"

// ContextControl 만들기
func returnContextControl(c *gin.Context) {
	params1 := map[string]string{
		"key1": "val1",
		"key2": "val2",
	}
	ctx := k.ContextControl{}.New()
	ctx.Values.Add(k.ContextValue{}.New("abc", 10, params1))
	ctx.Values.Add(k.ContextValue{}.New("ghi", 0, nil))

	c.PureJSON(200, ctx.Build())
}
```