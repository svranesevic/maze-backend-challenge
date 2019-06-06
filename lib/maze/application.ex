defmodule Maze.Application do
  # See https://hexdocs.pm/elixir/Application.html
  # for more information on OTP Applications
  @moduledoc false

  use Application

  def start(_type, _args) do
    # List all child processes to be supervised
    port = String.to_integer(System.get_env("PORT") || raise "missing $PORT environment variable")


    children = [
      # Starts a worker by calling: Maze.Worker.start_link(arg)
      # {Maze.Worker, arg}
      {Maze.Repo},
      {Task.Supervisor, name: Maze.TaskSupervisor},
      Supervisor.child_spec({Task, fn -> Maze.accept(port) end}, restart: :permanent)
    ]

    # See https://hexdocs.pm/elixir/Supervisor.html
    # for other strategies and supported options
    opts = [strategy: :one_for_one, name: Maze.Supervisor]
    Supervisor.start_link(children, opts)
  end
end
