package kakao

// Carousel 케로셀은 여러 장의 카드를 하나의 메세지에 일렬로 포함하는 타입입니다.
type Carousel struct {
	Type   string          `json:"type,omitempty"`   // Build()에서 지정됨. (commerceCard, basicCard)
	Cards  *Kakao          `json:"items,omitempty"`  // Carousel에 쓰일 카드 (한 종류의 카드만 담을 것)
	Header *CarouselHeader `json:"header,omitempty"` // 현재 CommerceCard만 지원

	isHeader bool // false
}

// CarouselHeader 케로셀의 커버를 제공합니다. (필수 여부 X)
type CarouselHeader struct {
	Title     string     `json:"title"`       // 필수 O, 최대 2줄
	Desc      string     `json:"description"` // 필수 O, 최대 3줄
	ThumbNail *ThumbNail `json:"thumbnail"`   // 필수 O, 현재 imageUrl 값만 지원합니다.
}

// BasicCard Title, Thumbnail 중 하나는 있어야 합니다.
//
// 기본 카드형 출력 요소입니다. 기본 카드는 소셜, 썸네일, 프로필 등을 통해서 사진이나 글, 인물 정보 등을 공유할 수 있습니다. 기본 카드는 제목과 설명 외에 썸네일 그룹, 프로필, 버튼 그룹, 소셜 정보를 추가로 포함합니다.
type BasicCard struct {
	Title     string     `json:"title,omitempty"`       // 필수 X, 최대 2줄
	Desc      string     `json:"description,omitempty"` // 필수 X, 최대 230자, Carousel에선 76자
	ThumbNail *ThumbNail `json:"thumbnail,omitempty"`   // 필수 X
	Buttons   *Kakao     `json:"buttons,omitempty"`     // 필수 X, 최대 3개
}

// ThumbNail BasicCard에 사용될 수 있는 요소
//
// 현재 ImageURL만 지원함
type ThumbNail struct {
	// ImageURL 필수 O
	//
	// 이미지의 url입니다.
	ImageURL string `json:"imageUrl"`

	// Link 필수 X
	//
	// 이미지 클릭시 작동하는 link입니다.
	Link *Link `json:"link,omitempty"`
	// FixedRatio 필수 X, 기본값 false.
	//
	// true: 이미지 영역을 1:1 비율로 두고 이미지의 원본 비율을 유지합니다. 이미지가 없는 영역은 흰색으로 노출합니다.
	//
	// false: 이미지 영역을 2:1 비율로 두고 이미지의 가운데를 크롭하여 노출합니다.
	FixedRatio bool `json:"fixedRatio,omitempty"`

	// fixedRatio를 true로 설정할 경우 필요한 값입니다. 실제 이미지 사이즈와 다른 값일 경우 원본이미지와 다르게 표현될 수 있습니다.
	Width int `json:"width,omitempty"`
	// fixedRatio를 true로 설정할 경우 필요한 값입니다. 실제 이미지 사이즈와 다른 값일 경우 원본이미지와 다르게 표현될 수 있습니다.
	Height int `json:"height,omitempty"`
}

// ListCard 리스트 카드형 출력 요소입니다. 리스트 카드는 표현하고자 하는 대상이 다수일 때, 이를 효과적으로 전달할 수 있습니다.
//
// 헤더와 아이템을 포함하며, 헤더는 리스트 카드의 상단에 위치합니다. 리스트 상의 아이템은 각각의 구체적인 형태를 보여주며, 제목과 썸네일, 상세 설명을 포함합니다.
type ListCard struct {
	Title        interface{} `json:"header"`                 //  필수 여부 O, 카드의 상단 항목
	Buttons      *Kakao      `json:"buttons"`                // 필수 여부 X, 최대 2개
	QuickReplies *Kakao      `json:"quickReplies,omitempty"` // 필수 여부 X
	Items        *Kakao      `json:"items"`                  // 필수 여부 O, 카드의 각각 아이템 (최대 5개)
}

// SimpleText 간단한 텍스트형 출력 요소입니다.
//
// text가 500자가 넘는 경우, 500자 이후의 글자는 생략되고 전체 보기 버튼을 통해서 전체 내용을 확인할 수 있습니다.
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

// SimpleImage 간단한 이미지형 출력 요소입니다.
//
// 이미지 링크 주소를 포함하면 이를 스크랩하여 사용자에게 전달합니다.
//
// 이미지 링크 주소가 유효하지 않을 수 있기 때문에, 대체 텍스트를 꼭 포함해야 합니다.
type SimpleImage struct {
	ImageURL string `json:"imageUrl"` // 필수, 전달하고자 하는 이미지의 url입니다
	AltText  string `json:"altText"`  // 필수, url이 유효하지 않은 경우, 전달되는 텍스트입니다 (최대 1000자)
}

// QuickReply 바로가기 응답은 발화와 동일합니다. 대신, 사용자가 직접 발화를 입력하지 않아도 선택을 통해서 발화를 전달하거나 다른 블록을 호출할 수 있습니다.
//
// Msg를 지정하지 않으면 Label로 발화가 됩니다.
//
// 사용법: k.QuickReply{}.New(label, msg)
type QuickReply struct {
	Action  string `json:"action"` // message 또는 block
	Label   string `json:"label"`  // 필수
	Msg     string `json:"messageText"`
	BlockID string `json:"blockId,omitempty"` // action "block"일 때 필수임
	Extra   *Kakao `json:"extra,omitempty"`
}

// * Buttons START

// ShareButton action="share"가 기본 값인 버튼입니다.
type ShareButton struct {
	Action string `json:"action"`                // 필수
	Label  string `json:"label"`                 // 필수
	MsgTxt string `json:"messageText,omitempty"` // MsgTxt가 있으면 이게 발화로 나감, 없으면 Label이 발화

}

// LinkButton action="webLink"가 기본 값인 버튼입니다.
type LinkButton struct {
	Action  string `json:"action"`                // 필수
	Label   string `json:"label"`                 // 필수
	WebLink string `json:"webLinkUrl"`            // 사용할 http 주소
	MsgTxt  string `json:"messageText,omitempty"` // MsgTxt가 있으면 이게 발화로 나감, 없으면 Label이 발화
}

// MsgButton action="message"가 기본 값인 버튼입니다. (단순 메시지 버튼)
type MsgButton struct {
	Label   string `json:"label"`                 // 버튼 14자(가로배열 2개 8자) 필수
	Action  string `json:"action"`                // 필수
	MsgTxt  string `json:"messageText,omitempty"` // MsgTxt가 있으면 이게 발화로 나감, 없으면 Label이 발화
	BlockID string `json:"blockId,omitempty"`
}

// CallButton action="phone"이 기본 값인 버튼입니다.
//
// MsgTxt를 지정하지 않으면 Label이 발화됩니다.
type CallButton struct {
	Label       string `json:"label"`                 // 버튼 14자(가로배열 2개 8자) 필수
	Action      string `json:"action"`                // phone 필수
	PhoneNumber string `json:"phoneNumber"`           // 필수
	MsgTxt      string `json:"messageText,omitempty"` // MsgTxt가 있으면 이게 발화로 나감, 없으면 Label이 발화
	BlockID     string `json:"blockId,omitempty"`
}

// * Buttons END

// * Items START

// ListItem 카드의 각각 아이템, Title은 필수 값입니다.
type ListItem struct {
	Image string `json:"imageUrl,omitempty"`
	Desc  string `json:"description,omitempty"`
	Title string `json:"title"`
}

// ListItemLink 카드의 각각 아이템 (+링크 값), Title, Link은 필수 값입니다.
type ListItemLink struct {
	Title string `json:"title"`                 // items에 들어가는 경우, 해당 항목의 제목이 됩니다.
	Desc  string `json:"description,omitempty"` // items에 들어가는 경우, 해당 항목의 설명이 됩니다.
	Image string `json:"imageUrl,omitempty"`    // items에 들어가는 경우, 해당 항목의 우측 안내 사진이 됩니다.
	Link  Link   `json:"link"`                  // 	클릭시 작동하는 링크입니다.
}

// Link for ListItemLink
type Link struct {
	Link string `json:"web"`
}

// ContextControl 블록에서 생성한 outputContext의 lifeSpan, params 등을 제어할 수 있습니다.
type ContextControl struct {
	Values *Kakao `json:"values"`
}

// ContextValue for ContextControl
type ContextValue struct {
	// 수정하려는 output 컨텍스트의 이름
	//
	// 필수 여부 O
	Name string `json:"name"` //
	// 수정하려는 ouptut 컨텍스트의 lifeSpan
	//
	// 필수 여부 O
	LifeSpan int `json:"lifeSpan"`
	// output 컨텍스트에 저장하는 추가 데이터
	//
	// 필수 여부 X
	Params map[string]string `json:"params,omitempty"`
}

// * Items END

// Request 카카오 서버에서 POST JSON 데이터용
//
// 예제) 유저 발화문 얻기: kjson.UserRequest.Utterance
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
