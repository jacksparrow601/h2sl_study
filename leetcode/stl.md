#C++总结
#### std::greater(),std::less(),自定义比较函数
结论：在排序和建堆两种场景有看似相反的结果
- 排序
  less<T>升序（从左到右遍历，数组元素从小到大）
  greater<T>降序（从左到右遍历，数组元素从大到小）
- 建堆
  less<T>降序（从左到右遍历，堆元素从大到小，同层之间随便）
  greater<T>降序（从左到右遍历，堆元素从小到大，同层之间随便）
原因，less的实质含义是让一个集合中的按照前一个元素小于后一个元素去排序，例如[a,b],a为前一个元素，在数组中很好理解；但是对于堆结构或者树结构，顺序是（子节点，父节点）所以如果有新元素插入，则父节点一定是更大的元素，最后的排序顺序就是小的位于子节点，大的元素位于父节点，最终表现为大根堆。
greater的实质是让集合按照前一个元素大于后一个元素去排序，所以导致了数组的降序和父节点永远小的小根堆。
[参考:C++：std::greater()、std::less()、自定义比较函数的规则](https://blog.csdn.net/sandalphon4869/article/details/105419706)
####Priority_queue优先队列
priority_queue//默认是大根堆存储
priority_queue<int, vector<int>,greater<int>>//小根堆
priority_queue<int, vector<int>, less<int>>//默认大根堆
重点：自定义比较函数，通用型
```cpp
auto cmp = [](int a, int b){return a < b;};//和默认的less函数有同样的效果
priority_queue<int, vector<int>, decltype(cmp)>pq(cmp);
```
#### lower_bound & upper_bound

|lower_bound:|returns an iterator to the first element not less than the given key|
|----|----|
|upper_bound:|returns an iterator to the first element greater than the given key|
####deque
deque全称double ended queue双端队列，元素可以从队头或者队尾出队和入队
用法
```cpp
//基础插入删除
mydeque.push_front(1); // 头部插入
mydeque.push_back(2);  // 尾部插入 

mydeque.pop_front(); // 头部删除
mydeque.pop_back();  // 尾部删除
//任意位置插入
mydeque.insert(mydeque.begin(), 0); // 在begin处插入0
mydeque.erase(mydeque.begin()+1);   // 删除第二个元素
//读元素
int first = mydeque[0]; 
int last = mydeque.back();
```
####substr用法
substr(I, j-i+1) 取s[i]-s[j]均为闭区间的这一段字符


```cpp
traversal(TreeNode* root){
    stack<TreeNode*> stk;
    stk.emplace(root);
    while(root || !stk.empty()){
        while(root){
            stk.push(root);
            root = root->left;
        }
        if()
    }
}
``` 