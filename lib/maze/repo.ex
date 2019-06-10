defmodule Maze.Repo do
  use Ecto.Repo,
    otp_app: :maze,
    adapter: Ecto.Adapters.Postgres
end
#
