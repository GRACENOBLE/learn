package data

type distance float64 //miles
type distanceKm float64 //Kilometers

// Method
func (miles distance) ToKm() distanceKm{
	return distanceKm(1.60934 * miles)
}
func (Kilometers distanceKm) ToMiles() distance{
	return distance(1.60934 * Kilometers)
}

func Test(){
	d := distance(4.5)
	km := d.ToKm()
	miles := km.ToMiles()
	println(d)
	println(km)
	println(miles)
}