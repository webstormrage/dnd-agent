function unitDefinition(attributes, inventory, options)
    -- Владения
    if attributes['proficiency'] == nil then
        attributes['proficiency'] = {}
    end
    attributes['proficiencies']['simple-weapon'] = true
    attributes['proficiencies']['armor/light'] = true
    attributes['proficiencies']['armor/medium'] = true
    -- Бонус мастерства
    attributes['proficiency-bonus'] = math.max(attributes['proficiency-bonus'], 2)
    -- Хиты
    attributes['hit-points'] = math.max(1, math.min(options['hit-points'], 18))
    -- Раса
    attributes['race'] = 'human' -- TODO: сделать опцией
    -- Экипировка
    inventory['scimitar'] = (inventory['scimitar'] or 0) + 1
    inventory['light-crossbow'] = (inventory['light-crossbow'] or 0) + 1
    inventory['bolt'] = (inventory['bolt'] or 0) + 10
    inventory['armor/leather'] = (inventory['armor/leather'] or 0) + 1
end

function optionsDefinition(attributes, choices)
    table.insert(choices, {
        name='hit-points', --2d8+2
        type='int'
    })
end