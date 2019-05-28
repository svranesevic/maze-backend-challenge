defmodule MazeTest do
  use ExUnit.Case
  doctest Maze

  test "greets the world" do
    assert Maze.hello() == :world
  end
end
