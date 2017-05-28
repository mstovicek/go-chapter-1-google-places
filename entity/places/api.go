package places

type Api interface {
	GetPlace(placeId string) Place
}
