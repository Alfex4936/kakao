package kakao

// Kakao = []interface{}
type Kakao []interface{}

// K = map[string]interface{}
type K map[string]interface{}

// Add 모든 객체를 Kakao 객체에 담습니다.
func (k *Kakao) Add(s ...interface{}) {
	for _, inter := range s {
		*k = append(*k, inter)
	}
}
