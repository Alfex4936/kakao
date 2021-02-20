package kakao

// Kakao is []interface{}
type Kakao []interface{}

// K is map[string]interface{}
type K map[string]interface{}

// Add all interfaces into Kakao object
func (k *Kakao) Add(s ...interface{}) {
	for _, inter := range s {
		*k = append(*k, inter)
	}
}
