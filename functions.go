/*
Package kakao 카카오 챗봇을 쉽게 만들 수 있게 도와줍니다.
*/
package kakao // import "github.com/Alfex4936/kakao"

// * Build() function

// Build (l *ListCard)
func (l *ListCard) Build() K {
	l.Title = K{"title": l.Title.(string)}
	template := K{"outputs": []K{{"listCard": K{"header": l.Title, "items": l.Items, "buttons": l.Buttons}}}}
	template["quickReplies"] = l.QuickReplies

	listCard := K{"version": "2.0", "template": template}

	return listCard
}

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

// * New() function

// New (l ListCard) Qr (bool) whether to use QuickReplies or not
func (l ListCard) New(Qr bool) *ListCard {
	l.Buttons = new(Kakao)
	l.Items = new(Kakao)
	if Qr {
		l.QuickReplies = new(Kakao)
	}
	return &l
}

// New (s ShareButton): label (string)
func (s ShareButton) New(label string) *ShareButton {
	s.Action = "share"
	s.Label = label
	return &s
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

// New (c CallButton): ...label, phoneNumber, msgTxt
func (c CallButton) New(params ...string) *CallButton {
	c.Label = params[0]
	c.Action = "phone"
	c.PhoneNumber = params[1]
	if len(params) >= 2 {
		c.MsgTxt = params[2]
	}
	return &c
}

// New (l ListItem): ...title, description, imageUrl
func (l ListItem) New(params ...string) *ListItem {
	n := len(params)

	l.Title = params[0]
	if n >= 1 {
		l.Desc = params[1]
	}
	if n >= 2 {
		l.Image = params[2]
	}
	return &l
}

// New (l ListItemLink): title, description, imageUrl, webLinkUrl
func (l ListItemLink) New(params ...string) *ListItemLink {
	n := len(params)

	l.Title = params[0]
	if n >= 1 {
		l.Desc = params[1]
	}
	if n >= 2 {
		l.Image = params[2]
	}
	if n >= 3 {
		l.Link.Link = params[3]
	}
	return &l
}
