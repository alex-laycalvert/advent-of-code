#include <fstream>
#include <iostream>
#include <string>
#include <vector>

using std::fstream;
using std::string;
using std::vector;

int getMostCalories(string filename, int numElves);

int main(int argc, char **argv) {
    int mostCalories = getMostCalories(argv[1], 3);
    std::cout << mostCalories << "\n";
    return 0;
}

int getMostCalories(string filename, int numElves) {
    vector<int> elvesWithMostCalories;
    int currentCalories = 0;

    fstream file(filename);
    string line;
    while (std::getline(file, line)) {
        if (line == "") {
            for (int i = 0; i < numElves; i++) {
            if (elvesWithMostCalories.size() < numElves) {
                elvesWithMostCalories.push_back(currentCalories);
                break;
            }
                if (currentCalories > elvesWithMostCalories.at(i)) {
                    elvesWithMostCalories.insert(
                        elvesWithMostCalories.begin() + i, currentCalories);
                    break;
                }
            }
            currentCalories = 0;
            continue;
        }
        int caloriesForLine = std::stoi(line);
        currentCalories += caloriesForLine;
    }

    int sum = 0;
    for (int i = 0; i < numElves; i++) {
        sum += elvesWithMostCalories.at(i);
    }

    return sum;
}
