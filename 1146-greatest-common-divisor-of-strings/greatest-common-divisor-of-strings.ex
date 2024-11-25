defmodule Solution do
  @spec gcd_of_strings(str1 :: String.t, str2 :: String.t) :: String.t
  def gcd(a, b) when b != 0 do
    IO.puts("#{a}, #{b}")
    gcd(b, rem(a, b))
  end

  def gcd(a, _b), do: a

  def gcd_of_strings(str1, str2) do
    if str1 <> str2 != str2 <> str1 do
      ""
    else
      str1 |> String.slice(0, gcd(String.length(str1), String.length(str2)))
    end
  end
end