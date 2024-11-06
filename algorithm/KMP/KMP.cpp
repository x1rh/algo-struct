#include <iostream>
#include <vector>

using namespace  std;

class KMP {
public:
    string& s, p;      // s: text,  p: pattern
    vector<int> nxt;   
    KMP(string &s, string &p):s(s), p(p) {
        nxt = vector<int>(p.size());
        nxt[0] = -1;
        for(int i=1, j=-1; i<p.size(); i++) {
            while(j>=0 && p[i] != p[j+1]) {
                j = nxt[j];
            }
            if(p[i] == p[j+1]) {
                j++;
            }
            nxt[i] = j;
        }
    }

    // return all i which satisfy s[i:i+len(p)] == p[0:len(p)], 0-index 
    vector<int>* match() {
        vector<int>* res = new vector<int>();
        for(int i=0, j=-1; i<s.size(); i++) {
            while(j>=0 && s[i] != p[j+1]) {
                j = nxt[j];
            }
            if(s[i] == p[j+1]) {
                ++j;
            }
            if(j == p.size()-1) {
                res->push_back(i-p.size()+1);
            }
        }
        return res;
    }
};


int main() {
    string s = "leetcode guardian guard";
    string p1 = "leetcode";
    string p2 = "code";
    string p3 = "guard";

    auto print = [](vector<int>& v) {
        for(auto &x : v) {
            cout<<x<<" ";
        } cout<<endl;
    };
    auto kmp1 = KMP(s, p1);
    print(*kmp1.match());

    auto kmp2 = KMP(s, p2);
    print(*kmp2.match());

    auto kmp3 = KMP(s, p3);
    print(*kmp3.match());

    return 0;
}
