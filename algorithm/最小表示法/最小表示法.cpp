// luogu  p1368 

#include <iostream>
#include <vector>

using namespace  std;

void getMin(vector<int> &s) {
    int n = s.size();
    vector<int> ss(n * 2);
    for(int i=0; i<n; i++) {
        ss[i] = s[i];
        ss[i+s.size()] = s[i];
    }

    int i=0, j=1, k=0;
    while(i<n && j<n) {
        for(k=0; k<n && ss[i+k] == ss[j+k]; k++);
        if(ss[i+k] > ss[j+k]) {
            i = i+k+1;
        } else {
            j = j+k+1;
        }
        if(i == j) {
            j++;
        }
    }


    auto idx = min(i, j);
    for(int i=0; i<n; i++) {
        if(i) cout<<" ";
        cout<<ss[i+idx];
    }cout<<endl;
}


int main() {
    int n;
    while(cin>>n) {
        vector<int> v(n);
        for(int i=0; i<n; i++) {
            cin>>v[i];
        }
        getMin(v);
    }
    return 0;
}


