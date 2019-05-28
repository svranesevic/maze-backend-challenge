defmodule Maze.Repo.Migrations.CreateUser do
  use Ecto.Migration

  def change do
    create table(:people) do
      add :user_number, :string
      add :followers, {:array, :integer}
      #add :age, :integer
    end
  end
end
