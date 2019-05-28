defmodule Maze.User do
  use Ecto.Schema

  schema "user" do
    field :user_number, :string
    field :followers, {:array, :integer}
  end
end
