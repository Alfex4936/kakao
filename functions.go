/*
Package kakao 카카오 챗봇을 쉽게 만들 수 있게 도와줍니다.
*/
package kakao // import "github.com/Alfex4936/kakao"

// BuildSimpleText takes msg (string)
func BuildSimpleText(msg string) *SimpleText {
	stext := &SimpleText{Version: "2.0"}

	var temp []struct {
		SimpleText Text `json:"simpleText"`
	}
	simpleText := Text{Text: msg}

	text := struct {
		SimpleText Text `json:"simpleText"`
	}{SimpleText: simpleText}

	temp = append(temp, text)

	stext.Template.Outputs = temp
	return stext
}

// BuildListCard takes title string, items, buttons, quickReplies []interface{}
func BuildListCard(title string, items, buttons, quickReplies []interface{}) K {
	// Card
	header := K{"title": title}

	// Make a template
	template := K{"outputs": []K{{"listCard": K{"header": header, "items": items, "buttons": buttons}}}}

	if len(quickReplies) > 0 {
		template["quickReplies"] = quickReplies // Optional
	}

	listCard := K{"version": "2.0", "template": template}

	return listCard
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
