# Go 練習筆記格式規範

## main.go 註解格式

```go
package main

import "fmt"

// ===== 概念名稱 =====
// 一句話說明這個概念是什麼

func main() {
    // --- 小節標題 ---
    x := 42 // 變數宣告：使用 := 進行短變數宣告，自動推斷型別

    fmt.Println(x) // 印出結果
}
```

**規則：**
- `// ===== 標題 =====`：主要概念區塊
- `// --- 標題 ---`：小節分隔
- `// 說明`：單行說明，放在程式碼右側或上方
- 每個重要概念都要有一行中文解釋

---

## NOTE.md 格式

```markdown
# 主題名稱

## 概念說明

用 2-4 句話解釋這個概念是什麼、為什麼存在、使用場景。

## 語法

​```go
// 基本語法示意（不需要能執行，著重說明結構）
var x int = 10
x := 10 // 短變數宣告
​```

## 程式碼範例

​```go
package main

import "fmt"

func main() {
    // 完整可執行的範例
    fmt.Println("Hello")
}
​```

**說明：**
- 重點 1
- 重點 2

## 常見錯誤

| 錯誤 | 原因 | 正確寫法 |
|------|------|----------|
| `var x = ` | 缺少值 | `var x = 0` |

## 相關概念

- 連結到其他相關筆記或主題（選填）
```

---

## 範例：Slice 筆記

### main.go 範例

```go
package main

import "fmt"

// ===== Slice（切片）=====
// Slice 是對陣列的動態視圖，長度可變，是 Go 中最常用的集合型別

func main() {
    // --- 建立 Slice ---
    s := []int{1, 2, 3}     // 字面值建立，不指定長度
    s2 := make([]int, 3)    // make 建立，長度 3，值為零值

    // --- append 新增元素 ---
    s = append(s, 4, 5)     // 新增多個元素，超過容量時自動擴容

    // --- 切片操作 ---
    fmt.Println(s[1:3])     // 取索引 1~2，左閉右開

    // --- len 與 cap ---
    fmt.Println(len(s))     // 長度：目前元素數量
    fmt.Println(cap(s))     // 容量：底層陣列可容納的元素數量

    _ = s2
}
```

### NOTE.md 範例

```markdown
# Slice（切片）

## 概念說明

Slice 是對底層陣列的動態視圖，長度可變。與陣列不同，Slice 不需要在宣告時指定長度，
是 Go 中處理集合資料最常用的型別。多個 Slice 可以共用同一個底層陣列。

## 語法

​```go
var s []int               // 宣告（nil slice）
s := []int{1, 2, 3}      // 字面值
s := make([]int, len, cap) // make 建立
​```

## 程式碼範例

​```go
package main

import "fmt"

func main() {
    s := []int{2, 3, 5, 7, 11}
    s = append(s, 13)
    fmt.Println(s[1:4]) // [3 5 7]
    fmt.Printf("len=%d cap=%d\n", len(s), cap(s))
}
​```

**說明：**
- `append` 若超過 cap，會建立新的底層陣列並複製
- 切片 `s[low:high]` 左閉右開，`high` 預設為 `len(s)`

## 常見錯誤

| 錯誤 | 原因 | 正確寫法 |
|------|------|----------|
| 修改 slice 意外影響原陣列 | 共用底層陣列 | 用 `copy()` 建立獨立副本 |
| index out of range | 存取超過 len | 先確認 `len(s)` |

## 相關概念

- Array（Slice 的底層結構）
- Map（另一種集合型別）
```
