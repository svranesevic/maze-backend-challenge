defmodule Maze.Application do
  # See https://hexdocs.pm/elixir/Application.html
  # for more information on OTP Applications
  @moduledoc false

  use Application

  def start(_type, _args) do
    port = String.to_integer(System.get_env("PORT") || "4040")

    children = [
      {Task.Supervisor, name: Maze.Server.TaskSupervisor},
      {Task, fn -> Maze.Server.accept(port) end}
    ]

    opts = [strategy: :one_for_one, name: Maze.Server.Supervisor]
    Supervisor.start_link(children, opts)
  end

end
