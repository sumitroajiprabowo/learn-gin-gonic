package example

import "github.com/gin-gonic/gin"

type Adress struct {
	Street string `form:"street" json:"street" binding:"required"`
	City   string `form:"city" json:"city" binding:"required"`
	Zip    string `form:"zip" json:"zip" binding:"required"`
}

type Teacher struct {
	Name         string `form:"name"`
	Age          int    `form:"age"`
	NestedStruct Adress
}

type Student struct {
	NestedStructPointer *Adress
	Name                string `form:"name"`
}

type Employee struct {
	NestedAnonyStruct struct {
		City string `form:"city" json:"city"`
	}
	Name string `form:"name"`
}

func GetBindRequestNestedStruct(c *gin.Context) {
	b := Teacher{
		Name: "Danu Budi Raharjo",
		Age:  23,
		NestedStruct: Adress{
			Street: "Jl. Raya",
			City:   "Jakarta",
			Zip:    "12345",
		},
	}
	c.Bind(&b)
	c.JSON(200, gin.H{
		"address": b.NestedStruct,
		"name":    b.Name,
		"age":     b.Age,
	})
}

func GetBindRequestNestedStructPointer(c *gin.Context) {
	b := Student{
		NestedStructPointer: &Adress{
			Street: "Jl. Raya",
			City:   "Jakarta",
			Zip:    "12345",
		},
		Name: "Danu Budi Raharjo",
	}
	c.Bind(&b)
	c.JSON(200, gin.H{
		"addresses": b.NestedStructPointer,
		"name":      b.Name,
	})
}

func GetBindRequestNestedAnonyStruct(c *gin.Context) {
	b := Employee{
		Name: "Danu Budi Raharjo",
		NestedAnonyStruct: struct {
			City string `form:"city" json:"city"`
		}{
			City: "Jakarta",
		},
	}
	c.Bind(&b)
	c.JSON(200, gin.H{
		"addresses": b.NestedAnonyStruct,
		"name":      b.Name,
	})
}
