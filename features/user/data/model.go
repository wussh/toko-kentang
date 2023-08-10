package data

import (
	c "github.com/wussh/tokokentang/features/cart/data"
	p "github.com/wussh/tokokentang/features/product/data"
	t "github.com/wussh/tokokentang/features/transaction/data"

	// t "github.com/wussh/tokokentang/features/transaction/data"
	"github.com/wussh/tokokentang/features/user"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Avatar      string
	Name        string
	Email       string
	Address     string
	Password    string
	Cartss      []c.Carts        `gorm:"foreignKey:UserID"`
	Product     []p.Products     `gorm:"foreignKey:UserID"`
	Transaction []t.Transactions `gorm:"foreignKey:UserID"`
}

func ToCore(data Users) user.Core {
	return user.Core{
		ID:       data.ID,
		Avatar:   data.Avatar,
		Name:     data.Name,
		Email:    data.Email,
		Address:  data.Address,
		Password: data.Password,
	}
}

func CoreToData(data user.Core) Users {
	return Users{
		Model:    gorm.Model{ID: data.ID},
		Avatar:   data.Avatar,
		Name:     data.Name,
		Email:    data.Email,
		Address:  data.Address,
		Password: data.Password,
	}
}
