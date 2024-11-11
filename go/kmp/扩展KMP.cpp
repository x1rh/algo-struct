#include <iostream>
#include <vector>
#include <string>
#include <cassert>
#include <algorithm>

using namespace  std;

// luogu p5410 
// p是模式串， s是文本串，用s的前缀和s的后缀的前缀匹配。 
// z[i]表示p[0:n-1]和p[i:n-1]的最长公共前缀的长度。
// 设p[l:r]为p[l:n-1]与p[0:n-1]的最长公共前缀，由此推广到i=l+1, z[i]的计算方法：
// (1)如果i<=r, 那么z[i] = min(z[i-l+1])
class ExtendKMP {
public:
    string p;
    vector<int> z, v;
    ExtendKMP(string& p):p(p) {
        get_z();
    }
    void get_z() {
        int n = p.size();
        z = vector<int>(n);
        z[0] = n;
        for(int i=1, l=0, r=0; i<n; i++) {
            if(i<=r) {
                z[i] = min(z[i-l], r-i+1);
            }
            while(i+z[i] < n && p[z[i]] == p[i+z[i]]) {
                z[i]++;
            }
            if(i+z[i]-1 > r) {
                l = i;
                r = i + z[i] - 1;
            }
        }
    }

    // 求p和s的后缀的LCP的长度
    // s[l:r]
    void match(string& s) {
        vector<int>(s.size()).swap(v);

        for(int i=0, l=0, r=-1; i<s.size(); i++){
            if(i<=r) {
                v[i] = min(z[i-l], r-i+1);
            }
            while(v[i]<p.size() && i+v[i]<s.size() && (p[v[i]] == s[i+v[i]])) {
                v[i]++;
            }
            if(i+v[i]-1 > r) {
                l = i;
                r = i + v[i] - 1;
            }
        }
    }
};

int main() {
    ios::sync_with_stdio(false);
    cin.tie(0);

    string a, b;
    cin>>a>>b;
    auto exkmp = ExtendKMP(b);
    exkmp.match(a);

    auto cal = [](vector<int>& v)  {
        long long x = 0;
        for(int i=0; i<v.size(); i++) {
            x ^= 1LL * (i + 1) * (v[i] + 1);
        }
        return x;
    };

    cout<<cal(exkmp.z)<<endl;
    cout<<cal(exkmp.v)<<endl;

    return 0;
}


