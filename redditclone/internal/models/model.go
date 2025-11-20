package models

type Models struct {
	UserMemory *UserMemory
	PostMemory *PostMemory
}

func InitModels(userMemory *UserMemory, postMemory *PostMemory) *Models {
	return &Models{
		UserMemory: userMemory,
		PostMemory: postMemory,
	}
}
