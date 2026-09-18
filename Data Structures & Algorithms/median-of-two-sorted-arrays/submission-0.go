func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    // 二分探索は短い配列側でやる
    // nums1 が短い方になるように入れ替える
    if len(nums1) > len(nums2) {
        nums1, nums2 = nums2, nums1
    }

    m, n := len(nums1), len(nums2)

    // 左側に置きたい個数
    // 例: 合計4個なら left 2個 / right 2個
    // 例: 合計5個なら left 2個 / right 3個
    half := (m + n) / 2

    // nums1 から左側に何個入れるかを二分探索する
    left, right := 0, m

    for left <= right {
        // i: nums1 から左側に入れる個数
        // j: nums2 から左側に入れる個数
        i := (left + right) / 2
        j := half - i

        // 範囲外対策
        // 左側が空なら -∞
        // 右側が空なら +∞
        const negInf = -1 << 60
        const posInf = 1 << 60

        nums1Left := negInf
        nums1Right := posInf
        nums2Left := negInf
        nums2Right := posInf

        // nums1 の切れ目の左側
        if i > 0 {
            nums1Left = nums1[i-1]
        }

        // nums1 の切れ目の右側
        if i < m {
            nums1Right = nums1[i]
        }

        // nums2 の切れ目の左側
        if j > 0 {
            nums2Left = nums2[j-1]
        }

        // nums2 の切れ目の右側
        if j < n {
            nums2Right = nums2[j]
        }

        // 正しい切れ目か確認する
        //
        // nums1: [ ... nums1Left ] | [ nums1Right ... ]
        // nums2: [ ... nums2Left ] | [ nums2Right ... ]
        //
        // 左側の値が、右側の値以下ならOK
        if nums1Left <= nums2Right && nums2Left <= nums1Right {
            // 合計が奇数なら、右側の一番小さい値が中央値
            if (m+n)%2 == 1 {
                return float64(min(nums1Right, nums2Right))
            }

            // 合計が偶数なら、
            // 左側の最大値と右側の最小値の平均が中央値
            leftMax := max(nums1Left, nums2Left)
            rightMin := min(nums1Right, nums2Right)

            return float64(leftMax+rightMin) / 2.0
        }

        // nums1Left が大きすぎる
        // nums1 から左側に入れすぎている
        if nums1Left > nums2Right {
            right = i - 1
        } else {
            // nums1 から左側に入れる数が少なすぎる
            left = i + 1
        }
    }

    return 0.0
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}