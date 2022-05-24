package serialize

import "fmt"

var (
	ErrSerializeResultNilDecoder = fmt.Errorf("Serialize decoder was nil.")
	ErrSerializeResultNilEncoder = fmt.Errorf("Serialize encoder was nil.")
)
