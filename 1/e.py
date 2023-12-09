import sys

def see_tower(input_string):
  """Returns the number of visible trees in a see tower.

  Args:
    input_string: A string representing the see tower. Each line represents a row of
      trees, and the character '#' represents a tree.

  Returns:
    The number of visible trees in the see tower.
  """

  # Split the input string into a list of lines.
  lines = input_string.splitlines()

  # Count the number of visible trees.
  visible_tree_count = 0
  for i in range(len(lines)):
    for j in range(len(lines[i])):
      if lines[i][j] == '#':
        # Check if the tree is visible from the north.
        if i == 0 or lines[i - 1][j] < lines[i][j]:
          visible_tree_count += 1
        # Check if the tree is visible from the east.
        if j == len(lines[i]) - 1 or lines[i][j + 1] < lines[i][j]:
          visible_tree_count += 1
        # Check if the tree is visible from the south.
        if i == len(lines) - 1 or lines[i + 1][j] < lines[i][j]:
          visible_tree_count += 1
        # Check if the tree is visible from the west.
        if j == 0 or lines[i][j - 1] < lines[i][j]:
          visible_tree_count += 1

  return visible_tree_count


def main():
  # Read the input string from standard input.
  input_string = sys.stdin.read()

  # Calculate the number of visible trees.
  visible_tree_count = see_tower(input_string)

  # Print the number of visible trees to standard output.
  print(visible_tree_count)


if __name__ == '__main__':
  main()
