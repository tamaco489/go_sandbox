package lib

import (
	"fmt"
	"sort"

	"github.com/samber/lo"
	"github.com/tamaco489/go_sandbox/lo/model"
)

func NewUserSlicerByLoGroupBy() {

	// サンプルデータを作成
	users, players := createSampleData()

	// ユーザーをProviderUIDでグループ化
	grouped := groupUsersByProviderUID(users)

	// プレイヤーがいるユーザーといないユーザーに分ける
	var withPlayers []*model.User
	var withoutPlayers []*model.User
	for _, group := range grouped {
		with, without := separateUsersByPlayerExistence(group, players)
		withPlayers = append(withPlayers, with...)
		withoutPlayers = append(withoutPlayers, without...)
	}

	// ユーザーをID順にソート
	withPlayers = sortUsersByID(withPlayers)
	withoutPlayers = sortUsersByID(withoutPlayers)

	// 最終的なリストを作成
	finalList := createFinalUserList(withPlayers, withoutPlayers)

	// 結果を表示
	for _, user := range finalList {
		fmt.Printf("uid: %d, name: %s, provider_uid: %s\n", user.ID, user.Name, user.ProviderUID)
	}
}

// 最終的なリストを作成（プレイヤーがいるユーザーを先に表示、プレイヤーがいないユーザーを後に表示）
func createFinalUserList(withPlayers, withoutPlayers []*model.User) []*model.User {
	// プレイヤーがいるユーザーを先に表示し、プレイヤーがいないユーザーを後に表示
	return append(withPlayers, withoutPlayers...)
}

// ユーザーをID順にソート（Go標準パッケージのsortを使用）
func sortUsersByID(users []*model.User) []*model.User {
	// sort.SliceでID順に並べ替える
	sort.Slice(users, func(i, j int) bool {
		return users[i].ID < users[j].ID
	})
	return users
}

// プレイヤーがいるユーザーといないユーザーを分ける
func separateUsersByPlayerExistence(users []*model.User, players []*model.Player) ([]*model.User, []*model.User) {
	// プレイヤーをMapに変換（効率的に存在チェックするため）
	playerMap := make(map[string]bool)
	for _, p := range players {
		playerMap[p.ProviderUID] = true
	}

	var withPlayers []*model.User
	var withoutPlayers []*model.User

	// ユーザーごとにプレイヤーの存在をチェック
	for _, user := range users {
		if playerMap[user.ProviderUID] {
			withPlayers = append(withPlayers, user)
		} else {
			withoutPlayers = append(withoutPlayers, user)
		}
	}

	return withPlayers, withoutPlayers
}

// ユーザーをProviderUIDでグループ化
func groupUsersByProviderUID(users []*model.User) map[string][]*model.User {
	// lo.GroupByを使って、ProviderUIDをキーにしてユーザーをグループ化
	return lo.GroupBy(users, func(user *model.User) string {
		return user.ProviderUID
	})
}

// サンプルデータ作成
func createSampleData() ([]*model.User, []*model.Player) {
	// ユーザーリストの作成
	users := []*model.User{
		model.NewUser(1, "User1", "user1@example.com", "admin", "active", "provide_uid_1"),
		model.NewUser(2, "User2", "user2@example.com", "user", "inactive", "provide_uid_2"),
		model.NewUser(3, "User3", "user3@example.com", "admin", "active", "provide_uid_3"),  // NOTE: プレイヤー側は存在しない
		model.NewUser(4, "User4", "user4@example.com", "user", "inactive", "provide_uid_4"), // NOTE: プレイヤー側は存在しない
		model.NewUser(5, "User5", "user5@example.com", "admin", "active", "provide_uid_5"),
	}

	// プレイヤーリストの作成
	players := []*model.Player{
		model.NewPlayer(100, "provide_uid_1", 10, 200, 500, nil, nil, nil, nil),
		model.NewPlayer(101, "provide_uid_2", 15, 300, 1000, nil, nil, nil, nil),
		model.NewPlayer(101, "provide_uid_5", 20, 350, 1500, nil, nil, nil, nil),
	}

	return users, players
}
