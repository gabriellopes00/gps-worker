package domain

type Position struct {
	Name      string  `json:"name,omitempty" valid:"-"`
	Latitude  float64 `json:"latitude,omitempty" valid:"latitude"`
	Longitude float64 `json:"longitude,omitempty" valid:"longitude"`
	Distance  float64 `json:"distance,omitempty" valid:"-"`
}
