package menurepository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/motikingo/resturant-api/entity"
	"github.com/motikingo/resturant-api/menu"
)

type CatagoryGormRepository struct{
	db *gorm.DB
}

func NewCatagoryGormRepository(db *gorm.DB) menu.CatagoryRepo{
	return &CatagoryGormRepository{db: db}
}

func(cat *CatagoryGormRepository) Catagories()([]entity.Catagory,[]error){
	catagories:=[]entity.Catagory{}
	err := cat.db.Find(&catagories).GetErrors()
	if len(err)>0{
		return nil,err
	}

	return catagories, nil

}

func(cat *CatagoryGormRepository)Catagory(id uint)(*entity.Catagory,[]error){

	var catagory entity.Catagory
	err:= cat.db.First(&catagory,id).GetErrors()
	if len(err)>0{
		return nil,err
	}
	return &catagory,nil

}

func(cat *CatagoryGormRepository)IsCatagoryNameExist(name string)bool{
	var catagory entity.Catagory
	ers := cat.db.First(&catagory,name).GetErrors()

	return ers == nil

}
func(cat *CatagoryGormRepository)UpdateCatagory(cata entity.Catagory)(*entity.Catagory,[]error){
	catagory,err:= cat.Catagory(cata.ID)
	if len(err)>0{
		return nil,err
	}
	
	catagory.Name = func ()string  {
		if catagory.Name != cata.Name{
			return cata.Name
		}
		return catagory.Name
	}()

	catagory.Items = func () []string  {
		if len(cata.Items) != len(catagory.Items) {
			return cata.Items
		}
		return catagory.Items
	}()

	err = cat.db.Save(&catagory).GetErrors()
	if len(err)>0{
		return nil,err
	}

	return catagory,nil

}
func(cat *CatagoryGormRepository)DeleteCatagory(id uint)(*entity.Catagory,[]error){

	catagory,err:= cat.Catagory(id)
	if len(err)>0{
		return nil,err
	}
	err = cat.db.Delete(&catagory).GetErrors()
	if len(err)>0{
		return nil,err
	}
	return catagory,nil
}
func(cat *CatagoryGormRepository)CreateCatagory(catagory entity.Catagory)(*entity.Catagory,[]error){
	cata:= catagory
	err := cat.db.Create(&cata).GetErrors() 

	if len(err)>0{
		fmt.Println("its here")
		return nil,err
	}
	return &cata,nil

}

