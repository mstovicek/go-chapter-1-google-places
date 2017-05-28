package places

type API interface {
	GetPlace(placeID string) (*Place, error)
}
