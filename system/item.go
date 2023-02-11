package system

import "sort"

type Item struct {
	Name   string
	Price  float64
	Rating int
}

func (i *Item) GiveRating(rating int) {
	i.Rating = rating
}

type ItemList []Item

func (l ItemList) Len() int {
	return len(l)
}

func (l ItemList) Less(i, j int) bool {
	return l[i].Price < l[j].Price
}

func (l ItemList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l ItemList) SortByPrice() {
	sort.Sort(l)
}

func (l ItemList) FilterByPrice(minPrice, maxPrice float64) ItemList {
	filteredList := make(ItemList, 0)
	for _, item := range l {
		if item.Price >= minPrice && item.Price <= maxPrice {
			filteredList = append(filteredList, item)
		}
	}
	return filteredList
}

func (l ItemList) FilterByRating(minRating, maxRating int) ItemList {
	filteredList := make(ItemList, 0)
	for _, item := range l {
		if item.Rating >= minRating && item.Rating <= maxRating {
			filteredList = append(filteredList, item)
		}
	}
	return filteredList
}

func (l ItemList) SearchByName(name string) ItemList {
	searchResult := make(ItemList, 0)
	for _, item := range l {
		if item.Name == name {
			searchResult = append(searchResult, item)
		}
	}
	return searchResult
}
