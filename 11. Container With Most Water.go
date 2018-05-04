// https://leetcode.com/problems/container-with-most-water/description/
package leetcode

func maxArea(height []int) int {
	area := 0
	i := 0
	j := len(height) - 1
	for i < j {
		h := height[i]
		if h > height[j] {
			h = height[j]
		}
		a := h * (j - i)
		if area < a {
			area = a
		}
		for height[i] <= h && i < j {
			i++
		}
		for height[j] <= h && i < j {
			j--
		}
	}
	return area
}

// int maxArea(int* heights, int n) {
//     int water = 0, *i = heights, *j = i + n - 1;
//     while (i < j) {
//         int h = *i < *j ? *i : *j;
//         int w = (j - i) * h;
//         if (w > water) water = w;
//         while (*i <= h && i < j) i++;
//         while (*j <= h && i < j) j--;
//     }
//     return water;
// }
