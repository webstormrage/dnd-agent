_G.core = {}

_G.core.filterBy = function(variantList, exceptionsDict)
    local resultList = {}
    for i, v in ipairs(variantList) do
        if exceptionsDict[v] ~= true then
            table.insert(resultList, v)
        end
    end
end
