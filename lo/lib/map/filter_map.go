package lib

import (
	"fmt"

	"github.com/samber/lo"
	"github.com/tamaco489/go_sandbox/lo/model"
)

// NewPlayerSlicerByLoFilterMap: ユーザーのリストからステータスが"active"のユーザーのEmailを抽出して出力する関数です。
//
// `lo.FilterMap`を使用して、ステータスが"active"のユーザーをフィルタし、そのEmailアドレスをマッピングします。
func NewPlayerSlicerByLoFilterMap() {

	// サンプルデータ（User）
	users := []*model.User{
		model.NewUser(1, "Alice", "alice@example.com", "admin", "active", "provider-001"),
		model.NewUser(2, "Bob", "bob@example.com", "user", "inactive", "provider-002"),
		model.NewUser(3, "Charlie", "charlie@example.com", "user", "active", "provider-003"),
		model.NewUser(4, "Mike", "mike@example.com", "admin", "active", "provider-004"),
		model.NewUser(5, "Alice_Duplicate_1", "alice.dup@example.com", "admin", "active", "provider-001"), // NOTE: 注意: このUserはID=1のUserと重複する`provider_id`を持っています。
		model.NewUser(6, "Alice_Duplicate_2", "alice.dup@example.com", "admin", "active", "provider-001"), // NOTE: 注意: このUserはID=1のUserと重複する`provider_id`を、ID=5のUserと重複する`email`を持っています。
	}

	// フィルタリング＆マッピング：statusが `active` のユーザーのIDとEmailを抽出
	activeUserEmailList := lo.FilterMap(users, func(user *model.User, index int) (string, bool) {
		if user.Status == "active" {
			return user.Email, true
		}
		return "", false
	})

	for i, email := range activeUserEmailList {
		fmt.Println("index:", i, "email:", email)
	}

	/* uid=2はstatusがinactiveのため、リストには含まれない（重複判定はしていないため、id=5,6のメールアドレス被りは含まれる）
	index: 0 email: alice@example.com
	index: 1 email: charlie@example.com
	index: 2 email: mike@example.com
	index: 3 email: alice.dup@example.com
	index: 4 email: alice.dup@example.com
	*/
}
