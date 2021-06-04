<div align="center">
<p>
    <img width="680" src="https://raw.githubusercontent.com/Alfex4936/kakaoChatbot-Ajou/main/imgs/chatbot.png">
</p>
<a href="https://pkg.go.dev/github.com/Alfex4936/kakao"><img src="https://pkg.go.dev/badge/github.com/Alfex4936/kakao.svg" alt="Go Reference"></a>
	
<a href="https://hits.seeyoufarm.com"><img src="https://hits.seeyoufarm.com/api/count/incr/badge.svg?url=https%3A%2F%2Fgithub.com%2FAlfex4936%2Fkakao&count_bg=%23000000&title_bg=%23000000&icon=strapi.svg&icon_color=%23FFFFFF&title=%7C&edge_flat=false"/></a>

<h2>카카오톡 챗봇 빌더 도우미</h2>
<h3>Go언어 전용</h3>
</div>

# 소개

Go언어로 카카오 챗봇 서버를 만들 때 좀 더 쉽게 JSON 메시지 응답을 만들 수 있게 도와줍니다.

SimpleText, SimpleImage, ListCard, Carousel, BasicCard, ContextControl JSON 데이터를 쉽게 만들 수 있도록 도와줍니다.

# 설치
```console
ubuntu:~$ go get -u github.com/Alfex4936/kakao
```

# 응답 타입별 아이템

Buttons: ShareButton (공유 버튼), LinkButton (링크 버튼), MsgButton (일반 메시지만), CallButton (전화 버튼)

Items: ListItem (일반), ListItemLink (링크 버전)

# 사용법 (gin 프레임워크 기준)

## 카카오 JSON 데이터 Bind

예제) 유저 발화문 얻기: kjson.UserRequest.Utterance

```go
import k "github.com/Alfex4936/kakao"

// JSON 요청 처리
var kjson k.Request
if err := c.BindJSON(&kjson); err != nil {
	c.AbortWithStatusJSON(200, k.SimpleText{}.Build("에러"))
	return
}

fmt.Println(kjson.UserRequest.Utterance)  // 유저 발화문
```

## SimpleText, SimpleImage, ListCard

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
func returnSimpleImage(c *gin.Context) {
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

Carousel은 New(커머스 카드, 케로셀헤더 bool) 초기화 후, 아이템들 입력 후 Build()

(Carousel + BasicCard + CarouselHeader도 정상적으로 작동합니다.)

```go
import k "github.com/Alfex4936/kakao"

// BasicCard 만들기
func returnBasicCard(c *gin.Context) {
	basicCard := k.BasicCard{}.New(true, true)  // 썸네일, 버튼 사용 여부
	basicCard.Title = "제목입니다."
	basicCard.Desc = "설명입니다."
	basicCard.ThumbNail = k.ThumbNail{}.New("http://썸네일링크")

	basicCard.Buttons.Add(k.LinkButton{}.New("날씨 홈피", "http://날씨 사이트"))
	c.PureJSON(200, basicCard.Build())
}

// Carousel 만들기
func returnCarousel(c *gin.Context) {
	carousel := k.Carousel{}.New(false, false)  // CommerceCard X, CarouselHeader X

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

## CommerceCard

New() -> Build()

설명, 가격, 할인, 통화 ("won"), 썸네일 1개 필수

```go
import . "github.com/Alfex4936/kakao"

// CommerceCard 만들기
func returnCommerceCard(c *gin.Context) {
	commerceCard := CommerceCard{}.New()
	
	commerceCard.Desc = "안녕하세요"
	commerceCard.Price = 10000
	commerceCard.Discount = 1000  // 할인
	commerceCard.Currency = "won"  // "won"만 지원
	commerceCard.ThumbNails.Add(ThumbNail{FixedRatio: true}.New("http://some.jpg"))  // 1개만 추가 가능, FixedRatio 변경 가능

	commerceCard.Buttons.Add(LinkButton{}.New("구매하기", "https://kakao/1542"))
	commerceCard.Buttons.Add(CallButton{}.New("전화하기", "354-86-00070"))
	commerceCard.Buttons.Add(ShareButton{}.New("공유하기"))

	c.PureJSON(200, commerceCard.Build())
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

# TODO

output 배열에 여러가지 응답을 저장할 수 있다.

그렇게 하려면 msg 리턴할 값을 하나 만들고 거기에 각 응답을 Build 하면 msg에 추가하는 방식이 나아보임
