package kakao

import (
	"encoding/json"
	"testing"
)

var JSON = `
{
	"intent": {
	  "id": "s1sabfeglft2g",
	  "name": "블록 이름"
	},
	"userRequest": {
	  "timezone": "Asia/Seoul",
	  "params": {
		"ignoreMe": "true"
	  },
	  "block": {
		"id": "s1sabfeglft2",
		"name": "블록 이름"
	  },
	  "utterance": "하이!",
	  "lang": null,
	  "user": {
		"id": "172514",
		"type": "accountId",
		"properties": {}
	  }
	},
	"bot": {
	  "id": "5fe45a6ddfbb1",
	  "name": "봇 이름"
	},
	"action": {
	  "name": "mbwnrkyh42",
	  "clientExtra": null,
	  "params": {
		"cate": "학사",
		"when": "yesterday",
		"sys_text": "코로나",
		"search": "소프트"
	  },
	  "id": "3f7ir2rgub3p",
	  "detailParams": {
		"sys_text": {
		  "origin": "코로나",
		  "value": "코로나",
		  "groupName": ""
		}
	  }
	}
  }
`

func TestUnMarshal(t *testing.T) {
	expected := "하이!"       // 발화 내용
	expectedParams := "소프트" // "search" 파라미터

	data := &Request{}

	_ = json.Unmarshal([]byte(JSON), data)

	if got := data.UserRequest.Utterance; got != expected {
		t.Errorf("Failed to UnMarshal Request: %q", got)
	} else {
		t.Logf("Correctly got utterance from request")
	}

	if got := data.Action.Params["search"]; got != expectedParams {
		t.Errorf("Failed to UnMarshal Request: %q", got)
	} else {
		t.Logf("Correctly got Params from request")
	}
}

func TestSimpleText(t *testing.T) {
	expected := json.RawMessage(`{"template":{"outputs":[{"simpleText":{"text":"안녕하세요."}}]},"version":"2.0"}`)

	data := SimpleText{}.Build("안녕하세요.")

	res, _ := json.Marshal(data)

	if got := string(res); got != string(expected) {
		t.Errorf("Failed to build SimpleText: %q, %q", got, expected)
	} else {
		t.Logf("Correctly built SimpleText")
	}
}

func TestListCard(t *testing.T) {
	expected := json.RawMessage(`{"template":{"outputs":[{"listCard":{"buttons":[{"action":"share","label":"공유하기"},{"action":"webLink","label":"네이버 링크","webLinkUrl":"https://www.naver.com/"}],"header":{"title":"Hello ListCard"},"items":[{"imageUrl":"http://image","title":"안녕하세요!"},{"title":"안녕하세요!","link":{"web":"https://www.naver.com/"}}]}}],"quickReplies":[{"action":"message","label":"오늘","messageText":"오늘 날씨 알려줘"},{"action":"message","label":"어제","messageText":"어제 날씨 알려줘"}]},"version":"2.0"}`)

	// Building
	listCard := ListCard{}.New(true) // whether to use quickReplies or not

	listCard.Title = "Hello ListCard"

	// ListItem: Title, Description, imageUrl
	listCard.Items.Add(ListItem{}.New("안녕하세요!", "", "http://image"))
	// LinkListItem: Title, Description, imageUrl, address
	listCard.Items.Add(ListItemLink{}.New("안녕하세요!", "", "", "https://www.naver.com/"))

	listCard.Buttons.Add(ShareButton{}.New("공유하기"))
	listCard.Buttons.Add(LinkButton{}.New("네이버 링크", "https://www.naver.com/"))

	// QuickReplies: label, message (메시지 없으면 라벨로 발화)
	listCard.QuickReplies.Add(QuickReply{}.New("오늘", "오늘 날씨 알려줘"))
	listCard.QuickReplies.Add(QuickReply{}.New("어제", "어제 날씨 알려줘"))

	res, _ := json.Marshal(listCard.Build())

	if got := string(res); got != string(expected) {
		t.Errorf("Failed to Marshal: %q, %q", got, string(expected))
	} else {
		t.Logf("Correctly built ListCard!")
	}

}

func TestBasicCard(t *testing.T) {
	expected := json.RawMessage(`{"template":{"outputs":[{"basicCard":{"title":"title!","description":"Descriptionss","thumbnail":{"imageUrl":"http://thumb"},"buttons":[{"action":"webLink","label":"날씨 홈피","webLinkUrl":"http://www"}]}}]},"version":"2.0"}`)

	// Building
	basicCard := BasicCard{}.New(true, true)
	basicCard.Title = "title!"
	basicCard.Desc = "Descriptionss"
	basicCard.ThumbNail = ThumbNail{}.New("http://thumb")

	basicCard.Buttons.Add(LinkButton{}.New("날씨 홈피", "http://www"))

	res, _ := json.Marshal(basicCard.Build())

	if got := string(res); got != string(expected) {
		t.Errorf("Failed to Marshal: %q, %q", got, string(expected))
	} else {
		t.Logf("Correctly built BasicCard!")
	}

}

func TestCarousel(t *testing.T) { // BasicCards
	expected := json.RawMessage(`{"template":{"outputs":[{"carousel":{"items":[{"title":"title1","description":"desc1","buttons":[{"label":"전화","action":"phone","phoneNumber":"031"},{"action":"webLink","label":"이메일","webLinkUrl":"mailto:example@world.com?subject=안녕하세요."}]},{"title":"title2","description":"desc2","buttons":[{"label":"전화","action":"phone","phoneNumber":"032"},{"action":"webLink","label":"이메일","webLinkUrl":"mailto:example@world.com?subject=안녕하세요."}]}],"type":"basicCard"}}]},"version":"2.0"}`)

	// Building
	carousel := Carousel{}.New(false)

	card1 := BasicCard{}.New(false, true)
	card1.Title = "title1"
	card1.Desc = "desc1"
	card1.Buttons.Add(CallButton{}.New("전화", "031"))
	card1.Buttons.Add(LinkButton{}.New("이메일", "mailto:example@world.com?subject=안녕하세요."))
	carousel.Cards.Add(card1)

	card2 := BasicCard{}.New(false, true)
	card2.Title = "title2"
	card2.Desc = "desc2"
	card2.Buttons.Add(CallButton{}.New("전화", "032"))
	card2.Buttons.Add(LinkButton{}.New("이메일", "mailto:example@world.com?subject=안녕하세요."))
	carousel.Cards.Add(card2)

	res, _ := json.Marshal(carousel.Build())

	if got := string(res); got != string(expected) {
		t.Errorf("Failed to Marshal: %q, %q", got, string(expected))
	} else {
		t.Logf("Correctly built CarouselCard (BasicCards)!")
	}

}

func TestSimpleImage(t *testing.T) {
	expected := json.RawMessage(`{"template":{"outputs":[{"simpleImage":{"altText":"ALT","imageUrl":"http://"}}]},"version":"2.0"}`)

	res, _ := json.Marshal(SimpleImage{}.Build("http://", "ALT"))

	if got := string(res); got != string(expected) {
		t.Errorf("Failed to Marshal: %q, %q", got, string(expected))
	} else {
		t.Logf("Correctly built SimpleImage!")
	}

}

func TestContextControl(t *testing.T) {
	expected := json.RawMessage(`{"context":{"values":[{"name":"abc","lifeSpan":10,"params":{"key1":"val1","key2":"val2"}},{"name":"ghi","lifeSpan":0}]},"version":"2.0"}`)

	params1 := map[string]string{
		"key1": "val1",
		"key2": "val2",
	}
	ctx := ContextControl{}.New()
	ctx.Values.Add(ContextValue{}.New("abc", 10, params1))
	ctx.Values.Add(ContextValue{}.New("ghi", 0, nil))

	res, _ := json.Marshal(ctx.Build())

	if got := string(res); got != string(expected) {
		t.Errorf("Failed to Marshal: %q, %q", got, string(expected))
	} else {
		t.Logf("Correctly built ContextControl!")
	}

}

// Benchmarks

func BenchmarkJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := &Request{}

		_ = json.Unmarshal([]byte(JSON), data)

		_ = data.UserRequest.Utterance

		_ = data.Action.Params["search"]
	}
}

func BenchmarkBuildAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// SimpleText
		stext := SimpleText{}.Build("안녕하세요.")
		json.Marshal(stext)

		// ListCard
		listCard := ListCard{}.New(true) // whether to use quickReplies or not
		listCard.Title = "Hello ListCard"
		// ListItem: Title, Description, imageUrl
		listCard.Items.Add(ListItem{}.New("안녕하세요!", "", "http://image"))
		// LinkListItem: Title, Description, imageUrl, address
		listCard.Items.Add(ListItemLink{}.New("안녕하세요!", "", "", "https://www.naver.com/"))
		listCard.Buttons.Add(ShareButton{}.New("공유하기"))
		listCard.Buttons.Add(LinkButton{}.New("네이버 링크", "https://www.naver.com/"))
		// QuickReplies: label, message (메시지 없으면 라벨로 발화)
		listCard.QuickReplies.Add(QuickReply{}.New("오늘", "오늘 날씨 알려줘"))
		listCard.QuickReplies.Add(QuickReply{}.New("어제", "어제 날씨 알려줘"))

		json.Marshal(listCard.Build())

		// BasicCard
		basicCard := BasicCard{}.New(true, true)
		basicCard.Title = "title!"
		basicCard.Desc = "Descriptionss"
		basicCard.ThumbNail = ThumbNail{}.New("http://thumb")
		basicCard.Buttons.Add(LinkButton{}.New("날씨 홈피", "http://www"))
		json.Marshal(basicCard.Build())

		// Carousel (BasicCard)
		carousel := Carousel{}.New(false)
		card1 := BasicCard{}.New(false, true)
		card1.Title = "title1"
		card1.Desc = "desc1"
		card1.Buttons.Add(CallButton{}.New("전화", "031"))
		card1.Buttons.Add(LinkButton{}.New("이메일", "mailto:example@world.com?subject=안녕하세요."))
		carousel.Cards.Add(card1)
		card2 := BasicCard{}.New(false, true)
		card2.Title = "title2"
		card2.Desc = "desc2"
		card2.Buttons.Add(CallButton{}.New("전화", "032"))
		card2.Buttons.Add(LinkButton{}.New("이메일", "mailto:example@world.com?subject=안녕하세요."))
		carousel.Cards.Add(card2)
		json.Marshal(carousel.Build())

		// SimpleImage
		json.Marshal(SimpleImage{}.Build("http://", "ALT"))

		// ContextControl
		params1 := map[string]string{
			"key1": "val1",
			"key2": "val2",
		}
		ctx := ContextControl{}.New()
		ctx.Values.Add(ContextValue{}.New("abc", 10, params1))
		ctx.Values.Add(ContextValue{}.New("ghi", 0, nil))
		json.Marshal(ctx.Build())
	}
}
