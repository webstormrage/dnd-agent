_G.core = {}

_G.core.filterBy = function(variantList, exceptionsDict)
    local resultList = {}
    for i, v in ipairs(variantList) do
        if exceptionsDict[v] ~= true then
            table.insert(resultList, v)
        end
    end
    return resultList
end

_G.core.omit = function(object, exceptionsDict)
    local omitted = {}
    for key, v in ipairs(object) do
        if exceptionsDict[key] ~= true then
            omitted[key] = v
        end
    end
    return omitted
end

_G.core.merge = function(target, source)
  for k, v in pairs(source) do
    target[k] = v
  end
end

_G.core.findIndex = function(list, value)
    local index = 0
    for k, v in pairs(list) do
        if v == value then
            index = k
        end
    end
    return index
end

_G.core.get = function(dict, field)
    if dict[field] == nil then
        dict[field] = {}
    end
    return dict[field]
end

_G.core.LANGS = {
    'language/abyssal',
    'language/celestial',
    'language/draconic',
    'language/deep-speech',
    'language/infernal',
    'language/primordial',
    'language/sylvan',
    'language/undercommon',
    'language/common',
    'language/dwarvish',
    'language/elvish',
    'language/giant',
    'language/gnomish',
    'language/goblin',
    'language/halfling',
    'language/orc'
}

_G.core.MUSIC = {
    'bagpipes',
    'drum',
    'dulcimer',
    'flute',
    'lute',
    'lyre',
    'horn',
    'pan-flute',
    'shawm',
    'viol'
}