/*
Package kakao 카카오 챗봇을 쉽게 만들 수 있게 도와줍니다.
*/
package kakao // import "github.com/Alfex4936/kakao"

// * Build() function
// ListCard, SimpleText, BasicCard, Carousel

// Build (l *ListCard)
func (l *ListCard) Build() K {
	l.Title = K{"title": l.Title.(string)}
	template := K{"outputs": []K{{"listCard": K{"header": l.Title, "items": l.Items, "buttons": l.Buttons}}}}
	template["quickReplies"] = l.QuickReplies

	listCard := K{"version": "2.0", "template": template}

	return listCard
}

// Build (s SimpleText)
func (s SimpleText) Build(msg string) *SimpleText {
	s.Version = "2.0"

	var temp []struct {
		SimpleText Text `json:"simpleText"`
	}
	simpleText := Text{Text: msg}

	text := struct {
		SimpleText Text `json:"simpleText"`
	}{SimpleText: simpleText}

	temp = append(temp, text)

	s.Template.Outputs = temp
	return &s
}

// Build (s SimpleImage)
func (s SimpleImage) Build(url, alt string) K {
	template := K{"outputs": []K{{"simpleImage": K{"imageUrl": url, "altText": alt}}}}
	simpleImage := K{"version": "2.0", "template": template}
	return simpleImage
}

// Build (c *Carousel)
func (c *Carousel) Build() K {
	var template K
	if c.isHeader {
		template = K{"outputs": []K{{"carousel": K{"type": "commerceCard", "header": c.Header, "items": c.Cards}}}}
	} else {
		template = K{"outputs": []K{{"carousel": K{"type": "basicCard", "items": c.Cards}}}}
	}
	carousel := K{"version": "2.0", "template": template}
	return carousel
}

// Build (b BasicCard)
func (b BasicCard) Build() K {
	template := K{"outputs": []K{{"basicCard": b}}}
	basicCard := K{"version": "2.0", "template": template}
	return basicCard
}

// Build (ctx ContextControl)
func (ctx ContextControl) Build() K {
	context := K{"values": ctx.Values}
	contextControl := K{"version": "2.0", "context": context}
	return contextControl
}

// * New() function

// New (l ListCard) Qr (bool) QuickReplies을 사용할지 true or false
func (l ListCard) New(Qr bool) *ListCard {
	l.Buttons = new(Kakao)
	l.Items = new(Kakao)
	if Qr {
		l.QuickReplies = new(Kakao)
	}
	return &l
}

// New (c Carousel): header (bool)
func (c Carousel) New(header bool) *Carousel {
	c.Type = "label"
	c.Cards = new(Kakao)
	if header {
		c.Header = new(CarouselHeader)
		c.isHeader = true
	}
	return &c
}

// New (b BasicCard): tb, btn (bool)  썸네일, 버튼 사용 여부
func (b BasicCard) New(tb, btn bool) *BasicCard {
	if tb {
		b.ThumbNail = new(ThumbNail)
	}
	if btn {
		b.Buttons = new(Kakao)
	}
	return &b
}

// New (t ThumbNail): ImageUrl
func (t ThumbNail) New(tm string) *ThumbNail {
	t.ImageURL = tm
	return &t
}

// New (q QuickReply): label (string), msg (string)
func (q QuickReply) New(label, msg string) *QuickReply {
	q.Action = "message"
	q.Label = label
	q.Msg = msg
	return &q
}

// New (l LinkButton): msg (string), link (string)
func (l LinkButton) New(msg, link string) *LinkButton {
	l.Action = "webLink"
	l.Label = msg
	l.WebLink = link
	return &l
}

// New (s ShareButton): label (string)
func (s ShareButton) New(label string) *ShareButton {
	s.Action = "share"
	s.Label = label
	return &s
}

// New (c CallButton): ...label, phoneNumber, msgTxt
func (c CallButton) New(params ...string) *CallButton {
	c.Label = params[0]
	c.Action = "phone"
	c.PhoneNumber = params[1]
	if len(params) >= 3 {
		c.MsgTxt = params[2]
	}
	return &c
}

// New (l ListItem): ...title, description, imageUrl
func (l ListItem) New(params ...string) *ListItem {
	n := len(params)

	l.Title = params[0]
	if n >= 2 {
		l.Desc = params[1]
	}
	if n >= 3 {
		l.Image = params[2]
	}
	return &l
}

// New (l ListItemLink): title, description, imageUrl, webLinkUrl
func (l ListItemLink) New(params ...string) *ListItemLink {
	n := len(params)

	l.Title = params[0]
	if n >= 2 {
		l.Desc = params[1]
	}
	if n >= 3 {
		l.Image = params[2]
	}
	if n >= 4 {
		l.Link.Link = params[3]
	}
	return &l
}

// New (ctx ContextControl)
func (ctx ContextControl) New() *ContextControl {
	ctx.Values = new(Kakao)
	return &ctx
}

// New (ctv ContextValue)
func (ctv ContextValue) New(name string, lifeSpan int, params map[string]string) *ContextValue {
	ctv.Name = name
	ctv.LifeSpan = lifeSpan
	if params != nil {
		ctv.Params = params
	}
	return &ctv
}
