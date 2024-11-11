#include <iostream>
#include <stack>
#include <string>
#include <map>
#include <cassert>

using namespace std;

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


int shuntingYardAlgorithm(string &expr){
    stack<int> operands;
    stack<char> operators;
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
                    operands.push(number);
                    number = 0;
                }

                while(!operators.empty() && operators.top() != '('){
                    assert(operands.size()>=2);       // if not pass, the expression is invalid
                    int operand1 = operands.top(); operands.pop();
                    int operand2 = operands.top(); operands.pop();
                    char op = operators.top(); operators.pop();
                    int temp = handle_operator(op, operand1, operand2);
                    operands.push(temp);
                }
                operators.pop();
            }
            else{
                if('0'<=expr[i-1] && expr[i-1]<='9'){
                    operands.push(number);
                    number = 0;
                }
                while(!operators.empty() && priority[operators.top()]>=priority[expr[i]]){
                    assert(operands.size()>=2);       // if not pass, the expression is invalid
                    int operand1 = operands.top(); operands.pop();
                    int operand2 = operands.top(); operands.pop();
                    char op = operators.top(); operators.pop();
                    int temp = handle_operator(op, operand1, operand2);
                    operands.push(temp);
                }
                operators.push(expr[i]);
            }
        }
    }
    if('0'<=expr.back() && expr.back()<='9')
        operands.push(number);

    while(!operators.empty()){
        assert(operands.size()>=2);       // if not pass, the expression is invalid
        int operand1 = operands.top(); operands.pop();
        int operand2 = operands.top(); operands.pop();
        char op = operators.top(); operators.pop();
        int temp = handle_operator(op, operand1, operand2);
        operands.push(temp);
    }
    assert(operands.size() == 1);
    return operands.top();
}

int main(){
    string expr("1+(2*((3-1)*2))");
    int result = shuntingYardAlgorithm(expr);
    cout<<result<<endl;
    return 0;
}