#include <fstream>
#include <iostream>
#include <string>
#include <unordered_map>
#include <vector>

using std::fstream;
using std::getline;
using std::string;
using std::unordered_map;

int solveProblem(string inputFilename);
int getPriorityForGroup(const string group[3]);
int getPriorityForChar(char compartmentChar);

int main(int argc, char **argv) {
    string filename = argv[1];
    int prioritySum = solveProblem(filename);
    std::cout << prioritySum << "\n";
    return 0;
}

int solveProblem(string inputFilename) {
    int prioritySum = 0;
    fstream file(inputFilename);
    string line;
    int lineNumber = 0;
    string group[3];
    while (getline(file, line)) {
        if (lineNumber == 3) {
            lineNumber = 0;
            prioritySum += getPriorityForGroup(group);
        }
        group[lineNumber] = line;
        lineNumber++;
    }
    if (lineNumber == 3) {
        lineNumber = 0;
        prioritySum += getPriorityForGroup(group);
    }
    return prioritySum;
}

int getPriorityForGroup(const string group[3]) {
    unordered_map<char, int> sack;
    unordered_map<char, int> currentLine;
    for (char i = 'a'; i <= 'z'; i++) {
        sack.insert({i, 0});
        currentLine.insert({i, 0});
    }
    for (char i = 'A'; i <= 'Z'; i++) {
        sack.insert({i, 0});
        currentLine.insert({i, 0});
    }
    for (int i = 0; i < 3; i++) {
        for (char i = 'a'; i <= 'z'; i++) {
            currentLine.at(i) = 0;
        }
        for (char i = 'A'; i <= 'Z'; i++) {
            currentLine.at(i) = 0;
        }
        for (int j = 0; j < group[i].size(); j++) {
            char currentItem = group[i].at(j);
            if (currentLine.at(currentItem) > 0) {
                continue;
            }
            if (sack.at(currentItem) >= 2) {
                return getPriorityForChar(currentItem);
            }
            sack.at(currentItem)++;
            currentLine.at(currentItem) = 1;
        }
    }
    return 0;
}

int getPriorityForChar(char compartmentChar) {
    if (compartmentChar >= 'a' && compartmentChar <= 'z') {
        return compartmentChar - 96;
    }
    return compartmentChar - 38;
}
