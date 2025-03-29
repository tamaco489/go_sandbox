package lib

import (
	"fmt"

	"github.com/samber/lo"
	"github.com/tamaco489/go_sandbox/lo/model"
)

// NewPlayerSlicerByLoSliceToMap: lo.SliceToMap 関数を使って、Userポインターのスライスを、
// 各Userの ProviderUID をキーとし、そのUserポインターを値とするマップに変換する処理を示す関数です。
//
// 重複するキーが存在する場合、後に出現した要素が前の要素を上書きすることに注意してください。
func NewPlayerSlicerByLoSliceToMap() {

	// Userのサンプルデータをコンストラクタを介して作成
	users := []*model.User{
		model.NewUser(1, "Alice", "alice@example.com", "admin", "active", "provider-001"),
		model.NewUser(2, "Bob", "bob@example.com", "user", "inactive", "provider-002"),
		model.NewUser(3, "Charlie", "charlie@example.com", "user", "active", "provider-003"),
		model.NewUser(4, "Mike", "mike@example.com", "admin", "active", "provider-004"),
		model.NewUser(5, "Alice_Duplicate", "alice.dup@example.com", "admin", "inactive", "provider-001"), // NOTE: 注意: このUserはID=1のUserと重複する`provider_id`を持っています。
	}

	// Userのスライスをprovider_idをkeyにして変換、重複しているものは除外されるため、指定したkeyで一意性を担保したスライスが出来上がる
	// NOTE: 但し、振る舞いとしては重複を除外しているのではなく後から処理するもので上書きしていることに注意
	userMap := lo.SliceToMap(users, func(user *model.User) (string, *model.User) {
		return user.ProviderUID, user
	})

	for pid, user := range userMap {
		fmt.Printf("provider_uid: %s uid: %d, name: %s, email: %s, role: %s, status: %s\n",
			pid, user.ID, user.Name, user.Email, user.Role, user.Status)
	}

	/* id=1と5は、provider_uidが重複しているため、id=5が上書きしてリストに格納される
	provider_uid: provider-001 uid: 5, name: Alice_Duplicate, email: alice.dup@example.com, role: admin, status: inactive
	provider_uid: provider-002 uid: 2, name: Bob, email: bob@example.com, role: user, status: inactive
	provider_uid: provider-003 uid: 3, name: Charlie, email: charlie@example.com, role: user, status: active
	provider_uid: provider-004 uid: 4, name: Mike, email: mike@example.com, role: admin, status: active
	*/
}
