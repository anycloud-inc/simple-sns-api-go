package account

import (
	"encoding/base64"
	"os"
	"path/filepath"
	"simple_sns_api/src/db"
	"testing"
)

func TestUploadImage(t *testing.T) {
	ctx := db.CreateTestConnection(t)
	user, _, err := AccountService{}.Register(ctx, RegisterParams{
		Name:     "Hoge",
		Email:    "hoge@example.com",
		Password: "hogehoge",
	})
	if err != nil {
		panic(err)
	}

	// base64からテストファイルを作成してtmpフォルダに保存
	data, err := base64.StdEncoding.DecodeString("SGVsbG8gV29ybGQ=")
	if err != nil {
		panic(err)
	}
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	file, err := os.Create(filepath.Join(wd, "/../../../tmp", "test.jpg"))
	if err != nil {
		panic(err)
	}
	_, err = file.Write(data)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	user, err = AccountService{}.UploadImage(ctx, user.ID, file)

	if err != nil {
		t.Error(err)
	}
	if user.IconImageUrl != "http://localhost:8080/uploads/test.jpg" {
		t.Error("Invalid IconImageUrl: " + user.IconImageUrl)
	}
}
