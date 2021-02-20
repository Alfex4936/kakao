package kakao

// Carousel ...
type Carousel struct {
	Type   string          `json:"type,omitempty"`
	Cards  *Kakao          `json:"items,omitempty"` // items
	Header *CarouselHeader `json:"header,omitempty"`
}

// CarouselHeader 필수 여부 X
type CarouselHeader struct {
	Title     string     `json:"title"`
	Desc      string     `json:"description"`
	ThumbNail *ThumbNail `json:"thumbnail"`
}

// BasicCard Title, Thumbnail 중 하나는 있어야 합니다.
type BasicCard struct {
	Title     string     `json:"title,omitempty"`       // 필수 X, 최대 2줄
	Desc      string     `json:"description,omitempty"` // 필수 X, 최대 230자, Carousel에선 76자
	ThumbNail *ThumbNail `json:"thumbnail,omitempty"`   // 필수 X
	Buttons   *Kakao     `json:"buttons,omitempty"`     // 필수 X, 최대 3개
}

/* ThumbNail ...
"thumbnails": [
	{
		"imageUrl": "http://k.kakaocdn.net/dn/83BvP/bl20duRC1Q1/lj3JUcmrzC53YIjNDkqbWK/i_6piz1p.jpg", 필수 O
		"link": {
			"web": "https://store.kakaofriends.com/kr/products/1542"
		} 필수 여부 X
	}
	]
*/
// ThumbNail
type ThumbNail struct {
	ImageURL   string `json:"imageUrl"`             // 필수 O
	Link       *Link  `json:"link,omitempty"`       // 필수 X
	FixedRatio bool   `json:"fixedRatio,omitempty"` // 필수 X
}

// ListCard ...
type ListCard struct {
	Title        interface{} `json:"header"`
	Buttons      *Kakao      `json:"buttons"`
	QuickReplies *Kakao      `json:"quickReplies,omitempty"`
	Items        *Kakao      `json:"items"`
}

// SimpleText ...
type SimpleText struct {
	Template struct {
		Outputs []struct {
			SimpleText Text `json:"simpleText"`
		} `json:"outputs"`
	} `json:"template"`
	Version string `json:"version"`
}

// Text for SimpleText
type Text struct {
	Text string `json:"text"`
}

// QuickReply ...
type QuickReply struct {
	Action  string `json:"action"` // message 또는 block
	Label   string `json:"label"`  // 필수
	Msg     string `json:"messageText"`
	BlockID string `json:"blockId,omitempty"` // action "block"일 때 필수임
	Extra   *Kakao `json:"extra,omitempty"`
}

// * Buttons START

// ShareButton ...
type ShareButton struct {
	Action string `json:"action"`                // 필수
	Label  string `json:"label"`                 // 필수
	MsgTxt string `json:"messageText,omitempty"` // MsgTxt가 있으면 이게 발화로 나감, 없으면 Label이 발화

}

// LinkButton ...
type LinkButton struct {
	Action  string `json:"action"` // 필수
	Label   string `json:"label"`  // 필수
	WebLink string `json:"webLinkUrl"`
	MsgTxt  string `json:"messageText,omitempty"` // MsgTxt가 있으면 이게 발화로 나감, 없으면 Label이 발화
}

// MsgButton ,...
type MsgButton struct {
	Label   string `json:"label"`                 // 버튼 14자(가로배열 2개 8자) 필수
	Action  string `json:"action"`                // 필수
	MsgTxt  string `json:"messageText,omitempty"` // MsgTxt가 있으면 이게 발화로 나감, 없으면 Label이 발화
	BlockID string `json:"blockId,omitempty"`
}

// CallButton ,...
type CallButton struct {
	Label       string `json:"label"`                 // 버튼 14자(가로배열 2개 8자) 필수
	Action      string `json:"action"`                // phone 필수
	PhoneNumber string `json:"phoneNumber"`           // 필수
	MsgTxt      string `json:"messageText,omitempty"` // MsgTxt가 있으면 이게 발화로 나감, 없으면 Label이 발화
	BlockID     string `json:"blockId,omitempty"`
}

// * Buttons END

// * Items START

// ListItem ...
type ListItem struct {
	Image string `json:"imageUrl,omitempty"`
	Desc  string `json:"description,omitempty"`
	Title string `json:"title"`
}

// ListItemLink ...
type ListItemLink struct {
	Title string `json:"title"`
	Desc  string `json:"description,omitempty"`
	Image string `json:"imageUrl,omitempty"`
	Link  Link   `json:"link"` // omit possible
}

// Link for ListItemLink
type Link struct {
	Link string `json:"web"`
}

// * Items END

// Request 카카오 서버에서 POST JSON 데이터용
type Request struct {
	Action struct {
		ID          string `json:"id"`
		ClientExtra struct {
		} `json:"clientExtra"`
		DetailParams map[string]interface{} `json:"detailParams"`
		Name         string                 `json:"name"`
		Params       map[string]interface{} `json:"params"`
	} `json:"action"`
	Bot struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"bot"`
	Contexts []interface{} `json:"contexts"`
	Intent   struct {
		ID    string `json:"id"`
		Extra struct {
			Reason struct {
				Code    int64  `json:"code"`
				Message string `json:"message"`
			} `json:"reason"`
		} `json:"extra"`
		Name string `json:"name"`
	} `json:"intent"`
	UserRequest struct {
		Block struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"block"`
		Lang   string `json:"lang"`
		Params struct {
			IgnoreMe bool   `json:"ignoreMe,string"`
			Surface  string `json:"surface"`
		} `json:"params"`
		Timezone string `json:"timezone"`
		User     struct {
			ID         string `json:"id"`
			Properties struct {
				BotUserKey  string `json:"botUserKey"`
				BotUserKey2 string `json:"bot_user_key"`
			} `json:"properties"`
			Type string `json:"type"`
		} `json:"user"`
		Utterance string `json:"utterance"`
	} `json:"userRequest"`
}
