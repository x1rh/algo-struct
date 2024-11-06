#include <iostream>
#include <string>
#include <vector>
#include <stack>
#include <map>
#include <cassert>

using namespace std;

struct Node{
    char op;
    int number;
    Node(char op, int number):op(op), number(number){}
};

map<char, int> priority{
        {'{', 0},
        {'[', 1},
        {'(', 2},
        {'+', 3}, {'-', 3},
        {'*', 4}, {'/', 4},
        {'_', 5}            // 用_代表负号
};



bool isDigit(char ch){
    return '0'<=ch && ch<='9';
}


// prepare 处理负数的负号
void prepare(string &expr){
    for(int i=0; i<=expr.size(); ++i){
        if(expr[i] == '-'){
            if(i==0){
                expr[i] = '_';
            }
            else if(expr[i-1]=='(' || expr[i-1]=='[' || expr[i-1]=='{'){
                expr[i] = '_';
            }
        }
    }
}

int cal(string &expr){
    prepare(expr);
    stack<char> operators;
    vector<Node> suffix;
    int number = 0;
    for(int i=0; i<expr.size(); ++i){
        if(isDigit(expr[i])){
            number = number * 10 + expr[i] - '0';
        } else if(expr[i]=='(' || expr[i]=='[' || expr[i]=='{'){
            operators.push(expr[i]);
        } else if(expr[i]==')' || expr[i]==']' || expr[i]=='}'){
            char target = ' ';
            if(expr[i] == ')'){
                target = '(';
            }
            else if(expr[i] == ']'){
                target = '[';
            }
            else if(expr[i] == '}'){
                target = '{';
            }

            // 右括号，i必定大于等于1；判断有数，则将数输出到suffix向量中
            if(isDigit(expr[i-1])){
                suffix.emplace_back(Node{0, number});
                number = 0;
            }

            // 将括号中的操作符全部输出到suffix向量中
            while(!operators.empty() && operators.top()!=target){
                suffix.emplace_back(Node{operators.top(), 0});
                operators.pop();
            }
            operators.pop();
        } else{
            // 遇到操作符，如果左边是操作数，将操作数输出到suffix向量中
            if(i>0 && isDigit(expr[i-1])){
                suffix.emplace_back(Node{0, number});
                number = 0;
            }

            // 将优先级高的操作符出栈，输出到suffix中
            while(!operators.empty() && priority[expr[i]]<=priority[operators.top()]){
                suffix.emplace_back(Node{operators.top(), 0});
                operators.pop();
            }
            operators.push(expr[i]);
        }
    }

    // 如果末尾是数字，那么把num输出到suffix向量
    if(isDigit(expr.back())){
        suffix.emplace_back(Node{0, number});
    }
    while(!operators.empty()){
        suffix.emplace_back(Node{operators.top(), 0});
        operators.pop();
    }
    

    stack<int> operands;
    for(auto el : suffix){
        if(el.op){
            if(el.op == '_'){
                int operand1 = operands.top(); operands.pop();
                operands.push(-operand1);
            }
            else{
                int operand1 = operands.top(); operands.pop();
                int operand2 = operands.top(); operands.pop();
                int res = 0;
                switch(el.op){
                    case '+':
                        res = operand1 + operand2;
                        break;
                    case '-':
                        res = operand2 - operand1;   // 注意减数和被减数 
                        break;
                    case '*':
                        res = operand1 * operand2;
                        break;
                    case '/':
                        res = operand2 / operand1;   // 注意除数和被除数
                        break;
                }
                operands.push(res);
            }
        }
        else{
            operands.push(el.number);
        }
    }
    assert(operands.size() == 1);
    return operands.top();
}

int main(){
    string expr;
    while(cin>>expr){
        cout<<cal(expr)<<endl;
        return 0;
    }
}
