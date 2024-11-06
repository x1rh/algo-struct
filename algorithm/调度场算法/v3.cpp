#include <iostream>
#include <stack>
#include <string>
#include <map>
#include <vector>

using namespace std;

struct Node{
    char op;
    int number;
    Node(char op, int number):op(op), number(number){}
};

map<char, int> priority{
    {'+', 1}, {'-', 1},
    {'*', 2}, {'/', 2},
    {'(', 0},
};

int handle_operator(char op, int operand1, int operand2){
    switch(op){
        case '+':
            return operand1 + operand2;
        case '-':
            return operand2 - operand1;
        case '*':
            return operand1 * operand2;
        case '/':
            return operand2 / operand1;      // maybe use double
    }
}


vector<Node> shuntingYardAlgorithm(string &expr){
    stack<char> operators;
    vector<Node> suffix;
    int number = 0;
    for(int i=0; i<expr.size(); ++i){
        if('0'<=expr[i] && expr[i]<='9'){
            number = number * 10 + expr[i] - '0';
        }
        else{
            if(expr[i] == '('){
                operators.push(expr[i]);
            }
            else if(expr[i] == ')'){
                if('0'<=expr[i-1] && expr[i-1]<='9'){
                    suffix.emplace_back(Node{0, number});
                    number = 0;
                }
                while(!operators.empty() && operators.top() != '('){
                    char op = operators.top(); operators.pop();
                    suffix.emplace_back(Node{op, 0});
                }
                operators.pop();
            }
            else{
                if('0'<=expr[i-1] && expr[i-1]<='9'){
                    suffix.emplace_back(Node{0, number});
                    number = 0;
                }
                while(!operators.empty() && priority[operators.top()]>=priority[expr[i]]){
                    char op = operators.top(); operators.pop();
                    operators.push(op);
                }
                operators.push(expr[i]);
            }
        }
    }
    if('0'<=expr.back() && expr.back()<='9'){
        suffix.emplace_back(Node{0, number});
    }

    while(!operators.empty()){
        char op = operators.top(); operators.pop();
        suffix.emplace_back(Node{op, 0});
    }
    return suffix;
}

int calculate(vector<Node> &suffix){
    stack<int> operands;
    for(auto el : suffix){
        if(el.op){
            int operand1 = operands.top(); operands.pop();
            int operand2 = operands.top(); operands.pop();
            int res = handle_operator(el.op, operand1, operand2);
            operands.push(res);
        }
        else{
            operands.push(el.number);
        }
    }
    return operands.top();
}

int main(){
    string expr("1+(2*((3-1)*2))");
    vector<Node> suffix = shuntingYardAlgorithm(expr);
    for(auto el : suffix){
        if(el.op) cout<<el.op<<" ";
        else cout<<el.number<<" ";
    }
    cout<<endl;
    int res = calculate(suffix);
    cout<<res<<endl;
    return 0;
}