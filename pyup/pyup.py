import sys
def up(t):
    return t.upper()

if __name__ == "__main__":
    if len(sys.argv) != 1:
        print(up(sys.argv[1]))
    else:
        print("usage!")