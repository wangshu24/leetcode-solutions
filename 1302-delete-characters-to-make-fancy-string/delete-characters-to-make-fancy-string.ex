defmodule Solution do
  @spec make_fancy_string(s :: String.t) :: String.t
  def make_fancy_string(s), do: do_make_fancy_string(s, nil, nil, <<>>)

  defp do_make_fancy_string(<<ch, rest::bytes>>, prev_ch, prev2_ch, acc) do
    cond do
      ch == prev_ch && ch == prev2_ch -> do_make_fancy_string(rest, prev_ch, prev2_ch, acc)
      true -> do_make_fancy_string(rest, ch, prev_ch, <<acc::bytes, ch>>)
    end
  end
  defp do_make_fancy_string(_, _, _, acc), do: acc
end