#include <fstream>
#include <iostream>
#include <string>
#include <vector>

using std::fstream;
using std::getline;
using std::string;

const char PLAYER_ROCK = 'A';
const char PLAYER_PAPER = 'B';
const char PLAYER_SCISSORS = 'C';
const char YOU_LOSE = 'X';
const char YOU_DRAW = 'Y';
const char YOU_WIN = 'Z';
const int PLAYER_COL = 0;
const int YOUR_COL = 2;

enum Scores {
    ROCK = 1,
    PAPER = 2,
    SCISSORS = 3,
    LOST = 0,
    DRAW = 3,
    WON = 6,
};

int calculateScore(string filename);
int getScoreForLine(string line);
int getRoundOutcome(char you);
int getHandScore(char player, char you);

int main(int argc, char **argv) {
    string filename = argv[1];
    int score = calculateScore(filename);
    std::cout << score << "\n";
    return 0;
}

int calculateScore(string filename) {
    int score = 0;
    fstream file(filename);
    string line;
    while (getline(file, line)) {
        score += getScoreForLine(line);
    }
    return score;
}

int getScoreForLine(string line) {
    char player = line.at(PLAYER_COL);
    char you = line.at(YOUR_COL);
    return getRoundOutcome(you) + getHandScore(player, you);
}

int getRoundOutcome(char you) {
    switch (you) {
        case YOU_WIN:
            return WON;
        case YOU_DRAW:
            return DRAW;
        case YOU_LOSE:
            return LOST;
    }
    return 0;
}

int getHandScore(char player, char you) {
    switch (player) {
        case PLAYER_ROCK:
            if (you == YOU_WIN) {
                return PAPER;
            }
            if (you == YOU_LOSE) {
                return SCISSORS;
            }
            return ROCK;
        case PLAYER_PAPER:
            if (you == YOU_WIN) {
                return SCISSORS;
            }
            if (you == YOU_LOSE) {
                return ROCK;
            }
            return PAPER;
        case PLAYER_SCISSORS:
            if (you == YOU_WIN) {
                return ROCK;
            }
            if (you == YOU_LOSE) {
                return PAPER;
            }
            return SCISSORS;
    }
    return 0;
}
