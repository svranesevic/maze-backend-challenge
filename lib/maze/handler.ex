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
    transport.recv(socket, 0, 5000)
    case transport.setopts(socket, [{:active, :once}]) do
      {:ok, socket, data} ->
        #io:format("data received: ~p~n", [Data]);
        IO.inspect(data)
        loop(socket, transport)
      {:closed, socket} ->
        IO.puts("socket got closed!~n");
        loop(socket, transport)
      {:error, socket, reason} ->
        IO.puts("error happened: ~p~n", [Reason])
        #transport.send(socket, "RECEIVED")
      _ ->
        :ok = transport.close(socket)
    end
  end
end
