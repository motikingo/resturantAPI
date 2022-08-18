package menuService

import (
	"github.com/motikingo/resturant-api/entity"
	"github.com/motikingo/resturant-api/menu"
)

type CatagoryGormService struct{
	repo menu.CatagoryRepo
}

func NewCatagoryGormService(repo menu.CatagoryRepo) menu.CatagorySrv{
	return &CatagoryGormService{repo:repo}
}

func(cat *CatagoryGormService) Catagories()([]entity.Catagory,[]error){
	catagories,err := cat.repo.Catagories()
	if len(err)>0{
		return nil,err
	}

	return catagories, nil

}

func(cat *CatagoryGormService)Catagory(id uint)(*entity.Catagory,[]error){

	catagory,err:= cat.repo.Catagory(id)
	if len(err)>0{
		return nil,err
	}
	return catagory,nil

}
func(cat *CatagoryGormService)IsCatagoryNameExist(name string)bool{
	return cat.repo.IsCatagoryNameExist(name)
}

func(cat *CatagoryGormService)UpdateCatagory(ct entity.Catagory)(*entity.Catagory,[]error){
	
	catagory,err:= cat.repo.UpdateCatagory(ct) 
	if len(err)>0{
		return nil,err
	}

	return catagory,nil

}
func(cat *CatagoryGormService)DeleteCatagory(id uint)(*entity.Catagory,[]error){

	catagory,err:= cat.repo.DeleteCatagory(id) 
	if len(err)>0{
		return nil,err
	}
	
	return catagory,nil
}
func(cat *CatagoryGormService)CreateCatagory(catagory entity.Catagory)(*entity.Catagory,[]error){
	cata,err := cat.repo.CreateCatagory(catagory)

	if len(err)>0{
		return nil,err
	}
	return cata,nil

}

