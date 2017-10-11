#pragma once
#include "stdafx.h"
#include <iostream>
#include <queue> 
#include <list>
#include <limits>
#include <fstream>
#include <string>
using namespace std;

#define INDEF 0
#define START 1
#define STOP 2

#define ASTAR 1
#define BFS 2
#define DIJKSTRA 3

void ReadFile();
void DrawBoard();
void SearchAlgorithm();
void AddChildren(int index);
bool IsNodeOpen(int thatNode);
bool IsNodeClosed(int thatNode);
void Calculations(int child, int parent);
void Traceback();
void SetSymbolOpenAndClosed();

class Node {
public:
	int x;
	int y;
	
	int Gcost;
	int Fcost;
	int Hcost;
	
	char symbol;
	bool wall = false;
	int difficulty;
	int parrentPlace;

	vector< int > kids;

	int Analyze(char c, int i);
};