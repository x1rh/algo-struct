template<typename DataType, typename ResultType>
class BIT {
public:
    int n;
    vector<DataType> d;
    
    BIT(int n):n(n) {
        d.resize(n+1);
    }

    void add(int i, DataType x) {
        for(; i <= n; i += (i & (-i) )) {
            d[i] += x;
        }
    }

    ResultType query(int i) {
        ResultType res = 0;
        for(; i > 0; i -= (i & (-i))) {
            res += d[i];
        }
        return res;
    }
};



struct BIT{
    int n;
    vector<int> tree;
    BIT(int n):n(n+1), tree(vector<int>(n+1)){}
    
    void build(vector<int>& arr){
        for(int i=0; i<this->n; i++) {
            update(i+1, arr[i]);
        }
    }

    void update(int index, int x) {
        for(int i=index; i<=this->n; i+=(i & (-i))){
            tree[i] += x;
        }
    }

    int query(int index) {
        int res = 0;
        for(int i=index; i>0; i-=(i & (-i))) {
            res += tree[i];
        }
        return res;
    }
};

