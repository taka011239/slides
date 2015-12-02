type Person struct {
	Name  string
	Likes []string
}
var people []*Person

likes := make(map[string][]*Person)
for _, p := range people {
	for _, l := range p.Likes {
		likes[l] = append(likes[l], p)
	}
}
