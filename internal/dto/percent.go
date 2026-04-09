package dto

import (
	"encoding/json"
	"fmt"
)

type Percent float64

func (p Percent) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("%.2f%%", float64(p))
	return json.Marshal(formatted)
}
