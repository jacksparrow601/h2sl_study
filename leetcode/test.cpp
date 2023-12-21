#include <vector>
#include <iostream>
#include <string>
#include <list>
#include <array>
using namespace std;
typedef char* pstring;
int main(){
    char a1[] = {'a', 'c', 'b'};
    char a3[] = "acb";
    auto a4 = a1;
    cout <<"size of a1: " <<  sizeof(a1) << " size of a3: " << sizeof(a3) << endl;
    cout << a3 <<endl << a4 << endl;
    int in1[] = {1,2};
    auto in2 = in1;
    cout << "in1 : " << *in1 << " in2 : " << in2 <<endl;
    cout << "size of in1: " << sizeof(in1) << endl;
    return 0;
}