package kakao

// * Build() function
// ListCard, SimpleText, BasicCard, Carousel

// Build (l *ListCard)
func (l *ListCard) Build() K {
	template := K{"outputs": &[]K{{"listCard": K{"header": &K{"title": l.Title.(string)}, "items": l.Items, "buttons": l.Buttons}}}}

	if l.QuickReplies != nil {
		template["quickReplies"] = l.QuickReplies
	}

	listCard := K{"version": "2.0", "template": template}

	return listCard
}

// Build (s SimpleText)
func (s SimpleText) Build(msg string, quickReplies Kakao) K {
	template := K{"outputs": []K{{"simpleText": K{"text": msg}}}}
	if quickReplies != nil {
		template["quickReplies"] = quickReplies
	}
	simpleText := K{"version": "2.0", "template": template}
	return simpleText
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

	switch {
	case !c.isCommerce && c.isHeader: // BasicCard + Header 테스트
		template = K{"outputs": []K{{"carousel": K{"type": "basicCard", "header": c.Header, "items": c.Cards}}}}
	case c.isCommerce && c.isHeader:
		template = K{"outputs": []K{{"carousel": K{"type": "commerceCard", "header": c.Header, "items": c.Cards}}}}
	case c.isCommerce && !c.isHeader:
		template = K{"outputs": []K{{"carousel": K{"type": "commerceCard", "items": c.Cards}}}}
	default:
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

// Build (cc CommerceCard)
func (cc CommerceCard) Build() K {
	template := K{"outputs": []K{{"commerceCard": cc}}}
	commerce := K{"version": "2.0", "template": template}
	return commerce
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
func (c Carousel) New(isCommerce, header bool) *Carousel {
	c.Type = "label"
	c.Cards = new(Kakao)
	if isCommerce {
		c.isCommerce = true
	}
	if header {
		c.Header = new(CarouselHeader)
		c.isHeader = true
	}
	return &c
}

// New (c CarouselHeader):
func (c CarouselHeader) New(title, description, imgURL string) *CarouselHeader {
	c.Title = title
	c.Desc = description
	c.ThumbNail = ThumbNail{}.New(imgURL)
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
func (t ThumbNail) New(imageURL string) *ThumbNail {
	t.ImageURL = imageURL
	return &t
}

// New (q QuickReply): label (string), msg (string)
func (q QuickReply) New(label, msg string) *QuickReply {
	q.Action = "message"
	q.Label = label
	q.Msg = msg
	return &q
}

// New (l LinkButton): label (string), link (string)
func (l LinkButton) New(label, link string) *LinkButton {
	l.Action = "webLink"
	l.Label = label
	l.WebLink = link
	return &l
}

// New (m MsgButton): label (string), link (string)
func (m MsgButton) New(label, msgTxt string) *MsgButton {
	m.Action = "message"
	m.Label = label
	m.MsgTxt = msgTxt
	return &m
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

// New (cc CommerceCard)
func (cc CommerceCard) New() *CommerceCard {
	cc.ThumbNails = new(Kakao)
	cc.Buttons = new(Kakao)
	return &cc
}
