#include <iostream>
#include <unordered_set>
#include <vector>

using namespace std;

// 利用unsigned溢出时取模 
class StringHash {
public:
    string s;
    vector<unsigned long long> p;     // p[i] = P^i
    vector<unsigned long long> hash;  // hash[i] = hash(s[0:i-1])

    // default p = 131
    StringHash(std::string &s, int P = 131) {
        this->s = s;
        p    = vector<unsigned long long>(s.size() + 1);
        hash = vector<unsigned long long>(s.size() + 1);
        p[0] = 1, hash[0] = 0;
        for(size_t i=1; i<=s.size(); i++) {
            p[i] = p[i-1] * P;
            hash[i] = hash[i-1]*P + s[i-1];
        }
    }

    // 1-index, get(l, r) = hash(s[l-1:r-1])
    unsigned long long get(int l, int r) {
        return hash[r] - hash[l-1] * p[r-l+1];
    }

    // 1-index
    bool isSubstr(int l1, int r1, int l2, int r2) {
        return get(l1, r1) == get(l2, r2);
    }
};
