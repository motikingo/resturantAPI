package menurepository

import (
	"github.com/jinzhu/gorm"
	"github.com/motikingo/resturant-api/entity"
	"github.com/motikingo/resturant-api/menu"
)


type ItemRepository struct{
	db *gorm.DB
}

func NewItemRepository(db * gorm.DB) menu.ItemRepo{

	return &ItemRepository{db: db}
}

func(itemRepo *ItemRepository) Items()([]entity.Item,[]error){
	items:=[]entity.Item{}
	err:= itemRepo.db.Find(&items).GetErrors()
	if len(err)>0{
		return nil,err
	}
	return items,nil

}

func(itemRepo *ItemRepository) Item(id uint)(*entity.Item,[]error){
	var item entity.Item
	err:= itemRepo.db.First(&item,id).GetErrors()
	if len(err)>0{
		return nil,err
	}
	return &item,nil


}

func(itemRepo *ItemRepository) UpdateItem(id uint)(*entity.Item,[]error){
	item,err := itemRepo.Item(id)
	if len(err)>0{
		return nil,err
	}
	err = itemRepo.db.Save(&item).GetErrors()
	if len(err)>0{
		return nil,err
	}
	return item,nil


}

func(itemRepo *ItemRepository) DeleteItem(id uint)(*entity.Item,[]error){
	item,err := itemRepo.Item(id)
	if len(err)>0{
		return nil,err
	}
	err = itemRepo.db.Delete(&item,id).GetErrors()
	if len(err)>0{
		return nil,err
	}
	return item,nil

}

func(itemRepo *ItemRepository) CreateItem(item *entity.Item)(*entity.Item,[]error){

	it := item
	err:= itemRepo.db.Create(&it).GetErrors()
	if len(err)>0{
		return nil,err
	}
	return it,nil

}

