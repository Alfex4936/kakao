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
	carousel := Carousel{}.New(false, false)

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

func TestCarouselCommerce(t *testing.T) { // CommerceCards
	expected := json.RawMessage(`""{\"template\":{\"outputs\":[{\"carousel\":{\"items\":[{\"description\":\"안녕하세요\",\"price\":10000,\"discount\":1000,\"currency\":\"won\",\"thumbnails\":[{\"imageUrl\":\"http://some.jpg\"}],\"buttons\":[{\"action\":\"webLink\",\"label\":\"구매하기\",\"webLinkUrl\":\"https://kakao/1542\"},{\"label\":\"전화하기\",\"action\":\"phone\",\"phoneNumber\":\"354-86-00070\"},{\"action\":\"share\",\"label\":\"공유하기\"}]},{\"description\":\"안녕하세요\",\"price\":5900,\"discount\":500,\"currency\":\"won\",\"thumbnails\":[{\"imageUrl\":\"http://some22.jpg\"}],\"buttons\":[{\"action\":\"webLink\",\"label\":\"구매하기\",\"webLinkUrl\":\"https://kakao/1543\"},{\"label\":\"전화하기\",\"action\":\"phone\",\"phoneNumber\":\"354-86-00071\"},{\"action\":\"share\",\"label\":\"공유하기\"}]}],\"type\":\"commerceCard\"}}]},\"version\":\"2.0\"}""`)

	// Building
	carousel := Carousel{}.New(true, false)

	commerceCard1 := CommerceCard{}.New()
	commerceCard1.Desc = "안녕하세요"
	commerceCard1.Price = 10000
	commerceCard1.Discount = 1000
	commerceCard1.Currency = "won"
	commerceCard1.ThumbNails.Add(ThumbNail{}.New("http://some.jpg"))
	commerceCard1.Buttons.Add(LinkButton{}.New("구매하기", "https://kakao/1542"))
	commerceCard1.Buttons.Add(CallButton{}.New("전화하기", "354-86-00070"))
	commerceCard1.Buttons.Add(ShareButton{}.New("공유하기"))
	carousel.Cards.Add(commerceCard1)

	commerceCard2 := CommerceCard{}.New()
	commerceCard2.Desc = "안녕하세요"
	commerceCard2.Price = 5900
	commerceCard2.Discount = 500
	commerceCard2.Currency = "won"
	commerceCard2.ThumbNails.Add(ThumbNail{}.New("http://some22.jpg"))
	commerceCard2.Buttons.Add(LinkButton{}.New("구매하기", "https://kakao/1543"))
	commerceCard2.Buttons.Add(CallButton{}.New("전화하기", "354-86-00071"))
	commerceCard2.Buttons.Add(ShareButton{}.New("공유하기"))
	carousel.Cards.Add(commerceCard2)

	res, _ := json.Marshal(carousel.Build())
	t.Logf("Correctly built CarouselCard (CommerceCards)! %q", string(res))

	if got := string(res); got != string(expected) {
		t.Errorf("Failed to Marshal: %q, %q", got, string(expected))
	} else {
		t.Logf("Correctly built CarouselCard (CommerceCards)!")
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

func TestCommerceCard(t *testing.T) {
	expected := json.RawMessage(`{"template":{"outputs":[{"commerceCard":{"description":"안녕하세요","price":10000,"discount":1000,"currency":"won","thumbnails":[{"imageUrl":"http://some.jpg"}],"buttons":[{"action":"webLink","label":"구매하기","webLinkUrl":"https://kakao/1542"},{"label":"전화하기","action":"phone","phoneNumber":"354-86-00070"},{"action":"share","label":"공유하기"}]}}]},"version":"2.0"}`)

	// Building
	commerceCard := CommerceCard{}.New()
	commerceCard.Desc = "안녕하세요"
	commerceCard.Price = 10000
	commerceCard.Discount = 1000
	commerceCard.Currency = "won"
	commerceCard.ThumbNails.Add(ThumbNail{}.New("http://some.jpg"))

	commerceCard.Buttons.Add(LinkButton{}.New("구매하기", "https://kakao/1542"))
	commerceCard.Buttons.Add(CallButton{}.New("전화하기", "354-86-00070"))
	commerceCard.Buttons.Add(ShareButton{}.New("공유하기"))

	res, _ := json.Marshal(commerceCard.Build())

	if got := string(res); got != string(expected) {
		t.Errorf("Failed to Marshal: %q, %q", got, string(expected))
	} else {
		t.Logf("Correctly built CommerceCard!")
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
		carousel := Carousel{}.New(false, false)
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

		// Carousel CommerceCards
		carouselCom := Carousel{}.New(true, false)

		commerceCard1 := CommerceCard{}.New()
		commerceCard1.Desc = "안녕하세요"
		commerceCard1.Price = 10000
		commerceCard1.Discount = 1000
		commerceCard1.Currency = "won"
		commerceCard1.ThumbNails.Add(ThumbNail{}.New("http://some.jpg"))
		commerceCard1.Buttons.Add(LinkButton{}.New("구매하기", "https://kakao/1542"))
		commerceCard1.Buttons.Add(CallButton{}.New("전화하기", "354-86-00070"))
		commerceCard1.Buttons.Add(ShareButton{}.New("공유하기"))
		carouselCom.Cards.Add(commerceCard1)

		commerceCard2 := CommerceCard{}.New()
		commerceCard2.Desc = "안녕하세요"
		commerceCard2.Price = 5900
		commerceCard2.Discount = 500
		commerceCard2.Currency = "won"
		commerceCard2.ThumbNails.Add(ThumbNail{}.New("http://some22.jpg"))
		commerceCard2.Buttons.Add(LinkButton{}.New("구매하기", "https://kakao/1543"))
		commerceCard2.Buttons.Add(CallButton{}.New("전화하기", "354-86-00071"))
		commerceCard2.Buttons.Add(ShareButton{}.New("공유하기"))
		carouselCom.Cards.Add(commerceCard2)

		json.Marshal(carouselCom.Build())
	}
}
