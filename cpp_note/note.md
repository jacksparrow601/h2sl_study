# C++ Primer
## Chapter 2
### 2.2变量
#### 2.2.1 变量定义
何为对象
通常情况，对象是指一块能存储数据并具有某种类型的内存空间。
列表初始化如果存在初始值丢失信息的风险，会报错
```cpp
long double ld = 3.141592636;
int a{ld}, b = {ld}; // error，转换未执行，因为存在丢失信息风险
int c(ld), d = ld; // right
```
#### 2.2.2 变量声明和定义
c++将声明和定义区分开来，声明(declaration)使得名字为程序所知，如果某个文件想要使用这个变量，则必须包括声明。定义(definition)创建与名字关联的实体。
如果想声明一个变量而非定义它，在变量名之前添加关键字extern，**并且不要显式初始化变量**：
```cpp
extern int i; // 声明i而不是定义i
int j; //声明并且定义了j
```
任何包含了显式初始化的声明即为定义
```cpp
extern double pi = 3.1416;
```
变量只能被定义一次，但是可以声明多次
### 2.3 复合类型
复合类型是指基于其他类型定义的类型。引用和指针均属于符合类型。
一条声明语句由一个**基本数据类型**和紧随其后的一个**声明符**(declarator)列表组成。
对于较为简单的声明语句，例如int  i = 3等，声明符就是变量名，变量类型就是基本数据类型。但是还可以更复杂
#### 2.3.1 引用
引用一般指“左值引用”(lvalue reference)
引用(reference)是为对象起了另一个名字，引用不是对象，是为已经存在的对象起的另一个名字，引用类型引用(refers to)另外一种类型。通过将声明符写成&d的形式来定义引用类型，其中d是声明的变量名:
```cpp
int ival = 1024;
int &refVal = ival; //refVal指向ival，是ival的另一个名字
int &refVal2; // 报错，引用必须初始化
int &refVal4 = 10; // 错误：引用对象必须是一个对象
double dval = 3.14;
int &refVal5 = dval; // 错误：引用类型初始值必须为int类型对象
```
一般初始化变量，初始值会被拷贝到新建变量中去。但是定义引用时，程序把引用和它的初始值绑定在一起，一直绑定，不能更改，所以引用必须初始化，并且保证了“顶层const”不变。(所以对于引用不谈顶层const)
引用并非对象，它只是为一个已经存在的对象起的别名。

#### 2.3.2 指针
指针是指向另外一种类型的复合类型。
- 指针本身是一个对象，所以可以对指针进行赋值和拷贝操作
- 指针无需在定义时赋初值，和其他内置类型一样，在块作用域内定义但未初始化，将有一个不确定的值
**& 取地址符
```cpp
int val = 42;
int *p = &ival;
```
像&和*这样符号，既能作表达式里的运算符，也能作为声的一部分出现，符号的上下文决定了符号的意义：
如果单纯*的是解引用
```cpp
int i = 42;
int &r = i; // &紧随类型名出现，因此是声明的一部分，r是一个引用
int *p; // *紧随类型名出现，因此是声明的一部分，p是一个指针
p = &i; // &出现在表达式中，是一个取地址符
*p = i; // *出现在表达式中，是一个解引用
int &r2 = *p; // &是声明的一部分，*是一个解引用符 
```
空指针不指向任何对象
```cpp
int *p1 = nullptr; // 注意这里是声明而不是解引用
int *p2 = 0; //和上面等价
int *p3 = NULL; // 和上面等价
```
指针使用建议要初始化，不然不知道它指向什么地方，访问会引起错误。
初始化为0是可以通过检测得知它未指向任何具体对象
引用和指针最大区别就是引用本身不是对象，定义完引用就不能将其绑定在其他对象上。
如果指针为0，在判断条件中为false
```cpp
int ival = 1032;
int *pi = 0; // 合法空指针
int *pi2 = &ival; // 合法指针，存放ival地址
if(pi) // false
 // ...
if(pi2) // true
 // ...
```
void\*指针可以存放任何对象的地址，但仅可以进行和其他指针比较，作为函数的输入输出、或者给其他void\*指针赋值，因为不确定该地址存放什么类型的对象，所以不能进行其他操作。
#### 2.3.2 理解复合类型的生命
变量 = 基本数据类型 + 声明符
在一条定义语句中，基本数据类型只有一个，但是声明符的形式可以不同。
```cpp
int i = 1024, *p = &i, &r = i; // int值，指针，引用
int* p1, p2; //指针，int值
```
因为引用本身不是一个对象，因此不能定义指向引用的指针。但是指针是对象，所以存在对指针的引用
```cpp
int i = 42;
int *p;
int *&r = p; // r是一个对指针p的引用
r = &i; // r解引用了一个指针，把&i赋值给r也就是把i的地址赋给p，p指向i
*r = 0; //解引用r得到i，也就是p指向的对象，将i的值改为0
```
**从右往左阅读r的定义，离变量r最近的符号对变量有最直接的影响，&r说明r是一个引用，\*&r说明r引用了一个指针，最后的int说明r引用了一个int指针
### 2.4 const限定符
const对象创建后其值不能更改，所以const对象必须初始化。
```cpp
const int i = get_size(); // 正确：运行时初始化
const int j = 42; // 正确：编译时初始化
const int k;  // 错误，未初始化
```
初始化只要不改变const对象，则是不是const无关紧要
```cpp
int i = 42;
const int ci = i; // 正确：i的值拷贝给了ci
int j = ci; // 正确：ci的值被拷贝给了j
```
希望在多个文件中共享使用同一个const量，只在一个文件中定义const，在其他多个文件中声明并使用它，对const变量的定义和使用都添加extern关键字
```cpp
// file_1.cc 定义并初始化了一个常量，该常量能被其他文件访问
extern const int bufSize = fcn();
// file_1.h头文件
extern const int bufSize; // 与file_1.cc中的bufSize是同一个
```
#### 2.4.1 const的引用
可以把引用绑定到const对象上，称为对常量的引用(const to reference)，但不能通过引用去修改所绑定的常量对象
```cpp
const int ci = 1024;
const int &r1 = ci; // 正确：引用及其对应的对象都是常量
r1 = 42; // 错误：r1是对常量的引用
int &r2 = ci; // 错误试图让一个非常量引用指向一个常量对象
```
常量引用一定是指引用的对象是常量，因为引用本身不是对象，并且引用不能更改绑定的对象，所以从这个角度理解，引用就是一个“常量”无需额外说明

例外：初始化常量引用时允许使用任意表达式，只要该结果能转换成引用类型，允许常量引用绑定非常量的对象、字面值，甚至是一般表达式:
```cpp
int i = 42;
const int &r1 = i; // 允许const int &绑定到普通int对象，但是不允许通过r1去修改i的值，因为r1是一个常量引用，常量引用仅对引用可参与的操作进行了限定
const int &r2 = 42; // 正确：r2是一个常量引用
const int &r3 = r1*2; // 正确：r3是一个常量引用
int &r4 = r1*2; // 错误：r4是一个普通的非常量引用
```
一个特例讨论
```cpp
double dval = 3.14;
const int &ri = dval;
```
编译器实际操作
```cpp
const int temp = dval; // 由双精度浮点数生成的一个临时的整型变量
const int &ri = temp; // 让ri绑定这个临时量
```
本意是希望ri去绑定dval，但实际绑定到temp上去了，该操作非法，如果不需要int和double转换倒是可以
#### 2.4.2 指针和const
与引用一样，指向常量的指针(pointer to const)不能用于改变其所指的对象的值。如果希望存放常量的地址，只能使用指向常量的指针:
```cpp
const double pi = 3.14; 
double *ptr = &pi; // 错误：ptr是一个普通指针
const double *cptr = &pi; // 正确：cptr可以指向一个双精度常量
*cptr = 42; // 错误：不能给*cptr赋值
```
允许一个指向常量的指针指向一个非常量对象，但是不能通过指向常量的指针来修改这个非常量
```cpp
double dval = 3.14;
cptr = &dval; // 正确：但是不能通过cptr改变dval的值
```
**const 指针**
指针是对象，所以允许把指针直接定义为常量，常量指针(const pointer)必须初始化，并且初始化完成其值不能改变，即这个指针指向固定地址。==把*放在const关键字之前说明指针是一个常量。==
```cpp
int errNumb = 0;
int *const curErr = &errNumb; // curErr将一直指向errNumb
const double pi = 3.14159;
const double *const pip = &pi // pip是一个指向常量对象的常量指针
```
同样从右向左读，curErr最近的是const，所以curErr是一个常量对象，声明符下一个是*，所以curErr是一个常量指针。最后声明语句的基本数据类型为int所以，常量指针指向int对象。
常量指针依然可以通过指针修改所值的对象，当然首先要求对象是可以修改的。
#### 2.4.3 顶层const
- 顶层const：指针本身是个常量，更一般地说顶层const表示任意对象是常量
- 底层const：指针所指对象是一个常量
```cpp
int i = 0;
int *const p1 = &i; // 不能改变p1的值，这是一个顶层const
const int ci = 42; // 不能改变ci的值，这是一个顶层const
const int *p2 = &ci; // 允许改变p2的值，这是一个底层const
const int *const p3 = p2; // 靠右的是顶层const，靠左的是底层const
const int &r= ci; // 声明引用的const都是底层const
```
执行拷贝对象的操作时，顶层const一般不会影响拷贝
```cpp
i = ci; // 正确：拷贝ci的值，ci是一个顶层，不影响
p2 = p3; // 正确： p2和p3指向的对象类型相同，p3顶层的const部分不影响
```
底层const的限制却不容忽视，拷入和拷出对象必须有相同的底层const资格，或者能够发生转换，一般来说非常量可以转换为常量，反之不行
```cpp
int *p = p3; //错误：p3包含底层const定义，而p没有
p2 = p3; // 正确：p2和p3都是底层const
p2 = &i; // 正确：int*能够转换成const int*
int &r = ci; // 错误：普通的int&不能绑定到int常量
const int &r2 = i; // 正确：const int& 可以绑定到一个普通的int
```
#### 2.4.4 constexpr和常量表达式
常量表达式(const expression)是指值不改变并且在编译过程中就能得到计算结果的表达式。字面值属于常量表达式，用常量表达式初始化的const对象也是。
```cpp
const int max_files = 20; // max_files是
const int limit = max_files + 1; // limit是
int staff_size = 27; // staff_size不是常量表达式，由于数据类型是int，而非const int
const int sz = get_size(); // sz不是常量表达式，sz的值要到运行的时候才能的到
```
**constexpr 变量**
c++11中可以把变量声明为constexpr类型以便由编译器来验证变量的值是否是一个常量表达式。声明为constexpr的量一定是一个常量，而且必须用常量表达式初始化
```cpp
constexpr int mf = 20; // 20是常量表达式
constexpr int limit = mf + 1; // mf+1是常量表达式
constexpr int sz = size(); //只有当size是一个constexpr函数时才是一条正确的声明语句
```
**指针和constexpr**
在constexpr声明中如果定义了一个指针，限定符constexpr仅对指针有效，与指针所指的对象无关：
```cpp
const int *p = nullptr; // p是一个指向整形常量的指针
constexpr int *q = nullptr; // q是一指向整数的常量指针
int j = 0;
const int i = 42;
constexpr const int *p = &i; // p是常量指针，指向整数常量i
conxtexpr int *p1 = &j; // p1是常指针，指向整数j
```
### 2.5 处理类型
#### 2.5.1 类型别名
类型别名(type alias)是一个名字，它是某种类型的同义词
**关键字typedef**
```cpp
typedef double wages; // wages是double的同义词
typedef wages base, *p; // base是double同义词，p是double* 同义词
```
**别名声明**
```cpp
using SI = Sales_item; // SI是Sales_item的同义词
```
**指针、常量和类型别名**
```cpp
typedef char *pstring;
const pstring cstr = 0; // cstr是指向char的常量指针
const pstring *ps; // ps是一个指针，它指向的对象是指向char的常量指针
```
声明语句的基本数据类型变成了const pstring，const是给定类型的修饰，pstring是指向char的指针，因此const pstring就是指向char的常量指针，而非指向常量字符的指针。
```cpp
const char *cstr = 0; // 是对const pstring cstr的错误理解
```
#### 2.5.2 auto类型说明符
c++11新标准引入了auto类型说明符
```cpp
auto i = 0, *p = &i; // 正确：i是整数、p是整形指针
```
auto一般会忽略顶层const，同时底层const则会保留下来，比如当初始值为一个指向常量的指针
```cpp
const int ci = i, &cr = ci;
auto b = ci; // b是一个整数(ci的顶层const属性被忽略掉了)
auto c = cr; // c是一个整数(cr是ci的别名，ci本身是一个顶层const)
auto d = &i; // d是一个整形指针指向i
auto e = &ci; //e是一个指向整数常量的指针(对常量取地址是一种底层const)
```
若希望推断出的auto是一个顶层const，则需要明确指出:
```cpp
const auto f = ci; // ci的推演类型是int，f则为const int
```
```cpp
auto &g = ci; // ci是一个常量，可以推导出底层const，g是一个整型常量引用，绑定到ci
auto &h = 42; // 错误：不能为非常量引用绑定字面值
const auto &j = 42; // 正确：可以为常量引用绑定字面值
```
&和*只从属于某个声明符，而非基本数据类型，因此初始值必须是同一种类型：
```cpp
auto k = ci, &l = i;
auto &m = ci, *p = &ci; // m是对整型常量的引用，p是指向整型常量的指针
//错误：i的类型是int，而&ci的类型是const int
auto &n = i, *p2 = &ci;
```
#### 2.5.3 decltype类型指示符
```cpp
decltype(f()) sum = x; // sum的类型只函数f的返回类型
```
decltype如果使用过的表达式是一个变量，则decltype返回该变量的类型(包括顶层const和引用在内)
```cpp
const int ci = 0, &cj = ci;
decltype(ci) x = 0; // x的类型是const int
decltype(cj) y = x; // y的类型是const int&（顶层const）,y绑定到x
decltype(cj) z; // 错误：z是一个引用，必须初始化
```
如果变量名加上一堆括号，则得到的类型与不加括号不同
```cpp
//decltype 的表达式如果是加上了括号的变量，结果将是引用
decltype((i)) d; //错误：d是int&，必须初始化
decltype(i) e; //正确：e是一个（未初始化的）int
```
### 2.6 自定义数据结构
#### 2.6.3 编写自己的头文件
**预处理器概述**
- \#define把一个名字设定为预处理变量
- \#ifdef当且仅当变量已定义时为真
- \#ifndef当且仅当变量未定义时为真
一旦检查结果为真，则执行后续操作直至\#endif为止，若为假则忽\#ifdef或者\#ifndef到\#endif之间的部分
```cpp
#ifndef SALES_DATA_H 
#define SALES_DATA_H
#include <string> 
struct Sales data {
    std::string bookNo; 
    unsigned units sold = 0; 
    double revenue = 0.0;
};
#endif
```
## Chapter 3 字符串、向量和数组
### 3.1 命名空间的using声明
有了using声明无须专门的前缀(形如命名空间::)也能使用所需的名字了。using namespace::name;
```cpp
using std::cin;
using std::cout;
using std::endl;
```
### 3.2 标准库类型string
#### 3.2.2 string对象上的操作
```cpp
is >> s; 从is中读取字符串赋给s，字符串以空白分割，返回is
getline(is, s); 从is中读取字符串赋给s，返回is
```
**读取数量未知的string对象**
```cpp
int main(){
    string word;
    while(cin >> word){//反复读取，直至到达文件末尾
        cout << word << endl;
    }
    return 0;
}
```
**使用getline读取一整行**
如果希望最终得到的字符串中保留输入时的空白符，应该使用getline函数代替原来的 >> 运算符。getline也会返回流参数，也能作为判断的条件
```cpp
int main(){
    string line;
    //每次读入一整行，直至到达文件末尾
    while(getline(cin, line)){
        cout << line << endl;
    }
    return 0;
}
```
**string::size_type类型**
size_type是一个无符号类型的值，足够存放下任何string对象的大小。所有存放string类的size函数返回值的变量，均为string::size_type类型。
```cpp
auto len = line.size(); // len是string::size_type类型
```
**字面值和string对象相加**
当把string对象和字符字面值及字符串字面值换在一条语句中使用时，必须确保每个加法运算符(+)两侧的运算对象至少有一个是string:
```cpp
string s1 = "hello", s2 = "world"; 
string s3 = s1 + ", " + s2 + '\n';
string s4 = s1 + ", "; // 正确：把一个string对象和一个字面值相加
string s5 = "hello" + ", "; // 错误：两个运算对象都不是string
//正确：每个加法运算符都有一个运算对象是string
string s6 = s1 + ", " + "world";
string s7 = "hello" + ", " + s2;// 错误：不能直接把字面值相加 
```
s6实质
```cpp
string s6 = (s1 + ", ") + "world";
```
#### 3.2.3 处理string对象中的字符
|cctype头文件中的函数||
|----|----|
|isalnum(c)|当c是字母或数字时为真|
|isaplha(c)|当c时字母时为真|
|iscntrl(c)|当c是控制字符时为真|
|isdigit(c)|当c是数字时为真|
|isxdigit(c)|当c是十六进制数字时为真|
|islower(c)|当c是小写字母时为真|
|isupper(c)|当c是大写字母时为真|
|ispunct(c)|当c时标点符号时为真|
|isspace(c)|当c时空白时为真（空格，换行，横向纵向制表符号）|
|tolower(c)|当c是大写字母时，输出对应的小写字母，否则原样输出c|
|toupper(c)|当c是小写字母时，输出对应的大写字母，否则原样输出c|

### 3.5 数组
#### 3.5.1 定义和初始化内置数组
##### 字符数组的特殊性
当使用字符串字面值对字符数组进行初始化的时候，字符串的结尾会多一个空字符
```cpp
char a1[] = {'C', '+', '+'};
char a2[] = {'C', '+', '+', '\0'};
char a3[] = "C++"; // 自动添加字符串结束的空字符
//size of a1: 3 size of a3: 4
const char a4[6] = "Daniel"; // 错误：没有康健可存放的空字符！
```
##### 理解复杂的数组声明
```cpp
int *ptrs[10]; // ptrs是含有10个整形指针的数组
int &refs[10] = /*?*/ ; // 错误：不存在引用的数组
int (*Parray)[10] = &arr; // Parray指向一个含有10个整数的数组
int (&arrRef)[10] = arr; // arrRef引用一个含有10个整数的数组
```
从括号内往括号外，从右向左。
对于ptrs从右向左，首先知道定义的时一个大小为10的数组，名字为ptrs，最后知道数组中存放的是指向int指针。
对于Parray，由于右括号存在，所以首先知道Parray是一个指针，接下来观察右边，可以知道指针指向的是有十个元素的数组，往左看知道元素类型为int。
```cpp
int *(&arry)[10] = ptrs; // arry是数组的引用，数组有10个指针
```
#### 3.5.2 访问数组元素
数组下标建议用size_t类型，是一种无符号类型，它足够大以便能表示内存中任意对象的大小。来自stddef.h库
#### 3.5.3 指针和数组
取地址符可以用于获取指向某个对象的指针，可作用于任何对象。
```cpp
string nums[] = {"one" , "two" , "three"}; // 数组的元素是string对象
string *p = &nums[0]; // p指向nums的第一个元素
//编译器一般会把数组名字自动替换为指向数组首元素的指针：
string *p2 = nums; // 等价于p2 = &nums[0]
```
使用auto推导
```cpp
auto p3 = nums; // 会推导出指针而非数组
int ia[] = {0 ,1, 2, 3, 4};
auto ia2(ia); // ia2是一个整型指针，指向ia第一个元素
ia2 = 42; // 错误： ia2是一个指针
//ia的操作过程实际如下一行代码
auto ia2(&ia[0]); // 显然ia2的类型是int*
```
使用decltype推导
```cpp
decltype(ia) ia3 = {0, 1, 2, 3, 4}; // ia3直接就是一个和ia size相同的数组
ia3 = p; // 错误：不能用整型指针给数组赋值
ia3[4] = i; // 正确：把i的值赋给ia3的一个元素
```
###### 数组指针也是迭代器
数组指针支持所有迭代器的能力
```cpp
int arr[] = {0, 1, 2, 3, 4}; 
int *p = arr; // p指向arr的第一个元素
++p; // p指向arr[]
```
##### 标准库函数begin和end
```cpp
int ia[] = {0, 1, 2, 3, 4}; 
int *beg = begin(ia); // 指向ia的首元素
int *last = end(ia); // 指向ia的尾元素的下一个位置的指针
```
##### 指针运算
```cpp
constexpr size_t = 5;
int arr[sz] = {1, 2, 3, 4, 5};
int *ip = arr; // 等价于int *ip = &arr[0];
int *ip2 = ip + 4; // ip2指向arr的尾元素arr[4]
int *p = arr + sz; // 正确：p指向数组的尾元素下一位置
int *p2 = arr + 10; // 错误： arr只有5个元素，p2的值未定义
```
和迭代器一样，两个指针相减的结果类型为一种名为ptrdiff_t的标准库类型，和size_t一样，ptrdiff_t也是一种定义在cstddef头文件中的机器相关的类型。
```cpp
auto n = end(arr) - begin(arr); // n的值是5，也就是arr中元素的数量，两个指针必须指向同一个数组
```
##### 下标和指针
```cpp
int *p = &ia[2]; // p指向索引为2的元素
int  k = p[-2]; // p[-2]是ia[0]所表示的元素
```
#### 3.5.4 C风格字符串
##### C标准库的String函数
c风格字符串习惯存放在字符数组中并以空字符(null teminated)，以空字符结束的意思是在字符串最后一个字符后面跟着一个空字符('\0')。
|C风格的字符串函数|cstring 是c语言头文件string.h的c++版本|
|----|----|
|strlen(p)|返回p的长度，空字符不计算在内，但是size of会计算|
|strcmp(p1, p2)|比较p1和p2相等性，p1==p2, 返回0， p1>p2，返回正值, p1 < p2 返回一个负值|
|strcat(p1, p2)|将p2附加给p1之后，返回p1|
|strcpy(p1, p2)|将p2拷贝给p1，返回p1|
传入此类函数的指针必须指向以空字符作为结束的数组
```cpp
char ca[] = {'C', '+', '+'};
cout << strlen(ca) <<endl; // 严重错误：ca没有以空字符结束
```
比较时不可像string一样使用
```cpp
string s1 = "A string example";
string s2 = "A different string";
if(s1 < s2) // 会去比较s1和s2

const char ca1[] = "A string example";
const char ca2[] = "A different string";
if(ca1 < ca2) // 为定义的：试图比较两个无关地址
//实际该使用strcmp
if(strcmp(ca1, ca2) < 0)
```
使用string要比C风格字符串更安全和高效
#### 3.5.5 与旧代码的接口
##### 混用string对象和C风格字符串
任何出现字符串字面值的地方，都可以用以空字符结束的字符串数组来代替
- 允许以空字符结束的字符数组来初始化string对象或为string对象赋值
- 在string对象的加法运算中允许以空字符结束的字符数组作为其中一个运算对象(不能两个运算对象都是)；在string对象的复合赋值运算中允许使用以空字符结束的字符数组作为右侧的运算对象。
但是对于需要C风格字符串，需要使用string的c_str成员函数
```cpp
char *str = s; // 错误： 不能用string对象初始化char*
const char *str = s.c_str(); // 正确
```
##### 使用数组初始化vector对象
```cpp
int int_arr[] = {0, 1, 2, 3, 4, 5};
vector<int> ivec(begin(int_arr), end(int_arr));
```
c++中应该尽量使用vector和迭代器，避免使用内置数组和指针，尽量使用string，避免使用C风格的基于数组的字符串。
### 3.6 多维数组
```cpp
int ia[3][4];
for(auto &row : ia){ // 这里的&是必须的，避免了数组被自动转换为指针，该指针式指向第一组的第一个元素
    for(auto col : row){
        cout << col << endl;
    }
}
```
## Chapter 4 表达式
表达式由一个或多个运算对象组成，对表达式求值将得到一个结果，字面值和变量是最简单的表达式。
### 4.1 基础
#### 4.1.1 基本概念
##### 左值和右值
左值(lvalue)可以位于赋值语句的左侧，右值(rvalue)不能。
- 取地址符作用于一个左值运算对象，返回一个指向该运算对象的指针，这个指针是一个右值
### 4.11 类型转换
#### 4.11.1 算数转换
##### 整型提升
对于char, signed char等类型来说，只要它们所有的值能存在int里面就会提升成int。较大char类型(wchar_t, char16_t, char32_t)提升成int, unsigned int, unsigned long等中最小的一种类型。前提是转换后的类型能容纳原类型所有可能的值。
##### 无符号类型的运算对象
首先执行整型提升，如果类型匹配，两个结果要么都是ID爱符号，要么都是无符号的，这类情况小类型的运算对象转换成较大类型。
如果一个运算对象是无符号类型，另外一个运算对象是带符号对象，只要无符号对象不小于带符号类型，那么有符号转换成无符号。
如果无符号类型小于有符号，并且有符号类型能够保存无符号类型的全部数据，这时候转换为有符号。否则依然转换为无符号
#### 4.11.3 显式转换
强制类型转换格式
```cpp
cast-name<type>(expression);
//type是目标类型，expression是待转换的值
```
##### static_cast
任何具有明确定义的类型转换，只要不包含底层const，都可以使用static_cast
```cpp
double slope = static_cast<double>(j) / i ;
```
当把一个较大的算数类型赋给较小的类型，可以使用该转换，忽略精度损失
```cpp
void *p = &d; // 正确：任何非常量对象的地址都能存入void*
//正确：将void*转换回初识的指针类型，前提是d之前就是double*类型，一旦不符合，将产生为定义后果
double *dp = static_cast<double*>(p);
```
##### const_cast
const_cast只能改变运算对象的底层const
```cpp
const char *pc;
char *p = const_cast<char*>(pc); // 正确：但是通过p写值是未定义的行为
```
将常量对象转换为非常量对象的行为，一般称为“去掉const性质(const away)"。一旦去掉某个对象的const性质，就可以通过该对象进行写操作，前提是对象本身不是一个常量。
```cpp
const char *cp;
//错误：static_cast不能转换掉const性质
char *q = static_cast<char*>cp;
static_cast<string>(cp); //正确：字符串字面值转换成string类型
const_cast<string>(cp); // 错误：const_cast只改变常量属性
```
##### reinterpret_cast
通常为运算对象的位模式提供较低位上的重新解释。

## Chapter 5 语句

### 5.5.3
## 5.6 try语句块和异常处理

## Chapter 9
### 9.1 顺序容器
|顺序容器|说明|
|----|----|
|vector|可变大小数组。 支持快速随机访问。在尾部之外的位置插入或删除元 素可能很慢|
|deque|双瑞队列。支持快速随机访问。在头尾位置插入/删除速度很快|
|list|双向链表。只支持双向顺序访问。在list 中任何位置进行插入/删除 操作速度都很快|
|forward_list|单向链表。只支持单向顺序访问。在链表任何位置进行插入/ 删除操作 速度都很快|
|array|固定大小数组。支持快速随机访问。不能添加或删除元素|
|string|与vector 相似的容器，但专门用于保存字符。随机访问快。在尾部插入/删除速度快|
###9.2容器库概览
某些class无默认构造函数，在构造这种容器时不能只传递给它一个元素数目的参数
```cpp
//假定noDefault是一个没有默认构造函数的类型
vector<noDefault> v1(10, init);//正确：提供了元素初始器
vector<noDefault>v2(10);//错误：必须提供一个元素初始化器
```
|操作、类型名|解释|
|----|----|
|size_type|无符号整数类型，足够保存此种容器最大类型最大可能容器的大小|
|C c(b, e);|构造c，将迭代器b和e指定范围内的元素拷贝到c(array不支持)|
|a.swap(b) , swap(a,b)|均为交换a,b中的元素|
|c.max_size()|c可保存的最大元素数目|
|c.begin(), c.end()|返回指向c的首元素和尾元素之后位置的迭代器|
|c.rbegin(), c.rend()|返回指向c的尾元素和首元素之前位置的迭代器|
rbegin自增也会到达rend,rbegin自增即可实现倒序。
```cpp
#include <vector>
#include <iostream>
using namespace std;
int main(){
    vector<int> num = {1, 2, 3, 4, 5};
    auto it = num.rend();
    cout <<"rend: " << *it << " rend--: " << *(--it) << endl;
    it = num.rbegin();
    cout << "rbegin: " << *it << " rbegin++: " << *(++it);
    return 0;
}
//结果：
// rend: 0 rend--: 1
// rbegin: 5 rbegin++: 4
```
####9.2.1 迭代器
迭代器范围(begin, end),左闭右开，实际为[begin, end)
####9.2.2 容器类型成员
迭代器写法:
```cpp
list<string>::iterator iter;
//容器名<元素>::iterator iter;
```
####9.2.4容器定义和初始化
对于array，如果要用另一个array去初始化需保证容器大小和元素类型都一致
```cpp
vector<const char*> articles = {"a", "an", "the"};
vector<string> words(articles);//错误：容器类型必须匹配
```
```cpp
#include <vector>
#include <iostream>
#include <string>
#include <list>
using namespace std;
int main(){
    vector<const char*> num = {"a", "an"};
    //报错：类型不匹配
    // vector<const char*> num1(num);
    //正确
    vector<string> num1(num.begin(), num.end());
    //甚至可以用来初始化其他容器
    list<string> lis(num.begin(), num.end());//注意end是在末元素之后，左闭右开涵盖了全部数组
    //报错，貌似string无法自动转换为const char*
    // vector<const char*> num2 (num1.begin(), num1.end());
    return 0;
}
```
array必须同时指定元素类型和容器大小
```cpp
array<int, 42>;//正确
array<int>;//错误
array<int, 10> ia3 = {42, 20};//前两个为42,20,后续均为0
```
数组不支持拷贝或者赋值，array可以，但要保证元素类型和array的size一样
####9.2.5 赋值和swap
```cpp
array<int, 10>a2 = {0};//初始化可以，不够的array会自己初始化为0
a2 = {0};//错误：不能将一个花括号列表赋予数组
```
assign可以用制定元素的拷贝替换左边容器的所有元素 
```cpp
a.assign(iter1, iter2);
```
```cpp
list<string> names;
vector<const char*> oldstyle;
names = oldstyle; // 错误：容器类型不匹配
//正确：可以将const char*转换为string
names.assign(oldstyle.cbegin(), oldstyle.cend());
```
第二种用法
```cpp
list<string> slist(1);//1个元素，为空string
slist.assign(10, "Hiya!"); // 10个元素，每个都是"Hiya"
```
swap对于array必须保证容器大小和元素类型一直，对于其他容器仅要求类型一致，swap速度远快于copy,array的swap会真正交换元素的值所以交换时间和array size成正比
```cpp
vector<string> svec1(10);
vector<string> svec2(24);
swap(svec1, svec2);//交换两个vec元素
```
####9.2.6
除forward_list不支持size，其余均支持size,max_size, empty
####9.2.7 关系运算符
容器比较必须保证元素类型和容器类型均相同
容器的比较实质上相当于比较字典序列
```cpp
vector<int> v1 = { 1, 3, 5, 7, 9, 12 };
vector<int> v2 = { 1, 3, 9 };
// v1 < v2, v1[2] < v2[2]
vector<int> v3 = { 1, 3, 5, 7 };
//v1 > v3, v1.size() > v3.size()
```
###9.3 顺序容器操作
####9.3.4 向顺序容器添加元素
|操作|说明|
|----|----|
|forward_list|有自己版本的insert和emplace，并且不支持push_back和emplace_back|
|c.push_front(t),c.emplace front(args)|在c的头部创建一个值为t或者是由args创建的元素,list,forward_list和deque支持|
|c.insert(p,t),c.emplace(p,args)|在迭代器p指向的元素之前创建一个值为t或由args创建的元素，返回指向新添加的元素的迭代器|
|c.insert(p, n, t )|在迭代器p指向的元素之前插入n个值为t的元素。返回新添加的第一个元素的迭代器，若n为0，则返回p|
|c.insert(p, b, e)|将迭代器b和e指定的范围内的元素插入到迭代器p指向的元素之前|
|c.insert(p, il)|il是一个花括号包围的元素列表|
```cpp
list<string> lst;
auto iter = list.begin();
while(cin >> word){//insert每次返回插入到第一个元素之前的元素就想弹雨iter每次都被重置为begin,所以能实现push_front()的效果
    iter = lst.insert(iter, word);
}
```
emplace家族：emplac_front, emplace_back, emplace,这些操作构造而不是拷贝元素,分别对应push_front, insert和push_back
调用push或者insert成员函数时，是将元素类型的对象传递给他们，这些对象被拷贝到容器中。当调用emplace时，则是将参数传递给元素类型的构造函数。emplace成员使用这些参数在容器管理的内存空间中直接构造元素。
```cpp
//Sales_data(const std::string &s, unsigned n, double p):bookNo (s), unints_sold(n), revenue(p*n){ }
//Sales_data(const std::string &s):bookNo (s) { }
//在c的末尾使用三个参数的sales_data构造函数构造一个Sales_data对象
c.emplace_back("978-0590353403", 25, 15.99);
//push_back没有接受三个参数的version
c.push_back("978-0590353403", 25, 15.99);
c.emplace_back();//虽然啥也没有，但是还是会调用默认构造函数
c.emplace(iter, "999-99999999");//使用Sales_data(string)
```


## Chapter ?
#### priority_queue
