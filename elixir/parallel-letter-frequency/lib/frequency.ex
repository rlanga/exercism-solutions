defmodule Frequency do
  @doc """
  Count letter frequency in parallel.

  Returns a map of characters to frequencies.

  The number of worker processes to use can be set with 'workers'.
  """
  @spec frequency([String.t()], pos_integer) :: map
  def frequency(texts, workers) do
    options = [
      name: MySupervisor,
      strategy: :one_for_one
    ]
    DynamicSupervisor.start_link(options)

    Enum.each(1..workers, fn _ -> 
      DynamicSupervisor.start_child(MySupervisor, {Worker, self()}) end)

    listen(texts, Enum.count(texts))
  end

  defp listen(tasks, count, result \\ %{})
  defp listen([], 0, result), do: result
  defp listen([], count, result) do
    receive do
      {:tally, value} ->
        listen([], count - 1, Map.merge(result, value, fn _k, v1, v2 -> v1 + v2 end))
    end
  end
  defp listen(tasks, count, result) do
    receive do
      {:ready, pid} ->
        GenServer.cast(pid, {:count, self(), hd(tasks)})
        listen(tl(tasks), count, result)
    end
  end
end


defmodule Worker do
  use GenServer

  def start_link(initiating_process) do
    {_ , pid} = GenServer.start_link(__MODULE__, %{})
    send(initiating_process, {:ready, pid})
  end
  
  @impl true
  def init(state \\ %{}) do
    {:ok, state}
  end

  @impl true
  def handle_cast({:count, server_pid, text}, state) do
    result = Regex.scan(~r/[[:alpha:]-]{1}/u, text)
    |> Stream.map(fn [w|_] -> String.downcase(w) end)
    |> Enum.frequencies()

    send(server_pid, {:tally, result})
    send(server_pid, {:ready, self()}) # Worker is indicating it is ready to do more work
    {:noreply, state}
  end
end
