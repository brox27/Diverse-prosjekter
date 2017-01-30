from random import *
from enum import Enum
seed(None)

class Monsters:
    #def __init__(
    items = []
    hp = 1

class Beastman(Monsters):
    hp = randint(10,40)
    attack = 5

class Player:
    hp = 100
    attack = 12
    xp = 0
    level = 1
    next_level = level*1000
    def level_up(level):
        level +=1
        attack +=1

INIT = 1
BATTLE = 2
FIELD = 3
SAFE = 4

Spiller = Player
var = INIT
while True:
    if var == INIT:
        Spiller = Player
        var = SAFE
    if var == SAFE:
        print("du er trygg nå....")
        print("men hva vil du gjøre?")
        print("tast 1 for å gjøre noe dumt")
        while True:
            a = int(input("ditt valg:"))
            if a == 1:
                var = FIELD
                print("dumt valg...")
                break
            else:
                print("tast 1 for å gjøre noe dumt")
            
    if var == FIELD:
        attacked = randint(1,10000)
        if attacked == 12:
            var = BATTLE
    if var == BATTLE:
        enemy = Monsters
