# leetcode笔记
## 暂时没用到的二级标题
### 暂时没用到的三级标题

##二分
#### [34.在排序数组中查找元素的第一个和最后一个位置](https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/description/)
二分法，优雅
下界函数始终找的是下界，所以第二次+1找到连续数组的最右边再往右偏移一位，获取需要的位数只要往左取一位即可
```cpp
class Solution {
public:
    int lower_bound1(vector<int>& nums, int target){
        int l = -1, r = nums.size();
        while(l+1 < r){
            int m = (l+r)>>1;
            if(nums[m] < target)
                l = m; 
            else
                r = m;        
        }
        return r;
    }
    vector<int> searchRange(vector<int>& nums, int target) {
        int n = nums.size();
        if(n < 1)
            return {-1, -1};
        int l = lower_bound1(nums, target);
        if(l >= n || nums[l] != target)
            return {-1, -1};
        int h = lower_bound1(nums, target+1)-1;
        return {l, h};
    }
};
```
1.dasda
## 二叉树
#### [145. 二叉树的后序遍历](https://leetcode.cn/problems/binary-tree-postorder-traversal/)
给你一棵二叉树的根节点 root ，返回其节点值的后序遍历 。
[前中后序全解析](https://leetcode.cn/problems/binary-tree-postorder-traversal/solution/bang-ni-dui-er-cha-shu-bu-zai-mi-mang-che-di-chi-t/)
```cpp
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    vector<int> postorderTraversal(TreeNode* root) {
        vector<int> res;//配合收藏夹内题解去使用
        stack<TreeNode*> stk;
        if(!root)
            return res;
        stk.emplace(root);
        while(!stk.empty()){//后序遍历即左右中，所以这里要按照
            auto node = stk.top();//中空右左去push
            if(node){
                stk.pop();//我们已经在初始，或者上一次循环push进了中节点，
                //这次循环还会push所以要先pop避免重复
                stk.emplace(node);
                stk.emplace(nullptr);//这个标记一定是放在中节点之前，因为每个节点都有当作中节点的
                // 时候，实质上对于递归也是每次只处理中节点，
                // 所以我们在遍历的时候在中节点前做标记即可
                if(node->right) stk.emplace(node->right);
                if(node->left) stk.emplace(node->left);
            }
            else{
                stk.pop();
                node = stk.top();
                stk.pop();
                res.push_back(node->val);
            }
        }
        return res;
    }
};
```
#### [96. 不同的二叉搜索树](https://leetcode.cn/problems/unique-binary-search-trees/) [题解](https://leetcode.cn/problems/unique-binary-search-trees/solution/hua-jie-suan-fa-96-bu-tong-de-er-cha-sou-suo-shu-b/)
#### 解题思路
- 假设n个节点存在二叉排序树的个数是$G(n)$，令 $f(i)$ 为以 i 为根的二叉搜索树的个数，则 $G(n) = f(1)+f(2)+...+f(n)$
- 当 i 为根节点时，其左子树节点个数为 i-1 个，右子树节点为 n-i，则 $f(i) = G(i-1)\times G(n-i)$
- 综上可得$G(n) = G(0)\times G(n-1)+G(1)\times G(n-2)...+G(n-1)\times G(0) $
```cpp
class Solution {
public:
    int numTrees(int n) {
        vector<int> g(n+1, 0);
        g[0] = 1;
        g[1] = 1;
        for(int t = 2; t <= n; t++){
            for(int i = 1; i <= t; i++)
            g[t] += g[i-1]*g[t-i];
        }
        return g[n];
    }
};
```
### [98. 验证二叉搜索树](https://leetcode.cn/problems/validate-binary-search-tree/)
- 中序遍历
```cpp
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    bool isValidBST(TreeNode* root) {
        stack<TreeNode*> stk;//中序遍历一定是单调递增，并且严格递增
        stk.emplace(root);
        long v_o = LONG_MIN, v_n = INT_MIN;
        while(!stk.empty()){
            root = stk.top();
            if(root){
                stk.pop();
                if(root->right)stk.emplace(root->right);
                stk.emplace(root);
                stk.emplace(nullptr);
                if(root->left)stk.emplace(root->left);
            }
            else{
                stk.pop();
                root = stk.top();
                stk.pop();
                v_n = root->val;
                if(v_o >= v_n)//这里包含=
                    return false;
                v_o = v_n;
            }
        }
        return true;
    }
};
```
- 递归
```cpp
class Solution {
public:
    bool check(TreeNode* root, long lower, long higher){
        if(!root)//二叉搜索树要确定左，右子树所有子节点都在一个范围内
            return true;
        if(root->val <= lower || root->val >= higher)return false;

        return check(root->left, lower, root->val)&&check(root->right, root->val, higher);
    }
    bool isValidBST(TreeNode* root) {
        return check(root, LONG_MIN, LONG_MAX);
    }
};
```

##动态规划
###股票问题
####[121. Best Time to Buy and Sell Stock](https://leetcode.cn/problems/best-time-to-hav-and-sell-stock/)
股票问题最重要的
```cpp
class Solution {
public:
    int maxProfit(vector<int>& prices) {
        int sell = 0, hav = INT_MIN;
        for(auto p:prices){
            sell = max(p+hav, sell);
            hav = max(0-p, hav);//实际相当于持有成本，找到最低的持有成本
        } 
        return sell;
    }
}
```
####[122. Best Time to Buy and Sell Stock II](https://leetcode.cn/problems/best-time-to-hav-and-sell-stock-ii/)
面对股票，最重要的就是要贪心。最终的结果一定是手里股票清空profit最大，但是因为知道每天股票的价格，并且通过了sell存储了昨天交易的最大利润，所以如果说昨天buy今天卖亏钱了，那实际上我们可以撤回交易，实现永远不亏，永远在上涨的时候卖钱
```cpp
class Solution {
public:
    int maxProfit(vector<int>& prices) {
        int sell = 0, hav = INT_MIN;
        for(auto p:prices){
            sell = max(p+hav, sell);
            hav = max(sell-p, hav);//把利润攒起来了
        } 
        return sell;
    }
}
```
[分享｜股票问题系列通解](https://leetcode.cn/circle/discuss/qiAgHn/)
[5行解决所有股票买卖问题](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/solutions/740557/5xing-jie-jue-suo-you-gu-piao-mai-mai-we-evro/)
##双指针

#### [75. 颜色分类](https://leetcode.cn/problems/sort-colors/description/)
```cpp
class Solution {
public://p0,p1分别表示0，1的结尾，遇到1就换到p1，再把p1++，因为p1是在0的后面所以不需要p0++，但是遇到0的时候要注意，如果此时p0<p1，很说明必然换的是1，所以在交换i和p0之后，要把1再换回p1.如果p0 !< p1,说明必然是等于，那就没必要换，因为此时还没有1被排列进来
    void sortColors(vector<int>& nums) {
        int n = nums.size();
        for(int i = 0, p0 = 0, p1 = 0; i < n; i++){
            if(nums[i] == 1){
                swap(nums[i], nums[p1++]);
            }
            else if(nums[i] == 0){
                swap(nums[i], nums[p0]);
                if(p0 < p1){//这步的含义，如果p0小于p1说明了其实刚刚swap的一定是1，所以要换回到1的尾部
                    swap(nums[i], nums[p1]);
                }
                p0++;
                p1++;
            }
        }
    }
};
```