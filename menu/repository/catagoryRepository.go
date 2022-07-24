package menurepository

import(
	"github.com/motikingo/resturant-api/entity"
	"github.com/motikingo/resturant-api/menu"
	"github.com/jinzhu/gorm"
)

type CatagoryGormRepository struct{
	db *gorm.DB
}

func NewCommentGormRepository(db *gorm.DB) menu.CatagoryRepo{
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
func(cat *CatagoryGormRepository)UpdateCatagory(id uint)(*entity.Catagory,[]error){
	catagory,err:= cat.Catagory(id)
	if len(err)>0{
		return nil,err
	}
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
		return nil,err
	}
	return &cata,nil

}

