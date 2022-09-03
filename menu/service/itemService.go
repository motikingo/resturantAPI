package menuService

import (
	"github.com/motikingo/resturant-api/entity"
	"github.com/motikingo/resturant-api/menu"
)

type ItemGormService struct {
	repo menu.ItemRepo
}

func NewItemGormService(repo menu.ItemRepo) menu.ItemService {

	return &ItemGormService{repo: repo}
}

func (itemRepo *ItemGormService) Items() ([]entity.Item, []error) {
	items, err := itemRepo.repo.Items()
	if len(err) > 0 {
		return nil, err
	}
	return items, nil

}

func (itemRepo *ItemGormService) Item(id uint) (*entity.Item, []error) {

	item, err := itemRepo.repo.Item(id)
	if len(err) > 0 {
		return nil, err
	}
	return item, nil

}

func (itemRepo *ItemGormService) IsItemNameExist(name string) bool {

	item := itemRepo.repo.IsItemNameExist(name)

	return item

}

func (itemRepo *ItemGormService) UpdateItem(it entity.Item) (*entity.Item, []error) {

	item, err := itemRepo.repo.UpdateItem(it)
	if len(err) > 0 {
		return nil, err
	}
	return item, nil

}

func (itemRepo *ItemGormService) DeleteItem(id uint) (*entity.Item, []error) {
	item, err := itemRepo.repo.DeleteItem(id)
	if len(err) > 0 {
		return nil, err
	}
	return item, nil

}

func (itemRepo *ItemGormService) CreateItem(item entity.Item) (*entity.Item, []error) {

	it, err := itemRepo.repo.CreateItem(item)
	if len(err) > 0 {
		return nil, err
	}
	return it, nil

}
