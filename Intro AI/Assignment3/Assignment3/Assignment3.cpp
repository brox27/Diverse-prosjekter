// Assignment3.cpp : Defines the entry point for the console application.
//

#include "stdafx.h"
#include "Assignment3.h"

using namespace std;
Node * Board;
int numNodes = 0;

list<int> OpenNodeList;
list<int> ClosedNodeList;


int StartPos, StopPos, CombinedWeight;
int SearchMethod = DIJKSTRA;


int main(){
	ReadFile();
    while(1){}
	return 1;
}


void ReadFile() {
	ifstream t3file("board-2-1.txt");
	char c;
	Board = new Node[999];

	if (t3file.is_open()) {
		while (t3file.get(c)) {
			int var = Board[numNodes].Analyze(c, numNodes);
			if (var == START) { StartPos = numNodes; };
			if (var == STOP) { StopPos = numNodes; };
			CombinedWeight += Board[numNodes].difficulty;
			numNodes++;
		}
		t3file.close();
		cout << "START!" << endl;
		DrawBoard();

		SearchAlgorithm();
	}
	else {
		printf("error! Couldn't open the file... \n");
	}
}

int Node::Analyze(char c, int i) {

	if (i == 0) {
		x = 1;
		y = 0;
	}
	else {
		x = Board[i - 1].x;
		x++;
		y = Board[i - 1].y;
	}

	parrentPlace = -1;
	difficulty = 1;
	Gcost = numeric_limits<int>::max();
	Fcost = numeric_limits<int>::max();
	Hcost = numeric_limits<int>::max();


	if (c == 'w') {
		symbol = 'w';
		difficulty = 100;
	}
	else if (c == 'm') {
		symbol = 'm';
		difficulty = 50;
	}
	else if (c == 'f') {
		symbol = 'f';
		difficulty = 10;
	}
	else if (c == 'g') {
		symbol = 'g';
		difficulty = 5;
	}
	else if (c == 'r') {
		symbol = 'r';
		difficulty = 1;
	}
	else if (c == 'A') {
		symbol = 'A';
		Gcost = 0;
		return START;
	}
	else if (c == 'B') {
		symbol = 'B';
		return STOP;
	}
	else if (c == '.') {
		symbol = '.';
	}
	else if (c == '#') {
		symbol = '#';
		wall = true;
	}
	else {
		symbol = '\n';
		wall = true;
		y++;
		x = 0;
	}
	return INDEF;
}

void AddChildren(int index) {

	for (int i = 0; i <= numNodes; i++) {
		if ((Board[i].x < Board[index].x + 2) && (Board[i].x > Board[index].x - 2)) {
			if ((Board[i].y < Board[index].y + 2) && (Board[i].y > Board[index].y - 2)) {
				if ((index != i && !Board[i].wall) && ((Board[i].x == Board[index].x) || (Board[i].y == Board[index].y))) {
					Board[index].kids.push_back(i);
				}
			}
		}
	}
}

bool compare_Fvalue3(const int& first, const int& second)
{

	if (Board[first].Fcost < Board[second].Fcost) {
		return true;
	}
	return false;
}

bool compare_Gvalue3(const int& first, const int& second)
{

	if (Board[first].Gcost < Board[second].Gcost) {
		return true;
	}
	return false;
}

bool IsNodeOpen(int thatNode) {
	if (OpenNodeList.empty()) { return false; }
	list<int>::iterator lit;
	for (lit = OpenNodeList.begin(); lit != OpenNodeList.end(); ++lit) {
		if (*lit == thatNode) { return true; }
	}
	return false;
}


bool IsNodeClosed(int thatNode) {
	if (ClosedNodeList.empty()) { return false; }
	list<int>::iterator it;
	for (it = ClosedNodeList.begin(); it != ClosedNodeList.end(); ++it) {
		if (*it == thatNode) { return true; }
	}
	return false;
}

void Calculations(int child, int parent) {
	int CurDad = Board[child].parrentPlace;
	if (Board[child].parrentPlace == -1 || (Board[parent].Hcost < Board[CurDad].Hcost)) {
		Board[child].parrentPlace = parent;
	}

	Board[child].Gcost = (Board[parent].Gcost + Board[child].difficulty);

	int manhattanDistance = ((abs(Board[StopPos].x - Board[child].x)) + (abs(Board[StopPos].y - Board[child].y)));
	Board[child].Hcost = manhattanDistance;

	Board[child].Fcost = Board[child].Gcost + Board[child].Hcost;	// lagt sammen
}

void SetSymbolOpenAndClosed() {
	list<int>::iterator lit;
	for (lit = ClosedNodeList.begin(); lit != ClosedNodeList.end(); ++lit) {
		Board[*lit].symbol = 'c';
	}

	list<int>::iterator it;
	for (it = OpenNodeList.begin(); it != OpenNodeList.end(); ++it) {
		Board[*it].symbol = 'o';
	}
}

void SearchAlgorithm() {
	bool finished = false;
	int curr = StartPos;


	while (!finished) {
		AddChildren(curr);

		int child, next, best;

		for (child = 0; child < Board[curr].kids.size(); child++) {
			if (!IsNodeClosed(Board[curr].kids[child])) {
				if (!IsNodeOpen(Board[curr].kids[child])) { OpenNodeList.push_back(Board[curr].kids[child]); }
				Calculations(Board[curr].kids[child], curr);
			}
		}

		if ((OpenNodeList.size() > 1) && SearchMethod == ASTAR) { OpenNodeList.sort(compare_Fvalue3); }
		if ((OpenNodeList.size() > 1) && SearchMethod == DIJKSTRA) { OpenNodeList.sort(compare_Gvalue3); }



		ClosedNodeList.push_back(curr);
		curr = OpenNodeList.front();
		OpenNodeList.pop_front();

		if (curr == StopPos) { finished = true; }
	}
	SetSymbolOpenAndClosed();
	Traceback();
}

void Traceback() {

	list<int> Path;
	int hest, solutonCost;
	solutonCost = 0;
	hest = Board[StopPos].parrentPlace;
	while (hest != StartPos) {
		solutonCost += Board[hest].difficulty;
		Board[hest].symbol = '@';
		char temp;
		Path.push_front(hest);
		hest = Board[hest].parrentPlace;
	}
	cout << endl << "Price of solution: " << solutonCost << endl;
	DrawBoard();
}


void DrawBoard() {
	int i;
	for (i = 0; i < numNodes; i++) {
		cout << Board[i].symbol;
	}
}