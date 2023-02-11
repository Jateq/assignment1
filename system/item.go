package system

import (
	"sort"
)

type Item struct {
	Name   string
	Price  float64
	Rating float64
}

type ItemList []*Item

type ItemRepository struct {
	orders []*Item
}

func (ir ItemRepository) Len() int {
	return len(ir.orders)
}

func (ir ItemRepository) Less(i, j int) bool {
	return ir.orders[i].Price < ir.orders[j].Price
}

func (ir ItemRepository) Swap(i, j int) {
	ir.orders[i], ir.orders[j] = ir.orders[j], ir.orders[i]
}

func (ir *ItemRepository) FilterItems(price float64, rating float64) ItemList {
	var items ItemList
	sort.Sort(ir)
	for _, item := range ir.orders {
		if item.Price >= price && item.Rating >= rating {
			items = append(items, item)
		}
	}
	return items
}func (i *Item) GiveRating(rating float64) {
	i.Rating = rating
}

func (ir *ItemRepository) AddItem(item *Item) {
	ir.orders = append(ir.orders, item)
}

func (ir *ItemRepository) SearchItems(name string) ItemList {
	var items ItemList
	for _, item := range ir.orders {
		if item.Name == name {
			items = append(items, item)
		}
	}
	return items
}
