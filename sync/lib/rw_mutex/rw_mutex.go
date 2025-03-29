package lib

import (
	"fmt"
	"sync"
	"time"
)

// sync.RWMutex (読み取り書き込みミューテックス)
//
// RWMutex は、複数のゴルーチンがデータを読み取ることができる一方で、書き込みは1つのゴルーチンのみが行えるように制御するミューテックスです。
//
// 読取り専用の操作が多い場合に、Mutex より効率的です。

// 読み取りと書き込みそれぞれからアクセスされる共有リソース（グローバル変数）
var data int

// 読み取りと書き込みの排他制御を行う。
var mu sync.RWMutex

// 読み取り操作
func readData() int {
	mu.RLock()         // 読み取りロックを取得（並行読み取りが可能）
	defer mu.RUnlock() // 関数終了時にロックを解除
	return data
}

// 書き込み操作
func writeData(value int) {
	mu.Lock()         // 書き込みロックを取得（書き込み中は読み取り禁止=RLockも取得不可）
	defer mu.Unlock() // 関数終了時にロックを解除
	data = value
}

// RWMutexProcess:
//
// sync.RWMutex を使用して並行処理における読み取りと書き込みの挙動をシミュレートします。
// RWMutexにより、複数のゴルーチンが同時にデータを読み取ることは可能だが、書き込みは1つのゴルーチンのみが行えるように制御されます。
//
// この関数では、10組のgoroutine（Reader/Writer）を作成し、それぞれ以下の挙動になることを期待します。
//
// - Reader（読み取り担当）は、readData 関数を呼び出し、共有データの同時読み取りを行う。
//
// - Writer（書き込み担当）は、writeData 関数を呼び出し、共有データの書き込みを行う。
//
// 各ゴルーチンの実行タイミングを分散させるため、time.Sleep を使用して、読み取りは 5ms 間隔、書き込みは 10ms 間隔で実行されるように調整しています。
//
// sync.WaitGroup を用いて、すべてのゴルーチンが終了するまで待機し、最後に最終的なデータの値を出力する。
func RWMutexProcess() {
	var wg sync.WaitGroup

	// Simultaneous reads and writes
	for i := range make([]struct{}, 10) {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Duration(i*5) * time.Millisecond) // Distribute read execution timing
			fmt.Println("Data:", readData())
		}(i)

		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Duration(i*10) * time.Millisecond) // Distribute write execution timing
			fmt.Println("書き込み中です:", i*10)
			writeData(i * 10)
		}(i)
	}

	wg.Wait()
	fmt.Println("================================================================")
	fmt.Println("Final Data:", readData())
}

/* @プログラムの挙動
go run cmd/main.go
Data: 0
書き込み中です: 0
Data: 0
Data: 0
書き込み中です: 10
Data: 10
Data: 10
書き込み中です: 20
Data: 20
Data: 20
書き込み中です: 30
Data: 30
Data: 30
書き込み中です: 40
Data: 40
書き込み中です: 50
書き込み中です: 60
書き込み中です: 70
書き込み中です: 80
書き込み中です: 90
================================================================
Final Data: 90
*/
