#include <fstream>
#include <iostream>
#include <string>

enum Indecies {
    FIRST_ELF_START,
    FIRST_ELF_SEPARATOR,
    FIRST_ELF_END,
    COMMA_SEPARATOR,
    SECOND_ELF_START,
    SECOND_ELF_SEPARATOR,
    SECOND_ELF_END,
};

using std::fstream;
using std::getline;
using std::string;

int solveProblem(string inputFilename);
bool lineHasOverlappingRanges(string line);
int *getElvesAssignments(string line);

int main(int argc, char **argv) {
    string filename = argv[1];
    int answer = solveProblem(filename);
    std::cout << answer << "\n";
    return 0;
}

int solveProblem(string inputFilename) {
    fstream file(inputFilename);
    string line;
    int sum = 0;
    while (getline(file, line)) {
        if (lineHasOverlappingRanges(line)) {
            sum++;
        }
    }
    return sum;
}

bool lineHasOverlappingRanges(string line) {
    int *assignments = getElvesAssignments(line);
    int firstElfStart = assignments[0];
    int firstElfEnd = assignments[1];
    int secondElfStart = assignments[2];
    int secondElfEnd = assignments[3];
    free(assignments);
    return (firstElfStart >= secondElfStart && firstElfStart <= secondElfEnd) ||
           (firstElfEnd >= secondElfStart && firstElfEnd <= secondElfEnd) ||
           (secondElfStart >= firstElfStart && secondElfStart <= firstElfEnd) ||
           (secondElfEnd >= firstElfStart && secondElfEnd <= firstElfEnd);
}

int *getElvesAssignments(string line) {
    int *assignments = (int *)malloc(4 * sizeof(int));
    string currentNumber = "";
    int currentIndex = 0;
    for (size_t i = 0; i < line.size(); i++) {
        char currentChar = line.at(i);
        if (currentChar == '-' || currentChar == ',') {
            assignments[currentIndex] = std::stoi(currentNumber);
            currentNumber = "";
            currentIndex++;
            continue;
        }
        currentNumber += currentChar;
    }
    assignments[currentIndex] = std::stoi(currentNumber);
    return assignments;
}
