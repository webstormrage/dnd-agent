generators = {}
core = _G.core or {}

generators['Unit.addBandit'] = function(args, state, stack)
    local attributes = core.get(state, 'attributes', args.unit.attributes)
    local inventory = core.get(state, 'inventory', args.unit.inventory)
    local proficiencies = core.get(attributes, 'proficiencies', attributes['proficiencies'])
    attributes['proficiencies'] = proficiencies

    attributes['proficiencies']['simple-weapon'] = true
    attributes['proficiencies']['armor/light'] = true
    attributes['proficiencies']['armor/medium'] = true
    -- Бонус мастерства
    attributes['proficiency-bonus'] = math.max(attributes['proficiency-bonus'], 2)
    -- Хиты
    attributes['hit-points'] = math.random(1, 8) + math.random(1, 8) + 2 -- TODO: сделать параметром
    -- Раса
    attributes['race'] = 'human' -- TODO: сделать параметром
    -- Экипировка
    inventory['scimitar'] = (inventory['scimitar'] or 0) + 1
    inventory['light-crossbow'] = (inventory['light-crossbow'] or 0) + 1
    inventory['bolt'] = (inventory['bolt'] or 0) + 10
    inventory['armor/leather'] = (inventory['armor/leather'] or 0) + 1

    stack.pop = {
        unit = {
            attributes=attributes,
            inventory=inventory
        }
    }
end