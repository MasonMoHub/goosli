package slicers

import (
	. "github.com/l1va/goosli/primitives"
	"strconv"
)

type Settings struct {
	DateTime            string
	Epsilon             float64
	LayerHeight         float64
	WallThickness       float64
	FillDensity         int
	BedTemperature      int
	ExtruderTemperature int
	PrintSpeed          int
	Nozzle              float64
	LayerCount          int
	RotationCenterZ     float64
	PlanesFile          string
}

func (s *Settings) ToMap() map[string]string {
	return map[string]string{
		"{datetime}":             s.DateTime,
		"{layer_height}":         StrF(s.LayerHeight),
		"{wall_thickness}":       StrF(s.WallThickness),
		"{fill_density}":         strconv.Itoa(s.FillDensity),
		"{bed_temperature}":      strconv.Itoa(s.BedTemperature),
		"{extruder_temperature}": strconv.Itoa(s.ExtruderTemperature),
		"{print_speed}":          strconv.Itoa(s.PrintSpeed),
		"{nozzle}":               StrF(s.Nozzle),
		"{layer_count}":          strconv.Itoa(s.LayerCount),
	}
}
