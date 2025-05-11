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
