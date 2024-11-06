class G {
public:
    struct node {
        int to, d;
        node(){}
        node(int to, int d):to(to), d(d){}
        bool operator< (const node& rhs) const {
            return this->d < rhs.d;
        }
    };

    int n; 
    const static int inf = 0x3f3f3f3f;
    vector<vector<node>> g;

    G(int n):n(n) {
        g = vector<vector<node>>(n, vector<node>());
    }

    int prim(int start) {
        int res = 0; 
        vector<int> d(n, inf);
        vector<bool> v(n, false);

        d[start] = 0; 
        for(int i=0; i<n; i++) {
        
            int to = -1;
            int max_ = inf;
            for(int j=0; j<n; j++) {
                if(!v[j] && d[j] < max_){
                    max_ = d[j];
                    to = j; 
                }
            }
            
            if(to==-1) {
                //return -1; // 未连通
            } 
            
            v[to] = true; 
            res += d[to];
            for(int j=0; j<g[to].size(); j++) {
                int u = g[to][j].to;
                int c = g[to][j].d;
                if(!v[u] && c < d[u]) {
                    d[u] = c;
                }
            }
        }

        return res;
    }
};