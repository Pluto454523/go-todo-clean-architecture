package user_repository

import (
	"github.com/pluto454523/go-todo-list/src/entity/user"
	"github.com/pluto454523/go-todo-list/src/repository/migrations"
	"github.com/pluto454523/go-todo-list/src/usecases/repository"
	"log"

	"gorm.io/gorm"
)

type userRepositoryDependency struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {

	// Create a new migration instance
	migrator := migrations.NewMigration(db, &userCollectionSchema{})

	// Start the migration
	err := migrator.Start()
	if err != nil {
		log.Fatal("error: migration failed userRepository")
	}

	return &userRepositoryDependency{
		DB: db,
	}
}

func (repo userRepositoryDependency) GetUserByID(id uint) (u user.UserEnity, err error) {

	userData := userCollectionSchema{}
	repo.DB.First(&userData, id)
	//fmt.Printf("%#v", userData)

	return user.UserEnity{}, nil
}
