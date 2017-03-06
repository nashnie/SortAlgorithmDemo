selectsort = {}

function swap(data, a, b)
    local temp = data[a]
    data[a] = data[b]
    data[b] = temp
end

function selectsort.sort(data)
    for i=1, #(data) - 1 do
    for j=i+1, #(data) do
        if data[j] < data[i] then
            swap(data, i, j)
        end
    end
    end
    return data
end

return selectsort