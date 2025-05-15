package dto

import "SplitSystemShop/internal/models"

type CatalogResponse struct {
	Total int           `json:"total"`
	Items []CatalogItem `json:"items"`
}

type CatalogItem struct {
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
	r.Items = make([]CatalogItem, len(allSystems))

	for i, system := range allSystems {
		_, inCart := systemsInCartIDs[system.ID]
		_, inFavorites := systemsInFavoritesIDs[system.ID]
		r.Items[i] = CatalogItem{
			SplitSystem: system,
			InCart:      inCart,
			InFavorites: inFavorites,
		}
	}
}

type CartModuleResponse struct {
	Cart struct {
		Total int
		Items []CatalogItem
	}
	Favorites struct {
		Total int
		Items []CatalogItem
	}
}

func NewCartModuleResponse(cart, favorites []models.SplitSystem) CartModuleResponse {
	cartIdx := make(map[uint]struct{})
	favoritesIdx := make(map[uint]struct{})

	for _, faveItem := range favorites {
		favoritesIdx[faveItem.ID] = struct{}{}
	}

	cartDto := make([]CatalogItem, len(cart))
	for i, _ := range cart {
		cartIdx[cart[i].ID] = struct{}{}
		_, inFave := favoritesIdx[cart[i].ID]
		cartDto[i] = CatalogItem{
			SplitSystem: cart[i],
			InCart:      true,
			InFavorites: inFave,
		}
	}

	favoritesDto := make([]CatalogItem, len(favorites))
	for i, _ := range favorites {
		_, inCart := cartIdx[favorites[i].ID]
		favoritesDto[i] = CatalogItem{
			SplitSystem: favorites[i],
			InCart:      inCart,
			InFavorites: true,
		}
	}

	return CartModuleResponse{
		Cart: struct {
			Total int
			Items []CatalogItem
		}{
			Total: len(cart),
			Items: cartDto,
		},
		Favorites: struct {
			Total int
			Items []CatalogItem
		}{
			Total: len(favorites),
			Items: favoritesDto,
		},
	}
}
