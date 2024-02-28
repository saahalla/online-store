package dto

func (l CartItemDataList) GetCartItemIDList() (cartItemIDList []int) {

	mapCartID := map[int]struct{}{}

	for _, ci := range l {

		if _, exists := mapCartID[ci.ID]; !exists {
			cartItemIDList = append(cartItemIDList, ci.ID)

			mapCartID[ci.ID] = struct{}{}
		}

	}

	return
}

func (l CartItemDBList) GetProductIDList() (productIDList []int) {
	mapProductID := map[int]struct{}{}

	for _, ci := range l {

		if _, exists := mapProductID[ci.ID]; !exists {
			productIDList = append(productIDList, ci.ProductID)

			mapProductID[ci.ID] = struct{}{}
		}

	}

	return
}
