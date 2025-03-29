package lib

import (
	"log"

	"github.com/samber/lo"
	"github.com/tamaco489/go_sandbox/lo/model"
)

func NewUserSlicerByLoUniqMap() {

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
		{2, "jane_doe", "jane@example.com", "user", "inactive", "provider-456"}, // 非アクティブなユーザ
		{3, "alice_smith", "alice@example.com", "admin", "active", "provider-789"},
		{4, "bob_marley", "bob@example.com", "user", "active", "provider-999"},
		{5, "john_doe", "john@example.com", "admin", "active", "provider-123"}, // provide_uidが重複しているデータ
	}

	// Step1: statusが"active"のユーザーのみを抽出（lo.Filterを使用）
	activeUsers := lo.Filter(fetchUserList, func(u struct {
		ID          int64
		Name        string
		Email       string
		Role        string
		Status      string
		ProviderUID string
	}, _ int) bool {
		return u.Status == "active"
	})

	// Step2: IDをキーにしてユニークなリストを作成（lo.UniqByを使用）
	uniqueUsers := lo.UniqBy(activeUsers, func(u struct {
		ID          int64
		Name        string
		Email       string
		Role        string
		Status      string
		ProviderUID string
	}) string {
		return u.ProviderUID // IDを基準に一意性を判定
	})

	// Step3: 重複を排除したリストをコンストラクタを使用して生成
	resultUsers := lo.Map(uniqueUsers, func(u struct {
		ID          int64
		Name        string
		Email       string
		Role        string
		Status      string
		ProviderUID string
	}, _ int) *model.User {
		return model.NewUser(u.ID, u.Name, u.Email, u.Role, u.Status, u.ProviderUID)
	})

	for _, v := range resultUsers {
		log.Println("value:", v)
	}

	/* [一意性を担保したリストを生成]
	- id:2: status=inactiveのためリストには含まれない
	- id:5: provider_uidが重複しているためリストには含まれない

	2025/03/29 12:02:50 value: &{1 john_doe john@example.com admin active provider-123 2025-03-29 12:02:50.752378344 +0900 JST m=+0.000122907 2025-03-29 12:02:50.752378344 +0900 JST m=+0.000122907}
	2025/03/29 12:02:50 value: &{3 alice_smith alice@example.com admin active provider-789 2025-03-29 12:02:50.75237867 +0900 JST m=+0.000123226 2025-03-29 12:02:50.75237867 +0900 JST m=+0.000123226}
	2025/03/29 12:02:50 value: &{4 bob_marley bob@example.com user active provider-999 2025-03-29 12:02:50.752378741 +0900 JST m=+0.000123297 2025-03-29 12:02:50.752378741 +0900 JST m=+0.000123297}
	*/
}
