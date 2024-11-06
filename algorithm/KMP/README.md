## 朴素的字符串匹配算法BF：
文本串`text[]`和模式`pattern[]`匹配，自然的想法是`O(n^2)`算法，拿`pattern[]`的字符从头到尾和`text[]`匹配，设`i`为`text[]`当前下标，`j`为`pattern[]`当前下标，每一次`text[i] != pattern[j]`，下标`j`将重新归`0`，`i`加一，然后重复。

```cpp
int Brute_Force(){
	int i=0, j=0;
	int m = strlen(text);
	int n = strlen(pattern);
	while(i<m && j<n){
		if(text[i] == pattern[j]){
			++i; ++j;
		}
		else{
			i -= (j-1); j = 0;
		}
	}
	//... 
}
```


## KMP算法：
观察上面的BF算法，不难发现每次匹配失败，`j`都要重置为`0`。KMP算法改进了BF算法，**使得 j 不需要每次都重置为 0 **.  

### kmp算法的思想：  
倘若 `text[]` 和 `pattern[]` 能局部匹配，那么当某次匹配失败时（设`text[]`此时下标为`i`, `pattern[]`下标为`j`），`pattern[j] != text[i]`， 那么我们可以知道，`text[i]`左边的若干个连续字符和`pattern[]`匹配，同时我们可以想象`text[]`保持不动，`pattern[]相对于text[]`右移若干个单位长度，即可以使得`pattern[]`和`text[]`重新匹配(若干个字符匹配)。  



完成上述操作，需要如下的几个条件的支持：  
（0）`设text[i]`, `pattern[j]`时出现匹配失败  

（1）`pattern[0, j) == text[ i-j, i)`     即失配位置的左边若干字符全匹配的。

（2）`pattern[0, t) == text[i-j, i)`      即pattern相对于text**右移了j-t个单位**后，使得`pattern[]`和`text[]`重新匹配。

（3）`pattern[j-t, j) == text[i-j, i)`     这点联系第一点不难得知。

（4）于是，对比（2），（3）我们得到：`pattern[0, t) == pattern[j-t, j)` 。即在`pattern[0, j)`中长度为t的真前缀，要和长度为t的后缀完全匹配。

（5）观察加想象，为了保证`pattern[]`右移`(j-t)`个长度不遗漏任何情况，**我们应该使得移动长度j-t尽可能的小，即t尽可能的大**。

（6）失配后的下一次匹配位置是`pattern[t]`

到目前为止，实现kmp算法还需要求解t的值。再观察以上得到的结论，我们发现，t值和text[]无关! 求t值现在转换成了求pattern[0, j)最长真前缀和最长真后缀，且真前缀和真后缀完全匹配(相同)。进一步，我们顺其自然的知道，可以先预处理出t的值，之后重复使用即可。我们将t的值写在一个next[]数组里。


ps: 以上是我17年写的对kmp的大致理解。next数组其实很多人都是闭区间，而上面是左闭右开



## 中场休息
现在看了[知乎专栏-KMP算法](https://zhuanlan.zhihu.com/p/105629613)，想用文章里的pmt表代替上文的next继续往下讲kmp。现在忘记上面详细的下标定义，只记住KMP从BF演化的过程以及节约时间的地方即可。






## pmt
pmt的定义为`pmt[i] = max{k | pattern[0, k-1]==pattern[i-k+1, i]}`，通俗的话讲就是遍历到下标i时，且，pattern前缀`pattern[0,i]`的真前缀等于真后缀时，它们的最长长度。  


pmt的第二层含义：相当于失配后，pattern指针回退后，重新获得匹配后，再后面的一位 （一定要理解这一点）  


为求pmt表，用了一种十分巧妙的设计：
想象有两个相同的pattern串（下文用P1、P2代替），将两个P串上下对齐，然后将下面的串右移一个单位长度。如：
```
abcd
 abcd
```
然后设两个指针i=1, j=0
i代表上面的P1串的指针，上面的P1串代表后缀；  
j代表下面的P2串的指针，下面的P2串代表前缀；  
然后可以发现，现在就是在做遍历到下标i时，前缀与后缀的匹配。（如i=1时，j=0，刚好相当于做ab的前缀b和后缀a的匹配）  

i需要明确一点，即i只做一次遍历，从0至pattern.size()-1
j的含义需要明确一点，那就是P[0, j-1]所代表的前缀，与P[x, i-1]表示的前缀匹配。即匹配了，j才加1。否则j进行回退操作。

回退操作，回退操作利用了已经求好的部分pmt表进行。我们在j的位置失配，意味着在j-1的位置是匹配的, j=pattern[j-1]（记住上文提到的pmt的第二层含义）

若全部失配，**j回退到下标-1**，然后j加1重回0（意味着P2重新从头开始匹配，也意味着遍历到i时，后缀与前缀没有匹配的可能。然后i加1）

详细实现看下面代码。
```cpp
string pattern;
cin>>pattern;
vector<int> pmt(pattern.size(), 0);
// pmt[0] = 0;  vector初始化是已经完成这个逻辑
for(int i=1, j=0; i<pattern.size(); ++i){
    while(j>=0 && pattern[i]!=pattern[j]){
        j = (j!=0)?pmt[j-1]:-1;
    }
    pmt[i] = ++j;
}
```


leetcode.28.实现strStr()

全部代码：
```cpp
class Solution {
public:
    int strStr(string haystack, string needle) {
        if(needle.size() == 0) return 0;
        vector<int> pmt(needle.size(), 0);
        for(int i=1, j=0; i<needle.size(); ++i){
            while(j>=0 && needle[i]!=needle[j]){
                j = (j!=0)?pmt[j-1]:-1;
            }
            pmt[i] = ++j;
        }
        for(int i=0, j=0; i<haystack.size(); ++i){
            while(j>=0 && haystack[i]!=needle[j]){
                j = (j!=0)?pmt[j-1]:-1;
            }
            ++j;
            if(j == needle.size()){
                return i-j+1;
            }
        }
        return -1;
    }
};


```



也如那篇专栏文章所讲，这只是MP，KMP还有Knuth的优化

## Knuth优化
请阅读[知乎专栏-KMP算法](https://zhuanlan.zhihu.com/p/105629613)。

pmt中存在可以优化的跳转，优化后pmt表不在满足其原定义，更改为next表


优化点在于

```
      i
abababd
abababc

      i
abababd
  abababc

      i
abababd
    abababc

      i
abababd
      abababc
```


在i时匹配失败后，跳转3次才能发现不能匹配，优化点在于连这3次也砍掉


好，现在暂时忘记上面的东西

优化的逻辑：
```
   i
abababc
  abac
   j
```
可见，遍历到当前i的位置b时，其下一个位置a也匹配。
按之前pmt的定义，pmt[i+1] = j+1+1

我们现在引入next表的概念，其定义已经与pmt表不同，其定义为pattern[]遍历到i失配时，最终跳转的位置（从而避免中间跳转）

此时思考一个问题
若前缀在i+1时失配，后缀此时指针为j+1, 则j=next[j+1-1]， 注意next的定义，则next[i+1] = next[j]

注意这个思路，这个思路是递归的，这意味着若在i处失配，则next[i] = next[j-1], 且下标小的是先生成了，则意味着next[i]最终指向的就是最后回退的地方（链式指向前，后面的都更新为最前面的那个值），而不需要多次回退


代码：
```cpp
next[0] = 0;
for(int i=1, j=0; i<pattern.size(); ++i){
    while(j>=0 && pattern[i]!=pattern[j]){
        j = (j!=0?)next[j-1]:-1;
    }
    next[i] = (j>=0&&pattern[i+1]==pattern[j+1])?next[j++]:++j;
}

```



ps：还是不太理解knuth的优化，待续


## kmp为什么主串s指针i不回退 
    |
ababababe
ababe


假设现在s和p出现了一次长度为n的匹配，如s[i:i+n] == p[j:j+n]， 假设在下一个字符失配，那么观察下一次匹配：  
    |
ababababe
..ababe

思考，s串中的下一个匹配位置到底在哪里？  

其实是相当于p的子串`abab`的真前缀 和 s的子串`abab`的真后缀做了一次最长匹配，得到2（即ab） 注意到这两个串是相等的（这很关键），可以用到next[]数组的定义求长度为4的p串 `abab`的当真前缀等于真后缀时的最大长度。  

这里便隐含了为什么主串i指针不会退的原因。

i指代的是当前迭代的主串的子串的右端点，而左端点是隐含移动的。  

