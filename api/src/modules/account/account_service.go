package account

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"simple_sns_api/src/db"
	"simple_sns_api/src/ent"
	"simple_sns_api/src/ent/user"
	"simple_sns_api/src/lib/auth"
)

type AccountService struct{}

type RegisterParams struct {
	Name     string
	Email    string
	Password string
}

type UpdateParams struct {
	Name  string
	Email string
}

func (as AccountService) Find(ctx context.Context, userId int) (*ent.User, error) {
	return db.Client.User.Query().Where(user.ID(userId)).First(ctx)
}

func (as AccountService) Register(ctx context.Context, params RegisterParams) (*ent.User, auth.AuthToken, error) {
	encrypted, err := auth.EncryptPassword(params.Password)
	if err != nil {
		return nil, "", err
	}
	user, err := db.Client.User.Create().
		SetName(params.Name).
		SetEmail(params.Email).
		SetPassword(string(encrypted)).
		Save(ctx)
	if err != nil {
		return nil, "", err
	}
	authToken, err := auth.MakeAuthToken(user.ID)
	if err != nil {
		return nil, "", err
	}

	return user, authToken, nil
}

func (as AccountService) Update(ctx context.Context, userId int, params UpdateParams) (*ent.User, error) {
	qb := db.Client.User.UpdateOneID(userId)
	if params.Name != "" {
		qb = qb.SetName(params.Name)
	}
	if params.Email != "" {
		qb = qb.SetEmail(params.Email)
	}
	err := qb.Exec(ctx)
	if err != nil {
		return nil, err
	}
	return as.Find(ctx, userId)
}

func (as AccountService) UploadImage(ctx context.Context, userId int, file *os.File) (*ent.User, error) {
	// ファイルの作成
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	dst, err := os.Create(fmt.Sprintf("%s/../../../uploads/%s", wd, filepath.Base(file.Name())))
	defer dst.Close()
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadFile(file.Name())
	if err != nil {
		panic(err)
	}
	_, err = dst.Write(bytes)
	if err != nil {
		return nil, err
	}

	// TODO: URLを動的に
	filePath := "http://localhost:8080/uploads/" + filepath.Base(dst.Name())

	err = db.Client.User.UpdateOneID(userId).SetIconImageUrl(filePath).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return as.Find(ctx, userId)
}
