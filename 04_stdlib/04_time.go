package main

import (
	"fmt"
	"time"
)

/*
【学習ポイント】
1. 時刻の取得と生成
2. 時刻のフォーマット
3. 時刻の計算
4. タイマーとTicker
5. 実用例
*/

func main() {
	// ========== 現在時刻の取得 ==========
	fmt.Println("=== 現在時刻の取得 ===")

	now := time.Now()
	fmt.Printf("現在時刻: %v\n", now)
	fmt.Printf("年: %d\n", now.Year())
	fmt.Printf("月: %d\n", now.Month())
	fmt.Printf("日: %d\n", now.Day())
	fmt.Printf("時: %d\n", now.Hour())
	fmt.Printf("分: %d\n", now.Minute())
	fmt.Printf("秒: %d\n", now.Second())
	fmt.Printf("曜日: %s\n", now.Weekday())

	// ========== 時刻の生成 ==========
	fmt.Println("\n=== 時刻の生成 ===")

	// 特定の日時を作成
	birthday := time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	fmt.Printf("誕生日: %v\n", birthday)

	// Unix時間から作成
	unixTime := time.Unix(1609459200, 0)
	fmt.Printf("Unix時間: %v\n", unixTime)

	// ========== 時刻のフォーマット ==========
	fmt.Println("\n=== 時刻のフォーマット ===")

	// 標準フォーマット
	fmt.Printf("RFC3339: %s\n", now.Format(time.RFC3339))
	fmt.Printf("RFC822: %s\n", now.Format(time.RFC822))

	// カスタムフォーマット（参照日時: 2006-01-02 15:04:05）
	fmt.Printf("YYYY-MM-DD: %s\n", now.Format("2006-01-02"))
	fmt.Printf("YYYY/MM/DD HH:MM:SS: %s\n", now.Format("2006/01/02 15:04:05"))
	fmt.Printf("日本語形式: %s\n", now.Format("2006年01月02日 15時04分05秒"))

	// ========== 文字列から時刻へ ==========
	fmt.Println("\n=== 文字列から時刻へ ===")

	dateStr := "2024-12-25 15:30:00"
	layout := "2006-01-02 15:04:05"
	parsedTime, err := time.Parse(layout, dateStr)
	if err != nil {
		fmt.Printf("パースエラー: %v\n", err)
	} else {
		fmt.Printf("パース結果: %v\n", parsedTime)
	}

	// タイムゾーン付きパース
	jst, _ := time.LoadLocation("Asia/Tokyo")
	parsedTimeJST, _ := time.ParseInLocation(layout, dateStr, jst)
	fmt.Printf("JST: %v\n", parsedTimeJST)

	// ========== 時刻の計算 ==========
	fmt.Println("\n=== 時刻の計算 ===")

	// 時刻の加算
	tomorrow := now.Add(24 * time.Hour)
	fmt.Printf("明日: %s\n", tomorrow.Format("2006-01-02"))

	nextWeek := now.AddDate(0, 0, 7)
	fmt.Printf("来週: %s\n", nextWeek.Format("2006-01-02"))

	nextMonth := now.AddDate(0, 1, 0)
	fmt.Printf("来月: %s\n", nextMonth.Format("2006-01-02"))

	nextYear := now.AddDate(1, 0, 0)
	fmt.Printf("来年: %s\n", nextYear.Format("2006-01-02"))

	// 時刻の減算
	yesterday := now.Add(-24 * time.Hour)
	fmt.Printf("昨日: %s\n", yesterday.Format("2006-01-02"))

	// ========== 時刻の差分 ==========
	fmt.Println("\n=== 時刻の差分 ===")

	start := time.Now()
	time.Sleep(100 * time.Millisecond)
	end := time.Now()

	duration := end.Sub(start)
	fmt.Printf("経過時間: %v\n", duration)
	fmt.Printf("ミリ秒: %d\n", duration.Milliseconds())
	fmt.Printf("秒: %.2f\n", duration.Seconds())

	// ========== 時刻の比較 ==========
	fmt.Println("\n=== 時刻の比較 ===")

	time1 := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	time2 := time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)

	fmt.Printf("time1 < time2: %t\n", time1.Before(time2))
	fmt.Printf("time1 > time2: %t\n", time1.After(time2))
	fmt.Printf("time1 == time1: %t\n", time1.Equal(time1))

	// ========== Duration（期間） ==========
	fmt.Println("\n=== Duration ===")

	d1 := 5 * time.Second
	d2 := 100 * time.Millisecond
	d3 := 2 * time.Minute

	fmt.Printf("5秒: %v\n", d1)
	fmt.Printf("100ミリ秒: %v\n", d2)
	fmt.Printf("2分: %v\n", d3)

	total := d1 + d2 + d3
	fmt.Printf("合計: %v\n", total)

	// ========== Timer（一度だけ実行） ==========
	fmt.Println("\n=== Timer ===")

	timer := time.NewTimer(1 * time.Second)
	fmt.Println("タイマー開始...")
	<-timer.C
	fmt.Println("タイマー終了!")

	// time.After を使った簡易版
	fmt.Println("1秒待機...")
	<-time.After(1 * time.Second)
	fmt.Println("完了!")

	// ========== Ticker（定期的に実行） ==========
	fmt.Println("\n=== Ticker ===")

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	count := 0
	for range ticker.C {
		count++
		fmt.Printf("Tick %d\n", count)
		if count >= 3 {
			break
		}
	}

	// ========== タイムゾーン ==========
	fmt.Println("\n=== タイムゾーン ===")

	// UTC
	utc := time.Now().UTC()
	fmt.Printf("UTC: %s\n", utc.Format("2006-01-02 15:04:05 MST"))

	// JST（日本標準時）
	jst2, _ := time.LoadLocation("Asia/Tokyo")
	jstTime := time.Now().In(jst2)
	fmt.Printf("JST: %s\n", jstTime.Format("2006-01-02 15:04:05 MST"))

	// EST（米国東部標準時）
	est, _ := time.LoadLocation("America/New_York")
	estTime := time.Now().In(est)
	fmt.Printf("EST: %s\n", estTime.Format("2006-01-02 15:04:05 MST"))

	// ========== 実用例：処理時間の計測 ==========
	fmt.Println("\n=== 実用例：処理時間の計測 ===")

	measureTime(func() {
		time.Sleep(200 * time.Millisecond)
		fmt.Println("処理実行中...")
	})

	// ========== 実用例：タイムアウト処理 ==========
	fmt.Println("\n=== 実用例：タイムアウト処理 ===")

	result := make(chan string, 1)

	go func() {
		time.Sleep(500 * time.Millisecond)
		result <- "処理完了"
	}()

	select {
	case res := <-result:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("タイムアウト")
	}

	// ========== 実用例：定期実行 ==========
	fmt.Println("\n=== 実用例：定期実行 ===")

	ticker2 := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		count := 0
		for {
			select {
			case <-ticker2.C:
				count++
				fmt.Printf("定期処理 %d 実行\n", count)
				if count >= 3 {
					done <- true
					return
				}
			}
		}
	}()

	<-done
	ticker2.Stop()
	fmt.Println("定期処理終了")

	// ========== 実用例：日付の範囲チェック ==========
	fmt.Println("\n=== 実用例：日付の範囲チェック ===")

	startDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)
	checkDate := time.Date(2024, 6, 15, 0, 0, 0, 0, time.UTC)

	if checkDate.After(startDate) && checkDate.Before(endDate) {
		fmt.Printf("%s は範囲内です\n", checkDate.Format("2006-01-02"))
	}

	// ========== 実用例：年齢計算 ==========
	fmt.Println("\n=== 実用例：年齢計算 ===")

	birthDate := time.Date(2000, 5, 15, 0, 0, 0, 0, time.UTC)
	age := calculateAge(birthDate)
	fmt.Printf("誕生日: %s, 年齢: %d歳\n", birthDate.Format("2006-01-02"), age)
}

// 処理時間を計測する関数
func measureTime(fn func()) {
	start := time.Now()
	fn()
	elapsed := time.Since(start)
	fmt.Printf("処理時間: %v\n", elapsed)
}

// 年齢を計算する関数
func calculateAge(birthDate time.Time) int {
	now := time.Now()
	age := now.Year() - birthDate.Year()

	// 誕生日がまだ来ていない場合は1歳引く
	if now.Month() < birthDate.Month() ||
		(now.Month() == birthDate.Month() && now.Day() < birthDate.Day()) {
		age--
	}

	return age
}

/*
【実行方法】
go run 04_time.go

【重要な概念】
1. time.Time は時刻を表す
2. time.Duration は期間を表す
3. タイムゾーンに注意
4. 参照日時は 2006-01-02 15:04:05

【時刻の取得】
time.Now()                    // 現在時刻
time.Date(y, m, d, h, m, s, ns, loc) // 特定の時刻
time.Unix(sec, nsec)          // Unix時間から

【フォーマット】
参照日時: Mon Jan 2 15:04:05 MST 2006
- 2006: 年
- 01: 月
- 02: 日
- 15: 時（24時間）
- 04: 分
- 05: 秒

【よく使うフォーマット】
"2006-01-02"              // YYYY-MM-DD
"2006-01-02 15:04:05"     // YYYY-MM-DD HH:MM:SS
"2006/01/02"              // YYYY/MM/DD
time.RFC3339              // 2006-01-02T15:04:05Z07:00

【時刻の操作】
t.Add(duration)           // 時刻を加算
t.AddDate(y, m, d)        // 年月日を加算
t.Sub(u)                  // 時刻の差分
t.Before(u)               // t < u
t.After(u)                // t > u
t.Equal(u)                // t == u

【Duration】
time.Nanosecond
time.Microsecond
time.Millisecond
time.Second
time.Minute
time.Hour

【Timer と Ticker】
time.NewTimer(d)          // 一度だけ実行
time.NewTicker(d)         // 定期的に実行
time.After(d)             // 簡易タイマー

【ベストプラクティス】
1. タイムゾーンを明示する
2. UTC で保存、表示時にローカル変換
3. Ticker は必ず Stop() する
4. 時刻の比較は Equal() を使う

【次のステップ】
05_regex.go で正規表現を学びましょう
*/
