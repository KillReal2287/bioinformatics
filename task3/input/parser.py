import sys

target_line = 4
val_pos = 4


def parse(file_name):
  lines = f.readlines()
  splitted = lines[target_line].split(" ")
  val = splitted[val_pos]
  result = val[1:-1]
  if float(result) > 90.0:
    print("OK")
  else:
    print("Not OK")


if __name__ == '__main__':
    name = sys.argv[1]
    f = open(name, 'r')
    parse(f)

