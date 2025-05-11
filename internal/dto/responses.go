package dto

import "SplitSystemShop/internal/models"

type CatalogResponse struct {
	Total int           `json:"total"`
	Items []catalogItem `json:"items"`
}

type catalogItem struct {
	models.SplitSystem
	InCart bool `json:"in_cart"`
}

func (r *CatalogResponse) New(userCart, allSystems []models.SplitSystem) {
	systemsInCartIDs := make(map[uint]struct{}, len(userCart))
	for _, inCartSystem := range userCart {
		systemsInCartIDs[inCartSystem.ID] = struct{}{}
	}

	r.Total = len(allSystems)
	r.Items = make([]catalogItem, len(allSystems))

	for i, system := range allSystems {
		_, inCart := systemsInCartIDs[system.ID]
		r.Items[i] = catalogItem{
			SplitSystem: system,
			InCart:      inCart,
		}
	}
}
