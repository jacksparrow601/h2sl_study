## chapter 0
### 二叉树纲领篇
二叉树两种分类
1. 是否可以通过遍历一遍二叉树得到答案？如果可以，用一个traverse函数配合外部变量来实现，这成为遍历的思维模式。
2. 是否可以定义一个递归函数，通过子问题（子树）的答案推导出原问题的答案？如果可以，写出这个递归函数的定义，并充分利用这个函数丶返回值，这叫分解问题的思维模式

如果单独抽出一个二叉树节点，它需要做什么事情？需要在什么时候（前/中/后序位置）做？其他的节点不用你操心，递归函数会帮你在所有节点上执行相同的操作。

快排就是二叉树的前序遍历，归并排序就是二叉树的后续遍历
```cpp
// quick sort
void sort(int nums[], int lo, int hi) {
    /****** 前序遍历位置 ******/
    // 通过交换元素构建分界点 p
    int p = partition(nums, lo, hi);
    /************************/

    sort(nums, lo, p - 1);
    sort(nums, p + 1, hi);
}
```
先构造分界点（前序位置），然后去左右子树构造分界点
```cpp
// merge sort

// 定义：排序 nums[lo..hi]
void sort(vector<int>& nums, int lo, int hi) {
    int mid = (lo + hi) / 2;
    // 对 nums[lo..mid] 进行排序
    sort(nums, lo, mid);
    // 对 nums[mid+1..hi] 进行排序
    sort(nums, mid + 1, hi);

    /****** 在后序位置进行归并操作 ******/
    // 合并 nums[lo..mid] 和 nums[mid+1..hi]
    merge(nums, lo, mid, hi);
    /**********************************/
}

```
先对[lo, mid],[mid+1, hi]进行排序，在后续位置合并数组，分治算法也是同样的思想
