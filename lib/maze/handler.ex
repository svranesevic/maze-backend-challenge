defmodule Maze.Handler do

  def start_link(ref, _socket, transport, opts) do
    pid = spawn_link(__MODULE__, :init, [ref, transport, opts])
    {:ok, pid}
  end

  def init(ref, transport, _opts) do
    {:ok, socket} = :ranch.handshake(ref)
    loop(socket, transport)
  end

  def loop(socket, transport) do
    case transport.recv(socket, 0, 5000) do
      {:ok, data} ->
        IO.inspect(data)
        transport.send(socket, "RECEIVED")
        loop(socket, transport)
      _ ->
        :ok = transport.close(socket)
    end
  end
end
