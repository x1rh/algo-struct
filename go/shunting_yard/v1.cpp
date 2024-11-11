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
    {'+', 1}, {'-', 1},
    {'*', 2}, {'/', 2},
    {'(', 0}, {'_', 3}            // 用_代表负号
};



bool isDigit(char ch){
    return '0'<=ch && ch<='9';
}

void prepare(string &expr){
    for(int i=0; i<=expr.size(); ++i){
        if(expr[i] == '-'){
            if(i==0){
                expr[i] = '_';
            }
            else if(expr[i-1]=='('){
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
        }
        else if(expr[i] == '('){
            operators.push('(');
        }
        else if(expr[i] == ')'){
            if(isDigit(expr[i-1])){
                suffix.emplace_back(Node{0, number});
                number = 0;
            }
            while(!operators.empty() && operators.top()!='('){
                suffix.emplace_back(Node{operators.top(), 0});
                operators.pop();
            }
            operators.pop();
        }
        else{
            if(isDigit(expr[i-1])){
                suffix.emplace_back(Node{0, number});
                number = 0;
            }
            while(!operators.empty() && priority[expr[i]]<=priority[operators.top()]){
                suffix.emplace_back(Node{operators.top(), 0});
                operators.pop();
            }
            operators.push(expr[i]);
        }
    }
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
                        res = operand2 - operand1;
                        break;
                    case '*':
                        res = operand1 * operand2;
                        break;
                    case '/':
                        res = operand2 / operand1;
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