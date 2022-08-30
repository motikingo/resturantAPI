package menurepository

import (
	"reflect"
	"fmt"
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
	err:= itemRepo.db.Model(entity.Item{}).Preload("Ingridients").Find(&items).GetErrors()
	if len(err)>0{
		return nil,err
	}
	return items,nil

}

func(itemRepo *ItemRepository) Item(id uint)(*entity.Item,[]error){
	var item entity.Item
	err:= itemRepo.db.Model(&entity.Item{}).Preload("Catagories").Preload("Ingridients").First(&item,id).GetErrors()
	if len(err)>0{
		return nil,err
	}
	return &item,nil

}

func(itemRepo *ItemRepository) IsItemNameExist(name string)bool{
	var item entity.Item
	err:= itemRepo.db.Where("name = ?",name ).First(&item).GetErrors()
	
	return  len(err)== 0

}

func(itemRepo *ItemRepository) UpdateItem(it entity.Item)(*entity.Item,[]error){
	item,err := itemRepo.Item(it.ID)
	if len(err)>0{
		return nil,err
	}
	item.Name = func () string {
		if item.Name != it.Name{
			return it.Name
		}
		return item.Name
	}()
	 
	item.Price = func() float64 {
		if item.Price != it.Price{
			return it.Price
		}
		return item.Price
	}()

	item.Description = func () string {
		if item.Description != it.Description{
			return it.Description
		}
		return item.Description
	}()
	item.Image = func () string {
		if it.Image != item.Image{
			return it.Image
		}
		return item.Image
	}()

	item.Number = func () int {
		if it.Number != item.Number{
			return it.Number
		}
		return item.Number
	}()
	item.Ingridients = func () []entity.Ingredient {
		fmt.Println("out")
		
		if ! reflect.DeepEqual(item.Ingridients,it.Ingridients){
			fmt.Println("here")
			return it.Ingridients
		}
		return item.Ingridients
	}()
	item.Catagories = func () []entity.Catagory {
		if reflect.DeepEqual(it.Catagories,item.Catagories){
			return it.Catagories
		}
		return item.Catagories
	}()

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

func(itemRepo *ItemRepository) CreateItem(item entity.Item)(*entity.Item,[]error){

	err:= itemRepo.db.Create(&item).GetErrors()
	if len(err)>0{
		return nil,err
	}
	return &item,nil

}

