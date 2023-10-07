package data

import (
	"encoding/json"
	"fmt"
	"time"
)

type Accessories struct {
	ID        int64     `json:"ID"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year"`
	Runtime   Runtime   `json:"runtime,omitempty"`
	Version   int32     `json:"version"`
}

func (a Accessories) MarshalJSON() ([]byte, error) {

	var runtime string
	if a.Runtime != 0 {
		runtime = fmt.Sprintf("%d mins", a.Runtime)
	}
	type AccessoryAlias Accessories

	aux := struct {
		AccessoryAlias
		Runtime string `json:"runtime,omitempty"`
	}{
		AccessoryAlias: AccessoryAlias(a),
		Runtime:        runtime,
	}
	return json.Marshal(aux)
}
