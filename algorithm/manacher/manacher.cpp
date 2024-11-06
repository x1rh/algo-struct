#include <iostream>
#include <vector>
#include <sstream>
#include <string>
#include <cassert>
#include <algorithm>

using namespace  std;

// luogo p3805 
// origin为原串
// s为填充#后的字符串，填充后字符串的长度必定为奇数
// d[i] 表示以i为中心的最长回文串的半径，包括中心。
// s[l:r]是一个动态维护的回文串，考虑i在区间[l, r]内、外时的情形，可以快速计算出d[i]的值:
// (1) l<=i<=r 时， 求出i的对称点j=r-i+l, d[i] = min(d[j], r-i+1), 且还要考虑i的对称半径可能比这更大的情况。
// (2) i > r时，暴力枚举以i为中心的最长回文串。
// l和r的值根据d[i]进行维护
class Manacher {
public:
    int idx, maxLen;
    string s;
    vector<int> d;

    Manacher(string &origin){
        maxLen = idx = 0;
        ostringstream oss;
        oss<<"#";
        for(auto &x : origin){
            oss<<x<<"#";
        }
        s = oss.str();

        d = vector<int>(s.size(), 0);
        d[0] = 1;
        for(int i=1, l=0, r=0; i<s.size(); i++) {
            if(i<=r) {
                d[i] = min(d[r-i+l], r-i+1);
            }
            while(i-d[i]>=0 && i+d[i]<s.size() && (s[i-d[i]] == s[i+d[i]])) {
                ++d[i];
            }
            if(i+d[i]-1 > r) {    // i+d[i]-1代表了一个合法的回文串的右端点
                l = i-d[i]+1;
                r = i+d[i]-1;
            }
            if(d[i] > maxLen) {
                maxLen = d[i];
                idx = i;
            }
        }
    }
    string get() {
        ostringstream oss;
        for(int i=idx; i<idx+maxLen; i++) {
           if(s[i] != '#') {
               oss<<s[i];
           }
        }
        string l = oss.str();
        string r = oss.str();
        reverse(l.begin(), l.end());
        if(s[idx] == '#') {
            return l + r;
        } else {
           return l + r.substr(1);
       }
    }
};

int main() {
    ios::sync_with_stdio(false);
    cin.tie(0);

    string s;
    cin>>s;
    auto m = Manacher(s);
    cout<<m.get().size()<<endl;
    return 0;
}


