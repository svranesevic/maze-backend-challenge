defmodule Maze.Application do
  # See https://hexdocs.pm/elixir/Application.html
  # for more information on OTP Applications
  @moduledoc false

  use Application
  use Agent

  def start(_type, _args) do
    port = String.to_integer(System.get_env("PORT") || "4040")

    children = [
      {Maze.Worker, port}
    ]

    Supervisor.start_link(children, strategy: :one_for_one)
  end

end
