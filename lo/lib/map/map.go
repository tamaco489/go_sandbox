package lib

import (
	"log"

	"github.com/samber/lo"
	"github.com/tamaco489/go_sandbox/lo/model"
)

// NewUserSlicerByLoMap: ユーザーのサンプルデータを生成し、lo.Mapを使用しそれをUser型のスライスに変換する関数です。
//
// 各UserはNewUserコンストラクタを通じて作成されます。
func NewUserSlicerByLoMap() []*model.User {

	// Userのサンプルデータ（例えばsqlc等で生成された構造体がこれにあたる）
	fetchUserList := []struct {
		ID          int64
		Name        string
		Email       string
		Role        string
		Status      string
		ProviderUID string
	}{
		{1, "john_doe", "john@example.com", "admin", "active", "provider-123"},
		{2, "jane_doe", "jane@example.com", "user", "inactive", "provider-456"},
		{3, "alice_smith", "alice@example.com", "admin", "active", "provider-789"},
	}

	// lo.Mapを使って、userDataの各要素をmodel.NewUserを使ってUser型に変換する
	// 変換関数は、userDataのフィールドをmodel.NewUserに渡して、新しいUserインスタンスを作成する
	users := lo.Map(fetchUserList, func(u struct {
		ID          int64
		Name        string
		Email       string
		Role        string
		Status      string
		ProviderUID string
	}, _ int) *model.User {
		return model.NewUser(u.ID, u.Name, u.Email, u.Role, u.Status, u.ProviderUID)
	})

	for i, v := range users {
		log.Println("idx:", i, "value:", v)
	}

	// 2025/03/29 10:27:13 idx: 0 value: &{1 john_doe john@example.com admin active provider-123 2025-03-29 10:27:13.222051971 +0900 JST m=+0.000143245 2025-03-29 10:27:13.222051971 +0900 JST m=+0.000143245}
	// 2025/03/29 10:27:13 idx: 1 value: &{2 jane_doe jane@example.com user inactive provider-456 2025-03-29 10:27:13.222052593 +0900 JST m=+0.000143870 2025-03-29 10:27:13.222052593 +0900 JST m=+0.000143870}
	// 2025/03/29 10:27:13 idx: 2 value: &{3 alice_smith alice@example.com admin active provider-789 2025-03-29 10:27:13.222052677 +0900 JST m=+0.000143945 2025-03-29 10:27:13.222052677 +0900 JST m=+0.000143945}

	return users
}
