package repositories

type categoryRepositoryImp struct {
}

func NewCategoryRepository() CategoryRepository {
	return categoryRepositoryImp{}
}
