defmodule Maze.Worker do
  use Agent
  def start_link(port) do
    opts = [port: port]
    {:ok, _} = :ranch.start_listener(:maze, 100, :ranch_tcp, opts, Maze.Handler, [])
  end
end
