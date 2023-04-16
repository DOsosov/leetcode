package main

func maxArea(height []int) int {
    maxelement := 0
    for i := 0; i < len(height); i++ {
        if height[i] > maxelement {
            maxelement = height[i]
        }
    }
    maxlast := 0
    v := 0
    vmax := 0
    for i := 0; i < len(height); i++ {
        for j := len(height) - 1; j > i; j-- {
            vmax = maxelement * (j - i)
            if vmax < maxlast {
                i = i + 1
                j = len(height) - 1
            }
            if height[i] >= height[j] {
                v = height[j] * (j - i)
                if v >= maxlast {
                    maxlast = v
                }
            } else {
                v = height[i] * (j - i)
                if v >= maxlast {
                    maxlast = v
                }
            }
        }
    }
    return maxlast
}