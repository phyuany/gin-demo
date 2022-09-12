package controller

import (
	"github.com/gin-gonic/gin"
	"jkdev.cn/api/model"
	"jkdev.cn/api/repository"
	"jkdev.cn/api/response"
	"jkdev.cn/api/vo"
	"strconv"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	Repository repository.CategoryRepository
}

func NewCategoryController() ICategoryController {
	repository := repository.NewCategoryRepository()
	repository.DB.AutoMigrate(model.Category{})

	return CategoryController{Repository: repository}
}

func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory vo.CreateCategoryRequest

	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Fail(ctx, "数据验证错误，分类名称必填", nil)
		return
	}

	category, err := c.Repository.Create(requestCategory.Name)
	if err != nil {
		panic(err)
		return
	}

	response.Success(ctx, gin.H{"category": category}, "")
}

func (c CategoryController) Update(ctx *gin.Context) {
	// 绑定body中的参数
	var requestCategory vo.CreateCategoryRequest

	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Fail(ctx, "数据验证错误，分类名称必填", nil)
		return
	}

	// 获取path中的参数,将字符串强转成int类型
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	updateCategory, err := c.Repository.SelectById(categoryId)
	if err != nil {
		response.Fail(ctx, "分类不存在", nil)
		return
	}

	// 更新分类
	// 参数类型 map / struct / name value
	category, err := c.Repository.Update(*updateCategory, requestCategory.Name)
	if err != nil {
		panic(err)
	}

	response.Success(ctx, gin.H{"category": category}, "修改成功")
}

func (c CategoryController) Show(ctx *gin.Context) {
	// 获取path中的参数,将字符串强转成int类型
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	category, err := c.Repository.SelectById(categoryId)
	if err != nil {
		response.Fail(ctx, "分类不存在", nil)
		return
	}

	response.Success(ctx, gin.H{"category": category}, "")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	// 获取path中的参数,将字符串强转成int类型
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	if err := c.Repository.DeleteById(categoryId); err != nil {
		response.Fail(ctx, "删除失败，请重试", nil)
		return
	}
	response.Success(ctx, nil, "")
}
