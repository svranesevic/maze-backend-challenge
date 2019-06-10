defmodule Maze.Worker do
  use Agent
  def start_link(port) do
    opts = [port: port]
    {:ok, _} = :ranch.start_listener(:listener, :ranch_tcp, opts, Maze.Handler, [])
  end
end
