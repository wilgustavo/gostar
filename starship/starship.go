package starship

// Starship contains the data of a Starship
type Starship struct {
	Name         string
	Model        string
	Manufacturer string
	Cost         string
	Length       string
}

// Deserializer get a list of Starships
type Deserializer interface {
	ListSharships() []Starship
}

// Serializer persist a list of Starships
type Serializer interface {
	SaveSharships(list []Starship) error
}

// ToSlice convert in an Slice of string
func (s Starship) ToSlice() []string {
	return []string{
		s.Name,
		s.Model,
		s.Manufacturer,
		s.Cost,
		s.Length,
	}
}
