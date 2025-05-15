package dto

import "SplitSystemShop/internal/models"

type CatalogResponse struct {
	Total int           `json:"total"`
	Items []catalogItem `json:"items"`
}

type catalogItem struct {
	models.SplitSystem
	InCart      bool `json:"in_cart"`
	InFavorites bool `json:"in_favorites"`
}

func (r *CatalogResponse) New(userCart, userFavorites, allSystems []models.SplitSystem) {
	systemsInCartIDs := make(map[uint]struct{}, len(userCart))
	for _, inCartSystem := range userCart {
		systemsInCartIDs[inCartSystem.ID] = struct{}{}
	}

	systemsInFavoritesIDs := make(map[uint]struct{}, len(userCart))
	for _, system := range userFavorites {
		systemsInFavoritesIDs[system.ID] = struct{}{}
	}

	r.Total = len(allSystems)
	r.Items = make([]catalogItem, len(allSystems))

	for i, system := range allSystems {
		_, inCart := systemsInCartIDs[system.ID]
		_, inFavorites := systemsInFavoritesIDs[system.ID]
		r.Items[i] = catalogItem{
			SplitSystem: system,
			InCart:      inCart,
			InFavorites: inFavorites,
		}
	}
}

type CartModuleResponse struct {
	Cart struct {
		Total int
		Items []catalogItem
	}
	Favorites struct {
		Total int
		Items []catalogItem
	}
}

func NewCartModuleResponse(cart, favorites []models.SplitSystem) CartModuleResponse {
	cartIdx := make(map[uint]struct{})
	favoritesIdx := make(map[uint]struct{})

	for _, faveItem := range favorites {
		favoritesIdx[faveItem.ID] = struct{}{}
	}

	cartDto := make([]catalogItem, len(cart))
	for i, _ := range cart {
		cartIdx[cart[i].ID] = struct{}{}
		_, inFave := favoritesIdx[cart[i].ID]
		cartDto[i] = catalogItem{
			SplitSystem: cart[i],
			InCart:      true,
			InFavorites: inFave,
		}
	}

	favoritesDto := make([]catalogItem, len(favorites))
	for i, _ := range favorites {
		_, inCart := cartIdx[favorites[i].ID]
		favoritesDto[i] = catalogItem{
			SplitSystem: favorites[i],
			InCart:      inCart,
			InFavorites: true,
		}
	}

	return CartModuleResponse{
		Cart: struct {
			Total int
			Items []catalogItem
		}{
			Total: len(cart),
			Items: cartDto,
		},
		Favorites: struct {
			Total int
			Items []catalogItem
		}{
			Total: len(favorites),
			Items: favoritesDto,
		},
	}
}
