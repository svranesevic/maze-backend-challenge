defmodule Maze do
  @moduledoc """
  Documentation for Maze.
  """

  @doc """
  Hello world.

  ## Examples

      iex> Maze.hello()
      :world

  """
  use Application

  def start(_type, _args) do
    Maze.Supervisor.start_link(name: Maze.Supervisor)
  end

end
